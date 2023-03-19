package widevine

import (
   "2a.pages.dev/rosso/protobuf"
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "github.com/chmike/cmac-go"
)

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
