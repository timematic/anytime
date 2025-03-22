
//line ragel/parse_datetime.go.rl:1
// GENERETED .hpp BY ragel AS:
//   ragel-go -G2 -e -o ragel_parse_datetime.go ragel_parse_datetime.go.rl
// it might be helpful to generate the FSM graph in debug:
//   ragel -Vp ragel_parse_datetime.go.rl -o ragel_parse_datetime.dot
//   dot ragel_parse_datetime.dot -Tpng -o ragel_parse_datetime.png

package ragel

import (
    "fmt"
    "strconv"
    "errors"
    "strings"
)


//line ragel/parse_datetime.go:20
const datetime_parser_start int = 1
const datetime_parser_first_final int = 867
const datetime_parser_error int = 0

const datetime_parser_en_main int = 1


//line ragel/parse_datetime.go.rl:24



//line ragel/parse_datetime.go:32
const start int = 1

const en_main int = 1


//line ragel/parse_datetime.go.rl:27

func Parse(data string) (st ParsedDatetime, err error) {
    cs, p, pe := 0, 0, len(data)
    eof := pe
    pb := p

    
//line ragel/parse_datetime.go:46
	{
	cs = start
	}

//line ragel/parse_datetime.go.rl:34
    
//line ragel/parse_datetime.go:53
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 867:
		goto st_case_867
	case 868:
		goto st_case_868
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 869:
		goto st_case_869
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 870:
		goto st_case_870
	case 17:
		goto st_case_17
	case 871:
		goto st_case_871
	case 18:
		goto st_case_18
	case 872:
		goto st_case_872
	case 19:
		goto st_case_19
	case 873:
		goto st_case_873
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 874:
		goto st_case_874
	case 23:
		goto st_case_23
	case 875:
		goto st_case_875
	case 24:
		goto st_case_24
	case 876:
		goto st_case_876
	case 877:
		goto st_case_877
	case 878:
		goto st_case_878
	case 879:
		goto st_case_879
	case 880:
		goto st_case_880
	case 881:
		goto st_case_881
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 882:
		goto st_case_882
	case 28:
		goto st_case_28
	case 883:
		goto st_case_883
	case 29:
		goto st_case_29
	case 884:
		goto st_case_884
	case 30:
		goto st_case_30
	case 885:
		goto st_case_885
	case 886:
		goto st_case_886
	case 887:
		goto st_case_887
	case 888:
		goto st_case_888
	case 889:
		goto st_case_889
	case 890:
		goto st_case_890
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 891:
		goto st_case_891
	case 892:
		goto st_case_892
	case 893:
		goto st_case_893
	case 34:
		goto st_case_34
	case 894:
		goto st_case_894
	case 895:
		goto st_case_895
	case 896:
		goto st_case_896
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 897:
		goto st_case_897
	case 898:
		goto st_case_898
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 899:
		goto st_case_899
	case 40:
		goto st_case_40
	case 900:
		goto st_case_900
	case 41:
		goto st_case_41
	case 901:
		goto st_case_901
	case 902:
		goto st_case_902
	case 903:
		goto st_case_903
	case 904:
		goto st_case_904
	case 905:
		goto st_case_905
	case 906:
		goto st_case_906
	case 907:
		goto st_case_907
	case 908:
		goto st_case_908
	case 909:
		goto st_case_909
	case 42:
		goto st_case_42
	case 910:
		goto st_case_910
	case 43:
		goto st_case_43
	case 911:
		goto st_case_911
	case 912:
		goto st_case_912
	case 44:
		goto st_case_44
	case 913:
		goto st_case_913
	case 45:
		goto st_case_45
	case 914:
		goto st_case_914
	case 915:
		goto st_case_915
	case 916:
		goto st_case_916
	case 917:
		goto st_case_917
	case 918:
		goto st_case_918
	case 919:
		goto st_case_919
	case 920:
		goto st_case_920
	case 921:
		goto st_case_921
	case 922:
		goto st_case_922
	case 923:
		goto st_case_923
	case 924:
		goto st_case_924
	case 925:
		goto st_case_925
	case 926:
		goto st_case_926
	case 46:
		goto st_case_46
	case 927:
		goto st_case_927
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 928:
		goto st_case_928
	case 929:
		goto st_case_929
	case 52:
		goto st_case_52
	case 930:
		goto st_case_930
	case 931:
		goto st_case_931
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 932:
		goto st_case_932
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 933:
		goto st_case_933
	case 934:
		goto st_case_934
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 935:
		goto st_case_935
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 936:
		goto st_case_936
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 937:
		goto st_case_937
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 177:
		goto st_case_177
	case 178:
		goto st_case_178
	case 179:
		goto st_case_179
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 185:
		goto st_case_185
	case 186:
		goto st_case_186
	case 187:
		goto st_case_187
	case 188:
		goto st_case_188
	case 189:
		goto st_case_189
	case 190:
		goto st_case_190
	case 191:
		goto st_case_191
	case 192:
		goto st_case_192
	case 193:
		goto st_case_193
	case 194:
		goto st_case_194
	case 195:
		goto st_case_195
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
	case 200:
		goto st_case_200
	case 201:
		goto st_case_201
	case 202:
		goto st_case_202
	case 203:
		goto st_case_203
	case 204:
		goto st_case_204
	case 205:
		goto st_case_205
	case 206:
		goto st_case_206
	case 207:
		goto st_case_207
	case 208:
		goto st_case_208
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	case 211:
		goto st_case_211
	case 212:
		goto st_case_212
	case 213:
		goto st_case_213
	case 214:
		goto st_case_214
	case 215:
		goto st_case_215
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 218:
		goto st_case_218
	case 219:
		goto st_case_219
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
	case 222:
		goto st_case_222
	case 223:
		goto st_case_223
	case 224:
		goto st_case_224
	case 225:
		goto st_case_225
	case 226:
		goto st_case_226
	case 227:
		goto st_case_227
	case 228:
		goto st_case_228
	case 229:
		goto st_case_229
	case 230:
		goto st_case_230
	case 231:
		goto st_case_231
	case 232:
		goto st_case_232
	case 233:
		goto st_case_233
	case 234:
		goto st_case_234
	case 235:
		goto st_case_235
	case 236:
		goto st_case_236
	case 237:
		goto st_case_237
	case 238:
		goto st_case_238
	case 239:
		goto st_case_239
	case 240:
		goto st_case_240
	case 241:
		goto st_case_241
	case 242:
		goto st_case_242
	case 243:
		goto st_case_243
	case 244:
		goto st_case_244
	case 245:
		goto st_case_245
	case 246:
		goto st_case_246
	case 247:
		goto st_case_247
	case 248:
		goto st_case_248
	case 249:
		goto st_case_249
	case 250:
		goto st_case_250
	case 251:
		goto st_case_251
	case 252:
		goto st_case_252
	case 253:
		goto st_case_253
	case 254:
		goto st_case_254
	case 255:
		goto st_case_255
	case 256:
		goto st_case_256
	case 257:
		goto st_case_257
	case 258:
		goto st_case_258
	case 259:
		goto st_case_259
	case 260:
		goto st_case_260
	case 261:
		goto st_case_261
	case 262:
		goto st_case_262
	case 263:
		goto st_case_263
	case 264:
		goto st_case_264
	case 265:
		goto st_case_265
	case 266:
		goto st_case_266
	case 267:
		goto st_case_267
	case 268:
		goto st_case_268
	case 269:
		goto st_case_269
	case 270:
		goto st_case_270
	case 271:
		goto st_case_271
	case 272:
		goto st_case_272
	case 273:
		goto st_case_273
	case 274:
		goto st_case_274
	case 275:
		goto st_case_275
	case 276:
		goto st_case_276
	case 277:
		goto st_case_277
	case 278:
		goto st_case_278
	case 279:
		goto st_case_279
	case 280:
		goto st_case_280
	case 281:
		goto st_case_281
	case 282:
		goto st_case_282
	case 283:
		goto st_case_283
	case 284:
		goto st_case_284
	case 285:
		goto st_case_285
	case 286:
		goto st_case_286
	case 287:
		goto st_case_287
	case 288:
		goto st_case_288
	case 289:
		goto st_case_289
	case 290:
		goto st_case_290
	case 291:
		goto st_case_291
	case 292:
		goto st_case_292
	case 293:
		goto st_case_293
	case 294:
		goto st_case_294
	case 295:
		goto st_case_295
	case 296:
		goto st_case_296
	case 297:
		goto st_case_297
	case 298:
		goto st_case_298
	case 299:
		goto st_case_299
	case 300:
		goto st_case_300
	case 301:
		goto st_case_301
	case 302:
		goto st_case_302
	case 303:
		goto st_case_303
	case 304:
		goto st_case_304
	case 305:
		goto st_case_305
	case 306:
		goto st_case_306
	case 307:
		goto st_case_307
	case 308:
		goto st_case_308
	case 309:
		goto st_case_309
	case 938:
		goto st_case_938
	case 310:
		goto st_case_310
	case 311:
		goto st_case_311
	case 939:
		goto st_case_939
	case 312:
		goto st_case_312
	case 313:
		goto st_case_313
	case 314:
		goto st_case_314
	case 940:
		goto st_case_940
	case 315:
		goto st_case_315
	case 941:
		goto st_case_941
	case 316:
		goto st_case_316
	case 942:
		goto st_case_942
	case 943:
		goto st_case_943
	case 944:
		goto st_case_944
	case 945:
		goto st_case_945
	case 946:
		goto st_case_946
	case 947:
		goto st_case_947
	case 317:
		goto st_case_317
	case 318:
		goto st_case_318
	case 319:
		goto st_case_319
	case 948:
		goto st_case_948
	case 320:
		goto st_case_320
	case 949:
		goto st_case_949
	case 950:
		goto st_case_950
	case 951:
		goto st_case_951
	case 952:
		goto st_case_952
	case 953:
		goto st_case_953
	case 954:
		goto st_case_954
	case 321:
		goto st_case_321
	case 322:
		goto st_case_322
	case 323:
		goto st_case_323
	case 955:
		goto st_case_955
	case 956:
		goto st_case_956
	case 324:
		goto st_case_324
	case 325:
		goto st_case_325
	case 326:
		goto st_case_326
	case 327:
		goto st_case_327
	case 328:
		goto st_case_328
	case 329:
		goto st_case_329
	case 330:
		goto st_case_330
	case 957:
		goto st_case_957
	case 958:
		goto st_case_958
	case 959:
		goto st_case_959
	case 960:
		goto st_case_960
	case 331:
		goto st_case_331
	case 961:
		goto st_case_961
	case 962:
		goto st_case_962
	case 332:
		goto st_case_332
	case 963:
		goto st_case_963
	case 964:
		goto st_case_964
	case 333:
		goto st_case_333
	case 334:
		goto st_case_334
	case 335:
		goto st_case_335
	case 336:
		goto st_case_336
	case 965:
		goto st_case_965
	case 966:
		goto st_case_966
	case 337:
		goto st_case_337
	case 967:
		goto st_case_967
	case 968:
		goto st_case_968
	case 338:
		goto st_case_338
	case 969:
		goto st_case_969
	case 339:
		goto st_case_339
	case 970:
		goto st_case_970
	case 340:
		goto st_case_340
	case 971:
		goto st_case_971
	case 972:
		goto st_case_972
	case 973:
		goto st_case_973
	case 974:
		goto st_case_974
	case 975:
		goto st_case_975
	case 976:
		goto st_case_976
	case 977:
		goto st_case_977
	case 978:
		goto st_case_978
	case 979:
		goto st_case_979
	case 341:
		goto st_case_341
	case 980:
		goto st_case_980
	case 342:
		goto st_case_342
	case 981:
		goto st_case_981
	case 982:
		goto st_case_982
	case 343:
		goto st_case_343
	case 983:
		goto st_case_983
	case 344:
		goto st_case_344
	case 984:
		goto st_case_984
	case 985:
		goto st_case_985
	case 986:
		goto st_case_986
	case 987:
		goto st_case_987
	case 988:
		goto st_case_988
	case 989:
		goto st_case_989
	case 990:
		goto st_case_990
	case 991:
		goto st_case_991
	case 992:
		goto st_case_992
	case 993:
		goto st_case_993
	case 994:
		goto st_case_994
	case 995:
		goto st_case_995
	case 996:
		goto st_case_996
	case 345:
		goto st_case_345
	case 997:
		goto st_case_997
	case 346:
		goto st_case_346
	case 347:
		goto st_case_347
	case 348:
		goto st_case_348
	case 349:
		goto st_case_349
	case 350:
		goto st_case_350
	case 351:
		goto st_case_351
	case 352:
		goto st_case_352
	case 353:
		goto st_case_353
	case 354:
		goto st_case_354
	case 355:
		goto st_case_355
	case 356:
		goto st_case_356
	case 357:
		goto st_case_357
	case 358:
		goto st_case_358
	case 359:
		goto st_case_359
	case 360:
		goto st_case_360
	case 361:
		goto st_case_361
	case 362:
		goto st_case_362
	case 363:
		goto st_case_363
	case 364:
		goto st_case_364
	case 365:
		goto st_case_365
	case 366:
		goto st_case_366
	case 367:
		goto st_case_367
	case 368:
		goto st_case_368
	case 369:
		goto st_case_369
	case 370:
		goto st_case_370
	case 371:
		goto st_case_371
	case 372:
		goto st_case_372
	case 373:
		goto st_case_373
	case 374:
		goto st_case_374
	case 375:
		goto st_case_375
	case 376:
		goto st_case_376
	case 377:
		goto st_case_377
	case 378:
		goto st_case_378
	case 379:
		goto st_case_379
	case 380:
		goto st_case_380
	case 381:
		goto st_case_381
	case 382:
		goto st_case_382
	case 383:
		goto st_case_383
	case 384:
		goto st_case_384
	case 998:
		goto st_case_998
	case 385:
		goto st_case_385
	case 999:
		goto st_case_999
	case 386:
		goto st_case_386
	case 387:
		goto st_case_387
	case 388:
		goto st_case_388
	case 389:
		goto st_case_389
	case 390:
		goto st_case_390
	case 1000:
		goto st_case_1000
	case 1001:
		goto st_case_1001
	case 391:
		goto st_case_391
	case 392:
		goto st_case_392
	case 393:
		goto st_case_393
	case 394:
		goto st_case_394
	case 395:
		goto st_case_395
	case 396:
		goto st_case_396
	case 397:
		goto st_case_397
	case 398:
		goto st_case_398
	case 399:
		goto st_case_399
	case 400:
		goto st_case_400
	case 401:
		goto st_case_401
	case 402:
		goto st_case_402
	case 403:
		goto st_case_403
	case 404:
		goto st_case_404
	case 405:
		goto st_case_405
	case 406:
		goto st_case_406
	case 407:
		goto st_case_407
	case 408:
		goto st_case_408
	case 409:
		goto st_case_409
	case 410:
		goto st_case_410
	case 411:
		goto st_case_411
	case 412:
		goto st_case_412
	case 413:
		goto st_case_413
	case 414:
		goto st_case_414
	case 415:
		goto st_case_415
	case 416:
		goto st_case_416
	case 417:
		goto st_case_417
	case 418:
		goto st_case_418
	case 419:
		goto st_case_419
	case 420:
		goto st_case_420
	case 421:
		goto st_case_421
	case 422:
		goto st_case_422
	case 423:
		goto st_case_423
	case 424:
		goto st_case_424
	case 425:
		goto st_case_425
	case 426:
		goto st_case_426
	case 427:
		goto st_case_427
	case 428:
		goto st_case_428
	case 429:
		goto st_case_429
	case 430:
		goto st_case_430
	case 431:
		goto st_case_431
	case 432:
		goto st_case_432
	case 433:
		goto st_case_433
	case 434:
		goto st_case_434
	case 435:
		goto st_case_435
	case 436:
		goto st_case_436
	case 437:
		goto st_case_437
	case 438:
		goto st_case_438
	case 439:
		goto st_case_439
	case 440:
		goto st_case_440
	case 441:
		goto st_case_441
	case 442:
		goto st_case_442
	case 1002:
		goto st_case_1002
	case 443:
		goto st_case_443
	case 444:
		goto st_case_444
	case 445:
		goto st_case_445
	case 446:
		goto st_case_446
	case 447:
		goto st_case_447
	case 448:
		goto st_case_448
	case 449:
		goto st_case_449
	case 450:
		goto st_case_450
	case 451:
		goto st_case_451
	case 452:
		goto st_case_452
	case 453:
		goto st_case_453
	case 454:
		goto st_case_454
	case 455:
		goto st_case_455
	case 456:
		goto st_case_456
	case 457:
		goto st_case_457
	case 458:
		goto st_case_458
	case 459:
		goto st_case_459
	case 460:
		goto st_case_460
	case 461:
		goto st_case_461
	case 462:
		goto st_case_462
	case 463:
		goto st_case_463
	case 464:
		goto st_case_464
	case 465:
		goto st_case_465
	case 466:
		goto st_case_466
	case 467:
		goto st_case_467
	case 468:
		goto st_case_468
	case 469:
		goto st_case_469
	case 470:
		goto st_case_470
	case 471:
		goto st_case_471
	case 472:
		goto st_case_472
	case 473:
		goto st_case_473
	case 474:
		goto st_case_474
	case 475:
		goto st_case_475
	case 476:
		goto st_case_476
	case 477:
		goto st_case_477
	case 478:
		goto st_case_478
	case 479:
		goto st_case_479
	case 480:
		goto st_case_480
	case 481:
		goto st_case_481
	case 482:
		goto st_case_482
	case 483:
		goto st_case_483
	case 484:
		goto st_case_484
	case 485:
		goto st_case_485
	case 486:
		goto st_case_486
	case 487:
		goto st_case_487
	case 488:
		goto st_case_488
	case 489:
		goto st_case_489
	case 490:
		goto st_case_490
	case 491:
		goto st_case_491
	case 492:
		goto st_case_492
	case 493:
		goto st_case_493
	case 494:
		goto st_case_494
	case 495:
		goto st_case_495
	case 496:
		goto st_case_496
	case 497:
		goto st_case_497
	case 498:
		goto st_case_498
	case 499:
		goto st_case_499
	case 500:
		goto st_case_500
	case 501:
		goto st_case_501
	case 502:
		goto st_case_502
	case 503:
		goto st_case_503
	case 504:
		goto st_case_504
	case 505:
		goto st_case_505
	case 506:
		goto st_case_506
	case 507:
		goto st_case_507
	case 508:
		goto st_case_508
	case 509:
		goto st_case_509
	case 510:
		goto st_case_510
	case 511:
		goto st_case_511
	case 512:
		goto st_case_512
	case 513:
		goto st_case_513
	case 514:
		goto st_case_514
	case 515:
		goto st_case_515
	case 516:
		goto st_case_516
	case 517:
		goto st_case_517
	case 518:
		goto st_case_518
	case 519:
		goto st_case_519
	case 520:
		goto st_case_520
	case 521:
		goto st_case_521
	case 522:
		goto st_case_522
	case 523:
		goto st_case_523
	case 524:
		goto st_case_524
	case 525:
		goto st_case_525
	case 1003:
		goto st_case_1003
	case 1004:
		goto st_case_1004
	case 526:
		goto st_case_526
	case 527:
		goto st_case_527
	case 1005:
		goto st_case_1005
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
	case 530:
		goto st_case_530
	case 1006:
		goto st_case_1006
	case 531:
		goto st_case_531
	case 1007:
		goto st_case_1007
	case 532:
		goto st_case_532
	case 1008:
		goto st_case_1008
	case 1009:
		goto st_case_1009
	case 1010:
		goto st_case_1010
	case 1011:
		goto st_case_1011
	case 1012:
		goto st_case_1012
	case 1013:
		goto st_case_1013
	case 533:
		goto st_case_533
	case 534:
		goto st_case_534
	case 535:
		goto st_case_535
	case 1014:
		goto st_case_1014
	case 536:
		goto st_case_536
	case 1015:
		goto st_case_1015
	case 537:
		goto st_case_537
	case 1016:
		goto st_case_1016
	case 538:
		goto st_case_538
	case 1017:
		goto st_case_1017
	case 1018:
		goto st_case_1018
	case 1019:
		goto st_case_1019
	case 1020:
		goto st_case_1020
	case 1021:
		goto st_case_1021
	case 1022:
		goto st_case_1022
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 541:
		goto st_case_541
	case 1023:
		goto st_case_1023
	case 542:
		goto st_case_542
	case 1024:
		goto st_case_1024
	case 543:
		goto st_case_543
	case 544:
		goto st_case_544
	case 1025:
		goto st_case_1025
	case 1026:
		goto st_case_1026
	case 545:
		goto st_case_545
	case 546:
		goto st_case_546
	case 547:
		goto st_case_547
	case 548:
		goto st_case_548
	case 549:
		goto st_case_549
	case 550:
		goto st_case_550
	case 551:
		goto st_case_551
	case 552:
		goto st_case_552
	case 553:
		goto st_case_553
	case 554:
		goto st_case_554
	case 555:
		goto st_case_555
	case 556:
		goto st_case_556
	case 557:
		goto st_case_557
	case 558:
		goto st_case_558
	case 559:
		goto st_case_559
	case 1027:
		goto st_case_1027
	case 1028:
		goto st_case_1028
	case 560:
		goto st_case_560
	case 561:
		goto st_case_561
	case 562:
		goto st_case_562
	case 563:
		goto st_case_563
	case 564:
		goto st_case_564
	case 565:
		goto st_case_565
	case 566:
		goto st_case_566
	case 567:
		goto st_case_567
	case 568:
		goto st_case_568
	case 569:
		goto st_case_569
	case 570:
		goto st_case_570
	case 571:
		goto st_case_571
	case 572:
		goto st_case_572
	case 573:
		goto st_case_573
	case 574:
		goto st_case_574
	case 575:
		goto st_case_575
	case 576:
		goto st_case_576
	case 577:
		goto st_case_577
	case 578:
		goto st_case_578
	case 579:
		goto st_case_579
	case 580:
		goto st_case_580
	case 581:
		goto st_case_581
	case 582:
		goto st_case_582
	case 583:
		goto st_case_583
	case 584:
		goto st_case_584
	case 585:
		goto st_case_585
	case 586:
		goto st_case_586
	case 1029:
		goto st_case_1029
	case 587:
		goto st_case_587
	case 588:
		goto st_case_588
	case 589:
		goto st_case_589
	case 590:
		goto st_case_590
	case 591:
		goto st_case_591
	case 592:
		goto st_case_592
	case 593:
		goto st_case_593
	case 594:
		goto st_case_594
	case 595:
		goto st_case_595
	case 596:
		goto st_case_596
	case 597:
		goto st_case_597
	case 598:
		goto st_case_598
	case 599:
		goto st_case_599
	case 600:
		goto st_case_600
	case 601:
		goto st_case_601
	case 602:
		goto st_case_602
	case 603:
		goto st_case_603
	case 604:
		goto st_case_604
	case 605:
		goto st_case_605
	case 606:
		goto st_case_606
	case 607:
		goto st_case_607
	case 608:
		goto st_case_608
	case 609:
		goto st_case_609
	case 610:
		goto st_case_610
	case 611:
		goto st_case_611
	case 612:
		goto st_case_612
	case 613:
		goto st_case_613
	case 614:
		goto st_case_614
	case 615:
		goto st_case_615
	case 616:
		goto st_case_616
	case 617:
		goto st_case_617
	case 618:
		goto st_case_618
	case 619:
		goto st_case_619
	case 620:
		goto st_case_620
	case 621:
		goto st_case_621
	case 622:
		goto st_case_622
	case 623:
		goto st_case_623
	case 624:
		goto st_case_624
	case 625:
		goto st_case_625
	case 626:
		goto st_case_626
	case 627:
		goto st_case_627
	case 628:
		goto st_case_628
	case 629:
		goto st_case_629
	case 630:
		goto st_case_630
	case 631:
		goto st_case_631
	case 632:
		goto st_case_632
	case 633:
		goto st_case_633
	case 634:
		goto st_case_634
	case 635:
		goto st_case_635
	case 636:
		goto st_case_636
	case 637:
		goto st_case_637
	case 638:
		goto st_case_638
	case 639:
		goto st_case_639
	case 640:
		goto st_case_640
	case 641:
		goto st_case_641
	case 642:
		goto st_case_642
	case 643:
		goto st_case_643
	case 644:
		goto st_case_644
	case 645:
		goto st_case_645
	case 646:
		goto st_case_646
	case 647:
		goto st_case_647
	case 648:
		goto st_case_648
	case 649:
		goto st_case_649
	case 650:
		goto st_case_650
	case 651:
		goto st_case_651
	case 652:
		goto st_case_652
	case 653:
		goto st_case_653
	case 654:
		goto st_case_654
	case 655:
		goto st_case_655
	case 656:
		goto st_case_656
	case 657:
		goto st_case_657
	case 658:
		goto st_case_658
	case 659:
		goto st_case_659
	case 660:
		goto st_case_660
	case 661:
		goto st_case_661
	case 662:
		goto st_case_662
	case 663:
		goto st_case_663
	case 664:
		goto st_case_664
	case 665:
		goto st_case_665
	case 666:
		goto st_case_666
	case 667:
		goto st_case_667
	case 668:
		goto st_case_668
	case 669:
		goto st_case_669
	case 670:
		goto st_case_670
	case 671:
		goto st_case_671
	case 672:
		goto st_case_672
	case 673:
		goto st_case_673
	case 674:
		goto st_case_674
	case 675:
		goto st_case_675
	case 676:
		goto st_case_676
	case 677:
		goto st_case_677
	case 678:
		goto st_case_678
	case 679:
		goto st_case_679
	case 680:
		goto st_case_680
	case 681:
		goto st_case_681
	case 682:
		goto st_case_682
	case 683:
		goto st_case_683
	case 684:
		goto st_case_684
	case 685:
		goto st_case_685
	case 686:
		goto st_case_686
	case 687:
		goto st_case_687
	case 688:
		goto st_case_688
	case 689:
		goto st_case_689
	case 690:
		goto st_case_690
	case 691:
		goto st_case_691
	case 692:
		goto st_case_692
	case 693:
		goto st_case_693
	case 694:
		goto st_case_694
	case 695:
		goto st_case_695
	case 696:
		goto st_case_696
	case 697:
		goto st_case_697
	case 698:
		goto st_case_698
	case 699:
		goto st_case_699
	case 700:
		goto st_case_700
	case 701:
		goto st_case_701
	case 702:
		goto st_case_702
	case 703:
		goto st_case_703
	case 704:
		goto st_case_704
	case 705:
		goto st_case_705
	case 706:
		goto st_case_706
	case 707:
		goto st_case_707
	case 708:
		goto st_case_708
	case 709:
		goto st_case_709
	case 710:
		goto st_case_710
	case 711:
		goto st_case_711
	case 712:
		goto st_case_712
	case 713:
		goto st_case_713
	case 714:
		goto st_case_714
	case 715:
		goto st_case_715
	case 716:
		goto st_case_716
	case 717:
		goto st_case_717
	case 718:
		goto st_case_718
	case 719:
		goto st_case_719
	case 720:
		goto st_case_720
	case 721:
		goto st_case_721
	case 722:
		goto st_case_722
	case 723:
		goto st_case_723
	case 724:
		goto st_case_724
	case 725:
		goto st_case_725
	case 726:
		goto st_case_726
	case 727:
		goto st_case_727
	case 728:
		goto st_case_728
	case 729:
		goto st_case_729
	case 730:
		goto st_case_730
	case 731:
		goto st_case_731
	case 732:
		goto st_case_732
	case 733:
		goto st_case_733
	case 734:
		goto st_case_734
	case 735:
		goto st_case_735
	case 736:
		goto st_case_736
	case 737:
		goto st_case_737
	case 738:
		goto st_case_738
	case 739:
		goto st_case_739
	case 740:
		goto st_case_740
	case 741:
		goto st_case_741
	case 742:
		goto st_case_742
	case 743:
		goto st_case_743
	case 744:
		goto st_case_744
	case 745:
		goto st_case_745
	case 746:
		goto st_case_746
	case 747:
		goto st_case_747
	case 748:
		goto st_case_748
	case 749:
		goto st_case_749
	case 750:
		goto st_case_750
	case 751:
		goto st_case_751
	case 752:
		goto st_case_752
	case 753:
		goto st_case_753
	case 754:
		goto st_case_754
	case 755:
		goto st_case_755
	case 756:
		goto st_case_756
	case 757:
		goto st_case_757
	case 758:
		goto st_case_758
	case 759:
		goto st_case_759
	case 760:
		goto st_case_760
	case 761:
		goto st_case_761
	case 762:
		goto st_case_762
	case 763:
		goto st_case_763
	case 764:
		goto st_case_764
	case 765:
		goto st_case_765
	case 766:
		goto st_case_766
	case 767:
		goto st_case_767
	case 768:
		goto st_case_768
	case 769:
		goto st_case_769
	case 770:
		goto st_case_770
	case 771:
		goto st_case_771
	case 772:
		goto st_case_772
	case 773:
		goto st_case_773
	case 774:
		goto st_case_774
	case 775:
		goto st_case_775
	case 776:
		goto st_case_776
	case 777:
		goto st_case_777
	case 778:
		goto st_case_778
	case 779:
		goto st_case_779
	case 780:
		goto st_case_780
	case 781:
		goto st_case_781
	case 782:
		goto st_case_782
	case 783:
		goto st_case_783
	case 784:
		goto st_case_784
	case 785:
		goto st_case_785
	case 786:
		goto st_case_786
	case 787:
		goto st_case_787
	case 788:
		goto st_case_788
	case 789:
		goto st_case_789
	case 790:
		goto st_case_790
	case 791:
		goto st_case_791
	case 792:
		goto st_case_792
	case 793:
		goto st_case_793
	case 794:
		goto st_case_794
	case 795:
		goto st_case_795
	case 796:
		goto st_case_796
	case 797:
		goto st_case_797
	case 798:
		goto st_case_798
	case 799:
		goto st_case_799
	case 800:
		goto st_case_800
	case 801:
		goto st_case_801
	case 802:
		goto st_case_802
	case 803:
		goto st_case_803
	case 804:
		goto st_case_804
	case 805:
		goto st_case_805
	case 806:
		goto st_case_806
	case 807:
		goto st_case_807
	case 808:
		goto st_case_808
	case 809:
		goto st_case_809
	case 810:
		goto st_case_810
	case 811:
		goto st_case_811
	case 812:
		goto st_case_812
	case 813:
		goto st_case_813
	case 814:
		goto st_case_814
	case 815:
		goto st_case_815
	case 816:
		goto st_case_816
	case 817:
		goto st_case_817
	case 818:
		goto st_case_818
	case 819:
		goto st_case_819
	case 820:
		goto st_case_820
	case 821:
		goto st_case_821
	case 822:
		goto st_case_822
	case 823:
		goto st_case_823
	case 824:
		goto st_case_824
	case 825:
		goto st_case_825
	case 826:
		goto st_case_826
	case 827:
		goto st_case_827
	case 828:
		goto st_case_828
	case 829:
		goto st_case_829
	case 830:
		goto st_case_830
	case 831:
		goto st_case_831
	case 832:
		goto st_case_832
	case 833:
		goto st_case_833
	case 834:
		goto st_case_834
	case 835:
		goto st_case_835
	case 836:
		goto st_case_836
	case 837:
		goto st_case_837
	case 838:
		goto st_case_838
	case 839:
		goto st_case_839
	case 840:
		goto st_case_840
	case 841:
		goto st_case_841
	case 842:
		goto st_case_842
	case 843:
		goto st_case_843
	case 844:
		goto st_case_844
	case 845:
		goto st_case_845
	case 846:
		goto st_case_846
	case 847:
		goto st_case_847
	case 848:
		goto st_case_848
	case 849:
		goto st_case_849
	case 850:
		goto st_case_850
	case 851:
		goto st_case_851
	case 852:
		goto st_case_852
	case 853:
		goto st_case_853
	case 854:
		goto st_case_854
	case 855:
		goto st_case_855
	case 856:
		goto st_case_856
	case 857:
		goto st_case_857
	case 858:
		goto st_case_858
	case 859:
		goto st_case_859
	case 860:
		goto st_case_860
	case 861:
		goto st_case_861
	case 862:
		goto st_case_862
	case 863:
		goto st_case_863
	case 864:
		goto st_case_864
	case 865:
		goto st_case_865
	case 866:
		goto st_case_866
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 48:
			goto tr0
		case 51:
			goto tr3
		case 65:
			goto st298
		case 68:
			goto st406
		case 70:
			goto st414
		case 74:
			goto st807
		case 77:
			goto st819
		case 78:
			goto st826
		case 79:
			goto st834
		case 83:
			goto st841
		case 84:
			goto st854
		case 87:
			goto st860
		case 97:
			goto st298
		case 100:
			goto st406
		case 102:
			goto st864
		case 106:
			goto st807
		case 109:
			goto st865
		case 110:
			goto st826
		case 111:
			goto st834
		case 115:
			goto st866
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr4
			}
		case data[p] >= 49:
			goto tr2
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
//line ragel/datetime.rl:5
 pb = p 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line ragel/parse_datetime.go:2186
		if data[p] == 48 {
			goto st3
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st146
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 32:
			goto tr22
		case 229:
			goto tr25
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr24
			}
		case data[p] >= 45:
			goto tr23
		}
		goto st0
tr22:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line ragel/parse_datetime.go:2243
		switch data[p] {
		case 48:
			goto tr26
		case 49:
			goto tr27
		case 65:
			goto st54
		case 68:
			goto st65
		case 70:
			goto st73
		case 74:
			goto st81
		case 77:
			goto st93
		case 78:
			goto st99
		case 79:
			goto st107
		case 83:
			goto st114
		case 97:
			goto st54
		case 100:
			goto st65
		case 102:
			goto st73
		case 106:
			goto st81
		case 109:
			goto st93
		case 110:
			goto st99
		case 111:
			goto st107
		case 115:
			goto st114
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr28
		}
		goto st0
tr26:
//line ragel/datetime.rl:5
 pb = p 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line ragel/parse_datetime.go:2295
		if 49 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto st0
tr28:
//line ragel/datetime.rl:5
 pb = p 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line ragel/parse_datetime.go:2309
		if data[p] == 32 {
			goto tr38
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr38
		}
		goto st0
tr38:
//line ragel/datetime.rl:9

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st9
tr100:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st9
tr106:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st9
tr113:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st9
tr122:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st9
tr132:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st9
tr140:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st9
tr143:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st9
tr149:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st9
tr153:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st9
tr157:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st9
tr166:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st9
tr174:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line ragel/parse_datetime.go:2376
		switch data[p] {
		case 48:
			goto tr39
		case 51:
			goto tr41
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr42
			}
		case data[p] >= 49:
			goto tr40
		}
		goto st0
tr39:
//line ragel/datetime.rl:5
 pb = p 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line ragel/parse_datetime.go:2401
		if 49 <= data[p] && data[p] <= 57 {
			goto st867
		}
		goto st0
tr42:
//line ragel/datetime.rl:5
 pb = p 
	goto st867
	st867:
		if p++; p == pe {
			goto _test_eof867
		}
	st_case_867:
//line ragel/parse_datetime.go:2415
		switch data[p] {
		case 32:
			goto tr1199
		case 43:
			goto tr1200
		case 44:
			goto tr1201
		case 45:
			goto tr1202
		case 47:
			goto tr1203
		case 84:
			goto tr1204
		case 90:
			goto tr1205
		case 95:
			goto tr1206
		case 116:
			goto tr1206
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1203
			}
		case data[p] >= 65:
			goto tr1203
		}
		goto st0
tr1356:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st868
tr1199:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st868
tr1335:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st868
tr1344:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st868
tr1364:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st868
	st868:
		if p++; p == pe {
			goto _test_eof868
		}
	st_case_868:
//line ragel/parse_datetime.go:2486
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st19
		case 45:
			goto st31
		case 47:
			goto tr1210
		case 50:
			goto tr95
		case 65:
			goto tr1211
		case 66:
			goto tr1212
		case 90:
			goto tr1213
		case 95:
			goto tr1210
		case 97:
			goto tr1214
		case 109:
			goto tr1215
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr94
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1210
				}
			case data[p] >= 67:
				goto tr1210
			}
		default:
			goto tr96
		}
		goto st0
