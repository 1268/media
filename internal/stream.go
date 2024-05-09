package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "154.pages.dev/widevine"
   "crypto/tls"
   "encoding/hex"
   "errors"
   "io"
   "log/slog"
   "net/http"
   "net/url"
   "os"
   "slices"
   "strconv"
   "strings"
)

func (s Stream) key(protect protection) ([]byte, error) {
   if protect.key_id == nil {
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
   if protect.pssh == nil {
      protect.pssh = widevine.PSSH(protect.key_id, nil)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, protect.pssh)
   if err != nil {
      return nil, err
   }
   key, err := module.Key(s.Poster, protect.key_id)
   if err != nil {
      return nil, err
   }
   slog.Debug("CDM", "key", hex.EncodeToString(key))
   return key, nil
}

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
type Stream struct {
   ClientId string
   PrivateKey string
   Name encoding.Namer
   Poster widevine.Poster
}

func (s Stream) segment_template(
   ext, initial string, rep *dash.Representation,
) error {
   base, err := url.Parse(rep.GetAdaptationSet().GetPeriod().GetMpd().BaseURL)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("GET", initial, nil)
   if err != nil {
      return err
   }
   req.URL = base.ResolveReference(req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   file, err := func() (*os.File, error) {
      s, err := encoding.Name(s.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(encoding.Clean(s) + ext)
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   var protect protection
   err = protect.init(file, res.Body)
   if err != nil {
      return err
   }
   key, err := s.key(protect)
   if err != nil {
      return err
   }
   var meter log.ProgressMeter
   log.SetTransport(nil)
   defer log.Transport{}.Set()
   template, ok := rep.GetSegmentTemplate()
   if !ok {
      return errors.New("GetSegmentTemplate")
   }
   media, err := template.GetMedia(rep)
   if err != nil {
      return err
   }
   meter.Set(len(media))
   client := http.Client{ // github.com/golang/go/issues/18639
      Transport: &http.Transport{
         Proxy: http.ProxyFromEnvironment,
         TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
      },
   }
   for _, medium := range media {
      req.URL, err = base.Parse(medium)
      if err != nil {
         return err
      }
      err := func() error {
         res, err := client.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if res.StatusCode != http.StatusOK {
            var b strings.Builder
            res.Write(&b)
            return errors.New(b.String())
         }
         return write_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (s *Stream) DASH(req *http.Request) ([]*dash.Representation, error) {
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   switch res.Status {
   case "200 OK", "403 OK":
   default:
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var media dash.MPD
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   err = media.Unmarshal(text)
   if err != nil {
      return nil, err
   }
   if media.BaseURL == "" {
      media.BaseURL = res.Request.URL.String()
   }
   var reps []*dash.Representation
   for _, v := range media.Period {
      seconds, err := v.Seconds()
      if err != nil {
         return nil, err
      }
      for _, v := range v.AdaptationSet {
         for _, v := range v.Representation {
            if seconds > 9 {
               if _, ok := v.Ext(); ok {
                  reps = append(reps, v)
               }
            }
         }
      }
   }
   slices.SortFunc(reps, func(a, b *dash.Representation) int {
      return int(a.Bandwidth - b.Bandwidth)
   })
   return reps, nil
}

func (s Stream) Download(rep *dash.Representation) error {
   ext, ok := rep.Ext()
   if !ok {
      return errors.New("Ext")
   }
   if v, ok := rep.GetSegmentTemplate(); ok {
      if v, ok := v.GetInitialization(rep); ok {
         return s.segment_template(ext, v, rep)
      }
   }
   return s.segment_base(ext, *rep.BaseURL, rep)
}


func (s Stream) TimedText(url string) error {
   res, err := http.Get(url)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := func() (*os.File, error) {
      s, err := encoding.Name(s.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(encoding.Clean(s) + ".vtt")
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   _, err = file.ReadFrom(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (s Stream) segment_base(
   ext, base_url string, rep *dash.Representation,
) error {
   sb := rep.SegmentBase
   req, err := http.NewRequest("GET", base_url, nil)
   if err != nil {
      return err
   }
   req.Header.Set("Range", "bytes=" + string(sb.Initialization.Range))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := func() (*os.File, error) {
      s, err := encoding.Name(s.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(encoding.Clean(s) + ext)
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   var protect protection
   err = protect.init(file, res.Body)
   if err != nil {
      return err
   }
   key, err := s.key(protect)
   if err != nil {
      return err
   }
   references, err := write_sidx(base_url, sb.IndexRange)
   if err != nil {
      return err
   }
   var meter log.ProgressMeter
   meter.Set(len(references))
   var start uint64
   end, err := func() (uint64, error) {
      _, s, _ := sb.IndexRange.Cut()
      return strconv.ParseUint(s, 10, 64)
   }()
   if err != nil {
      return err
   }
   log.SetTransport(nil)
   defer log.Transport{}.Set()
   for _, reference := range references {
      start = end + 1
      end += uint64(reference.ReferencedSize())
      bytes := func() string {
         b := []byte("bytes=")
         b = strconv.AppendUint(b, start, 10)
         b = append(b, '-')
         b = strconv.AppendUint(b, end, 10)
         return string(b)
      }()
      err := func() error {
         req, err := http.NewRequest("GET", base_url, nil)
         if err != nil {
            return err
         }
         req.Header.Set("Range", bytes)
         res, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if res.StatusCode != http.StatusPartialContent {
            return errors.New(res.Status)
         }
         return write_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}
