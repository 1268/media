package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/sofia/container"
   "41.neocities.org/sofia/pssh"
   "encoding/base64"
)

var widevine = [8]uint16{
   0xEDEF, 0x8BA9, 0x79D6, 0x4ACE, 0xA3C8, 0x27DC, 0xD51D, 0x21ED,
}

func (s *Stream) init_protect(data []byte) ([]byte, error) {
   var file container.File
   err := file.Read(data)
   if err != nil {
      return nil, err
   }
   if moov, ok := file.GetMoov(); ok {
      for _, pssh1 := range moov.Pssh {
         if pssh1.Widevine() {
            s.pssh = pssh1.Data
         }
         copy(pssh1.BoxHeader.Type[:], "free") // Firefox
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