tr1223:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line ragel/parse_datetime.go:2538
		switch data[p] {
		case 65:
			goto st12
		case 66:
			goto st18
		case 109:
			goto st14
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 68 {
			goto st869
		}
		goto st0
	st869:
		if p++; p == pe {
			goto _test_eof869
		}
	st_case_869:
		if data[p] == 32 {
			goto st13
		}
		goto st0
tr1499:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st13
tr1376:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st13
tr1219:
//line ragel/datetime.rl:67

    st.Ad_bc = ADBC_BC;

	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line ragel/parse_datetime.go:2587
		if data[p] == 109 {
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 61 {
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 43:
			goto st16
		case 45:
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr50
		}
		goto st0
tr50:
//line ragel/datetime.rl:5
 pb = p 
	goto st870
	st870:
		if p++; p == pe {
			goto _test_eof870
		}
	st_case_870:
//line ragel/parse_datetime.go:2631
		if data[p] == 46 {
			goto st17
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st870
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if 48 <= data[p] && data[p] <= 57 {
			goto st871
		}
		goto st0
	st871:
		if p++; p == pe {
			goto _test_eof871
		}
	st_case_871:
		if 48 <= data[p] && data[p] <= 57 {
			goto st871
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 67 {
			goto st872
		}
		goto st0
	st872:
		if p++; p == pe {
			goto _test_eof872
		}
	st_case_872:
		if data[p] == 32 {
			goto tr1219
		}
		goto st0
tr1357:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st19
tr1200:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st19
tr1243:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st19
tr1255:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st19
tr1260:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st19
tr1275:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st19
tr1268:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st19
tr1288:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st19
tr1297:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st19
tr1305:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st19
tr1325:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st19
tr1336:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st19
tr1345:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st19
tr1365:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line ragel/parse_datetime.go:2825
		if data[p] == 50 {
			goto tr54
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr55
			}
		case data[p] >= 48:
			goto tr53
		}
		goto st0
tr53:
//line ragel/datetime.rl:5
 pb = p 
	goto st873
tr75:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st873
	st873:
		if p++; p == pe {
			goto _test_eof873
		}
	st_case_873:
//line ragel/parse_datetime.go:2853
		switch data[p] {
		case 32:
			goto tr1220
		case 58:
			goto tr1222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st885
		}
		goto st0
tr1220:
//line ragel/datetime.rl:197

    st.ZoneOffsetIsValid = true

    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st20
tr1235:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st20
tr1238:
//line ragel/datetime.rl:187

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line ragel/parse_datetime.go:2921
		switch data[p] {
		case 40:
			goto st21
		case 43:
			goto st23
		case 45:
			goto st25
		case 47:
			goto tr59
		case 65:
			goto tr60
		case 66:
			goto tr61
		case 95:
			goto tr59
		case 109:
			goto tr62
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr59
			}
		case data[p] >= 67:
			goto tr59
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 32:
			goto st22
		case 43:
			goto st22
		case 45:
			goto st22
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st22
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		default:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 32:
			goto st22
		case 41:
			goto st874
		case 43:
			goto st22
		case 45:
			goto st22
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st22
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		default:
			goto st22
		}
		goto st0
	st874:
		if p++; p == pe {
			goto _test_eof874
		}
	st_case_874:
		if data[p] == 32 {
			goto tr1223
		}
		goto st0
tr1240:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line ragel/parse_datetime.go:3024
		if data[p] == 50 {
			goto tr66
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr67
			}
		case data[p] >= 48:
			goto tr65
		}
		goto st0
tr65:
//line ragel/datetime.rl:5
 pb = p 
	goto st875
tr68:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st875
	st875:
		if p++; p == pe {
			goto _test_eof875
		}
	st_case_875:
//line ragel/parse_datetime.go:3052
		switch data[p] {
		case 32:
			goto tr1224
		case 58:
			goto tr1226
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st876
		}
		goto st0
tr1224:
//line ragel/datetime.rl:197

    st.ZoneOffsetIsValid = true

    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st24
tr1228:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st24
tr1231:
//line ragel/datetime.rl:187

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st24
tr1233:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line ragel/parse_datetime.go:3129
		switch data[p] {
		case 40:
			goto st21
		case 65:
			goto st12
		case 66:
			goto st18
		case 109:
			goto st14
		}
		goto st0
tr67:
//line ragel/datetime.rl:5
 pb = p 
	goto st876
tr70:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st876
	st876:
		if p++; p == pe {
			goto _test_eof876
		}
	st_case_876:
//line ragel/parse_datetime.go:3156
		switch data[p] {
		case 32:
			goto tr1224
		case 58:
			goto tr1226
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st877
		}
		goto st0
	st877:
		if p++; p == pe {
			goto _test_eof877
		}
	st_case_877:
		if data[p] == 32 {
			goto tr1224
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st877
		}
		goto st0
tr1226:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st878
	st878:
		if p++; p == pe {
			goto _test_eof878
		}
	st_case_878:
//line ragel/parse_datetime.go:3196
		if data[p] == 32 {
			goto tr1228
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1230
			}
		case data[p] >= 48:
			goto tr1229
		}
		goto st0
tr1229:
//line ragel/datetime.rl:5
 pb = p 
	goto st879
	st879:
		if p++; p == pe {
			goto _test_eof879
		}
	st_case_879:
//line ragel/parse_datetime.go:3218
		if data[p] == 32 {
			goto tr1231
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st880
		}
		goto st0
tr1230:
//line ragel/datetime.rl:5
 pb = p 
	goto st880
	st880:
		if p++; p == pe {
			goto _test_eof880
		}
	st_case_880:
//line ragel/parse_datetime.go:3235
		if data[p] == 32 {
			goto tr1231
		}
		goto st0
tr66:
//line ragel/datetime.rl:5
 pb = p 
	goto st881
tr69:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st881
	st881:
		if p++; p == pe {
			goto _test_eof881
		}
	st_case_881:
//line ragel/parse_datetime.go:3255
		switch data[p] {
		case 32:
			goto tr1224
		case 58:
			goto tr1226
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st877
			}
		case data[p] >= 48:
			goto st876
		}
		goto st0
tr1241:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line ragel/parse_datetime.go:3283
		if data[p] == 50 {
			goto tr69
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr70
			}
		case data[p] >= 48:
			goto tr68
		}
		goto st0
tr59:
//line ragel/datetime.rl:5
 pb = p 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line ragel/parse_datetime.go:3305
		switch data[p] {
		case 47:
			goto st27
		case 95:
			goto st27
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st27
			}
		case data[p] >= 65:
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 47:
			goto st882
		case 95:
			goto st882
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st882
			}
		case data[p] >= 65:
			goto st882
		}
		goto st0
	st882:
		if p++; p == pe {
			goto _test_eof882
		}
	st_case_882:
		switch data[p] {
		case 32:
			goto tr1233
		case 47:
			goto st882
		case 95:
			goto st882
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st882
			}
		case data[p] >= 65:
			goto st882
		}
		goto st0
tr60:
//line ragel/datetime.rl:5
 pb = p 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line ragel/parse_datetime.go:3372
		switch data[p] {
		case 47:
			goto st27
		case 68:
			goto st883
		case 95:
			goto st27
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st27
			}
		case data[p] >= 65:
			goto st27
		}
		goto st0
	st883:
		if p++; p == pe {
			goto _test_eof883
		}
	st_case_883:
		switch data[p] {
		case 32:
			goto st13
		case 47:
			goto st882
		case 95:
			goto st882
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st882
			}
		case data[p] >= 65:
			goto st882
		}
		goto st0
tr61:
//line ragel/datetime.rl:5
 pb = p 
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line ragel/parse_datetime.go:3421
		switch data[p] {
		case 47:
			goto st27
		case 67:
			goto st884
		case 95:
			goto st27
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st27
			}
		case data[p] >= 65:
			goto st27
		}
		goto st0
	st884:
		if p++; p == pe {
			goto _test_eof884
		}
	st_case_884:
		switch data[p] {
		case 32:
			goto tr1219
		case 47:
			goto st882
		case 95:
			goto st882
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st882
			}
		case data[p] >= 65:
			goto st882
		}
		goto st0
tr62:
//line ragel/datetime.rl:5
 pb = p 
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line ragel/parse_datetime.go:3470
		switch data[p] {
		case 47:
			goto st27
		case 61:
			goto st15
		case 95:
			goto st27
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st27
			}
		case data[p] >= 65:
			goto st27
		}
		goto st0
tr55:
//line ragel/datetime.rl:5
 pb = p 
	goto st885
tr77:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st885
	st885:
		if p++; p == pe {
			goto _test_eof885
		}
	st_case_885:
//line ragel/parse_datetime.go:3503
		switch data[p] {
		case 32:
			goto tr1220
		case 58:
			goto tr1222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st886
		}
		goto st0
	st886:
		if p++; p == pe {
			goto _test_eof886
		}
	st_case_886:
		if data[p] == 32 {
			goto tr1220
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st886
		}
		goto st0
tr1222:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st887
	st887:
		if p++; p == pe {
			goto _test_eof887
		}
	st_case_887:
//line ragel/parse_datetime.go:3543
		if data[p] == 32 {
			goto tr1235
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1237
			}
		case data[p] >= 48:
			goto tr1236
		}
		goto st0
tr1236:
//line ragel/datetime.rl:5
 pb = p 
	goto st888
	st888:
		if p++; p == pe {
			goto _test_eof888
		}
	st_case_888:
//line ragel/parse_datetime.go:3565
		if data[p] == 32 {
			goto tr1238
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st889
		}
		goto st0
tr1237:
//line ragel/datetime.rl:5
 pb = p 
	goto st889
	st889:
		if p++; p == pe {
			goto _test_eof889
		}
	st_case_889:
//line ragel/parse_datetime.go:3582
		if data[p] == 32 {
			goto tr1238
		}
		goto st0
tr54:
//line ragel/datetime.rl:5
 pb = p 
	goto st890
tr76:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st890
	st890:
		if p++; p == pe {
			goto _test_eof890
		}
	st_case_890:
//line ragel/parse_datetime.go:3602
		switch data[p] {
		case 32:
			goto tr1220
		case 58:
			goto tr1222
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st886
			}
		case data[p] >= 48:
			goto st885
		}
		goto st0
tr1359:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st31
tr1202:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st31
tr1244:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st31
tr1256:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st31
tr1261:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st31
tr1276:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st31
tr1270:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st31
tr1289:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st31
tr1298:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st31
tr1307:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st31
tr1327:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st31
tr1338:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st31
tr1347:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st31
tr1367:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line ragel/parse_datetime.go:3768
		if data[p] == 50 {
			goto tr76
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr77
			}
		case data[p] >= 48:
			goto tr75
		}
		goto st0
tr1210:
//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1360:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1245:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1262:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1290:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1299:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1308:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1277:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1328:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1271:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1203:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1339:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1348:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
tr1368:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line ragel/parse_datetime.go:3936
		switch data[p] {
		case 47:
			goto st33
		case 95:
			goto st33
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
tr1333:
//line ragel/datetime.rl:5
 pb = p 
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line ragel/parse_datetime.go:3961
		switch data[p] {
		case 47:
			goto st891
		case 95:
			goto st891
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st891
			}
		case data[p] >= 65:
			goto st891
		}
		goto st0
tr1257:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st891
	st891:
		if p++; p == pe {
			goto _test_eof891
		}
	st_case_891:
//line ragel/parse_datetime.go:4009
		switch data[p] {
		case 32:
			goto tr1233
		case 43:
			goto tr1240
		case 45:
			goto tr1241
		case 47:
			goto st891
		case 95:
			goto st891
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st891
			}
		case data[p] >= 65:
			goto st891
		}
		goto st0
tr94:
//line ragel/datetime.rl:5
 pb = p 
	goto st892
	st892:
		if p++; p == pe {
			goto _test_eof892
		}
	st_case_892:
//line ragel/parse_datetime.go:4040
		switch data[p] {
		case 32:
			goto tr1242
		case 43:
			goto tr1243
		case 45:
			goto tr1244
		case 47:
			goto tr1245
		case 58:
			goto tr1247
		case 65:
			goto tr1248
		case 80:
			goto tr1248
		case 90:
			goto tr1249
		case 95:
			goto tr1245
		case 97:
			goto tr1250
		case 112:
			goto tr1250
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st899
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1245
			}
		default:
			goto tr1245
		}
		goto st0
tr1242:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st893
tr1259:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st893
tr1313:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st893
tr1287:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st893
tr1296:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st893
tr1304:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st893
tr1324:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st893
	st893:
		if p++; p == pe {
			goto _test_eof893
		}
	st_case_893:
//line ragel/parse_datetime.go:4150
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st19
		case 45:
			goto st31
		case 47:
			goto tr1210
		case 65:
			goto tr1251
		case 66:
			goto tr1212
		case 80:
			goto tr1252
		case 90:
			goto tr1213
		case 95:
			goto tr1210
		case 97:
			goto tr1253
		case 109:
			goto tr1215
		case 112:
			goto tr1253
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1210
			}
		case data[p] >= 67:
			goto tr1210
		}
		goto st0
tr1251:
//line ragel/datetime.rl:5
 pb = p 
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line ragel/parse_datetime.go:4195
		switch data[p] {
		case 47:
			goto st33
		case 68:
			goto st894
		case 77:
			goto st895
		case 95:
			goto st33
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
	st894:
		if p++; p == pe {
			goto _test_eof894
		}
	st_case_894:
		switch data[p] {
		case 32:
			goto st13
		case 47:
			goto st891
		case 95:
			goto st891
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st891
			}
		case data[p] >= 65:
			goto st891
		}
		goto st0
	st895:
		if p++; p == pe {
			goto _test_eof895
		}
	st_case_895:
		switch data[p] {
		case 32:
			goto tr1254
		case 43:
			goto tr1255
		case 45:
			goto tr1256
		case 47:
			goto tr1257
		case 95:
			goto tr1257
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1257
			}
		case data[p] >= 65:
			goto tr1257
		}
		goto st0
tr1254:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st896
tr1274:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st896
tr1267:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st896
tr1402:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st896
	st896:
		if p++; p == pe {
			goto _test_eof896
		}
	st_case_896:
//line ragel/parse_datetime.go:4362
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st19
		case 45:
			goto st31
		case 47:
			goto tr1210
		case 65:
			goto tr1258
		case 66:
			goto tr1212
		case 90:
			goto tr1213
		case 95:
			goto tr1210
		case 109:
			goto tr1215
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1210
			}
		case data[p] >= 67:
			goto tr1210
		}
		goto st0
tr1258:
//line ragel/datetime.rl:5
 pb = p 
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line ragel/parse_datetime.go:4401
		switch data[p] {
		case 47:
			goto st33
		case 68:
			goto st894
		case 95:
			goto st33
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
tr1212:
//line ragel/datetime.rl:5
 pb = p 
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line ragel/parse_datetime.go:4428
		switch data[p] {
		case 47:
			goto st33
		case 67:
			goto st897
		case 95:
			goto st33
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
	st897:
		if p++; p == pe {
			goto _test_eof897
		}
	st_case_897:
		switch data[p] {
		case 32:
			goto tr1219
		case 47:
			goto st891
		case 95:
			goto st891
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st891
			}
		case data[p] >= 65:
			goto st891
		}
		goto st0
tr1213:
//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1362:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1249:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1265:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1294:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1302:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1311:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1279:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1330:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1273:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1205:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1341:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1350:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
tr1370:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st898
	st898:
		if p++; p == pe {
			goto _test_eof898
		}
	st_case_898:
//line ragel/parse_datetime.go:4623
		switch data[p] {
		case 32:
			goto tr1228
		case 47:
			goto st33
		case 95:
			goto st33
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
tr1215:
//line ragel/datetime.rl:5
 pb = p 
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line ragel/parse_datetime.go:4650
		switch data[p] {
		case 47:
			goto st33
		case 61:
			goto st15
		case 95:
			goto st33
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
tr1252:
//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1248:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1264:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1293:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1301:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1310:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1315:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1329:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line ragel/parse_datetime.go:4758
		switch data[p] {
		case 47:
			goto st33
		case 77:
			goto st895
		case 95:
			goto st33
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
tr1253:
//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1250:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1266:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1295:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1303:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1312:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1316:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1331:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line ragel/parse_datetime.go:4866
		switch data[p] {
		case 47:
			goto st33
		case 95:
			goto st33
		case 109:
			goto st895
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
	st899:
		if p++; p == pe {
			goto _test_eof899
		}
	st_case_899:
		switch data[p] {
		case 32:
			goto tr1259
		case 43:
			goto tr1260
		case 45:
			goto tr1261
		case 47:
			goto tr1262
		case 58:
			goto tr1263
		case 65:
			goto tr1264
		case 80:
			goto tr1264
		case 90:
			goto tr1265
		case 95:
			goto tr1262
		case 97:
			goto tr1266
		case 112:
			goto tr1266
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st40
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1262
			}
		default:
			goto tr1262
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if 48 <= data[p] && data[p] <= 57 {
			goto st900
		}
		goto st0
	st900:
		if p++; p == pe {
			goto _test_eof900
		}
	st_case_900:
		switch data[p] {
		case 32:
			goto tr1267
		case 43:
			goto tr1268
		case 45:
			goto tr1270
		case 47:
			goto tr1271
		case 90:
			goto tr1273
		case 95:
			goto tr1271
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1269
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1271
				}
			case data[p] >= 65:
				goto tr1271
			}
		default:
			goto st42
		}
		goto st0
tr1269:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line ragel/parse_datetime.go:4994
		if 48 <= data[p] && data[p] <= 57 {
			goto tr84
		}
		goto st0
tr84:
//line ragel/datetime.rl:5
 pb = p 
	goto st901
	st901:
		if p++; p == pe {
			goto _test_eof901
		}
	st_case_901:
//line ragel/parse_datetime.go:5008
		switch data[p] {
		case 32:
			goto tr1274
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st902
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st902:
		if p++; p == pe {
			goto _test_eof902
		}
	st_case_902:
		switch data[p] {
		case 32:
			goto tr1274
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st903
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st903:
		if p++; p == pe {
			goto _test_eof903
		}
	st_case_903:
		switch data[p] {
		case 32:
			goto tr1274
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st904
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st904:
		if p++; p == pe {
			goto _test_eof904
		}
	st_case_904:
		switch data[p] {
		case 32:
			goto tr1274
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st905
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st905:
		if p++; p == pe {
			goto _test_eof905
		}
	st_case_905:
		switch data[p] {
		case 32:
			goto tr1274
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st906
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st906:
		if p++; p == pe {
			goto _test_eof906
		}
	st_case_906:
		switch data[p] {
		case 32:
			goto tr1274
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st907
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st907:
		if p++; p == pe {
			goto _test_eof907
		}
	st_case_907:
		switch data[p] {
		case 32:
			goto tr1274
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st908
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st908:
		if p++; p == pe {
			goto _test_eof908
		}
	st_case_908:
		switch data[p] {
		case 32:
			goto tr1274
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st909
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st909:
		if p++; p == pe {
			goto _test_eof909
		}
	st_case_909:
		switch data[p] {
		case 32:
			goto tr1274
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		case data[p] >= 65:
			goto tr1277
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if 48 <= data[p] && data[p] <= 57 {
			goto st910
		}
		goto st0
	st910:
		if p++; p == pe {
			goto _test_eof910
		}
	st_case_910:
		switch data[p] {
		case 32:
			goto tr1267
		case 43:
			goto tr1268
		case 45:
			goto tr1270
		case 47:
			goto tr1271
		case 90:
			goto tr1273
		case 95:
			goto tr1271
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1269
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1271
			}
		default:
			goto tr1271
		}
		goto st0
tr1247:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st43
tr1263:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line ragel/parse_datetime.go:5346
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr87
			}
		case data[p] >= 48:
			goto tr86
		}
		goto st0
tr86:
//line ragel/datetime.rl:5
 pb = p 
	goto st911
	st911:
		if p++; p == pe {
			goto _test_eof911
		}
	st_case_911:
//line ragel/parse_datetime.go:5365
		switch data[p] {
		case 32:
			goto tr1287
		case 43:
			goto tr1288
		case 45:
			goto tr1289
		case 47:
			goto tr1290
		case 58:
			goto tr1292
		case 65:
			goto tr1293
		case 80:
			goto tr1293
		case 90:
			goto tr1294
		case 95:
			goto tr1290
		case 97:
			goto tr1295
		case 112:
			goto tr1295
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st912
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1290
			}
		default:
			goto tr1290
		}
		goto st0
	st912:
		if p++; p == pe {
			goto _test_eof912
		}
	st_case_912:
		switch data[p] {
		case 32:
			goto tr1296
		case 43:
			goto tr1297
		case 45:
			goto tr1298
		case 47:
			goto tr1299
		case 58:
			goto tr1300
		case 65:
			goto tr1301
		case 80:
			goto tr1301
		case 90:
			goto tr1302
		case 95:
			goto tr1299
		case 97:
			goto tr1303
		case 112:
			goto tr1303
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1299
			}
		case data[p] >= 66:
			goto tr1299
		}
		goto st0
tr1292:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st44
tr1300:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line ragel/parse_datetime.go:5458
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr89
			}
		case data[p] >= 48:
			goto tr88
		}
		goto st0
tr88:
//line ragel/datetime.rl:5
 pb = p 
	goto st913
	st913:
		if p++; p == pe {
			goto _test_eof913
		}
	st_case_913:
//line ragel/parse_datetime.go:5477
		switch data[p] {
		case 32:
			goto tr1304
		case 43:
			goto tr1305
		case 45:
			goto tr1307
		case 47:
			goto tr1308
		case 65:
			goto tr1310
		case 80:
			goto tr1310
		case 90:
			goto tr1311
		case 95:
			goto tr1308
		case 97:
			goto tr1312
		case 112:
			goto tr1312
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1306
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1308
				}
			case data[p] >= 66:
				goto tr1308
			}
		default:
			goto st923
		}
		goto st0
tr1306:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st45
tr1326:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line ragel/parse_datetime.go:5535
		if 48 <= data[p] && data[p] <= 57 {
			goto tr90
		}
		goto st0
tr90:
//line ragel/datetime.rl:5
 pb = p 
	goto st914
	st914:
		if p++; p == pe {
			goto _test_eof914
		}
	st_case_914:
//line ragel/parse_datetime.go:5549
		switch data[p] {
		case 32:
			goto tr1313
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 65:
			goto tr1315
		case 80:
			goto tr1315
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		case 97:
			goto tr1316
		case 112:
			goto tr1316
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st915
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st915:
		if p++; p == pe {
			goto _test_eof915
		}
	st_case_915:
		switch data[p] {
		case 32:
			goto tr1313
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 65:
			goto tr1315
		case 80:
			goto tr1315
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		case 97:
			goto tr1316
		case 112:
			goto tr1316
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st916
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st916:
		if p++; p == pe {
			goto _test_eof916
		}
	st_case_916:
		switch data[p] {
		case 32:
			goto tr1313
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 65:
			goto tr1315
		case 80:
			goto tr1315
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		case 97:
			goto tr1316
		case 112:
			goto tr1316
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st917
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st917:
		if p++; p == pe {
			goto _test_eof917
		}
	st_case_917:
		switch data[p] {
		case 32:
			goto tr1313
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 65:
			goto tr1315
		case 80:
			goto tr1315
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		case 97:
			goto tr1316
		case 112:
			goto tr1316
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st918
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st918:
		if p++; p == pe {
			goto _test_eof918
		}
	st_case_918:
		switch data[p] {
		case 32:
			goto tr1313
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 65:
			goto tr1315
		case 80:
			goto tr1315
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		case 97:
			goto tr1316
		case 112:
			goto tr1316
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st919
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st919:
		if p++; p == pe {
			goto _test_eof919
		}
	st_case_919:
		switch data[p] {
		case 32:
			goto tr1313
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 65:
			goto tr1315
		case 80:
			goto tr1315
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		case 97:
			goto tr1316
		case 112:
			goto tr1316
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st920
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st920:
		if p++; p == pe {
			goto _test_eof920
		}
	st_case_920:
		switch data[p] {
		case 32:
			goto tr1313
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 65:
			goto tr1315
		case 80:
			goto tr1315
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		case 97:
			goto tr1316
		case 112:
			goto tr1316
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st921
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st921:
		if p++; p == pe {
			goto _test_eof921
		}
	st_case_921:
		switch data[p] {
		case 32:
			goto tr1313
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 65:
			goto tr1315
		case 80:
			goto tr1315
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		case 97:
			goto tr1316
		case 112:
			goto tr1316
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st922
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		default:
			goto tr1277
		}
		goto st0
	st922:
		if p++; p == pe {
			goto _test_eof922
		}
	st_case_922:
		switch data[p] {
		case 32:
			goto tr1313
		case 43:
			goto tr1275
		case 45:
			goto tr1276
		case 47:
			goto tr1277
		case 65:
			goto tr1315
		case 80:
			goto tr1315
		case 90:
			goto tr1279
		case 95:
			goto tr1277
		case 97:
			goto tr1316
		case 112:
			goto tr1316
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1277
			}
		case data[p] >= 66:
			goto tr1277
		}
		goto st0
	st923:
		if p++; p == pe {
			goto _test_eof923
		}
	st_case_923:
		switch data[p] {
		case 32:
			goto tr1324
		case 43:
			goto tr1325
		case 45:
			goto tr1327
		case 47:
			goto tr1328
		case 65:
			goto tr1329
		case 80:
			goto tr1329
		case 90:
			goto tr1330
		case 95:
			goto tr1328
		case 97:
			goto tr1331
		case 112:
			goto tr1331
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1326
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1328
			}
		default:
			goto tr1328
		}
		goto st0
tr89:
//line ragel/datetime.rl:5
 pb = p 
	goto st924
	st924:
		if p++; p == pe {
			goto _test_eof924
		}
	st_case_924:
//line ragel/parse_datetime.go:5950
		switch data[p] {
		case 32:
			goto tr1304
		case 43:
			goto tr1305
		case 45:
			goto tr1307
		case 47:
			goto tr1308
		case 65:
			goto tr1310
		case 80:
			goto tr1310
		case 90:
			goto tr1311
		case 95:
			goto tr1308
		case 97:
			goto tr1312
		case 112:
			goto tr1312
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1306
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1308
			}
		default:
			goto tr1308
		}
		goto st0
tr87:
//line ragel/datetime.rl:5
 pb = p 
	goto st925
	st925:
		if p++; p == pe {
			goto _test_eof925
		}
	st_case_925:
//line ragel/parse_datetime.go:5995
		switch data[p] {
		case 32:
			goto tr1287
		case 43:
			goto tr1288
		case 45:
			goto tr1289
		case 47:
			goto tr1290
		case 58:
			goto tr1292
		case 65:
			goto tr1293
		case 80:
			goto tr1293
		case 90:
			goto tr1294
		case 95:
			goto tr1290
		case 97:
			goto tr1295
		case 112:
			goto tr1295
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1290
			}
		case data[p] >= 66:
			goto tr1290
		}
		goto st0
tr95:
//line ragel/datetime.rl:5
 pb = p 
	goto st926
	st926:
		if p++; p == pe {
			goto _test_eof926
		}
	st_case_926:
//line ragel/parse_datetime.go:6038
		switch data[p] {
		case 32:
			goto tr1242
		case 43:
			goto tr1243
		case 45:
			goto tr1244
		case 47:
			goto tr1245
		case 58:
			goto tr1247
		case 65:
			goto tr1248
		case 80:
			goto tr1248
		case 90:
			goto tr1249
		case 95:
			goto tr1245
		case 97:
			goto tr1250
		case 112:
			goto tr1250
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st899
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1245
				}
			case data[p] >= 66:
				goto tr1245
			}
		default:
			goto st46
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if 48 <= data[p] && data[p] <= 57 {
			goto st40
		}
		goto st0
tr96:
//line ragel/datetime.rl:5
 pb = p 
	goto st927
	st927:
		if p++; p == pe {
			goto _test_eof927
		}
	st_case_927:
//line ragel/parse_datetime.go:6099
		switch data[p] {
		case 32:
			goto tr1242
		case 43:
			goto tr1243
		case 45:
			goto tr1244
		case 47:
			goto tr1245
		case 58:
			goto tr1247
		case 65:
			goto tr1248
		case 80:
			goto tr1248
		case 90:
			goto tr1249
		case 95:
			goto tr1245
		case 97:
			goto tr1250
		case 112:
			goto tr1250
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st46
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1245
			}
		default:
			goto tr1245
		}
		goto st0
tr1211:
//line ragel/datetime.rl:5
 pb = p 
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line ragel/parse_datetime.go:6146
		switch data[p] {
		case 47:
			goto st33
		case 68:
			goto st894
		case 84:
			goto st48
		case 95:
			goto st33
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 32:
			goto st49
		case 47:
			goto st891
		case 95:
			goto st891
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st891
			}
		case data[p] >= 65:
			goto st891
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 50 {
			goto tr95
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr96
			}
		case data[p] >= 48:
			goto tr94
		}
		goto st0
tr1214:
//line ragel/datetime.rl:5
 pb = p 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line ragel/parse_datetime.go:6214
		switch data[p] {
		case 47:
			goto st33
		case 95:
			goto st33
		case 116:
			goto st48
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		case data[p] >= 65:
			goto st33
		}
		goto st0
tr1358:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st51
tr1201:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st51
tr1337:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st51
tr1346:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st51
tr1366:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line ragel/parse_datetime.go:6273
		if data[p] == 32 {
			goto st49
		}
		goto st0
tr1354:
//line ragel/datetime.rl:5
 pb = p 
	goto st928
tr1361:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st928
tr1204:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st928
tr1340:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st928
tr1349:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st928
tr1369:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st928
	st928:
		if p++; p == pe {
			goto _test_eof928
		}
	st_case_928:
//line ragel/parse_datetime.go:6333
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st19
		case 45:
			goto st31
		case 47:
			goto tr1333
		case 50:
			goto tr95
		case 90:
			goto tr1334
		case 95:
			goto tr1333
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr94
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1333
				}
			case data[p] >= 65:
				goto tr1333
			}
		default:
			goto tr96
		}
		goto st0
tr1334:
//line ragel/datetime.rl:5
 pb = p 
	goto st929
	st929:
		if p++; p == pe {
			goto _test_eof929
		}
	st_case_929:
//line ragel/parse_datetime.go:6377
		switch data[p] {
		case 32:
			goto tr1228
		case 47:
			goto st891
		case 95:
			goto st891
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st891
			}
		case data[p] >= 65:
			goto st891
		}
		goto st0
tr1355:
//line ragel/datetime.rl:5
 pb = p 
	goto st52
tr1363:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st52
tr1206:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st52
tr1342:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st52
tr1351:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st52
tr1371:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line ragel/parse_datetime.go:6450
		switch data[p] {
		case 47:
			goto st33
		case 50:
			goto tr95
		case 95:
			goto st33
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr94
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st33
				}
			case data[p] >= 65:
				goto st33
			}
		default:
			goto tr96
		}
		goto st0
tr40:
//line ragel/datetime.rl:5
 pb = p 
	goto st930
	st930:
		if p++; p == pe {
			goto _test_eof930
		}
	st_case_930:
//line ragel/parse_datetime.go:6486
		switch data[p] {
		case 32:
			goto tr1199
		case 43:
			goto tr1200
		case 44:
			goto tr1201
		case 45:
			goto tr1202
		case 47:
			goto tr1203
		case 84:
			goto tr1204
		case 90:
			goto tr1205
		case 95:
			goto tr1206
		case 116:
			goto tr1206
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st867
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1203
			}
		default:
			goto tr1203
		}
		goto st0
tr41:
//line ragel/datetime.rl:5
 pb = p 
	goto st931
	st931:
		if p++; p == pe {
			goto _test_eof931
		}
	st_case_931:
//line ragel/parse_datetime.go:6529
		switch data[p] {
		case 32:
			goto tr1199
		case 43:
			goto tr1200
		case 44:
			goto tr1201
		case 45:
			goto tr1202
		case 47:
			goto tr1203
		case 84:
			goto tr1204
		case 90:
			goto tr1205
		case 95:
			goto tr1206
		case 116:
			goto tr1206
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 49 {
				goto st867
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1203
			}
		default:
			goto tr1203
		}
		goto st0
tr27:
//line ragel/datetime.rl:5
 pb = p 
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line ragel/parse_datetime.go:6572
		if data[p] == 32 {
			goto tr38
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 50 {
				goto st8
			}
		case data[p] >= 45:
			goto tr38
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 112:
			goto st55
		case 117:
			goto st60
		}
		goto st0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 114 {
			goto st56
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 32:
			goto tr100
		case 46:
			goto tr101
		case 105:
			goto st58
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr100
		}
		goto st0
tr101:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st57
tr107:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st57
tr114:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st57
tr123:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st57
tr133:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st57
tr141:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st57
tr144:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st57
tr150:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st57
tr154:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st57
tr158:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st57
tr167:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st57
tr175:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line ragel/parse_datetime.go:6676
		switch data[p] {
		case 32:
			goto st9
		case 48:
			goto tr39
		case 51:
			goto tr41
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st9
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr42
			}
		default:
			goto tr40
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 108 {
			goto st59
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 32:
			goto tr100
		case 46:
			goto tr101
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr100
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if data[p] == 103 {
			goto st61
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 32:
			goto tr106
		case 46:
			goto tr107
		case 117:
			goto st62
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr106
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if data[p] == 115 {
			goto st63
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 116 {
			goto st64
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 32:
			goto tr106
		case 46:
			goto tr107
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr106
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 101 {
			goto st66
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		if data[p] == 99 {
			goto st67
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 32:
			goto tr113
		case 46:
			goto tr114
		case 101:
			goto st68
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr113
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if data[p] == 109 {
			goto st69
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		if data[p] == 98 {
			goto st70
		}
		goto st0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		if data[p] == 101 {
			goto st71
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		if data[p] == 114 {
			goto st72
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 32:
			goto tr113
		case 46:
			goto tr114
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr113
		}
		goto st0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 101 {
			goto st74
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		if data[p] == 98 {
			goto st75
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 32:
			goto tr122
		case 46:
			goto tr123
		case 114:
			goto st76
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr122
		}
		goto st0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		if data[p] == 117 {
			goto st77
		}
		goto st0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		if data[p] == 97 {
			goto st78
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		if data[p] == 114 {
			goto st79
		}
		goto st0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		if data[p] == 121 {
			goto st80
		}
		goto st0
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 32:
			goto tr122
		case 46:
			goto tr123
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr122
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 97:
			goto st82
		case 117:
			goto st88
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		if data[p] == 110 {
			goto st83
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 32:
			goto tr132
		case 46:
			goto tr133
		case 117:
			goto st84
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr132
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 97 {
			goto st85
		}
		goto st0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		if data[p] == 114 {
			goto st86
		}
		goto st0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		if data[p] == 121 {
			goto st87
		}
		goto st0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 32:
			goto tr132
		case 46:
			goto tr133
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr132
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 108:
			goto st89
		case 110:
			goto st91
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 32:
			goto tr140
		case 46:
			goto tr141
		case 121:
			goto st90
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr140
		}
		goto st0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 32:
			goto tr140
		case 46:
			goto tr141
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr140
		}
		goto st0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 32:
			goto tr143
		case 46:
			goto tr144
		case 101:
			goto st92
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr143
		}
		goto st0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 32:
			goto tr143
		case 46:
			goto tr144
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr143
		}
		goto st0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		if data[p] == 97 {
			goto st94
		}
		goto st0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 114:
			goto st95
		case 121:
			goto st98
		}
		goto st0
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 32:
			goto tr149
		case 46:
			goto tr150
		case 99:
			goto st96
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr149
		}
		goto st0
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		if data[p] == 104 {
			goto st97
		}
		goto st0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 32:
			goto tr149
		case 46:
			goto tr150
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr149
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 32:
			goto tr153
		case 46:
			goto tr154
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr153
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		if data[p] == 111 {
			goto st100
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		if data[p] == 118 {
			goto st101
		}
		goto st0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 32:
			goto tr157
		case 46:
			goto tr158
		case 101:
			goto st102
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr157
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		if data[p] == 109 {
			goto st103
		}
		goto st0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		if data[p] == 98 {
			goto st104
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		if data[p] == 101 {
			goto st105
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		if data[p] == 114 {
			goto st106
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 32:
			goto tr157
		case 46:
			goto tr158
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr157
		}
		goto st0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		if data[p] == 99 {
			goto st108
		}
		goto st0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		if data[p] == 116 {
			goto st109
		}
		goto st0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 32:
			goto tr166
		case 46:
			goto tr167
		case 111:
			goto st110
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr166
		}
		goto st0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		if data[p] == 98 {
			goto st111
		}
		goto st0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		if data[p] == 101 {
			goto st112
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		if data[p] == 114 {
			goto st113
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 32:
			goto tr166
		case 46:
			goto tr167
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr166
		}
		goto st0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		if data[p] == 101 {
			goto st115
		}
		goto st0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		if data[p] == 112 {
			goto st116
		}
		goto st0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 32:
			goto tr174
		case 46:
			goto tr175
		case 116:
			goto st117
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr174
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 32:
			goto tr174
		case 46:
			goto tr175
		case 101:
			goto st118
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr174
		}
		goto st0
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		if data[p] == 109 {
			goto st119
		}
		goto st0
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		if data[p] == 98 {
			goto st120
		}
		goto st0
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		if data[p] == 101 {
			goto st121
		}
		goto st0
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		if data[p] == 114 {
			goto st122
		}
		goto st0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 32:
			goto tr174
		case 46:
			goto tr175
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr174
		}
		goto st0
tr23:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
//line ragel/parse_datetime.go:7463
		switch data[p] {
		case 48:
			goto tr182
		case 49:
			goto tr183
		case 65:
			goto st54
		case 68:
			goto st65
		case 70:
			goto st73
		case 74:
			goto st81
		case 77:
			goto st93
		case 78:
			goto st99
		case 79:
			goto st107
		case 83:
			goto st114
		case 97:
			goto st54
		case 100:
			goto st65
		case 102:
			goto st73
		case 106:
			goto st81
		case 109:
			goto st93
		case 110:
			goto st99
		case 111:
			goto st107
		case 115:
			goto st114
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr184
		}
		goto st0
tr182:
//line ragel/datetime.rl:5
 pb = p 
	goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line ragel/parse_datetime.go:7515
		if data[p] == 48 {
			goto st125
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st126
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		if 48 <= data[p] && data[p] <= 57 {
			goto st932
		}
		goto st0
	st932:
		if p++; p == pe {
			goto _test_eof932
		}
	st_case_932:
		switch data[p] {
		case 32:
			goto tr1335
		case 43:
			goto tr1336
		case 44:
			goto tr1337
		case 45:
			goto tr1338
		case 47:
			goto tr1339
		case 84:
			goto tr1340
		case 90:
			goto tr1341
		case 95:
			goto tr1342
		case 116:
			goto tr1342
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1339
			}
		case data[p] >= 65:
			goto tr1339
		}
		goto st0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		if data[p] == 32 {
			goto tr38
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st932
			}
		case data[p] >= 45:
			goto tr38
		}
		goto st0
tr183:
//line ragel/datetime.rl:5
 pb = p 
	goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line ragel/parse_datetime.go:7592
		if data[p] == 32 {
			goto tr38
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr38
			}
		case data[p] > 50:
			if 51 <= data[p] && data[p] <= 57 {
				goto st125
			}
		default:
			goto st126
		}
		goto st0
tr184:
//line ragel/datetime.rl:5
 pb = p 
	goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
//line ragel/parse_datetime.go:7618
		if data[p] == 32 {
			goto tr38
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st125
			}
		case data[p] >= 45:
			goto tr38
		}
		goto st0
tr24:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st129
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
//line ragel/parse_datetime.go:7644
		if 48 <= data[p] && data[p] <= 57 {
			goto st130
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		if 48 <= data[p] && data[p] <= 57 {
			goto st933
		}
		goto st0
	st933:
		if p++; p == pe {
			goto _test_eof933
		}
	st_case_933:
		switch data[p] {
		case 32:
			goto tr1335
		case 43:
			goto tr1336
		case 44:
			goto tr1337
		case 45:
			goto tr1338
		case 47:
			goto tr1339
		case 84:
			goto tr1340
		case 90:
			goto tr1341
		case 95:
			goto tr1342
		case 116:
			goto tr1342
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st934
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1339
			}
		default:
			goto tr1339
		}
		goto st0
	st934:
		if p++; p == pe {
			goto _test_eof934
		}
	st_case_934:
		switch data[p] {
		case 32:
			goto tr1344
		case 43:
			goto tr1345
		case 44:
			goto tr1346
		case 45:
			goto tr1347
		case 47:
			goto tr1348
		case 84:
			goto tr1349
		case 90:
			goto tr1350
		case 95:
			goto tr1351
		case 116:
			goto tr1351
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1348
			}
		case data[p] >= 65:
			goto tr1348
		}
		goto st0
tr25:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
//line ragel/parse_datetime.go:7741
		if data[p] == 185 {
			goto st132
		}
		goto st0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		if data[p] == 180 {
			goto st133
		}
		goto st0
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 48:
			goto tr192
		case 49:
			goto tr193
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr194
		}
		goto st0
tr192:
//line ragel/datetime.rl:5
 pb = p 
	goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line ragel/parse_datetime.go:7779
		if 49 <= data[p] && data[p] <= 57 {
			goto st135
		}
		goto st0
tr194:
//line ragel/datetime.rl:5
 pb = p 
	goto st135
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
//line ragel/parse_datetime.go:7793
		if data[p] == 230 {
			goto tr196
		}
		goto st0
tr196:
//line ragel/datetime.rl:9

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line ragel/parse_datetime.go:7809
		if data[p] == 156 {
			goto st137
		}
		goto st0
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		if data[p] == 136 {
			goto st138
		}
		goto st0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 48:
			goto tr199
		case 51:
			goto tr201
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr202
			}
		case data[p] >= 49:
			goto tr200
		}
		goto st0
