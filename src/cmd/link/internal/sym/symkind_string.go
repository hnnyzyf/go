// Code generated by "stringer -type=SymKind"; DO NOT EDIT.

package sym

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Sxxx-0]
	_ = x[STEXT-1]
	_ = x[SELFRXSECT-2]
	_ = x[STYPE-3]
	_ = x[SSTRING-4]
	_ = x[SGOSTRING-5]
	_ = x[SGOFUNC-6]
	_ = x[SGCBITS-7]
	_ = x[SRODATA-8]
	_ = x[SFUNCTAB-9]
	_ = x[SELFROSECT-10]
	_ = x[SMACHOPLT-11]
	_ = x[STYPERELRO-12]
	_ = x[SSTRINGRELRO-13]
	_ = x[SGOSTRINGRELRO-14]
	_ = x[SGOFUNCRELRO-15]
	_ = x[SGCBITSRELRO-16]
	_ = x[SRODATARELRO-17]
	_ = x[SFUNCTABRELRO-18]
	_ = x[STYPELINK-19]
	_ = x[SITABLINK-20]
	_ = x[SSYMTAB-21]
	_ = x[SPCLNTAB-22]
	_ = x[SFirstWritable-23]
	_ = x[SBUILDINFO-24]
	_ = x[SELFSECT-25]
	_ = x[SMACHO-26]
	_ = x[SMACHOGOT-27]
	_ = x[SWINDOWS-28]
	_ = x[SELFGOT-29]
	_ = x[SNOPTRDATA-30]
	_ = x[SINITARR-31]
	_ = x[SDATA-32]
	_ = x[SXCOFFTOC-33]
	_ = x[SBSS-34]
	_ = x[SNOPTRBSS-35]
	_ = x[STLSBSS-36]
	_ = x[SXREF-37]
	_ = x[SMACHOSYMSTR-38]
	_ = x[SMACHOSYMTAB-39]
	_ = x[SMACHOINDIRECTPLT-40]
	_ = x[SMACHOINDIRECTGOT-41]
	_ = x[SFILEPATH-42]
	_ = x[SCONST-43]
	_ = x[SDYNIMPORT-44]
	_ = x[SHOSTOBJ-45]
	_ = x[SDWARFSECT-46]
	_ = x[SDWARFINFO-47]
	_ = x[SDWARFRANGE-48]
	_ = x[SDWARFLOC-49]
	_ = x[SDWARFMISC-50]
	_ = x[SABIALIAS-51]
}

const _SymKind_name = "SxxxSTEXTSELFRXSECTSTYPESSTRINGSGOSTRINGSGOFUNCSGCBITSSRODATASFUNCTABSELFROSECTSMACHOPLTSTYPERELROSSTRINGRELROSGOSTRINGRELROSGOFUNCRELROSGCBITSRELROSRODATARELROSFUNCTABRELROSTYPELINKSITABLINKSSYMTABSPCLNTABSFirstWritableSBUILDINFOSELFSECTSMACHOSMACHOGOTSWINDOWSSELFGOTSNOPTRDATASINITARRSDATASXCOFFTOCSBSSSNOPTRBSSSTLSBSSSXREFSMACHOSYMSTRSMACHOSYMTABSMACHOINDIRECTPLTSMACHOINDIRECTGOTSFILEPATHSCONSTSDYNIMPORTSHOSTOBJSDWARFSECTSDWARFINFOSDWARFRANGESDWARFLOCSDWARFMISCSABIALIAS"

var _SymKind_index = [...]uint16{0, 4, 9, 19, 24, 31, 40, 47, 54, 61, 69, 79, 88, 98, 110, 124, 136, 148, 160, 173, 182, 191, 198, 206, 220, 230, 238, 244, 253, 261, 268, 278, 286, 291, 300, 304, 313, 320, 325, 337, 349, 366, 383, 392, 398, 408, 416, 426, 436, 447, 456, 466, 475}

func (i SymKind) String() string {
	if i >= SymKind(len(_SymKind_index)-1) {
		return "SymKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SymKind_name[_SymKind_index[i]:_SymKind_index[i+1]]
}
