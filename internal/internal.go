package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/sofia/container"
   "41.neocities.org/sofia/pssh"
   "41.neocities.org/sofia/sidx"
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "bytes"
   "encoding/base64"
   "errors"
   "io"
   "log/slog"
   "net/http"
   "net/url"
   "os"
   "slices"
   "strings"
)

func (s *Stream) segment_base(represent *dash.Representation, ext string) error {
   file, err := os.Create(ext)
   if err != nil {
      return err
   }
   defer file.Close()
   segment := represent.SegmentBase
   data, _ := segment.Initialization.Range.MarshalText()
   var req http.Request
   req.URL = represent.BaseUrl.Url
   req.Header = http.Header{}
   // need to use Set for lower case
   req.Header.Set("range", "bytes=" + string(data))
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusPartialContent {
      return errors.New(resp.Status)
   }
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   data, err = s.init_protect(data)
   if err != nil {
      return err
   }
   _, err = file.Write(data)
   if err != nil {
      return err
   }
   key, err := s.key()
   if err != nil {
      return err
   }
   references, err := write_sidx(&req, segment.IndexRange)
   if err != nil {
      return err
   }
   var meter text.ProgressMeter
   meter.Set(len(references))
   http.DefaultClient.Transport = nil
   for _, reference := range references {
      segment.IndexRange[0] = segment.IndexRange[1] + 1
      segment.IndexRange[1] += uint64(reference.Size())
      data, _ := segment.IndexRange.MarshalText()
      data, err = func() ([]byte, error) {
         req.Header.Set("range", "bytes=" + string(data))
         resp, err := http.DefaultClient.Do(&req)
         if err != nil {
            return nil, err
         }
         defer resp.Body.Close()
         if resp.StatusCode != http.StatusPartialContent {
            return nil, errors.New(resp.Status)
         }
         return io.ReadAll(meter.Reader(resp))
      }()
      if err != nil {
         return err
      }
      data, err = write_segment(data, key)
      if err != nil {
         return err
      }
      _, err = file.Write(data)
      if err != nil {
         return err
      }
   }
   return nil
}

func (s *Stream) segment_list(
   represent *dash.Representation, ext string,
) error {
   file, err := os.Create(ext)
   if err != nil {
      return err
   }
   defer file.Close()
   initial, err := represent.SegmentList.Initialization.SourceUrl.Url(represent)
   if err != nil {
      return err
   }
   resp, err := http.Get(initial.String())
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   data, err = s.init_protect(data)
   if err != nil {
      return err
   }
   _, err = file.Write(data)
   if err != nil {
      return err
   }
   key, err := s.key()
   if err != nil {
      return err
   }
   var meter text.ProgressMeter
   meter.Set(len(represent.SegmentList.SegmentUrl))
   http.DefaultClient.Transport = nil
   for _, segment := range represent.SegmentList.SegmentUrl {
      media, err := segment.Media.Url(represent)
      if err != nil {
         return err
      }
      data, err := get(media, &meter)
      if err != nil {
         return err
      }
      data, err = write_segment(data, key)
      if err != nil {
         return err
      }
      _, err = file.Write(data)
      if err != nil {
         return err
      }
   }
   return nil
}

func (s *Stream) segment_template(
   represent *dash.Representation, ext string,
) error {
   file, err := os.Create(ext)
   if err != nil {
      return err
   }
   defer file.Close()
   var initial *url.URL
   if i := represent.SegmentTemplate.Initialization; i != nil {
      initial, err = i.Url(represent)
      if err != nil {
         return err
      }
   }
   if initial != nil {
      resp, err := http.Get(initial.String())
      if err != nil {
         return err
      }
      defer resp.Body.Close()
      if resp.StatusCode != http.StatusOK {
         return errors.New(resp.Status)
      }
      data, err := io.ReadAll(resp.Body)
      if err != nil {
         return err
      }
      data, err = s.init_protect(data)
      if err != nil {
         return err
      }
      _, err = file.Write(data)
      if err != nil {
         return err
      }
   }
   key, err := s.key()
   if err != nil {
      return err
   }
   var segments []int
   for r := range represent.Representation() {
      segments = slices.AppendSeq(segments, r.Segment())
   }
   var meter text.ProgressMeter
   meter.Set(len(segments))
   http.DefaultClient.Transport = nil
   for _, segment := range segments {
      media, err := represent.SegmentTemplate.Media.Url(represent, segment)
      if err != nil {
         return err
      }
      data, err := get(media, &meter)
      if err != nil {
         return err
      }
      data, err = write_segment(data, key)
      if err != nil {
         return err
      }
      _, err = file.Write(data)
      if err != nil {
         return err
      }
   }
   return nil
}

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
type Stream struct {
   ClientId string
   PrivateKey string
   Wrapper widevine.Wrapper
   key_id []byte
   pssh []byte
}
func (s *Stream) Download(represent *dash.Representation) error {
   for _, protect := range represent.ContentProtection {
      if protect.SchemeIdUri.Widevine() {
         if len(protect.Pssh) >= 1 {
            var box pssh.Box
            n, err := box.BoxHeader.Decode(protect.Pssh)
            if err != nil {
               return err
            }
            err = box.Read(protect.Pssh[n:])
            if err != nil {
               return err
            }
            s.pssh = box.Data
            break
         }
      }
   }
   var ext string
   switch *represent.MimeType {
   case "audio/mp4":
      ext = ".m4a"
   case "text/vtt":
      ext = ".vtt"
   case "video/mp4":
      ext = ".m4v"
   }
   if represent.SegmentBase != nil {
      return s.segment_base(represent, ext)
   }
   if represent.SegmentList != nil {
      return s.segment_list(represent, ext)
   }
   return s.segment_template(represent, ext)
}