tr199:
//line ragel/datetime.rl:5
 pb = p 
	goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
//line ragel/parse_datetime.go:7852
		if 49 <= data[p] && data[p] <= 57 {
			goto st140
		}
		goto st0
tr202:
//line ragel/datetime.rl:5
 pb = p 
	goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
//line ragel/parse_datetime.go:7866
		if data[p] == 230 {
			goto tr204
		}
		goto st0
tr204:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
//line ragel/parse_datetime.go:7887
		if data[p] == 151 {
			goto st142
		}
		goto st0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		if data[p] == 165 {
			goto st935
		}
		goto st0
	st935:
		if p++; p == pe {
			goto _test_eof935
		}
	st_case_935:
		switch data[p] {
		case 32:
			goto st868
		case 43:
			goto st19
		case 44:
			goto st51
		case 45:
			goto st31
		case 47:
			goto tr1210
		case 84:
			goto tr1354
		case 90:
			goto tr1213
		case 95:
			goto tr1355
		case 116:
			goto tr1355
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1210
			}
		case data[p] >= 65:
			goto tr1210
		}
		goto st0
tr200:
//line ragel/datetime.rl:5
 pb = p 
	goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
//line ragel/parse_datetime.go:7944
		if data[p] == 230 {
			goto tr204
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st140
		}
		goto st0
tr201:
//line ragel/datetime.rl:5
 pb = p 
	goto st144
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
//line ragel/parse_datetime.go:7961
		if data[p] == 230 {
			goto tr204
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st140
		}
		goto st0
tr193:
//line ragel/datetime.rl:5
 pb = p 
	goto st145
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
//line ragel/parse_datetime.go:7978
		if data[p] == 230 {
			goto tr196
		}
		if 48 <= data[p] && data[p] <= 50 {
			goto st135
		}
		goto st0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		if data[p] == 32 {
			goto tr207
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] >= 45:
			goto tr208
		}
		goto st0
tr207:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
//line ragel/parse_datetime.go:8019
		switch data[p] {
		case 65:
			goto st154
		case 68:
			goto st167
		case 70:
			goto st175
		case 74:
			goto st183
		case 77:
			goto st195
		case 78:
			goto st201
		case 79:
			goto st209
		case 83:
			goto st216
		case 97:
			goto st154
		case 100:
			goto st167
		case 102:
			goto st175
		case 106:
			goto st183
		case 109:
			goto st195
		case 110:
			goto st201
		case 111:
			goto st209
		case 115:
			goto st216
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr209
		}
		goto st0
tr209:
//line ragel/datetime.rl:5
 pb = p 
	goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line ragel/parse_datetime.go:8067
		if data[p] == 32 {
			goto tr218
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st153
			}
		case data[p] >= 45:
			goto tr218
		}
		goto st0
tr228:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st149
tr238:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st149
tr246:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st149
tr256:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st149
tr267:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st149
tr276:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st149
tr280:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st149
tr287:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st149
tr292:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st149
tr297:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st149
tr307:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st149
tr316:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st149
tr418:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st149
tr218:
//line ragel/datetime.rl:29

    if st.Day == 0 {
        value, _ := strconv.Atoi(data[pb:pb+2])
        st.Day = value
    }else {
        value, _ := strconv.Atoi(data[pb:pb+2])
        max_v := max(st.Day, value)
        min_v := min(st.Day, value)
        st.Day, st.Month = max_v, min_v
        if st.Month > 12 {
            err = errors.New("month value overflow")
        } else if st.Day <=12 && st.Day != st.Month {
            err = errors.New("ambiguous day/month")
        }
    }

	goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
//line ragel/parse_datetime.go:8163
		if 48 <= data[p] && data[p] <= 57 {
			goto tr220
		}
		goto st0
tr220:
//line ragel/datetime.rl:5
 pb = p 
	goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line ragel/parse_datetime.go:8177
		if 48 <= data[p] && data[p] <= 57 {
			goto st151
		}
		goto st0
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		if 48 <= data[p] && data[p] <= 57 {
			goto st152
		}
		goto st0
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		if 48 <= data[p] && data[p] <= 57 {
			goto st936
		}
		goto st0
	st936:
		if p++; p == pe {
			goto _test_eof936
		}
	st_case_936:
		switch data[p] {
		case 32:
			goto tr1356
		case 43:
			goto tr1357
		case 44:
			goto tr1358
		case 45:
			goto tr1359
		case 47:
			goto tr1360
		case 84:
			goto tr1361
		case 90:
			goto tr1362
		case 95:
			goto tr1363
		case 116:
			goto tr1363
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1360
			}
		case data[p] >= 65:
			goto tr1360
		}
		goto st0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		if data[p] == 32 {
			goto tr218
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr218
		}
		goto st0
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 112:
			goto st155
		case 117:
			goto st162
		}
		goto st0
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		if data[p] == 114 {
			goto st156
		}
		goto st0
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 32:
			goto tr227
		case 46:
			goto tr229
		case 105:
			goto st160
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr228
		}
		goto st0
tr227:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st157
tr237:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st157
tr245:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st157
tr255:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st157
tr266:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st157
tr275:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st157
tr279:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st157
tr286:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st157
tr291:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st157
tr296:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st157
tr306:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st157
tr315:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st157
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
//line ragel/parse_datetime.go:8337
		if 48 <= data[p] && data[p] <= 57 {
			goto tr231
		}
		goto st0
tr231:
//line ragel/datetime.rl:5
 pb = p 
	goto st158
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
//line ragel/parse_datetime.go:8351
		if 48 <= data[p] && data[p] <= 57 {
			goto st937
		}
		goto st0
	st937:
		if p++; p == pe {
			goto _test_eof937
		}
	st_case_937:
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 44:
			goto tr1366
		case 45:
			goto tr1367
		case 47:
			goto tr1368
		case 84:
			goto tr1369
		case 90:
			goto tr1370
		case 95:
			goto tr1371
		case 116:
			goto tr1371
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st152
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1368
			}
		default:
			goto tr1368
		}
		goto st0
tr229:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st159
tr239:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st159
tr247:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st159
tr257:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st159
tr268:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st159
tr277:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st159
tr281:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st159
tr288:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st159
tr293:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st159
tr298:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st159
tr308:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st159
tr317:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st159
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
//line ragel/parse_datetime.go:8447
		if data[p] == 32 {
			goto st157
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr220
			}
		case data[p] >= 45:
			goto st149
		}
		goto st0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		if data[p] == 108 {
			goto st161
		}
		goto st0
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 32:
			goto tr227
		case 46:
			goto tr229
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr228
		}
		goto st0
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		if data[p] == 103 {
			goto st163
		}
		goto st0
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 32:
			goto tr237
		case 46:
			goto tr239
		case 117:
			goto st164
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr238
		}
		goto st0
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		if data[p] == 115 {
			goto st165
		}
		goto st0
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		if data[p] == 116 {
			goto st166
		}
		goto st0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 32:
			goto tr237
		case 46:
			goto tr239
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr238
		}
		goto st0
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		if data[p] == 101 {
			goto st168
		}
		goto st0
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		if data[p] == 99 {
			goto st169
		}
		goto st0
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch data[p] {
		case 32:
			goto tr245
		case 46:
			goto tr247
		case 101:
			goto st170
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr246
		}
		goto st0
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		if data[p] == 109 {
			goto st171
		}
		goto st0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		if data[p] == 98 {
			goto st172
		}
		goto st0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		if data[p] == 101 {
			goto st173
		}
		goto st0
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		if data[p] == 114 {
			goto st174
		}
		goto st0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 32:
			goto tr245
		case 46:
			goto tr247
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr246
		}
		goto st0
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		if data[p] == 101 {
			goto st176
		}
		goto st0
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		if data[p] == 98 {
			goto st177
		}
		goto st0
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 32:
			goto tr255
		case 46:
			goto tr257
		case 114:
			goto st178
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr256
		}
		goto st0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		if data[p] == 117 {
			goto st179
		}
		goto st0
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		if data[p] == 97 {
			goto st180
		}
		goto st0
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		if data[p] == 114 {
			goto st181
		}
		goto st0
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		if data[p] == 121 {
			goto st182
		}
		goto st0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		switch data[p] {
		case 32:
			goto tr255
		case 46:
			goto tr257
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr256
		}
		goto st0
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		switch data[p] {
		case 97:
			goto st184
		case 117:
			goto st190
		}
		goto st0
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		if data[p] == 110 {
			goto st185
		}
		goto st0
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		switch data[p] {
		case 32:
			goto tr266
		case 46:
			goto tr268
		case 117:
			goto st186
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr267
		}
		goto st0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		if data[p] == 97 {
			goto st187
		}
		goto st0
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		if data[p] == 114 {
			goto st188
		}
		goto st0
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		if data[p] == 121 {
			goto st189
		}
		goto st0
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 32:
			goto tr266
		case 46:
			goto tr268
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr267
		}
		goto st0
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		switch data[p] {
		case 108:
			goto st191
		case 110:
			goto st193
		}
		goto st0
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch data[p] {
		case 32:
			goto tr275
		case 46:
			goto tr277
		case 121:
			goto st192
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr276
		}
		goto st0
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		switch data[p] {
		case 32:
			goto tr275
		case 46:
			goto tr277
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr276
		}
		goto st0
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		switch data[p] {
		case 32:
			goto tr279
		case 46:
			goto tr281
		case 101:
			goto st194
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr280
		}
		goto st0
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		switch data[p] {
		case 32:
			goto tr279
		case 46:
			goto tr281
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr280
		}
		goto st0
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		if data[p] == 97 {
			goto st196
		}
		goto st0
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 114:
			goto st197
		case 121:
			goto st200
		}
		goto st0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 32:
			goto tr286
		case 46:
			goto tr288
		case 99:
			goto st198
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr287
		}
		goto st0
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		if data[p] == 104 {
			goto st199
		}
		goto st0
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		switch data[p] {
		case 32:
			goto tr286
		case 46:
			goto tr288
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr287
		}
		goto st0
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		switch data[p] {
		case 32:
			goto tr291
		case 46:
			goto tr293
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr292
		}
		goto st0
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		if data[p] == 111 {
			goto st202
		}
		goto st0
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		if data[p] == 118 {
			goto st203
		}
		goto st0
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		switch data[p] {
		case 32:
			goto tr296
		case 46:
			goto tr298
		case 101:
			goto st204
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr297
		}
		goto st0
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		if data[p] == 109 {
			goto st205
		}
		goto st0
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		if data[p] == 98 {
			goto st206
		}
		goto st0
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		if data[p] == 101 {
			goto st207
		}
		goto st0
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		if data[p] == 114 {
			goto st208
		}
		goto st0
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		switch data[p] {
		case 32:
			goto tr296
		case 46:
			goto tr298
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr297
		}
		goto st0
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		if data[p] == 99 {
			goto st210
		}
		goto st0
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		if data[p] == 116 {
			goto st211
		}
		goto st0
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		switch data[p] {
		case 32:
			goto tr306
		case 46:
			goto tr308
		case 111:
			goto st212
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr307
		}
		goto st0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		if data[p] == 98 {
			goto st213
		}
		goto st0
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		if data[p] == 101 {
			goto st214
		}
		goto st0
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		if data[p] == 114 {
			goto st215
		}
		goto st0
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		switch data[p] {
		case 32:
			goto tr306
		case 46:
			goto tr308
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr307
		}
		goto st0
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		if data[p] == 101 {
			goto st217
		}
		goto st0
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		if data[p] == 112 {
			goto st218
		}
		goto st0
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 32:
			goto tr315
		case 46:
			goto tr317
		case 116:
			goto st219
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr316
		}
		goto st0
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		switch data[p] {
		case 32:
			goto tr315
		case 46:
			goto tr317
		case 101:
			goto st220
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr316
		}
		goto st0
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		if data[p] == 109 {
			goto st221
		}
		goto st0
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		if data[p] == 98 {
			goto st222
		}
		goto st0
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		if data[p] == 101 {
			goto st223
		}
		goto st0
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		if data[p] == 114 {
			goto st224
		}
		goto st0
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 32:
			goto tr315
		case 46:
			goto tr317
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr316
		}
		goto st0
tr208:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st225
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
//line ragel/parse_datetime.go:9230
		switch data[p] {
		case 65:
			goto st226
		case 68:
			goto st237
		case 70:
			goto st245
		case 74:
			goto st253
		case 77:
			goto st265
		case 78:
			goto st271
		case 79:
			goto st279
		case 83:
			goto st286
		case 97:
			goto st226
		case 100:
			goto st237
		case 102:
			goto st245
		case 106:
			goto st253
		case 109:
			goto st265
		case 110:
			goto st271
		case 111:
			goto st279
		case 115:
			goto st286
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr209
		}
		goto st0
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 112:
			goto st227
		case 117:
			goto st232
		}
		goto st0
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		if data[p] == 114 {
			goto st228
		}
		goto st0
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		switch data[p] {
		case 32:
			goto tr228
		case 46:
			goto tr335
		case 105:
			goto st230
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr228
		}
		goto st0
tr335:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st229
tr339:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st229
tr345:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st229
tr353:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st229
tr362:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st229
tr369:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st229
tr371:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st229
tr376:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st229
tr379:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st229
tr382:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st229
tr390:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st229
tr397:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line ragel/parse_datetime.go:9360
		if data[p] == 32 {
			goto st149
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr220
			}
		case data[p] >= 45:
			goto st149
		}
		goto st0
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		if data[p] == 108 {
			goto st231
		}
		goto st0
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		switch data[p] {
		case 32:
			goto tr228
		case 46:
			goto tr335
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr228
		}
		goto st0
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		if data[p] == 103 {
			goto st233
		}
		goto st0
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		switch data[p] {
		case 32:
			goto tr238
		case 46:
			goto tr339
		case 117:
			goto st234
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr238
		}
		goto st0
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		if data[p] == 115 {
			goto st235
		}
		goto st0
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		if data[p] == 116 {
			goto st236
		}
		goto st0
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		switch data[p] {
		case 32:
			goto tr238
		case 46:
			goto tr339
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr238
		}
		goto st0
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		if data[p] == 101 {
			goto st238
		}
		goto st0
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		if data[p] == 99 {
			goto st239
		}
		goto st0
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		switch data[p] {
		case 32:
			goto tr246
		case 46:
			goto tr345
		case 101:
			goto st240
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr246
		}
		goto st0
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		if data[p] == 109 {
			goto st241
		}
		goto st0
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		if data[p] == 98 {
			goto st242
		}
		goto st0
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		if data[p] == 101 {
			goto st243
		}
		goto st0
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		if data[p] == 114 {
			goto st244
		}
		goto st0
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		switch data[p] {
		case 32:
			goto tr246
		case 46:
			goto tr345
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr246
		}
		goto st0
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		if data[p] == 101 {
			goto st246
		}
		goto st0
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		if data[p] == 98 {
			goto st247
		}
		goto st0
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		switch data[p] {
		case 32:
			goto tr256
		case 46:
			goto tr353
		case 114:
			goto st248
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr256
		}
		goto st0
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		if data[p] == 117 {
			goto st249
		}
		goto st0
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		if data[p] == 97 {
			goto st250
		}
		goto st0
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		if data[p] == 114 {
			goto st251
		}
		goto st0
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		if data[p] == 121 {
			goto st252
		}
		goto st0
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		switch data[p] {
		case 32:
			goto tr256
		case 46:
			goto tr353
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr256
		}
		goto st0
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		switch data[p] {
		case 97:
			goto st254
		case 117:
			goto st260
		}
		goto st0
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		if data[p] == 110 {
			goto st255
		}
		goto st0
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 32:
			goto tr267
		case 46:
			goto tr362
		case 117:
			goto st256
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr267
		}
		goto st0
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		if data[p] == 97 {
			goto st257
		}
		goto st0
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		if data[p] == 114 {
			goto st258
		}
		goto st0
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		if data[p] == 121 {
			goto st259
		}
		goto st0
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		switch data[p] {
		case 32:
			goto tr267
		case 46:
			goto tr362
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr267
		}
		goto st0
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 108:
			goto st261
		case 110:
			goto st263
		}
		goto st0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		switch data[p] {
		case 32:
			goto tr276
		case 46:
			goto tr369
		case 121:
			goto st262
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr276
		}
		goto st0
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		switch data[p] {
		case 32:
			goto tr276
		case 46:
			goto tr369
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr276
		}
		goto st0
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		switch data[p] {
		case 32:
			goto tr280
		case 46:
			goto tr371
		case 101:
			goto st264
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr280
		}
		goto st0
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 32:
			goto tr280
		case 46:
			goto tr371
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr280
		}
		goto st0
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		if data[p] == 97 {
			goto st266
		}
		goto st0
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
		switch data[p] {
		case 114:
			goto st267
		case 121:
			goto st270
		}
		goto st0
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		switch data[p] {
		case 32:
			goto tr287
		case 46:
			goto tr376
		case 99:
			goto st268
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr287
		}
		goto st0
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		if data[p] == 104 {
			goto st269
		}
		goto st0
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		switch data[p] {
		case 32:
			goto tr287
		case 46:
			goto tr376
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr287
		}
		goto st0
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		switch data[p] {
		case 32:
			goto tr292
		case 46:
			goto tr379
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr292
		}
		goto st0
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		if data[p] == 111 {
			goto st272
		}
		goto st0
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		if data[p] == 118 {
			goto st273
		}
		goto st0
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
		switch data[p] {
		case 32:
			goto tr297
		case 46:
			goto tr382
		case 101:
			goto st274
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr297
		}
		goto st0
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		if data[p] == 109 {
			goto st275
		}
		goto st0
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		if data[p] == 98 {
			goto st276
		}
		goto st0
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		if data[p] == 101 {
			goto st277
		}
		goto st0
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		if data[p] == 114 {
			goto st278
		}
		goto st0
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		switch data[p] {
		case 32:
			goto tr297
		case 46:
			goto tr382
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr297
		}
		goto st0
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		if data[p] == 99 {
			goto st280
		}
		goto st0
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		if data[p] == 116 {
			goto st281
		}
		goto st0
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		switch data[p] {
		case 32:
			goto tr307
		case 46:
			goto tr390
		case 111:
			goto st282
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr307
		}
		goto st0
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		if data[p] == 98 {
			goto st283
		}
		goto st0
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		if data[p] == 101 {
			goto st284
		}
		goto st0
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		if data[p] == 114 {
			goto st285
		}
		goto st0
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		switch data[p] {
		case 32:
			goto tr307
		case 46:
			goto tr390
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr307
		}
		goto st0
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		if data[p] == 101 {
			goto st287
		}
		goto st0
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		if data[p] == 112 {
			goto st288
		}
		goto st0
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		switch data[p] {
		case 32:
			goto tr316
		case 46:
			goto tr397
		case 116:
			goto st289
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr316
		}
		goto st0
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		switch data[p] {
		case 32:
			goto tr316
		case 46:
			goto tr397
		case 101:
			goto st290
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr316
		}
		goto st0
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
		if data[p] == 109 {
			goto st291
		}
		goto st0
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		if data[p] == 98 {
			goto st292
		}
		goto st0
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		if data[p] == 101 {
			goto st293
		}
		goto st0
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
		if data[p] == 114 {
			goto st294
		}
		goto st0
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		switch data[p] {
		case 32:
			goto tr316
		case 46:
			goto tr397
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr316
		}
		goto st0
tr2:
//line ragel/datetime.rl:5
 pb = p 
	goto st295
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
//line ragel/parse_datetime.go:10136
		if data[p] == 32 {
			goto tr207
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st146
			}
		case data[p] >= 45:
			goto tr208
		}
		goto st0
tr3:
//line ragel/datetime.rl:5
 pb = p 
	goto st296
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
//line ragel/parse_datetime.go:10158
		if data[p] == 32 {
			goto tr207
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr208
			}
		case data[p] > 49:
			if 50 <= data[p] && data[p] <= 57 {
				goto st3
			}
		default:
			goto st146
		}
		goto st0
tr4:
//line ragel/datetime.rl:5
 pb = p 
	goto st297
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
//line ragel/parse_datetime.go:10184
		if data[p] == 32 {
			goto tr207
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] >= 45:
			goto tr208
		}
		goto st0
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		switch data[p] {
		case 112:
			goto st299
		case 117:
			goto st401
		}
		goto st0
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		if data[p] == 114 {
			goto st300
		}
		goto st0
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		switch data[p] {
		case 32:
			goto tr407
		case 46:
			goto tr409
		case 105:
			goto st399
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr408
		}
		goto st0
tr407:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st301
tr555:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st301
tr563:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st301
tr574:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st301
tr1136:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st301
tr1144:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st301
tr1147:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st301
tr1154:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st301
tr1158:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st301
tr1162:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st301
tr1171:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st301
tr1183:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st301
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
//line ragel/parse_datetime.go:10288
		switch data[p] {
		case 48:
			goto tr411
		case 51:
			goto tr413
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr414
			}
		case data[p] >= 49:
			goto tr412
		}
		goto st0
tr411:
//line ragel/datetime.rl:5
 pb = p 
	goto st302
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
//line ragel/parse_datetime.go:10313
		if 49 <= data[p] && data[p] <= 57 {
			goto st303
		}
		goto st0
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		switch data[p] {
		case 32:
			goto tr416
		case 44:
			goto tr417
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr418
		}
		goto st0
tr416:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st304
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
//line ragel/parse_datetime.go:10349
		if data[p] == 50 {
			goto tr420
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr421
			}
		case data[p] >= 48:
			goto tr419
		}
		goto st0
tr419:
//line ragel/datetime.rl:5
 pb = p 
	goto st305
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
//line ragel/parse_datetime.go:10371
		switch data[p] {
		case 32:
			goto tr422
		case 58:
			goto tr424
		case 65:
			goto tr425
		case 80:
			goto tr425
		case 97:
			goto tr426
		case 112:
			goto tr426
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st329
		}
		goto st0
tr422:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st306
tr463:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st306
tr520:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st306
tr503:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st306
tr508:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st306
tr514:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st306
tr531:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st306
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
//line ragel/parse_datetime.go:10462
		switch data[p] {
		case 65:
			goto tr428
		case 80:
			goto tr428
		case 97:
			goto tr429
		case 112:
			goto tr429
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr427
		}
		goto st0
tr427:
//line ragel/datetime.rl:5
 pb = p 
	goto st307
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
//line ragel/parse_datetime.go:10486
		if 48 <= data[p] && data[p] <= 57 {
			goto st308
		}
		goto st0
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		if 48 <= data[p] && data[p] <= 57 {
			goto st309
		}
		goto st0
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		if 48 <= data[p] && data[p] <= 57 {
			goto st938
		}
		goto st0
	st938:
		if p++; p == pe {
			goto _test_eof938
		}
	st_case_938:
		if data[p] == 32 {
			goto tr1372
		}
		goto st0
tr1372:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st310
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
//line ragel/parse_datetime.go:10529
		switch data[p] {
		case 43:
			goto st311
		case 45:
			goto st321
		case 47:
			goto tr435
		case 90:
			goto tr436
		case 95:
			goto tr435
		case 109:
			goto tr437
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr435
			}
		case data[p] >= 65:
			goto tr435
		}
		goto st0
tr1406:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st311
tr1417:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st311
tr1421:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st311
tr1436:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st311
tr1429:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st311
tr1449:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st311
tr1458:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st311
tr1466:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st311
tr1486:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st311
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
//line ragel/parse_datetime.go:10667
		if data[p] == 50 {
			goto tr439
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr440
			}
		case data[p] >= 48:
			goto tr438
		}
		goto st0
tr438:
//line ragel/datetime.rl:5
 pb = p 
	goto st939
tr456:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st939
	st939:
		if p++; p == pe {
			goto _test_eof939
		}
	st_case_939:
//line ragel/parse_datetime.go:10695
		switch data[p] {
		case 32:
			goto tr1373
		case 58:
			goto tr1375
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st949
		}
		goto st0
tr1373:
//line ragel/datetime.rl:197

    st.ZoneOffsetIsValid = true

    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st312
tr1388:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st312
tr1391:
//line ragel/datetime.rl:187

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st312
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
//line ragel/parse_datetime.go:10763
		switch data[p] {
		case 40:
			goto st313
		case 43:
			goto st315
		case 45:
			goto st317
		case 47:
			goto tr444
		case 95:
			goto tr444
		case 109:
			goto tr445
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr444
			}
		case data[p] >= 65:
			goto tr444
		}
		goto st0
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 32:
			goto st314
		case 43:
			goto st314
		case 45:
			goto st314
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st314
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st314
			}
		default:
			goto st314
		}
		goto st0
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		switch data[p] {
		case 32:
			goto st314
		case 41:
			goto st940
		case 43:
			goto st314
		case 45:
			goto st314
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st314
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st314
			}
		default:
			goto st314
		}
		goto st0
	st940:
		if p++; p == pe {
			goto _test_eof940
		}
	st_case_940:
		if data[p] == 32 {
			goto tr1376
		}
		goto st0
tr1393:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st315
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
//line ragel/parse_datetime.go:10862
		if data[p] == 50 {
			goto tr449
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr450
			}
		case data[p] >= 48:
			goto tr448
		}
		goto st0
tr448:
//line ragel/datetime.rl:5
 pb = p 
	goto st941
tr451:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st941
	st941:
		if p++; p == pe {
			goto _test_eof941
		}
	st_case_941:
//line ragel/parse_datetime.go:10890
		switch data[p] {
		case 32:
			goto tr1377
		case 58:
			goto tr1379
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st942
		}
		goto st0
tr1377:
//line ragel/datetime.rl:197

    st.ZoneOffsetIsValid = true

    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st316
tr1381:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st316
tr1384:
//line ragel/datetime.rl:187

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st316
tr1386:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st316
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
//line ragel/parse_datetime.go:10967
		switch data[p] {
		case 40:
			goto st313
		case 109:
			goto st14
		}
		goto st0
tr450:
//line ragel/datetime.rl:5
 pb = p 
	goto st942
tr453:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st942
	st942:
		if p++; p == pe {
			goto _test_eof942
		}
	st_case_942:
//line ragel/parse_datetime.go:10990
		switch data[p] {
		case 32:
			goto tr1377
		case 58:
			goto tr1379
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st943
		}
		goto st0
	st943:
		if p++; p == pe {
			goto _test_eof943
		}
	st_case_943:
		if data[p] == 32 {
			goto tr1377
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st943
		}
		goto st0
tr1379:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st944
	st944:
		if p++; p == pe {
			goto _test_eof944
		}
	st_case_944:
//line ragel/parse_datetime.go:11030
		if data[p] == 32 {
			goto tr1381
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1383
			}
		case data[p] >= 48:
			goto tr1382
		}
		goto st0
tr1382:
//line ragel/datetime.rl:5
 pb = p 
	goto st945
	st945:
		if p++; p == pe {
			goto _test_eof945
		}
	st_case_945:
//line ragel/parse_datetime.go:11052
		if data[p] == 32 {
			goto tr1384
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st946
		}
		goto st0
tr1383:
//line ragel/datetime.rl:5
 pb = p 
	goto st946
	st946:
		if p++; p == pe {
			goto _test_eof946
		}
	st_case_946:
//line ragel/parse_datetime.go:11069
		if data[p] == 32 {
			goto tr1384
		}
		goto st0
tr449:
//line ragel/datetime.rl:5
 pb = p 
	goto st947
tr452:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st947
	st947:
		if p++; p == pe {
			goto _test_eof947
		}
	st_case_947:
//line ragel/parse_datetime.go:11089
		switch data[p] {
		case 32:
			goto tr1377
		case 58:
			goto tr1379
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st943
			}
		case data[p] >= 48:
			goto st942
		}
		goto st0
tr1394:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st317
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
//line ragel/parse_datetime.go:11117
		if data[p] == 50 {
			goto tr452
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr453
			}
		case data[p] >= 48:
			goto tr451
		}
		goto st0
tr444:
//line ragel/datetime.rl:5
 pb = p 
	goto st318
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
//line ragel/parse_datetime.go:11139
		switch data[p] {
		case 47:
			goto st319
		case 95:
			goto st319
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st319
			}
		case data[p] >= 65:
			goto st319
		}
		goto st0
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		switch data[p] {
		case 47:
			goto st948
		case 95:
			goto st948
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st948
			}
		case data[p] >= 65:
			goto st948
		}
		goto st0
	st948:
		if p++; p == pe {
			goto _test_eof948
		}
	st_case_948:
		switch data[p] {
		case 32:
			goto tr1386
		case 47:
			goto st948
		case 95:
			goto st948
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st948
			}
		case data[p] >= 65:
			goto st948
		}
		goto st0
tr445:
//line ragel/datetime.rl:5
 pb = p 
	goto st320
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
//line ragel/parse_datetime.go:11206
		switch data[p] {
		case 47:
			goto st319
		case 61:
			goto st15
		case 95:
			goto st319
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st319
			}
		case data[p] >= 65:
			goto st319
		}
		goto st0
tr440:
//line ragel/datetime.rl:5
 pb = p 
	goto st949
tr458:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st949
	st949:
		if p++; p == pe {
			goto _test_eof949
		}
	st_case_949:
//line ragel/parse_datetime.go:11239
		switch data[p] {
		case 32:
			goto tr1373
		case 58:
			goto tr1375
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st950
		}
		goto st0
	st950:
		if p++; p == pe {
			goto _test_eof950
		}
	st_case_950:
		if data[p] == 32 {
			goto tr1373
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st950
		}
		goto st0
tr1375:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st951
	st951:
		if p++; p == pe {
			goto _test_eof951
		}
	st_case_951:
//line ragel/parse_datetime.go:11279
		if data[p] == 32 {
			goto tr1388
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1390
			}
		case data[p] >= 48:
			goto tr1389
		}
		goto st0
tr1389:
//line ragel/datetime.rl:5
 pb = p 
	goto st952
	st952:
		if p++; p == pe {
			goto _test_eof952
		}
	st_case_952:
//line ragel/parse_datetime.go:11301
		if data[p] == 32 {
			goto tr1391
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st953
		}
		goto st0
tr1390:
//line ragel/datetime.rl:5
 pb = p 
	goto st953
	st953:
		if p++; p == pe {
			goto _test_eof953
		}
	st_case_953:
//line ragel/parse_datetime.go:11318
		if data[p] == 32 {
			goto tr1391
		}
		goto st0
tr439:
//line ragel/datetime.rl:5
 pb = p 
	goto st954
tr457:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st954
	st954:
		if p++; p == pe {
			goto _test_eof954
		}
	st_case_954:
//line ragel/parse_datetime.go:11338
		switch data[p] {
		case 32:
			goto tr1373
		case 58:
			goto tr1375
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st950
			}
		case data[p] >= 48:
			goto st949
		}
		goto st0
tr1407:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st321
tr1418:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st321
tr1422:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st321
tr1437:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st321
tr1431:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st321
tr1450:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st321
tr1459:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st321
tr1468:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st321
tr1488:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st321
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
//line ragel/parse_datetime.go:11468
		if data[p] == 50 {
			goto tr457
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr458
			}
		case data[p] >= 48:
			goto tr456
		}
		goto st0
tr435:
//line ragel/datetime.rl:5
 pb = p 
	goto st322
tr1408:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st322
tr1423:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st322
tr1451:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st322
tr1460:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st322
tr1469:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st322
tr1438:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st322
tr1489:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st322
tr1432:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st322
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
//line ragel/parse_datetime.go:11590
		switch data[p] {
		case 47:
			goto st323
		case 95:
			goto st323
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st323
			}
		case data[p] >= 65:
			goto st323
		}
		goto st0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 47:
			goto st955
		case 95:
			goto st955
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st955
			}
		case data[p] >= 65:
			goto st955
		}
		goto st0
tr1419:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st955
	st955:
		if p++; p == pe {
			goto _test_eof955
		}
	st_case_955:
//line ragel/parse_datetime.go:11658
		switch data[p] {
		case 32:
			goto tr1386
		case 43:
			goto tr1393
		case 45:
			goto tr1394
		case 47:
			goto st955
		case 95:
			goto st955
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st955
			}
		case data[p] >= 65:
			goto st955
		}
		goto st0
tr436:
//line ragel/datetime.rl:5
 pb = p 
	goto st956
tr1412:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st956
tr1426:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st956
tr1455:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st956
tr1463:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st956
tr1472:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st956
tr1440:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st956
tr1491:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st956
tr1434:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st956
	st956:
		if p++; p == pe {
			goto _test_eof956
		}
	st_case_956:
//line ragel/parse_datetime.go:11789
		switch data[p] {
		case 32:
			goto tr1381
		case 47:
			goto st323
		case 95:
			goto st323
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st323
			}
		case data[p] >= 65:
			goto st323
		}
		goto st0
tr437:
//line ragel/datetime.rl:5
 pb = p 
	goto st324
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
//line ragel/parse_datetime.go:11816
		switch data[p] {
		case 47:
			goto st323
		case 61:
			goto st15
		case 95:
			goto st323
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st323
			}
		case data[p] >= 65:
			goto st323
		}
		goto st0
