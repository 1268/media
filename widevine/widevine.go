package widevine

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/protobuf"
   "bytes"
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "encoding/hex"
   "encoding/pem"
   "github.com/chmike/cmac-go"
   "io"
)

// some videos require key_id and content_id, so entire PSSH is needed
func New_Module(private_key, client_ID, pssh []byte) (*Module, error) {
   block, _ := pem.Decode(private_key)
   var (
      err error
      mod Module
   )
   mod.private_key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   mod.license_request = protobuf.Message{
      1: protobuf.Bytes(client_ID),
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Bytes(pssh[32:]),
         },
      },
   }.Marshal()
   return &mod, nil
}

var Client = http.Default_Client

func unpad(buf []byte) []byte {
   if len(buf) >= 1 {
      pad := buf[len(buf)-1]
      if len(buf) >= int(pad) {
         buf = buf[:len(buf)-int(pad)]
      }
   }
   return buf
}

type Container struct {
   Key []byte
   Type uint64
}

func (c Container) String() string {
   return hex.EncodeToString(c.Key)
}

type Containers []Container

func (c Containers) Content() *Container {
   for _, container := range c {
      if container.Type == 2 {
         return &container
      }
   }
   return nil
}

type Module struct {
   license_request []byte
   private_key *rsa.PrivateKey
}

func (m Module) Post(post Poster) (Containers, error) {
   signed_request, err := m.signed_request()
   if err != nil {
      return nil, err
   }
   body, err := post.Request_Body(signed_request)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", post.Request_URL(), bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   if head := post.Request_Header(); head != nil {
      req.Header = head
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   body, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   body, err = post.Response_Body(body)
   if err != nil {
      return nil, err
   }
   return m.signed_response(body)
}

func (m Module) signed_request() ([]byte, error) {
   hash := sha1.Sum(m.license_request)
   signature, err := rsa.SignPSS(
      no_operation{},
      m.private_key,
      crypto.SHA1,
      hash[:],
      &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
   )
   if err != nil {
      return nil, err
   }
   signed_request := protobuf.Message{
      2: protobuf.Bytes(m.license_request),
      3: protobuf.Bytes(signature),
   }
   return signed_request.Marshal(), nil
}

func (m Module) signed_response(response []byte) (Containers, error) {
   // key
   signed_response, err := protobuf.Unmarshal(response)
   if err != nil {
      return nil, err
   }
   raw_key, err := signed_response.Get_Bytes(4)
   if err != nil {
      return nil, err
   }
   session_key, err := rsa.DecryptOAEP(
      sha1.New(), nil, m.private_key, raw_key, nil,
   )
   if err != nil {
      return nil, err
   }
   // message
   var enc_key []byte
   enc_key = append(enc_key, 1)
   enc_key = append(enc_key, "ENCRYPTION"...)
   enc_key = append(enc_key, 0)
   enc_key = append(enc_key, m.license_request...)
   enc_key = append(enc_key, 0, 0, 0, 0x80)
   // CMAC
   key_CMAC, err := cmac.New(aes.NewCipher, session_key)
   if err != nil {
      return nil, err
   }
   key_CMAC.Write(enc_key)
   key_cipher, err := aes.NewCipher(key_CMAC.Sum(nil))
   if err != nil {
      return nil, err
   }
   var keys Containers
   // .Msg.Key
   for _, message := range signed_response.Get(2).Get_Messages(3) {
      var key Container
      iv, err := message.Get_Bytes(2)
      if err != nil {
         return nil, err
      }
      key.Key, err = message.Get_Bytes(3)
      if err != nil {
         return nil, err
      }
      key.Type, err = message.Get_Varint(4)
      if err != nil {
         return nil, err
      }
      cipher.NewCBCDecrypter(key_cipher, iv).CryptBlocks(key.Key, key.Key)
      key.Key = unpad(key.Key)
      keys = append(keys, key)
   }
   return keys, nil
}

type Poster interface {
   Request_URL() string
   Request_Header() http.Header
   Request_Body([]byte) ([]byte, error)
   Response_Body([]byte) ([]byte, error)
}

type no_operation struct{}

func (no_operation) Read(buf []byte) (int, error) {
   return len(buf), nil
}