func get(address *url.URL, meter *text.ProgressMeter) ([]byte, error) {
   resp, err := http.Get(address.String())
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   return io.ReadAll(meter.Reader(resp))
}

func write_segment(data, key []byte) ([]byte, error) {
   if key == nil {
      return data, nil
   }
   var file container.File
   err := file.Read(data)
   if err != nil {
      return nil, err
   }
   track := file.Moof.Traf
   if senc := track.Senc; senc != nil {
      for i, data := range file.Mdat.Data(&track) {
         err = senc.Sample[i].DecryptCenc(data, key)
         if err != nil {
            return nil, err
         }
      }
   }
   return file.Append(nil)
}

func (s *Stream) init_protect(data []byte) ([]byte, error) {
   var file container.File
   err := file.Read(data)
   if err != nil {
      return nil, err
   }
   if moov, ok := file.GetMoov(); ok {
      for _, value := range moov.Pssh {
         if value.Widevine() {
            s.pssh = value.Data
         }
         copy(value.BoxHeader.Type[:], "free") // Firefox
      }
      description := moov.Trak.Mdia.Minf.Stbl.Stsd
      if sinf, ok := description.Sinf(); ok {
         s.key_id = sinf.Schi.Tenc.S.DefaultKid[:]
         // Firefox
         copy(sinf.BoxHeader.Type[:], "free")
         if sample, ok := description.SampleEntry(); ok {
            // Firefox
            copy(sample.BoxHeader.Type[:], sinf.Frma.DataFormat[:])
         }
      }
   }
   return file.Append(nil)
}

var Forward = []struct{
   Country string
   Ip string
}{
{"Argentina", "186.128.0.0"},
{"Australia", "1.128.0.0"},
{"Bolivia", "179.58.0.0"},
{"Brazil", "179.192.0.0"},
{"Canada", "99.224.0.0"},
{"Chile", "191.112.0.0"},
{"Colombia", "181.128.0.0"},
{"Costa Rica", "201.192.0.0"},
{"Denmark", "2.104.0.0"},
{"Ecuador", "186.68.0.0"},
{"Egypt", "197.32.0.0"},
{"Germany", "53.0.0.0"},
{"Guatemala", "190.56.0.0"},
{"India", "106.192.0.0"},
{"Indonesia", "39.192.0.0"},
{"Ireland", "87.32.0.0"},
{"Italy", "79.0.0.0"},
{"Latvia", "78.84.0.0"},
{"Malaysia", "175.136.0.0"},
{"Mexico", "189.128.0.0"},
{"Netherlands", "145.160.0.0"},
{"New Zealand", "49.224.0.0"},
{"Norway", "88.88.0.0"},
{"Peru", "190.232.0.0"},
{"Russia", "95.24.0.0"},
{"South Africa", "105.0.0.0"},
{"South Korea", "175.192.0.0"},
{"Spain", "88.0.0.0"},
{"Sweden", "78.64.0.0"},
{"Taiwan", "120.96.0.0"},
{"United Kingdom", "25.0.0.0"},
{"Venezuela", "190.72.0.0"},
}

func write_sidx(req *http.Request, index dash.Range) ([]sidx.Reference, error) {
   data, _ := index.MarshalText()
   req.Header.Set("range", "bytes=" + string(data))
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var file container.File
   err = file.Read(data)
   if err != nil {
      return nil, err
   }
   return file.Sidx.Reference, nil
}

func (s *Stream) key() ([]byte, error) {
   if s.key_id == nil {
      return nil, nil
   }
   private_key, err := os.ReadFile(s.PrivateKey)
   if err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(s.ClientId)
   if err != nil {
      return nil, err
   }
   if s.pssh == nil {
      var pssh widevine.PsshData
      pssh.KeyIds = [][]byte{s.key_id}
      s.pssh = pssh.Marshal()
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, s.pssh)
   if err != nil {
      return nil, err
   }
   data, err := module.RequestBody()
   if err != nil {
      return nil, err
   }
   data, err = s.Wrapper.Wrap(data)
   if err != nil {
      return nil, err
   }
   var body widevine.ResponseBody
   err = body.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   block, err := module.Block(body)
   if err != nil {
      return nil, err
   }
   containers := body.Container()
   for {
      container, ok := containers()
      if !ok {
         return nil, errors.New("ResponseBody.Container")
      }
      if bytes.Equal(container.Id(), s.key_id) {
         key := container.Key(block)
         slog.Info(
            "CDM",
            "PSSH", base64.StdEncoding.EncodeToString(s.pssh),
            "key", base64.StdEncoding.EncodeToString(key),
         )
         return key, nil
      }
   }
}