tr428:
//line ragel/datetime.rl:5
 pb = p 
	goto st325
tr425:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st325
tr466:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st325
tr506:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st325
tr510:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st325
tr517:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st325
tr522:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st325
tr533:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st325
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
//line ragel/parse_datetime.go:11924
		if data[p] == 77 {
			goto st326
		}
		goto st0
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		if data[p] == 32 {
			goto tr462
		}
		goto st0
tr462:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st327
tr488:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st327
tr499:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st327
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
//line ragel/parse_datetime.go:12016
		if 48 <= data[p] && data[p] <= 57 {
			goto tr427
		}
		goto st0
tr429:
//line ragel/datetime.rl:5
 pb = p 
	goto st328
tr426:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st328
tr467:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st328
tr507:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st328
tr511:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st328
tr518:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st328
tr523:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st328
tr534:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st328
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
//line ragel/parse_datetime.go:12111
		if data[p] == 109 {
			goto st326
		}
		goto st0
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		switch data[p] {
		case 32:
			goto tr463
		case 58:
			goto tr465
		case 65:
			goto tr466
		case 80:
			goto tr466
		case 97:
			goto tr467
		case 112:
			goto tr467
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st330
		}
		goto st0
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		if 48 <= data[p] && data[p] <= 57 {
			goto st957
		}
		goto st0
	st957:
		if p++; p == pe {
			goto _test_eof957
		}
	st_case_957:
		switch data[p] {
		case 32:
			goto tr1395
		case 43:
			goto tr1357
		case 44:
			goto tr1396
		case 45:
			goto tr1359
		case 46:
			goto tr500
		case 47:
			goto tr1360
		case 84:
			goto tr1361
		case 90:
			goto tr1362
		case 95:
			goto tr1363
		case 116:
			goto tr1363
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st357
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1360
			}
		default:
			goto tr1360
		}
		goto st0
tr1395:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st958
	st958:
		if p++; p == pe {
			goto _test_eof958
		}
	st_case_958:
//line ragel/parse_datetime.go:12214
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st19
		case 45:
			goto st31
		case 47:
			goto tr1210
		case 50:
			goto tr1399
		case 65:
			goto tr1211
		case 66:
			goto tr1212
		case 90:
			goto tr1213
		case 95:
			goto tr1210
		case 97:
			goto tr1214
		case 109:
			goto tr1215
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr1398
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1210
				}
			case data[p] >= 67:
				goto tr1210
			}
		default:
			goto tr1400
		}
		goto st0
tr1398:
//line ragel/datetime.rl:5
 pb = p 
	goto st959
	st959:
		if p++; p == pe {
			goto _test_eof959
		}
	st_case_959:
//line ragel/parse_datetime.go:12266
		switch data[p] {
		case 32:
			goto tr1242
		case 43:
			goto tr1243
		case 45:
			goto tr1244
		case 47:
			goto tr1245
		case 58:
			goto tr1247
		case 65:
			goto tr1248
		case 80:
			goto tr1248
		case 90:
			goto tr1249
		case 95:
			goto tr1245
		case 97:
			goto tr1250
		case 112:
			goto tr1250
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st960
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1245
			}
		default:
			goto tr1245
		}
		goto st0
	st960:
		if p++; p == pe {
			goto _test_eof960
		}
	st_case_960:
		switch data[p] {
		case 32:
			goto tr1259
		case 43:
			goto tr1260
		case 45:
			goto tr1261
		case 47:
			goto tr1262
		case 58:
			goto tr1263
		case 65:
			goto tr1264
		case 80:
			goto tr1264
		case 90:
			goto tr1265
		case 95:
			goto tr1262
		case 97:
			goto tr1266
		case 112:
			goto tr1266
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st331
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1262
			}
		default:
			goto tr1262
		}
		goto st0
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		if 48 <= data[p] && data[p] <= 57 {
			goto st961
		}
		goto st0
	st961:
		if p++; p == pe {
			goto _test_eof961
		}
	st_case_961:
		switch data[p] {
		case 32:
			goto tr1402
		case 43:
			goto tr1268
		case 45:
			goto tr1270
		case 47:
			goto tr1271
		case 90:
			goto tr1273
		case 95:
			goto tr1271
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1269
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1271
				}
			case data[p] >= 65:
				goto tr1271
			}
		default:
			goto st42
		}
		goto st0
tr1399:
//line ragel/datetime.rl:5
 pb = p 
	goto st962
	st962:
		if p++; p == pe {
			goto _test_eof962
		}
	st_case_962:
//line ragel/parse_datetime.go:12401
		switch data[p] {
		case 32:
			goto tr1242
		case 43:
			goto tr1243
		case 45:
			goto tr1244
		case 47:
			goto tr1245
		case 58:
			goto tr1247
		case 65:
			goto tr1248
		case 80:
			goto tr1248
		case 90:
			goto tr1249
		case 95:
			goto tr1245
		case 97:
			goto tr1250
		case 112:
			goto tr1250
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st960
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1245
				}
			case data[p] >= 66:
				goto tr1245
			}
		default:
			goto st332
		}
		goto st0
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		if 48 <= data[p] && data[p] <= 57 {
			goto st331
		}
		goto st0
tr1400:
//line ragel/datetime.rl:5
 pb = p 
	goto st963
	st963:
		if p++; p == pe {
			goto _test_eof963
		}
	st_case_963:
//line ragel/parse_datetime.go:12462
		switch data[p] {
		case 32:
			goto tr1242
		case 43:
			goto tr1243
		case 45:
			goto tr1244
		case 47:
			goto tr1245
		case 58:
			goto tr1247
		case 65:
			goto tr1248
		case 80:
			goto tr1248
		case 90:
			goto tr1249
		case 95:
			goto tr1245
		case 97:
			goto tr1250
		case 112:
			goto tr1250
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st332
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1245
			}
		default:
			goto tr1245
		}
		goto st0
tr1396:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st964
	st964:
		if p++; p == pe {
			goto _test_eof964
		}
	st_case_964:
//line ragel/parse_datetime.go:12526
		switch data[p] {
		case 32:
			goto st333
		case 44:
			goto st335
		case 84:
			goto st336
		case 95:
			goto st336
		case 116:
			goto st336
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr497
		}
		goto st0
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 50:
			goto tr95
		case 65:
			goto st334
		case 97:
			goto st346
		case 109:
			goto st14
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr96
			}
		case data[p] >= 48:
			goto tr94
		}
		goto st0
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		if data[p] == 84 {
			goto st335
		}
		goto st0
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		if data[p] == 32 {
			goto st336
		}
		goto st0
tr1496:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st336
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
//line ragel/parse_datetime.go:12596
		if data[p] == 50 {
			goto tr476
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr477
			}
		case data[p] >= 48:
			goto tr475
		}
		goto st0
tr475:
//line ragel/datetime.rl:5
 pb = p 
	goto st965
	st965:
		if p++; p == pe {
			goto _test_eof965
		}
	st_case_965:
//line ragel/parse_datetime.go:12618
		switch data[p] {
		case 32:
			goto tr1405
		case 43:
			goto tr1406
		case 45:
			goto tr1407
		case 47:
			goto tr1408
		case 58:
			goto tr1410
		case 65:
			goto tr1411
		case 80:
			goto tr1411
		case 90:
			goto tr1412
		case 95:
			goto tr1408
		case 97:
			goto tr1413
		case 112:
			goto tr1413
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st969
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1408
			}
		default:
			goto tr1408
		}
		goto st0
tr1405:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st966
tr1420:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st966
tr1474:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st966
tr1448:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st966
tr1457:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st966
tr1465:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st966
tr1485:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st966
	st966:
		if p++; p == pe {
			goto _test_eof966
		}
	st_case_966:
//line ragel/parse_datetime.go:12728
		switch data[p] {
		case 32:
			goto st13
		case 43:
			goto st311
		case 45:
			goto st321
		case 47:
			goto tr435
		case 65:
			goto tr1414
		case 80:
			goto tr1414
		case 90:
			goto tr436
		case 95:
			goto tr435
		case 97:
			goto tr1415
		case 109:
			goto tr437
		case 112:
			goto tr1415
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr435
			}
		case data[p] >= 66:
			goto tr435
		}
		goto st0
tr1414:
//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1411:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1425:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1454:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1462:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1471:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1476:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1490:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
//line ragel/parse_datetime.go:12852
		switch data[p] {
		case 47:
			goto st323
		case 77:
			goto st967
		case 95:
			goto st323
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st323
			}
		case data[p] >= 65:
			goto st323
		}
		goto st0
	st967:
		if p++; p == pe {
			goto _test_eof967
		}
	st_case_967:
		switch data[p] {
		case 32:
			goto tr1416
		case 43:
			goto tr1417
		case 45:
			goto tr1418
		case 47:
			goto tr1419
		case 95:
			goto tr1419
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1419
			}
		case data[p] >= 65:
			goto tr1419
		}
		goto st0
tr1416:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st968
tr1435:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st968
tr1428:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st968
	st968:
		if p++; p == pe {
			goto _test_eof968
		}
	st_case_968:
//line ragel/parse_datetime.go:12974
		switch data[p] {
		case 32:
			goto st13
		case 43:
			goto st311
		case 45:
			goto st321
		case 47:
			goto tr435
		case 90:
			goto tr436
		case 95:
			goto tr435
		case 109:
			goto tr437
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr435
			}
		case data[p] >= 65:
			goto tr435
		}
		goto st0
tr1415:
//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1413:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1427:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1456:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1464:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1473:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1477:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1492:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
//line ragel/parse_datetime.go:13090
		switch data[p] {
		case 47:
			goto st323
		case 95:
			goto st323
		case 109:
			goto st967
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st323
			}
		case data[p] >= 65:
			goto st323
		}
		goto st0
	st969:
		if p++; p == pe {
			goto _test_eof969
		}
	st_case_969:
		switch data[p] {
		case 32:
			goto tr1420
		case 43:
			goto tr1421
		case 45:
			goto tr1422
		case 47:
			goto tr1423
		case 58:
			goto tr1424
		case 65:
			goto tr1425
		case 80:
			goto tr1425
		case 90:
			goto tr1426
		case 95:
			goto tr1423
		case 97:
			goto tr1427
		case 112:
			goto tr1427
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st339
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1423
			}
		default:
			goto tr1423
		}
		goto st0
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		if 48 <= data[p] && data[p] <= 57 {
			goto st970
		}
		goto st0
	st970:
		if p++; p == pe {
			goto _test_eof970
		}
	st_case_970:
		switch data[p] {
		case 32:
			goto tr1428
		case 43:
			goto tr1429
		case 45:
			goto tr1431
		case 47:
			goto tr1432
		case 90:
			goto tr1434
		case 95:
			goto tr1432
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1430
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1432
				}
			case data[p] >= 65:
				goto tr1432
			}
		default:
			goto st341
		}
		goto st0
tr1430:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st340
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
//line ragel/parse_datetime.go:13218
		if 48 <= data[p] && data[p] <= 57 {
			goto tr480
		}
		goto st0
tr480:
//line ragel/datetime.rl:5
 pb = p 
	goto st971
	st971:
		if p++; p == pe {
			goto _test_eof971
		}
	st_case_971:
//line ragel/parse_datetime.go:13232
		switch data[p] {
		case 32:
			goto tr1435
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st972
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st972:
		if p++; p == pe {
			goto _test_eof972
		}
	st_case_972:
		switch data[p] {
		case 32:
			goto tr1435
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st973
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st973:
		if p++; p == pe {
			goto _test_eof973
		}
	st_case_973:
		switch data[p] {
		case 32:
			goto tr1435
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st974
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st974:
		if p++; p == pe {
			goto _test_eof974
		}
	st_case_974:
		switch data[p] {
		case 32:
			goto tr1435
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st975
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st975:
		if p++; p == pe {
			goto _test_eof975
		}
	st_case_975:
		switch data[p] {
		case 32:
			goto tr1435
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st976
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st976:
		if p++; p == pe {
			goto _test_eof976
		}
	st_case_976:
		switch data[p] {
		case 32:
			goto tr1435
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st977
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st977:
		if p++; p == pe {
			goto _test_eof977
		}
	st_case_977:
		switch data[p] {
		case 32:
			goto tr1435
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st978
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st978:
		if p++; p == pe {
			goto _test_eof978
		}
	st_case_978:
		switch data[p] {
		case 32:
			goto tr1435
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st979
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st979:
		if p++; p == pe {
			goto _test_eof979
		}
	st_case_979:
		switch data[p] {
		case 32:
			goto tr1435
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		case data[p] >= 65:
			goto tr1438
		}
		goto st0
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		if 48 <= data[p] && data[p] <= 57 {
			goto st980
		}
		goto st0
	st980:
		if p++; p == pe {
			goto _test_eof980
		}
	st_case_980:
		switch data[p] {
		case 32:
			goto tr1428
		case 43:
			goto tr1429
		case 45:
			goto tr1431
		case 47:
			goto tr1432
		case 90:
			goto tr1434
		case 95:
			goto tr1432
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1430
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1432
			}
		default:
			goto tr1432
		}
		goto st0
tr1410:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st342
tr1424:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st342
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
//line ragel/parse_datetime.go:13570
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr483
			}
		case data[p] >= 48:
			goto tr482
		}
		goto st0
tr482:
//line ragel/datetime.rl:5
 pb = p 
	goto st981
	st981:
		if p++; p == pe {
			goto _test_eof981
		}
	st_case_981:
//line ragel/parse_datetime.go:13589
		switch data[p] {
		case 32:
			goto tr1448
		case 43:
			goto tr1449
		case 45:
			goto tr1450
		case 47:
			goto tr1451
		case 58:
			goto tr1453
		case 65:
			goto tr1454
		case 80:
			goto tr1454
		case 90:
			goto tr1455
		case 95:
			goto tr1451
		case 97:
			goto tr1456
		case 112:
			goto tr1456
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st982
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1451
			}
		default:
			goto tr1451
		}
		goto st0
	st982:
		if p++; p == pe {
			goto _test_eof982
		}
	st_case_982:
		switch data[p] {
		case 32:
			goto tr1457
		case 43:
			goto tr1458
		case 45:
			goto tr1459
		case 47:
			goto tr1460
		case 58:
			goto tr1461
		case 65:
			goto tr1462
		case 80:
			goto tr1462
		case 90:
			goto tr1463
		case 95:
			goto tr1460
		case 97:
			goto tr1464
		case 112:
			goto tr1464
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1460
			}
		case data[p] >= 66:
			goto tr1460
		}
		goto st0
tr1453:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st343
tr1461:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st343
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
//line ragel/parse_datetime.go:13682
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr485
			}
		case data[p] >= 48:
			goto tr484
		}
		goto st0
tr484:
//line ragel/datetime.rl:5
 pb = p 
	goto st983
	st983:
		if p++; p == pe {
			goto _test_eof983
		}
	st_case_983:
//line ragel/parse_datetime.go:13701
		switch data[p] {
		case 32:
			goto tr1465
		case 43:
			goto tr1466
		case 45:
			goto tr1468
		case 47:
			goto tr1469
		case 65:
			goto tr1471
		case 80:
			goto tr1471
		case 90:
			goto tr1472
		case 95:
			goto tr1469
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1467
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1469
				}
			case data[p] >= 66:
				goto tr1469
			}
		default:
			goto st993
		}
		goto st0
tr1467:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st344
tr1487:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st344
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
//line ragel/parse_datetime.go:13759
		if 48 <= data[p] && data[p] <= 57 {
			goto tr486
		}
		goto st0
tr486:
//line ragel/datetime.rl:5
 pb = p 
	goto st984
	st984:
		if p++; p == pe {
			goto _test_eof984
		}
	st_case_984:
//line ragel/parse_datetime.go:13773
		switch data[p] {
		case 32:
			goto tr1474
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 65:
			goto tr1476
		case 80:
			goto tr1476
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		case 97:
			goto tr1477
		case 112:
			goto tr1477
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st985
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st985:
		if p++; p == pe {
			goto _test_eof985
		}
	st_case_985:
		switch data[p] {
		case 32:
			goto tr1474
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 65:
			goto tr1476
		case 80:
			goto tr1476
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		case 97:
			goto tr1477
		case 112:
			goto tr1477
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st986
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st986:
		if p++; p == pe {
			goto _test_eof986
		}
	st_case_986:
		switch data[p] {
		case 32:
			goto tr1474
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 65:
			goto tr1476
		case 80:
			goto tr1476
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		case 97:
			goto tr1477
		case 112:
			goto tr1477
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st987
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st987:
		if p++; p == pe {
			goto _test_eof987
		}
	st_case_987:
		switch data[p] {
		case 32:
			goto tr1474
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 65:
			goto tr1476
		case 80:
			goto tr1476
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		case 97:
			goto tr1477
		case 112:
			goto tr1477
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st988
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st988:
		if p++; p == pe {
			goto _test_eof988
		}
	st_case_988:
		switch data[p] {
		case 32:
			goto tr1474
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 65:
			goto tr1476
		case 80:
			goto tr1476
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		case 97:
			goto tr1477
		case 112:
			goto tr1477
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st989
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st989:
		if p++; p == pe {
			goto _test_eof989
		}
	st_case_989:
		switch data[p] {
		case 32:
			goto tr1474
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 65:
			goto tr1476
		case 80:
			goto tr1476
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		case 97:
			goto tr1477
		case 112:
			goto tr1477
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st990
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st990:
		if p++; p == pe {
			goto _test_eof990
		}
	st_case_990:
		switch data[p] {
		case 32:
			goto tr1474
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 65:
			goto tr1476
		case 80:
			goto tr1476
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		case 97:
			goto tr1477
		case 112:
			goto tr1477
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st991
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st991:
		if p++; p == pe {
			goto _test_eof991
		}
	st_case_991:
		switch data[p] {
		case 32:
			goto tr1474
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 65:
			goto tr1476
		case 80:
			goto tr1476
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		case 97:
			goto tr1477
		case 112:
			goto tr1477
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st992
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		default:
			goto tr1438
		}
		goto st0
	st992:
		if p++; p == pe {
			goto _test_eof992
		}
	st_case_992:
		switch data[p] {
		case 32:
			goto tr1474
		case 43:
			goto tr1436
		case 45:
			goto tr1437
		case 47:
			goto tr1438
		case 65:
			goto tr1476
		case 80:
			goto tr1476
		case 90:
			goto tr1440
		case 95:
			goto tr1438
		case 97:
			goto tr1477
		case 112:
			goto tr1477
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1438
			}
		case data[p] >= 66:
			goto tr1438
		}
		goto st0
	st993:
		if p++; p == pe {
			goto _test_eof993
		}
	st_case_993:
		switch data[p] {
		case 32:
			goto tr1485
		case 43:
			goto tr1486
		case 45:
			goto tr1488
		case 47:
			goto tr1489
		case 65:
			goto tr1490
		case 80:
			goto tr1490
		case 90:
			goto tr1491
		case 95:
			goto tr1489
		case 97:
			goto tr1492
		case 112:
			goto tr1492
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1487
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1489
			}
		default:
			goto tr1489
		}
		goto st0
tr485:
//line ragel/datetime.rl:5
 pb = p 
	goto st994
	st994:
		if p++; p == pe {
			goto _test_eof994
		}
	st_case_994:
//line ragel/parse_datetime.go:14174
		switch data[p] {
		case 32:
			goto tr1465
		case 43:
			goto tr1466
		case 45:
			goto tr1468
		case 47:
			goto tr1469
		case 65:
			goto tr1471
		case 80:
			goto tr1471
		case 90:
			goto tr1472
		case 95:
			goto tr1469
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1467
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1469
			}
		default:
			goto tr1469
		}
		goto st0
tr483:
//line ragel/datetime.rl:5
 pb = p 
	goto st995
	st995:
		if p++; p == pe {
			goto _test_eof995
		}
	st_case_995:
//line ragel/parse_datetime.go:14219
		switch data[p] {
		case 32:
			goto tr1448
		case 43:
			goto tr1449
		case 45:
			goto tr1450
		case 47:
			goto tr1451
		case 58:
			goto tr1453
		case 65:
			goto tr1454
		case 80:
			goto tr1454
		case 90:
			goto tr1455
		case 95:
			goto tr1451
		case 97:
			goto tr1456
		case 112:
			goto tr1456
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1451
			}
		case data[p] >= 66:
			goto tr1451
		}
		goto st0
tr476:
//line ragel/datetime.rl:5
 pb = p 
	goto st996
	st996:
		if p++; p == pe {
			goto _test_eof996
		}
	st_case_996:
//line ragel/parse_datetime.go:14262
		switch data[p] {
		case 32:
			goto tr1405
		case 43:
			goto tr1406
		case 45:
			goto tr1407
		case 47:
			goto tr1408
		case 58:
			goto tr1410
		case 65:
			goto tr1411
		case 80:
			goto tr1411
		case 90:
			goto tr1412
		case 95:
			goto tr1408
		case 97:
			goto tr1413
		case 112:
			goto tr1413
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st969
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1408
				}
			case data[p] >= 66:
				goto tr1408
			}
		default:
			goto st345
		}
		goto st0
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		if 48 <= data[p] && data[p] <= 57 {
			goto st339
		}
		goto st0
tr477:
//line ragel/datetime.rl:5
 pb = p 
	goto st997
	st997:
		if p++; p == pe {
			goto _test_eof997
		}
	st_case_997:
//line ragel/parse_datetime.go:14323
		switch data[p] {
		case 32:
			goto tr1405
		case 43:
			goto tr1406
		case 45:
			goto tr1407
		case 47:
			goto tr1408
		case 58:
			goto tr1410
		case 65:
			goto tr1411
		case 80:
			goto tr1411
		case 90:
			goto tr1412
		case 95:
			goto tr1408
		case 97:
			goto tr1413
		case 112:
			goto tr1413
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st345
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1408
			}
		default:
			goto tr1408
		}
		goto st0
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		if data[p] == 116 {
			goto st335
		}
		goto st0
tr497:
//line ragel/datetime.rl:5
 pb = p 
	goto st347
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
//line ragel/parse_datetime.go:14379
		if data[p] == 32 {
			goto tr488
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st348
		}
		goto st0
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		if data[p] == 32 {
			goto tr488
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st349
		}
		goto st0
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		if data[p] == 32 {
			goto tr488
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st350
		}
		goto st0
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		if data[p] == 32 {
			goto tr488
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st351
		}
		goto st0
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		if data[p] == 32 {
			goto tr488
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st352
		}
		goto st0
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		if data[p] == 32 {
			goto tr488
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st353
		}
		goto st0
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		if data[p] == 32 {
			goto tr488
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st354
		}
		goto st0
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		if data[p] == 32 {
			goto tr488
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st355
		}
		goto st0
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		if data[p] == 32 {
			goto tr488
		}
		goto st0
tr500:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st356
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
//line ragel/parse_datetime.go:14502
		if 48 <= data[p] && data[p] <= 57 {
			goto tr497
		}
		goto st0
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		if 48 <= data[p] && data[p] <= 57 {
			goto st358
		}
		goto st0
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		switch data[p] {
		case 32:
			goto tr499
		case 44:
			goto tr500
		case 46:
			goto tr500
		}
		goto st0
tr424:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st359
tr465:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st359
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
//line ragel/parse_datetime.go:14547
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr502
			}
		case data[p] >= 48:
			goto tr501
		}
		goto st0
tr501:
//line ragel/datetime.rl:5
 pb = p 
	goto st360
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
//line ragel/parse_datetime.go:14566
		switch data[p] {
		case 32:
			goto tr503
		case 58:
			goto tr505
		case 65:
			goto tr506
		case 80:
			goto tr506
		case 97:
			goto tr507
		case 112:
			goto tr507
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st361
		}
		goto st0
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		switch data[p] {
		case 32:
			goto tr508
		case 58:
			goto tr509
		case 65:
			goto tr510
		case 80:
			goto tr510
		case 97:
			goto tr511
		case 112:
			goto tr511
		}
		goto st0
tr505:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st362
tr509:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st362
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
//line ragel/parse_datetime.go:14622
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr513
			}
		case data[p] >= 48:
			goto tr512
		}
		goto st0
tr512:
//line ragel/datetime.rl:5
 pb = p 
	goto st363
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
//line ragel/parse_datetime.go:14641
		switch data[p] {
		case 32:
			goto tr514
		case 44:
			goto tr515
		case 46:
			goto tr515
		case 65:
			goto tr517
		case 80:
			goto tr517
		case 97:
			goto tr518
		case 112:
			goto tr518
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st374
		}
		goto st0
tr515:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st364
tr532:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st364
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
//line ragel/parse_datetime.go:14679
		if 48 <= data[p] && data[p] <= 57 {
			goto tr519
		}
		goto st0
tr519:
//line ragel/datetime.rl:5
 pb = p 
	goto st365
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
//line ragel/parse_datetime.go:14693
		switch data[p] {
		case 32:
			goto tr520
		case 65:
			goto tr522
		case 80:
			goto tr522
		case 97:
			goto tr523
		case 112:
			goto tr523
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st366
		}
		goto st0
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		switch data[p] {
		case 32:
			goto tr520
		case 65:
			goto tr522
		case 80:
			goto tr522
		case 97:
			goto tr523
		case 112:
			goto tr523
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st367
		}
		goto st0
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		switch data[p] {
		case 32:
			goto tr520
		case 65:
			goto tr522
		case 80:
			goto tr522
		case 97:
			goto tr523
		case 112:
			goto tr523
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st368
		}
		goto st0
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		switch data[p] {
		case 32:
			goto tr520
		case 65:
			goto tr522
		case 80:
			goto tr522
		case 97:
			goto tr523
		case 112:
			goto tr523
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st369
		}
		goto st0
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
		switch data[p] {
		case 32:
			goto tr520
		case 65:
			goto tr522
		case 80:
			goto tr522
		case 97:
			goto tr523
		case 112:
			goto tr523
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st370
		}
		goto st0
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		switch data[p] {
		case 32:
			goto tr520
		case 65:
			goto tr522
		case 80:
			goto tr522
		case 97:
			goto tr523
		case 112:
			goto tr523
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st371
		}
		goto st0
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		switch data[p] {
		case 32:
			goto tr520
		case 65:
			goto tr522
		case 80:
			goto tr522
		case 97:
			goto tr523
		case 112:
			goto tr523
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st372
		}
		goto st0
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		switch data[p] {
		case 32:
			goto tr520
		case 65:
			goto tr522
		case 80:
			goto tr522
		case 97:
			goto tr523
		case 112:
			goto tr523
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st373
		}
		goto st0
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		switch data[p] {
		case 32:
			goto tr520
		case 65:
			goto tr522
		case 80:
			goto tr522
		case 97:
			goto tr523
		case 112:
			goto tr523
		}
		goto st0
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		switch data[p] {
		case 32:
			goto tr531
		case 44:
			goto tr532
		case 46:
			goto tr532
		case 65:
			goto tr533
		case 80:
			goto tr533
		case 97:
			goto tr534
		case 112:
			goto tr534
		}
		goto st0
tr513:
//line ragel/datetime.rl:5
 pb = p 
	goto st375
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
//line ragel/parse_datetime.go:14906
		switch data[p] {
		case 32:
			goto tr514
		case 44:
			goto tr515
		case 46:
			goto tr515
		case 65:
			goto tr517
		case 80:
			goto tr517
		case 97:
			goto tr518
		case 112:
			goto tr518
		}
		goto st0
tr502:
//line ragel/datetime.rl:5
 pb = p 
	goto st376
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
//line ragel/parse_datetime.go:14933
		switch data[p] {
		case 32:
			goto tr503
		case 58:
			goto tr505
		case 65:
			goto tr506
		case 80:
			goto tr506
		case 97:
			goto tr507
		case 112:
			goto tr507
		}
		goto st0
tr420:
//line ragel/datetime.rl:5
 pb = p 
	goto st377
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
//line ragel/parse_datetime.go:14958
		switch data[p] {
		case 32:
			goto tr422
		case 58:
			goto tr424
		case 65:
			goto tr425
		case 80:
			goto tr425
		case 97:
			goto tr426
		case 112:
			goto tr426
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st378
			}
		case data[p] >= 48:
			goto st329
		}
		goto st0
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		if 48 <= data[p] && data[p] <= 57 {
			goto st330
		}
		goto st0
tr421:
//line ragel/datetime.rl:5
 pb = p 
	goto st379
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
//line ragel/parse_datetime.go:15000
		switch data[p] {
		case 32:
			goto tr422
		case 58:
			goto tr424
		case 65:
			goto tr425
		case 80:
			goto tr425
		case 97:
			goto tr426
		case 112:
			goto tr426
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st378
		}
		goto st0
tr417:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st380
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
//line ragel/parse_datetime.go:15035
		if data[p] == 32 {
			goto st381
		}
		goto st0
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr537
		}
		goto st0
tr537:
//line ragel/datetime.rl:5
 pb = p 
	goto st382
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
//line ragel/parse_datetime.go:15058
		if 48 <= data[p] && data[p] <= 57 {
			goto st383
		}
		goto st0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		if 48 <= data[p] && data[p] <= 57 {
			goto st384
		}
		goto st0
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		if 48 <= data[p] && data[p] <= 57 {
			goto st998
		}
		goto st0
	st998:
		if p++; p == pe {
			goto _test_eof998
		}
	st_case_998:
		switch data[p] {
		case 32:
			goto tr1494
		case 44:
			goto tr1495
		case 84:
			goto tr1496
		case 95:
			goto tr1496
		case 116:
			goto tr1496
		}
		goto st0
tr1494:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st385
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
//line ragel/parse_datetime.go:15110
		switch data[p] {
		case 50:
			goto tr476
		case 65:
			goto st334
		case 97:
			goto st346
		case 109:
			goto st14
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr477
			}
		case data[p] >= 48:
			goto tr475
		}
		goto st0
tr1495:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st999
	st999:
		if p++; p == pe {
			goto _test_eof999
		}
	st_case_999:
//line ragel/parse_datetime.go:15141
		switch data[p] {
		case 32:
			goto st385
		case 44:
			goto st335
		case 84:
			goto st336
		case 95:
			goto st336
		case 116:
			goto st336
		}
		goto st0
tr412:
//line ragel/datetime.rl:5
 pb = p 
	goto st386
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
//line ragel/parse_datetime.go:15164
		switch data[p] {
		case 32:
			goto tr541
		case 44:
			goto tr417
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st303
			}
		case data[p] >= 45:
			goto tr418
		}
		goto st0
tr541:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st387
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
//line ragel/parse_datetime.go:15196
		if 48 <= data[p] && data[p] <= 57 {
			goto tr542
		}
		goto st0
tr542:
//line ragel/datetime.rl:5
 pb = p 
	goto st388
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
//line ragel/parse_datetime.go:15210
		if 48 <= data[p] && data[p] <= 57 {
			goto st389
		}
		goto st0
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		if 48 <= data[p] && data[p] <= 57 {
			goto st390
		}
		goto st0
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1000
		}
		goto st0
	st1000:
		if p++; p == pe {
			goto _test_eof1000
		}
	st_case_1000:
		switch data[p] {
		case 32:
			goto tr1356
		case 43:
			goto tr1357
		case 44:
			goto tr1498
		case 45:
			goto tr1359
		case 47:
			goto tr1360
		case 84:
			goto tr1361
		case 90:
			goto tr1362
		case 95:
			goto tr1363
		case 116:
			goto tr1363
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1360
			}
		case data[p] >= 65:
			goto tr1360
		}
		goto st0
tr1498:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st1001
	st1001:
		if p++; p == pe {
			goto _test_eof1001
		}
	st_case_1001:
//line ragel/parse_datetime.go:15278
		switch data[p] {
		case 32:
			goto st333
		case 44:
			goto st335
		case 84:
			goto st336
		case 95:
			goto st336
		case 116:
			goto st336
		}
		goto st0
tr413:
//line ragel/datetime.rl:5
 pb = p 
	goto st391
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
//line ragel/parse_datetime.go:15301
		switch data[p] {
		case 32:
			goto tr541
		case 44:
			goto tr417
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st303
			}
		case data[p] >= 45:
			goto tr418
		}
		goto st0
tr414:
//line ragel/datetime.rl:5
 pb = p 
	goto st392
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
//line ragel/parse_datetime.go:15326
		switch data[p] {
		case 32:
			goto tr541
		case 44:
			goto tr417
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr418
		}
		goto st0
tr408:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st393
tr556:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st393
tr564:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st393
tr575:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st393
tr907:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st393
tr916:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st393
tr920:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st393
tr927:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st393
tr932:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st393
tr937:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st393
tr947:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st393
tr956:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st393
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
//line ragel/parse_datetime.go:15390
		switch data[p] {
		case 48:
			goto tr546
		case 51:
			goto tr548
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr549
			}
		case data[p] >= 49:
			goto tr547
		}
		goto st0
tr546:
//line ragel/datetime.rl:5
 pb = p 
	goto st394
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
//line ragel/parse_datetime.go:15415
		if 49 <= data[p] && data[p] <= 57 {
			goto st395
		}
		goto st0
tr549:
//line ragel/datetime.rl:5
 pb = p 
	goto st395
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
//line ragel/parse_datetime.go:15429
		if data[p] == 32 {
			goto tr418
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr418
		}
		goto st0
tr547:
//line ragel/datetime.rl:5
 pb = p 
	goto st396
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
//line ragel/parse_datetime.go:15446
		if data[p] == 32 {
			goto tr418
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st395
			}
		case data[p] >= 45:
			goto tr418
		}
		goto st0
tr548:
//line ragel/datetime.rl:5
 pb = p 
	goto st397
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
//line ragel/parse_datetime.go:15468
		if data[p] == 32 {
			goto tr418
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st395
			}
		case data[p] >= 45:
			goto tr418
		}
		goto st0
tr409:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st398
tr557:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st398
tr565:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st398
tr576:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st398
tr1137:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st398
tr1145:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st398
tr1148:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st398
tr1155:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st398
tr1159:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st398
tr1163:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st398
tr1172:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st398
tr1184:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st398
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
//line ragel/parse_datetime.go:15534
		switch data[p] {
		case 32:
			goto st301
		case 48:
			goto tr546
		case 51:
			goto tr548
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st393
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr549
			}
		default:
			goto tr547
		}
		goto st0
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		if data[p] == 108 {
			goto st400
		}
		goto st0
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		switch data[p] {
		case 32:
			goto tr407
		case 46:
			goto tr409
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr408
		}
		goto st0
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		if data[p] == 103 {
			goto st402
		}
		goto st0
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
		switch data[p] {
		case 32:
			goto tr555
		case 46:
			goto tr557
		case 117:
			goto st403
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr556
		}
		goto st0
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		if data[p] == 115 {
			goto st404
		}
		goto st0
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		if data[p] == 116 {
			goto st405
		}
		goto st0
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 32:
			goto tr555
		case 46:
			goto tr557
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr556
		}
		goto st0
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		if data[p] == 101 {
			goto st407
		}
		goto st0
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		if data[p] == 99 {
			goto st408
		}
		goto st0
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
		switch data[p] {
		case 32:
			goto tr563
		case 46:
			goto tr565
		case 101:
			goto st409
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr564
		}
		goto st0
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		if data[p] == 109 {
			goto st410
		}
		goto st0
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		if data[p] == 98 {
			goto st411
		}
		goto st0
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		if data[p] == 101 {
			goto st412
		}
		goto st0
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		if data[p] == 114 {
			goto st413
		}
		goto st0
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
		switch data[p] {
		case 32:
			goto tr563
		case 46:
			goto tr565
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr564
		}
		goto st0
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		switch data[p] {
		case 101:
			goto st415
		case 114:
			goto st422
		}
		goto st0
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		if data[p] == 98 {
			goto st416
		}
		goto st0
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		switch data[p] {
		case 32:
			goto tr574
		case 46:
			goto tr576
		case 114:
			goto st417
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr575
		}
		goto st0
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		if data[p] == 117 {
			goto st418
		}
		goto st0
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		if data[p] == 97 {
			goto st419
		}
		goto st0
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		if data[p] == 114 {
			goto st420
		}
		goto st0
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
		if data[p] == 121 {
			goto st421
		}
		goto st0
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		switch data[p] {
		case 32:
			goto tr574
		case 46:
			goto tr576
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr575
		}
		goto st0
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		if data[p] == 105 {
			goto st423
		}
		goto st0
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		switch data[p] {
		case 32:
			goto st424
		case 44:
			goto st659
		case 100:
			goto st804
		}
		goto st0
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 48:
			goto tr586
		case 51:
			goto tr588
		case 65:
			goto st429
		case 68:
			goto st601
		case 70:
			goto st609
		case 74:
			goto st617
		case 77:
			goto st629
		case 78:
			goto st635
		case 79:
			goto st643
		case 83:
			goto st650
		case 97:
			goto st429
		case 100:
			goto st601
		case 102:
			goto st609
		case 106:
			goto st617
		case 109:
			goto st629
		case 110:
			goto st635
		case 111:
			goto st643
		case 115:
			goto st650
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr589
			}
		case data[p] >= 49:
			goto tr587
		}
		goto st0
