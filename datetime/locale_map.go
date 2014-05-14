package main

import "dlib"

var (
	localDescMap = map[string]string{
		"aa_DJ.UTF-8":      dlib.Tr("Afar, Djibouti"),
		"aa_ER":            dlib.Tr("Afar, Eritrea"),
		"aa_ER@saaho":      dlib.Tr("Afar, Eritrea (Saaho)"),
		"aa_ET":            dlib.Tr("Afar, Ethiopia"),
		"af_ZA.UTF-8":      dlib.Tr("Afrikaans, South Africa"),
		"am_ET":            dlib.Tr("Amharic, Ethiopia"),
		"an_ES.UTF-8":      dlib.Tr("Aragonese, Spain"),
		"ar_AE.UTF-8":      dlib.Tr("Arabic, United Arab Emirates"),
		"ar_BH.UTF-8":      dlib.Tr("Arabic, Bahrain"),
		"ar_DZ.UTF-8":      dlib.Tr("Arabic, Algeria"),
		"ar_EG.UTF-8":      dlib.Tr("Arabic, Egypt"),
		"ar_IN":            dlib.Tr("Arabic, India"),
		"ar_IQ.UTF-8":      dlib.Tr("Arabic, Iraq"),
		"ar_JO.UTF-8":      dlib.Tr("Arabic, Jordan"),
		"ar_KW.UTF-8":      dlib.Tr("Arabic, Kuwait"),
		"ar_LB.UTF-8":      dlib.Tr("Arabic, Lebanon"),
		"ar_LY.UTF-8":      dlib.Tr("Arabic, Libya"),
		"ar_MA.UTF-8":      dlib.Tr("Arabic, Morocco"),
		"ar_OM.UTF-8":      dlib.Tr("Arabic, Oman"),
		"ar_QA.UTF-8":      dlib.Tr("Arabic, Qatar"),
		"ar_SA.UTF-8":      dlib.Tr("Arabic, Saudi Arabia"),
		"ar_SD.UTF-8":      dlib.Tr("Arabic, Sudan"),
		"ar_SY.UTF-8":      dlib.Tr("Arabic, Syria"),
		"ar_TN.UTF-8":      dlib.Tr("Arabic, Tunisia"),
		"ar_YE.UTF-8":      dlib.Tr("Arabic, Yemen"),
		"as_IN":            dlib.Tr("Assamese, India"),
		"ast_ES.UTF-8":     dlib.Tr("Asturian, Spain"),
		"az_AZ":            dlib.Tr("Azeri, Azerbaijan"),
		"be_BY.UTF-8":      dlib.Tr("Belarusian, Belarus"),
		"be_BY@latin":      dlib.Tr("Belarusian, Belarus"),
		"bem_ZM":           dlib.Tr("Bemba, Zambia"),
		"ber_DZ":           dlib.Tr("Amazigh, Algeria"),
		"ber_MA":           dlib.Tr("Amazigh, Morocco"),
		"bg_BG.UTF-8":      dlib.Tr("Bulgarian, Bulgaria"),
		"bho_IN":           dlib.Tr("Bhojpuri, India"),
		"bn_BD":            dlib.Tr("Bengali, Bangladesh"),
		"bn_IN":            dlib.Tr("Bengali, India"),
		"bo_CN":            dlib.Tr("Tibetan, China"),
		"bo_IN":            dlib.Tr("Tibetan, India"),
		"br_FR.UTF-8":      dlib.Tr("Breton, France"),
		"brx_IN":           dlib.Tr("Bodo, India"),
		"bs_BA.UTF-8":      dlib.Tr("Bosnian, Bosnia and Herzegowina"),
		"byn_ER":           dlib.Tr("Blin, Eritrea"),
		"ca_AD.UTF-8":      dlib.Tr("Catalan, Andorra "),
		"ca_ES.UTF-8":      dlib.Tr("Catalan, Spain"),
		"ca_FR.UTF-8":      dlib.Tr("Catalan, France "),
		"ca_IT.UTF-8":      dlib.Tr("Catalan, Italy"),
		"crh_UA":           dlib.Tr("Crimean Tatar, Ukraine"),
		"csb_PL":           dlib.Tr("Kashubian, Poland"),
		"cs_CZ.UTF-8":      dlib.Tr("Czech, Czech Republic"),
		"cv_RU":            dlib.Tr("Chuvash, Russia"),
		"cy_GB.UTF-8":      dlib.Tr("Welsh, United Kingdom"),
		"da_DK.UTF-8":      dlib.Tr("Danish, Denmark"),
		"de_AT.UTF-8":      dlib.Tr("German, Austria"),
		"de_BE.UTF-8":      dlib.Tr("German, Belgium"),
		"de_CH.UTF-8":      dlib.Tr("German, Switzerland"),
		"de_DE.UTF-8":      dlib.Tr("German, Germany"),
		"de_LI.UTF-8":      dlib.Tr("German, Liechtenstein"),
		"de_LU.UTF-8":      dlib.Tr("German, Luxemburg"),
		"dv_MV":            dlib.Tr("Dhivehi, Maldives"),
		"dz_BT":            dlib.Tr("Dzongkha, Bhutan"),
		"el_CY.UTF-8":      dlib.Tr("Greek, Cyprus"),
		"el_GR.UTF-8":      dlib.Tr("Greek, Greece"),
		"en_AG":            dlib.Tr("English, Antigua and Barbuda"),
		"en_AU.UTF-8":      dlib.Tr("English, Australia"),
		"en_BW.UTF-8":      dlib.Tr("English, Botswana"),
		"en_CA.UTF-8":      dlib.Tr("English, Canada"),
		"en_DK.UTF-8":      dlib.Tr("English, Denmark"),
		"en_GB.UTF-8":      dlib.Tr("English, United Kingdom"),
		"en_HK.UTF-8":      dlib.Tr("English, Hong Kong"),
		"en_IE.UTF-8":      dlib.Tr("English, Ireland"),
		"en_IN":            dlib.Tr("English, India"),
		"en_NG":            dlib.Tr("English, Nigeria"),
		"en_NZ.UTF-8":      dlib.Tr("English, New Zealand"),
		"en_PH.UTF-8":      dlib.Tr("English, Philippines"),
		"en_SG.UTF-8":      dlib.Tr("English, Singapore"),
		"en_US.UTF-8":      dlib.Tr("English, United States"),
		"en_ZA.UTF-8":      dlib.Tr("English, South Africa"),
		"en_ZM":            dlib.Tr("English, Zambia"),
		"en_ZW.UTF-8":      dlib.Tr("English, Zimbabwe"),
		"eo.UTF-8":         dlib.Tr("Esperanto"),
		"eo_US.UTF-8":      dlib.Tr("Esperanto, United States"),
		"es_AR.UTF-8":      dlib.Tr("Spanish, Argentina"),
		"es_BO.UTF-8":      dlib.Tr("Spanish, Bolivia"),
		"es_CL.UTF-8":      dlib.Tr("Spanish, Chile"),
		"es_CO.UTF-8":      dlib.Tr("Spanish, Colombia"),
		"es_CR.UTF-8":      dlib.Tr("Spanish, Costa Rica"),
		"es_CU":            dlib.Tr("Spanish, Cuba"),
		"es_DO.UTF-8":      dlib.Tr("Spanish, Dominican Republic"),
		"es_EC.UTF-8":      dlib.Tr("Spanish, Ecuador"),
		"es_ES.UTF-8":      dlib.Tr("Spanish, Spain"),
		"es_GT.UTF-8":      dlib.Tr("Spanish, Guatemala"),
		"es_HN.UTF-8":      dlib.Tr("Spanish, Honduras"),
		"es_MX.UTF-8":      dlib.Tr("Spanish, Mexico"),
		"es_NI.UTF-8":      dlib.Tr("Spanish, Nicaragua"),
		"es_PA.UTF-8":      dlib.Tr("Spanish, Panama"),
		"es_PE.UTF-8":      dlib.Tr("Spanish, Peru"),
		"es_PR.UTF-8":      dlib.Tr("Spanish, Puerto Rico"),
		"es_PY.UTF-8":      dlib.Tr("Spanish, Paraguay"),
		"es_SV.UTF-8":      dlib.Tr("Spanish, El Salvador"),
		"es_US.UTF-8":      dlib.Tr("Spanish, United States"),
		"es_UY.UTF-8":      dlib.Tr("Spanish, Uruguay"),
		"es_VE.UTF-8":      dlib.Tr("Spanish, Venezuela"),
		"et_EE.UTF-8":      dlib.Tr("Estonian, Estonia"),
		"eu_ES.UTF-8":      dlib.Tr("Basque, Spain"),
		"eu_FR.UTF-8":      dlib.Tr("Basque, France"),
		"fa_IR":            dlib.Tr("Persian, Iran"),
		"ff_SN":            dlib.Tr("Fulah, Senegal"),
		"fi_FI.UTF-8":      dlib.Tr("Finnish, Finland"),
		"fil_PH":           dlib.Tr("Filipino, Philippines"),
		"fo_FO.UTF-8":      dlib.Tr("Faroese, Faroe Islands"),
		"fr_BE.UTF-8":      dlib.Tr("French, Belgium"),
		"fr_CA.UTF-8":      dlib.Tr("French, Canada"),
		"fr_CH.UTF-8":      dlib.Tr("French, Switzerland"),
		"fr_FR.UTF-8":      dlib.Tr("French, France"),
		"fr_LU.UTF-8":      dlib.Tr("French, Luxemburg"),
		"fur_IT":           dlib.Tr("Furlan, Italy"),
		"fy_DE":            dlib.Tr("Frisian, Germany"),
		"fy_NL":            dlib.Tr("Frisian, the Netherlands"),
		"ga_IE.UTF-8":      dlib.Tr("Irish, Ireland"),
		"gd_GB.UTF-8":      dlib.Tr("Gaelic, United Kingdom"),
		"gez_ER@abegede":   dlib.Tr("Ge'ez, Eritrea (with Abegede collation)"),
		"gez_ER":           dlib.Tr("Ge'ez, Eritrea"),
		"gez_ET@abegede":   dlib.Tr("Ge'ez, Ethiopia (Abegede)"),
		"gez_ET":           dlib.Tr("Ge'ez, Ethiopia"),
		"gl_ES.UTF-8":      dlib.Tr("Galician, Spain"),
		"gu_IN":            dlib.Tr("Gujarati, India"),
		"gv_GB.UTF-8":      dlib.Tr("Manx, United Kingdom"),
		"ha_NG":            dlib.Tr("Hausa, Nigeria"),
		"he_IL.UTF-8":      dlib.Tr("Hebrew, Israel"),
		"hi_IN":            dlib.Tr("Hindi, India"),
		"hne_IN":           dlib.Tr("Chhattisgarhi, India"),
		"hr_HR.UTF-8":      dlib.Tr("Croatian, Croatia"),
		"hsb_DE.UTF-8":     dlib.Tr("Upper Sorbian, Germany"),
		"ht_HT":            dlib.Tr("Kreyol, Haiti"),
		"hu_HU.UTF-8":      dlib.Tr("Hungarian, Hungary"),
		"hy_AM":            dlib.Tr("Armenian, Armenia"),
		"ia":               dlib.Tr("Interlingua"),
		"id_ID.UTF-8":      dlib.Tr("Indonesian, Indonesia"),
		"ig_NG":            dlib.Tr("Igbo, Nigeria"),
		"ik_CA":            dlib.Tr("Inupiaq, Canada"),
		"is_IS.UTF-8":      dlib.Tr("Icelandic, Iceland"),
		"it_CH.UTF-8":      dlib.Tr("Italian, Switzerland"),
		"it_IT.UTF-8":      dlib.Tr("Italian, Italy"),
		"iu_CA":            dlib.Tr("Inuktitut, Canada"),
		"iw_IL.UTF-8":      dlib.Tr("Hebrew, Israel"),
		"ja_JP.UTF-8":      dlib.Tr("Japanese, Japan"),
		"ka_GE.UTF-8":      dlib.Tr("Georgian, Georgia"),
		"kk_KZ.UTF-8":      dlib.Tr("Kazakh, Kazakhstan"),
		"kl_GL.UTF-8":      dlib.Tr("Greenlandic, Greenland"),
		"km_KH":            dlib.Tr("Khmer, Cambodia"),
		"kn_IN":            dlib.Tr("Kannada, India"),
		"kok_IN":           dlib.Tr("Konkani, India"),
		"ko_KR.UTF-8":      dlib.Tr("Korean, Republic of Korea"),
		"ks_IN@devanagari": dlib.Tr("Kashmiri, India (devanagari)"),
		"ks_IN":            dlib.Tr("Kashmiri, India"),
		"ku_TR.UTF-8":      dlib.Tr("Kurdish, Turkey"),
		"kw_GB.UTF-8":      dlib.Tr("Cornish, United Kingdom"),
		"ky_KG":            dlib.Tr("Kyrgyz, Kyrgyzstan"),
		"lb_LU":            dlib.Tr("Luxembourgish, Luxembourg"),
		"lg_UG.UTF-8":      dlib.Tr("Luganda, Uganda"),
		"li_BE":            dlib.Tr("Limburgish, Belgium"),
		"lij_IT":           dlib.Tr("Ligurian, Italy"),
		"li_NL":            dlib.Tr("Limburgish, the Netherlands"),
		"lo_LA":            dlib.Tr("Lao, Laos"),
		"lt_LT.UTF-8":      dlib.Tr("Lithuanian, Lithuania"),
		"lv_LV.UTF-8":      dlib.Tr("Latvian, Latvia"),
		"mai_IN":           dlib.Tr("Maithili, India"),
		"mg_MG.UTF-8":      dlib.Tr("Malagasy, Madagascar"),
		"mhr_RU":           dlib.Tr("Mari, Russia"),
		"mi_NZ.UTF-8":      dlib.Tr("Maori, New Zealand"),
		"mk_MK.UTF-8":      dlib.Tr("Macedonian, Macedonia"),
		"ml_IN":            dlib.Tr("Malayalam, India"),
		"mn_MN":            dlib.Tr("Mongolian, Mongolia"),
		"mr_IN":            dlib.Tr("Marathi, India"),
		"ms_MY.UTF-8":      dlib.Tr("Malay, Malaysia"),
		"mt_MT.UTF-8":      dlib.Tr("Maltese, Malta"),
		"my_MM":            dlib.Tr("Burmese, Myanmar"),
		"nan_TW@latin":     dlib.Tr("Chinese, Taiwan (Minnan dialect, Latin alphabet)"),
		"nb_NO.UTF-8":      dlib.Tr("Norwegian Bokmal, Norway"),
		"nds_DE":           dlib.Tr("Low German, Germany"),
		"nds_NL":           dlib.Tr("Low German, the Netherlands"),
		"ne_NP":            dlib.Tr("Nepali, Nepal"),
		"nl_AW":            dlib.Tr("Dutch, Aruba"),
		"nl_BE.UTF-8":      dlib.Tr("Dutch, Belgium"),
		"nl_NL.UTF-8":      dlib.Tr("Dutch, the Netherlands"),
		"nn_NO.UTF-8":      dlib.Tr("Nynorsk, Norway"),
		"nr_ZA":            dlib.Tr("Southern Ndebele, South Africa"),
		"nso_ZA":           dlib.Tr("Northern Sotho, South Africa"),
		"oc_FR.UTF-8":      dlib.Tr("Occitan, France"),
		"om_ET":            dlib.Tr("Oromo, Ethiopia"),
		"om_KE.UTF-8":      dlib.Tr("Oromo, Kenya"),
		"or_IN":            dlib.Tr("Oriya, India"),
		"os_RU":            dlib.Tr("Ossetian, Russia"),
		"pa_IN":            dlib.Tr("Punjabi, India"),
		"pap_AN":           dlib.Tr("Papiamento, Antilles"),
		"pa_PK":            dlib.Tr("Punjabi, Pakistan"),
		"pl_PL.UTF-8":      dlib.Tr("Polish, Poland"),
		"pt_BR.UTF-8":      dlib.Tr("Portuguese, Brasil"),
		"pt_PT.UTF-8":      dlib.Tr("Portuguese, Portugal"),
		"ro_RO.UTF-8":      dlib.Tr("Romanian, Romania"),
		"ru_RU.UTF-8":      dlib.Tr("Russian, Russia"),
		"ru_UA.UTF-8":      dlib.Tr("Russian, Ukraine"),
		"rw_RW":            dlib.Tr("Kinyarwanda, Rwanda"),
		"sa_IN":            dlib.Tr("Sanskrit, India"),
		"sc_IT":            dlib.Tr("Sardinian, Italy"),
		"sd_IN@devanagari": dlib.Tr("Sindhi, India"),
		"sd_IN":            dlib.Tr("Sindhi, India"),
		"sd_PK":            dlib.Tr("Sindhi, Pakistan"),
		"se_NO":            dlib.Tr("Northern Sami, Norway"),
		"shs_CA":           dlib.Tr("Secwepemctsin, Canada"),
		"sid_ET":           dlib.Tr("Sidama, Ethiopia."),
		"si_LK":            dlib.Tr("Sinhala, Sri Lanka"),
		"sk_SK.UTF-8":      dlib.Tr("Slovak, Slovak"),
		"sl_SI.UTF-8":      dlib.Tr("Slovenian, Slovenia"),
		"so_DJ.UTF-8":      dlib.Tr("Somali, Djibouti."),
		"so_ET":            dlib.Tr("Somali, Ethiopia"),
		"so_KE.UTF-8":      dlib.Tr("Somali, Kenya"),
		"so_SO.UTF-8":      dlib.Tr("Somali, Somalia"),
		"sq_AL.UTF-8":      dlib.Tr("Albanian, Albania"),
		"sq_MK":            dlib.Tr("Albanian, Macedonia"),
		"sr_ME":            dlib.Tr("Serbian, Montenegro"),
		"sr_RS@latin":      dlib.Tr("Serbian, Serbia (Latin)"),
		"sr_RS":            dlib.Tr("Serbian, Serbia"),
		"ss_ZA":            dlib.Tr("Swati, South Africa"),
		"st_ZA.UTF-8":      dlib.Tr("Sotho, South Africa"),
		"sv_FI.UTF-8":      dlib.Tr("Swedish, Finland"),
		"sv_SE.UTF-8":      dlib.Tr("Swedish, Sweden"),
		"sw_KE":            dlib.Tr("Swahili, Kenya"),
		"sw_TZ":            dlib.Tr("Swahili, Tanzania"),
		"ta_IN":            dlib.Tr("Tamil, India"),
		"ta_LK":            dlib.Tr("Tamil, Sri Lanka"),
		"te_IN":            dlib.Tr("Telugu, India"),
		"tg_TJ.UTF-8":      dlib.Tr("Tajik, Tajikistan"),
		"th_TH.UTF-8":      dlib.Tr("Thai, Thailand"),
		"ti_ER":            dlib.Tr("Tigrigna, Eritrea."),
		"ti_ET":            dlib.Tr("Tigrigna, Ethiopia."),
		"tig_ER":           dlib.Tr("Tigre, Eritrea"),
		"tk_TM":            dlib.Tr("Turkmen, Turkmenistan"),
		"tl_PH.UTF-8":      dlib.Tr("Tagalog, Philippines"),
		"tn_ZA":            dlib.Tr("Tswana, South Africa"),
		"tr_CY.UTF-8":      dlib.Tr("Turkish, Cyprus"),
		"tr_TR.UTF-8":      dlib.Tr("Turkish, Turkey"),
		"ts_ZA":            dlib.Tr("Tsonga, South Africa"),
		"tt_RU@iqtelif":    dlib.Tr("Tatar, Russia (IQTElif alphabet)"),
		"tt_RU":            dlib.Tr("Tatar, Russia"),
		"ug_CN@latin":      dlib.Tr("Uyghur, China (Latin alphabet)"),
		"ug_CN":            dlib.Tr("Uyghur, China"),
		"uk_UA.UTF-8":      dlib.Tr("Ukrainian, Ukraine"),
		"unm_US":           dlib.Tr("Unami Delaware, United States"),
		"ur_IN":            dlib.Tr("Urdu, India"),
		"ur_PK":            dlib.Tr("Urdu, Pakistan"),
		"uz_UZ@cyrillic":   dlib.Tr("Uzbek, Uzbekistan (cyrillic)"),
		"uz_UZ.UTF-8":      dlib.Tr("Uzbek, Uzbekistan"),
		"ve_ZA":            dlib.Tr("Venda, South Africa"),
		"vi_VN":            dlib.Tr("Vietnamese, Vietnam"),
		"wa_BE.UTF-8":      dlib.Tr("Walloon, Belgium"),
		"wae_CH":           dlib.Tr("Walser, Switzerland"),
		"wal_ET":           dlib.Tr("Walaita, Ethiopia"),
		"wo_SN":            dlib.Tr("Wolof, Senegal"),
		"xh_ZA.UTF-8":      dlib.Tr("Xhosa, South Africa"),
		"yi_US.UTF-8":      dlib.Tr("Yiddish, United States"),
		"yo_NG":            dlib.Tr("Yoruba, Nigeria"),
		"yue_HK":           dlib.Tr("Chinese, Hong Kong (Yue dialect)"),
		"zh_CN.UTF-8":      dlib.Tr("Chinese, China"),
		"zh_HK.UTF-8":      dlib.Tr("Chinese, Hong Kong"),
		"zh_SG.UTF-8":      dlib.Tr("Chinese, Singapore"),
		"zh_TW.UTF-8":      dlib.Tr("Chinese, Taiwan"),
		"zu_ZA.UTF-8":      dlib.Tr("Zulu, South Africa"),
	}
)
