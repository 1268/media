package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/sofia/pssh"
   "encoding/base64"
)

// dashif.org/identifiers/content_protection
// func (b *Box) Widevine() bool {
//    return b.SystemId.String() == "edef8ba979d64acea3c827dcd51d21ed"
// }

func (s *Stream) Download(represent *dash.Representation) error {
   for _, p := range represent.ContentProtection {
      if p.SchemeIdUri == "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed" {
         if p.Pssh != "" {
            data, err := base64.StdEncoding.DecodeString(p.Pssh)
            if err != nil {
               return err
            }
            var box pssh.Box
            n, err := box.BoxHeader.Decode(data)
            if err != nil {
               return err
            }
            err = box.Read(data[n:])
            if err != nil {
               return err
            }
            s.pssh = box.Data
            // fallback to INIT
            break
         }
      }
   }
   ext, err := get_ext(represent)
   if err != nil {
      return err
   }
   if represent.SegmentBase != nil {
      return s.segment_base(represent, ext)
   }
   if represent.SegmentList != nil {
      return s.segment_list(represent, ext)
   }
   return s.segment_template(represent, ext)
}