tr586:
//line ragel/datetime.rl:5
 pb = p 
	goto st425
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
//line ragel/parse_datetime.go:15898
		if 49 <= data[p] && data[p] <= 57 {
			goto st426
		}
		goto st0
tr589:
//line ragel/datetime.rl:5
 pb = p 
	goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
//line ragel/parse_datetime.go:15912
		if data[p] == 32 {
			goto tr208
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr208
		}
		goto st0
tr587:
//line ragel/datetime.rl:5
 pb = p 
	goto st427
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
//line ragel/parse_datetime.go:15929
		if data[p] == 32 {
			goto tr208
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st426
			}
		case data[p] >= 45:
			goto tr208
		}
		goto st0
tr588:
//line ragel/datetime.rl:5
 pb = p 
	goto st428
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
//line ragel/parse_datetime.go:15951
		if data[p] == 32 {
			goto tr208
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st426
			}
		case data[p] >= 45:
			goto tr208
		}
		goto st0
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
		switch data[p] {
		case 112:
			goto st430
		case 117:
			goto st596
		}
		goto st0
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
		if data[p] == 114 {
			goto st431
		}
		goto st0
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		switch data[p] {
		case 32:
			goto tr602
		case 46:
			goto tr603
		case 105:
			goto st594
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr408
		}
		goto st0
tr602:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st432
tr880:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st432
tr887:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st432
tr896:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st432
tr906:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st432
tr915:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st432
tr919:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st432
tr926:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st432
tr931:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st432
tr936:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st432
tr946:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st432
tr955:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st432
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
//line ragel/parse_datetime.go:16055
		switch data[p] {
		case 32:
			goto st433
		case 48:
			goto tr606
		case 51:
			goto tr608
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr609
			}
		case data[p] >= 49:
			goto tr607
		}
		goto st0
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 48:
			goto tr610
		case 51:
			goto tr612
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr613
			}
		case data[p] >= 49:
			goto tr611
		}
		goto st0
tr610:
//line ragel/datetime.rl:5
 pb = p 
	goto st434
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
//line ragel/parse_datetime.go:16102
		if 49 <= data[p] && data[p] <= 57 {
			goto st435
		}
		goto st0
tr613:
//line ragel/datetime.rl:5
 pb = p 
	goto st435
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
//line ragel/parse_datetime.go:16116
		if data[p] == 32 {
			goto tr615
		}
		goto st0
tr615:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st436
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
//line ragel/parse_datetime.go:16137
		if data[p] == 50 {
			goto tr617
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr618
			}
		case data[p] >= 48:
			goto tr616
		}
		goto st0
tr616:
//line ragel/datetime.rl:5
 pb = p 
	goto st437
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
//line ragel/parse_datetime.go:16159
		switch data[p] {
		case 32:
			goto tr619
		case 43:
			goto tr620
		case 45:
			goto tr621
		case 47:
			goto tr622
		case 58:
			goto tr624
		case 65:
			goto tr625
		case 80:
			goto tr625
		case 90:
			goto tr626
		case 95:
			goto tr622
		case 97:
			goto tr627
		case 112:
			goto tr627
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st477
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr622
			}
		default:
			goto tr622
		}
		goto st0
tr619:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st438
tr688:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st438
tr751:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st438
tr722:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st438
tr731:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st438
tr741:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st438
tr762:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st438
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
//line ragel/parse_datetime.go:16269
		switch data[p] {
		case 32:
			goto st439
		case 43:
			goto st443
		case 45:
			goto st468
		case 47:
			goto tr631
		case 65:
			goto tr633
		case 80:
			goto tr633
		case 90:
			goto tr634
		case 95:
			goto tr631
		case 97:
			goto tr635
		case 112:
			goto tr635
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr632
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr631
			}
		default:
			goto tr631
		}
		goto st0
tr651:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st439
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
//line ragel/parse_datetime.go:16314
		if 48 <= data[p] && data[p] <= 57 {
			goto tr632
		}
		goto st0
tr632:
//line ragel/datetime.rl:5
 pb = p 
	goto st440
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
//line ragel/parse_datetime.go:16328
		if 48 <= data[p] && data[p] <= 57 {
			goto st441
		}
		goto st0
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		if 48 <= data[p] && data[p] <= 57 {
			goto st442
		}
		goto st0
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1002
		}
		goto st0
	st1002:
		if p++; p == pe {
			goto _test_eof1002
		}
	st_case_1002:
		if data[p] == 32 {
			goto tr1499
		}
		goto st0
tr620:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st443
tr685:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st443
tr689:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st443
tr707:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st443
tr699:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st443
tr723:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st443
tr732:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st443
tr742:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st443
tr763:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
//line ragel/parse_datetime.go:16474
		if data[p] == 50 {
			goto tr640
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr641
			}
		case data[p] >= 48:
			goto tr639
		}
		goto st0
tr639:
//line ragel/datetime.rl:5
 pb = p 
	goto st444
tr676:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st444
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
//line ragel/parse_datetime.go:16502
		switch data[p] {
		case 32:
			goto tr642
		case 58:
			goto tr644
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st462
		}
		goto st0
tr642:
//line ragel/datetime.rl:197

    st.ZoneOffsetIsValid = true

    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st445
tr671:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st445
tr674:
//line ragel/datetime.rl:187

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st445
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
//line ragel/parse_datetime.go:16570
		switch data[p] {
		case 40:
			goto st446
		case 43:
			goto st449
		case 45:
			goto st458
		case 47:
			goto tr648
		case 95:
			goto tr648
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr632
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr648
			}
		default:
			goto tr648
		}
		goto st0
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		switch data[p] {
		case 32:
			goto st447
		case 43:
			goto st447
		case 45:
			goto st447
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st447
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st447
			}
		default:
			goto st447
		}
		goto st0
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
		switch data[p] {
		case 32:
			goto st447
		case 41:
			goto st448
		case 43:
			goto st447
		case 45:
			goto st447
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st447
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st447
			}
		default:
			goto st447
		}
		goto st0
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		if data[p] == 32 {
			goto tr651
		}
		goto st0
tr681:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st449
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
//line ragel/parse_datetime.go:16671
		if data[p] == 50 {
			goto tr653
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr654
			}
		case data[p] >= 48:
			goto tr652
		}
		goto st0
tr652:
//line ragel/datetime.rl:5
 pb = p 
	goto st450
tr664:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st450
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
//line ragel/parse_datetime.go:16699
		switch data[p] {
		case 32:
			goto tr655
		case 58:
			goto tr657
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st452
		}
		goto st0
tr655:
//line ragel/datetime.rl:197

    st.ZoneOffsetIsValid = true

    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st451
tr659:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st451
tr662:
//line ragel/datetime.rl:187

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st451
tr669:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st451
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
//line ragel/parse_datetime.go:16776
		if data[p] == 40 {
			goto st446
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr632
		}
		goto st0
tr654:
//line ragel/datetime.rl:5
 pb = p 
	goto st452
tr666:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st452
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
//line ragel/parse_datetime.go:16799
		switch data[p] {
		case 32:
			goto tr655
		case 58:
			goto tr657
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st453
		}
		goto st0
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		if data[p] == 32 {
			goto tr655
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st453
		}
		goto st0
tr657:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st454
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
//line ragel/parse_datetime.go:16839
		if data[p] == 32 {
			goto tr659
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr661
			}
		case data[p] >= 48:
			goto tr660
		}
		goto st0
tr660:
//line ragel/datetime.rl:5
 pb = p 
	goto st455
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
//line ragel/parse_datetime.go:16861
		if data[p] == 32 {
			goto tr662
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st456
		}
		goto st0
tr661:
//line ragel/datetime.rl:5
 pb = p 
	goto st456
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
//line ragel/parse_datetime.go:16878
		if data[p] == 32 {
			goto tr662
		}
		goto st0
tr653:
//line ragel/datetime.rl:5
 pb = p 
	goto st457
tr665:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st457
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
//line ragel/parse_datetime.go:16898
		switch data[p] {
		case 32:
			goto tr655
		case 58:
			goto tr657
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st453
			}
		case data[p] >= 48:
			goto st452
		}
		goto st0
tr682:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st458
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
//line ragel/parse_datetime.go:16926
		if data[p] == 50 {
			goto tr665
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr666
			}
		case data[p] >= 48:
			goto tr664
		}
		goto st0
tr648:
//line ragel/datetime.rl:5
 pb = p 
	goto st459
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
//line ragel/parse_datetime.go:16948
		switch data[p] {
		case 47:
			goto st460
		case 95:
			goto st460
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st460
			}
		case data[p] >= 65:
			goto st460
		}
		goto st0
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
		switch data[p] {
		case 47:
			goto st461
		case 95:
			goto st461
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st461
			}
		case data[p] >= 65:
			goto st461
		}
		goto st0
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		switch data[p] {
		case 32:
			goto tr669
		case 47:
			goto st461
		case 95:
			goto st461
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st461
			}
		case data[p] >= 65:
			goto st461
		}
		goto st0
tr641:
//line ragel/datetime.rl:5
 pb = p 
	goto st462
tr678:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st462
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
//line ragel/parse_datetime.go:17021
		switch data[p] {
		case 32:
			goto tr642
		case 58:
			goto tr644
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st463
		}
		goto st0
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
		if data[p] == 32 {
			goto tr642
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st463
		}
		goto st0
tr644:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st464
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
//line ragel/parse_datetime.go:17061
		if data[p] == 32 {
			goto tr671
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr673
			}
		case data[p] >= 48:
			goto tr672
		}
		goto st0
tr672:
//line ragel/datetime.rl:5
 pb = p 
	goto st465
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
//line ragel/parse_datetime.go:17083
		if data[p] == 32 {
			goto tr674
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st466
		}
		goto st0
tr673:
//line ragel/datetime.rl:5
 pb = p 
	goto st466
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
//line ragel/parse_datetime.go:17100
		if data[p] == 32 {
			goto tr674
		}
		goto st0
tr640:
//line ragel/datetime.rl:5
 pb = p 
	goto st467
tr677:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st467
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
//line ragel/parse_datetime.go:17120
		switch data[p] {
		case 32:
			goto tr642
		case 58:
			goto tr644
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st463
			}
		case data[p] >= 48:
			goto st462
		}
		goto st0
tr621:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st468
tr686:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st468
tr690:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st468
tr708:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st468
tr701:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st468
tr724:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st468
tr733:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st468
tr744:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st468
tr765:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
//line ragel/parse_datetime.go:17250
		if data[p] == 50 {
			goto tr677
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr678
			}
		case data[p] >= 48:
			goto tr676
		}
		goto st0
tr631:
//line ragel/datetime.rl:5
 pb = p 
	goto st469
tr622:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st469
tr691:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st469
tr725:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st469
tr734:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st469
tr745:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st469
tr709:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st469
tr766:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st469
tr702:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st469
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
//line ragel/parse_datetime.go:17372
		switch data[p] {
		case 47:
			goto st470
		case 95:
			goto st470
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st470
			}
		case data[p] >= 65:
			goto st470
		}
		goto st0
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		switch data[p] {
		case 47:
			goto st471
		case 95:
			goto st471
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st471
			}
		case data[p] >= 65:
			goto st471
		}
		goto st0
tr687:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st471
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
//line ragel/parse_datetime.go:17440
		switch data[p] {
		case 32:
			goto tr669
		case 43:
			goto tr681
		case 45:
			goto tr682
		case 47:
			goto st471
		case 95:
			goto st471
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st471
			}
		case data[p] >= 65:
			goto st471
		}
		goto st0
tr633:
//line ragel/datetime.rl:5
 pb = p 
	goto st472
tr625:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st472
tr694:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st472
tr728:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st472
tr736:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st472
tr747:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st472
tr753:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st472
tr767:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st472
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
//line ragel/parse_datetime.go:17552
		switch data[p] {
		case 47:
			goto st470
		case 77:
			goto st473
		case 95:
			goto st470
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st470
			}
		case data[p] >= 65:
			goto st470
		}
		goto st0
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		switch data[p] {
		case 32:
			goto tr684
		case 43:
			goto tr685
		case 45:
			goto tr686
		case 47:
			goto tr687
		case 95:
			goto tr687
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr687
			}
		case data[p] >= 65:
			goto tr687
		}
		goto st0
tr684:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st474
tr706:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st474
tr698:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st474
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
//line ragel/parse_datetime.go:17674
		switch data[p] {
		case 32:
			goto st439
		case 43:
			goto st443
		case 45:
			goto st468
		case 47:
			goto tr631
		case 90:
			goto tr634
		case 95:
			goto tr631
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr632
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr631
			}
		default:
			goto tr631
		}
		goto st0
tr634:
//line ragel/datetime.rl:5
 pb = p 
	goto st475
tr626:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st475
tr695:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st475
tr729:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st475
tr737:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st475
tr748:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st475
tr711:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st475
tr768:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st475
tr704:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st475
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
//line ragel/parse_datetime.go:17811
		switch data[p] {
		case 32:
			goto tr659
		case 47:
			goto st470
		case 95:
			goto st470
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st470
			}
		case data[p] >= 65:
			goto st470
		}
		goto st0
tr635:
//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr627:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr696:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr730:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr738:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr749:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr754:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr769:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
//line ragel/parse_datetime.go:17919
		switch data[p] {
		case 47:
			goto st470
		case 95:
			goto st470
		case 109:
			goto st473
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st470
			}
		case data[p] >= 65:
			goto st470
		}
		goto st0
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		switch data[p] {
		case 32:
			goto tr688
		case 43:
			goto tr689
		case 45:
			goto tr690
		case 47:
			goto tr691
		case 58:
			goto tr693
		case 65:
			goto tr694
		case 80:
			goto tr694
		case 90:
			goto tr695
		case 95:
			goto tr691
		case 97:
			goto tr696
		case 112:
			goto tr696
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st478
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr691
			}
		default:
			goto tr691
		}
		goto st0
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		if 48 <= data[p] && data[p] <= 57 {
			goto st479
		}
		goto st0
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		switch data[p] {
		case 32:
			goto tr698
		case 43:
			goto tr699
		case 45:
			goto tr701
		case 47:
			goto tr702
		case 90:
			goto tr704
		case 95:
			goto tr702
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr700
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr702
				}
			case data[p] >= 65:
				goto tr702
			}
		default:
			goto st490
		}
		goto st0
tr700:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st480
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
//line ragel/parse_datetime.go:18047
		if 48 <= data[p] && data[p] <= 57 {
			goto tr705
		}
		goto st0
tr705:
//line ragel/datetime.rl:5
 pb = p 
	goto st481
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
//line ragel/parse_datetime.go:18061
		switch data[p] {
		case 32:
			goto tr706
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st482
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
		switch data[p] {
		case 32:
			goto tr706
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st483
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
		switch data[p] {
		case 32:
			goto tr706
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st484
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
		switch data[p] {
		case 32:
			goto tr706
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st485
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
		switch data[p] {
		case 32:
			goto tr706
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st486
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		switch data[p] {
		case 32:
			goto tr706
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st487
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
		switch data[p] {
		case 32:
			goto tr706
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st488
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
		switch data[p] {
		case 32:
			goto tr706
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st489
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
		switch data[p] {
		case 32:
			goto tr706
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		case data[p] >= 65:
			goto tr709
		}
		goto st0
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		if 48 <= data[p] && data[p] <= 57 {
			goto st491
		}
		goto st0
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		switch data[p] {
		case 32:
			goto tr698
		case 43:
			goto tr699
		case 45:
			goto tr701
		case 47:
			goto tr702
		case 90:
			goto tr704
		case 95:
			goto tr702
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr700
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr702
			}
		default:
			goto tr702
		}
		goto st0
tr624:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st492
tr693:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st492
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
//line ragel/parse_datetime.go:18399
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr721
			}
		case data[p] >= 48:
			goto tr720
		}
		goto st0
tr720:
//line ragel/datetime.rl:5
 pb = p 
	goto st493
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
//line ragel/parse_datetime.go:18418
		switch data[p] {
		case 32:
			goto tr722
		case 43:
			goto tr723
		case 45:
			goto tr724
		case 47:
			goto tr725
		case 58:
			goto tr727
		case 65:
			goto tr728
		case 80:
			goto tr728
		case 90:
			goto tr729
		case 95:
			goto tr725
		case 97:
			goto tr730
		case 112:
			goto tr730
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st494
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr725
			}
		default:
			goto tr725
		}
		goto st0
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		switch data[p] {
		case 32:
			goto tr731
		case 43:
			goto tr732
		case 45:
			goto tr733
		case 47:
			goto tr734
		case 58:
			goto tr735
		case 65:
			goto tr736
		case 80:
			goto tr736
		case 90:
			goto tr737
		case 95:
			goto tr734
		case 97:
			goto tr738
		case 112:
			goto tr738
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr734
			}
		case data[p] >= 66:
			goto tr734
		}
		goto st0
tr727:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st495
tr735:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st495
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
//line ragel/parse_datetime.go:18511
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr740
			}
		case data[p] >= 48:
			goto tr739
		}
		goto st0
tr739:
//line ragel/datetime.rl:5
 pb = p 
	goto st496
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
//line ragel/parse_datetime.go:18530
		switch data[p] {
		case 32:
			goto tr741
		case 43:
			goto tr742
		case 45:
			goto tr744
		case 47:
			goto tr745
		case 65:
			goto tr747
		case 80:
			goto tr747
		case 90:
			goto tr748
		case 95:
			goto tr745
		case 97:
			goto tr749
		case 112:
			goto tr749
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr743
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr745
				}
			case data[p] >= 66:
				goto tr745
			}
		default:
			goto st507
		}
		goto st0
tr743:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st497
tr764:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st497
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
//line ragel/parse_datetime.go:18588
		if 48 <= data[p] && data[p] <= 57 {
			goto tr750
		}
		goto st0
tr750:
//line ragel/datetime.rl:5
 pb = p 
	goto st498
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
//line ragel/parse_datetime.go:18602
		switch data[p] {
		case 32:
			goto tr751
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr753
		case 80:
			goto tr753
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr754
		case 112:
			goto tr754
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st499
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		switch data[p] {
		case 32:
			goto tr751
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr753
		case 80:
			goto tr753
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr754
		case 112:
			goto tr754
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st500
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		switch data[p] {
		case 32:
			goto tr751
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr753
		case 80:
			goto tr753
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr754
		case 112:
			goto tr754
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st501
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		switch data[p] {
		case 32:
			goto tr751
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr753
		case 80:
			goto tr753
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr754
		case 112:
			goto tr754
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st502
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		switch data[p] {
		case 32:
			goto tr751
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr753
		case 80:
			goto tr753
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr754
		case 112:
			goto tr754
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st503
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		switch data[p] {
		case 32:
			goto tr751
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr753
		case 80:
			goto tr753
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr754
		case 112:
			goto tr754
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st504
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		switch data[p] {
		case 32:
			goto tr751
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr753
		case 80:
			goto tr753
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr754
		case 112:
			goto tr754
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st505
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		switch data[p] {
		case 32:
			goto tr751
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr753
		case 80:
			goto tr753
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr754
		case 112:
			goto tr754
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st506
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		switch data[p] {
		case 32:
			goto tr751
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr753
		case 80:
			goto tr753
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr754
		case 112:
			goto tr754
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		case data[p] >= 66:
			goto tr709
		}
		goto st0
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		switch data[p] {
		case 32:
			goto tr762
		case 43:
			goto tr763
		case 45:
			goto tr765
		case 47:
			goto tr766
		case 65:
			goto tr767
		case 80:
			goto tr767
		case 90:
			goto tr768
		case 95:
			goto tr766
		case 97:
			goto tr769
		case 112:
			goto tr769
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr764
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr766
			}
		default:
			goto tr766
		}
		goto st0
tr740:
//line ragel/datetime.rl:5
 pb = p 
	goto st508
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
//line ragel/parse_datetime.go:19003
		switch data[p] {
		case 32:
			goto tr741
		case 43:
			goto tr742
		case 45:
			goto tr744
		case 47:
			goto tr745
		case 65:
			goto tr747
		case 80:
			goto tr747
		case 90:
			goto tr748
		case 95:
			goto tr745
		case 97:
			goto tr749
		case 112:
			goto tr749
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr743
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr745
			}
		default:
			goto tr745
		}
		goto st0
tr721:
//line ragel/datetime.rl:5
 pb = p 
	goto st509
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
//line ragel/parse_datetime.go:19048
		switch data[p] {
		case 32:
			goto tr722
		case 43:
			goto tr723
		case 45:
			goto tr724
		case 47:
			goto tr725
		case 58:
			goto tr727
		case 65:
			goto tr728
		case 80:
			goto tr728
		case 90:
			goto tr729
		case 95:
			goto tr725
		case 97:
			goto tr730
		case 112:
			goto tr730
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr725
			}
		case data[p] >= 66:
			goto tr725
		}
		goto st0
tr617:
//line ragel/datetime.rl:5
 pb = p 
	goto st510
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
//line ragel/parse_datetime.go:19091
		switch data[p] {
		case 32:
			goto tr619
		case 43:
			goto tr620
		case 45:
			goto tr621
		case 47:
			goto tr622
		case 58:
			goto tr624
		case 65:
			goto tr625
		case 80:
			goto tr625
		case 90:
			goto tr626
		case 95:
			goto tr622
		case 97:
			goto tr627
		case 112:
			goto tr627
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st477
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr622
				}
			case data[p] >= 66:
				goto tr622
			}
		default:
			goto st511
		}
		goto st0
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		if 48 <= data[p] && data[p] <= 57 {
			goto st478
		}
		goto st0
tr618:
//line ragel/datetime.rl:5
 pb = p 
	goto st512
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
//line ragel/parse_datetime.go:19152
		switch data[p] {
		case 32:
			goto tr619
		case 43:
			goto tr620
		case 45:
			goto tr621
		case 47:
			goto tr622
		case 58:
			goto tr624
		case 65:
			goto tr625
		case 80:
			goto tr625
		case 90:
			goto tr626
		case 95:
			goto tr622
		case 97:
			goto tr627
		case 112:
			goto tr627
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st511
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr622
			}
		default:
			goto tr622
		}
		goto st0
tr611:
//line ragel/datetime.rl:5
 pb = p 
	goto st513
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
//line ragel/parse_datetime.go:19199
		if data[p] == 32 {
			goto tr615
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st435
		}
		goto st0
tr612:
//line ragel/datetime.rl:5
 pb = p 
	goto st514
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
//line ragel/parse_datetime.go:19216
		if data[p] == 32 {
			goto tr615
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st435
		}
		goto st0
tr606:
//line ragel/datetime.rl:5
 pb = p 
	goto st515
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
//line ragel/parse_datetime.go:19233
		if 49 <= data[p] && data[p] <= 57 {
			goto st516
		}
		goto st0
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		if data[p] == 32 {
			goto tr772
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr418
		}
		goto st0
tr772:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st517
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
//line ragel/parse_datetime.go:19266
		if data[p] == 50 {
			goto tr774
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr775
			}
		case data[p] >= 48:
			goto tr773
		}
		goto st0
tr773:
//line ragel/datetime.rl:5
 pb = p 
	goto st518
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
//line ragel/parse_datetime.go:19288
		switch data[p] {
		case 32:
			goto tr776
		case 43:
			goto tr620
		case 45:
			goto tr621
		case 47:
			goto tr622
		case 58:
			goto tr778
		case 65:
			goto tr779
		case 80:
			goto tr779
		case 90:
			goto tr626
		case 95:
			goto tr622
		case 97:
			goto tr780
		case 112:
			goto tr780
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st524
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr622
			}
		default:
			goto tr622
		}
		goto st0
tr776:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st519
tr785:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st519
tr853:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st519
tr836:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st519
tr841:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st519
tr847:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st519
tr864:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st519
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
//line ragel/parse_datetime.go:19398
		switch data[p] {
		case 32:
			goto st439
		case 43:
			goto st443
		case 45:
			goto st468
		case 47:
			goto tr631
		case 65:
			goto tr781
		case 80:
			goto tr781
		case 90:
			goto tr634
		case 95:
			goto tr631
		case 97:
			goto tr782
		case 112:
			goto tr782
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr427
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr631
			}
		default:
			goto tr631
		}
		goto st0
tr781:
//line ragel/datetime.rl:5
 pb = p 
	goto st520
tr779:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st520
tr788:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st520
tr839:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st520
tr843:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st520
tr850:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st520
tr855:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st520
tr866:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st520
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
//line ragel/parse_datetime.go:19524
		switch data[p] {
		case 47:
			goto st470
		case 77:
			goto st521
		case 95:
			goto st470
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st470
			}
		case data[p] >= 65:
			goto st470
		}
		goto st0
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		switch data[p] {
		case 32:
			goto tr784
		case 43:
			goto tr685
		case 45:
			goto tr686
		case 47:
			goto tr687
		case 95:
			goto tr687
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr687
			}
		case data[p] >= 65:
			goto tr687
		}
		goto st0
tr784:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st522
tr822:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

	goto st522
tr832:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st522
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
//line ragel/parse_datetime.go:19646
		switch data[p] {
		case 32:
			goto st439
		case 43:
			goto st443
		case 45:
			goto st468
		case 47:
			goto tr631
		case 90:
			goto tr634
		case 95:
			goto tr631
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr427
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr631
			}
		default:
			goto tr631
		}
		goto st0
tr782:
//line ragel/datetime.rl:5
 pb = p 
	goto st523
tr780:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st523
tr789:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st523
tr840:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st523
tr844:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st523
tr851:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st523
tr856:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st523
tr867:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st523
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
//line ragel/parse_datetime.go:19764
		switch data[p] {
		case 47:
			goto st470
		case 95:
			goto st470
		case 109:
			goto st521
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st470
			}
		case data[p] >= 65:
			goto st470
		}
		goto st0
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
		switch data[p] {
		case 32:
			goto tr785
		case 43:
			goto tr689
		case 45:
			goto tr690
		case 47:
			goto tr691
		case 58:
			goto tr787
		case 65:
			goto tr788
		case 80:
			goto tr788
		case 90:
			goto tr695
		case 95:
			goto tr691
		case 97:
			goto tr789
		case 112:
			goto tr789
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st525
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr691
			}
		default:
			goto tr691
		}
		goto st0
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1003
		}
		goto st0
	st1003:
		if p++; p == pe {
			goto _test_eof1003
		}
	st_case_1003:
		switch data[p] {
		case 32:
			goto tr1500
		case 43:
			goto tr1501
		case 44:
			goto tr1502
		case 45:
			goto tr1503
		case 46:
			goto tr833
		case 47:
			goto tr1504
		case 84:
			goto tr1506
		case 90:
			goto tr1507
		case 95:
			goto tr1508
		case 116:
			goto tr1508
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st558
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1504
			}
		default:
			goto tr1504
		}
		goto st0
tr1500:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st1004
	st1004:
		if p++; p == pe {
			goto _test_eof1004
		}
	st_case_1004:
//line ragel/parse_datetime.go:19899
		switch data[p] {
		case 32:
			goto st526
		case 43:
			goto st527
		case 45:
			goto st539
		case 47:
			goto tr1512
		case 50:
			goto tr1399
		case 65:
			goto tr1513
		case 66:
			goto tr1514
		case 90:
			goto tr1515
		case 95:
			goto tr1512
		case 97:
			goto tr1516
		case 109:
			goto tr1517
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr1398
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1512
				}
			case data[p] >= 67:
				goto tr1512
			}
		default:
			goto tr1400
		}
		goto st0
tr1521:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st526
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
//line ragel/parse_datetime.go:19951
		switch data[p] {
		case 65:
			goto st12
		case 66:
			goto st18
		case 109:
			goto st14
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr632
		}
		goto st0
tr1501:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st527
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
//line ragel/parse_datetime.go:19990
		if data[p] == 50 {
			goto tr792
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr793
			}
		case data[p] >= 48:
			goto tr791
		}
		goto st0
tr791:
//line ragel/datetime.rl:5
 pb = p 
	goto st1005
tr813:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1005
	st1005:
		if p++; p == pe {
			goto _test_eof1005
		}
	st_case_1005:
//line ragel/parse_datetime.go:20018
		switch data[p] {
		case 32:
			goto tr1518
		case 58:
			goto tr1520
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1017
		}
		goto st0
tr1518:
//line ragel/datetime.rl:197

    st.ZoneOffsetIsValid = true

    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st528
tr1533:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st528
tr1536:
//line ragel/datetime.rl:187

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st528
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
//line ragel/parse_datetime.go:20086
		switch data[p] {
		case 40:
			goto st529
		case 43:
			goto st531
		case 45:
			goto st533
		case 47:
			goto tr797
		case 65:
			goto tr798
		case 66:
			goto tr799
		case 95:
			goto tr797
		case 109:
			goto tr800
		}
		switch {
		case data[p] < 67:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr632
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr797
			}
		default:
			goto tr797
		}
		goto st0
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
		switch data[p] {
		case 32:
			goto st530
		case 43:
			goto st530
		case 45:
			goto st530
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st530
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st530
			}
		default:
			goto st530
		}
		goto st0
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
		switch data[p] {
		case 32:
			goto st530
		case 41:
			goto st1006
		case 43:
			goto st530
		case 45:
			goto st530
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st530
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st530
			}
		default:
			goto st530
		}
		goto st0
	st1006:
		if p++; p == pe {
			goto _test_eof1006
		}
	st_case_1006:
		if data[p] == 32 {
			goto tr1521
		}
		goto st0
tr1538:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st531
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
//line ragel/parse_datetime.go:20193
		if data[p] == 50 {
			goto tr804
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr805
			}
		case data[p] >= 48:
			goto tr803
		}
		goto st0
tr803:
//line ragel/datetime.rl:5
 pb = p 
	goto st1007
tr806:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1007
	st1007:
		if p++; p == pe {
			goto _test_eof1007
		}
	st_case_1007:
//line ragel/parse_datetime.go:20221
		switch data[p] {
		case 32:
			goto tr1522
		case 58:
			goto tr1524
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1008
		}
		goto st0
tr1522:
//line ragel/datetime.rl:197

    st.ZoneOffsetIsValid = true

    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st532
tr1526:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st532
tr1529:
//line ragel/datetime.rl:187

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st532
tr1531:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st532
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
//line ragel/parse_datetime.go:20298
		switch data[p] {
		case 40:
			goto st529
		case 65:
			goto st12
		case 66:
			goto st18
		case 109:
			goto st14
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr632
		}
		goto st0
tr805:
//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr808:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1008
	st1008:
		if p++; p == pe {
			goto _test_eof1008
		}
	st_case_1008:
//line ragel/parse_datetime.go:20328
		switch data[p] {
		case 32:
			goto tr1522
		case 58:
			goto tr1524
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1009
		}
		goto st0
	st1009:
		if p++; p == pe {
			goto _test_eof1009
		}
	st_case_1009:
		if data[p] == 32 {
			goto tr1522
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1009
		}
		goto st0
tr1524:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st1010
	st1010:
		if p++; p == pe {
			goto _test_eof1010
		}
	st_case_1010:
//line ragel/parse_datetime.go:20368
		if data[p] == 32 {
			goto tr1526
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1528
			}
		case data[p] >= 48:
			goto tr1527
		}
		goto st0
tr1527:
//line ragel/datetime.rl:5
 pb = p 
	goto st1011
	st1011:
		if p++; p == pe {
			goto _test_eof1011
		}
	st_case_1011:
//line ragel/parse_datetime.go:20390
		if data[p] == 32 {
			goto tr1529
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1012
		}
		goto st0
tr1528:
//line ragel/datetime.rl:5
 pb = p 
	goto st1012
	st1012:
		if p++; p == pe {
			goto _test_eof1012
		}
	st_case_1012:
//line ragel/parse_datetime.go:20407
		if data[p] == 32 {
			goto tr1529
		}
		goto st0
tr804:
//line ragel/datetime.rl:5
 pb = p 
	goto st1013
tr807:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1013
	st1013:
		if p++; p == pe {
			goto _test_eof1013
		}
	st_case_1013:
//line ragel/parse_datetime.go:20427
		switch data[p] {
		case 32:
			goto tr1522
		case 58:
			goto tr1524
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st1009
			}
		case data[p] >= 48:
			goto st1008
		}
		goto st0
tr1539:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st533
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
//line ragel/parse_datetime.go:20455
		if data[p] == 50 {
			goto tr807
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr808
			}
		case data[p] >= 48:
			goto tr806
		}
		goto st0
tr797:
//line ragel/datetime.rl:5
 pb = p 
	goto st534
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
//line ragel/parse_datetime.go:20477
		switch data[p] {
		case 47:
			goto st535
		case 95:
			goto st535
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st535
			}
		case data[p] >= 65:
			goto st535
		}
		goto st0
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		switch data[p] {
		case 47:
			goto st1014
		case 95:
			goto st1014
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1014
			}
		case data[p] >= 65:
			goto st1014
		}
		goto st0
	st1014:
		if p++; p == pe {
			goto _test_eof1014
		}
	st_case_1014:
		switch data[p] {
		case 32:
			goto tr1531
		case 47:
			goto st1014
		case 95:
			goto st1014
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1014
			}
		case data[p] >= 65:
			goto st1014
		}
		goto st0
tr798:
//line ragel/datetime.rl:5
 pb = p 
	goto st536
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
//line ragel/parse_datetime.go:20544
		switch data[p] {
		case 47:
			goto st535
		case 68:
			goto st1015
		case 95:
			goto st535
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st535
			}
		case data[p] >= 65:
			goto st535
		}
		goto st0
	st1015:
		if p++; p == pe {
			goto _test_eof1015
		}
	st_case_1015:
		switch data[p] {
		case 32:
			goto st13
		case 47:
			goto st1014
		case 95:
			goto st1014
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1014
			}
		case data[p] >= 65:
			goto st1014
		}
		goto st0
tr799:
//line ragel/datetime.rl:5
 pb = p 
	goto st537
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
//line ragel/parse_datetime.go:20593
		switch data[p] {
		case 47:
			goto st535
		case 67:
			goto st1016
		case 95:
			goto st535
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st535
			}
		case data[p] >= 65:
			goto st535
		}
		goto st0
	st1016:
		if p++; p == pe {
			goto _test_eof1016
		}
	st_case_1016:
		switch data[p] {
		case 32:
			goto tr1219
		case 47:
			goto st1014
		case 95:
			goto st1014
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1014
			}
		case data[p] >= 65:
			goto st1014
		}
		goto st0
tr800:
//line ragel/datetime.rl:5
 pb = p 
	goto st538
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
//line ragel/parse_datetime.go:20642
		switch data[p] {
		case 47:
			goto st535
		case 61:
			goto st15
		case 95:
			goto st535
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st535
			}
		case data[p] >= 65:
			goto st535
		}
		goto st0
tr793:
//line ragel/datetime.rl:5
 pb = p 
	goto st1017
tr815:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1017
	st1017:
		if p++; p == pe {
			goto _test_eof1017
		}
	st_case_1017:
//line ragel/parse_datetime.go:20675
		switch data[p] {
		case 32:
			goto tr1518
		case 58:
			goto tr1520
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1018
		}
		goto st0
	st1018:
		if p++; p == pe {
			goto _test_eof1018
		}
	st_case_1018:
		if data[p] == 32 {
			goto tr1518
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1018
		}
		goto st0
tr1520:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st1019
	st1019:
		if p++; p == pe {
			goto _test_eof1019
		}
	st_case_1019:
//line ragel/parse_datetime.go:20715
		if data[p] == 32 {
			goto tr1533
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1535
			}
		case data[p] >= 48:
			goto tr1534
		}
		goto st0
tr1534:
//line ragel/datetime.rl:5
 pb = p 
	goto st1020
	st1020:
		if p++; p == pe {
			goto _test_eof1020
		}
	st_case_1020:
//line ragel/parse_datetime.go:20737
		if data[p] == 32 {
			goto tr1536
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1021
		}
		goto st0
tr1535:
//line ragel/datetime.rl:5
 pb = p 
	goto st1021
	st1021:
		if p++; p == pe {
			goto _test_eof1021
		}
	st_case_1021:
//line ragel/parse_datetime.go:20754
		if data[p] == 32 {
			goto tr1536
		}
		goto st0
tr792:
//line ragel/datetime.rl:5
 pb = p 
	goto st1022
tr814:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1022
	st1022:
		if p++; p == pe {
			goto _test_eof1022
		}
	st_case_1022:
//line ragel/parse_datetime.go:20774
		switch data[p] {
		case 32:
			goto tr1518
		case 58:
			goto tr1520
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st1018
			}
		case data[p] >= 48:
			goto st1017
		}
		goto st0
tr1503:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st539
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
//line ragel/parse_datetime.go:20816
		if data[p] == 50 {
			goto tr814
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr815
			}
		case data[p] >= 48:
			goto tr813
		}
		goto st0
tr1512:
//line ragel/datetime.rl:5
 pb = p 
	goto st540
tr1504:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st540
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
//line ragel/parse_datetime.go:20861
		switch data[p] {
		case 47:
			goto st541
		case 95:
			goto st541
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st541
			}
		case data[p] >= 65:
			goto st541
		}
		goto st0
tr1540:
//line ragel/datetime.rl:5
 pb = p 
	goto st541
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
//line ragel/parse_datetime.go:20886
		switch data[p] {
		case 47:
			goto st1023
		case 95:
			goto st1023
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1023
			}
		case data[p] >= 65:
			goto st1023
		}
		goto st0
	st1023:
		if p++; p == pe {
			goto _test_eof1023
		}
	st_case_1023:
		switch data[p] {
		case 32:
			goto tr1531
		case 43:
			goto tr1538
		case 45:
			goto tr1539
		case 47:
			goto st1023
		case 95:
			goto st1023
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1023
			}
		case data[p] >= 65:
			goto st1023
		}
		goto st0
tr1513:
//line ragel/datetime.rl:5
 pb = p 
	goto st542
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
//line ragel/parse_datetime.go:20937
		switch data[p] {
		case 47:
			goto st541
		case 68:
			goto st1024
		case 84:
			goto st543
		case 95:
			goto st541
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st541
			}
		case data[p] >= 65:
			goto st541
		}
		goto st0
	st1024:
		if p++; p == pe {
			goto _test_eof1024
		}
	st_case_1024:
		switch data[p] {
		case 32:
			goto st13
		case 47:
			goto st1023
		case 95:
			goto st1023
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1023
			}
		case data[p] >= 65:
			goto st1023
		}
		goto st0
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
		switch data[p] {
		case 32:
			goto st49
		case 47:
			goto st1023
		case 95:
			goto st1023
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1023
			}
		case data[p] >= 65:
			goto st1023
		}
		goto st0
tr1514:
//line ragel/datetime.rl:5
 pb = p 
	goto st544
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
//line ragel/parse_datetime.go:21010
		switch data[p] {
		case 47:
			goto st541
		case 67:
			goto st1025
		case 95:
			goto st541
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st541
			}
		case data[p] >= 65:
			goto st541
		}
		goto st0
	st1025:
		if p++; p == pe {
			goto _test_eof1025
		}
	st_case_1025:
		switch data[p] {
		case 32:
			goto tr1219
		case 47:
			goto st1023
		case 95:
			goto st1023
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1023
			}
		case data[p] >= 65:
			goto st1023
		}
		goto st0
tr1515:
//line ragel/datetime.rl:5
 pb = p 
	goto st1026
tr1507:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st1026
	st1026:
		if p++; p == pe {
			goto _test_eof1026
		}
	st_case_1026:
//line ragel/parse_datetime.go:21082
		switch data[p] {
		case 32:
			goto tr1526
		case 47:
			goto st541
		case 95:
			goto st541
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st541
			}
		case data[p] >= 65:
			goto st541
		}
		goto st0
tr1516:
//line ragel/datetime.rl:5
 pb = p 
	goto st545
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
//line ragel/parse_datetime.go:21109
		switch data[p] {
		case 47:
			goto st541
		case 95:
			goto st541
		case 116:
			goto st543
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st541
			}
		case data[p] >= 65:
			goto st541
		}
		goto st0
tr1517:
//line ragel/datetime.rl:5
 pb = p 
	goto st546
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
//line ragel/parse_datetime.go:21136
		switch data[p] {
		case 47:
			goto st541
		case 61:
			goto st15
		case 95:
			goto st541
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st541
			}
		case data[p] >= 65:
			goto st541
		}
		goto st0
tr1502:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st547
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
//line ragel/parse_datetime.go:21180
		if data[p] == 32 {
			goto st49
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr821
		}
		goto st0
tr821:
//line ragel/datetime.rl:5
 pb = p 
	goto st548
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
//line ragel/parse_datetime.go:21197
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st549
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st550
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st551
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st552
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st553
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st554
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st555
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st556
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 90:
			goto tr711
		case 95:
			goto tr709
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		case data[p] >= 65:
			goto tr709
		}
		goto st0
tr833:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st557
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
//line ragel/parse_datetime.go:21499
		if 48 <= data[p] && data[p] <= 57 {
			goto tr821
		}
		goto st0
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
		if 48 <= data[p] && data[p] <= 57 {
			goto st559
		}
		goto st0
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
		switch data[p] {
		case 32:
			goto tr832
		case 43:
			goto tr699
		case 45:
			goto tr701
		case 47:
			goto tr702
		case 90:
			goto tr704
		case 95:
			goto tr702
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr833
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr702
			}
		default:
			goto tr702
		}
		goto st0
tr1506:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st1027
	st1027:
		if p++; p == pe {
			goto _test_eof1027
		}
	st_case_1027:
//line ragel/parse_datetime.go:21573
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st19
		case 45:
			goto st31
		case 47:
			goto tr1540
		case 50:
			goto tr95
		case 90:
			goto tr1541
		case 95:
			goto tr1540
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr94
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1540
				}
			case data[p] >= 65:
				goto tr1540
			}
		default:
			goto tr96
		}
		goto st0
tr1541:
//line ragel/datetime.rl:5
 pb = p 
	goto st1028
	st1028:
		if p++; p == pe {
			goto _test_eof1028
		}
	st_case_1028:
//line ragel/parse_datetime.go:21617
		switch data[p] {
		case 32:
			goto tr1228
		case 47:
			goto st1023
		case 95:
			goto st1023
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1023
			}
		case data[p] >= 65:
			goto st1023
		}
		goto st0
tr1508:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st560
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
//line ragel/parse_datetime.go:21663
		switch data[p] {
		case 47:
			goto st541
		case 50:
			goto tr95
		case 95:
			goto st541
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr94
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st541
				}
			case data[p] >= 65:
				goto st541
			}
		default:
			goto tr96
		}
		goto st0
tr778:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st561
tr787:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st561
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
//line ragel/parse_datetime.go:21707
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr835
			}
		case data[p] >= 48:
			goto tr834
		}
		goto st0
tr834:
//line ragel/datetime.rl:5
 pb = p 
	goto st562
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
//line ragel/parse_datetime.go:21726
		switch data[p] {
		case 32:
			goto tr836
		case 43:
			goto tr723
		case 45:
			goto tr724
		case 47:
			goto tr725
		case 58:
			goto tr838
		case 65:
			goto tr839
		case 80:
			goto tr839
		case 90:
			goto tr729
		case 95:
			goto tr725
		case 97:
			goto tr840
		case 112:
			goto tr840
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st563
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr725
			}
		default:
			goto tr725
		}
		goto st0
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
		switch data[p] {
		case 32:
			goto tr841
		case 43:
			goto tr732
		case 45:
			goto tr733
		case 47:
			goto tr734
		case 58:
			goto tr842
		case 65:
			goto tr843
		case 80:
			goto tr843
		case 90:
			goto tr737
		case 95:
			goto tr734
		case 97:
			goto tr844
		case 112:
			goto tr844
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr734
			}
		case data[p] >= 66:
			goto tr734
		}
		goto st0
tr838:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st564
tr842:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st564
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
//line ragel/parse_datetime.go:21819
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr846
			}
		case data[p] >= 48:
			goto tr845
		}
		goto st0
tr845:
//line ragel/datetime.rl:5
 pb = p 
	goto st565
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
//line ragel/parse_datetime.go:21838
		switch data[p] {
		case 32:
			goto tr847
		case 43:
			goto tr742
		case 45:
			goto tr744
		case 47:
			goto tr745
		case 65:
			goto tr850
		case 80:
			goto tr850
		case 90:
			goto tr748
		case 95:
			goto tr745
		case 97:
			goto tr851
		case 112:
			goto tr851
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr848
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr745
				}
			case data[p] >= 66:
				goto tr745
			}
		default:
			goto st576
		}
		goto st0
tr848:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st566
tr865:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st566
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
//line ragel/parse_datetime.go:21896
		if 48 <= data[p] && data[p] <= 57 {
			goto tr852
		}
		goto st0
tr852:
//line ragel/datetime.rl:5
 pb = p 
	goto st567
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
//line ragel/parse_datetime.go:21910
		switch data[p] {
		case 32:
			goto tr853
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr855
		case 80:
			goto tr855
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr856
		case 112:
			goto tr856
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st568
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
		switch data[p] {
		case 32:
			goto tr853
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr855
		case 80:
			goto tr855
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr856
		case 112:
			goto tr856
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st569
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		switch data[p] {
		case 32:
			goto tr853
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr855
		case 80:
			goto tr855
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr856
		case 112:
			goto tr856
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st570
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		switch data[p] {
		case 32:
			goto tr853
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr855
		case 80:
			goto tr855
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr856
		case 112:
			goto tr856
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st571
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
		switch data[p] {
		case 32:
			goto tr853
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr855
		case 80:
			goto tr855
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr856
		case 112:
			goto tr856
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st572
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
		switch data[p] {
		case 32:
			goto tr853
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr855
		case 80:
			goto tr855
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr856
		case 112:
			goto tr856
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st573
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		switch data[p] {
		case 32:
			goto tr853
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr855
		case 80:
			goto tr855
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr856
		case 112:
			goto tr856
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st574
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
		switch data[p] {
		case 32:
			goto tr853
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr855
		case 80:
			goto tr855
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr856
		case 112:
			goto tr856
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st575
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		default:
			goto tr709
		}
		goto st0
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
		switch data[p] {
		case 32:
			goto tr853
		case 43:
			goto tr707
		case 45:
			goto tr708
		case 47:
			goto tr709
		case 65:
			goto tr855
		case 80:
			goto tr855
		case 90:
			goto tr711
		case 95:
			goto tr709
		case 97:
			goto tr856
		case 112:
			goto tr856
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr709
			}
		case data[p] >= 66:
			goto tr709
		}
		goto st0
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
		switch data[p] {
		case 32:
			goto tr864
		case 43:
			goto tr763
		case 45:
			goto tr765
		case 47:
			goto tr766
		case 65:
			goto tr866
		case 80:
			goto tr866
		case 90:
			goto tr768
		case 95:
			goto tr766
		case 97:
			goto tr867
		case 112:
			goto tr867
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr865
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr766
			}
		default:
			goto tr766
		}
		goto st0
tr846:
//line ragel/datetime.rl:5
 pb = p 
	goto st577
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
//line ragel/parse_datetime.go:22311
		switch data[p] {
		case 32:
			goto tr847
		case 43:
			goto tr742
		case 45:
			goto tr744
		case 47:
			goto tr745
		case 65:
			goto tr850
		case 80:
			goto tr850
		case 90:
			goto tr748
		case 95:
			goto tr745
		case 97:
			goto tr851
		case 112:
			goto tr851
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr848
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr745
			}
		default:
			goto tr745
		}
		goto st0
tr835:
//line ragel/datetime.rl:5
 pb = p 
	goto st578
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
//line ragel/parse_datetime.go:22356
		switch data[p] {
		case 32:
			goto tr836
		case 43:
			goto tr723
		case 45:
			goto tr724
		case 47:
			goto tr725
		case 58:
			goto tr838
		case 65:
			goto tr839
		case 80:
			goto tr839
		case 90:
			goto tr729
		case 95:
			goto tr725
		case 97:
			goto tr840
		case 112:
			goto tr840
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr725
			}
		case data[p] >= 66:
			goto tr725
		}
		goto st0
tr774:
//line ragel/datetime.rl:5
 pb = p 
	goto st579
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
//line ragel/parse_datetime.go:22399
		switch data[p] {
		case 32:
			goto tr776
		case 43:
			goto tr620
		case 45:
			goto tr621
		case 47:
			goto tr622
		case 58:
			goto tr778
		case 65:
			goto tr779
		case 80:
			goto tr779
		case 90:
			goto tr626
		case 95:
			goto tr622
		case 97:
			goto tr780
		case 112:
			goto tr780
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st524
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr622
				}
			case data[p] >= 66:
				goto tr622
			}
		default:
			goto st580
		}
		goto st0
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		if 48 <= data[p] && data[p] <= 57 {
			goto st525
		}
		goto st0
tr775:
//line ragel/datetime.rl:5
 pb = p 
	goto st581
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
//line ragel/parse_datetime.go:22460
		switch data[p] {
		case 32:
			goto tr776
		case 43:
			goto tr620
		case 45:
			goto tr621
		case 47:
			goto tr622
		case 58:
			goto tr778
		case 65:
			goto tr779
		case 80:
			goto tr779
		case 90:
			goto tr626
		case 95:
			goto tr622
		case 97:
			goto tr780
		case 112:
			goto tr780
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st580
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr622
			}
		default:
			goto tr622
		}
		goto st0
tr607:
//line ragel/datetime.rl:5
 pb = p 
	goto st582
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
//line ragel/parse_datetime.go:22507
		if data[p] == 32 {
			goto tr869
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st516
			}
		case data[p] >= 45:
			goto tr418
		}
		goto st0
tr869:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st583
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
//line ragel/parse_datetime.go:22536
		if data[p] == 50 {
			goto tr871
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr872
			}
		case data[p] >= 48:
			goto tr870
		}
		goto st0
tr870:
//line ragel/datetime.rl:5
 pb = p 
	goto st584
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
//line ragel/parse_datetime.go:22558
		switch data[p] {
		case 32:
			goto tr619
		case 43:
			goto tr620
		case 45:
			goto tr621
		case 47:
			goto tr622
		case 58:
			goto tr624
		case 65:
			goto tr625
		case 80:
			goto tr625
		case 90:
			goto tr626
		case 95:
			goto tr622
		case 97:
			goto tr627
		case 112:
			goto tr627
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st585
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr622
			}
		default:
			goto tr622
		}
		goto st0
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		switch data[p] {
		case 32:
			goto tr688
		case 43:
			goto tr689
		case 45:
			goto tr690
		case 47:
			goto tr691
		case 58:
			goto tr693
		case 65:
			goto tr694
		case 80:
			goto tr694
		case 90:
			goto tr695
		case 95:
			goto tr691
		case 97:
			goto tr696
		case 112:
			goto tr696
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st586
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr691
			}
		default:
			goto tr691
		}
		goto st0
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1029
		}
		goto st0
	st1029:
		if p++; p == pe {
			goto _test_eof1029
		}
	st_case_1029:
		switch data[p] {
		case 32:
			goto tr1500
		case 43:
			goto tr1501
		case 44:
			goto tr1542
		case 45:
			goto tr1503
		case 46:
			goto tr700
		case 47:
			goto tr1504
		case 84:
			goto tr1506
		case 90:
			goto tr1507
		case 95:
			goto tr1508
		case 116:
			goto tr1508
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st490
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1504
			}
		default:
			goto tr1504
		}
		goto st0
tr1542:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

	goto st587
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
//line ragel/parse_datetime.go:22713
		if data[p] == 32 {
			goto st49
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr705
		}
		goto st0
tr871:
//line ragel/datetime.rl:5
 pb = p 
	goto st588
	st588:
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
//line ragel/parse_datetime.go:22730
		switch data[p] {
		case 32:
			goto tr619
		case 43:
			goto tr620
		case 45:
			goto tr621
		case 47:
			goto tr622
		case 58:
			goto tr624
		case 65:
			goto tr625
		case 80:
			goto tr625
		case 90:
			goto tr626
		case 95:
			goto tr622
		case 97:
			goto tr627
		case 112:
			goto tr627
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st585
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr622
				}
			case data[p] >= 66:
				goto tr622
			}
		default:
			goto st589
		}
		goto st0
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		if 48 <= data[p] && data[p] <= 57 {
			goto st586
		}
		goto st0
tr872:
//line ragel/datetime.rl:5
 pb = p 
	goto st590
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
//line ragel/parse_datetime.go:22791
		switch data[p] {
		case 32:
			goto tr619
		case 43:
			goto tr620
		case 45:
			goto tr621
		case 47:
			goto tr622
		case 58:
			goto tr624
		case 65:
			goto tr625
		case 80:
			goto tr625
		case 90:
			goto tr626
		case 95:
			goto tr622
		case 97:
			goto tr627
		case 112:
			goto tr627
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st589
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr622
			}
		default:
			goto tr622
		}
		goto st0
tr608:
//line ragel/datetime.rl:5
 pb = p 
	goto st591
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
//line ragel/parse_datetime.go:22838
		if data[p] == 32 {
			goto tr869
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st516
			}
		case data[p] >= 45:
			goto tr418
		}
		goto st0
tr609:
//line ragel/datetime.rl:5
 pb = p 
	goto st592
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
//line ragel/parse_datetime.go:22860
		if data[p] == 32 {
			goto tr869
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr418
		}
		goto st0
tr603:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st593
tr881:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st593
tr888:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st593
tr897:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st593
tr908:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st593
tr917:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st593
tr921:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st593
tr928:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st593
tr933:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st593
tr938:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st593
tr948:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st593
tr957:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st593
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
	st_case_593:
//line ragel/parse_datetime.go:22921
		switch data[p] {
		case 32:
			goto st432
		case 48:
			goto tr546
		case 51:
			goto tr548
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st393
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr549
			}
		default:
			goto tr547
		}
		goto st0
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		if data[p] == 108 {
			goto st595
		}
		goto st0
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		switch data[p] {
		case 32:
			goto tr602
		case 46:
			goto tr603
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr408
		}
		goto st0
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		if data[p] == 103 {
			goto st597
		}
		goto st0
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
	st_case_597:
		switch data[p] {
		case 32:
			goto tr880
		case 46:
			goto tr881
		case 117:
			goto st598
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr556
		}
		goto st0
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		if data[p] == 115 {
			goto st599
		}
		goto st0
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		if data[p] == 116 {
			goto st600
		}
		goto st0
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		switch data[p] {
		case 32:
			goto tr880
		case 46:
			goto tr881
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr556
		}
		goto st0
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
		if data[p] == 101 {
			goto st602
		}
		goto st0
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
		if data[p] == 99 {
			goto st603
		}
		goto st0
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
		switch data[p] {
		case 32:
			goto tr887
		case 46:
			goto tr888
		case 101:
			goto st604
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr564
		}
		goto st0
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		if data[p] == 109 {
			goto st605
		}
		goto st0
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
		if data[p] == 98 {
			goto st606
		}
		goto st0
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
		if data[p] == 101 {
			goto st607
		}
		goto st0
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
		if data[p] == 114 {
			goto st608
		}
		goto st0
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		switch data[p] {
		case 32:
			goto tr887
		case 46:
			goto tr888
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr564
		}
		goto st0
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
		if data[p] == 101 {
			goto st610
		}
		goto st0
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		if data[p] == 98 {
			goto st611
		}
		goto st0
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		switch data[p] {
		case 32:
			goto tr896
		case 46:
			goto tr897
		case 114:
			goto st612
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr575
		}
		goto st0
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
		if data[p] == 117 {
			goto st613
		}
		goto st0
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		if data[p] == 97 {
			goto st614
		}
		goto st0
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
		if data[p] == 114 {
			goto st615
		}
		goto st0
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
		if data[p] == 121 {
			goto st616
		}
		goto st0
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
		switch data[p] {
		case 32:
			goto tr896
		case 46:
			goto tr897
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr575
		}
		goto st0
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
		switch data[p] {
		case 97:
			goto st618
		case 117:
			goto st624
		}
		goto st0
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
		if data[p] == 110 {
			goto st619
		}
		goto st0
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		switch data[p] {
		case 32:
			goto tr906
		case 46:
			goto tr908
		case 117:
			goto st620
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr907
		}
		goto st0
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		if data[p] == 97 {
			goto st621
		}
		goto st0
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		if data[p] == 114 {
			goto st622
		}
		goto st0
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		if data[p] == 121 {
			goto st623
		}
		goto st0
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
		switch data[p] {
		case 32:
			goto tr906
		case 46:
			goto tr908
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr907
		}
		goto st0
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
		switch data[p] {
		case 108:
			goto st625
		case 110:
			goto st627
		}
		goto st0
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
		switch data[p] {
		case 32:
			goto tr915
		case 46:
			goto tr917
		case 121:
			goto st626
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr916
		}
		goto st0
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
		switch data[p] {
		case 32:
			goto tr915
		case 46:
			goto tr917
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr916
		}
		goto st0
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
		switch data[p] {
		case 32:
			goto tr919
		case 46:
			goto tr921
		case 101:
			goto st628
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr920
		}
		goto st0
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
		switch data[p] {
		case 32:
			goto tr919
		case 46:
			goto tr921
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr920
		}
		goto st0
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
		if data[p] == 97 {
			goto st630
		}
		goto st0
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
		switch data[p] {
		case 114:
			goto st631
		case 121:
			goto st634
		}
		goto st0
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		switch data[p] {
		case 32:
			goto tr926
		case 46:
			goto tr928
		case 99:
			goto st632
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr927
		}
		goto st0
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
		if data[p] == 104 {
			goto st633
		}
		goto st0
	st633:
		if p++; p == pe {
			goto _test_eof633
		}
	st_case_633:
		switch data[p] {
		case 32:
			goto tr926
		case 46:
			goto tr928
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr927
		}
		goto st0
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
		switch data[p] {
		case 32:
			goto tr931
		case 46:
			goto tr933
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr932
		}
		goto st0
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
		if data[p] == 111 {
			goto st636
		}
		goto st0
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
		if data[p] == 118 {
			goto st637
		}
		goto st0
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
		switch data[p] {
		case 32:
			goto tr936
		case 46:
			goto tr938
		case 101:
			goto st638
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr937
		}
		goto st0
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
		if data[p] == 109 {
			goto st639
		}
		goto st0
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
		if data[p] == 98 {
			goto st640
		}
		goto st0
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
		if data[p] == 101 {
			goto st641
		}
		goto st0
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
		if data[p] == 114 {
			goto st642
		}
		goto st0
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
		switch data[p] {
		case 32:
			goto tr936
		case 46:
			goto tr938
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr937
		}
		goto st0
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
		if data[p] == 99 {
			goto st644
		}
		goto st0
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
		if data[p] == 116 {
			goto st645
		}
		goto st0
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
		switch data[p] {
		case 32:
			goto tr946
		case 46:
			goto tr948
		case 111:
			goto st646
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr947
		}
		goto st0
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
		if data[p] == 98 {
			goto st647
		}
		goto st0
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
		if data[p] == 101 {
			goto st648
		}
		goto st0
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
		if data[p] == 114 {
			goto st649
		}
		goto st0
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
		switch data[p] {
		case 32:
			goto tr946
		case 46:
			goto tr948
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr947
		}
		goto st0
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
		if data[p] == 101 {
			goto st651
		}
		goto st0
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		if data[p] == 112 {
			goto st652
		}
		goto st0
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
		switch data[p] {
		case 32:
			goto tr955
		case 46:
			goto tr957
		case 116:
			goto st653
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr956
		}
		goto st0
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		switch data[p] {
		case 32:
			goto tr955
		case 46:
			goto tr957
		case 101:
			goto st654
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr956
		}
		goto st0
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
		if data[p] == 109 {
			goto st655
		}
		goto st0
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
		if data[p] == 98 {
			goto st656
		}
		goto st0
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
		if data[p] == 101 {
			goto st657
		}
		goto st0
	st657:
		if p++; p == pe {
			goto _test_eof657
		}
	st_case_657:
		if data[p] == 114 {
			goto st658
		}
		goto st0
	st658:
		if p++; p == pe {
			goto _test_eof658
		}
	st_case_658:
		switch data[p] {
		case 32:
			goto tr955
		case 46:
			goto tr957
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr956
		}
		goto st0
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
		if data[p] == 32 {
			goto st660
		}
		goto st0
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
		switch data[p] {
		case 48:
			goto tr965
		case 51:
			goto tr967
		case 65:
			goto st735
		case 68:
			goto st746
		case 70:
			goto st754
		case 74:
			goto st762
		case 77:
			goto st774
		case 78:
			goto st780
		case 79:
			goto st788
		case 83:
			goto st795
		case 97:
			goto st735
		case 100:
			goto st746
		case 102:
			goto st754
		case 106:
			goto st762
		case 109:
			goto st774
		case 110:
			goto st780
		case 111:
			goto st788
		case 115:
			goto st795
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr968
			}
		case data[p] >= 49:
			goto tr966
		}
		goto st0
tr965:
//line ragel/datetime.rl:5
 pb = p 
	goto st661
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
	st_case_661:
//line ragel/parse_datetime.go:23767
		if 49 <= data[p] && data[p] <= 57 {
			goto st662
		}
		goto st0
tr968:
//line ragel/datetime.rl:5
 pb = p 
	goto st662
	st662:
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
//line ragel/parse_datetime.go:23781
		switch data[p] {
		case 32:
			goto tr208
		case 45:
			goto tr978
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto tr208
		}
		goto st0
tr978:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st663
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
//line ragel/parse_datetime.go:23808
		switch data[p] {
		case 65:
			goto st664
		case 68:
			goto st675
		case 70:
			goto st683
		case 74:
			goto st691
		case 77:
			goto st703
		case 78:
			goto st709
		case 79:
			goto st717
		case 83:
			goto st724
		case 97:
			goto st664
		case 100:
			goto st675
		case 102:
			goto st683
		case 106:
			goto st691
		case 109:
			goto st703
		case 110:
			goto st709
		case 111:
			goto st717
		case 115:
			goto st724
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr209
		}
		goto st0
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
		switch data[p] {
		case 112:
			goto st665
		case 117:
			goto st670
		}
		goto st0
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
		if data[p] == 114 {
			goto st666
		}
		goto st0
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
		switch data[p] {
		case 32:
			goto tr228
		case 45:
			goto tr227
		case 46:
			goto tr990
		case 47:
			goto tr228
		case 105:
			goto st668
		}
		goto st0
tr990:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st667
tr994:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st667
tr1000:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st667
tr1008:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st667
tr1017:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st667
tr1024:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st667
tr1026:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st667
tr1031:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st667
tr1034:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st667
tr1037:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st667
tr1045:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st667
tr1052:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st667
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
	st_case_667:
//line ragel/parse_datetime.go:23939
		switch data[p] {
		case 32:
			goto st149
		case 45:
			goto st157
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr220
			}
		case data[p] >= 46:
			goto st149
		}
		goto st0
	st668:
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
		if data[p] == 108 {
			goto st669
		}
		goto st0
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
		switch data[p] {
		case 32:
			goto tr228
		case 45:
			goto tr227
		case 46:
			goto tr990
		case 47:
			goto tr228
		}
		goto st0
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
		if data[p] == 103 {
			goto st671
		}
		goto st0
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
		switch data[p] {
		case 32:
			goto tr238
		case 45:
			goto tr237
		case 46:
			goto tr994
		case 47:
			goto tr238
		case 117:
			goto st672
		}
		goto st0
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
	st_case_672:
		if data[p] == 115 {
			goto st673
		}
		goto st0
	st673:
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
		if data[p] == 116 {
			goto st674
		}
		goto st0
	st674:
		if p++; p == pe {
			goto _test_eof674
		}
	st_case_674:
		switch data[p] {
		case 32:
			goto tr238
		case 45:
			goto tr237
		case 46:
			goto tr994
		case 47:
			goto tr238
		}
		goto st0
	st675:
		if p++; p == pe {
			goto _test_eof675
		}
	st_case_675:
		if data[p] == 101 {
			goto st676
		}
		goto st0
	st676:
		if p++; p == pe {
			goto _test_eof676
		}
	st_case_676:
		if data[p] == 99 {
			goto st677
		}
		goto st0
	st677:
		if p++; p == pe {
			goto _test_eof677
		}
	st_case_677:
		switch data[p] {
		case 32:
			goto tr246
		case 45:
			goto tr245
		case 46:
			goto tr1000
		case 47:
			goto tr246
		case 101:
			goto st678
		}
		goto st0
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
		if data[p] == 109 {
			goto st679
		}
		goto st0
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
		if data[p] == 98 {
			goto st680
		}
		goto st0
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
		if data[p] == 101 {
			goto st681
		}
		goto st0
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
	st_case_681:
		if data[p] == 114 {
			goto st682
		}
		goto st0
	st682:
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
		switch data[p] {
		case 32:
			goto tr246
		case 45:
			goto tr245
		case 46:
			goto tr1000
		case 47:
			goto tr246
		}
		goto st0
	st683:
		if p++; p == pe {
			goto _test_eof683
		}
	st_case_683:
		if data[p] == 101 {
			goto st684
		}
		goto st0
	st684:
		if p++; p == pe {
			goto _test_eof684
		}
	st_case_684:
		if data[p] == 98 {
			goto st685
		}
		goto st0
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
	st_case_685:
		switch data[p] {
		case 32:
			goto tr256
		case 45:
			goto tr255
		case 46:
			goto tr1008
		case 47:
			goto tr256
		case 114:
			goto st686
		}
		goto st0
	st686:
		if p++; p == pe {
			goto _test_eof686
		}
	st_case_686:
		if data[p] == 117 {
			goto st687
		}
		goto st0
	st687:
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
		if data[p] == 97 {
			goto st688
		}
		goto st0
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
		if data[p] == 114 {
			goto st689
		}
		goto st0
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
		if data[p] == 121 {
			goto st690
		}
		goto st0
	st690:
		if p++; p == pe {
			goto _test_eof690
		}
	st_case_690:
		switch data[p] {
		case 32:
			goto tr256
		case 45:
			goto tr255
		case 46:
			goto tr1008
		case 47:
			goto tr256
		}
		goto st0
	st691:
		if p++; p == pe {
			goto _test_eof691
		}
	st_case_691:
		switch data[p] {
		case 97:
			goto st692
		case 117:
			goto st698
		}
		goto st0
	st692:
		if p++; p == pe {
			goto _test_eof692
		}
	st_case_692:
		if data[p] == 110 {
			goto st693
		}
		goto st0
	st693:
		if p++; p == pe {
			goto _test_eof693
		}
	st_case_693:
		switch data[p] {
		case 32:
			goto tr267
		case 45:
			goto tr266
		case 46:
			goto tr1017
		case 47:
			goto tr267
		case 117:
			goto st694
		}
		goto st0
	st694:
		if p++; p == pe {
			goto _test_eof694
		}
	st_case_694:
		if data[p] == 97 {
			goto st695
		}
		goto st0
	st695:
		if p++; p == pe {
			goto _test_eof695
		}
	st_case_695:
		if data[p] == 114 {
			goto st696
		}
		goto st0
	st696:
		if p++; p == pe {
			goto _test_eof696
		}
	st_case_696:
		if data[p] == 121 {
			goto st697
		}
		goto st0
	st697:
		if p++; p == pe {
			goto _test_eof697
		}
	st_case_697:
		switch data[p] {
		case 32:
			goto tr267
		case 45:
			goto tr266
		case 46:
			goto tr1017
		case 47:
			goto tr267
		}
		goto st0
	st698:
		if p++; p == pe {
			goto _test_eof698
		}
	st_case_698:
		switch data[p] {
		case 108:
			goto st699
		case 110:
			goto st701
		}
		goto st0
	st699:
		if p++; p == pe {
			goto _test_eof699
		}
	st_case_699:
		switch data[p] {
		case 32:
			goto tr276
		case 45:
			goto tr275
		case 46:
			goto tr1024
		case 47:
			goto tr276
		case 121:
			goto st700
		}
		goto st0
	st700:
		if p++; p == pe {
			goto _test_eof700
		}
	st_case_700:
		switch data[p] {
		case 32:
			goto tr276
		case 45:
			goto tr275
		case 46:
			goto tr1024
		case 47:
			goto tr276
		}
		goto st0
	st701:
		if p++; p == pe {
			goto _test_eof701
		}
	st_case_701:
		switch data[p] {
		case 32:
			goto tr280
		case 45:
			goto tr279
		case 46:
			goto tr1026
		case 47:
			goto tr280
		case 101:
			goto st702
		}
		goto st0
	st702:
		if p++; p == pe {
			goto _test_eof702
		}
	st_case_702:
		switch data[p] {
		case 32:
			goto tr280
		case 45:
			goto tr279
		case 46:
			goto tr1026
		case 47:
			goto tr280
		}
		goto st0
	st703:
		if p++; p == pe {
			goto _test_eof703
		}
	st_case_703:
		if data[p] == 97 {
			goto st704
		}
		goto st0
	st704:
		if p++; p == pe {
			goto _test_eof704
		}
	st_case_704:
		switch data[p] {
		case 114:
			goto st705
		case 121:
			goto st708
		}
		goto st0
	st705:
		if p++; p == pe {
			goto _test_eof705
		}
	st_case_705:
		switch data[p] {
		case 32:
			goto tr287
		case 45:
			goto tr286
		case 46:
			goto tr1031
		case 47:
			goto tr287
		case 99:
			goto st706
		}
		goto st0
	st706:
		if p++; p == pe {
			goto _test_eof706
		}
	st_case_706:
		if data[p] == 104 {
			goto st707
		}
		goto st0
	st707:
		if p++; p == pe {
			goto _test_eof707
		}
	st_case_707:
		switch data[p] {
		case 32:
			goto tr287
		case 45:
			goto tr286
		case 46:
			goto tr1031
		case 47:
			goto tr287
		}
		goto st0
	st708:
		if p++; p == pe {
			goto _test_eof708
		}
	st_case_708:
		switch data[p] {
		case 32:
			goto tr292
		case 45:
			goto tr291
		case 46:
			goto tr1034
		case 47:
			goto tr292
		}
		goto st0
	st709:
		if p++; p == pe {
			goto _test_eof709
		}
	st_case_709:
		if data[p] == 111 {
			goto st710
		}
		goto st0
	st710:
		if p++; p == pe {
			goto _test_eof710
		}
	st_case_710:
		if data[p] == 118 {
			goto st711
		}
		goto st0
	st711:
		if p++; p == pe {
			goto _test_eof711
		}
	st_case_711:
		switch data[p] {
		case 32:
			goto tr297
		case 45:
			goto tr296
		case 46:
			goto tr1037
		case 47:
			goto tr297
		case 101:
			goto st712
		}
		goto st0
	st712:
		if p++; p == pe {
			goto _test_eof712
		}
	st_case_712:
		if data[p] == 109 {
			goto st713
		}
		goto st0
	st713:
		if p++; p == pe {
			goto _test_eof713
		}
	st_case_713:
		if data[p] == 98 {
			goto st714
		}
		goto st0
	st714:
		if p++; p == pe {
			goto _test_eof714
		}
	st_case_714:
		if data[p] == 101 {
			goto st715
		}
		goto st0
	st715:
		if p++; p == pe {
			goto _test_eof715
		}
	st_case_715:
		if data[p] == 114 {
			goto st716
		}
		goto st0
	st716:
		if p++; p == pe {
			goto _test_eof716
		}
	st_case_716:
		switch data[p] {
		case 32:
			goto tr297
		case 45:
			goto tr296
		case 46:
			goto tr1037
		case 47:
			goto tr297
		}
		goto st0
	st717:
		if p++; p == pe {
			goto _test_eof717
		}
	st_case_717:
		if data[p] == 99 {
			goto st718
		}
		goto st0
	st718:
		if p++; p == pe {
			goto _test_eof718
		}
	st_case_718:
		if data[p] == 116 {
			goto st719
		}
		goto st0
	st719:
		if p++; p == pe {
			goto _test_eof719
		}
	st_case_719:
		switch data[p] {
		case 32:
			goto tr307
		case 45:
			goto tr306
		case 46:
			goto tr1045
		case 47:
			goto tr307
		case 111:
			goto st720
		}
		goto st0
	st720:
		if p++; p == pe {
			goto _test_eof720
		}
	st_case_720:
		if data[p] == 98 {
			goto st721
		}
		goto st0
	st721:
		if p++; p == pe {
			goto _test_eof721
		}
	st_case_721:
		if data[p] == 101 {
			goto st722
		}
		goto st0
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
		if data[p] == 114 {
			goto st723
		}
		goto st0
	st723:
		if p++; p == pe {
			goto _test_eof723
		}
	st_case_723:
		switch data[p] {
		case 32:
			goto tr307
		case 45:
			goto tr306
		case 46:
			goto tr1045
		case 47:
			goto tr307
		}
		goto st0
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
	st_case_724:
		if data[p] == 101 {
			goto st725
		}
		goto st0
	st725:
		if p++; p == pe {
			goto _test_eof725
		}
	st_case_725:
		if data[p] == 112 {
			goto st726
		}
		goto st0
	st726:
		if p++; p == pe {
			goto _test_eof726
		}
	st_case_726:
		switch data[p] {
		case 32:
			goto tr316
		case 45:
			goto tr315
		case 46:
			goto tr1052
		case 47:
			goto tr316
		case 116:
			goto st727
		}
		goto st0
	st727:
		if p++; p == pe {
			goto _test_eof727
		}
	st_case_727:
		switch data[p] {
		case 32:
			goto tr316
		case 45:
			goto tr315
		case 46:
			goto tr1052
		case 47:
			goto tr316
		case 101:
			goto st728
		}
		goto st0
	st728:
		if p++; p == pe {
			goto _test_eof728
		}
	st_case_728:
		if data[p] == 109 {
			goto st729
		}
		goto st0
	st729:
		if p++; p == pe {
			goto _test_eof729
		}
	st_case_729:
		if data[p] == 98 {
			goto st730
		}
		goto st0
	st730:
		if p++; p == pe {
			goto _test_eof730
		}
	st_case_730:
		if data[p] == 101 {
			goto st731
		}
		goto st0
	st731:
		if p++; p == pe {
			goto _test_eof731
		}
	st_case_731:
		if data[p] == 114 {
			goto st732
		}
		goto st0
	st732:
		if p++; p == pe {
			goto _test_eof732
		}
	st_case_732:
		switch data[p] {
		case 32:
			goto tr316
		case 45:
			goto tr315
		case 46:
			goto tr1052
		case 47:
			goto tr316
		}
		goto st0
tr966:
//line ragel/datetime.rl:5
 pb = p 
	goto st733
	st733:
		if p++; p == pe {
			goto _test_eof733
		}
	st_case_733:
//line ragel/parse_datetime.go:24741
		switch data[p] {
		case 32:
			goto tr208
		case 45:
			goto tr978
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st662
			}
		case data[p] >= 46:
			goto tr208
		}
		goto st0
tr967:
//line ragel/datetime.rl:5
 pb = p 
	goto st734
	st734:
		if p++; p == pe {
			goto _test_eof734
		}
	st_case_734:
//line ragel/parse_datetime.go:24766
		switch data[p] {
		case 32:
			goto tr208
		case 45:
			goto tr978
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st662
			}
		case data[p] >= 46:
			goto tr208
		}
		goto st0
	st735:
		if p++; p == pe {
			goto _test_eof735
		}
	st_case_735:
		switch data[p] {
		case 112:
			goto st736
		case 117:
			goto st741
		}
		goto st0
	st736:
		if p++; p == pe {
			goto _test_eof736
		}
	st_case_736:
		if data[p] == 114 {
			goto st737
		}
		goto st0
	st737:
		if p++; p == pe {
			goto _test_eof737
		}
	st_case_737:
		switch data[p] {
		case 32:
			goto tr408
		case 46:
			goto tr1062
		case 105:
			goto st739
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr408
		}
		goto st0
tr1062:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st738
tr1066:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st738
tr1072:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st738
tr1080:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st738
tr1089:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st738
tr1096:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st738
tr1098:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st738
tr1103:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st738
tr1106:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st738
tr1109:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st738
tr1117:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st738
tr1124:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st738
	st738:
		if p++; p == pe {
			goto _test_eof738
		}
	st_case_738:
//line ragel/parse_datetime.go:24873
		switch data[p] {
		case 32:
			goto st393
		case 48:
			goto tr546
		case 51:
			goto tr548
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st393
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr549
			}
		default:
			goto tr547
		}
		goto st0
	st739:
		if p++; p == pe {
			goto _test_eof739
		}
	st_case_739:
		if data[p] == 108 {
			goto st740
		}
		goto st0
	st740:
		if p++; p == pe {
			goto _test_eof740
		}
	st_case_740:
		switch data[p] {
		case 32:
			goto tr408
		case 46:
			goto tr1062
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr408
		}
		goto st0
	st741:
		if p++; p == pe {
			goto _test_eof741
		}
	st_case_741:
		if data[p] == 103 {
			goto st742
		}
		goto st0
	st742:
		if p++; p == pe {
			goto _test_eof742
		}
	st_case_742:
		switch data[p] {
		case 32:
			goto tr556
		case 46:
			goto tr1066
		case 117:
			goto st743
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr556
		}
		goto st0
	st743:
		if p++; p == pe {
			goto _test_eof743
		}
	st_case_743:
		if data[p] == 115 {
			goto st744
		}
		goto st0
	st744:
		if p++; p == pe {
			goto _test_eof744
		}
	st_case_744:
		if data[p] == 116 {
			goto st745
		}
		goto st0
	st745:
		if p++; p == pe {
			goto _test_eof745
		}
	st_case_745:
		switch data[p] {
		case 32:
			goto tr556
		case 46:
			goto tr1066
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr556
		}
		goto st0
	st746:
		if p++; p == pe {
			goto _test_eof746
		}
	st_case_746:
		if data[p] == 101 {
			goto st747
		}
		goto st0
	st747:
		if p++; p == pe {
			goto _test_eof747
		}
	st_case_747:
		if data[p] == 99 {
			goto st748
		}
		goto st0
	st748:
		if p++; p == pe {
			goto _test_eof748
		}
	st_case_748:
		switch data[p] {
		case 32:
			goto tr564
		case 46:
			goto tr1072
		case 101:
			goto st749
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr564
		}
		goto st0
	st749:
		if p++; p == pe {
			goto _test_eof749
		}
	st_case_749:
		if data[p] == 109 {
			goto st750
		}
		goto st0
	st750:
		if p++; p == pe {
			goto _test_eof750
		}
	st_case_750:
		if data[p] == 98 {
			goto st751
		}
		goto st0
	st751:
		if p++; p == pe {
			goto _test_eof751
		}
	st_case_751:
		if data[p] == 101 {
			goto st752
		}
		goto st0
	st752:
		if p++; p == pe {
			goto _test_eof752
		}
	st_case_752:
		if data[p] == 114 {
			goto st753
		}
		goto st0
	st753:
		if p++; p == pe {
			goto _test_eof753
		}
	st_case_753:
		switch data[p] {
		case 32:
			goto tr564
		case 46:
			goto tr1072
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr564
		}
		goto st0
	st754:
		if p++; p == pe {
			goto _test_eof754
		}
	st_case_754:
		if data[p] == 101 {
			goto st755
		}
		goto st0
	st755:
		if p++; p == pe {
			goto _test_eof755
		}
	st_case_755:
		if data[p] == 98 {
			goto st756
		}
		goto st0
	st756:
		if p++; p == pe {
			goto _test_eof756
		}
	st_case_756:
		switch data[p] {
		case 32:
			goto tr575
		case 46:
			goto tr1080
		case 114:
			goto st757
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr575
		}
		goto st0
	st757:
		if p++; p == pe {
			goto _test_eof757
		}
	st_case_757:
		if data[p] == 117 {
			goto st758
		}
		goto st0
	st758:
		if p++; p == pe {
			goto _test_eof758
		}
	st_case_758:
		if data[p] == 97 {
			goto st759
		}
		goto st0
	st759:
		if p++; p == pe {
			goto _test_eof759
		}
	st_case_759:
		if data[p] == 114 {
			goto st760
		}
		goto st0
	st760:
		if p++; p == pe {
			goto _test_eof760
		}
	st_case_760:
		if data[p] == 121 {
			goto st761
		}
		goto st0
	st761:
		if p++; p == pe {
			goto _test_eof761
		}
	st_case_761:
		switch data[p] {
		case 32:
			goto tr575
		case 46:
			goto tr1080
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr575
		}
		goto st0
	st762:
		if p++; p == pe {
			goto _test_eof762
		}
	st_case_762:
		switch data[p] {
		case 97:
			goto st763
		case 117:
			goto st769
		}
		goto st0
	st763:
		if p++; p == pe {
			goto _test_eof763
		}
	st_case_763:
		if data[p] == 110 {
			goto st764
		}
		goto st0
	st764:
		if p++; p == pe {
			goto _test_eof764
		}
	st_case_764:
		switch data[p] {
		case 32:
			goto tr907
		case 46:
			goto tr1089
		case 117:
			goto st765
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr907
		}
		goto st0
	st765:
		if p++; p == pe {
			goto _test_eof765
		}
	st_case_765:
		if data[p] == 97 {
			goto st766
		}
		goto st0
	st766:
		if p++; p == pe {
			goto _test_eof766
		}
	st_case_766:
		if data[p] == 114 {
			goto st767
		}
		goto st0
	st767:
		if p++; p == pe {
			goto _test_eof767
		}
	st_case_767:
		if data[p] == 121 {
			goto st768
		}
		goto st0
	st768:
		if p++; p == pe {
			goto _test_eof768
		}
	st_case_768:
		switch data[p] {
		case 32:
			goto tr907
		case 46:
			goto tr1089
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr907
		}
		goto st0
	st769:
		if p++; p == pe {
			goto _test_eof769
		}
	st_case_769:
		switch data[p] {
		case 108:
			goto st770
		case 110:
			goto st772
		}
		goto st0
	st770:
		if p++; p == pe {
			goto _test_eof770
		}
	st_case_770:
		switch data[p] {
		case 32:
			goto tr916
		case 46:
			goto tr1096
		case 121:
			goto st771
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr916
		}
		goto st0
	st771:
		if p++; p == pe {
			goto _test_eof771
		}
	st_case_771:
		switch data[p] {
		case 32:
			goto tr916
		case 46:
			goto tr1096
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr916
		}
		goto st0
	st772:
		if p++; p == pe {
			goto _test_eof772
		}
	st_case_772:
		switch data[p] {
		case 32:
			goto tr920
		case 46:
			goto tr1098
		case 101:
			goto st773
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr920
		}
		goto st0
	st773:
		if p++; p == pe {
			goto _test_eof773
		}
	st_case_773:
		switch data[p] {
		case 32:
			goto tr920
		case 46:
			goto tr1098
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr920
		}
		goto st0
	st774:
		if p++; p == pe {
			goto _test_eof774
		}
	st_case_774:
		if data[p] == 97 {
			goto st775
		}
		goto st0
	st775:
		if p++; p == pe {
			goto _test_eof775
		}
	st_case_775:
		switch data[p] {
		case 114:
			goto st776
		case 121:
			goto st779
		}
		goto st0
	st776:
		if p++; p == pe {
			goto _test_eof776
		}
	st_case_776:
		switch data[p] {
		case 32:
			goto tr927
		case 46:
			goto tr1103
		case 99:
			goto st777
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr927
		}
		goto st0
	st777:
		if p++; p == pe {
			goto _test_eof777
		}
	st_case_777:
		if data[p] == 104 {
			goto st778
		}
		goto st0
	st778:
		if p++; p == pe {
			goto _test_eof778
		}
	st_case_778:
		switch data[p] {
		case 32:
			goto tr927
		case 46:
			goto tr1103
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr927
		}
		goto st0
	st779:
		if p++; p == pe {
			goto _test_eof779
		}
	st_case_779:
		switch data[p] {
		case 32:
			goto tr932
		case 46:
			goto tr1106
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr932
		}
		goto st0
	st780:
		if p++; p == pe {
			goto _test_eof780
		}
	st_case_780:
		if data[p] == 111 {
			goto st781
		}
		goto st0
	st781:
		if p++; p == pe {
			goto _test_eof781
		}
	st_case_781:
		if data[p] == 118 {
			goto st782
		}
		goto st0
	st782:
		if p++; p == pe {
			goto _test_eof782
		}
	st_case_782:
		switch data[p] {
		case 32:
			goto tr937
		case 46:
			goto tr1109
		case 101:
			goto st783
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr937
		}
		goto st0
	st783:
		if p++; p == pe {
			goto _test_eof783
		}
	st_case_783:
		if data[p] == 109 {
			goto st784
		}
		goto st0
	st784:
		if p++; p == pe {
			goto _test_eof784
		}
	st_case_784:
		if data[p] == 98 {
			goto st785
		}
		goto st0
	st785:
		if p++; p == pe {
			goto _test_eof785
		}
	st_case_785:
		if data[p] == 101 {
			goto st786
		}
		goto st0
	st786:
		if p++; p == pe {
			goto _test_eof786
		}
	st_case_786:
		if data[p] == 114 {
			goto st787
		}
		goto st0
	st787:
		if p++; p == pe {
			goto _test_eof787
		}
	st_case_787:
		switch data[p] {
		case 32:
			goto tr937
		case 46:
			goto tr1109
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr937
		}
		goto st0
	st788:
		if p++; p == pe {
			goto _test_eof788
		}
	st_case_788:
		if data[p] == 99 {
			goto st789
		}
		goto st0
	st789:
		if p++; p == pe {
			goto _test_eof789
		}
	st_case_789:
		if data[p] == 116 {
			goto st790
		}
		goto st0
	st790:
		if p++; p == pe {
			goto _test_eof790
		}
	st_case_790:
		switch data[p] {
		case 32:
			goto tr947
		case 46:
			goto tr1117
		case 111:
			goto st791
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr947
		}
		goto st0
	st791:
		if p++; p == pe {
			goto _test_eof791
		}
	st_case_791:
		if data[p] == 98 {
			goto st792
		}
		goto st0
	st792:
		if p++; p == pe {
			goto _test_eof792
		}
	st_case_792:
		if data[p] == 101 {
			goto st793
		}
		goto st0
	st793:
		if p++; p == pe {
			goto _test_eof793
		}
	st_case_793:
		if data[p] == 114 {
			goto st794
		}
		goto st0
	st794:
		if p++; p == pe {
			goto _test_eof794
		}
	st_case_794:
		switch data[p] {
		case 32:
			goto tr947
		case 46:
			goto tr1117
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr947
		}
		goto st0
	st795:
		if p++; p == pe {
			goto _test_eof795
		}
	st_case_795:
		if data[p] == 101 {
			goto st796
		}
		goto st0
	st796:
		if p++; p == pe {
			goto _test_eof796
		}
	st_case_796:
		if data[p] == 112 {
			goto st797
		}
		goto st0
	st797:
		if p++; p == pe {
			goto _test_eof797
		}
	st_case_797:
		switch data[p] {
		case 32:
			goto tr956
		case 46:
			goto tr1124
		case 116:
			goto st798
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr956
		}
		goto st0
	st798:
		if p++; p == pe {
			goto _test_eof798
		}
	st_case_798:
		switch data[p] {
		case 32:
			goto tr956
		case 46:
			goto tr1124
		case 101:
			goto st799
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr956
		}
		goto st0
	st799:
		if p++; p == pe {
			goto _test_eof799
		}
	st_case_799:
		if data[p] == 109 {
			goto st800
		}
		goto st0
	st800:
		if p++; p == pe {
			goto _test_eof800
		}
	st_case_800:
		if data[p] == 98 {
			goto st801
		}
		goto st0
	st801:
		if p++; p == pe {
			goto _test_eof801
		}
	st_case_801:
		if data[p] == 101 {
			goto st802
		}
		goto st0
	st802:
		if p++; p == pe {
			goto _test_eof802
		}
	st_case_802:
		if data[p] == 114 {
			goto st803
		}
		goto st0
	st803:
		if p++; p == pe {
			goto _test_eof803
		}
	st_case_803:
		switch data[p] {
		case 32:
			goto tr956
		case 46:
			goto tr1124
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr956
		}
		goto st0
	st804:
		if p++; p == pe {
			goto _test_eof804
		}
	st_case_804:
		if data[p] == 97 {
			goto st805
		}
		goto st0
	st805:
		if p++; p == pe {
			goto _test_eof805
		}
	st_case_805:
		if data[p] == 121 {
			goto st806
		}
		goto st0
	st806:
		if p++; p == pe {
			goto _test_eof806
		}
	st_case_806:
		switch data[p] {
		case 32:
			goto st424
		case 44:
			goto st659
		}
		goto st0
	st807:
		if p++; p == pe {
			goto _test_eof807
		}
	st_case_807:
		switch data[p] {
		case 97:
			goto st808
		case 117:
			goto st814
		}
		goto st0
	st808:
		if p++; p == pe {
			goto _test_eof808
		}
	st_case_808:
		if data[p] == 110 {
			goto st809
		}
		goto st0
	st809:
		if p++; p == pe {
			goto _test_eof809
		}
	st_case_809:
		switch data[p] {
		case 32:
			goto tr1136
		case 46:
			goto tr1137
		case 117:
			goto st810
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr907
		}
		goto st0
	st810:
		if p++; p == pe {
			goto _test_eof810
		}
	st_case_810:
		if data[p] == 97 {
			goto st811
		}
		goto st0
	st811:
		if p++; p == pe {
			goto _test_eof811
		}
	st_case_811:
		if data[p] == 114 {
			goto st812
		}
		goto st0
	st812:
		if p++; p == pe {
			goto _test_eof812
		}
	st_case_812:
		if data[p] == 121 {
			goto st813
		}
		goto st0
	st813:
		if p++; p == pe {
			goto _test_eof813
		}
	st_case_813:
		switch data[p] {
		case 32:
			goto tr1136
		case 46:
			goto tr1137
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr907
		}
		goto st0
	st814:
		if p++; p == pe {
			goto _test_eof814
		}
	st_case_814:
		switch data[p] {
		case 108:
			goto st815
		case 110:
			goto st817
		}
		goto st0
	st815:
		if p++; p == pe {
			goto _test_eof815
		}
	st_case_815:
		switch data[p] {
		case 32:
			goto tr1144
		case 46:
			goto tr1145
		case 121:
			goto st816
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr916
		}
		goto st0
	st816:
		if p++; p == pe {
			goto _test_eof816
		}
	st_case_816:
		switch data[p] {
		case 32:
			goto tr1144
		case 46:
			goto tr1145
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr916
		}
		goto st0
	st817:
		if p++; p == pe {
			goto _test_eof817
		}
	st_case_817:
		switch data[p] {
		case 32:
			goto tr1147
		case 46:
			goto tr1148
		case 101:
			goto st818
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr920
		}
		goto st0
	st818:
		if p++; p == pe {
			goto _test_eof818
		}
	st_case_818:
		switch data[p] {
		case 32:
			goto tr1147
		case 46:
			goto tr1148
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr920
		}
		goto st0
	st819:
		if p++; p == pe {
			goto _test_eof819
		}
	st_case_819:
		switch data[p] {
		case 97:
			goto st820
		case 111:
			goto st825
		}
		goto st0
	st820:
		if p++; p == pe {
			goto _test_eof820
		}
	st_case_820:
		switch data[p] {
		case 114:
			goto st821
		case 121:
			goto st824
		}
		goto st0
	st821:
		if p++; p == pe {
			goto _test_eof821
		}
	st_case_821:
		switch data[p] {
		case 32:
			goto tr1154
		case 46:
			goto tr1155
		case 99:
			goto st822
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr927
		}
		goto st0
	st822:
		if p++; p == pe {
			goto _test_eof822
		}
	st_case_822:
		if data[p] == 104 {
			goto st823
		}
		goto st0
	st823:
		if p++; p == pe {
			goto _test_eof823
		}
	st_case_823:
		switch data[p] {
		case 32:
			goto tr1154
		case 46:
			goto tr1155
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr927
		}
		goto st0
	st824:
		if p++; p == pe {
			goto _test_eof824
		}
	st_case_824:
		switch data[p] {
		case 32:
			goto tr1158
		case 46:
			goto tr1159
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr932
		}
		goto st0
	st825:
		if p++; p == pe {
			goto _test_eof825
		}
	st_case_825:
		if data[p] == 110 {
			goto st423
		}
		goto st0
	st826:
		if p++; p == pe {
			goto _test_eof826
		}
	st_case_826:
		if data[p] == 111 {
			goto st827
		}
		goto st0
	st827:
		if p++; p == pe {
			goto _test_eof827
		}
	st_case_827:
		if data[p] == 118 {
			goto st828
		}
		goto st0
	st828:
		if p++; p == pe {
			goto _test_eof828
		}
	st_case_828:
		switch data[p] {
		case 32:
			goto tr1162
		case 46:
			goto tr1163
		case 101:
			goto st829
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr937
		}
		goto st0
	st829:
		if p++; p == pe {
			goto _test_eof829
		}
	st_case_829:
		if data[p] == 109 {
			goto st830
		}
		goto st0
	st830:
		if p++; p == pe {
			goto _test_eof830
		}
	st_case_830:
		if data[p] == 98 {
			goto st831
		}
		goto st0
	st831:
		if p++; p == pe {
			goto _test_eof831
		}
	st_case_831:
		if data[p] == 101 {
			goto st832
		}
		goto st0
	st832:
		if p++; p == pe {
			goto _test_eof832
		}
	st_case_832:
		if data[p] == 114 {
			goto st833
		}
		goto st0
	st833:
		if p++; p == pe {
			goto _test_eof833
		}
	st_case_833:
		switch data[p] {
		case 32:
			goto tr1162
		case 46:
			goto tr1163
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr937
		}
		goto st0
	st834:
		if p++; p == pe {
			goto _test_eof834
		}
	st_case_834:
		if data[p] == 99 {
			goto st835
		}
		goto st0
	st835:
		if p++; p == pe {
			goto _test_eof835
		}
	st_case_835:
		if data[p] == 116 {
			goto st836
		}
		goto st0
	st836:
		if p++; p == pe {
			goto _test_eof836
		}
	st_case_836:
		switch data[p] {
		case 32:
			goto tr1171
		case 46:
			goto tr1172
		case 111:
			goto st837
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr947
		}
		goto st0
	st837:
		if p++; p == pe {
			goto _test_eof837
		}
	st_case_837:
		if data[p] == 98 {
			goto st838
		}
		goto st0
	st838:
		if p++; p == pe {
			goto _test_eof838
		}
	st_case_838:
		if data[p] == 101 {
			goto st839
		}
		goto st0
	st839:
		if p++; p == pe {
			goto _test_eof839
		}
	st_case_839:
		if data[p] == 114 {
			goto st840
		}
		goto st0
	st840:
		if p++; p == pe {
			goto _test_eof840
		}
	st_case_840:
		switch data[p] {
		case 32:
			goto tr1171
		case 46:
			goto tr1172
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr947
		}
		goto st0
	st841:
		if p++; p == pe {
			goto _test_eof841
		}
	st_case_841:
		switch data[p] {
		case 97:
			goto st842
		case 101:
			goto st846
		case 117:
			goto st825
		}
		goto st0
	st842:
		if p++; p == pe {
			goto _test_eof842
		}
	st_case_842:
		if data[p] == 116 {
			goto st843
		}
		goto st0
	st843:
		if p++; p == pe {
			goto _test_eof843
		}
	st_case_843:
		switch data[p] {
		case 32:
			goto st424
		case 44:
			goto st659
		case 117:
			goto st844
		}
		goto st0
	st844:
		if p++; p == pe {
			goto _test_eof844
		}
	st_case_844:
		if data[p] == 114 {
			goto st845
		}
		goto st0
	st845:
		if p++; p == pe {
			goto _test_eof845
		}
	st_case_845:
		if data[p] == 100 {
			goto st804
		}
		goto st0
	st846:
		if p++; p == pe {
			goto _test_eof846
		}
	st_case_846:
		if data[p] == 112 {
			goto st847
		}
		goto st0
	st847:
		if p++; p == pe {
			goto _test_eof847
		}
	st_case_847:
		switch data[p] {
		case 32:
			goto tr1183
		case 46:
			goto tr1184
		case 116:
			goto st848
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr956
		}
		goto st0
	st848:
		if p++; p == pe {
			goto _test_eof848
		}
	st_case_848:
		switch data[p] {
		case 32:
			goto tr1183
		case 46:
			goto tr1184
		case 101:
			goto st849
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr956
		}
		goto st0
	st849:
		if p++; p == pe {
			goto _test_eof849
		}
	st_case_849:
		if data[p] == 109 {
			goto st850
		}
		goto st0
	st850:
		if p++; p == pe {
			goto _test_eof850
		}
	st_case_850:
		if data[p] == 98 {
			goto st851
		}
		goto st0
	st851:
		if p++; p == pe {
			goto _test_eof851
		}
	st_case_851:
		if data[p] == 101 {
			goto st852
		}
		goto st0
	st852:
		if p++; p == pe {
			goto _test_eof852
		}
	st_case_852:
		if data[p] == 114 {
			goto st853
		}
		goto st0
	st853:
		if p++; p == pe {
			goto _test_eof853
		}
	st_case_853:
		switch data[p] {
		case 32:
			goto tr1183
		case 46:
			goto tr1184
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr956
		}
		goto st0
	st854:
		if p++; p == pe {
			goto _test_eof854
		}
	st_case_854:
		switch data[p] {
		case 104:
			goto st855
		case 117:
			goto st858
		}
		goto st0
	st855:
		if p++; p == pe {
			goto _test_eof855
		}
	st_case_855:
		if data[p] == 117 {
			goto st856
		}
		goto st0
	st856:
		if p++; p == pe {
			goto _test_eof856
		}
	st_case_856:
		switch data[p] {
		case 32:
			goto st424
		case 44:
			goto st659
		case 114:
			goto st857
		}
		goto st0
	st857:
		if p++; p == pe {
			goto _test_eof857
		}
	st_case_857:
		if data[p] == 115 {
			goto st845
		}
		goto st0
	st858:
		if p++; p == pe {
			goto _test_eof858
		}
	st_case_858:
		if data[p] == 101 {
			goto st859
		}
		goto st0
	st859:
		if p++; p == pe {
			goto _test_eof859
		}
	st_case_859:
		switch data[p] {
		case 32:
			goto st424
		case 44:
			goto st659
		case 115:
			goto st845
		}
		goto st0
	st860:
		if p++; p == pe {
			goto _test_eof860
		}
	st_case_860:
		if data[p] == 101 {
			goto st861
		}
		goto st0
	st861:
		if p++; p == pe {
			goto _test_eof861
		}
	st_case_861:
		if data[p] == 100 {
			goto st862
		}
		goto st0
	st862:
		if p++; p == pe {
			goto _test_eof862
		}
	st_case_862:
		switch data[p] {
		case 32:
			goto st424
		case 44:
			goto st659
		case 110:
			goto st863
		}
		goto st0
	st863:
		if p++; p == pe {
			goto _test_eof863
		}
	st_case_863:
		if data[p] == 101 {
			goto st857
		}
		goto st0
	st864:
		if p++; p == pe {
			goto _test_eof864
		}
	st_case_864:
		if data[p] == 101 {
			goto st415
		}
		goto st0
	st865:
		if p++; p == pe {
			goto _test_eof865
		}
	st_case_865:
		if data[p] == 97 {
			goto st820
		}
		goto st0
	st866:
		if p++; p == pe {
			goto _test_eof866
		}
	st_case_866:
		if data[p] == 101 {
			goto st846
		}
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof867: cs = 867; goto _test_eof
	_test_eof868: cs = 868; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof869: cs = 869; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof870: cs = 870; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof871: cs = 871; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof872: cs = 872; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof873: cs = 873; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof874: cs = 874; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof875: cs = 875; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof876: cs = 876; goto _test_eof
	_test_eof877: cs = 877; goto _test_eof
	_test_eof878: cs = 878; goto _test_eof
	_test_eof879: cs = 879; goto _test_eof
	_test_eof880: cs = 880; goto _test_eof
	_test_eof881: cs = 881; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof882: cs = 882; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof883: cs = 883; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof884: cs = 884; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof885: cs = 885; goto _test_eof
	_test_eof886: cs = 886; goto _test_eof
	_test_eof887: cs = 887; goto _test_eof
	_test_eof888: cs = 888; goto _test_eof
	_test_eof889: cs = 889; goto _test_eof
	_test_eof890: cs = 890; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof891: cs = 891; goto _test_eof
	_test_eof892: cs = 892; goto _test_eof
	_test_eof893: cs = 893; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof894: cs = 894; goto _test_eof
	_test_eof895: cs = 895; goto _test_eof
	_test_eof896: cs = 896; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof897: cs = 897; goto _test_eof
	_test_eof898: cs = 898; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof899: cs = 899; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof900: cs = 900; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof901: cs = 901; goto _test_eof
	_test_eof902: cs = 902; goto _test_eof
	_test_eof903: cs = 903; goto _test_eof
	_test_eof904: cs = 904; goto _test_eof
	_test_eof905: cs = 905; goto _test_eof
	_test_eof906: cs = 906; goto _test_eof
	_test_eof907: cs = 907; goto _test_eof
	_test_eof908: cs = 908; goto _test_eof
	_test_eof909: cs = 909; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof910: cs = 910; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof911: cs = 911; goto _test_eof
	_test_eof912: cs = 912; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof913: cs = 913; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof914: cs = 914; goto _test_eof
	_test_eof915: cs = 915; goto _test_eof
	_test_eof916: cs = 916; goto _test_eof
	_test_eof917: cs = 917; goto _test_eof
	_test_eof918: cs = 918; goto _test_eof
	_test_eof919: cs = 919; goto _test_eof
	_test_eof920: cs = 920; goto _test_eof
	_test_eof921: cs = 921; goto _test_eof
	_test_eof922: cs = 922; goto _test_eof
	_test_eof923: cs = 923; goto _test_eof
	_test_eof924: cs = 924; goto _test_eof
	_test_eof925: cs = 925; goto _test_eof
	_test_eof926: cs = 926; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof927: cs = 927; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof928: cs = 928; goto _test_eof
	_test_eof929: cs = 929; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof930: cs = 930; goto _test_eof
	_test_eof931: cs = 931; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof932: cs = 932; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof933: cs = 933; goto _test_eof
	_test_eof934: cs = 934; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof935: cs = 935; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof150: cs = 150; goto _test_eof
	_test_eof151: cs = 151; goto _test_eof
	_test_eof152: cs = 152; goto _test_eof
	_test_eof936: cs = 936; goto _test_eof
	_test_eof153: cs = 153; goto _test_eof
	_test_eof154: cs = 154; goto _test_eof
	_test_eof155: cs = 155; goto _test_eof
	_test_eof156: cs = 156; goto _test_eof
	_test_eof157: cs = 157; goto _test_eof
	_test_eof158: cs = 158; goto _test_eof
	_test_eof937: cs = 937; goto _test_eof
	_test_eof159: cs = 159; goto _test_eof
	_test_eof160: cs = 160; goto _test_eof
	_test_eof161: cs = 161; goto _test_eof
	_test_eof162: cs = 162; goto _test_eof
	_test_eof163: cs = 163; goto _test_eof
	_test_eof164: cs = 164; goto _test_eof
	_test_eof165: cs = 165; goto _test_eof
	_test_eof166: cs = 166; goto _test_eof
	_test_eof167: cs = 167; goto _test_eof
	_test_eof168: cs = 168; goto _test_eof
	_test_eof169: cs = 169; goto _test_eof
	_test_eof170: cs = 170; goto _test_eof
	_test_eof171: cs = 171; goto _test_eof
	_test_eof172: cs = 172; goto _test_eof
	_test_eof173: cs = 173; goto _test_eof
	_test_eof174: cs = 174; goto _test_eof
	_test_eof175: cs = 175; goto _test_eof
	_test_eof176: cs = 176; goto _test_eof
	_test_eof177: cs = 177; goto _test_eof
	_test_eof178: cs = 178; goto _test_eof
	_test_eof179: cs = 179; goto _test_eof
	_test_eof180: cs = 180; goto _test_eof
	_test_eof181: cs = 181; goto _test_eof
	_test_eof182: cs = 182; goto _test_eof
	_test_eof183: cs = 183; goto _test_eof
	_test_eof184: cs = 184; goto _test_eof
	_test_eof185: cs = 185; goto _test_eof
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
	_test_eof189: cs = 189; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof193: cs = 193; goto _test_eof
	_test_eof194: cs = 194; goto _test_eof
	_test_eof195: cs = 195; goto _test_eof
	_test_eof196: cs = 196; goto _test_eof
	_test_eof197: cs = 197; goto _test_eof
	_test_eof198: cs = 198; goto _test_eof
	_test_eof199: cs = 199; goto _test_eof
	_test_eof200: cs = 200; goto _test_eof
	_test_eof201: cs = 201; goto _test_eof
	_test_eof202: cs = 202; goto _test_eof
	_test_eof203: cs = 203; goto _test_eof
	_test_eof204: cs = 204; goto _test_eof
	_test_eof205: cs = 205; goto _test_eof
	_test_eof206: cs = 206; goto _test_eof
	_test_eof207: cs = 207; goto _test_eof
	_test_eof208: cs = 208; goto _test_eof
	_test_eof209: cs = 209; goto _test_eof
	_test_eof210: cs = 210; goto _test_eof
	_test_eof211: cs = 211; goto _test_eof
	_test_eof212: cs = 212; goto _test_eof
	_test_eof213: cs = 213; goto _test_eof
	_test_eof214: cs = 214; goto _test_eof
	_test_eof215: cs = 215; goto _test_eof
	_test_eof216: cs = 216; goto _test_eof
	_test_eof217: cs = 217; goto _test_eof
	_test_eof218: cs = 218; goto _test_eof
	_test_eof219: cs = 219; goto _test_eof
	_test_eof220: cs = 220; goto _test_eof
	_test_eof221: cs = 221; goto _test_eof
	_test_eof222: cs = 222; goto _test_eof
	_test_eof223: cs = 223; goto _test_eof
	_test_eof224: cs = 224; goto _test_eof
	_test_eof225: cs = 225; goto _test_eof
	_test_eof226: cs = 226; goto _test_eof
	_test_eof227: cs = 227; goto _test_eof
	_test_eof228: cs = 228; goto _test_eof
	_test_eof229: cs = 229; goto _test_eof
	_test_eof230: cs = 230; goto _test_eof
	_test_eof231: cs = 231; goto _test_eof
	_test_eof232: cs = 232; goto _test_eof
	_test_eof233: cs = 233; goto _test_eof
	_test_eof234: cs = 234; goto _test_eof
	_test_eof235: cs = 235; goto _test_eof
	_test_eof236: cs = 236; goto _test_eof
	_test_eof237: cs = 237; goto _test_eof
	_test_eof238: cs = 238; goto _test_eof
	_test_eof239: cs = 239; goto _test_eof
	_test_eof240: cs = 240; goto _test_eof
	_test_eof241: cs = 241; goto _test_eof
	_test_eof242: cs = 242; goto _test_eof
	_test_eof243: cs = 243; goto _test_eof
	_test_eof244: cs = 244; goto _test_eof
	_test_eof245: cs = 245; goto _test_eof
	_test_eof246: cs = 246; goto _test_eof
	_test_eof247: cs = 247; goto _test_eof
	_test_eof248: cs = 248; goto _test_eof
	_test_eof249: cs = 249; goto _test_eof
	_test_eof250: cs = 250; goto _test_eof
	_test_eof251: cs = 251; goto _test_eof
	_test_eof252: cs = 252; goto _test_eof
	_test_eof253: cs = 253; goto _test_eof
	_test_eof254: cs = 254; goto _test_eof
	_test_eof255: cs = 255; goto _test_eof
	_test_eof256: cs = 256; goto _test_eof
	_test_eof257: cs = 257; goto _test_eof
	_test_eof258: cs = 258; goto _test_eof
	_test_eof259: cs = 259; goto _test_eof
	_test_eof260: cs = 260; goto _test_eof
	_test_eof261: cs = 261; goto _test_eof
	_test_eof262: cs = 262; goto _test_eof
	_test_eof263: cs = 263; goto _test_eof
	_test_eof264: cs = 264; goto _test_eof
	_test_eof265: cs = 265; goto _test_eof
	_test_eof266: cs = 266; goto _test_eof
	_test_eof267: cs = 267; goto _test_eof
	_test_eof268: cs = 268; goto _test_eof
	_test_eof269: cs = 269; goto _test_eof
	_test_eof270: cs = 270; goto _test_eof
	_test_eof271: cs = 271; goto _test_eof
	_test_eof272: cs = 272; goto _test_eof
	_test_eof273: cs = 273; goto _test_eof
	_test_eof274: cs = 274; goto _test_eof
	_test_eof275: cs = 275; goto _test_eof
	_test_eof276: cs = 276; goto _test_eof
	_test_eof277: cs = 277; goto _test_eof
	_test_eof278: cs = 278; goto _test_eof
	_test_eof279: cs = 279; goto _test_eof
	_test_eof280: cs = 280; goto _test_eof
	_test_eof281: cs = 281; goto _test_eof
	_test_eof282: cs = 282; goto _test_eof
	_test_eof283: cs = 283; goto _test_eof
	_test_eof284: cs = 284; goto _test_eof
	_test_eof285: cs = 285; goto _test_eof
	_test_eof286: cs = 286; goto _test_eof
	_test_eof287: cs = 287; goto _test_eof
	_test_eof288: cs = 288; goto _test_eof
	_test_eof289: cs = 289; goto _test_eof
	_test_eof290: cs = 290; goto _test_eof
	_test_eof291: cs = 291; goto _test_eof
	_test_eof292: cs = 292; goto _test_eof
	_test_eof293: cs = 293; goto _test_eof
	_test_eof294: cs = 294; goto _test_eof
	_test_eof295: cs = 295; goto _test_eof
	_test_eof296: cs = 296; goto _test_eof
	_test_eof297: cs = 297; goto _test_eof
	_test_eof298: cs = 298; goto _test_eof
	_test_eof299: cs = 299; goto _test_eof
	_test_eof300: cs = 300; goto _test_eof
	_test_eof301: cs = 301; goto _test_eof
	_test_eof302: cs = 302; goto _test_eof
	_test_eof303: cs = 303; goto _test_eof
	_test_eof304: cs = 304; goto _test_eof
	_test_eof305: cs = 305; goto _test_eof
	_test_eof306: cs = 306; goto _test_eof
	_test_eof307: cs = 307; goto _test_eof
	_test_eof308: cs = 308; goto _test_eof
	_test_eof309: cs = 309; goto _test_eof
	_test_eof938: cs = 938; goto _test_eof
	_test_eof310: cs = 310; goto _test_eof
	_test_eof311: cs = 311; goto _test_eof
	_test_eof939: cs = 939; goto _test_eof
	_test_eof312: cs = 312; goto _test_eof
	_test_eof313: cs = 313; goto _test_eof
	_test_eof314: cs = 314; goto _test_eof
	_test_eof940: cs = 940; goto _test_eof
	_test_eof315: cs = 315; goto _test_eof
	_test_eof941: cs = 941; goto _test_eof
	_test_eof316: cs = 316; goto _test_eof
	_test_eof942: cs = 942; goto _test_eof
	_test_eof943: cs = 943; goto _test_eof
	_test_eof944: cs = 944; goto _test_eof
	_test_eof945: cs = 945; goto _test_eof
	_test_eof946: cs = 946; goto _test_eof
	_test_eof947: cs = 947; goto _test_eof
	_test_eof317: cs = 317; goto _test_eof
	_test_eof318: cs = 318; goto _test_eof
	_test_eof319: cs = 319; goto _test_eof
	_test_eof948: cs = 948; goto _test_eof
	_test_eof320: cs = 320; goto _test_eof
	_test_eof949: cs = 949; goto _test_eof
	_test_eof950: cs = 950; goto _test_eof
	_test_eof951: cs = 951; goto _test_eof
	_test_eof952: cs = 952; goto _test_eof
	_test_eof953: cs = 953; goto _test_eof
	_test_eof954: cs = 954; goto _test_eof
	_test_eof321: cs = 321; goto _test_eof
	_test_eof322: cs = 322; goto _test_eof
	_test_eof323: cs = 323; goto _test_eof
	_test_eof955: cs = 955; goto _test_eof
	_test_eof956: cs = 956; goto _test_eof
	_test_eof324: cs = 324; goto _test_eof
	_test_eof325: cs = 325; goto _test_eof
	_test_eof326: cs = 326; goto _test_eof
	_test_eof327: cs = 327; goto _test_eof
	_test_eof328: cs = 328; goto _test_eof
	_test_eof329: cs = 329; goto _test_eof
	_test_eof330: cs = 330; goto _test_eof
	_test_eof957: cs = 957; goto _test_eof
	_test_eof958: cs = 958; goto _test_eof
	_test_eof959: cs = 959; goto _test_eof
	_test_eof960: cs = 960; goto _test_eof
	_test_eof331: cs = 331; goto _test_eof
	_test_eof961: cs = 961; goto _test_eof
	_test_eof962: cs = 962; goto _test_eof
	_test_eof332: cs = 332; goto _test_eof
	_test_eof963: cs = 963; goto _test_eof
	_test_eof964: cs = 964; goto _test_eof
	_test_eof333: cs = 333; goto _test_eof
	_test_eof334: cs = 334; goto _test_eof
	_test_eof335: cs = 335; goto _test_eof
	_test_eof336: cs = 336; goto _test_eof
	_test_eof965: cs = 965; goto _test_eof
	_test_eof966: cs = 966; goto _test_eof
	_test_eof337: cs = 337; goto _test_eof
	_test_eof967: cs = 967; goto _test_eof
	_test_eof968: cs = 968; goto _test_eof
	_test_eof338: cs = 338; goto _test_eof
	_test_eof969: cs = 969; goto _test_eof
	_test_eof339: cs = 339; goto _test_eof
	_test_eof970: cs = 970; goto _test_eof
	_test_eof340: cs = 340; goto _test_eof
	_test_eof971: cs = 971; goto _test_eof
	_test_eof972: cs = 972; goto _test_eof
	_test_eof973: cs = 973; goto _test_eof
	_test_eof974: cs = 974; goto _test_eof
	_test_eof975: cs = 975; goto _test_eof
	_test_eof976: cs = 976; goto _test_eof
	_test_eof977: cs = 977; goto _test_eof
	_test_eof978: cs = 978; goto _test_eof
	_test_eof979: cs = 979; goto _test_eof
	_test_eof341: cs = 341; goto _test_eof
	_test_eof980: cs = 980; goto _test_eof
	_test_eof342: cs = 342; goto _test_eof
	_test_eof981: cs = 981; goto _test_eof
	_test_eof982: cs = 982; goto _test_eof
	_test_eof343: cs = 343; goto _test_eof
	_test_eof983: cs = 983; goto _test_eof
	_test_eof344: cs = 344; goto _test_eof
	_test_eof984: cs = 984; goto _test_eof
	_test_eof985: cs = 985; goto _test_eof
	_test_eof986: cs = 986; goto _test_eof
	_test_eof987: cs = 987; goto _test_eof
	_test_eof988: cs = 988; goto _test_eof
	_test_eof989: cs = 989; goto _test_eof
	_test_eof990: cs = 990; goto _test_eof
	_test_eof991: cs = 991; goto _test_eof
	_test_eof992: cs = 992; goto _test_eof
	_test_eof993: cs = 993; goto _test_eof
	_test_eof994: cs = 994; goto _test_eof
	_test_eof995: cs = 995; goto _test_eof
	_test_eof996: cs = 996; goto _test_eof
	_test_eof345: cs = 345; goto _test_eof
	_test_eof997: cs = 997; goto _test_eof
	_test_eof346: cs = 346; goto _test_eof
	_test_eof347: cs = 347; goto _test_eof
	_test_eof348: cs = 348; goto _test_eof
	_test_eof349: cs = 349; goto _test_eof
	_test_eof350: cs = 350; goto _test_eof
	_test_eof351: cs = 351; goto _test_eof
	_test_eof352: cs = 352; goto _test_eof
	_test_eof353: cs = 353; goto _test_eof
	_test_eof354: cs = 354; goto _test_eof
	_test_eof355: cs = 355; goto _test_eof
	_test_eof356: cs = 356; goto _test_eof
	_test_eof357: cs = 357; goto _test_eof
	_test_eof358: cs = 358; goto _test_eof
	_test_eof359: cs = 359; goto _test_eof
	_test_eof360: cs = 360; goto _test_eof
	_test_eof361: cs = 361; goto _test_eof
	_test_eof362: cs = 362; goto _test_eof
	_test_eof363: cs = 363; goto _test_eof
	_test_eof364: cs = 364; goto _test_eof
	_test_eof365: cs = 365; goto _test_eof
	_test_eof366: cs = 366; goto _test_eof
	_test_eof367: cs = 367; goto _test_eof
	_test_eof368: cs = 368; goto _test_eof
	_test_eof369: cs = 369; goto _test_eof
	_test_eof370: cs = 370; goto _test_eof
	_test_eof371: cs = 371; goto _test_eof
	_test_eof372: cs = 372; goto _test_eof
	_test_eof373: cs = 373; goto _test_eof
	_test_eof374: cs = 374; goto _test_eof
	_test_eof375: cs = 375; goto _test_eof
	_test_eof376: cs = 376; goto _test_eof
	_test_eof377: cs = 377; goto _test_eof
	_test_eof378: cs = 378; goto _test_eof
	_test_eof379: cs = 379; goto _test_eof
	_test_eof380: cs = 380; goto _test_eof
	_test_eof381: cs = 381; goto _test_eof
	_test_eof382: cs = 382; goto _test_eof
	_test_eof383: cs = 383; goto _test_eof
	_test_eof384: cs = 384; goto _test_eof
	_test_eof998: cs = 998; goto _test_eof
	_test_eof385: cs = 385; goto _test_eof
	_test_eof999: cs = 999; goto _test_eof
	_test_eof386: cs = 386; goto _test_eof
	_test_eof387: cs = 387; goto _test_eof
	_test_eof388: cs = 388; goto _test_eof
	_test_eof389: cs = 389; goto _test_eof
	_test_eof390: cs = 390; goto _test_eof
	_test_eof1000: cs = 1000; goto _test_eof
	_test_eof1001: cs = 1001; goto _test_eof
	_test_eof391: cs = 391; goto _test_eof
	_test_eof392: cs = 392; goto _test_eof
	_test_eof393: cs = 393; goto _test_eof
	_test_eof394: cs = 394; goto _test_eof
	_test_eof395: cs = 395; goto _test_eof
	_test_eof396: cs = 396; goto _test_eof
	_test_eof397: cs = 397; goto _test_eof
	_test_eof398: cs = 398; goto _test_eof
	_test_eof399: cs = 399; goto _test_eof
	_test_eof400: cs = 400; goto _test_eof
	_test_eof401: cs = 401; goto _test_eof
	_test_eof402: cs = 402; goto _test_eof
	_test_eof403: cs = 403; goto _test_eof
	_test_eof404: cs = 404; goto _test_eof
	_test_eof405: cs = 405; goto _test_eof
	_test_eof406: cs = 406; goto _test_eof
	_test_eof407: cs = 407; goto _test_eof
	_test_eof408: cs = 408; goto _test_eof
	_test_eof409: cs = 409; goto _test_eof
	_test_eof410: cs = 410; goto _test_eof
	_test_eof411: cs = 411; goto _test_eof
	_test_eof412: cs = 412; goto _test_eof
	_test_eof413: cs = 413; goto _test_eof
	_test_eof414: cs = 414; goto _test_eof
	_test_eof415: cs = 415; goto _test_eof
	_test_eof416: cs = 416; goto _test_eof
	_test_eof417: cs = 417; goto _test_eof
	_test_eof418: cs = 418; goto _test_eof
	_test_eof419: cs = 419; goto _test_eof
	_test_eof420: cs = 420; goto _test_eof
	_test_eof421: cs = 421; goto _test_eof
	_test_eof422: cs = 422; goto _test_eof
	_test_eof423: cs = 423; goto _test_eof
	_test_eof424: cs = 424; goto _test_eof
	_test_eof425: cs = 425; goto _test_eof
	_test_eof426: cs = 426; goto _test_eof
	_test_eof427: cs = 427; goto _test_eof
	_test_eof428: cs = 428; goto _test_eof
	_test_eof429: cs = 429; goto _test_eof
	_test_eof430: cs = 430; goto _test_eof
	_test_eof431: cs = 431; goto _test_eof
	_test_eof432: cs = 432; goto _test_eof
	_test_eof433: cs = 433; goto _test_eof
	_test_eof434: cs = 434; goto _test_eof
	_test_eof435: cs = 435; goto _test_eof
	_test_eof436: cs = 436; goto _test_eof
	_test_eof437: cs = 437; goto _test_eof
	_test_eof438: cs = 438; goto _test_eof
	_test_eof439: cs = 439; goto _test_eof
	_test_eof440: cs = 440; goto _test_eof
	_test_eof441: cs = 441; goto _test_eof
	_test_eof442: cs = 442; goto _test_eof
	_test_eof1002: cs = 1002; goto _test_eof
	_test_eof443: cs = 443; goto _test_eof
	_test_eof444: cs = 444; goto _test_eof
	_test_eof445: cs = 445; goto _test_eof
	_test_eof446: cs = 446; goto _test_eof
	_test_eof447: cs = 447; goto _test_eof
	_test_eof448: cs = 448; goto _test_eof
	_test_eof449: cs = 449; goto _test_eof
	_test_eof450: cs = 450; goto _test_eof
	_test_eof451: cs = 451; goto _test_eof
	_test_eof452: cs = 452; goto _test_eof
	_test_eof453: cs = 453; goto _test_eof
	_test_eof454: cs = 454; goto _test_eof
	_test_eof455: cs = 455; goto _test_eof
	_test_eof456: cs = 456; goto _test_eof
	_test_eof457: cs = 457; goto _test_eof
	_test_eof458: cs = 458; goto _test_eof
	_test_eof459: cs = 459; goto _test_eof
	_test_eof460: cs = 460; goto _test_eof
	_test_eof461: cs = 461; goto _test_eof
	_test_eof462: cs = 462; goto _test_eof
	_test_eof463: cs = 463; goto _test_eof
	_test_eof464: cs = 464; goto _test_eof
	_test_eof465: cs = 465; goto _test_eof
	_test_eof466: cs = 466; goto _test_eof
	_test_eof467: cs = 467; goto _test_eof
	_test_eof468: cs = 468; goto _test_eof
	_test_eof469: cs = 469; goto _test_eof
	_test_eof470: cs = 470; goto _test_eof
	_test_eof471: cs = 471; goto _test_eof
	_test_eof472: cs = 472; goto _test_eof
	_test_eof473: cs = 473; goto _test_eof
	_test_eof474: cs = 474; goto _test_eof
	_test_eof475: cs = 475; goto _test_eof
	_test_eof476: cs = 476; goto _test_eof
	_test_eof477: cs = 477; goto _test_eof
	_test_eof478: cs = 478; goto _test_eof
	_test_eof479: cs = 479; goto _test_eof
	_test_eof480: cs = 480; goto _test_eof
	_test_eof481: cs = 481; goto _test_eof
	_test_eof482: cs = 482; goto _test_eof
	_test_eof483: cs = 483; goto _test_eof
	_test_eof484: cs = 484; goto _test_eof
	_test_eof485: cs = 485; goto _test_eof
	_test_eof486: cs = 486; goto _test_eof
	_test_eof487: cs = 487; goto _test_eof
	_test_eof488: cs = 488; goto _test_eof
	_test_eof489: cs = 489; goto _test_eof
	_test_eof490: cs = 490; goto _test_eof
	_test_eof491: cs = 491; goto _test_eof
	_test_eof492: cs = 492; goto _test_eof
	_test_eof493: cs = 493; goto _test_eof
	_test_eof494: cs = 494; goto _test_eof
	_test_eof495: cs = 495; goto _test_eof
	_test_eof496: cs = 496; goto _test_eof
	_test_eof497: cs = 497; goto _test_eof
	_test_eof498: cs = 498; goto _test_eof
	_test_eof499: cs = 499; goto _test_eof
	_test_eof500: cs = 500; goto _test_eof
	_test_eof501: cs = 501; goto _test_eof
	_test_eof502: cs = 502; goto _test_eof
	_test_eof503: cs = 503; goto _test_eof
	_test_eof504: cs = 504; goto _test_eof
	_test_eof505: cs = 505; goto _test_eof
	_test_eof506: cs = 506; goto _test_eof
	_test_eof507: cs = 507; goto _test_eof
	_test_eof508: cs = 508; goto _test_eof
	_test_eof509: cs = 509; goto _test_eof
	_test_eof510: cs = 510; goto _test_eof
	_test_eof511: cs = 511; goto _test_eof
	_test_eof512: cs = 512; goto _test_eof
	_test_eof513: cs = 513; goto _test_eof
	_test_eof514: cs = 514; goto _test_eof
	_test_eof515: cs = 515; goto _test_eof
	_test_eof516: cs = 516; goto _test_eof
	_test_eof517: cs = 517; goto _test_eof
	_test_eof518: cs = 518; goto _test_eof
	_test_eof519: cs = 519; goto _test_eof
	_test_eof520: cs = 520; goto _test_eof
	_test_eof521: cs = 521; goto _test_eof
	_test_eof522: cs = 522; goto _test_eof
	_test_eof523: cs = 523; goto _test_eof
	_test_eof524: cs = 524; goto _test_eof
	_test_eof525: cs = 525; goto _test_eof
	_test_eof1003: cs = 1003; goto _test_eof
	_test_eof1004: cs = 1004; goto _test_eof
	_test_eof526: cs = 526; goto _test_eof
	_test_eof527: cs = 527; goto _test_eof
	_test_eof1005: cs = 1005; goto _test_eof
	_test_eof528: cs = 528; goto _test_eof
	_test_eof529: cs = 529; goto _test_eof
	_test_eof530: cs = 530; goto _test_eof
	_test_eof1006: cs = 1006; goto _test_eof
	_test_eof531: cs = 531; goto _test_eof
	_test_eof1007: cs = 1007; goto _test_eof
	_test_eof532: cs = 532; goto _test_eof
	_test_eof1008: cs = 1008; goto _test_eof
	_test_eof1009: cs = 1009; goto _test_eof
	_test_eof1010: cs = 1010; goto _test_eof
	_test_eof1011: cs = 1011; goto _test_eof
	_test_eof1012: cs = 1012; goto _test_eof
	_test_eof1013: cs = 1013; goto _test_eof
	_test_eof533: cs = 533; goto _test_eof
	_test_eof534: cs = 534; goto _test_eof
	_test_eof535: cs = 535; goto _test_eof
	_test_eof1014: cs = 1014; goto _test_eof
	_test_eof536: cs = 536; goto _test_eof
	_test_eof1015: cs = 1015; goto _test_eof
	_test_eof537: cs = 537; goto _test_eof
	_test_eof1016: cs = 1016; goto _test_eof
	_test_eof538: cs = 538; goto _test_eof
	_test_eof1017: cs = 1017; goto _test_eof
	_test_eof1018: cs = 1018; goto _test_eof
	_test_eof1019: cs = 1019; goto _test_eof
	_test_eof1020: cs = 1020; goto _test_eof
	_test_eof1021: cs = 1021; goto _test_eof
	_test_eof1022: cs = 1022; goto _test_eof
	_test_eof539: cs = 539; goto _test_eof
	_test_eof540: cs = 540; goto _test_eof
	_test_eof541: cs = 541; goto _test_eof
	_test_eof1023: cs = 1023; goto _test_eof
	_test_eof542: cs = 542; goto _test_eof
	_test_eof1024: cs = 1024; goto _test_eof
	_test_eof543: cs = 543; goto _test_eof
	_test_eof544: cs = 544; goto _test_eof
	_test_eof1025: cs = 1025; goto _test_eof
	_test_eof1026: cs = 1026; goto _test_eof
	_test_eof545: cs = 545; goto _test_eof
	_test_eof546: cs = 546; goto _test_eof
	_test_eof547: cs = 547; goto _test_eof
	_test_eof548: cs = 548; goto _test_eof
	_test_eof549: cs = 549; goto _test_eof
	_test_eof550: cs = 550; goto _test_eof
	_test_eof551: cs = 551; goto _test_eof
	_test_eof552: cs = 552; goto _test_eof
	_test_eof553: cs = 553; goto _test_eof
	_test_eof554: cs = 554; goto _test_eof
	_test_eof555: cs = 555; goto _test_eof
	_test_eof556: cs = 556; goto _test_eof
	_test_eof557: cs = 557; goto _test_eof
	_test_eof558: cs = 558; goto _test_eof
	_test_eof559: cs = 559; goto _test_eof
	_test_eof1027: cs = 1027; goto _test_eof
	_test_eof1028: cs = 1028; goto _test_eof
	_test_eof560: cs = 560; goto _test_eof
	_test_eof561: cs = 561; goto _test_eof
	_test_eof562: cs = 562; goto _test_eof
	_test_eof563: cs = 563; goto _test_eof
	_test_eof564: cs = 564; goto _test_eof
	_test_eof565: cs = 565; goto _test_eof
	_test_eof566: cs = 566; goto _test_eof
	_test_eof567: cs = 567; goto _test_eof
	_test_eof568: cs = 568; goto _test_eof
	_test_eof569: cs = 569; goto _test_eof
	_test_eof570: cs = 570; goto _test_eof
	_test_eof571: cs = 571; goto _test_eof
	_test_eof572: cs = 572; goto _test_eof
	_test_eof573: cs = 573; goto _test_eof
	_test_eof574: cs = 574; goto _test_eof
	_test_eof575: cs = 575; goto _test_eof
	_test_eof576: cs = 576; goto _test_eof
	_test_eof577: cs = 577; goto _test_eof
	_test_eof578: cs = 578; goto _test_eof
	_test_eof579: cs = 579; goto _test_eof
	_test_eof580: cs = 580; goto _test_eof
	_test_eof581: cs = 581; goto _test_eof
	_test_eof582: cs = 582; goto _test_eof
	_test_eof583: cs = 583; goto _test_eof
	_test_eof584: cs = 584; goto _test_eof
	_test_eof585: cs = 585; goto _test_eof
	_test_eof586: cs = 586; goto _test_eof
	_test_eof1029: cs = 1029; goto _test_eof
	_test_eof587: cs = 587; goto _test_eof
	_test_eof588: cs = 588; goto _test_eof
	_test_eof589: cs = 589; goto _test_eof
	_test_eof590: cs = 590; goto _test_eof
	_test_eof591: cs = 591; goto _test_eof
	_test_eof592: cs = 592; goto _test_eof
	_test_eof593: cs = 593; goto _test_eof
	_test_eof594: cs = 594; goto _test_eof
	_test_eof595: cs = 595; goto _test_eof
	_test_eof596: cs = 596; goto _test_eof
	_test_eof597: cs = 597; goto _test_eof
	_test_eof598: cs = 598; goto _test_eof
	_test_eof599: cs = 599; goto _test_eof
	_test_eof600: cs = 600; goto _test_eof
	_test_eof601: cs = 601; goto _test_eof
	_test_eof602: cs = 602; goto _test_eof
	_test_eof603: cs = 603; goto _test_eof
	_test_eof604: cs = 604; goto _test_eof
	_test_eof605: cs = 605; goto _test_eof
	_test_eof606: cs = 606; goto _test_eof
	_test_eof607: cs = 607; goto _test_eof
	_test_eof608: cs = 608; goto _test_eof
	_test_eof609: cs = 609; goto _test_eof
	_test_eof610: cs = 610; goto _test_eof
	_test_eof611: cs = 611; goto _test_eof
	_test_eof612: cs = 612; goto _test_eof
	_test_eof613: cs = 613; goto _test_eof
	_test_eof614: cs = 614; goto _test_eof
	_test_eof615: cs = 615; goto _test_eof
	_test_eof616: cs = 616; goto _test_eof
	_test_eof617: cs = 617; goto _test_eof
	_test_eof618: cs = 618; goto _test_eof
	_test_eof619: cs = 619; goto _test_eof
	_test_eof620: cs = 620; goto _test_eof
	_test_eof621: cs = 621; goto _test_eof
	_test_eof622: cs = 622; goto _test_eof
	_test_eof623: cs = 623; goto _test_eof
	_test_eof624: cs = 624; goto _test_eof
	_test_eof625: cs = 625; goto _test_eof
	_test_eof626: cs = 626; goto _test_eof
	_test_eof627: cs = 627; goto _test_eof
	_test_eof628: cs = 628; goto _test_eof
	_test_eof629: cs = 629; goto _test_eof
	_test_eof630: cs = 630; goto _test_eof
	_test_eof631: cs = 631; goto _test_eof
	_test_eof632: cs = 632; goto _test_eof
	_test_eof633: cs = 633; goto _test_eof
	_test_eof634: cs = 634; goto _test_eof
	_test_eof635: cs = 635; goto _test_eof
	_test_eof636: cs = 636; goto _test_eof
	_test_eof637: cs = 637; goto _test_eof
	_test_eof638: cs = 638; goto _test_eof
	_test_eof639: cs = 639; goto _test_eof
	_test_eof640: cs = 640; goto _test_eof
	_test_eof641: cs = 641; goto _test_eof
	_test_eof642: cs = 642; goto _test_eof
	_test_eof643: cs = 643; goto _test_eof
	_test_eof644: cs = 644; goto _test_eof
	_test_eof645: cs = 645; goto _test_eof
	_test_eof646: cs = 646; goto _test_eof
	_test_eof647: cs = 647; goto _test_eof
	_test_eof648: cs = 648; goto _test_eof
	_test_eof649: cs = 649; goto _test_eof
	_test_eof650: cs = 650; goto _test_eof
	_test_eof651: cs = 651; goto _test_eof
	_test_eof652: cs = 652; goto _test_eof
	_test_eof653: cs = 653; goto _test_eof
	_test_eof654: cs = 654; goto _test_eof
	_test_eof655: cs = 655; goto _test_eof
	_test_eof656: cs = 656; goto _test_eof
	_test_eof657: cs = 657; goto _test_eof
	_test_eof658: cs = 658; goto _test_eof
	_test_eof659: cs = 659; goto _test_eof
	_test_eof660: cs = 660; goto _test_eof
	_test_eof661: cs = 661; goto _test_eof
	_test_eof662: cs = 662; goto _test_eof
	_test_eof663: cs = 663; goto _test_eof
	_test_eof664: cs = 664; goto _test_eof
	_test_eof665: cs = 665; goto _test_eof
	_test_eof666: cs = 666; goto _test_eof
	_test_eof667: cs = 667; goto _test_eof
	_test_eof668: cs = 668; goto _test_eof
	_test_eof669: cs = 669; goto _test_eof
	_test_eof670: cs = 670; goto _test_eof
	_test_eof671: cs = 671; goto _test_eof
	_test_eof672: cs = 672; goto _test_eof
	_test_eof673: cs = 673; goto _test_eof
	_test_eof674: cs = 674; goto _test_eof
	_test_eof675: cs = 675; goto _test_eof
	_test_eof676: cs = 676; goto _test_eof
	_test_eof677: cs = 677; goto _test_eof
	_test_eof678: cs = 678; goto _test_eof
	_test_eof679: cs = 679; goto _test_eof
	_test_eof680: cs = 680; goto _test_eof
	_test_eof681: cs = 681; goto _test_eof
	_test_eof682: cs = 682; goto _test_eof
	_test_eof683: cs = 683; goto _test_eof
	_test_eof684: cs = 684; goto _test_eof
	_test_eof685: cs = 685; goto _test_eof
	_test_eof686: cs = 686; goto _test_eof
	_test_eof687: cs = 687; goto _test_eof
	_test_eof688: cs = 688; goto _test_eof
	_test_eof689: cs = 689; goto _test_eof
	_test_eof690: cs = 690; goto _test_eof
	_test_eof691: cs = 691; goto _test_eof
	_test_eof692: cs = 692; goto _test_eof
	_test_eof693: cs = 693; goto _test_eof
	_test_eof694: cs = 694; goto _test_eof
	_test_eof695: cs = 695; goto _test_eof
	_test_eof696: cs = 696; goto _test_eof
	_test_eof697: cs = 697; goto _test_eof
	_test_eof698: cs = 698; goto _test_eof
	_test_eof699: cs = 699; goto _test_eof
	_test_eof700: cs = 700; goto _test_eof
	_test_eof701: cs = 701; goto _test_eof
	_test_eof702: cs = 702; goto _test_eof
	_test_eof703: cs = 703; goto _test_eof
	_test_eof704: cs = 704; goto _test_eof
	_test_eof705: cs = 705; goto _test_eof
	_test_eof706: cs = 706; goto _test_eof
	_test_eof707: cs = 707; goto _test_eof
	_test_eof708: cs = 708; goto _test_eof
	_test_eof709: cs = 709; goto _test_eof
	_test_eof710: cs = 710; goto _test_eof
	_test_eof711: cs = 711; goto _test_eof
	_test_eof712: cs = 712; goto _test_eof
	_test_eof713: cs = 713; goto _test_eof
	_test_eof714: cs = 714; goto _test_eof
	_test_eof715: cs = 715; goto _test_eof
	_test_eof716: cs = 716; goto _test_eof
	_test_eof717: cs = 717; goto _test_eof
	_test_eof718: cs = 718; goto _test_eof
	_test_eof719: cs = 719; goto _test_eof
	_test_eof720: cs = 720; goto _test_eof
	_test_eof721: cs = 721; goto _test_eof
	_test_eof722: cs = 722; goto _test_eof
	_test_eof723: cs = 723; goto _test_eof
	_test_eof724: cs = 724; goto _test_eof
	_test_eof725: cs = 725; goto _test_eof
	_test_eof726: cs = 726; goto _test_eof
	_test_eof727: cs = 727; goto _test_eof
	_test_eof728: cs = 728; goto _test_eof
	_test_eof729: cs = 729; goto _test_eof
	_test_eof730: cs = 730; goto _test_eof
	_test_eof731: cs = 731; goto _test_eof
	_test_eof732: cs = 732; goto _test_eof
	_test_eof733: cs = 733; goto _test_eof
	_test_eof734: cs = 734; goto _test_eof
	_test_eof735: cs = 735; goto _test_eof
	_test_eof736: cs = 736; goto _test_eof
	_test_eof737: cs = 737; goto _test_eof
	_test_eof738: cs = 738; goto _test_eof
	_test_eof739: cs = 739; goto _test_eof
	_test_eof740: cs = 740; goto _test_eof
	_test_eof741: cs = 741; goto _test_eof
	_test_eof742: cs = 742; goto _test_eof
	_test_eof743: cs = 743; goto _test_eof
	_test_eof744: cs = 744; goto _test_eof
	_test_eof745: cs = 745; goto _test_eof
	_test_eof746: cs = 746; goto _test_eof
	_test_eof747: cs = 747; goto _test_eof
	_test_eof748: cs = 748; goto _test_eof
	_test_eof749: cs = 749; goto _test_eof
	_test_eof750: cs = 750; goto _test_eof
	_test_eof751: cs = 751; goto _test_eof
	_test_eof752: cs = 752; goto _test_eof
	_test_eof753: cs = 753; goto _test_eof
	_test_eof754: cs = 754; goto _test_eof
	_test_eof755: cs = 755; goto _test_eof
	_test_eof756: cs = 756; goto _test_eof
	_test_eof757: cs = 757; goto _test_eof
	_test_eof758: cs = 758; goto _test_eof
	_test_eof759: cs = 759; goto _test_eof
	_test_eof760: cs = 760; goto _test_eof
	_test_eof761: cs = 761; goto _test_eof
	_test_eof762: cs = 762; goto _test_eof
	_test_eof763: cs = 763; goto _test_eof
	_test_eof764: cs = 764; goto _test_eof
	_test_eof765: cs = 765; goto _test_eof
	_test_eof766: cs = 766; goto _test_eof
	_test_eof767: cs = 767; goto _test_eof
	_test_eof768: cs = 768; goto _test_eof
	_test_eof769: cs = 769; goto _test_eof
	_test_eof770: cs = 770; goto _test_eof
	_test_eof771: cs = 771; goto _test_eof
	_test_eof772: cs = 772; goto _test_eof
	_test_eof773: cs = 773; goto _test_eof
	_test_eof774: cs = 774; goto _test_eof
	_test_eof775: cs = 775; goto _test_eof
	_test_eof776: cs = 776; goto _test_eof
	_test_eof777: cs = 777; goto _test_eof
	_test_eof778: cs = 778; goto _test_eof
	_test_eof779: cs = 779; goto _test_eof
	_test_eof780: cs = 780; goto _test_eof
	_test_eof781: cs = 781; goto _test_eof
	_test_eof782: cs = 782; goto _test_eof
	_test_eof783: cs = 783; goto _test_eof
	_test_eof784: cs = 784; goto _test_eof
	_test_eof785: cs = 785; goto _test_eof
	_test_eof786: cs = 786; goto _test_eof
	_test_eof787: cs = 787; goto _test_eof
	_test_eof788: cs = 788; goto _test_eof
	_test_eof789: cs = 789; goto _test_eof
	_test_eof790: cs = 790; goto _test_eof
	_test_eof791: cs = 791; goto _test_eof
	_test_eof792: cs = 792; goto _test_eof
	_test_eof793: cs = 793; goto _test_eof
	_test_eof794: cs = 794; goto _test_eof
	_test_eof795: cs = 795; goto _test_eof
	_test_eof796: cs = 796; goto _test_eof
	_test_eof797: cs = 797; goto _test_eof
	_test_eof798: cs = 798; goto _test_eof
	_test_eof799: cs = 799; goto _test_eof
	_test_eof800: cs = 800; goto _test_eof
	_test_eof801: cs = 801; goto _test_eof
	_test_eof802: cs = 802; goto _test_eof
	_test_eof803: cs = 803; goto _test_eof
	_test_eof804: cs = 804; goto _test_eof
	_test_eof805: cs = 805; goto _test_eof
	_test_eof806: cs = 806; goto _test_eof
	_test_eof807: cs = 807; goto _test_eof
	_test_eof808: cs = 808; goto _test_eof
	_test_eof809: cs = 809; goto _test_eof
	_test_eof810: cs = 810; goto _test_eof
	_test_eof811: cs = 811; goto _test_eof
	_test_eof812: cs = 812; goto _test_eof
	_test_eof813: cs = 813; goto _test_eof
	_test_eof814: cs = 814; goto _test_eof
	_test_eof815: cs = 815; goto _test_eof
	_test_eof816: cs = 816; goto _test_eof
	_test_eof817: cs = 817; goto _test_eof
	_test_eof818: cs = 818; goto _test_eof
	_test_eof819: cs = 819; goto _test_eof
	_test_eof820: cs = 820; goto _test_eof
	_test_eof821: cs = 821; goto _test_eof
	_test_eof822: cs = 822; goto _test_eof
	_test_eof823: cs = 823; goto _test_eof
	_test_eof824: cs = 824; goto _test_eof
	_test_eof825: cs = 825; goto _test_eof
	_test_eof826: cs = 826; goto _test_eof
	_test_eof827: cs = 827; goto _test_eof
	_test_eof828: cs = 828; goto _test_eof
	_test_eof829: cs = 829; goto _test_eof
	_test_eof830: cs = 830; goto _test_eof
	_test_eof831: cs = 831; goto _test_eof
	_test_eof832: cs = 832; goto _test_eof
	_test_eof833: cs = 833; goto _test_eof
	_test_eof834: cs = 834; goto _test_eof
	_test_eof835: cs = 835; goto _test_eof
	_test_eof836: cs = 836; goto _test_eof
	_test_eof837: cs = 837; goto _test_eof
	_test_eof838: cs = 838; goto _test_eof
	_test_eof839: cs = 839; goto _test_eof
	_test_eof840: cs = 840; goto _test_eof
	_test_eof841: cs = 841; goto _test_eof
	_test_eof842: cs = 842; goto _test_eof
	_test_eof843: cs = 843; goto _test_eof
	_test_eof844: cs = 844; goto _test_eof
	_test_eof845: cs = 845; goto _test_eof
	_test_eof846: cs = 846; goto _test_eof
	_test_eof847: cs = 847; goto _test_eof
	_test_eof848: cs = 848; goto _test_eof
	_test_eof849: cs = 849; goto _test_eof
	_test_eof850: cs = 850; goto _test_eof
	_test_eof851: cs = 851; goto _test_eof
	_test_eof852: cs = 852; goto _test_eof
	_test_eof853: cs = 853; goto _test_eof
	_test_eof854: cs = 854; goto _test_eof
	_test_eof855: cs = 855; goto _test_eof
	_test_eof856: cs = 856; goto _test_eof
	_test_eof857: cs = 857; goto _test_eof
	_test_eof858: cs = 858; goto _test_eof
	_test_eof859: cs = 859; goto _test_eof
	_test_eof860: cs = 860; goto _test_eof
	_test_eof861: cs = 861; goto _test_eof
	_test_eof862: cs = 862; goto _test_eof
	_test_eof863: cs = 863; goto _test_eof
	_test_eof864: cs = 864; goto _test_eof
	_test_eof865: cs = 865; goto _test_eof
	_test_eof866: cs = 866; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 874, 878, 887, 898, 929, 940, 944, 951, 956, 1006, 1010, 1019, 1026, 1028:
//line ragel/datetime.rl:7
 st.Zoned = true 
		case 934:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

		case 936, 938, 957, 998, 1000, 1002, 1003, 1029:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

		case 937:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

		case 932, 933:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

		case 900, 910, 970, 980:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

		case 872, 884, 897, 1016, 1025:
//line ragel/datetime.rl:67

    st.Ad_bc = ADBC_BC;

		case 895, 967:
//line ragel/datetime.rl:71

    if st.Hour > 12 {
        err = errors.New("hour value overflow")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

		case 867, 930, 931:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

		case 899, 960, 969:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

		case 892, 926, 927, 959, 962, 963, 965, 996, 997:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

		case 912, 982:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

		case 911, 925, 981, 995:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

		case 923, 993:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

		case 913, 924, 983, 994:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

		case 901, 902, 903, 904, 905, 906, 907, 908, 909, 914, 915, 916, 917, 918, 919, 920, 921, 922, 971, 972, 973, 974, 975, 976, 977, 978, 979, 984, 985, 986, 987, 988, 989, 990, 991, 992:
//line ragel/datetime.rl:134

    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }

		case 870, 871:
//line ragel/datetime.rl:164

    if dot_index := strings.IndexRune(data[pb:p], '.'); dot_index == -1 { // no '.'
        monotonic_offset_sec := parse_digits(data[pb:p])
        st.MonotonicOffsetNanosecond = int64(monotonic_offset_sec) * 1000000000
    }else {
        monotonic_offset_sec := parse_digits(data[pb:pb+dot_index])
        monotonic_offset_nsec := parse_digits(data[pb+dot_index+1:p])
        st.MonotonicOffsetNanosecond = int64(monotonic_offset_sec) * 1000000000 + int64(monotonic_offset_nsec)
    }

		case 961:
//line ragel/datetime.rl:52

    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }

//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

		case 879, 880, 888, 889, 945, 946, 952, 953, 1011, 1012, 1020, 1021:
//line ragel/datetime.rl:187

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
		case 873, 875, 876, 877, 881, 885, 886, 890, 939, 941, 942, 943, 947, 949, 950, 954, 1005, 1007, 1008, 1009, 1013, 1017, 1018, 1022:
//line ragel/datetime.rl:197

    st.ZoneOffsetIsValid = true

    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
		case 882, 891, 948, 955, 1014, 1023:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
//line ragel/parse_datetime.go:27629
		}
	}

	_out: {}
	}

//line ragel/parse_datetime.go.rl:35

    if p != eof {  // input date not fully consumed
        err = errors.New("input data not fully consumed");
        return
    }

    if cs < datetime_parser_first_final {
        err = fmt.Errorf("time parse error: %s", data)
    }
    return
}

