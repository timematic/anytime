
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
)


//line ragel/parse_datetime.go:19
const datetime_parser_start int = 1
const datetime_parser_first_final int = 818
const datetime_parser_error int = 0

const datetime_parser_en_main int = 1


//line ragel/parse_datetime.go.rl:23



//line ragel/parse_datetime.go:31
const start int = 1

const en_main int = 1


//line ragel/parse_datetime.go.rl:26

func Parse(data string) (st ParsedDatetime, err error) {
    cs, p, pe := 0, 0, len(data)
    eof := pe
    pb := p

    
//line ragel/parse_datetime.go:45
	{
	cs = start
	}

//line ragel/parse_datetime.go.rl:33
    
//line ragel/parse_datetime.go:52
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
	case 818:
		goto st_case_818
	case 819:
		goto st_case_819
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 820:
		goto st_case_820
	case 13:
		goto st_case_13
	case 821:
		goto st_case_821
	case 14:
		goto st_case_14
	case 822:
		goto st_case_822
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 823:
		goto st_case_823
	case 18:
		goto st_case_18
	case 824:
		goto st_case_824
	case 19:
		goto st_case_19
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
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 832:
		goto st_case_832
	case 23:
		goto st_case_23
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
	case 24:
		goto st_case_24
	case 840:
		goto st_case_840
	case 841:
		goto st_case_841
	case 25:
		goto st_case_25
	case 842:
		goto st_case_842
	case 843:
		goto st_case_843
	case 844:
		goto st_case_844
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 845:
		goto st_case_845
	case 846:
		goto st_case_846
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 847:
		goto st_case_847
	case 30:
		goto st_case_30
	case 848:
		goto st_case_848
	case 31:
		goto st_case_31
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
	case 32:
		goto st_case_32
	case 858:
		goto st_case_858
	case 33:
		goto st_case_33
	case 859:
		goto st_case_859
	case 860:
		goto st_case_860
	case 34:
		goto st_case_34
	case 861:
		goto st_case_861
	case 35:
		goto st_case_35
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
	case 867:
		goto st_case_867
	case 868:
		goto st_case_868
	case 869:
		goto st_case_869
	case 870:
		goto st_case_870
	case 871:
		goto st_case_871
	case 872:
		goto st_case_872
	case 873:
		goto st_case_873
	case 874:
		goto st_case_874
	case 36:
		goto st_case_36
	case 875:
		goto st_case_875
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 876:
		goto st_case_876
	case 877:
		goto st_case_877
	case 39:
		goto st_case_39
	case 878:
		goto st_case_878
	case 879:
		goto st_case_879
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
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
	case 52:
		goto st_case_52
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
	case 880:
		goto st_case_880
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
	case 881:
		goto st_case_881
	case 882:
		goto st_case_882
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
	case 883:
		goto st_case_883
	case 125:
		goto st_case_125
	case 884:
		goto st_case_884
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
	case 885:
		goto st_case_885
	case 280:
		goto st_case_280
	case 281:
		goto st_case_281
	case 886:
		goto st_case_886
	case 282:
		goto st_case_282
	case 283:
		goto st_case_283
	case 284:
		goto st_case_284
	case 887:
		goto st_case_887
	case 888:
		goto st_case_888
	case 889:
		goto st_case_889
	case 890:
		goto st_case_890
	case 891:
		goto st_case_891
	case 892:
		goto st_case_892
	case 893:
		goto st_case_893
	case 285:
		goto st_case_285
	case 286:
		goto st_case_286
	case 287:
		goto st_case_287
	case 894:
		goto st_case_894
	case 288:
		goto st_case_288
	case 895:
		goto st_case_895
	case 896:
		goto st_case_896
	case 897:
		goto st_case_897
	case 898:
		goto st_case_898
	case 899:
		goto st_case_899
	case 900:
		goto st_case_900
	case 901:
		goto st_case_901
	case 289:
		goto st_case_289
	case 902:
		goto st_case_902
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
	case 903:
		goto st_case_903
	case 904:
		goto st_case_904
	case 905:
		goto st_case_905
	case 906:
		goto st_case_906
	case 296:
		goto st_case_296
	case 907:
		goto st_case_907
	case 908:
		goto st_case_908
	case 297:
		goto st_case_297
	case 909:
		goto st_case_909
	case 910:
		goto st_case_910
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
	case 310:
		goto st_case_310
	case 311:
		goto st_case_311
	case 312:
		goto st_case_312
	case 313:
		goto st_case_313
	case 314:
		goto st_case_314
	case 315:
		goto st_case_315
	case 316:
		goto st_case_316
	case 317:
		goto st_case_317
	case 318:
		goto st_case_318
	case 319:
		goto st_case_319
	case 320:
		goto st_case_320
	case 321:
		goto st_case_321
	case 322:
		goto st_case_322
	case 323:
		goto st_case_323
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
	case 331:
		goto st_case_331
	case 332:
		goto st_case_332
	case 333:
		goto st_case_333
	case 334:
		goto st_case_334
	case 335:
		goto st_case_335
	case 911:
		goto st_case_911
	case 336:
		goto st_case_336
	case 912:
		goto st_case_912
	case 913:
		goto st_case_913
	case 337:
		goto st_case_337
	case 914:
		goto st_case_914
	case 915:
		goto st_case_915
	case 338:
		goto st_case_338
	case 916:
		goto st_case_916
	case 339:
		goto st_case_339
	case 917:
		goto st_case_917
	case 340:
		goto st_case_340
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
	case 341:
		goto st_case_341
	case 927:
		goto st_case_927
	case 342:
		goto st_case_342
	case 928:
		goto st_case_928
	case 929:
		goto st_case_929
	case 343:
		goto st_case_343
	case 930:
		goto st_case_930
	case 344:
		goto st_case_344
	case 931:
		goto st_case_931
	case 932:
		goto st_case_932
	case 933:
		goto st_case_933
	case 934:
		goto st_case_934
	case 935:
		goto st_case_935
	case 936:
		goto st_case_936
	case 937:
		goto st_case_937
	case 938:
		goto st_case_938
	case 939:
		goto st_case_939
	case 940:
		goto st_case_940
	case 941:
		goto st_case_941
	case 942:
		goto st_case_942
	case 943:
		goto st_case_943
	case 345:
		goto st_case_345
	case 944:
		goto st_case_944
	case 945:
		goto st_case_945
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
	case 946:
		goto st_case_946
	case 947:
		goto st_case_947
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
	case 385:
		goto st_case_385
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
	case 948:
		goto st_case_948
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
	case 949:
		goto st_case_949
	case 950:
		goto st_case_950
	case 482:
		goto st_case_482
	case 483:
		goto st_case_483
	case 951:
		goto st_case_951
	case 484:
		goto st_case_484
	case 485:
		goto st_case_485
	case 486:
		goto st_case_486
	case 952:
		goto st_case_952
	case 487:
		goto st_case_487
	case 953:
		goto st_case_953
	case 488:
		goto st_case_488
	case 954:
		goto st_case_954
	case 955:
		goto st_case_955
	case 956:
		goto st_case_956
	case 957:
		goto st_case_957
	case 958:
		goto st_case_958
	case 959:
		goto st_case_959
	case 960:
		goto st_case_960
	case 489:
		goto st_case_489
	case 490:
		goto st_case_490
	case 491:
		goto st_case_491
	case 961:
		goto st_case_961
	case 492:
		goto st_case_492
	case 962:
		goto st_case_962
	case 963:
		goto st_case_963
	case 964:
		goto st_case_964
	case 965:
		goto st_case_965
	case 966:
		goto st_case_966
	case 967:
		goto st_case_967
	case 968:
		goto st_case_968
	case 493:
		goto st_case_493
	case 494:
		goto st_case_494
	case 969:
		goto st_case_969
	case 495:
		goto st_case_495
	case 970:
		goto st_case_970
	case 971:
		goto st_case_971
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
	case 972:
		goto st_case_972
	case 973:
		goto st_case_973
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
	case 526:
		goto st_case_526
	case 527:
		goto st_case_527
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
	case 530:
		goto st_case_530
	case 531:
		goto st_case_531
	case 532:
		goto st_case_532
	case 533:
		goto st_case_533
	case 534:
		goto st_case_534
	case 535:
		goto st_case_535
	case 974:
		goto st_case_974
	case 975:
		goto st_case_975
	case 976:
		goto st_case_976
	case 977:
		goto st_case_977
	case 536:
		goto st_case_536
	case 978:
		goto st_case_978
	case 979:
		goto st_case_979
	case 537:
		goto st_case_537
	case 980:
		goto st_case_980
	case 538:
		goto st_case_538
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 541:
		goto st_case_541
	case 542:
		goto st_case_542
	case 543:
		goto st_case_543
	case 544:
		goto st_case_544
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
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 48:
			goto tr0
		case 51:
			goto tr3
		case 65:
			goto st268
		case 68:
			goto st366
		case 70:
			goto st374
		case 74:
			goto st758
		case 77:
			goto st770
		case 78:
			goto st777
		case 79:
			goto st785
		case 83:
			goto st792
		case 84:
			goto st805
		case 87:
			goto st811
		case 97:
			goto st268
		case 100:
			goto st366
		case 102:
			goto st815
		case 106:
			goto st758
		case 109:
			goto st816
		case 110:
			goto st777
		case 111:
			goto st785
		case 115:
			goto st817
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
//line ragel/parse_datetime.go:2087
		if data[p] == 48 {
			goto st3
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st118
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
		if data[p] == 32 {
			goto tr22
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
//line ragel/parse_datetime.go:2141
		switch data[p] {
		case 48:
			goto tr25
		case 49:
			goto tr26
		case 65:
			goto st41
		case 68:
			goto st52
		case 70:
			goto st60
		case 74:
			goto st68
		case 77:
			goto st80
		case 78:
			goto st86
		case 79:
			goto st94
		case 83:
			goto st101
		case 97:
			goto st41
		case 100:
			goto st52
		case 102:
			goto st60
		case 106:
			goto st68
		case 109:
			goto st80
		case 110:
			goto st86
		case 111:
			goto st94
		case 115:
			goto st101
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr27
		}
		goto st0
tr25:
//line ragel/datetime.rl:5
 pb = p 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line ragel/parse_datetime.go:2193
		if 49 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto st0
tr27:
//line ragel/datetime.rl:5
 pb = p 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line ragel/parse_datetime.go:2207
		if data[p] == 32 {
			goto tr37
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr37
		}
		goto st0
tr37:
//line ragel/datetime.rl:9

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st9
tr87:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st9
tr93:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st9
tr100:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st9
tr109:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st9
tr119:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st9
tr127:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st9
tr130:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st9
tr136:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st9
tr140:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st9
tr144:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st9
tr153:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st9
tr161:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line ragel/parse_datetime.go:2274
		switch data[p] {
		case 48:
			goto tr38
		case 51:
			goto tr40
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr41
			}
		case data[p] >= 49:
			goto tr39
		}
		goto st0
tr38:
//line ragel/datetime.rl:5
 pb = p 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line ragel/parse_datetime.go:2299
		if 49 <= data[p] && data[p] <= 57 {
			goto st818
		}
		goto st0
tr41:
//line ragel/datetime.rl:5
 pb = p 
	goto st818
	st818:
		if p++; p == pe {
			goto _test_eof818
		}
	st_case_818:
//line ragel/parse_datetime.go:2313
		switch data[p] {
		case 32:
			goto tr1144
		case 43:
			goto tr1145
		case 44:
			goto tr1146
		case 45:
			goto tr1147
		case 47:
			goto tr1148
		case 84:
			goto tr1149
		case 90:
			goto tr1150
		case 95:
			goto tr1151
		case 116:
			goto tr1151
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1148
			}
		case data[p] >= 65:
			goto tr1148
		}
		goto st0
tr1297:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st819
tr1144:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st819
tr1272:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st819
tr1281:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st819
tr1289:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st819
	st819:
		if p++; p == pe {
			goto _test_eof819
		}
	st_case_819:
//line ragel/parse_datetime.go:2384
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr1155
		case 50:
			goto tr82
		case 65:
			goto tr1156
		case 66:
			goto tr1157
		case 90:
			goto tr1158
		case 95:
			goto tr1155
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr81
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1155
				}
			case data[p] >= 67:
				goto tr1155
			}
		default:
			goto tr83
		}
		goto st0
tr1171:
//line ragel/datetime.rl:166

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
	goto st11
tr1162:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st11
tr1175:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st11
tr1178:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line ragel/parse_datetime.go:2486
		switch data[p] {
		case 65:
			goto st12
		case 66:
			goto st13
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 68 {
			goto st820
		}
		goto st0
	st820:
		if p++; p == pe {
			goto _test_eof820
		}
	st_case_820:
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 67 {
			goto st821
		}
		goto st0
	st821:
		if p++; p == pe {
			goto _test_eof821
		}
	st_case_821:
		goto st0
tr1298:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st14
tr1145:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st14
tr1181:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st14
tr1193:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st14
tr1197:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st14
tr1212:
//line ragel/datetime.rl:117

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

	goto st14
tr1205:
//line ragel/datetime.rl:35

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

	goto st14
tr1225:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st14
tr1234:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st14
tr1242:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st14
tr1262:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st14
tr1273:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st14
tr1282:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st14
tr1290:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line ragel/parse_datetime.go:2674
		if data[p] == 50 {
			goto tr48
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr49
			}
		case data[p] >= 48:
			goto tr47
		}
		goto st0
tr47:
//line ragel/datetime.rl:5
 pb = p 
	goto st822
tr57:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st822
	st822:
		if p++; p == pe {
			goto _test_eof822
		}
	st_case_822:
//line ragel/parse_datetime.go:2702
		switch data[p] {
		case 32:
			goto tr1159
		case 58:
			goto tr1161
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st826
		}
		goto st0
tr1159:
//line ragel/datetime.rl:166

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
	goto st15
tr1164:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st15
tr1167:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line ragel/parse_datetime.go:2767
		switch data[p] {
		case 47:
			goto tr50
		case 65:
			goto tr51
		case 66:
			goto tr52
		case 95:
			goto tr50
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr50
			}
		case data[p] >= 67:
			goto tr50
		}
		goto st0
tr50:
//line ragel/datetime.rl:5
 pb = p 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line ragel/parse_datetime.go:2796
		switch data[p] {
		case 47:
			goto st17
		case 95:
			goto st17
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 47:
			goto st823
		case 95:
			goto st823
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st823
			}
		case data[p] >= 65:
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
			goto tr1162
		case 47:
			goto st823
		case 95:
			goto st823
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st823
			}
		case data[p] >= 65:
			goto st823
		}
		goto st0
tr51:
//line ragel/datetime.rl:5
 pb = p 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line ragel/parse_datetime.go:2863
		switch data[p] {
		case 47:
			goto st17
		case 68:
			goto st824
		case 95:
			goto st17
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto st0
	st824:
		if p++; p == pe {
			goto _test_eof824
		}
	st_case_824:
		switch data[p] {
		case 47:
			goto st823
		case 95:
			goto st823
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st823
			}
		case data[p] >= 65:
			goto st823
		}
		goto st0
tr52:
//line ragel/datetime.rl:5
 pb = p 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line ragel/parse_datetime.go:2910
		switch data[p] {
		case 47:
			goto st17
		case 67:
			goto st825
		case 95:
			goto st17
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto st0
	st825:
		if p++; p == pe {
			goto _test_eof825
		}
	st_case_825:
		switch data[p] {
		case 47:
			goto st823
		case 95:
			goto st823
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st823
			}
		case data[p] >= 65:
			goto st823
		}
		goto st0
tr49:
//line ragel/datetime.rl:5
 pb = p 
	goto st826
tr59:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st826
	st826:
		if p++; p == pe {
			goto _test_eof826
		}
	st_case_826:
//line ragel/parse_datetime.go:2963
		switch data[p] {
		case 32:
			goto tr1159
		case 58:
			goto tr1161
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st827
		}
		goto st0
	st827:
		if p++; p == pe {
			goto _test_eof827
		}
	st_case_827:
		if data[p] == 32 {
			goto tr1159
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st827
		}
		goto st0
tr1161:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st828
	st828:
		if p++; p == pe {
			goto _test_eof828
		}
	st_case_828:
//line ragel/parse_datetime.go:3002
		if data[p] == 32 {
			goto tr1164
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1166
			}
		case data[p] >= 48:
			goto tr1165
		}
		goto st0
tr1165:
//line ragel/datetime.rl:5
 pb = p 
	goto st829
	st829:
		if p++; p == pe {
			goto _test_eof829
		}
	st_case_829:
//line ragel/parse_datetime.go:3024
		if data[p] == 32 {
			goto tr1167
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st830
		}
		goto st0
tr1166:
//line ragel/datetime.rl:5
 pb = p 
	goto st830
	st830:
		if p++; p == pe {
			goto _test_eof830
		}
	st_case_830:
//line ragel/parse_datetime.go:3041
		if data[p] == 32 {
			goto tr1167
		}
		goto st0
tr48:
//line ragel/datetime.rl:5
 pb = p 
	goto st831
tr58:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st831
	st831:
		if p++; p == pe {
			goto _test_eof831
		}
	st_case_831:
//line ragel/parse_datetime.go:3061
		switch data[p] {
		case 32:
			goto tr1159
		case 58:
			goto tr1161
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st827
			}
		case data[p] >= 48:
			goto st826
		}
		goto st0
tr1300:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st20
tr1147:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st20
tr1182:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st20
tr1194:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st20
tr1198:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st20
tr1213:
//line ragel/datetime.rl:117

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

	goto st20
tr1207:
//line ragel/datetime.rl:35

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

	goto st20
tr1226:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st20
tr1235:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st20
tr1244:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st20
tr1264:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st20
tr1275:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st20
tr1284:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st20
tr1292:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line ragel/parse_datetime.go:3227
		if data[p] == 50 {
			goto tr58
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr59
			}
		case data[p] >= 48:
			goto tr57
		}
		goto st0
tr1155:
//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1301:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1183:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1199:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1227:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1236:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1245:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1214:
//line ragel/datetime.rl:117

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
	goto st21
tr1265:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1208:
//line ragel/datetime.rl:35

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
	goto st21
tr1148:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1276:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1285:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
tr1293:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line ragel/parse_datetime.go:3395
		switch data[p] {
		case 47:
			goto st22
		case 95:
			goto st22
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		case data[p] >= 65:
			goto st22
		}
		goto st0
tr1270:
//line ragel/datetime.rl:5
 pb = p 
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line ragel/parse_datetime.go:3420
		switch data[p] {
		case 47:
			goto st832
		case 95:
			goto st832
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st832
			}
		case data[p] >= 65:
			goto st832
		}
		goto st0
tr1195:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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
	goto st832
	st832:
		if p++; p == pe {
			goto _test_eof832
		}
	st_case_832:
//line ragel/parse_datetime.go:3468
		switch data[p] {
		case 32:
			goto tr1162
		case 43:
			goto tr1169
		case 45:
			goto tr1170
		case 47:
			goto st832
		case 95:
			goto st832
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st832
			}
		case data[p] >= 65:
			goto st832
		}
		goto st0
tr1169:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line ragel/parse_datetime.go:3502
		if data[p] == 50 {
			goto tr63
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr64
			}
		case data[p] >= 48:
			goto tr62
		}
		goto st0
tr62:
//line ragel/datetime.rl:5
 pb = p 
	goto st833
tr65:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st833
	st833:
		if p++; p == pe {
			goto _test_eof833
		}
	st_case_833:
//line ragel/parse_datetime.go:3530
		switch data[p] {
		case 32:
			goto tr1171
		case 58:
			goto tr1173
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st834
		}
		goto st0
tr64:
//line ragel/datetime.rl:5
 pb = p 
	goto st834
tr67:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st834
	st834:
		if p++; p == pe {
			goto _test_eof834
		}
	st_case_834:
//line ragel/parse_datetime.go:3556
		switch data[p] {
		case 32:
			goto tr1171
		case 58:
			goto tr1173
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st835
		}
		goto st0
	st835:
		if p++; p == pe {
			goto _test_eof835
		}
	st_case_835:
		if data[p] == 32 {
			goto tr1171
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st835
		}
		goto st0
tr1173:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st836
	st836:
		if p++; p == pe {
			goto _test_eof836
		}
	st_case_836:
//line ragel/parse_datetime.go:3595
		if data[p] == 32 {
			goto tr1175
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1177
			}
		case data[p] >= 48:
			goto tr1176
		}
		goto st0
tr1176:
//line ragel/datetime.rl:5
 pb = p 
	goto st837
	st837:
		if p++; p == pe {
			goto _test_eof837
		}
	st_case_837:
//line ragel/parse_datetime.go:3617
		if data[p] == 32 {
			goto tr1178
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st838
		}
		goto st0
tr1177:
//line ragel/datetime.rl:5
 pb = p 
	goto st838
	st838:
		if p++; p == pe {
			goto _test_eof838
		}
	st_case_838:
//line ragel/parse_datetime.go:3634
		if data[p] == 32 {
			goto tr1178
		}
		goto st0
tr63:
//line ragel/datetime.rl:5
 pb = p 
	goto st839
tr66:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st839
	st839:
		if p++; p == pe {
			goto _test_eof839
		}
	st_case_839:
//line ragel/parse_datetime.go:3654
		switch data[p] {
		case 32:
			goto tr1171
		case 58:
			goto tr1173
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st835
			}
		case data[p] >= 48:
			goto st834
		}
		goto st0
tr1170:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line ragel/parse_datetime.go:3682
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
tr81:
//line ragel/datetime.rl:5
 pb = p 
	goto st840
	st840:
		if p++; p == pe {
			goto _test_eof840
		}
	st_case_840:
//line ragel/parse_datetime.go:3704
		switch data[p] {
		case 32:
			goto tr1180
		case 43:
			goto tr1181
		case 45:
			goto tr1182
		case 47:
			goto tr1183
		case 58:
			goto tr1185
		case 65:
			goto tr1186
		case 80:
			goto tr1186
		case 90:
			goto tr1187
		case 95:
			goto tr1183
		case 97:
			goto tr1188
		case 112:
			goto tr1188
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st847
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1183
			}
		default:
			goto tr1183
		}
		goto st0
tr1180:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st841
tr1196:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st841
tr1250:
//line ragel/datetime.rl:117

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

	goto st841
tr1224:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st841
tr1233:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st841
tr1241:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st841
tr1261:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st841
	st841:
		if p++; p == pe {
			goto _test_eof841
		}
	st_case_841:
//line ragel/parse_datetime.go:3814
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr1155
		case 65:
			goto tr1189
		case 66:
			goto tr1157
		case 80:
			goto tr1190
		case 90:
			goto tr1158
		case 95:
			goto tr1155
		case 97:
			goto tr1191
		case 112:
			goto tr1191
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1155
			}
		case data[p] >= 67:
			goto tr1155
		}
		goto st0
tr1189:
//line ragel/datetime.rl:5
 pb = p 
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line ragel/parse_datetime.go:3857
		switch data[p] {
		case 47:
			goto st22
		case 68:
			goto st842
		case 77:
			goto st843
		case 95:
			goto st22
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		case data[p] >= 65:
			goto st22
		}
		goto st0
	st842:
		if p++; p == pe {
			goto _test_eof842
		}
	st_case_842:
		switch data[p] {
		case 47:
			goto st832
		case 95:
			goto st832
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st832
			}
		case data[p] >= 65:
			goto st832
		}
		goto st0
	st843:
		if p++; p == pe {
			goto _test_eof843
		}
	st_case_843:
		switch data[p] {
		case 32:
			goto tr1192
		case 43:
			goto tr1193
		case 45:
			goto tr1194
		case 47:
			goto tr1195
		case 95:
			goto tr1195
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1195
			}
		case data[p] >= 65:
			goto tr1195
		}
		goto st0
tr1192:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st844
tr1211:
//line ragel/datetime.rl:117

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

	goto st844
tr1204:
//line ragel/datetime.rl:35

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

	goto st844
tr1330:
//line ragel/datetime.rl:35

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

	goto st844
	st844:
		if p++; p == pe {
			goto _test_eof844
		}
	st_case_844:
//line ragel/parse_datetime.go:4022
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr1155
		case 65:
			goto tr1156
		case 66:
			goto tr1157
		case 90:
			goto tr1158
		case 95:
			goto tr1155
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1155
			}
		case data[p] >= 67:
			goto tr1155
		}
		goto st0
tr1156:
//line ragel/datetime.rl:5
 pb = p 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line ragel/parse_datetime.go:4059
		switch data[p] {
		case 47:
			goto st22
		case 68:
			goto st842
		case 95:
			goto st22
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		case data[p] >= 65:
			goto st22
		}
		goto st0
tr1157:
//line ragel/datetime.rl:5
 pb = p 
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line ragel/parse_datetime.go:4086
		switch data[p] {
		case 47:
			goto st22
		case 67:
			goto st845
		case 95:
			goto st22
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		case data[p] >= 65:
			goto st22
		}
		goto st0
	st845:
		if p++; p == pe {
			goto _test_eof845
		}
	st_case_845:
		switch data[p] {
		case 47:
			goto st832
		case 95:
			goto st832
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st832
			}
		case data[p] >= 65:
			goto st832
		}
		goto st0
tr1158:
//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1303:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1187:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1202:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1231:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1239:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1248:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1216:
//line ragel/datetime.rl:117

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
	goto st846
tr1267:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1210:
//line ragel/datetime.rl:35

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
	goto st846
tr1150:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1278:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1287:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
tr1295:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st846
	st846:
		if p++; p == pe {
			goto _test_eof846
		}
	st_case_846:
//line ragel/parse_datetime.go:4279
		switch data[p] {
		case 32:
			goto tr1175
		case 47:
			goto st22
		case 95:
			goto st22
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		case data[p] >= 65:
			goto st22
		}
		goto st0
tr1190:
//line ragel/datetime.rl:5
 pb = p 
	goto st28
tr1186:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st28
tr1201:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st28
tr1230:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st28
tr1238:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st28
tr1247:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st28
tr1252:
//line ragel/datetime.rl:117

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
	goto st28
tr1266:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line ragel/parse_datetime.go:4387
		switch data[p] {
		case 47:
			goto st22
		case 77:
			goto st843
		case 95:
			goto st22
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		case data[p] >= 65:
			goto st22
		}
		goto st0
tr1191:
//line ragel/datetime.rl:5
 pb = p 
	goto st29
tr1188:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st29
tr1203:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st29
tr1232:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st29
tr1240:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st29
tr1249:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st29
tr1253:
//line ragel/datetime.rl:117

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
	goto st29
tr1268:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line ragel/parse_datetime.go:4495
		switch data[p] {
		case 47:
			goto st22
		case 95:
			goto st22
		case 109:
			goto st843
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		case data[p] >= 65:
			goto st22
		}
		goto st0
	st847:
		if p++; p == pe {
			goto _test_eof847
		}
	st_case_847:
		switch data[p] {
		case 32:
			goto tr1196
		case 43:
			goto tr1197
		case 45:
			goto tr1198
		case 47:
			goto tr1199
		case 58:
			goto tr1200
		case 65:
			goto tr1201
		case 80:
			goto tr1201
		case 90:
			goto tr1202
		case 95:
			goto tr1199
		case 97:
			goto tr1203
		case 112:
			goto tr1203
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st30
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1199
			}
		default:
			goto tr1199
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if 48 <= data[p] && data[p] <= 57 {
			goto st848
		}
		goto st0
	st848:
		if p++; p == pe {
			goto _test_eof848
		}
	st_case_848:
		switch data[p] {
		case 32:
			goto tr1204
		case 43:
			goto tr1205
		case 45:
			goto tr1207
		case 47:
			goto tr1208
		case 90:
			goto tr1210
		case 95:
			goto tr1208
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1206
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1208
				}
			case data[p] >= 65:
				goto tr1208
			}
		default:
			goto st32
		}
		goto st0
tr1206:
//line ragel/datetime.rl:35

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
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line ragel/parse_datetime.go:4623
		if 48 <= data[p] && data[p] <= 57 {
			goto tr72
		}
		goto st0
tr72:
//line ragel/datetime.rl:5
 pb = p 
	goto st849
	st849:
		if p++; p == pe {
			goto _test_eof849
		}
	st_case_849:
//line ragel/parse_datetime.go:4637
		switch data[p] {
		case 32:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st850
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st850:
		if p++; p == pe {
			goto _test_eof850
		}
	st_case_850:
		switch data[p] {
		case 32:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st851
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st851:
		if p++; p == pe {
			goto _test_eof851
		}
	st_case_851:
		switch data[p] {
		case 32:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st852
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st852:
		if p++; p == pe {
			goto _test_eof852
		}
	st_case_852:
		switch data[p] {
		case 32:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st853
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st853:
		if p++; p == pe {
			goto _test_eof853
		}
	st_case_853:
		switch data[p] {
		case 32:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st854
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st854:
		if p++; p == pe {
			goto _test_eof854
		}
	st_case_854:
		switch data[p] {
		case 32:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st855
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st855:
		if p++; p == pe {
			goto _test_eof855
		}
	st_case_855:
		switch data[p] {
		case 32:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st856
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st856:
		if p++; p == pe {
			goto _test_eof856
		}
	st_case_856:
		switch data[p] {
		case 32:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st857
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st857:
		if p++; p == pe {
			goto _test_eof857
		}
	st_case_857:
		switch data[p] {
		case 32:
			goto tr1211
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		case data[p] >= 65:
			goto tr1214
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if 48 <= data[p] && data[p] <= 57 {
			goto st858
		}
		goto st0
	st858:
		if p++; p == pe {
			goto _test_eof858
		}
	st_case_858:
		switch data[p] {
		case 32:
			goto tr1204
		case 43:
			goto tr1205
		case 45:
			goto tr1207
		case 47:
			goto tr1208
		case 90:
			goto tr1210
		case 95:
			goto tr1208
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1206
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1208
			}
		default:
			goto tr1208
		}
		goto st0
tr1185:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st33
tr1200:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line ragel/parse_datetime.go:4975
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr75
			}
		case data[p] >= 48:
			goto tr74
		}
		goto st0
tr74:
//line ragel/datetime.rl:5
 pb = p 
	goto st859
	st859:
		if p++; p == pe {
			goto _test_eof859
		}
	st_case_859:
//line ragel/parse_datetime.go:4994
		switch data[p] {
		case 32:
			goto tr1224
		case 43:
			goto tr1225
		case 45:
			goto tr1226
		case 47:
			goto tr1227
		case 58:
			goto tr1229
		case 65:
			goto tr1230
		case 80:
			goto tr1230
		case 90:
			goto tr1231
		case 95:
			goto tr1227
		case 97:
			goto tr1232
		case 112:
			goto tr1232
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st860
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1227
			}
		default:
			goto tr1227
		}
		goto st0
	st860:
		if p++; p == pe {
			goto _test_eof860
		}
	st_case_860:
		switch data[p] {
		case 32:
			goto tr1233
		case 43:
			goto tr1234
		case 45:
			goto tr1235
		case 47:
			goto tr1236
		case 58:
			goto tr1237
		case 65:
			goto tr1238
		case 80:
			goto tr1238
		case 90:
			goto tr1239
		case 95:
			goto tr1236
		case 97:
			goto tr1240
		case 112:
			goto tr1240
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1236
			}
		case data[p] >= 66:
			goto tr1236
		}
		goto st0
tr1229:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st34
tr1237:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line ragel/parse_datetime.go:5087
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr77
			}
		case data[p] >= 48:
			goto tr76
		}
		goto st0
tr76:
//line ragel/datetime.rl:5
 pb = p 
	goto st861
	st861:
		if p++; p == pe {
			goto _test_eof861
		}
	st_case_861:
//line ragel/parse_datetime.go:5106
		switch data[p] {
		case 32:
			goto tr1241
		case 43:
			goto tr1242
		case 45:
			goto tr1244
		case 47:
			goto tr1245
		case 65:
			goto tr1247
		case 80:
			goto tr1247
		case 90:
			goto tr1248
		case 95:
			goto tr1245
		case 97:
			goto tr1249
		case 112:
			goto tr1249
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1243
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
			goto st871
		}
		goto st0
tr1243:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st35
tr1263:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line ragel/parse_datetime.go:5164
		if 48 <= data[p] && data[p] <= 57 {
			goto tr78
		}
		goto st0
tr78:
//line ragel/datetime.rl:5
 pb = p 
	goto st862
	st862:
		if p++; p == pe {
			goto _test_eof862
		}
	st_case_862:
//line ragel/parse_datetime.go:5178
		switch data[p] {
		case 32:
			goto tr1250
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		case 97:
			goto tr1253
		case 112:
			goto tr1253
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st863
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st863:
		if p++; p == pe {
			goto _test_eof863
		}
	st_case_863:
		switch data[p] {
		case 32:
			goto tr1250
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		case 97:
			goto tr1253
		case 112:
			goto tr1253
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st864
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st864:
		if p++; p == pe {
			goto _test_eof864
		}
	st_case_864:
		switch data[p] {
		case 32:
			goto tr1250
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		case 97:
			goto tr1253
		case 112:
			goto tr1253
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st865
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st865:
		if p++; p == pe {
			goto _test_eof865
		}
	st_case_865:
		switch data[p] {
		case 32:
			goto tr1250
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		case 97:
			goto tr1253
		case 112:
			goto tr1253
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st866
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st866:
		if p++; p == pe {
			goto _test_eof866
		}
	st_case_866:
		switch data[p] {
		case 32:
			goto tr1250
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		case 97:
			goto tr1253
		case 112:
			goto tr1253
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st867
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st867:
		if p++; p == pe {
			goto _test_eof867
		}
	st_case_867:
		switch data[p] {
		case 32:
			goto tr1250
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		case 97:
			goto tr1253
		case 112:
			goto tr1253
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st868
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st868:
		if p++; p == pe {
			goto _test_eof868
		}
	st_case_868:
		switch data[p] {
		case 32:
			goto tr1250
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		case 97:
			goto tr1253
		case 112:
			goto tr1253
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st869
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st869:
		if p++; p == pe {
			goto _test_eof869
		}
	st_case_869:
		switch data[p] {
		case 32:
			goto tr1250
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		case 97:
			goto tr1253
		case 112:
			goto tr1253
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st870
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		default:
			goto tr1214
		}
		goto st0
	st870:
		if p++; p == pe {
			goto _test_eof870
		}
	st_case_870:
		switch data[p] {
		case 32:
			goto tr1250
		case 43:
			goto tr1212
		case 45:
			goto tr1213
		case 47:
			goto tr1214
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1216
		case 95:
			goto tr1214
		case 97:
			goto tr1253
		case 112:
			goto tr1253
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1214
			}
		case data[p] >= 66:
			goto tr1214
		}
		goto st0
	st871:
		if p++; p == pe {
			goto _test_eof871
		}
	st_case_871:
		switch data[p] {
		case 32:
			goto tr1261
		case 43:
			goto tr1262
		case 45:
			goto tr1264
		case 47:
			goto tr1265
		case 65:
			goto tr1266
		case 80:
			goto tr1266
		case 90:
			goto tr1267
		case 95:
			goto tr1265
		case 97:
			goto tr1268
		case 112:
			goto tr1268
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1263
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1265
			}
		default:
			goto tr1265
		}
		goto st0
tr77:
//line ragel/datetime.rl:5
 pb = p 
	goto st872
	st872:
		if p++; p == pe {
			goto _test_eof872
		}
	st_case_872:
//line ragel/parse_datetime.go:5579
		switch data[p] {
		case 32:
			goto tr1241
		case 43:
			goto tr1242
		case 45:
			goto tr1244
		case 47:
			goto tr1245
		case 65:
			goto tr1247
		case 80:
			goto tr1247
		case 90:
			goto tr1248
		case 95:
			goto tr1245
		case 97:
			goto tr1249
		case 112:
			goto tr1249
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1243
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1245
			}
		default:
			goto tr1245
		}
		goto st0
tr75:
//line ragel/datetime.rl:5
 pb = p 
	goto st873
	st873:
		if p++; p == pe {
			goto _test_eof873
		}
	st_case_873:
//line ragel/parse_datetime.go:5624
		switch data[p] {
		case 32:
			goto tr1224
		case 43:
			goto tr1225
		case 45:
			goto tr1226
		case 47:
			goto tr1227
		case 58:
			goto tr1229
		case 65:
			goto tr1230
		case 80:
			goto tr1230
		case 90:
			goto tr1231
		case 95:
			goto tr1227
		case 97:
			goto tr1232
		case 112:
			goto tr1232
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1227
			}
		case data[p] >= 66:
			goto tr1227
		}
		goto st0
tr82:
//line ragel/datetime.rl:5
 pb = p 
	goto st874
	st874:
		if p++; p == pe {
			goto _test_eof874
		}
	st_case_874:
//line ragel/parse_datetime.go:5667
		switch data[p] {
		case 32:
			goto tr1180
		case 43:
			goto tr1181
		case 45:
			goto tr1182
		case 47:
			goto tr1183
		case 58:
			goto tr1185
		case 65:
			goto tr1186
		case 80:
			goto tr1186
		case 90:
			goto tr1187
		case 95:
			goto tr1183
		case 97:
			goto tr1188
		case 112:
			goto tr1188
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st847
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1183
				}
			case data[p] >= 66:
				goto tr1183
			}
		default:
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if 48 <= data[p] && data[p] <= 57 {
			goto st30
		}
		goto st0
tr83:
//line ragel/datetime.rl:5
 pb = p 
	goto st875
	st875:
		if p++; p == pe {
			goto _test_eof875
		}
	st_case_875:
//line ragel/parse_datetime.go:5728
		switch data[p] {
		case 32:
			goto tr1180
		case 43:
			goto tr1181
		case 45:
			goto tr1182
		case 47:
			goto tr1183
		case 58:
			goto tr1185
		case 65:
			goto tr1186
		case 80:
			goto tr1186
		case 90:
			goto tr1187
		case 95:
			goto tr1183
		case 97:
			goto tr1188
		case 112:
			goto tr1188
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st36
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1183
			}
		default:
			goto tr1183
		}
		goto st0
tr1299:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st37
tr1146:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st37
tr1274:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st37
tr1283:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st37
tr1291:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line ragel/parse_datetime.go:5807
		if data[p] == 32 {
			goto st38
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 50 {
			goto tr82
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr83
			}
		case data[p] >= 48:
			goto tr81
		}
		goto st0
tr1302:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st876
tr1149:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st876
tr1277:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st876
tr1286:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st876
tr1294:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st876
	st876:
		if p++; p == pe {
			goto _test_eof876
		}
	st_case_876:
//line ragel/parse_datetime.go:5880
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr1270
		case 50:
			goto tr82
		case 90:
			goto tr1271
		case 95:
			goto tr1270
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr81
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1270
				}
			case data[p] >= 65:
				goto tr1270
			}
		default:
			goto tr83
		}
		goto st0
tr1271:
//line ragel/datetime.rl:5
 pb = p 
	goto st877
	st877:
		if p++; p == pe {
			goto _test_eof877
		}
	st_case_877:
//line ragel/parse_datetime.go:5924
		switch data[p] {
		case 32:
			goto tr1175
		case 47:
			goto st832
		case 95:
			goto st832
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st832
			}
		case data[p] >= 65:
			goto st832
		}
		goto st0
tr1304:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1151:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1279:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1288:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
tr1296:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line ragel/parse_datetime.go:5993
		switch data[p] {
		case 47:
			goto st22
		case 50:
			goto tr82
		case 95:
			goto st22
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr81
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st22
				}
			case data[p] >= 65:
				goto st22
			}
		default:
			goto tr83
		}
		goto st0
tr39:
//line ragel/datetime.rl:5
 pb = p 
	goto st878
	st878:
		if p++; p == pe {
			goto _test_eof878
		}
	st_case_878:
//line ragel/parse_datetime.go:6029
		switch data[p] {
		case 32:
			goto tr1144
		case 43:
			goto tr1145
		case 44:
			goto tr1146
		case 45:
			goto tr1147
		case 47:
			goto tr1148
		case 84:
			goto tr1149
		case 90:
			goto tr1150
		case 95:
			goto tr1151
		case 116:
			goto tr1151
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st818
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1148
			}
		default:
			goto tr1148
		}
		goto st0
tr40:
//line ragel/datetime.rl:5
 pb = p 
	goto st879
	st879:
		if p++; p == pe {
			goto _test_eof879
		}
	st_case_879:
//line ragel/parse_datetime.go:6072
		switch data[p] {
		case 32:
			goto tr1144
		case 43:
			goto tr1145
		case 44:
			goto tr1146
		case 45:
			goto tr1147
		case 47:
			goto tr1148
		case 84:
			goto tr1149
		case 90:
			goto tr1150
		case 95:
			goto tr1151
		case 116:
			goto tr1151
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 49 {
				goto st818
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1148
			}
		default:
			goto tr1148
		}
		goto st0
tr26:
//line ragel/datetime.rl:5
 pb = p 
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line ragel/parse_datetime.go:6115
		if data[p] == 32 {
			goto tr37
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 50 {
				goto st8
			}
		case data[p] >= 45:
			goto tr37
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 112:
			goto st42
		case 117:
			goto st47
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 114 {
			goto st43
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 32:
			goto tr87
		case 46:
			goto tr88
		case 105:
			goto st45
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr87
		}
		goto st0
tr88:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st44
tr94:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st44
tr101:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st44
tr110:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st44
tr120:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st44
tr128:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st44
tr131:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st44
tr137:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st44
tr141:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st44
tr145:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st44
tr154:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st44
tr162:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line ragel/parse_datetime.go:6219
		switch data[p] {
		case 32:
			goto st9
		case 48:
			goto tr38
		case 51:
			goto tr40
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st9
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr41
			}
		default:
			goto tr39
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if data[p] == 108 {
			goto st46
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 32:
			goto tr87
		case 46:
			goto tr88
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr87
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 103 {
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 32:
			goto tr93
		case 46:
			goto tr94
		case 117:
			goto st49
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr93
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 115 {
			goto st50
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		if data[p] == 116 {
			goto st51
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 32:
			goto tr93
		case 46:
			goto tr94
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr93
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 101 {
			goto st53
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 99 {
			goto st54
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 32:
			goto tr100
		case 46:
			goto tr101
		case 101:
			goto st55
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr100
		}
		goto st0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 109 {
			goto st56
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 98 {
			goto st57
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 101 {
			goto st58
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 114 {
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
		if data[p] == 101 {
			goto st61
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 98 {
			goto st62
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 32:
			goto tr109
		case 46:
			goto tr110
		case 114:
			goto st63
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr109
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 117 {
			goto st64
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 97 {
			goto st65
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 114 {
			goto st66
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		if data[p] == 121 {
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
			goto tr109
		case 46:
			goto tr110
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr109
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 97:
			goto st69
		case 117:
			goto st75
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		if data[p] == 110 {
			goto st70
		}
		goto st0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 32:
			goto tr119
		case 46:
			goto tr120
		case 117:
			goto st71
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr119
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		if data[p] == 97 {
			goto st72
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 114 {
			goto st73
		}
		goto st0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 121 {
			goto st74
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 32:
			goto tr119
		case 46:
			goto tr120
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr119
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 108:
			goto st76
		case 110:
			goto st78
		}
		goto st0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 32:
			goto tr127
		case 46:
			goto tr128
		case 121:
			goto st77
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr127
		}
		goto st0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 32:
			goto tr127
		case 46:
			goto tr128
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr127
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 32:
			goto tr130
		case 46:
			goto tr131
		case 101:
			goto st79
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr130
		}
		goto st0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 32:
			goto tr130
		case 46:
			goto tr131
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr130
		}
		goto st0
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		if data[p] == 97 {
			goto st81
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 114:
			goto st82
		case 121:
			goto st85
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 32:
			goto tr136
		case 46:
			goto tr137
		case 99:
			goto st83
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr136
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		if data[p] == 104 {
			goto st84
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 32:
			goto tr136
		case 46:
			goto tr137
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr136
		}
		goto st0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
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
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		if data[p] == 111 {
			goto st87
		}
		goto st0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		if data[p] == 118 {
			goto st88
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 32:
			goto tr144
		case 46:
			goto tr145
		case 101:
			goto st89
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr144
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		if data[p] == 109 {
			goto st90
		}
		goto st0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		if data[p] == 98 {
			goto st91
		}
		goto st0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		if data[p] == 101 {
			goto st92
		}
		goto st0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		if data[p] == 114 {
			goto st93
		}
		goto st0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 32:
			goto tr144
		case 46:
			goto tr145
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr144
		}
		goto st0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		if data[p] == 99 {
			goto st95
		}
		goto st0
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		if data[p] == 116 {
			goto st96
		}
		goto st0
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 32:
			goto tr153
		case 46:
			goto tr154
		case 111:
			goto st97
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr153
		}
		goto st0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		if data[p] == 98 {
			goto st98
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		if data[p] == 101 {
			goto st99
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		if data[p] == 114 {
			goto st100
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
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
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		if data[p] == 101 {
			goto st102
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		if data[p] == 112 {
			goto st103
		}
		goto st0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 32:
			goto tr161
		case 46:
			goto tr162
		case 116:
			goto st104
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr161
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 32:
			goto tr161
		case 46:
			goto tr162
		case 101:
			goto st105
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr161
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		if data[p] == 109 {
			goto st106
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		if data[p] == 98 {
			goto st107
		}
		goto st0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		if data[p] == 101 {
			goto st108
		}
		goto st0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		if data[p] == 114 {
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
			goto tr161
		case 46:
			goto tr162
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr161
		}
		goto st0
tr23:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line ragel/parse_datetime.go:7006
		switch data[p] {
		case 48:
			goto tr169
		case 49:
			goto tr170
		case 65:
			goto st41
		case 68:
			goto st52
		case 70:
			goto st60
		case 74:
			goto st68
		case 77:
			goto st80
		case 78:
			goto st86
		case 79:
			goto st94
		case 83:
			goto st101
		case 97:
			goto st41
		case 100:
			goto st52
		case 102:
			goto st60
		case 106:
			goto st68
		case 109:
			goto st80
		case 110:
			goto st86
		case 111:
			goto st94
		case 115:
			goto st101
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr171
		}
		goto st0
tr169:
//line ragel/datetime.rl:5
 pb = p 
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line ragel/parse_datetime.go:7058
		if data[p] == 48 {
			goto st112
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st113
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		if 48 <= data[p] && data[p] <= 57 {
			goto st880
		}
		goto st0
	st880:
		if p++; p == pe {
			goto _test_eof880
		}
	st_case_880:
		switch data[p] {
		case 32:
			goto tr1272
		case 43:
			goto tr1273
		case 44:
			goto tr1274
		case 45:
			goto tr1275
		case 47:
			goto tr1276
		case 84:
			goto tr1277
		case 90:
			goto tr1278
		case 95:
			goto tr1279
		case 116:
			goto tr1279
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1276
			}
		case data[p] >= 65:
			goto tr1276
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		if data[p] == 32 {
			goto tr37
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st880
			}
		case data[p] >= 45:
			goto tr37
		}
		goto st0
tr170:
//line ragel/datetime.rl:5
 pb = p 
	goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
//line ragel/parse_datetime.go:7135
		if data[p] == 32 {
			goto tr37
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr37
			}
		case data[p] > 50:
			if 51 <= data[p] && data[p] <= 57 {
				goto st112
			}
		default:
			goto st113
		}
		goto st0
tr171:
//line ragel/datetime.rl:5
 pb = p 
	goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
//line ragel/parse_datetime.go:7161
		if data[p] == 32 {
			goto tr37
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st112
			}
		case data[p] >= 45:
			goto tr37
		}
		goto st0
tr24:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line ragel/parse_datetime.go:7187
		if 48 <= data[p] && data[p] <= 57 {
			goto st117
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		if 48 <= data[p] && data[p] <= 57 {
			goto st881
		}
		goto st0
	st881:
		if p++; p == pe {
			goto _test_eof881
		}
	st_case_881:
		switch data[p] {
		case 32:
			goto tr1272
		case 43:
			goto tr1273
		case 44:
			goto tr1274
		case 45:
			goto tr1275
		case 47:
			goto tr1276
		case 84:
			goto tr1277
		case 90:
			goto tr1278
		case 95:
			goto tr1279
		case 116:
			goto tr1279
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st882
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1276
			}
		default:
			goto tr1276
		}
		goto st0
	st882:
		if p++; p == pe {
			goto _test_eof882
		}
	st_case_882:
		switch data[p] {
		case 32:
			goto tr1281
		case 43:
			goto tr1282
		case 44:
			goto tr1283
		case 45:
			goto tr1284
		case 47:
			goto tr1285
		case 84:
			goto tr1286
		case 90:
			goto tr1287
		case 95:
			goto tr1288
		case 116:
			goto tr1288
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1285
			}
		case data[p] >= 65:
			goto tr1285
		}
		goto st0
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		if data[p] == 32 {
			goto tr177
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] >= 45:
			goto tr178
		}
		goto st0
tr177:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line ragel/parse_datetime.go:7306
		switch data[p] {
		case 65:
			goto st120
		case 68:
			goto st137
		case 70:
			goto st145
		case 74:
			goto st153
		case 77:
			goto st165
		case 78:
			goto st171
		case 79:
			goto st179
		case 83:
			goto st186
		case 97:
			goto st120
		case 100:
			goto st137
		case 102:
			goto st145
		case 106:
			goto st153
		case 109:
			goto st165
		case 110:
			goto st171
		case 111:
			goto st179
		case 115:
			goto st186
		}
		goto st0
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 112:
			goto st121
		case 117:
			goto st132
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
			goto tr190
		case 46:
			goto tr192
		case 105:
			goto st130
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr191
		}
		goto st0
tr190:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st123
tr204:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st123
tr212:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st123
tr222:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st123
tr233:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st123
tr242:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st123
tr246:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st123
tr253:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st123
tr258:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st123
tr263:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st123
tr273:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st123
tr282:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
//line ragel/parse_datetime.go:7433
		if 48 <= data[p] && data[p] <= 57 {
			goto tr194
		}
		goto st0
tr194:
//line ragel/datetime.rl:5
 pb = p 
	goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line ragel/parse_datetime.go:7447
		if 48 <= data[p] && data[p] <= 57 {
			goto st883
		}
		goto st0
	st883:
		if p++; p == pe {
			goto _test_eof883
		}
	st_case_883:
		switch data[p] {
		case 32:
			goto tr1289
		case 43:
			goto tr1290
		case 44:
			goto tr1291
		case 45:
			goto tr1292
		case 47:
			goto tr1293
		case 84:
			goto tr1294
		case 90:
			goto tr1295
		case 95:
			goto tr1296
		case 116:
			goto tr1296
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st125
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1293
			}
		default:
			goto tr1293
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		if 48 <= data[p] && data[p] <= 57 {
			goto st884
		}
		goto st0
	st884:
		if p++; p == pe {
			goto _test_eof884
		}
	st_case_884:
		switch data[p] {
		case 32:
			goto tr1297
		case 43:
			goto tr1298
		case 44:
			goto tr1299
		case 45:
			goto tr1300
		case 47:
			goto tr1301
		case 84:
			goto tr1302
		case 90:
			goto tr1303
		case 95:
			goto tr1304
		case 116:
			goto tr1304
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1301
			}
		case data[p] >= 65:
			goto tr1301
		}
		goto st0
tr191:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st126
tr205:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st126
tr213:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st126
tr223:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st126
tr234:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st126
tr243:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st126
tr247:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st126
tr254:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st126
tr259:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st126
tr264:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st126
tr274:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st126
tr283:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st126
tr385:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line ragel/parse_datetime.go:7597
		if 48 <= data[p] && data[p] <= 57 {
			goto tr197
		}
		goto st0
tr197:
//line ragel/datetime.rl:5
 pb = p 
	goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line ragel/parse_datetime.go:7611
		if 48 <= data[p] && data[p] <= 57 {
			goto st128
		}
		goto st0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
		goto st0
tr192:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st129
tr206:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st129
tr214:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st129
tr224:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st129
tr235:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st129
tr244:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st129
tr248:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st129
tr255:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st129
tr260:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st129
tr265:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st129
tr275:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st129
tr284:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st129
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
//line ragel/parse_datetime.go:7678
		if data[p] == 32 {
			goto st123
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr197
			}
		case data[p] >= 45:
			goto st126
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		if data[p] == 108 {
			goto st131
		}
		goto st0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 32:
			goto tr190
		case 46:
			goto tr192
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr191
		}
		goto st0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		if data[p] == 103 {
			goto st133
		}
		goto st0
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 32:
			goto tr204
		case 46:
			goto tr206
		case 117:
			goto st134
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr205
		}
		goto st0
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		if data[p] == 115 {
			goto st135
		}
		goto st0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		if data[p] == 116 {
			goto st136
		}
		goto st0
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 32:
			goto tr204
		case 46:
			goto tr206
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr205
		}
		goto st0
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		if data[p] == 101 {
			goto st138
		}
		goto st0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		if data[p] == 99 {
			goto st139
		}
		goto st0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 32:
			goto tr212
		case 46:
			goto tr214
		case 101:
			goto st140
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr213
		}
		goto st0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		if data[p] == 109 {
			goto st141
		}
		goto st0
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		if data[p] == 98 {
			goto st142
		}
		goto st0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		if data[p] == 101 {
			goto st143
		}
		goto st0
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		if data[p] == 114 {
			goto st144
		}
		goto st0
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 32:
			goto tr212
		case 46:
			goto tr214
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr213
		}
		goto st0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		if data[p] == 101 {
			goto st146
		}
		goto st0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		if data[p] == 98 {
			goto st147
		}
		goto st0
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 32:
			goto tr222
		case 46:
			goto tr224
		case 114:
			goto st148
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr223
		}
		goto st0
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		if data[p] == 117 {
			goto st149
		}
		goto st0
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		if data[p] == 97 {
			goto st150
		}
		goto st0
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		if data[p] == 114 {
			goto st151
		}
		goto st0
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		if data[p] == 121 {
			goto st152
		}
		goto st0
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 32:
			goto tr222
		case 46:
			goto tr224
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr223
		}
		goto st0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 97:
			goto st154
		case 117:
			goto st160
		}
		goto st0
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		if data[p] == 110 {
			goto st155
		}
		goto st0
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 32:
			goto tr233
		case 46:
			goto tr235
		case 117:
			goto st156
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr234
		}
		goto st0
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		if data[p] == 97 {
			goto st157
		}
		goto st0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		if data[p] == 114 {
			goto st158
		}
		goto st0
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		if data[p] == 121 {
			goto st159
		}
		goto st0
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 32:
			goto tr233
		case 46:
			goto tr235
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr234
		}
		goto st0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 108:
			goto st161
		case 110:
			goto st163
		}
		goto st0
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 32:
			goto tr242
		case 46:
			goto tr244
		case 121:
			goto st162
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr243
		}
		goto st0
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 32:
			goto tr242
		case 46:
			goto tr244
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr243
		}
		goto st0
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 32:
			goto tr246
		case 46:
			goto tr248
		case 101:
			goto st164
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr247
		}
		goto st0
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch data[p] {
		case 32:
			goto tr246
		case 46:
			goto tr248
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr247
		}
		goto st0
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		if data[p] == 97 {
			goto st166
		}
		goto st0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 114:
			goto st167
		case 121:
			goto st170
		}
		goto st0
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		switch data[p] {
		case 32:
			goto tr253
		case 46:
			goto tr255
		case 99:
			goto st168
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr254
		}
		goto st0
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		if data[p] == 104 {
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
			goto tr253
		case 46:
			goto tr255
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr254
		}
		goto st0
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 32:
			goto tr258
		case 46:
			goto tr260
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr259
		}
		goto st0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		if data[p] == 111 {
			goto st172
		}
		goto st0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		if data[p] == 118 {
			goto st173
		}
		goto st0
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 32:
			goto tr263
		case 46:
			goto tr265
		case 101:
			goto st174
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr264
		}
		goto st0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		if data[p] == 109 {
			goto st175
		}
		goto st0
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		if data[p] == 98 {
			goto st176
		}
		goto st0
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		if data[p] == 101 {
			goto st177
		}
		goto st0
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		if data[p] == 114 {
			goto st178
		}
		goto st0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 32:
			goto tr263
		case 46:
			goto tr265
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr264
		}
		goto st0
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		if data[p] == 99 {
			goto st180
		}
		goto st0
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		if data[p] == 116 {
			goto st181
		}
		goto st0
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 32:
			goto tr273
		case 46:
			goto tr275
		case 111:
			goto st182
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr274
		}
		goto st0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		if data[p] == 98 {
			goto st183
		}
		goto st0
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		if data[p] == 101 {
			goto st184
		}
		goto st0
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		if data[p] == 114 {
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
			goto tr273
		case 46:
			goto tr275
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr274
		}
		goto st0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		if data[p] == 101 {
			goto st187
		}
		goto st0
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		if data[p] == 112 {
			goto st188
		}
		goto st0
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch data[p] {
		case 32:
			goto tr282
		case 46:
			goto tr284
		case 116:
			goto st189
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr283
		}
		goto st0
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 32:
			goto tr282
		case 46:
			goto tr284
		case 101:
			goto st190
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr283
		}
		goto st0
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		if data[p] == 109 {
			goto st191
		}
		goto st0
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		if data[p] == 98 {
			goto st192
		}
		goto st0
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		if data[p] == 101 {
			goto st193
		}
		goto st0
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		if data[p] == 114 {
			goto st194
		}
		goto st0
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		switch data[p] {
		case 32:
			goto tr282
		case 46:
			goto tr284
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr283
		}
		goto st0
tr178:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st195
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
//line ragel/parse_datetime.go:8461
		switch data[p] {
		case 65:
			goto st196
		case 68:
			goto st207
		case 70:
			goto st215
		case 74:
			goto st223
		case 77:
			goto st235
		case 78:
			goto st241
		case 79:
			goto st249
		case 83:
			goto st256
		case 97:
			goto st196
		case 100:
			goto st207
		case 102:
			goto st215
		case 106:
			goto st223
		case 109:
			goto st235
		case 110:
			goto st241
		case 111:
			goto st249
		case 115:
			goto st256
		}
		goto st0
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 112:
			goto st197
		case 117:
			goto st202
		}
		goto st0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		if data[p] == 114 {
			goto st198
		}
		goto st0
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 32:
			goto tr191
		case 46:
			goto tr302
		case 105:
			goto st200
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr191
		}
		goto st0
tr302:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st199
tr306:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st199
tr312:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st199
tr320:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st199
tr329:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st199
tr336:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st199
tr338:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st199
tr343:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st199
tr346:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st199
tr349:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st199
tr357:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st199
tr364:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
//line ragel/parse_datetime.go:8588
		if data[p] == 32 {
			goto st126
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr197
			}
		case data[p] >= 45:
			goto st126
		}
		goto st0
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		if data[p] == 108 {
			goto st201
		}
		goto st0
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		switch data[p] {
		case 32:
			goto tr191
		case 46:
			goto tr302
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr191
		}
		goto st0
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		if data[p] == 103 {
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
			goto tr205
		case 46:
			goto tr306
		case 117:
			goto st204
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr205
		}
		goto st0
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		if data[p] == 115 {
			goto st205
		}
		goto st0
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		if data[p] == 116 {
			goto st206
		}
		goto st0
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		switch data[p] {
		case 32:
			goto tr205
		case 46:
			goto tr306
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr205
		}
		goto st0
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		if data[p] == 101 {
			goto st208
		}
		goto st0
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		if data[p] == 99 {
			goto st209
		}
		goto st0
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		switch data[p] {
		case 32:
			goto tr213
		case 46:
			goto tr312
		case 101:
			goto st210
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr213
		}
		goto st0
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		if data[p] == 109 {
			goto st211
		}
		goto st0
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		if data[p] == 98 {
			goto st212
		}
		goto st0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		if data[p] == 101 {
			goto st213
		}
		goto st0
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		if data[p] == 114 {
			goto st214
		}
		goto st0
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		switch data[p] {
		case 32:
			goto tr213
		case 46:
			goto tr312
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr213
		}
		goto st0
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		if data[p] == 101 {
			goto st216
		}
		goto st0
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		if data[p] == 98 {
			goto st217
		}
		goto st0
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		switch data[p] {
		case 32:
			goto tr223
		case 46:
			goto tr320
		case 114:
			goto st218
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr223
		}
		goto st0
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		if data[p] == 117 {
			goto st219
		}
		goto st0
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		if data[p] == 97 {
			goto st220
		}
		goto st0
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		if data[p] == 114 {
			goto st221
		}
		goto st0
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		if data[p] == 121 {
			goto st222
		}
		goto st0
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 32:
			goto tr223
		case 46:
			goto tr320
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr223
		}
		goto st0
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 97:
			goto st224
		case 117:
			goto st230
		}
		goto st0
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		if data[p] == 110 {
			goto st225
		}
		goto st0
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 32:
			goto tr234
		case 46:
			goto tr329
		case 117:
			goto st226
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr234
		}
		goto st0
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		if data[p] == 97 {
			goto st227
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
		if data[p] == 121 {
			goto st229
		}
		goto st0
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
		switch data[p] {
		case 32:
			goto tr234
		case 46:
			goto tr329
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr234
		}
		goto st0
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 108:
			goto st231
		case 110:
			goto st233
		}
		goto st0
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		switch data[p] {
		case 32:
			goto tr243
		case 46:
			goto tr336
		case 121:
			goto st232
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr243
		}
		goto st0
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		switch data[p] {
		case 32:
			goto tr243
		case 46:
			goto tr336
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr243
		}
		goto st0
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		switch data[p] {
		case 32:
			goto tr247
		case 46:
			goto tr338
		case 101:
			goto st234
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr247
		}
		goto st0
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		switch data[p] {
		case 32:
			goto tr247
		case 46:
			goto tr338
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr247
		}
		goto st0
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		if data[p] == 97 {
			goto st236
		}
		goto st0
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		switch data[p] {
		case 114:
			goto st237
		case 121:
			goto st240
		}
		goto st0
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		switch data[p] {
		case 32:
			goto tr254
		case 46:
			goto tr343
		case 99:
			goto st238
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr254
		}
		goto st0
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		if data[p] == 104 {
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
			goto tr254
		case 46:
			goto tr343
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr254
		}
		goto st0
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		switch data[p] {
		case 32:
			goto tr259
		case 46:
			goto tr346
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr259
		}
		goto st0
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		if data[p] == 111 {
			goto st242
		}
		goto st0
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		if data[p] == 118 {
			goto st243
		}
		goto st0
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		switch data[p] {
		case 32:
			goto tr264
		case 46:
			goto tr349
		case 101:
			goto st244
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr264
		}
		goto st0
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		if data[p] == 109 {
			goto st245
		}
		goto st0
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		if data[p] == 98 {
			goto st246
		}
		goto st0
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		if data[p] == 101 {
			goto st247
		}
		goto st0
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		if data[p] == 114 {
			goto st248
		}
		goto st0
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		switch data[p] {
		case 32:
			goto tr264
		case 46:
			goto tr349
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr264
		}
		goto st0
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		if data[p] == 99 {
			goto st250
		}
		goto st0
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		if data[p] == 116 {
			goto st251
		}
		goto st0
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		switch data[p] {
		case 32:
			goto tr274
		case 46:
			goto tr357
		case 111:
			goto st252
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr274
		}
		goto st0
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		if data[p] == 98 {
			goto st253
		}
		goto st0
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		if data[p] == 101 {
			goto st254
		}
		goto st0
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		if data[p] == 114 {
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
			goto tr274
		case 46:
			goto tr357
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr274
		}
		goto st0
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		if data[p] == 101 {
			goto st257
		}
		goto st0
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		if data[p] == 112 {
			goto st258
		}
		goto st0
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		switch data[p] {
		case 32:
			goto tr283
		case 46:
			goto tr364
		case 116:
			goto st259
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr283
		}
		goto st0
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		switch data[p] {
		case 32:
			goto tr283
		case 46:
			goto tr364
		case 101:
			goto st260
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr283
		}
		goto st0
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		if data[p] == 109 {
			goto st261
		}
		goto st0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		if data[p] == 98 {
			goto st262
		}
		goto st0
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		if data[p] == 101 {
			goto st263
		}
		goto st0
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		if data[p] == 114 {
			goto st264
		}
		goto st0
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 32:
			goto tr283
		case 46:
			goto tr364
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr283
		}
		goto st0
tr2:
//line ragel/datetime.rl:5
 pb = p 
	goto st265
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
//line ragel/parse_datetime.go:9364
		if data[p] == 32 {
			goto tr177
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st118
			}
		case data[p] >= 45:
			goto tr178
		}
		goto st0
tr3:
//line ragel/datetime.rl:5
 pb = p 
	goto st266
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
//line ragel/parse_datetime.go:9386
		if data[p] == 32 {
			goto tr177
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr178
			}
		case data[p] > 49:
			if 50 <= data[p] && data[p] <= 57 {
				goto st3
			}
		default:
			goto st118
		}
		goto st0
tr4:
//line ragel/datetime.rl:5
 pb = p 
	goto st267
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
//line ragel/parse_datetime.go:9412
		if data[p] == 32 {
			goto tr177
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] >= 45:
			goto tr178
		}
		goto st0
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		switch data[p] {
		case 112:
			goto st269
		case 117:
			goto st361
		}
		goto st0
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		if data[p] == 114 {
			goto st270
		}
		goto st0
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		switch data[p] {
		case 32:
			goto tr374
		case 46:
			goto tr376
		case 105:
			goto st359
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr375
		}
		goto st0
tr374:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st271
tr511:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st271
tr519:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st271
tr530:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st271
tr1081:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st271
tr1089:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st271
tr1092:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st271
tr1099:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st271
tr1103:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st271
tr1107:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st271
tr1116:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st271
tr1128:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st271
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
//line ragel/parse_datetime.go:9516
		switch data[p] {
		case 48:
			goto tr378
		case 51:
			goto tr380
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr381
			}
		case data[p] >= 49:
			goto tr379
		}
		goto st0
tr378:
//line ragel/datetime.rl:5
 pb = p 
	goto st272
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
//line ragel/parse_datetime.go:9541
		if 49 <= data[p] && data[p] <= 57 {
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
			goto tr383
		case 44:
			goto tr384
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr385
		}
		goto st0
tr383:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st274
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
//line ragel/parse_datetime.go:9577
		if data[p] == 50 {
			goto tr387
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr388
			}
		case data[p] >= 48:
			goto tr386
		}
		goto st0
tr386:
//line ragel/datetime.rl:5
 pb = p 
	goto st275
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
//line ragel/parse_datetime.go:9599
		switch data[p] {
		case 32:
			goto tr389
		case 58:
			goto tr391
		case 65:
			goto tr392
		case 80:
			goto tr392
		case 97:
			goto tr393
		case 112:
			goto tr393
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st294
		}
		goto st0
tr389:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st276
tr423:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st276
tr463:
//line ragel/datetime.rl:117

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

	goto st276
tr446:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st276
tr451:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st276
tr457:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st276
tr474:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st276
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
//line ragel/parse_datetime.go:9690
		switch data[p] {
		case 65:
			goto tr395
		case 80:
			goto tr395
		case 97:
			goto tr396
		case 112:
			goto tr396
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr394
		}
		goto st0
tr394:
//line ragel/datetime.rl:5
 pb = p 
	goto st277
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
//line ragel/parse_datetime.go:9714
		if 48 <= data[p] && data[p] <= 57 {
			goto st278
		}
		goto st0
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		if 48 <= data[p] && data[p] <= 57 {
			goto st279
		}
		goto st0
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		if 48 <= data[p] && data[p] <= 57 {
			goto st885
		}
		goto st0
	st885:
		if p++; p == pe {
			goto _test_eof885
		}
	st_case_885:
		if data[p] == 32 {
			goto tr1305
		}
		goto st0
tr1305:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st280
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
//line ragel/parse_datetime.go:9757
		switch data[p] {
		case 43:
			goto st281
		case 45:
			goto st285
		case 47:
			goto tr402
		case 90:
			goto tr403
		case 95:
			goto tr402
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr402
			}
		case data[p] >= 65:
			goto tr402
		}
		goto st0
tr1335:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st281
tr1346:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st281
tr1350:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st281
tr1365:
//line ragel/datetime.rl:117

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

	goto st281
tr1358:
//line ragel/datetime.rl:35

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

	goto st281
tr1378:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st281
tr1387:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st281
tr1395:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st281
tr1415:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st281
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
//line ragel/parse_datetime.go:9893
		if data[p] == 50 {
			goto tr405
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr406
			}
		case data[p] >= 48:
			goto tr404
		}
		goto st0
tr404:
//line ragel/datetime.rl:5
 pb = p 
	goto st886
tr410:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st886
	st886:
		if p++; p == pe {
			goto _test_eof886
		}
	st_case_886:
//line ragel/parse_datetime.go:9921
		switch data[p] {
		case 32:
			goto tr1306
		case 58:
			goto tr1308
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st888
		}
		goto st0
tr1306:
//line ragel/datetime.rl:166

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

	goto st282
tr1313:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

	goto st282
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
//line ragel/parse_datetime.go:9978
		switch data[p] {
		case 47:
			goto tr407
		case 95:
			goto tr407
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr407
			}
		case data[p] >= 65:
			goto tr407
		}
		goto st0
tr407:
//line ragel/datetime.rl:5
 pb = p 
	goto st283
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
//line ragel/parse_datetime.go:10003
		switch data[p] {
		case 47:
			goto st284
		case 95:
			goto st284
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st284
			}
		case data[p] >= 65:
			goto st284
		}
		goto st0
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		switch data[p] {
		case 47:
			goto st887
		case 95:
			goto st887
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st887
			}
		case data[p] >= 65:
			goto st887
		}
		goto st0
	st887:
		if p++; p == pe {
			goto _test_eof887
		}
	st_case_887:
		switch data[p] {
		case 47:
			goto st887
		case 95:
			goto st887
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st887
			}
		case data[p] >= 65:
			goto st887
		}
		goto st0
tr406:
//line ragel/datetime.rl:5
 pb = p 
	goto st888
tr412:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st888
	st888:
		if p++; p == pe {
			goto _test_eof888
		}
	st_case_888:
//line ragel/parse_datetime.go:10074
		switch data[p] {
		case 32:
			goto tr1306
		case 58:
			goto tr1308
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st889
		}
		goto st0
	st889:
		if p++; p == pe {
			goto _test_eof889
		}
	st_case_889:
		if data[p] == 32 {
			goto tr1306
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st889
		}
		goto st0
tr1308:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st890
	st890:
		if p++; p == pe {
			goto _test_eof890
		}
	st_case_890:
//line ragel/parse_datetime.go:10113
		if data[p] == 32 {
			goto st282
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1312
			}
		case data[p] >= 48:
			goto tr1311
		}
		goto st0
tr1311:
//line ragel/datetime.rl:5
 pb = p 
	goto st891
	st891:
		if p++; p == pe {
			goto _test_eof891
		}
	st_case_891:
//line ragel/parse_datetime.go:10135
		if data[p] == 32 {
			goto tr1313
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st892
		}
		goto st0
tr1312:
//line ragel/datetime.rl:5
 pb = p 
	goto st892
	st892:
		if p++; p == pe {
			goto _test_eof892
		}
	st_case_892:
//line ragel/parse_datetime.go:10152
		if data[p] == 32 {
			goto tr1313
		}
		goto st0
tr405:
//line ragel/datetime.rl:5
 pb = p 
	goto st893
tr411:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st893
	st893:
		if p++; p == pe {
			goto _test_eof893
		}
	st_case_893:
//line ragel/parse_datetime.go:10172
		switch data[p] {
		case 32:
			goto tr1306
		case 58:
			goto tr1308
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st889
			}
		case data[p] >= 48:
			goto st888
		}
		goto st0
tr1336:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st285
tr1347:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st285
tr1351:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st285
tr1366:
//line ragel/datetime.rl:117

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

	goto st285
tr1360:
//line ragel/datetime.rl:35

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

	goto st285
tr1379:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st285
tr1388:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st285
tr1397:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st285
tr1417:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st285
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
//line ragel/parse_datetime.go:10302
		if data[p] == 50 {
			goto tr411
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr412
			}
		case data[p] >= 48:
			goto tr410
		}
		goto st0
tr402:
//line ragel/datetime.rl:5
 pb = p 
	goto st286
tr1337:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st286
tr1352:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st286
tr1380:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st286
tr1389:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st286
tr1398:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st286
tr1367:
//line ragel/datetime.rl:117

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
	goto st286
tr1418:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st286
tr1361:
//line ragel/datetime.rl:35

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
	goto st286
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
//line ragel/parse_datetime.go:10424
		switch data[p] {
		case 47:
			goto st287
		case 95:
			goto st287
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st287
			}
		case data[p] >= 65:
			goto st287
		}
		goto st0
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		switch data[p] {
		case 47:
			goto st894
		case 95:
			goto st894
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st894
			}
		case data[p] >= 65:
			goto st894
		}
		goto st0
tr1348:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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
	goto st894
	st894:
		if p++; p == pe {
			goto _test_eof894
		}
	st_case_894:
//line ragel/parse_datetime.go:10492
		switch data[p] {
		case 43:
			goto tr1315
		case 45:
			goto tr1316
		case 47:
			goto st894
		case 95:
			goto st894
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st894
			}
		case data[p] >= 65:
			goto st894
		}
		goto st0
tr1315:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st288
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
//line ragel/parse_datetime.go:10524
		if data[p] == 50 {
			goto tr416
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr417
			}
		case data[p] >= 48:
			goto tr415
		}
		goto st0
tr415:
//line ragel/datetime.rl:5
 pb = p 
	goto st895
tr418:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st895
	st895:
		if p++; p == pe {
			goto _test_eof895
		}
	st_case_895:
//line ragel/parse_datetime.go:10552
		if data[p] == 58 {
			goto tr1318
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st896
		}
		goto st0
tr417:
//line ragel/datetime.rl:5
 pb = p 
	goto st896
tr420:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st896
	st896:
		if p++; p == pe {
			goto _test_eof896
		}
	st_case_896:
//line ragel/parse_datetime.go:10575
		if data[p] == 58 {
			goto tr1318
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st897
		}
		goto st0
	st897:
		if p++; p == pe {
			goto _test_eof897
		}
	st_case_897:
		if 48 <= data[p] && data[p] <= 57 {
			goto st897
		}
		goto st0
tr1318:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st898
	st898:
		if p++; p == pe {
			goto _test_eof898
		}
	st_case_898:
//line ragel/parse_datetime.go:10608
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1321
			}
		case data[p] >= 48:
			goto tr1320
		}
		goto st0
tr1320:
//line ragel/datetime.rl:5
 pb = p 
	goto st899
	st899:
		if p++; p == pe {
			goto _test_eof899
		}
	st_case_899:
//line ragel/parse_datetime.go:10627
		if 48 <= data[p] && data[p] <= 57 {
			goto st900
		}
		goto st0
tr1321:
//line ragel/datetime.rl:5
 pb = p 
	goto st900
	st900:
		if p++; p == pe {
			goto _test_eof900
		}
	st_case_900:
//line ragel/parse_datetime.go:10641
		goto st0
tr416:
//line ragel/datetime.rl:5
 pb = p 
	goto st901
tr419:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st901
	st901:
		if p++; p == pe {
			goto _test_eof901
		}
	st_case_901:
//line ragel/parse_datetime.go:10658
		if data[p] == 58 {
			goto tr1318
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st897
			}
		case data[p] >= 48:
			goto st896
		}
		goto st0
tr1316:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st289
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
//line ragel/parse_datetime.go:10683
		if data[p] == 50 {
			goto tr419
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr420
			}
		case data[p] >= 48:
			goto tr418
		}
		goto st0
tr403:
//line ragel/datetime.rl:5
 pb = p 
	goto st902
tr1341:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st902
tr1355:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st902
tr1384:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st902
tr1392:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st902
tr1401:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st902
tr1369:
//line ragel/datetime.rl:117

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
	goto st902
tr1420:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st902
tr1363:
//line ragel/datetime.rl:35

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
	goto st902
	st902:
		if p++; p == pe {
			goto _test_eof902
		}
	st_case_902:
//line ragel/parse_datetime.go:10805
		switch data[p] {
		case 47:
			goto st287
		case 95:
			goto st287
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st287
			}
		case data[p] >= 65:
			goto st287
		}
		goto st0
tr395:
//line ragel/datetime.rl:5
 pb = p 
	goto st290
tr392:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st290
tr426:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st290
tr449:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st290
tr453:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st290
tr460:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st290
tr465:
//line ragel/datetime.rl:117

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
	goto st290
tr476:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st290
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
//line ragel/parse_datetime.go:10911
		if data[p] == 77 {
			goto st291
		}
		goto st0
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		if data[p] == 32 {
			goto tr422
		}
		goto st0
tr422:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st292
tr431:
//line ragel/datetime.rl:117

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

	goto st292
tr442:
//line ragel/datetime.rl:35

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

	goto st292
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
//line ragel/parse_datetime.go:11003
		if 48 <= data[p] && data[p] <= 57 {
			goto tr394
		}
		goto st0
tr396:
//line ragel/datetime.rl:5
 pb = p 
	goto st293
tr393:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st293
tr427:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st293
tr450:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st293
tr454:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st293
tr461:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st293
tr466:
//line ragel/datetime.rl:117

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
	goto st293
tr477:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st293
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
//line ragel/parse_datetime.go:11098
		if data[p] == 109 {
			goto st291
		}
		goto st0
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		switch data[p] {
		case 32:
			goto tr423
		case 58:
			goto tr425
		case 65:
			goto tr426
		case 80:
			goto tr426
		case 97:
			goto tr427
		case 112:
			goto tr427
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st295
		}
		goto st0
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		if 48 <= data[p] && data[p] <= 57 {
			goto st903
		}
		goto st0
	st903:
		if p++; p == pe {
			goto _test_eof903
		}
	st_case_903:
		switch data[p] {
		case 32:
			goto tr1323
		case 43:
			goto tr1298
		case 44:
			goto tr1324
		case 45:
			goto tr1300
		case 46:
			goto tr443
		case 47:
			goto tr1301
		case 84:
			goto tr1302
		case 90:
			goto tr1303
		case 95:
			goto tr1304
		case 116:
			goto tr1304
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st308
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1301
			}
		default:
			goto tr1301
		}
		goto st0
tr1323:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:35

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

	goto st904
	st904:
		if p++; p == pe {
			goto _test_eof904
		}
	st_case_904:
//line ragel/parse_datetime.go:11201
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr1155
		case 50:
			goto tr1327
		case 65:
			goto tr1156
		case 66:
			goto tr1157
		case 90:
			goto tr1158
		case 95:
			goto tr1155
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr1326
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1155
				}
			case data[p] >= 67:
				goto tr1155
			}
		default:
			goto tr1328
		}
		goto st0
tr1326:
//line ragel/datetime.rl:5
 pb = p 
	goto st905
	st905:
		if p++; p == pe {
			goto _test_eof905
		}
	st_case_905:
//line ragel/parse_datetime.go:11249
		switch data[p] {
		case 32:
			goto tr1180
		case 43:
			goto tr1181
		case 45:
			goto tr1182
		case 47:
			goto tr1183
		case 58:
			goto tr1185
		case 65:
			goto tr1186
		case 80:
			goto tr1186
		case 90:
			goto tr1187
		case 95:
			goto tr1183
		case 97:
			goto tr1188
		case 112:
			goto tr1188
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st906
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1183
			}
		default:
			goto tr1183
		}
		goto st0
	st906:
		if p++; p == pe {
			goto _test_eof906
		}
	st_case_906:
		switch data[p] {
		case 32:
			goto tr1196
		case 43:
			goto tr1197
		case 45:
			goto tr1198
		case 47:
			goto tr1199
		case 58:
			goto tr1200
		case 65:
			goto tr1201
		case 80:
			goto tr1201
		case 90:
			goto tr1202
		case 95:
			goto tr1199
		case 97:
			goto tr1203
		case 112:
			goto tr1203
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st296
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1199
			}
		default:
			goto tr1199
		}
		goto st0
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		if 48 <= data[p] && data[p] <= 57 {
			goto st907
		}
		goto st0
	st907:
		if p++; p == pe {
			goto _test_eof907
		}
	st_case_907:
		switch data[p] {
		case 32:
			goto tr1330
		case 43:
			goto tr1205
		case 45:
			goto tr1207
		case 47:
			goto tr1208
		case 90:
			goto tr1210
		case 95:
			goto tr1208
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1206
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1208
				}
			case data[p] >= 65:
				goto tr1208
			}
		default:
			goto st32
		}
		goto st0
tr1327:
//line ragel/datetime.rl:5
 pb = p 
	goto st908
	st908:
		if p++; p == pe {
			goto _test_eof908
		}
	st_case_908:
//line ragel/parse_datetime.go:11384
		switch data[p] {
		case 32:
			goto tr1180
		case 43:
			goto tr1181
		case 45:
			goto tr1182
		case 47:
			goto tr1183
		case 58:
			goto tr1185
		case 65:
			goto tr1186
		case 80:
			goto tr1186
		case 90:
			goto tr1187
		case 95:
			goto tr1183
		case 97:
			goto tr1188
		case 112:
			goto tr1188
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st906
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1183
				}
			case data[p] >= 66:
				goto tr1183
			}
		default:
			goto st297
		}
		goto st0
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		if 48 <= data[p] && data[p] <= 57 {
			goto st296
		}
		goto st0
tr1328:
//line ragel/datetime.rl:5
 pb = p 
	goto st909
	st909:
		if p++; p == pe {
			goto _test_eof909
		}
	st_case_909:
//line ragel/parse_datetime.go:11445
		switch data[p] {
		case 32:
			goto tr1180
		case 43:
			goto tr1181
		case 45:
			goto tr1182
		case 47:
			goto tr1183
		case 58:
			goto tr1185
		case 65:
			goto tr1186
		case 80:
			goto tr1186
		case 90:
			goto tr1187
		case 95:
			goto tr1183
		case 97:
			goto tr1188
		case 112:
			goto tr1188
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st297
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1183
			}
		default:
			goto tr1183
		}
		goto st0
tr1324:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:35

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

	goto st910
	st910:
		if p++; p == pe {
			goto _test_eof910
		}
	st_case_910:
//line ragel/parse_datetime.go:11509
		if data[p] == 32 {
			goto st38
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr440
		}
		goto st0
tr440:
//line ragel/datetime.rl:5
 pb = p 
	goto st298
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
//line ragel/parse_datetime.go:11526
		if data[p] == 32 {
			goto tr431
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st299
		}
		goto st0
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		if data[p] == 32 {
			goto tr431
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st300
		}
		goto st0
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		if data[p] == 32 {
			goto tr431
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st301
		}
		goto st0
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		if data[p] == 32 {
			goto tr431
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st302
		}
		goto st0
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		if data[p] == 32 {
			goto tr431
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st303
		}
		goto st0
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		if data[p] == 32 {
			goto tr431
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st304
		}
		goto st0
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		if data[p] == 32 {
			goto tr431
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st305
		}
		goto st0
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		if data[p] == 32 {
			goto tr431
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st306
		}
		goto st0
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		if data[p] == 32 {
			goto tr431
		}
		goto st0
tr443:
//line ragel/datetime.rl:35

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

	goto st307
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
//line ragel/parse_datetime.go:11649
		if 48 <= data[p] && data[p] <= 57 {
			goto tr440
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
		switch data[p] {
		case 32:
			goto tr442
		case 44:
			goto tr443
		case 46:
			goto tr443
		}
		goto st0
tr391:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st310
tr425:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st310
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
//line ragel/parse_datetime.go:11694
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr445
			}
		case data[p] >= 48:
			goto tr444
		}
		goto st0
tr444:
//line ragel/datetime.rl:5
 pb = p 
	goto st311
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
//line ragel/parse_datetime.go:11713
		switch data[p] {
		case 32:
			goto tr446
		case 58:
			goto tr448
		case 65:
			goto tr449
		case 80:
			goto tr449
		case 97:
			goto tr450
		case 112:
			goto tr450
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st312
		}
		goto st0
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		switch data[p] {
		case 32:
			goto tr451
		case 58:
			goto tr452
		case 65:
			goto tr453
		case 80:
			goto tr453
		case 97:
			goto tr454
		case 112:
			goto tr454
		}
		goto st0
tr448:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st313
tr452:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st313
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
//line ragel/parse_datetime.go:11769
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr456
			}
		case data[p] >= 48:
			goto tr455
		}
		goto st0
tr455:
//line ragel/datetime.rl:5
 pb = p 
	goto st314
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
//line ragel/parse_datetime.go:11788
		switch data[p] {
		case 32:
			goto tr457
		case 44:
			goto tr458
		case 46:
			goto tr458
		case 65:
			goto tr460
		case 80:
			goto tr460
		case 97:
			goto tr461
		case 112:
			goto tr461
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st325
		}
		goto st0
tr458:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st315
tr475:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st315
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
//line ragel/parse_datetime.go:11826
		if 48 <= data[p] && data[p] <= 57 {
			goto tr462
		}
		goto st0
tr462:
//line ragel/datetime.rl:5
 pb = p 
	goto st316
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
//line ragel/parse_datetime.go:11840
		switch data[p] {
		case 32:
			goto tr463
		case 65:
			goto tr465
		case 80:
			goto tr465
		case 97:
			goto tr466
		case 112:
			goto tr466
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st317
		}
		goto st0
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		switch data[p] {
		case 32:
			goto tr463
		case 65:
			goto tr465
		case 80:
			goto tr465
		case 97:
			goto tr466
		case 112:
			goto tr466
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st318
		}
		goto st0
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		switch data[p] {
		case 32:
			goto tr463
		case 65:
			goto tr465
		case 80:
			goto tr465
		case 97:
			goto tr466
		case 112:
			goto tr466
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st319
		}
		goto st0
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		switch data[p] {
		case 32:
			goto tr463
		case 65:
			goto tr465
		case 80:
			goto tr465
		case 97:
			goto tr466
		case 112:
			goto tr466
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st320
		}
		goto st0
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		switch data[p] {
		case 32:
			goto tr463
		case 65:
			goto tr465
		case 80:
			goto tr465
		case 97:
			goto tr466
		case 112:
			goto tr466
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st321
		}
		goto st0
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		switch data[p] {
		case 32:
			goto tr463
		case 65:
			goto tr465
		case 80:
			goto tr465
		case 97:
			goto tr466
		case 112:
			goto tr466
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st322
		}
		goto st0
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		switch data[p] {
		case 32:
			goto tr463
		case 65:
			goto tr465
		case 80:
			goto tr465
		case 97:
			goto tr466
		case 112:
			goto tr466
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st323
		}
		goto st0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 32:
			goto tr463
		case 65:
			goto tr465
		case 80:
			goto tr465
		case 97:
			goto tr466
		case 112:
			goto tr466
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st324
		}
		goto st0
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		switch data[p] {
		case 32:
			goto tr463
		case 65:
			goto tr465
		case 80:
			goto tr465
		case 97:
			goto tr466
		case 112:
			goto tr466
		}
		goto st0
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
		switch data[p] {
		case 32:
			goto tr474
		case 44:
			goto tr475
		case 46:
			goto tr475
		case 65:
			goto tr476
		case 80:
			goto tr476
		case 97:
			goto tr477
		case 112:
			goto tr477
		}
		goto st0
tr456:
//line ragel/datetime.rl:5
 pb = p 
	goto st326
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
//line ragel/parse_datetime.go:12053
		switch data[p] {
		case 32:
			goto tr457
		case 44:
			goto tr458
		case 46:
			goto tr458
		case 65:
			goto tr460
		case 80:
			goto tr460
		case 97:
			goto tr461
		case 112:
			goto tr461
		}
		goto st0
tr445:
//line ragel/datetime.rl:5
 pb = p 
	goto st327
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
//line ragel/parse_datetime.go:12080
		switch data[p] {
		case 32:
			goto tr446
		case 58:
			goto tr448
		case 65:
			goto tr449
		case 80:
			goto tr449
		case 97:
			goto tr450
		case 112:
			goto tr450
		}
		goto st0
tr387:
//line ragel/datetime.rl:5
 pb = p 
	goto st328
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
//line ragel/parse_datetime.go:12105
		switch data[p] {
		case 32:
			goto tr389
		case 58:
			goto tr391
		case 65:
			goto tr392
		case 80:
			goto tr392
		case 97:
			goto tr393
		case 112:
			goto tr393
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st329
			}
		case data[p] >= 48:
			goto st294
		}
		goto st0
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		if 48 <= data[p] && data[p] <= 57 {
			goto st295
		}
		goto st0
tr388:
//line ragel/datetime.rl:5
 pb = p 
	goto st330
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
//line ragel/parse_datetime.go:12147
		switch data[p] {
		case 32:
			goto tr389
		case 58:
			goto tr391
		case 65:
			goto tr392
		case 80:
			goto tr392
		case 97:
			goto tr393
		case 112:
			goto tr393
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st329
		}
		goto st0
tr384:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st331
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
//line ragel/parse_datetime.go:12182
		if data[p] == 32 {
			goto st332
		}
		goto st0
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr480
		}
		goto st0
tr480:
//line ragel/datetime.rl:5
 pb = p 
	goto st333
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
//line ragel/parse_datetime.go:12205
		if 48 <= data[p] && data[p] <= 57 {
			goto st334
		}
		goto st0
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		if 48 <= data[p] && data[p] <= 57 {
			goto st335
		}
		goto st0
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		if 48 <= data[p] && data[p] <= 57 {
			goto st911
		}
		goto st0
	st911:
		if p++; p == pe {
			goto _test_eof911
		}
	st_case_911:
		switch data[p] {
		case 32:
			goto tr1332
		case 44:
			goto tr1333
		}
		goto st0
tr1332:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st336
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
//line ragel/parse_datetime.go:12251
		if data[p] == 50 {
			goto tr485
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr486
			}
		case data[p] >= 48:
			goto tr484
		}
		goto st0
tr484:
//line ragel/datetime.rl:5
 pb = p 
	goto st912
	st912:
		if p++; p == pe {
			goto _test_eof912
		}
	st_case_912:
//line ragel/parse_datetime.go:12273
		switch data[p] {
		case 32:
			goto tr1334
		case 43:
			goto tr1335
		case 45:
			goto tr1336
		case 47:
			goto tr1337
		case 58:
			goto tr1339
		case 65:
			goto tr1340
		case 80:
			goto tr1340
		case 90:
			goto tr1341
		case 95:
			goto tr1337
		case 97:
			goto tr1342
		case 112:
			goto tr1342
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st916
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1337
			}
		default:
			goto tr1337
		}
		goto st0
tr1334:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st913
tr1349:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st913
tr1403:
//line ragel/datetime.rl:117

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

	goto st913
tr1377:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st913
tr1386:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st913
tr1394:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st913
tr1414:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st913
	st913:
		if p++; p == pe {
			goto _test_eof913
		}
	st_case_913:
//line ragel/parse_datetime.go:12383
		switch data[p] {
		case 43:
			goto st281
		case 45:
			goto st285
		case 47:
			goto tr402
		case 65:
			goto tr1343
		case 80:
			goto tr1343
		case 90:
			goto tr403
		case 95:
			goto tr402
		case 97:
			goto tr1344
		case 112:
			goto tr1344
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr402
			}
		case data[p] >= 66:
			goto tr402
		}
		goto st0
tr1343:
//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1340:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1354:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1383:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1391:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1400:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
tr1405:
//line ragel/datetime.rl:117

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
tr1419:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st337
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
//line ragel/parse_datetime.go:12503
		switch data[p] {
		case 47:
			goto st287
		case 77:
			goto st914
		case 95:
			goto st287
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st287
			}
		case data[p] >= 65:
			goto st287
		}
		goto st0
	st914:
		if p++; p == pe {
			goto _test_eof914
		}
	st_case_914:
		switch data[p] {
		case 32:
			goto tr1345
		case 43:
			goto tr1346
		case 45:
			goto tr1347
		case 47:
			goto tr1348
		case 95:
			goto tr1348
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1348
			}
		case data[p] >= 65:
			goto tr1348
		}
		goto st0
tr1345:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st915
tr1364:
//line ragel/datetime.rl:117

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

	goto st915
tr1357:
//line ragel/datetime.rl:35

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

	goto st915
	st915:
		if p++; p == pe {
			goto _test_eof915
		}
	st_case_915:
//line ragel/parse_datetime.go:12625
		switch data[p] {
		case 43:
			goto st281
		case 45:
			goto st285
		case 47:
			goto tr402
		case 90:
			goto tr403
		case 95:
			goto tr402
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr402
			}
		case data[p] >= 65:
			goto tr402
		}
		goto st0
tr1344:
//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1342:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1356:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1385:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1393:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1402:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
tr1406:
//line ragel/datetime.rl:117

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
tr1421:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st338
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
//line ragel/parse_datetime.go:12737
		switch data[p] {
		case 47:
			goto st287
		case 95:
			goto st287
		case 109:
			goto st914
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st287
			}
		case data[p] >= 65:
			goto st287
		}
		goto st0
	st916:
		if p++; p == pe {
			goto _test_eof916
		}
	st_case_916:
		switch data[p] {
		case 32:
			goto tr1349
		case 43:
			goto tr1350
		case 45:
			goto tr1351
		case 47:
			goto tr1352
		case 58:
			goto tr1353
		case 65:
			goto tr1354
		case 80:
			goto tr1354
		case 90:
			goto tr1355
		case 95:
			goto tr1352
		case 97:
			goto tr1356
		case 112:
			goto tr1356
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st339
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1352
			}
		default:
			goto tr1352
		}
		goto st0
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		if 48 <= data[p] && data[p] <= 57 {
			goto st917
		}
		goto st0
	st917:
		if p++; p == pe {
			goto _test_eof917
		}
	st_case_917:
		switch data[p] {
		case 32:
			goto tr1357
		case 43:
			goto tr1358
		case 45:
			goto tr1360
		case 47:
			goto tr1361
		case 90:
			goto tr1363
		case 95:
			goto tr1361
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1359
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1361
				}
			case data[p] >= 65:
				goto tr1361
			}
		default:
			goto st341
		}
		goto st0
tr1359:
//line ragel/datetime.rl:35

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
//line ragel/parse_datetime.go:12865
		if 48 <= data[p] && data[p] <= 57 {
			goto tr489
		}
		goto st0
tr489:
//line ragel/datetime.rl:5
 pb = p 
	goto st918
	st918:
		if p++; p == pe {
			goto _test_eof918
		}
	st_case_918:
//line ragel/parse_datetime.go:12879
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st919
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st919:
		if p++; p == pe {
			goto _test_eof919
		}
	st_case_919:
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st920
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st920:
		if p++; p == pe {
			goto _test_eof920
		}
	st_case_920:
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st921
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st921:
		if p++; p == pe {
			goto _test_eof921
		}
	st_case_921:
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st922
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st922:
		if p++; p == pe {
			goto _test_eof922
		}
	st_case_922:
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st923
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st923:
		if p++; p == pe {
			goto _test_eof923
		}
	st_case_923:
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st924
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st924:
		if p++; p == pe {
			goto _test_eof924
		}
	st_case_924:
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st925
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st925:
		if p++; p == pe {
			goto _test_eof925
		}
	st_case_925:
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st926
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st926:
		if p++; p == pe {
			goto _test_eof926
		}
	st_case_926:
		switch data[p] {
		case 32:
			goto tr1364
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		case data[p] >= 65:
			goto tr1367
		}
		goto st0
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		if 48 <= data[p] && data[p] <= 57 {
			goto st927
		}
		goto st0
	st927:
		if p++; p == pe {
			goto _test_eof927
		}
	st_case_927:
		switch data[p] {
		case 32:
			goto tr1357
		case 43:
			goto tr1358
		case 45:
			goto tr1360
		case 47:
			goto tr1361
		case 90:
			goto tr1363
		case 95:
			goto tr1361
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1359
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1361
			}
		default:
			goto tr1361
		}
		goto st0
tr1339:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st342
tr1353:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st342
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
//line ragel/parse_datetime.go:13217
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr492
			}
		case data[p] >= 48:
			goto tr491
		}
		goto st0
tr491:
//line ragel/datetime.rl:5
 pb = p 
	goto st928
	st928:
		if p++; p == pe {
			goto _test_eof928
		}
	st_case_928:
//line ragel/parse_datetime.go:13236
		switch data[p] {
		case 32:
			goto tr1377
		case 43:
			goto tr1378
		case 45:
			goto tr1379
		case 47:
			goto tr1380
		case 58:
			goto tr1382
		case 65:
			goto tr1383
		case 80:
			goto tr1383
		case 90:
			goto tr1384
		case 95:
			goto tr1380
		case 97:
			goto tr1385
		case 112:
			goto tr1385
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st929
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1380
			}
		default:
			goto tr1380
		}
		goto st0
	st929:
		if p++; p == pe {
			goto _test_eof929
		}
	st_case_929:
		switch data[p] {
		case 32:
			goto tr1386
		case 43:
			goto tr1387
		case 45:
			goto tr1388
		case 47:
			goto tr1389
		case 58:
			goto tr1390
		case 65:
			goto tr1391
		case 80:
			goto tr1391
		case 90:
			goto tr1392
		case 95:
			goto tr1389
		case 97:
			goto tr1393
		case 112:
			goto tr1393
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1389
			}
		case data[p] >= 66:
			goto tr1389
		}
		goto st0
tr1382:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st343
tr1390:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st343
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
//line ragel/parse_datetime.go:13329
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr494
			}
		case data[p] >= 48:
			goto tr493
		}
		goto st0
tr493:
//line ragel/datetime.rl:5
 pb = p 
	goto st930
	st930:
		if p++; p == pe {
			goto _test_eof930
		}
	st_case_930:
//line ragel/parse_datetime.go:13348
		switch data[p] {
		case 32:
			goto tr1394
		case 43:
			goto tr1395
		case 45:
			goto tr1397
		case 47:
			goto tr1398
		case 65:
			goto tr1400
		case 80:
			goto tr1400
		case 90:
			goto tr1401
		case 95:
			goto tr1398
		case 97:
			goto tr1402
		case 112:
			goto tr1402
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1396
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1398
				}
			case data[p] >= 66:
				goto tr1398
			}
		default:
			goto st940
		}
		goto st0
tr1396:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st344
tr1416:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st344
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
//line ragel/parse_datetime.go:13406
		if 48 <= data[p] && data[p] <= 57 {
			goto tr495
		}
		goto st0
tr495:
//line ragel/datetime.rl:5
 pb = p 
	goto st931
	st931:
		if p++; p == pe {
			goto _test_eof931
		}
	st_case_931:
//line ragel/parse_datetime.go:13420
		switch data[p] {
		case 32:
			goto tr1403
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		case 97:
			goto tr1406
		case 112:
			goto tr1406
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st932
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st932:
		if p++; p == pe {
			goto _test_eof932
		}
	st_case_932:
		switch data[p] {
		case 32:
			goto tr1403
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		case 97:
			goto tr1406
		case 112:
			goto tr1406
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st933
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st933:
		if p++; p == pe {
			goto _test_eof933
		}
	st_case_933:
		switch data[p] {
		case 32:
			goto tr1403
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		case 97:
			goto tr1406
		case 112:
			goto tr1406
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st934
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st934:
		if p++; p == pe {
			goto _test_eof934
		}
	st_case_934:
		switch data[p] {
		case 32:
			goto tr1403
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		case 97:
			goto tr1406
		case 112:
			goto tr1406
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st935
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st935:
		if p++; p == pe {
			goto _test_eof935
		}
	st_case_935:
		switch data[p] {
		case 32:
			goto tr1403
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		case 97:
			goto tr1406
		case 112:
			goto tr1406
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st936
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st936:
		if p++; p == pe {
			goto _test_eof936
		}
	st_case_936:
		switch data[p] {
		case 32:
			goto tr1403
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		case 97:
			goto tr1406
		case 112:
			goto tr1406
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st937
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st937:
		if p++; p == pe {
			goto _test_eof937
		}
	st_case_937:
		switch data[p] {
		case 32:
			goto tr1403
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		case 97:
			goto tr1406
		case 112:
			goto tr1406
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st938
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st938:
		if p++; p == pe {
			goto _test_eof938
		}
	st_case_938:
		switch data[p] {
		case 32:
			goto tr1403
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		case 97:
			goto tr1406
		case 112:
			goto tr1406
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st939
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		default:
			goto tr1367
		}
		goto st0
	st939:
		if p++; p == pe {
			goto _test_eof939
		}
	st_case_939:
		switch data[p] {
		case 32:
			goto tr1403
		case 43:
			goto tr1365
		case 45:
			goto tr1366
		case 47:
			goto tr1367
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1369
		case 95:
			goto tr1367
		case 97:
			goto tr1406
		case 112:
			goto tr1406
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		case data[p] >= 66:
			goto tr1367
		}
		goto st0
	st940:
		if p++; p == pe {
			goto _test_eof940
		}
	st_case_940:
		switch data[p] {
		case 32:
			goto tr1414
		case 43:
			goto tr1415
		case 45:
			goto tr1417
		case 47:
			goto tr1418
		case 65:
			goto tr1419
		case 80:
			goto tr1419
		case 90:
			goto tr1420
		case 95:
			goto tr1418
		case 97:
			goto tr1421
		case 112:
			goto tr1421
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1416
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1418
			}
		default:
			goto tr1418
		}
		goto st0
tr494:
//line ragel/datetime.rl:5
 pb = p 
	goto st941
	st941:
		if p++; p == pe {
			goto _test_eof941
		}
	st_case_941:
//line ragel/parse_datetime.go:13821
		switch data[p] {
		case 32:
			goto tr1394
		case 43:
			goto tr1395
		case 45:
			goto tr1397
		case 47:
			goto tr1398
		case 65:
			goto tr1400
		case 80:
			goto tr1400
		case 90:
			goto tr1401
		case 95:
			goto tr1398
		case 97:
			goto tr1402
		case 112:
			goto tr1402
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1396
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1398
			}
		default:
			goto tr1398
		}
		goto st0
tr492:
//line ragel/datetime.rl:5
 pb = p 
	goto st942
	st942:
		if p++; p == pe {
			goto _test_eof942
		}
	st_case_942:
//line ragel/parse_datetime.go:13866
		switch data[p] {
		case 32:
			goto tr1377
		case 43:
			goto tr1378
		case 45:
			goto tr1379
		case 47:
			goto tr1380
		case 58:
			goto tr1382
		case 65:
			goto tr1383
		case 80:
			goto tr1383
		case 90:
			goto tr1384
		case 95:
			goto tr1380
		case 97:
			goto tr1385
		case 112:
			goto tr1385
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1380
			}
		case data[p] >= 66:
			goto tr1380
		}
		goto st0
tr485:
//line ragel/datetime.rl:5
 pb = p 
	goto st943
	st943:
		if p++; p == pe {
			goto _test_eof943
		}
	st_case_943:
//line ragel/parse_datetime.go:13909
		switch data[p] {
		case 32:
			goto tr1334
		case 43:
			goto tr1335
		case 45:
			goto tr1336
		case 47:
			goto tr1337
		case 58:
			goto tr1339
		case 65:
			goto tr1340
		case 80:
			goto tr1340
		case 90:
			goto tr1341
		case 95:
			goto tr1337
		case 97:
			goto tr1342
		case 112:
			goto tr1342
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st916
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1337
				}
			case data[p] >= 66:
				goto tr1337
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
tr486:
//line ragel/datetime.rl:5
 pb = p 
	goto st944
	st944:
		if p++; p == pe {
			goto _test_eof944
		}
	st_case_944:
//line ragel/parse_datetime.go:13970
		switch data[p] {
		case 32:
			goto tr1334
		case 43:
			goto tr1335
		case 45:
			goto tr1336
		case 47:
			goto tr1337
		case 58:
			goto tr1339
		case 65:
			goto tr1340
		case 80:
			goto tr1340
		case 90:
			goto tr1341
		case 95:
			goto tr1337
		case 97:
			goto tr1342
		case 112:
			goto tr1342
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st345
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1337
			}
		default:
			goto tr1337
		}
		goto st0
tr1333:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st945
	st945:
		if p++; p == pe {
			goto _test_eof945
		}
	st_case_945:
//line ragel/parse_datetime.go:14019
		if data[p] == 32 {
			goto st336
		}
		goto st0
tr379:
//line ragel/datetime.rl:5
 pb = p 
	goto st346
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
//line ragel/parse_datetime.go:14033
		switch data[p] {
		case 32:
			goto tr497
		case 44:
			goto tr384
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st273
			}
		case data[p] >= 45:
			goto tr385
		}
		goto st0
tr497:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st347
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
//line ragel/parse_datetime.go:14065
		if 48 <= data[p] && data[p] <= 57 {
			goto tr498
		}
		goto st0
tr498:
//line ragel/datetime.rl:5
 pb = p 
	goto st348
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
//line ragel/parse_datetime.go:14079
		if 48 <= data[p] && data[p] <= 57 {
			goto st349
		}
		goto st0
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		if 48 <= data[p] && data[p] <= 57 {
			goto st350
		}
		goto st0
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		if 48 <= data[p] && data[p] <= 57 {
			goto st946
		}
		goto st0
	st946:
		if p++; p == pe {
			goto _test_eof946
		}
	st_case_946:
		switch data[p] {
		case 32:
			goto tr1297
		case 43:
			goto tr1298
		case 44:
			goto tr1424
		case 45:
			goto tr1300
		case 47:
			goto tr1301
		case 84:
			goto tr1302
		case 90:
			goto tr1303
		case 95:
			goto tr1304
		case 116:
			goto tr1304
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1301
			}
		case data[p] >= 65:
			goto tr1301
		}
		goto st0
tr1424:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st947
	st947:
		if p++; p == pe {
			goto _test_eof947
		}
	st_case_947:
//line ragel/parse_datetime.go:14147
		if data[p] == 32 {
			goto st38
		}
		goto st0
tr380:
//line ragel/datetime.rl:5
 pb = p 
	goto st351
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
//line ragel/parse_datetime.go:14161
		switch data[p] {
		case 32:
			goto tr497
		case 44:
			goto tr384
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st273
			}
		case data[p] >= 45:
			goto tr385
		}
		goto st0
tr381:
//line ragel/datetime.rl:5
 pb = p 
	goto st352
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
//line ragel/parse_datetime.go:14186
		switch data[p] {
		case 32:
			goto tr497
		case 44:
			goto tr384
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr385
		}
		goto st0
tr375:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st353
tr512:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st353
tr520:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st353
tr531:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st353
tr852:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st353
tr861:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st353
tr865:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st353
tr872:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st353
tr877:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st353
tr882:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st353
tr892:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st353
tr901:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st353
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
//line ragel/parse_datetime.go:14250
		switch data[p] {
		case 48:
			goto tr502
		case 51:
			goto tr504
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr505
			}
		case data[p] >= 49:
			goto tr503
		}
		goto st0
tr502:
//line ragel/datetime.rl:5
 pb = p 
	goto st354
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
//line ragel/parse_datetime.go:14275
		if 49 <= data[p] && data[p] <= 57 {
			goto st355
		}
		goto st0
tr505:
//line ragel/datetime.rl:5
 pb = p 
	goto st355
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
//line ragel/parse_datetime.go:14289
		if data[p] == 32 {
			goto tr385
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr385
		}
		goto st0
tr503:
//line ragel/datetime.rl:5
 pb = p 
	goto st356
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
//line ragel/parse_datetime.go:14306
		if data[p] == 32 {
			goto tr385
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st355
			}
		case data[p] >= 45:
			goto tr385
		}
		goto st0
tr504:
//line ragel/datetime.rl:5
 pb = p 
	goto st357
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
//line ragel/parse_datetime.go:14328
		if data[p] == 32 {
			goto tr385
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st355
			}
		case data[p] >= 45:
			goto tr385
		}
		goto st0
tr376:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st358
tr513:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st358
tr521:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st358
tr532:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st358
tr1082:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st358
tr1090:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st358
tr1093:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st358
tr1100:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st358
tr1104:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st358
tr1108:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st358
tr1117:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st358
tr1129:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st358
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
//line ragel/parse_datetime.go:14394
		switch data[p] {
		case 32:
			goto st271
		case 48:
			goto tr502
		case 51:
			goto tr504
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st353
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr505
			}
		default:
			goto tr503
		}
		goto st0
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		if data[p] == 108 {
			goto st360
		}
		goto st0
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		switch data[p] {
		case 32:
			goto tr374
		case 46:
			goto tr376
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr375
		}
		goto st0
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		if data[p] == 103 {
			goto st362
		}
		goto st0
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		switch data[p] {
		case 32:
			goto tr511
		case 46:
			goto tr513
		case 117:
			goto st363
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr512
		}
		goto st0
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		if data[p] == 115 {
			goto st364
		}
		goto st0
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		if data[p] == 116 {
			goto st365
		}
		goto st0
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		switch data[p] {
		case 32:
			goto tr511
		case 46:
			goto tr513
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr512
		}
		goto st0
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		if data[p] == 101 {
			goto st367
		}
		goto st0
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		if data[p] == 99 {
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
			goto tr519
		case 46:
			goto tr521
		case 101:
			goto st369
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr520
		}
		goto st0
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
		if data[p] == 109 {
			goto st370
		}
		goto st0
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		if data[p] == 98 {
			goto st371
		}
		goto st0
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		if data[p] == 101 {
			goto st372
		}
		goto st0
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		if data[p] == 114 {
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
			goto tr519
		case 46:
			goto tr521
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr520
		}
		goto st0
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		switch data[p] {
		case 101:
			goto st375
		case 114:
			goto st382
		}
		goto st0
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		if data[p] == 98 {
			goto st376
		}
		goto st0
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		switch data[p] {
		case 32:
			goto tr530
		case 46:
			goto tr532
		case 114:
			goto st377
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr531
		}
		goto st0
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		if data[p] == 117 {
			goto st378
		}
		goto st0
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		if data[p] == 97 {
			goto st379
		}
		goto st0
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		if data[p] == 114 {
			goto st380
		}
		goto st0
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		if data[p] == 121 {
			goto st381
		}
		goto st0
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		switch data[p] {
		case 32:
			goto tr530
		case 46:
			goto tr532
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr531
		}
		goto st0
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		if data[p] == 105 {
			goto st383
		}
		goto st0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 32:
			goto st384
		case 44:
			goto st610
		case 100:
			goto st755
		}
		goto st0
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		switch data[p] {
		case 48:
			goto tr542
		case 51:
			goto tr544
		case 65:
			goto st389
		case 68:
			goto st552
		case 70:
			goto st560
		case 74:
			goto st568
		case 77:
			goto st580
		case 78:
			goto st586
		case 79:
			goto st594
		case 83:
			goto st601
		case 97:
			goto st389
		case 100:
			goto st552
		case 102:
			goto st560
		case 106:
			goto st568
		case 109:
			goto st580
		case 110:
			goto st586
		case 111:
			goto st594
		case 115:
			goto st601
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr545
			}
		case data[p] >= 49:
			goto tr543
		}
		goto st0
tr542:
//line ragel/datetime.rl:5
 pb = p 
	goto st385
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
//line ragel/parse_datetime.go:14758
		if 49 <= data[p] && data[p] <= 57 {
			goto st386
		}
		goto st0
tr545:
//line ragel/datetime.rl:5
 pb = p 
	goto st386
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
//line ragel/parse_datetime.go:14772
		if data[p] == 32 {
			goto tr178
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr178
		}
		goto st0
tr543:
//line ragel/datetime.rl:5
 pb = p 
	goto st387
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
//line ragel/parse_datetime.go:14789
		if data[p] == 32 {
			goto tr178
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st386
			}
		case data[p] >= 45:
			goto tr178
		}
		goto st0
tr544:
//line ragel/datetime.rl:5
 pb = p 
	goto st388
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
//line ragel/parse_datetime.go:14811
		if data[p] == 32 {
			goto tr178
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st386
			}
		case data[p] >= 45:
			goto tr178
		}
		goto st0
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 112:
			goto st390
		case 117:
			goto st547
		}
		goto st0
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		if data[p] == 114 {
			goto st391
		}
		goto st0
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
		switch data[p] {
		case 32:
			goto tr558
		case 46:
			goto tr559
		case 105:
			goto st545
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr375
		}
		goto st0
tr558:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st392
tr825:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st392
tr832:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st392
tr841:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st392
tr851:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st392
tr860:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st392
tr864:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st392
tr871:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st392
tr876:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st392
tr881:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st392
tr891:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st392
tr900:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st392
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
//line ragel/parse_datetime.go:14915
		switch data[p] {
		case 32:
			goto st393
		case 48:
			goto tr562
		case 51:
			goto tr564
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr565
			}
		case data[p] >= 49:
			goto tr563
		}
		goto st0
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		switch data[p] {
		case 48:
			goto tr566
		case 51:
			goto tr568
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr569
			}
		case data[p] >= 49:
			goto tr567
		}
		goto st0
tr566:
//line ragel/datetime.rl:5
 pb = p 
	goto st394
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
//line ragel/parse_datetime.go:14962
		if 49 <= data[p] && data[p] <= 57 {
			goto st395
		}
		goto st0
tr569:
//line ragel/datetime.rl:5
 pb = p 
	goto st395
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
//line ragel/parse_datetime.go:14976
		if data[p] == 32 {
			goto tr571
		}
		goto st0
tr571:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st396
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
//line ragel/parse_datetime.go:14997
		if data[p] == 50 {
			goto tr573
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr574
			}
		case data[p] >= 48:
			goto tr572
		}
		goto st0
tr572:
//line ragel/datetime.rl:5
 pb = p 
	goto st397
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
//line ragel/parse_datetime.go:15019
		switch data[p] {
		case 32:
			goto tr575
		case 43:
			goto tr576
		case 45:
			goto tr577
		case 47:
			goto tr578
		case 58:
			goto tr580
		case 65:
			goto tr581
		case 80:
			goto tr581
		case 90:
			goto tr582
		case 95:
			goto tr578
		case 97:
			goto tr583
		case 112:
			goto tr583
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st433
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr578
			}
		default:
			goto tr578
		}
		goto st0
tr575:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st398
tr638:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st398
tr701:
//line ragel/datetime.rl:117

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

	goto st398
tr672:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st398
tr681:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st398
tr691:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st398
tr712:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st398
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
//line ragel/parse_datetime.go:15129
		switch data[p] {
		case 32:
			goto st399
		case 43:
			goto st403
		case 45:
			goto st415
		case 47:
			goto tr587
		case 65:
			goto tr589
		case 80:
			goto tr589
		case 90:
			goto tr590
		case 95:
			goto tr587
		case 97:
			goto tr591
		case 112:
			goto tr591
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr588
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr587
			}
		default:
			goto tr587
		}
		goto st0
tr621:
//line ragel/datetime.rl:166

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
	goto st399
tr604:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st399
tr625:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st399
tr628:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st399
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
//line ragel/parse_datetime.go:15228
		if 48 <= data[p] && data[p] <= 57 {
			goto tr588
		}
		goto st0
tr588:
//line ragel/datetime.rl:5
 pb = p 
	goto st400
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
//line ragel/parse_datetime.go:15242
		if 48 <= data[p] && data[p] <= 57 {
			goto st401
		}
		goto st0
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		if 48 <= data[p] && data[p] <= 57 {
			goto st402
		}
		goto st0
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
		if 48 <= data[p] && data[p] <= 57 {
			goto st948
		}
		goto st0
	st948:
		if p++; p == pe {
			goto _test_eof948
		}
	st_case_948:
		goto st0
tr576:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st403
tr635:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st403
tr639:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st403
tr657:
//line ragel/datetime.rl:117

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

	goto st403
tr649:
//line ragel/datetime.rl:35

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

	goto st403
tr673:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st403
tr682:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st403
tr692:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st403
tr713:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st403
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
//line ragel/parse_datetime.go:15385
		if data[p] == 50 {
			goto tr596
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr597
			}
		case data[p] >= 48:
			goto tr595
		}
		goto st0
tr595:
//line ragel/datetime.rl:5
 pb = p 
	goto st404
tr611:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st404
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
//line ragel/parse_datetime.go:15413
		switch data[p] {
		case 32:
			goto tr598
		case 58:
			goto tr600
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st409
		}
		goto st0
tr598:
//line ragel/datetime.rl:166

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
	goto st405
tr606:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st405
tr609:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st405
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
//line ragel/parse_datetime.go:15478
		switch data[p] {
		case 47:
			goto tr601
		case 95:
			goto tr601
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr588
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr601
			}
		default:
			goto tr601
		}
		goto st0
tr601:
//line ragel/datetime.rl:5
 pb = p 
	goto st406
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
//line ragel/parse_datetime.go:15507
		switch data[p] {
		case 47:
			goto st407
		case 95:
			goto st407
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st407
			}
		case data[p] >= 65:
			goto st407
		}
		goto st0
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 47:
			goto st408
		case 95:
			goto st408
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st408
			}
		case data[p] >= 65:
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
			goto tr604
		case 47:
			goto st408
		case 95:
			goto st408
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st408
			}
		case data[p] >= 65:
			goto st408
		}
		goto st0
tr597:
//line ragel/datetime.rl:5
 pb = p 
	goto st409
tr613:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st409
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
//line ragel/parse_datetime.go:15580
		switch data[p] {
		case 32:
			goto tr598
		case 58:
			goto tr600
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st410
		}
		goto st0
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		if data[p] == 32 {
			goto tr598
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st410
		}
		goto st0
tr600:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st411
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
//line ragel/parse_datetime.go:15619
		if data[p] == 32 {
			goto tr606
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr608
			}
		case data[p] >= 48:
			goto tr607
		}
		goto st0
tr607:
//line ragel/datetime.rl:5
 pb = p 
	goto st412
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
//line ragel/parse_datetime.go:15641
		if data[p] == 32 {
			goto tr609
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st413
		}
		goto st0
tr608:
//line ragel/datetime.rl:5
 pb = p 
	goto st413
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
//line ragel/parse_datetime.go:15658
		if data[p] == 32 {
			goto tr609
		}
		goto st0
tr596:
//line ragel/datetime.rl:5
 pb = p 
	goto st414
tr612:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st414
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
//line ragel/parse_datetime.go:15678
		switch data[p] {
		case 32:
			goto tr598
		case 58:
			goto tr600
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st410
			}
		case data[p] >= 48:
			goto st409
		}
		goto st0
tr577:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st415
tr636:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st415
tr640:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st415
tr658:
//line ragel/datetime.rl:117

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

	goto st415
tr651:
//line ragel/datetime.rl:35

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

	goto st415
tr674:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st415
tr683:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st415
tr694:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st415
tr715:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st415
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
//line ragel/parse_datetime.go:15808
		if data[p] == 50 {
			goto tr612
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr613
			}
		case data[p] >= 48:
			goto tr611
		}
		goto st0
tr587:
//line ragel/datetime.rl:5
 pb = p 
	goto st416
tr578:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st416
tr641:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st416
tr675:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st416
tr684:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st416
tr695:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st416
tr659:
//line ragel/datetime.rl:117

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
	goto st416
tr716:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st416
tr652:
//line ragel/datetime.rl:35

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
	goto st416
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
//line ragel/parse_datetime.go:15930
		switch data[p] {
		case 47:
			goto st417
		case 95:
			goto st417
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st417
			}
		case data[p] >= 65:
			goto st417
		}
		goto st0
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		switch data[p] {
		case 47:
			goto st418
		case 95:
			goto st418
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st418
			}
		case data[p] >= 65:
			goto st418
		}
		goto st0
tr637:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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
	goto st418
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
//line ragel/parse_datetime.go:15998
		switch data[p] {
		case 32:
			goto tr604
		case 43:
			goto tr616
		case 45:
			goto tr617
		case 47:
			goto st418
		case 95:
			goto st418
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st418
			}
		case data[p] >= 65:
			goto st418
		}
		goto st0
tr616:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st419
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
//line ragel/parse_datetime.go:16032
		if data[p] == 50 {
			goto tr619
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr620
			}
		case data[p] >= 48:
			goto tr618
		}
		goto st0
tr618:
//line ragel/datetime.rl:5
 pb = p 
	goto st420
tr630:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st420
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
//line ragel/parse_datetime.go:16060
		switch data[p] {
		case 32:
			goto tr621
		case 58:
			goto tr623
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st421
		}
		goto st0
tr620:
//line ragel/datetime.rl:5
 pb = p 
	goto st421
tr632:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st421
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
//line ragel/parse_datetime.go:16086
		switch data[p] {
		case 32:
			goto tr621
		case 58:
			goto tr623
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st422
		}
		goto st0
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		if data[p] == 32 {
			goto tr621
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st422
		}
		goto st0
tr623:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st423
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
//line ragel/parse_datetime.go:16125
		if data[p] == 32 {
			goto tr625
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr627
			}
		case data[p] >= 48:
			goto tr626
		}
		goto st0
tr626:
//line ragel/datetime.rl:5
 pb = p 
	goto st424
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
//line ragel/parse_datetime.go:16147
		if data[p] == 32 {
			goto tr628
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st425
		}
		goto st0
tr627:
//line ragel/datetime.rl:5
 pb = p 
	goto st425
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
//line ragel/parse_datetime.go:16164
		if data[p] == 32 {
			goto tr628
		}
		goto st0
tr619:
//line ragel/datetime.rl:5
 pb = p 
	goto st426
tr631:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
//line ragel/parse_datetime.go:16184
		switch data[p] {
		case 32:
			goto tr621
		case 58:
			goto tr623
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st422
			}
		case data[p] >= 48:
			goto st421
		}
		goto st0
tr617:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st427
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
//line ragel/parse_datetime.go:16212
		if data[p] == 50 {
			goto tr631
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr632
			}
		case data[p] >= 48:
			goto tr630
		}
		goto st0
tr589:
//line ragel/datetime.rl:5
 pb = p 
	goto st428
tr581:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st428
tr644:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st428
tr678:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st428
tr686:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st428
tr697:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st428
tr703:
//line ragel/datetime.rl:117

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
	goto st428
tr717:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st428
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
//line ragel/parse_datetime.go:16315
		switch data[p] {
		case 47:
			goto st417
		case 77:
			goto st429
		case 95:
			goto st417
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st417
			}
		case data[p] >= 65:
			goto st417
		}
		goto st0
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
		switch data[p] {
		case 32:
			goto tr634
		case 43:
			goto tr635
		case 45:
			goto tr636
		case 47:
			goto tr637
		case 95:
			goto tr637
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr637
			}
		case data[p] >= 65:
			goto tr637
		}
		goto st0
tr634:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st430
tr656:
//line ragel/datetime.rl:117

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

	goto st430
tr648:
//line ragel/datetime.rl:35

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

	goto st430
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
//line ragel/parse_datetime.go:16437
		switch data[p] {
		case 32:
			goto st399
		case 43:
			goto st403
		case 45:
			goto st415
		case 47:
			goto tr587
		case 90:
			goto tr590
		case 95:
			goto tr587
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr588
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr587
			}
		default:
			goto tr587
		}
		goto st0
tr590:
//line ragel/datetime.rl:5
 pb = p 
	goto st431
tr582:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st431
tr645:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st431
tr679:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st431
tr687:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st431
tr698:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st431
tr661:
//line ragel/datetime.rl:117

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
	goto st431
tr718:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st431
tr654:
//line ragel/datetime.rl:35

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
	goto st431
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
//line ragel/parse_datetime.go:16574
		switch data[p] {
		case 32:
			goto tr625
		case 47:
			goto st417
		case 95:
			goto st417
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st417
			}
		case data[p] >= 65:
			goto st417
		}
		goto st0
tr591:
//line ragel/datetime.rl:5
 pb = p 
	goto st432
tr583:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st432
tr646:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st432
tr680:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st432
tr688:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st432
tr699:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st432
tr704:
//line ragel/datetime.rl:117

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
	goto st432
tr719:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st432
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
//line ragel/parse_datetime.go:16682
		switch data[p] {
		case 47:
			goto st417
		case 95:
			goto st417
		case 109:
			goto st429
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st417
			}
		case data[p] >= 65:
			goto st417
		}
		goto st0
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 32:
			goto tr638
		case 43:
			goto tr639
		case 45:
			goto tr640
		case 47:
			goto tr641
		case 58:
			goto tr643
		case 65:
			goto tr644
		case 80:
			goto tr644
		case 90:
			goto tr645
		case 95:
			goto tr641
		case 97:
			goto tr646
		case 112:
			goto tr646
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st434
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr641
			}
		default:
			goto tr641
		}
		goto st0
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		if 48 <= data[p] && data[p] <= 57 {
			goto st435
		}
		goto st0
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		switch data[p] {
		case 32:
			goto tr648
		case 43:
			goto tr649
		case 45:
			goto tr651
		case 47:
			goto tr652
		case 90:
			goto tr654
		case 95:
			goto tr652
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr650
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr652
				}
			case data[p] >= 65:
				goto tr652
			}
		default:
			goto st446
		}
		goto st0
tr650:
//line ragel/datetime.rl:35

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

	goto st436
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
//line ragel/parse_datetime.go:16810
		if 48 <= data[p] && data[p] <= 57 {
			goto tr655
		}
		goto st0
tr655:
//line ragel/datetime.rl:5
 pb = p 
	goto st437
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
//line ragel/parse_datetime.go:16824
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st438
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st439
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st440
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st441
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st442
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st443
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st444
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st445
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		case data[p] >= 65:
			goto tr659
		}
		goto st0
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr648
		case 43:
			goto tr649
		case 45:
			goto tr651
		case 47:
			goto tr652
		case 90:
			goto tr654
		case 95:
			goto tr652
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr650
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr652
			}
		default:
			goto tr652
		}
		goto st0
tr580:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st448
tr643:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st448
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
//line ragel/parse_datetime.go:17162
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr671
			}
		case data[p] >= 48:
			goto tr670
		}
		goto st0
tr670:
//line ragel/datetime.rl:5
 pb = p 
	goto st449
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
//line ragel/parse_datetime.go:17181
		switch data[p] {
		case 32:
			goto tr672
		case 43:
			goto tr673
		case 45:
			goto tr674
		case 47:
			goto tr675
		case 58:
			goto tr677
		case 65:
			goto tr678
		case 80:
			goto tr678
		case 90:
			goto tr679
		case 95:
			goto tr675
		case 97:
			goto tr680
		case 112:
			goto tr680
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st450
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr675
			}
		default:
			goto tr675
		}
		goto st0
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr682
		case 45:
			goto tr683
		case 47:
			goto tr684
		case 58:
			goto tr685
		case 65:
			goto tr686
		case 80:
			goto tr686
		case 90:
			goto tr687
		case 95:
			goto tr684
		case 97:
			goto tr688
		case 112:
			goto tr688
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr684
			}
		case data[p] >= 66:
			goto tr684
		}
		goto st0
tr677:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st451
tr685:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st451
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
//line ragel/parse_datetime.go:17274
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr690
			}
		case data[p] >= 48:
			goto tr689
		}
		goto st0
tr689:
//line ragel/datetime.rl:5
 pb = p 
	goto st452
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
//line ragel/parse_datetime.go:17293
		switch data[p] {
		case 32:
			goto tr691
		case 43:
			goto tr692
		case 45:
			goto tr694
		case 47:
			goto tr695
		case 65:
			goto tr697
		case 80:
			goto tr697
		case 90:
			goto tr698
		case 95:
			goto tr695
		case 97:
			goto tr699
		case 112:
			goto tr699
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr693
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr695
				}
			case data[p] >= 66:
				goto tr695
			}
		default:
			goto st463
		}
		goto st0
tr693:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st453
tr714:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st453
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
//line ragel/parse_datetime.go:17351
		if 48 <= data[p] && data[p] <= 57 {
			goto tr700
		}
		goto st0
tr700:
//line ragel/datetime.rl:5
 pb = p 
	goto st454
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
//line ragel/parse_datetime.go:17365
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr703
		case 80:
			goto tr703
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr704
		case 112:
			goto tr704
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st455
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr703
		case 80:
			goto tr703
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr704
		case 112:
			goto tr704
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st456
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr703
		case 80:
			goto tr703
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr704
		case 112:
			goto tr704
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st457
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr703
		case 80:
			goto tr703
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr704
		case 112:
			goto tr704
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st458
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr703
		case 80:
			goto tr703
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr704
		case 112:
			goto tr704
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st459
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr703
		case 80:
			goto tr703
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr704
		case 112:
			goto tr704
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st460
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr703
		case 80:
			goto tr703
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr704
		case 112:
			goto tr704
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st461
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr703
		case 80:
			goto tr703
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr704
		case 112:
			goto tr704
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st462
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr703
		case 80:
			goto tr703
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr704
		case 112:
			goto tr704
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		case data[p] >= 66:
			goto tr659
		}
		goto st0
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
		switch data[p] {
		case 32:
			goto tr712
		case 43:
			goto tr713
		case 45:
			goto tr715
		case 47:
			goto tr716
		case 65:
			goto tr717
		case 80:
			goto tr717
		case 90:
			goto tr718
		case 95:
			goto tr716
		case 97:
			goto tr719
		case 112:
			goto tr719
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr714
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr716
			}
		default:
			goto tr716
		}
		goto st0
tr690:
//line ragel/datetime.rl:5
 pb = p 
	goto st464
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
//line ragel/parse_datetime.go:17766
		switch data[p] {
		case 32:
			goto tr691
		case 43:
			goto tr692
		case 45:
			goto tr694
		case 47:
			goto tr695
		case 65:
			goto tr697
		case 80:
			goto tr697
		case 90:
			goto tr698
		case 95:
			goto tr695
		case 97:
			goto tr699
		case 112:
			goto tr699
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr693
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr695
			}
		default:
			goto tr695
		}
		goto st0
tr671:
//line ragel/datetime.rl:5
 pb = p 
	goto st465
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
//line ragel/parse_datetime.go:17811
		switch data[p] {
		case 32:
			goto tr672
		case 43:
			goto tr673
		case 45:
			goto tr674
		case 47:
			goto tr675
		case 58:
			goto tr677
		case 65:
			goto tr678
		case 80:
			goto tr678
		case 90:
			goto tr679
		case 95:
			goto tr675
		case 97:
			goto tr680
		case 112:
			goto tr680
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr675
			}
		case data[p] >= 66:
			goto tr675
		}
		goto st0
tr573:
//line ragel/datetime.rl:5
 pb = p 
	goto st466
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
//line ragel/parse_datetime.go:17854
		switch data[p] {
		case 32:
			goto tr575
		case 43:
			goto tr576
		case 45:
			goto tr577
		case 47:
			goto tr578
		case 58:
			goto tr580
		case 65:
			goto tr581
		case 80:
			goto tr581
		case 90:
			goto tr582
		case 95:
			goto tr578
		case 97:
			goto tr583
		case 112:
			goto tr583
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st433
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr578
				}
			case data[p] >= 66:
				goto tr578
			}
		default:
			goto st467
		}
		goto st0
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
		if 48 <= data[p] && data[p] <= 57 {
			goto st434
		}
		goto st0
tr574:
//line ragel/datetime.rl:5
 pb = p 
	goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
//line ragel/parse_datetime.go:17915
		switch data[p] {
		case 32:
			goto tr575
		case 43:
			goto tr576
		case 45:
			goto tr577
		case 47:
			goto tr578
		case 58:
			goto tr580
		case 65:
			goto tr581
		case 80:
			goto tr581
		case 90:
			goto tr582
		case 95:
			goto tr578
		case 97:
			goto tr583
		case 112:
			goto tr583
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st467
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr578
			}
		default:
			goto tr578
		}
		goto st0
tr567:
//line ragel/datetime.rl:5
 pb = p 
	goto st469
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
//line ragel/parse_datetime.go:17962
		if data[p] == 32 {
			goto tr571
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st395
		}
		goto st0
tr568:
//line ragel/datetime.rl:5
 pb = p 
	goto st470
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
//line ragel/parse_datetime.go:17979
		if data[p] == 32 {
			goto tr571
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st395
		}
		goto st0
tr562:
//line ragel/datetime.rl:5
 pb = p 
	goto st471
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
//line ragel/parse_datetime.go:17996
		if 49 <= data[p] && data[p] <= 57 {
			goto st472
		}
		goto st0
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
		if data[p] == 32 {
			goto tr722
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr385
		}
		goto st0
tr722:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st473
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
//line ragel/parse_datetime.go:18029
		if data[p] == 50 {
			goto tr724
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr725
			}
		case data[p] >= 48:
			goto tr723
		}
		goto st0
tr723:
//line ragel/datetime.rl:5
 pb = p 
	goto st474
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
//line ragel/parse_datetime.go:18051
		switch data[p] {
		case 32:
			goto tr726
		case 43:
			goto tr576
		case 45:
			goto tr577
		case 47:
			goto tr578
		case 58:
			goto tr728
		case 65:
			goto tr729
		case 80:
			goto tr729
		case 90:
			goto tr582
		case 95:
			goto tr578
		case 97:
			goto tr730
		case 112:
			goto tr730
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st480
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr578
			}
		default:
			goto tr578
		}
		goto st0
tr726:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st475
tr735:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st475
tr796:
//line ragel/datetime.rl:117

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

	goto st475
tr779:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st475
tr784:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st475
tr790:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st475
tr807:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st475
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
//line ragel/parse_datetime.go:18161
		switch data[p] {
		case 32:
			goto st399
		case 43:
			goto st403
		case 45:
			goto st415
		case 47:
			goto tr587
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr590
		case 95:
			goto tr587
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr394
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr587
			}
		default:
			goto tr587
		}
		goto st0
tr731:
//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr729:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr738:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr782:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr786:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr793:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
tr798:
//line ragel/datetime.rl:117

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
tr809:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st476
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
//line ragel/parse_datetime.go:18287
		switch data[p] {
		case 47:
			goto st417
		case 77:
			goto st477
		case 95:
			goto st417
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st417
			}
		case data[p] >= 65:
			goto st417
		}
		goto st0
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		switch data[p] {
		case 32:
			goto tr734
		case 43:
			goto tr635
		case 45:
			goto tr636
		case 47:
			goto tr637
		case 95:
			goto tr637
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr637
			}
		case data[p] >= 65:
			goto tr637
		}
		goto st0
tr734:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

	goto st478
tr765:
//line ragel/datetime.rl:117

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

	goto st478
tr775:
//line ragel/datetime.rl:35

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

	goto st478
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
//line ragel/parse_datetime.go:18409
		switch data[p] {
		case 32:
			goto st399
		case 43:
			goto st403
		case 45:
			goto st415
		case 47:
			goto tr587
		case 90:
			goto tr590
		case 95:
			goto tr587
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr394
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr587
			}
		default:
			goto tr587
		}
		goto st0
tr732:
//line ragel/datetime.rl:5
 pb = p 
	goto st479
tr730:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st479
tr739:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st479
tr783:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st479
tr787:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st479
tr794:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st479
tr799:
//line ragel/datetime.rl:117

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
	goto st479
tr810:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st479
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
//line ragel/parse_datetime.go:18527
		switch data[p] {
		case 47:
			goto st417
		case 95:
			goto st417
		case 109:
			goto st477
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st417
			}
		case data[p] >= 65:
			goto st417
		}
		goto st0
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		switch data[p] {
		case 32:
			goto tr735
		case 43:
			goto tr639
		case 45:
			goto tr640
		case 47:
			goto tr641
		case 58:
			goto tr737
		case 65:
			goto tr738
		case 80:
			goto tr738
		case 90:
			goto tr645
		case 95:
			goto tr641
		case 97:
			goto tr739
		case 112:
			goto tr739
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st481
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr641
			}
		default:
			goto tr641
		}
		goto st0
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		if 48 <= data[p] && data[p] <= 57 {
			goto st949
		}
		goto st0
	st949:
		if p++; p == pe {
			goto _test_eof949
		}
	st_case_949:
		switch data[p] {
		case 32:
			goto tr1425
		case 43:
			goto tr1426
		case 44:
			goto tr1427
		case 45:
			goto tr1428
		case 46:
			goto tr776
		case 47:
			goto tr1429
		case 84:
			goto tr1431
		case 90:
			goto tr1432
		case 95:
			goto tr1433
		case 116:
			goto tr1433
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st507
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1429
			}
		default:
			goto tr1429
		}
		goto st0
tr1425:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:35

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

	goto st950
	st950:
		if p++; p == pe {
			goto _test_eof950
		}
	st_case_950:
//line ragel/parse_datetime.go:18662
		switch data[p] {
		case 32:
			goto st482
		case 43:
			goto st483
		case 45:
			goto st489
		case 47:
			goto tr1437
		case 50:
			goto tr1327
		case 65:
			goto tr1438
		case 66:
			goto tr1439
		case 90:
			goto tr1440
		case 95:
			goto tr1437
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr1326
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1437
				}
			case data[p] >= 67:
				goto tr1437
			}
		default:
			goto tr1328
		}
		goto st0
tr1453:
//line ragel/datetime.rl:166

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
	goto st482
tr1444:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st482
tr1457:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st482
tr1460:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st482
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
//line ragel/parse_datetime.go:18764
		switch data[p] {
		case 65:
			goto st12
		case 66:
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr588
		}
		goto st0
tr1426:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:35

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

	goto st483
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
//line ragel/parse_datetime.go:18801
		if data[p] == 50 {
			goto tr742
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr743
			}
		case data[p] >= 48:
			goto tr741
		}
		goto st0
tr741:
//line ragel/datetime.rl:5
 pb = p 
	goto st951
tr751:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st951
	st951:
		if p++; p == pe {
			goto _test_eof951
		}
	st_case_951:
//line ragel/parse_datetime.go:18829
		switch data[p] {
		case 32:
			goto tr1441
		case 58:
			goto tr1443
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st955
		}
		goto st0
tr1441:
//line ragel/datetime.rl:166

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
	goto st484
tr1446:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st484
tr1449:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st484
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
//line ragel/parse_datetime.go:18894
		switch data[p] {
		case 47:
			goto tr744
		case 65:
			goto tr745
		case 66:
			goto tr746
		case 95:
			goto tr744
		}
		switch {
		case data[p] < 67:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr588
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr744
			}
		default:
			goto tr744
		}
		goto st0
tr744:
//line ragel/datetime.rl:5
 pb = p 
	goto st485
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
//line ragel/parse_datetime.go:18927
		switch data[p] {
		case 47:
			goto st486
		case 95:
			goto st486
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st486
			}
		case data[p] >= 65:
			goto st486
		}
		goto st0
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		switch data[p] {
		case 47:
			goto st952
		case 95:
			goto st952
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st952
			}
		case data[p] >= 65:
			goto st952
		}
		goto st0
	st952:
		if p++; p == pe {
			goto _test_eof952
		}
	st_case_952:
		switch data[p] {
		case 32:
			goto tr1444
		case 47:
			goto st952
		case 95:
			goto st952
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st952
			}
		case data[p] >= 65:
			goto st952
		}
		goto st0
tr745:
//line ragel/datetime.rl:5
 pb = p 
	goto st487
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
//line ragel/parse_datetime.go:18994
		switch data[p] {
		case 47:
			goto st486
		case 68:
			goto st953
		case 95:
			goto st486
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st486
			}
		case data[p] >= 65:
			goto st486
		}
		goto st0
	st953:
		if p++; p == pe {
			goto _test_eof953
		}
	st_case_953:
		switch data[p] {
		case 47:
			goto st952
		case 95:
			goto st952
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st952
			}
		case data[p] >= 65:
			goto st952
		}
		goto st0
tr746:
//line ragel/datetime.rl:5
 pb = p 
	goto st488
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
//line ragel/parse_datetime.go:19041
		switch data[p] {
		case 47:
			goto st486
		case 67:
			goto st954
		case 95:
			goto st486
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st486
			}
		case data[p] >= 65:
			goto st486
		}
		goto st0
	st954:
		if p++; p == pe {
			goto _test_eof954
		}
	st_case_954:
		switch data[p] {
		case 47:
			goto st952
		case 95:
			goto st952
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st952
			}
		case data[p] >= 65:
			goto st952
		}
		goto st0
tr743:
//line ragel/datetime.rl:5
 pb = p 
	goto st955
tr753:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st955
	st955:
		if p++; p == pe {
			goto _test_eof955
		}
	st_case_955:
//line ragel/parse_datetime.go:19094
		switch data[p] {
		case 32:
			goto tr1441
		case 58:
			goto tr1443
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st956
		}
		goto st0
	st956:
		if p++; p == pe {
			goto _test_eof956
		}
	st_case_956:
		if data[p] == 32 {
			goto tr1441
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st956
		}
		goto st0
tr1443:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st957
	st957:
		if p++; p == pe {
			goto _test_eof957
		}
	st_case_957:
//line ragel/parse_datetime.go:19133
		if data[p] == 32 {
			goto tr1446
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1448
			}
		case data[p] >= 48:
			goto tr1447
		}
		goto st0
tr1447:
//line ragel/datetime.rl:5
 pb = p 
	goto st958
	st958:
		if p++; p == pe {
			goto _test_eof958
		}
	st_case_958:
//line ragel/parse_datetime.go:19155
		if data[p] == 32 {
			goto tr1449
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st959
		}
		goto st0
tr1448:
//line ragel/datetime.rl:5
 pb = p 
	goto st959
	st959:
		if p++; p == pe {
			goto _test_eof959
		}
	st_case_959:
//line ragel/parse_datetime.go:19172
		if data[p] == 32 {
			goto tr1449
		}
		goto st0
tr742:
//line ragel/datetime.rl:5
 pb = p 
	goto st960
tr752:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st960
	st960:
		if p++; p == pe {
			goto _test_eof960
		}
	st_case_960:
//line ragel/parse_datetime.go:19192
		switch data[p] {
		case 32:
			goto tr1441
		case 58:
			goto tr1443
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st956
			}
		case data[p] >= 48:
			goto st955
		}
		goto st0
tr1428:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:35

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

	goto st489
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
//line ragel/parse_datetime.go:19234
		if data[p] == 50 {
			goto tr752
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr753
			}
		case data[p] >= 48:
			goto tr751
		}
		goto st0
tr1437:
//line ragel/datetime.rl:5
 pb = p 
	goto st490
tr1429:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
//line ragel/datetime.rl:35

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

	goto st490
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
//line ragel/parse_datetime.go:19279
		switch data[p] {
		case 47:
			goto st491
		case 95:
			goto st491
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st491
			}
		case data[p] >= 65:
			goto st491
		}
		goto st0
tr1462:
//line ragel/datetime.rl:5
 pb = p 
	goto st491
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
//line ragel/parse_datetime.go:19304
		switch data[p] {
		case 47:
			goto st961
		case 95:
			goto st961
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st961
			}
		case data[p] >= 65:
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
			goto tr1444
		case 43:
			goto tr1451
		case 45:
			goto tr1452
		case 47:
			goto st961
		case 95:
			goto st961
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st961
			}
		case data[p] >= 65:
			goto st961
		}
		goto st0
tr1451:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st492
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
//line ragel/parse_datetime.go:19358
		if data[p] == 50 {
			goto tr757
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr758
			}
		case data[p] >= 48:
			goto tr756
		}
		goto st0
tr756:
//line ragel/datetime.rl:5
 pb = p 
	goto st962
tr759:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st962
	st962:
		if p++; p == pe {
			goto _test_eof962
		}
	st_case_962:
//line ragel/parse_datetime.go:19386
		switch data[p] {
		case 32:
			goto tr1453
		case 58:
			goto tr1455
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st963
		}
		goto st0
tr758:
//line ragel/datetime.rl:5
 pb = p 
	goto st963
tr761:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st963
	st963:
		if p++; p == pe {
			goto _test_eof963
		}
	st_case_963:
//line ragel/parse_datetime.go:19412
		switch data[p] {
		case 32:
			goto tr1453
		case 58:
			goto tr1455
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st964
		}
		goto st0
	st964:
		if p++; p == pe {
			goto _test_eof964
		}
	st_case_964:
		if data[p] == 32 {
			goto tr1453
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st964
		}
		goto st0
tr1455:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st965
	st965:
		if p++; p == pe {
			goto _test_eof965
		}
	st_case_965:
//line ragel/parse_datetime.go:19451
		if data[p] == 32 {
			goto tr1457
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1459
			}
		case data[p] >= 48:
			goto tr1458
		}
		goto st0
tr1458:
//line ragel/datetime.rl:5
 pb = p 
	goto st966
	st966:
		if p++; p == pe {
			goto _test_eof966
		}
	st_case_966:
//line ragel/parse_datetime.go:19473
		if data[p] == 32 {
			goto tr1460
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st967
		}
		goto st0
tr1459:
//line ragel/datetime.rl:5
 pb = p 
	goto st967
	st967:
		if p++; p == pe {
			goto _test_eof967
		}
	st_case_967:
//line ragel/parse_datetime.go:19490
		if data[p] == 32 {
			goto tr1460
		}
		goto st0
tr757:
//line ragel/datetime.rl:5
 pb = p 
	goto st968
tr760:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st968
	st968:
		if p++; p == pe {
			goto _test_eof968
		}
	st_case_968:
//line ragel/parse_datetime.go:19510
		switch data[p] {
		case 32:
			goto tr1453
		case 58:
			goto tr1455
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st964
			}
		case data[p] >= 48:
			goto st963
		}
		goto st0
tr1452:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st493
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
//line ragel/parse_datetime.go:19538
		if data[p] == 50 {
			goto tr760
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr761
			}
		case data[p] >= 48:
			goto tr759
		}
		goto st0
tr1438:
//line ragel/datetime.rl:5
 pb = p 
	goto st494
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
//line ragel/parse_datetime.go:19560
		switch data[p] {
		case 47:
			goto st491
		case 68:
			goto st969
		case 95:
			goto st491
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st491
			}
		case data[p] >= 65:
			goto st491
		}
		goto st0
	st969:
		if p++; p == pe {
			goto _test_eof969
		}
	st_case_969:
		switch data[p] {
		case 47:
			goto st961
		case 95:
			goto st961
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st961
			}
		case data[p] >= 65:
			goto st961
		}
		goto st0
tr1439:
//line ragel/datetime.rl:5
 pb = p 
	goto st495
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
//line ragel/parse_datetime.go:19607
		switch data[p] {
		case 47:
			goto st491
		case 67:
			goto st970
		case 95:
			goto st491
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st491
			}
		case data[p] >= 65:
			goto st491
		}
		goto st0
	st970:
		if p++; p == pe {
			goto _test_eof970
		}
	st_case_970:
		switch data[p] {
		case 47:
			goto st961
		case 95:
			goto st961
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st961
			}
		case data[p] >= 65:
			goto st961
		}
		goto st0
tr1440:
//line ragel/datetime.rl:5
 pb = p 
	goto st971
tr1432:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
//line ragel/datetime.rl:35

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

	goto st971
	st971:
		if p++; p == pe {
			goto _test_eof971
		}
	st_case_971:
//line ragel/parse_datetime.go:19677
		switch data[p] {
		case 32:
			goto tr1457
		case 47:
			goto st491
		case 95:
			goto st491
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st491
			}
		case data[p] >= 65:
			goto st491
		}
		goto st0
tr1427:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:35

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

	goto st496
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
//line ragel/parse_datetime.go:19721
		if data[p] == 32 {
			goto st38
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr764
		}
		goto st0
tr764:
//line ragel/datetime.rl:5
 pb = p 
	goto st497
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
//line ragel/parse_datetime.go:19738
		switch data[p] {
		case 32:
			goto tr765
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st498
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		switch data[p] {
		case 32:
			goto tr765
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st499
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		switch data[p] {
		case 32:
			goto tr765
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st500
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		switch data[p] {
		case 32:
			goto tr765
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st501
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		switch data[p] {
		case 32:
			goto tr765
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st502
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		switch data[p] {
		case 32:
			goto tr765
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st503
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		switch data[p] {
		case 32:
			goto tr765
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st504
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		switch data[p] {
		case 32:
			goto tr765
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st505
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		switch data[p] {
		case 32:
			goto tr765
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 90:
			goto tr661
		case 95:
			goto tr659
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		case data[p] >= 65:
			goto tr659
		}
		goto st0
tr776:
//line ragel/datetime.rl:35

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

	goto st506
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
//line ragel/parse_datetime.go:20040
		if 48 <= data[p] && data[p] <= 57 {
			goto tr764
		}
		goto st0
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		if 48 <= data[p] && data[p] <= 57 {
			goto st508
		}
		goto st0
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		switch data[p] {
		case 32:
			goto tr775
		case 43:
			goto tr649
		case 45:
			goto tr651
		case 47:
			goto tr652
		case 90:
			goto tr654
		case 95:
			goto tr652
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr776
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr652
			}
		default:
			goto tr652
		}
		goto st0
tr1431:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
//line ragel/datetime.rl:35

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

	goto st972
	st972:
		if p++; p == pe {
			goto _test_eof972
		}
	st_case_972:
//line ragel/parse_datetime.go:20114
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr1462
		case 50:
			goto tr82
		case 90:
			goto tr1463
		case 95:
			goto tr1462
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr81
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1462
				}
			case data[p] >= 65:
				goto tr1462
			}
		default:
			goto tr83
		}
		goto st0
tr1463:
//line ragel/datetime.rl:5
 pb = p 
	goto st973
	st973:
		if p++; p == pe {
			goto _test_eof973
		}
	st_case_973:
//line ragel/parse_datetime.go:20158
		switch data[p] {
		case 32:
			goto tr1175
		case 47:
			goto st961
		case 95:
			goto st961
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st961
			}
		case data[p] >= 65:
			goto st961
		}
		goto st0
tr1433:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
//line ragel/datetime.rl:35

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

	goto st509
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
//line ragel/parse_datetime.go:20204
		switch data[p] {
		case 47:
			goto st491
		case 50:
			goto tr82
		case 95:
			goto st491
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr81
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st491
				}
			case data[p] >= 65:
				goto st491
			}
		default:
			goto tr83
		}
		goto st0
tr728:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st510
tr737:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st510
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
//line ragel/parse_datetime.go:20248
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr778
			}
		case data[p] >= 48:
			goto tr777
		}
		goto st0
tr777:
//line ragel/datetime.rl:5
 pb = p 
	goto st511
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
//line ragel/parse_datetime.go:20267
		switch data[p] {
		case 32:
			goto tr779
		case 43:
			goto tr673
		case 45:
			goto tr674
		case 47:
			goto tr675
		case 58:
			goto tr781
		case 65:
			goto tr782
		case 80:
			goto tr782
		case 90:
			goto tr679
		case 95:
			goto tr675
		case 97:
			goto tr783
		case 112:
			goto tr783
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st512
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr675
			}
		default:
			goto tr675
		}
		goto st0
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		switch data[p] {
		case 32:
			goto tr784
		case 43:
			goto tr682
		case 45:
			goto tr683
		case 47:
			goto tr684
		case 58:
			goto tr785
		case 65:
			goto tr786
		case 80:
			goto tr786
		case 90:
			goto tr687
		case 95:
			goto tr684
		case 97:
			goto tr787
		case 112:
			goto tr787
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr684
			}
		case data[p] >= 66:
			goto tr684
		}
		goto st0
tr781:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st513
tr785:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st513
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
//line ragel/parse_datetime.go:20360
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr789
			}
		case data[p] >= 48:
			goto tr788
		}
		goto st0
tr788:
//line ragel/datetime.rl:5
 pb = p 
	goto st514
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
//line ragel/parse_datetime.go:20379
		switch data[p] {
		case 32:
			goto tr790
		case 43:
			goto tr692
		case 45:
			goto tr694
		case 47:
			goto tr695
		case 65:
			goto tr793
		case 80:
			goto tr793
		case 90:
			goto tr698
		case 95:
			goto tr695
		case 97:
			goto tr794
		case 112:
			goto tr794
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr791
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr695
				}
			case data[p] >= 66:
				goto tr695
			}
		default:
			goto st525
		}
		goto st0
tr791:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st515
tr808:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st515
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
//line ragel/parse_datetime.go:20437
		if 48 <= data[p] && data[p] <= 57 {
			goto tr795
		}
		goto st0
tr795:
//line ragel/datetime.rl:5
 pb = p 
	goto st516
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
//line ragel/parse_datetime.go:20451
		switch data[p] {
		case 32:
			goto tr796
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr798
		case 80:
			goto tr798
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr799
		case 112:
			goto tr799
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st517
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		switch data[p] {
		case 32:
			goto tr796
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr798
		case 80:
			goto tr798
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr799
		case 112:
			goto tr799
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st518
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		switch data[p] {
		case 32:
			goto tr796
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr798
		case 80:
			goto tr798
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr799
		case 112:
			goto tr799
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st519
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		switch data[p] {
		case 32:
			goto tr796
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr798
		case 80:
			goto tr798
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr799
		case 112:
			goto tr799
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st520
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
		switch data[p] {
		case 32:
			goto tr796
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr798
		case 80:
			goto tr798
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr799
		case 112:
			goto tr799
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st521
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		switch data[p] {
		case 32:
			goto tr796
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr798
		case 80:
			goto tr798
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr799
		case 112:
			goto tr799
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st522
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		switch data[p] {
		case 32:
			goto tr796
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr798
		case 80:
			goto tr798
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr799
		case 112:
			goto tr799
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st523
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
		switch data[p] {
		case 32:
			goto tr796
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr798
		case 80:
			goto tr798
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr799
		case 112:
			goto tr799
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st524
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		default:
			goto tr659
		}
		goto st0
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
		switch data[p] {
		case 32:
			goto tr796
		case 43:
			goto tr657
		case 45:
			goto tr658
		case 47:
			goto tr659
		case 65:
			goto tr798
		case 80:
			goto tr798
		case 90:
			goto tr661
		case 95:
			goto tr659
		case 97:
			goto tr799
		case 112:
			goto tr799
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr659
			}
		case data[p] >= 66:
			goto tr659
		}
		goto st0
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		switch data[p] {
		case 32:
			goto tr807
		case 43:
			goto tr713
		case 45:
			goto tr715
		case 47:
			goto tr716
		case 65:
			goto tr809
		case 80:
			goto tr809
		case 90:
			goto tr718
		case 95:
			goto tr716
		case 97:
			goto tr810
		case 112:
			goto tr810
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr808
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr716
			}
		default:
			goto tr716
		}
		goto st0
tr789:
//line ragel/datetime.rl:5
 pb = p 
	goto st526
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
//line ragel/parse_datetime.go:20852
		switch data[p] {
		case 32:
			goto tr790
		case 43:
			goto tr692
		case 45:
			goto tr694
		case 47:
			goto tr695
		case 65:
			goto tr793
		case 80:
			goto tr793
		case 90:
			goto tr698
		case 95:
			goto tr695
		case 97:
			goto tr794
		case 112:
			goto tr794
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr791
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr695
			}
		default:
			goto tr695
		}
		goto st0
tr778:
//line ragel/datetime.rl:5
 pb = p 
	goto st527
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
//line ragel/parse_datetime.go:20897
		switch data[p] {
		case 32:
			goto tr779
		case 43:
			goto tr673
		case 45:
			goto tr674
		case 47:
			goto tr675
		case 58:
			goto tr781
		case 65:
			goto tr782
		case 80:
			goto tr782
		case 90:
			goto tr679
		case 95:
			goto tr675
		case 97:
			goto tr783
		case 112:
			goto tr783
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr675
			}
		case data[p] >= 66:
			goto tr675
		}
		goto st0
tr724:
//line ragel/datetime.rl:5
 pb = p 
	goto st528
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
//line ragel/parse_datetime.go:20940
		switch data[p] {
		case 32:
			goto tr726
		case 43:
			goto tr576
		case 45:
			goto tr577
		case 47:
			goto tr578
		case 58:
			goto tr728
		case 65:
			goto tr729
		case 80:
			goto tr729
		case 90:
			goto tr582
		case 95:
			goto tr578
		case 97:
			goto tr730
		case 112:
			goto tr730
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st480
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr578
				}
			case data[p] >= 66:
				goto tr578
			}
		default:
			goto st529
		}
		goto st0
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
		if 48 <= data[p] && data[p] <= 57 {
			goto st481
		}
		goto st0
tr725:
//line ragel/datetime.rl:5
 pb = p 
	goto st530
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
//line ragel/parse_datetime.go:21001
		switch data[p] {
		case 32:
			goto tr726
		case 43:
			goto tr576
		case 45:
			goto tr577
		case 47:
			goto tr578
		case 58:
			goto tr728
		case 65:
			goto tr729
		case 80:
			goto tr729
		case 90:
			goto tr582
		case 95:
			goto tr578
		case 97:
			goto tr730
		case 112:
			goto tr730
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st529
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr578
			}
		default:
			goto tr578
		}
		goto st0
tr563:
//line ragel/datetime.rl:5
 pb = p 
	goto st531
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
//line ragel/parse_datetime.go:21048
		if data[p] == 32 {
			goto tr812
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st472
			}
		case data[p] >= 45:
			goto tr385
		}
		goto st0
tr812:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st532
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
//line ragel/parse_datetime.go:21077
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
tr813:
//line ragel/datetime.rl:5
 pb = p 
	goto st533
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
//line ragel/parse_datetime.go:21099
		switch data[p] {
		case 32:
			goto tr575
		case 43:
			goto tr576
		case 45:
			goto tr577
		case 47:
			goto tr578
		case 58:
			goto tr580
		case 65:
			goto tr581
		case 80:
			goto tr581
		case 90:
			goto tr582
		case 95:
			goto tr578
		case 97:
			goto tr583
		case 112:
			goto tr583
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st534
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr578
			}
		default:
			goto tr578
		}
		goto st0
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
		switch data[p] {
		case 32:
			goto tr638
		case 43:
			goto tr639
		case 45:
			goto tr640
		case 47:
			goto tr641
		case 58:
			goto tr643
		case 65:
			goto tr644
		case 80:
			goto tr644
		case 90:
			goto tr645
		case 95:
			goto tr641
		case 97:
			goto tr646
		case 112:
			goto tr646
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st535
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr641
			}
		default:
			goto tr641
		}
		goto st0
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		if 48 <= data[p] && data[p] <= 57 {
			goto st974
		}
		goto st0
	st974:
		if p++; p == pe {
			goto _test_eof974
		}
	st_case_974:
		switch data[p] {
		case 32:
			goto tr1464
		case 43:
			goto tr1426
		case 44:
			goto tr1465
		case 45:
			goto tr1428
		case 46:
			goto tr650
		case 47:
			goto tr1429
		case 84:
			goto tr1431
		case 90:
			goto tr1432
		case 95:
			goto tr1433
		case 116:
			goto tr1433
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st446
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1429
			}
		default:
			goto tr1429
		}
		goto st0
tr1464:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:35

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

	goto st975
	st975:
		if p++; p == pe {
			goto _test_eof975
		}
	st_case_975:
//line ragel/parse_datetime.go:21254
		switch data[p] {
		case 32:
			goto st482
		case 43:
			goto st483
		case 45:
			goto st489
		case 47:
			goto tr1437
		case 50:
			goto tr1467
		case 65:
			goto tr1438
		case 66:
			goto tr1439
		case 90:
			goto tr1440
		case 95:
			goto tr1437
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr1466
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1437
				}
			case data[p] >= 67:
				goto tr1437
			}
		default:
			goto tr1468
		}
		goto st0
tr1466:
//line ragel/datetime.rl:5
 pb = p 
	goto st976
	st976:
		if p++; p == pe {
			goto _test_eof976
		}
	st_case_976:
//line ragel/parse_datetime.go:21302
		switch data[p] {
		case 32:
			goto tr1180
		case 43:
			goto tr1181
		case 45:
			goto tr1182
		case 47:
			goto tr1183
		case 58:
			goto tr1185
		case 65:
			goto tr1186
		case 80:
			goto tr1186
		case 90:
			goto tr1187
		case 95:
			goto tr1183
		case 97:
			goto tr1188
		case 112:
			goto tr1188
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st977
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1183
			}
		default:
			goto tr1183
		}
		goto st0
	st977:
		if p++; p == pe {
			goto _test_eof977
		}
	st_case_977:
		switch data[p] {
		case 32:
			goto tr1196
		case 43:
			goto tr1197
		case 45:
			goto tr1198
		case 47:
			goto tr1199
		case 58:
			goto tr1200
		case 65:
			goto tr1201
		case 80:
			goto tr1201
		case 90:
			goto tr1202
		case 95:
			goto tr1199
		case 97:
			goto tr1203
		case 112:
			goto tr1203
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st536
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1199
			}
		default:
			goto tr1199
		}
		goto st0
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		if 48 <= data[p] && data[p] <= 57 {
			goto st978
		}
		goto st0
	st978:
		if p++; p == pe {
			goto _test_eof978
		}
	st_case_978:
		switch data[p] {
		case 32:
			goto tr1204
		case 43:
			goto tr1205
		case 45:
			goto tr1207
		case 47:
			goto tr1208
		case 90:
			goto tr1210
		case 95:
			goto tr1208
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1206
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1208
				}
			case data[p] >= 65:
				goto tr1208
			}
		default:
			goto st32
		}
		goto st0
tr1467:
//line ragel/datetime.rl:5
 pb = p 
	goto st979
	st979:
		if p++; p == pe {
			goto _test_eof979
		}
	st_case_979:
//line ragel/parse_datetime.go:21437
		switch data[p] {
		case 32:
			goto tr1180
		case 43:
			goto tr1181
		case 45:
			goto tr1182
		case 47:
			goto tr1183
		case 58:
			goto tr1185
		case 65:
			goto tr1186
		case 80:
			goto tr1186
		case 90:
			goto tr1187
		case 95:
			goto tr1183
		case 97:
			goto tr1188
		case 112:
			goto tr1188
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st977
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1183
				}
			case data[p] >= 66:
				goto tr1183
			}
		default:
			goto st537
		}
		goto st0
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		if 48 <= data[p] && data[p] <= 57 {
			goto st536
		}
		goto st0
tr1468:
//line ragel/datetime.rl:5
 pb = p 
	goto st980
	st980:
		if p++; p == pe {
			goto _test_eof980
		}
	st_case_980:
//line ragel/parse_datetime.go:21498
		switch data[p] {
		case 32:
			goto tr1180
		case 43:
			goto tr1181
		case 45:
			goto tr1182
		case 47:
			goto tr1183
		case 58:
			goto tr1185
		case 65:
			goto tr1186
		case 80:
			goto tr1186
		case 90:
			goto tr1187
		case 95:
			goto tr1183
		case 97:
			goto tr1188
		case 112:
			goto tr1188
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st537
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1183
			}
		default:
			goto tr1183
		}
		goto st0
tr1465:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:35

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

	goto st538
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
//line ragel/parse_datetime.go:21562
		if data[p] == 32 {
			goto st38
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr655
		}
		goto st0
tr814:
//line ragel/datetime.rl:5
 pb = p 
	goto st539
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
//line ragel/parse_datetime.go:21579
		switch data[p] {
		case 32:
			goto tr575
		case 43:
			goto tr576
		case 45:
			goto tr577
		case 47:
			goto tr578
		case 58:
			goto tr580
		case 65:
			goto tr581
		case 80:
			goto tr581
		case 90:
			goto tr582
		case 95:
			goto tr578
		case 97:
			goto tr583
		case 112:
			goto tr583
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st534
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr578
				}
			case data[p] >= 66:
				goto tr578
			}
		default:
			goto st540
		}
		goto st0
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
		if 48 <= data[p] && data[p] <= 57 {
			goto st535
		}
		goto st0
tr815:
//line ragel/datetime.rl:5
 pb = p 
	goto st541
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
//line ragel/parse_datetime.go:21640
		switch data[p] {
		case 32:
			goto tr575
		case 43:
			goto tr576
		case 45:
			goto tr577
		case 47:
			goto tr578
		case 58:
			goto tr580
		case 65:
			goto tr581
		case 80:
			goto tr581
		case 90:
			goto tr582
		case 95:
			goto tr578
		case 97:
			goto tr583
		case 112:
			goto tr583
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st540
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr578
			}
		default:
			goto tr578
		}
		goto st0
tr564:
//line ragel/datetime.rl:5
 pb = p 
	goto st542
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
//line ragel/parse_datetime.go:21687
		if data[p] == 32 {
			goto tr812
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st472
			}
		case data[p] >= 45:
			goto tr385
		}
		goto st0
tr565:
//line ragel/datetime.rl:5
 pb = p 
	goto st543
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
//line ragel/parse_datetime.go:21709
		if data[p] == 32 {
			goto tr812
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr385
		}
		goto st0
tr559:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st544
tr826:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st544
tr833:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st544
tr842:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st544
tr853:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st544
tr862:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st544
tr866:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st544
tr873:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st544
tr878:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st544
tr883:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st544
tr893:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st544
tr902:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st544
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
//line ragel/parse_datetime.go:21770
		switch data[p] {
		case 32:
			goto st392
		case 48:
			goto tr502
		case 51:
			goto tr504
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st353
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr505
			}
		default:
			goto tr503
		}
		goto st0
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
		if data[p] == 108 {
			goto st546
		}
		goto st0
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
		switch data[p] {
		case 32:
			goto tr558
		case 46:
			goto tr559
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr375
		}
		goto st0
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
		if data[p] == 103 {
			goto st548
		}
		goto st0
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
		switch data[p] {
		case 32:
			goto tr825
		case 46:
			goto tr826
		case 117:
			goto st549
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr512
		}
		goto st0
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
		if data[p] == 115 {
			goto st550
		}
		goto st0
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
		if data[p] == 116 {
			goto st551
		}
		goto st0
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
		switch data[p] {
		case 32:
			goto tr825
		case 46:
			goto tr826
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr512
		}
		goto st0
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
		if data[p] == 101 {
			goto st553
		}
		goto st0
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
		if data[p] == 99 {
			goto st554
		}
		goto st0
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
		switch data[p] {
		case 32:
			goto tr832
		case 46:
			goto tr833
		case 101:
			goto st555
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr520
		}
		goto st0
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
		if data[p] == 109 {
			goto st556
		}
		goto st0
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
		if data[p] == 98 {
			goto st557
		}
		goto st0
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
		if data[p] == 101 {
			goto st558
		}
		goto st0
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
		if data[p] == 114 {
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
		case 46:
			goto tr833
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr520
		}
		goto st0
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
		if data[p] == 101 {
			goto st561
		}
		goto st0
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
		if data[p] == 98 {
			goto st562
		}
		goto st0
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
		switch data[p] {
		case 32:
			goto tr841
		case 46:
			goto tr842
		case 114:
			goto st563
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr531
		}
		goto st0
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
		if data[p] == 117 {
			goto st564
		}
		goto st0
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
		if data[p] == 97 {
			goto st565
		}
		goto st0
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
		if data[p] == 114 {
			goto st566
		}
		goto st0
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
		if data[p] == 121 {
			goto st567
		}
		goto st0
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		switch data[p] {
		case 32:
			goto tr841
		case 46:
			goto tr842
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr531
		}
		goto st0
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
		switch data[p] {
		case 97:
			goto st569
		case 117:
			goto st575
		}
		goto st0
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		if data[p] == 110 {
			goto st570
		}
		goto st0
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		switch data[p] {
		case 32:
			goto tr851
		case 46:
			goto tr853
		case 117:
			goto st571
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr852
		}
		goto st0
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
		if data[p] == 97 {
			goto st572
		}
		goto st0
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
		if data[p] == 114 {
			goto st573
		}
		goto st0
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		if data[p] == 121 {
			goto st574
		}
		goto st0
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
		switch data[p] {
		case 32:
			goto tr851
		case 46:
			goto tr853
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr852
		}
		goto st0
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
		switch data[p] {
		case 108:
			goto st576
		case 110:
			goto st578
		}
		goto st0
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
		switch data[p] {
		case 32:
			goto tr860
		case 46:
			goto tr862
		case 121:
			goto st577
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr861
		}
		goto st0
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
		switch data[p] {
		case 32:
			goto tr860
		case 46:
			goto tr862
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr861
		}
		goto st0
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
		switch data[p] {
		case 32:
			goto tr864
		case 46:
			goto tr866
		case 101:
			goto st579
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr865
		}
		goto st0
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
		switch data[p] {
		case 32:
			goto tr864
		case 46:
			goto tr866
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr865
		}
		goto st0
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		if data[p] == 97 {
			goto st581
		}
		goto st0
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
		switch data[p] {
		case 114:
			goto st582
		case 121:
			goto st585
		}
		goto st0
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		switch data[p] {
		case 32:
			goto tr871
		case 46:
			goto tr873
		case 99:
			goto st583
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr872
		}
		goto st0
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		if data[p] == 104 {
			goto st584
		}
		goto st0
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		switch data[p] {
		case 32:
			goto tr871
		case 46:
			goto tr873
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr872
		}
		goto st0
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		switch data[p] {
		case 32:
			goto tr876
		case 46:
			goto tr878
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr877
		}
		goto st0
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		if data[p] == 111 {
			goto st587
		}
		goto st0
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
		if data[p] == 118 {
			goto st588
		}
		goto st0
	st588:
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
		switch data[p] {
		case 32:
			goto tr881
		case 46:
			goto tr883
		case 101:
			goto st589
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr882
		}
		goto st0
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		if data[p] == 109 {
			goto st590
		}
		goto st0
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
		if data[p] == 98 {
			goto st591
		}
		goto st0
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		if data[p] == 101 {
			goto st592
		}
		goto st0
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
		if data[p] == 114 {
			goto st593
		}
		goto st0
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
	st_case_593:
		switch data[p] {
		case 32:
			goto tr881
		case 46:
			goto tr883
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr882
		}
		goto st0
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		if data[p] == 99 {
			goto st595
		}
		goto st0
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		if data[p] == 116 {
			goto st596
		}
		goto st0
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		switch data[p] {
		case 32:
			goto tr891
		case 46:
			goto tr893
		case 111:
			goto st597
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr892
		}
		goto st0
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
	st_case_597:
		if data[p] == 98 {
			goto st598
		}
		goto st0
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		if data[p] == 101 {
			goto st599
		}
		goto st0
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		if data[p] == 114 {
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
			goto tr891
		case 46:
			goto tr893
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr892
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
		if data[p] == 112 {
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
			goto tr900
		case 46:
			goto tr902
		case 116:
			goto st604
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr901
		}
		goto st0
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		switch data[p] {
		case 32:
			goto tr900
		case 46:
			goto tr902
		case 101:
			goto st605
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr901
		}
		goto st0
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
		if data[p] == 109 {
			goto st606
		}
		goto st0
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
		if data[p] == 98 {
			goto st607
		}
		goto st0
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
		if data[p] == 101 {
			goto st608
		}
		goto st0
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		if data[p] == 114 {
			goto st609
		}
		goto st0
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
		switch data[p] {
		case 32:
			goto tr900
		case 46:
			goto tr902
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr901
		}
		goto st0
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		if data[p] == 32 {
			goto st611
		}
		goto st0
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		switch data[p] {
		case 48:
			goto tr910
		case 51:
			goto tr912
		case 65:
			goto st686
		case 68:
			goto st697
		case 70:
			goto st705
		case 74:
			goto st713
		case 77:
			goto st725
		case 78:
			goto st731
		case 79:
			goto st739
		case 83:
			goto st746
		case 97:
			goto st686
		case 100:
			goto st697
		case 102:
			goto st705
		case 106:
			goto st713
		case 109:
			goto st725
		case 110:
			goto st731
		case 111:
			goto st739
		case 115:
			goto st746
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr913
			}
		case data[p] >= 49:
			goto tr911
		}
		goto st0
tr910:
//line ragel/datetime.rl:5
 pb = p 
	goto st612
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
//line ragel/parse_datetime.go:22616
		if 49 <= data[p] && data[p] <= 57 {
			goto st613
		}
		goto st0
tr913:
//line ragel/datetime.rl:5
 pb = p 
	goto st613
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
//line ragel/parse_datetime.go:22630
		switch data[p] {
		case 32:
			goto tr178
		case 45:
			goto tr923
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto tr178
		}
		goto st0
tr923:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st614
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
//line ragel/parse_datetime.go:22657
		switch data[p] {
		case 65:
			goto st615
		case 68:
			goto st626
		case 70:
			goto st634
		case 74:
			goto st642
		case 77:
			goto st654
		case 78:
			goto st660
		case 79:
			goto st668
		case 83:
			goto st675
		case 97:
			goto st615
		case 100:
			goto st626
		case 102:
			goto st634
		case 106:
			goto st642
		case 109:
			goto st654
		case 110:
			goto st660
		case 111:
			goto st668
		case 115:
			goto st675
		}
		goto st0
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
		switch data[p] {
		case 112:
			goto st616
		case 117:
			goto st621
		}
		goto st0
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
		if data[p] == 114 {
			goto st617
		}
		goto st0
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
		switch data[p] {
		case 32:
			goto tr191
		case 45:
			goto tr190
		case 46:
			goto tr935
		case 47:
			goto tr191
		case 105:
			goto st619
		}
		goto st0
tr935:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st618
tr939:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st618
tr945:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st618
tr953:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st618
tr962:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st618
tr969:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st618
tr971:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st618
tr976:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st618
tr979:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st618
tr982:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st618
tr990:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st618
tr997:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st618
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
//line ragel/parse_datetime.go:22785
		switch data[p] {
		case 32:
			goto st126
		case 45:
			goto st123
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr197
			}
		case data[p] >= 46:
			goto st126
		}
		goto st0
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		if data[p] == 108 {
			goto st620
		}
		goto st0
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		switch data[p] {
		case 32:
			goto tr191
		case 45:
			goto tr190
		case 46:
			goto tr935
		case 47:
			goto tr191
		}
		goto st0
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		if data[p] == 103 {
			goto st622
		}
		goto st0
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		switch data[p] {
		case 32:
			goto tr205
		case 45:
			goto tr204
		case 46:
			goto tr939
		case 47:
			goto tr205
		case 117:
			goto st623
		}
		goto st0
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
		if data[p] == 115 {
			goto st624
		}
		goto st0
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
		if data[p] == 116 {
			goto st625
		}
		goto st0
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
		switch data[p] {
		case 32:
			goto tr205
		case 45:
			goto tr204
		case 46:
			goto tr939
		case 47:
			goto tr205
		}
		goto st0
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
		if data[p] == 101 {
			goto st627
		}
		goto st0
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
		if data[p] == 99 {
			goto st628
		}
		goto st0
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
		switch data[p] {
		case 32:
			goto tr213
		case 45:
			goto tr212
		case 46:
			goto tr945
		case 47:
			goto tr213
		case 101:
			goto st629
		}
		goto st0
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
		if data[p] == 109 {
			goto st630
		}
		goto st0
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
		if data[p] == 98 {
			goto st631
		}
		goto st0
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		if data[p] == 101 {
			goto st632
		}
		goto st0
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
		if data[p] == 114 {
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
			goto tr213
		case 45:
			goto tr212
		case 46:
			goto tr945
		case 47:
			goto tr213
		}
		goto st0
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
		if data[p] == 101 {
			goto st635
		}
		goto st0
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
		if data[p] == 98 {
			goto st636
		}
		goto st0
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
		switch data[p] {
		case 32:
			goto tr223
		case 45:
			goto tr222
		case 46:
			goto tr953
		case 47:
			goto tr223
		case 114:
			goto st637
		}
		goto st0
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
		if data[p] == 117 {
			goto st638
		}
		goto st0
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
		if data[p] == 97 {
			goto st639
		}
		goto st0
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
		if data[p] == 114 {
			goto st640
		}
		goto st0
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
		if data[p] == 121 {
			goto st641
		}
		goto st0
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
		switch data[p] {
		case 32:
			goto tr223
		case 45:
			goto tr222
		case 46:
			goto tr953
		case 47:
			goto tr223
		}
		goto st0
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
		switch data[p] {
		case 97:
			goto st643
		case 117:
			goto st649
		}
		goto st0
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
		if data[p] == 110 {
			goto st644
		}
		goto st0
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
		switch data[p] {
		case 32:
			goto tr234
		case 45:
			goto tr233
		case 46:
			goto tr962
		case 47:
			goto tr234
		case 117:
			goto st645
		}
		goto st0
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
		if data[p] == 97 {
			goto st646
		}
		goto st0
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
		if data[p] == 114 {
			goto st647
		}
		goto st0
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
		if data[p] == 121 {
			goto st648
		}
		goto st0
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
		switch data[p] {
		case 32:
			goto tr234
		case 45:
			goto tr233
		case 46:
			goto tr962
		case 47:
			goto tr234
		}
		goto st0
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
		switch data[p] {
		case 108:
			goto st650
		case 110:
			goto st652
		}
		goto st0
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
		switch data[p] {
		case 32:
			goto tr243
		case 45:
			goto tr242
		case 46:
			goto tr969
		case 47:
			goto tr243
		case 121:
			goto st651
		}
		goto st0
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		switch data[p] {
		case 32:
			goto tr243
		case 45:
			goto tr242
		case 46:
			goto tr969
		case 47:
			goto tr243
		}
		goto st0
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
		switch data[p] {
		case 32:
			goto tr247
		case 45:
			goto tr246
		case 46:
			goto tr971
		case 47:
			goto tr247
		case 101:
			goto st653
		}
		goto st0
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		switch data[p] {
		case 32:
			goto tr247
		case 45:
			goto tr246
		case 46:
			goto tr971
		case 47:
			goto tr247
		}
		goto st0
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
		if data[p] == 97 {
			goto st655
		}
		goto st0
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
		switch data[p] {
		case 114:
			goto st656
		case 121:
			goto st659
		}
		goto st0
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
		switch data[p] {
		case 32:
			goto tr254
		case 45:
			goto tr253
		case 46:
			goto tr976
		case 47:
			goto tr254
		case 99:
			goto st657
		}
		goto st0
	st657:
		if p++; p == pe {
			goto _test_eof657
		}
	st_case_657:
		if data[p] == 104 {
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
			goto tr254
		case 45:
			goto tr253
		case 46:
			goto tr976
		case 47:
			goto tr254
		}
		goto st0
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
		switch data[p] {
		case 32:
			goto tr259
		case 45:
			goto tr258
		case 46:
			goto tr979
		case 47:
			goto tr259
		}
		goto st0
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
		if data[p] == 111 {
			goto st661
		}
		goto st0
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
	st_case_661:
		if data[p] == 118 {
			goto st662
		}
		goto st0
	st662:
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
		switch data[p] {
		case 32:
			goto tr264
		case 45:
			goto tr263
		case 46:
			goto tr982
		case 47:
			goto tr264
		case 101:
			goto st663
		}
		goto st0
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
		if data[p] == 109 {
			goto st664
		}
		goto st0
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
		if data[p] == 98 {
			goto st665
		}
		goto st0
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
		if data[p] == 101 {
			goto st666
		}
		goto st0
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
		if data[p] == 114 {
			goto st667
		}
		goto st0
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
	st_case_667:
		switch data[p] {
		case 32:
			goto tr264
		case 45:
			goto tr263
		case 46:
			goto tr982
		case 47:
			goto tr264
		}
		goto st0
	st668:
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
		if data[p] == 99 {
			goto st669
		}
		goto st0
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
		if data[p] == 116 {
			goto st670
		}
		goto st0
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
		switch data[p] {
		case 32:
			goto tr274
		case 45:
			goto tr273
		case 46:
			goto tr990
		case 47:
			goto tr274
		case 111:
			goto st671
		}
		goto st0
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
		if data[p] == 98 {
			goto st672
		}
		goto st0
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
	st_case_672:
		if data[p] == 101 {
			goto st673
		}
		goto st0
	st673:
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
		if data[p] == 114 {
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
			goto tr274
		case 45:
			goto tr273
		case 46:
			goto tr990
		case 47:
			goto tr274
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
		if data[p] == 112 {
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
			goto tr283
		case 45:
			goto tr282
		case 46:
			goto tr997
		case 47:
			goto tr283
		case 116:
			goto st678
		}
		goto st0
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
		switch data[p] {
		case 32:
			goto tr283
		case 45:
			goto tr282
		case 46:
			goto tr997
		case 47:
			goto tr283
		case 101:
			goto st679
		}
		goto st0
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
		if data[p] == 109 {
			goto st680
		}
		goto st0
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
		if data[p] == 98 {
			goto st681
		}
		goto st0
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
	st_case_681:
		if data[p] == 101 {
			goto st682
		}
		goto st0
	st682:
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
		if data[p] == 114 {
			goto st683
		}
		goto st0
	st683:
		if p++; p == pe {
			goto _test_eof683
		}
	st_case_683:
		switch data[p] {
		case 32:
			goto tr283
		case 45:
			goto tr282
		case 46:
			goto tr997
		case 47:
			goto tr283
		}
		goto st0
tr911:
//line ragel/datetime.rl:5
 pb = p 
	goto st684
	st684:
		if p++; p == pe {
			goto _test_eof684
		}
	st_case_684:
//line ragel/parse_datetime.go:23587
		switch data[p] {
		case 32:
			goto tr178
		case 45:
			goto tr923
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st613
			}
		case data[p] >= 46:
			goto tr178
		}
		goto st0
tr912:
//line ragel/datetime.rl:5
 pb = p 
	goto st685
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
	st_case_685:
//line ragel/parse_datetime.go:23612
		switch data[p] {
		case 32:
			goto tr178
		case 45:
			goto tr923
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st613
			}
		case data[p] >= 46:
			goto tr178
		}
		goto st0
	st686:
		if p++; p == pe {
			goto _test_eof686
		}
	st_case_686:
		switch data[p] {
		case 112:
			goto st687
		case 117:
			goto st692
		}
		goto st0
	st687:
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
		if data[p] == 114 {
			goto st688
		}
		goto st0
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
		switch data[p] {
		case 32:
			goto tr375
		case 46:
			goto tr1007
		case 105:
			goto st690
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr375
		}
		goto st0
tr1007:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st689
tr1011:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st689
tr1017:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st689
tr1025:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st689
tr1034:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st689
tr1041:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st689
tr1043:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st689
tr1048:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st689
tr1051:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st689
tr1054:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st689
tr1062:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st689
tr1069:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st689
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
//line ragel/parse_datetime.go:23719
		switch data[p] {
		case 32:
			goto st353
		case 48:
			goto tr502
		case 51:
			goto tr504
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st353
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr505
			}
		default:
			goto tr503
		}
		goto st0
	st690:
		if p++; p == pe {
			goto _test_eof690
		}
	st_case_690:
		if data[p] == 108 {
			goto st691
		}
		goto st0
	st691:
		if p++; p == pe {
			goto _test_eof691
		}
	st_case_691:
		switch data[p] {
		case 32:
			goto tr375
		case 46:
			goto tr1007
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr375
		}
		goto st0
	st692:
		if p++; p == pe {
			goto _test_eof692
		}
	st_case_692:
		if data[p] == 103 {
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
			goto tr512
		case 46:
			goto tr1011
		case 117:
			goto st694
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr512
		}
		goto st0
	st694:
		if p++; p == pe {
			goto _test_eof694
		}
	st_case_694:
		if data[p] == 115 {
			goto st695
		}
		goto st0
	st695:
		if p++; p == pe {
			goto _test_eof695
		}
	st_case_695:
		if data[p] == 116 {
			goto st696
		}
		goto st0
	st696:
		if p++; p == pe {
			goto _test_eof696
		}
	st_case_696:
		switch data[p] {
		case 32:
			goto tr512
		case 46:
			goto tr1011
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr512
		}
		goto st0
	st697:
		if p++; p == pe {
			goto _test_eof697
		}
	st_case_697:
		if data[p] == 101 {
			goto st698
		}
		goto st0
	st698:
		if p++; p == pe {
			goto _test_eof698
		}
	st_case_698:
		if data[p] == 99 {
			goto st699
		}
		goto st0
	st699:
		if p++; p == pe {
			goto _test_eof699
		}
	st_case_699:
		switch data[p] {
		case 32:
			goto tr520
		case 46:
			goto tr1017
		case 101:
			goto st700
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr520
		}
		goto st0
	st700:
		if p++; p == pe {
			goto _test_eof700
		}
	st_case_700:
		if data[p] == 109 {
			goto st701
		}
		goto st0
	st701:
		if p++; p == pe {
			goto _test_eof701
		}
	st_case_701:
		if data[p] == 98 {
			goto st702
		}
		goto st0
	st702:
		if p++; p == pe {
			goto _test_eof702
		}
	st_case_702:
		if data[p] == 101 {
			goto st703
		}
		goto st0
	st703:
		if p++; p == pe {
			goto _test_eof703
		}
	st_case_703:
		if data[p] == 114 {
			goto st704
		}
		goto st0
	st704:
		if p++; p == pe {
			goto _test_eof704
		}
	st_case_704:
		switch data[p] {
		case 32:
			goto tr520
		case 46:
			goto tr1017
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr520
		}
		goto st0
	st705:
		if p++; p == pe {
			goto _test_eof705
		}
	st_case_705:
		if data[p] == 101 {
			goto st706
		}
		goto st0
	st706:
		if p++; p == pe {
			goto _test_eof706
		}
	st_case_706:
		if data[p] == 98 {
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
			goto tr531
		case 46:
			goto tr1025
		case 114:
			goto st708
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr531
		}
		goto st0
	st708:
		if p++; p == pe {
			goto _test_eof708
		}
	st_case_708:
		if data[p] == 117 {
			goto st709
		}
		goto st0
	st709:
		if p++; p == pe {
			goto _test_eof709
		}
	st_case_709:
		if data[p] == 97 {
			goto st710
		}
		goto st0
	st710:
		if p++; p == pe {
			goto _test_eof710
		}
	st_case_710:
		if data[p] == 114 {
			goto st711
		}
		goto st0
	st711:
		if p++; p == pe {
			goto _test_eof711
		}
	st_case_711:
		if data[p] == 121 {
			goto st712
		}
		goto st0
	st712:
		if p++; p == pe {
			goto _test_eof712
		}
	st_case_712:
		switch data[p] {
		case 32:
			goto tr531
		case 46:
			goto tr1025
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr531
		}
		goto st0
	st713:
		if p++; p == pe {
			goto _test_eof713
		}
	st_case_713:
		switch data[p] {
		case 97:
			goto st714
		case 117:
			goto st720
		}
		goto st0
	st714:
		if p++; p == pe {
			goto _test_eof714
		}
	st_case_714:
		if data[p] == 110 {
			goto st715
		}
		goto st0
	st715:
		if p++; p == pe {
			goto _test_eof715
		}
	st_case_715:
		switch data[p] {
		case 32:
			goto tr852
		case 46:
			goto tr1034
		case 117:
			goto st716
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr852
		}
		goto st0
	st716:
		if p++; p == pe {
			goto _test_eof716
		}
	st_case_716:
		if data[p] == 97 {
			goto st717
		}
		goto st0
	st717:
		if p++; p == pe {
			goto _test_eof717
		}
	st_case_717:
		if data[p] == 114 {
			goto st718
		}
		goto st0
	st718:
		if p++; p == pe {
			goto _test_eof718
		}
	st_case_718:
		if data[p] == 121 {
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
			goto tr852
		case 46:
			goto tr1034
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr852
		}
		goto st0
	st720:
		if p++; p == pe {
			goto _test_eof720
		}
	st_case_720:
		switch data[p] {
		case 108:
			goto st721
		case 110:
			goto st723
		}
		goto st0
	st721:
		if p++; p == pe {
			goto _test_eof721
		}
	st_case_721:
		switch data[p] {
		case 32:
			goto tr861
		case 46:
			goto tr1041
		case 121:
			goto st722
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr861
		}
		goto st0
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
		switch data[p] {
		case 32:
			goto tr861
		case 46:
			goto tr1041
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr861
		}
		goto st0
	st723:
		if p++; p == pe {
			goto _test_eof723
		}
	st_case_723:
		switch data[p] {
		case 32:
			goto tr865
		case 46:
			goto tr1043
		case 101:
			goto st724
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr865
		}
		goto st0
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
	st_case_724:
		switch data[p] {
		case 32:
			goto tr865
		case 46:
			goto tr1043
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr865
		}
		goto st0
	st725:
		if p++; p == pe {
			goto _test_eof725
		}
	st_case_725:
		if data[p] == 97 {
			goto st726
		}
		goto st0
	st726:
		if p++; p == pe {
			goto _test_eof726
		}
	st_case_726:
		switch data[p] {
		case 114:
			goto st727
		case 121:
			goto st730
		}
		goto st0
	st727:
		if p++; p == pe {
			goto _test_eof727
		}
	st_case_727:
		switch data[p] {
		case 32:
			goto tr872
		case 46:
			goto tr1048
		case 99:
			goto st728
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr872
		}
		goto st0
	st728:
		if p++; p == pe {
			goto _test_eof728
		}
	st_case_728:
		if data[p] == 104 {
			goto st729
		}
		goto st0
	st729:
		if p++; p == pe {
			goto _test_eof729
		}
	st_case_729:
		switch data[p] {
		case 32:
			goto tr872
		case 46:
			goto tr1048
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr872
		}
		goto st0
	st730:
		if p++; p == pe {
			goto _test_eof730
		}
	st_case_730:
		switch data[p] {
		case 32:
			goto tr877
		case 46:
			goto tr1051
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr877
		}
		goto st0
	st731:
		if p++; p == pe {
			goto _test_eof731
		}
	st_case_731:
		if data[p] == 111 {
			goto st732
		}
		goto st0
	st732:
		if p++; p == pe {
			goto _test_eof732
		}
	st_case_732:
		if data[p] == 118 {
			goto st733
		}
		goto st0
	st733:
		if p++; p == pe {
			goto _test_eof733
		}
	st_case_733:
		switch data[p] {
		case 32:
			goto tr882
		case 46:
			goto tr1054
		case 101:
			goto st734
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr882
		}
		goto st0
	st734:
		if p++; p == pe {
			goto _test_eof734
		}
	st_case_734:
		if data[p] == 109 {
			goto st735
		}
		goto st0
	st735:
		if p++; p == pe {
			goto _test_eof735
		}
	st_case_735:
		if data[p] == 98 {
			goto st736
		}
		goto st0
	st736:
		if p++; p == pe {
			goto _test_eof736
		}
	st_case_736:
		if data[p] == 101 {
			goto st737
		}
		goto st0
	st737:
		if p++; p == pe {
			goto _test_eof737
		}
	st_case_737:
		if data[p] == 114 {
			goto st738
		}
		goto st0
	st738:
		if p++; p == pe {
			goto _test_eof738
		}
	st_case_738:
		switch data[p] {
		case 32:
			goto tr882
		case 46:
			goto tr1054
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr882
		}
		goto st0
	st739:
		if p++; p == pe {
			goto _test_eof739
		}
	st_case_739:
		if data[p] == 99 {
			goto st740
		}
		goto st0
	st740:
		if p++; p == pe {
			goto _test_eof740
		}
	st_case_740:
		if data[p] == 116 {
			goto st741
		}
		goto st0
	st741:
		if p++; p == pe {
			goto _test_eof741
		}
	st_case_741:
		switch data[p] {
		case 32:
			goto tr892
		case 46:
			goto tr1062
		case 111:
			goto st742
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr892
		}
		goto st0
	st742:
		if p++; p == pe {
			goto _test_eof742
		}
	st_case_742:
		if data[p] == 98 {
			goto st743
		}
		goto st0
	st743:
		if p++; p == pe {
			goto _test_eof743
		}
	st_case_743:
		if data[p] == 101 {
			goto st744
		}
		goto st0
	st744:
		if p++; p == pe {
			goto _test_eof744
		}
	st_case_744:
		if data[p] == 114 {
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
			goto tr892
		case 46:
			goto tr1062
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr892
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
		if data[p] == 112 {
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
			goto tr901
		case 46:
			goto tr1069
		case 116:
			goto st749
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr901
		}
		goto st0
	st749:
		if p++; p == pe {
			goto _test_eof749
		}
	st_case_749:
		switch data[p] {
		case 32:
			goto tr901
		case 46:
			goto tr1069
		case 101:
			goto st750
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr901
		}
		goto st0
	st750:
		if p++; p == pe {
			goto _test_eof750
		}
	st_case_750:
		if data[p] == 109 {
			goto st751
		}
		goto st0
	st751:
		if p++; p == pe {
			goto _test_eof751
		}
	st_case_751:
		if data[p] == 98 {
			goto st752
		}
		goto st0
	st752:
		if p++; p == pe {
			goto _test_eof752
		}
	st_case_752:
		if data[p] == 101 {
			goto st753
		}
		goto st0
	st753:
		if p++; p == pe {
			goto _test_eof753
		}
	st_case_753:
		if data[p] == 114 {
			goto st754
		}
		goto st0
	st754:
		if p++; p == pe {
			goto _test_eof754
		}
	st_case_754:
		switch data[p] {
		case 32:
			goto tr901
		case 46:
			goto tr1069
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr901
		}
		goto st0
	st755:
		if p++; p == pe {
			goto _test_eof755
		}
	st_case_755:
		if data[p] == 97 {
			goto st756
		}
		goto st0
	st756:
		if p++; p == pe {
			goto _test_eof756
		}
	st_case_756:
		if data[p] == 121 {
			goto st757
		}
		goto st0
	st757:
		if p++; p == pe {
			goto _test_eof757
		}
	st_case_757:
		switch data[p] {
		case 32:
			goto st384
		case 44:
			goto st610
		}
		goto st0
	st758:
		if p++; p == pe {
			goto _test_eof758
		}
	st_case_758:
		switch data[p] {
		case 97:
			goto st759
		case 117:
			goto st765
		}
		goto st0
	st759:
		if p++; p == pe {
			goto _test_eof759
		}
	st_case_759:
		if data[p] == 110 {
			goto st760
		}
		goto st0
	st760:
		if p++; p == pe {
			goto _test_eof760
		}
	st_case_760:
		switch data[p] {
		case 32:
			goto tr1081
		case 46:
			goto tr1082
		case 117:
			goto st761
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr852
		}
		goto st0
	st761:
		if p++; p == pe {
			goto _test_eof761
		}
	st_case_761:
		if data[p] == 97 {
			goto st762
		}
		goto st0
	st762:
		if p++; p == pe {
			goto _test_eof762
		}
	st_case_762:
		if data[p] == 114 {
			goto st763
		}
		goto st0
	st763:
		if p++; p == pe {
			goto _test_eof763
		}
	st_case_763:
		if data[p] == 121 {
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
			goto tr1081
		case 46:
			goto tr1082
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr852
		}
		goto st0
	st765:
		if p++; p == pe {
			goto _test_eof765
		}
	st_case_765:
		switch data[p] {
		case 108:
			goto st766
		case 110:
			goto st768
		}
		goto st0
	st766:
		if p++; p == pe {
			goto _test_eof766
		}
	st_case_766:
		switch data[p] {
		case 32:
			goto tr1089
		case 46:
			goto tr1090
		case 121:
			goto st767
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr861
		}
		goto st0
	st767:
		if p++; p == pe {
			goto _test_eof767
		}
	st_case_767:
		switch data[p] {
		case 32:
			goto tr1089
		case 46:
			goto tr1090
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr861
		}
		goto st0
	st768:
		if p++; p == pe {
			goto _test_eof768
		}
	st_case_768:
		switch data[p] {
		case 32:
			goto tr1092
		case 46:
			goto tr1093
		case 101:
			goto st769
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr865
		}
		goto st0
	st769:
		if p++; p == pe {
			goto _test_eof769
		}
	st_case_769:
		switch data[p] {
		case 32:
			goto tr1092
		case 46:
			goto tr1093
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr865
		}
		goto st0
	st770:
		if p++; p == pe {
			goto _test_eof770
		}
	st_case_770:
		switch data[p] {
		case 97:
			goto st771
		case 111:
			goto st776
		}
		goto st0
	st771:
		if p++; p == pe {
			goto _test_eof771
		}
	st_case_771:
		switch data[p] {
		case 114:
			goto st772
		case 121:
			goto st775
		}
		goto st0
	st772:
		if p++; p == pe {
			goto _test_eof772
		}
	st_case_772:
		switch data[p] {
		case 32:
			goto tr1099
		case 46:
			goto tr1100
		case 99:
			goto st773
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr872
		}
		goto st0
	st773:
		if p++; p == pe {
			goto _test_eof773
		}
	st_case_773:
		if data[p] == 104 {
			goto st774
		}
		goto st0
	st774:
		if p++; p == pe {
			goto _test_eof774
		}
	st_case_774:
		switch data[p] {
		case 32:
			goto tr1099
		case 46:
			goto tr1100
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr872
		}
		goto st0
	st775:
		if p++; p == pe {
			goto _test_eof775
		}
	st_case_775:
		switch data[p] {
		case 32:
			goto tr1103
		case 46:
			goto tr1104
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr877
		}
		goto st0
	st776:
		if p++; p == pe {
			goto _test_eof776
		}
	st_case_776:
		if data[p] == 110 {
			goto st383
		}
		goto st0
	st777:
		if p++; p == pe {
			goto _test_eof777
		}
	st_case_777:
		if data[p] == 111 {
			goto st778
		}
		goto st0
	st778:
		if p++; p == pe {
			goto _test_eof778
		}
	st_case_778:
		if data[p] == 118 {
			goto st779
		}
		goto st0
	st779:
		if p++; p == pe {
			goto _test_eof779
		}
	st_case_779:
		switch data[p] {
		case 32:
			goto tr1107
		case 46:
			goto tr1108
		case 101:
			goto st780
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr882
		}
		goto st0
	st780:
		if p++; p == pe {
			goto _test_eof780
		}
	st_case_780:
		if data[p] == 109 {
			goto st781
		}
		goto st0
	st781:
		if p++; p == pe {
			goto _test_eof781
		}
	st_case_781:
		if data[p] == 98 {
			goto st782
		}
		goto st0
	st782:
		if p++; p == pe {
			goto _test_eof782
		}
	st_case_782:
		if data[p] == 101 {
			goto st783
		}
		goto st0
	st783:
		if p++; p == pe {
			goto _test_eof783
		}
	st_case_783:
		if data[p] == 114 {
			goto st784
		}
		goto st0
	st784:
		if p++; p == pe {
			goto _test_eof784
		}
	st_case_784:
		switch data[p] {
		case 32:
			goto tr1107
		case 46:
			goto tr1108
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr882
		}
		goto st0
	st785:
		if p++; p == pe {
			goto _test_eof785
		}
	st_case_785:
		if data[p] == 99 {
			goto st786
		}
		goto st0
	st786:
		if p++; p == pe {
			goto _test_eof786
		}
	st_case_786:
		if data[p] == 116 {
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
			goto tr1116
		case 46:
			goto tr1117
		case 111:
			goto st788
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr892
		}
		goto st0
	st788:
		if p++; p == pe {
			goto _test_eof788
		}
	st_case_788:
		if data[p] == 98 {
			goto st789
		}
		goto st0
	st789:
		if p++; p == pe {
			goto _test_eof789
		}
	st_case_789:
		if data[p] == 101 {
			goto st790
		}
		goto st0
	st790:
		if p++; p == pe {
			goto _test_eof790
		}
	st_case_790:
		if data[p] == 114 {
			goto st791
		}
		goto st0
	st791:
		if p++; p == pe {
			goto _test_eof791
		}
	st_case_791:
		switch data[p] {
		case 32:
			goto tr1116
		case 46:
			goto tr1117
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr892
		}
		goto st0
	st792:
		if p++; p == pe {
			goto _test_eof792
		}
	st_case_792:
		switch data[p] {
		case 97:
			goto st793
		case 101:
			goto st797
		case 117:
			goto st776
		}
		goto st0
	st793:
		if p++; p == pe {
			goto _test_eof793
		}
	st_case_793:
		if data[p] == 116 {
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
			goto st384
		case 44:
			goto st610
		case 117:
			goto st795
		}
		goto st0
	st795:
		if p++; p == pe {
			goto _test_eof795
		}
	st_case_795:
		if data[p] == 114 {
			goto st796
		}
		goto st0
	st796:
		if p++; p == pe {
			goto _test_eof796
		}
	st_case_796:
		if data[p] == 100 {
			goto st755
		}
		goto st0
	st797:
		if p++; p == pe {
			goto _test_eof797
		}
	st_case_797:
		if data[p] == 112 {
			goto st798
		}
		goto st0
	st798:
		if p++; p == pe {
			goto _test_eof798
		}
	st_case_798:
		switch data[p] {
		case 32:
			goto tr1128
		case 46:
			goto tr1129
		case 116:
			goto st799
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr901
		}
		goto st0
	st799:
		if p++; p == pe {
			goto _test_eof799
		}
	st_case_799:
		switch data[p] {
		case 32:
			goto tr1128
		case 46:
			goto tr1129
		case 101:
			goto st800
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr901
		}
		goto st0
	st800:
		if p++; p == pe {
			goto _test_eof800
		}
	st_case_800:
		if data[p] == 109 {
			goto st801
		}
		goto st0
	st801:
		if p++; p == pe {
			goto _test_eof801
		}
	st_case_801:
		if data[p] == 98 {
			goto st802
		}
		goto st0
	st802:
		if p++; p == pe {
			goto _test_eof802
		}
	st_case_802:
		if data[p] == 101 {
			goto st803
		}
		goto st0
	st803:
		if p++; p == pe {
			goto _test_eof803
		}
	st_case_803:
		if data[p] == 114 {
			goto st804
		}
		goto st0
	st804:
		if p++; p == pe {
			goto _test_eof804
		}
	st_case_804:
		switch data[p] {
		case 32:
			goto tr1128
		case 46:
			goto tr1129
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr901
		}
		goto st0
	st805:
		if p++; p == pe {
			goto _test_eof805
		}
	st_case_805:
		switch data[p] {
		case 104:
			goto st806
		case 117:
			goto st809
		}
		goto st0
	st806:
		if p++; p == pe {
			goto _test_eof806
		}
	st_case_806:
		if data[p] == 117 {
			goto st807
		}
		goto st0
	st807:
		if p++; p == pe {
			goto _test_eof807
		}
	st_case_807:
		switch data[p] {
		case 32:
			goto st384
		case 44:
			goto st610
		case 114:
			goto st808
		}
		goto st0
	st808:
		if p++; p == pe {
			goto _test_eof808
		}
	st_case_808:
		if data[p] == 115 {
			goto st796
		}
		goto st0
	st809:
		if p++; p == pe {
			goto _test_eof809
		}
	st_case_809:
		if data[p] == 101 {
			goto st810
		}
		goto st0
	st810:
		if p++; p == pe {
			goto _test_eof810
		}
	st_case_810:
		switch data[p] {
		case 32:
			goto st384
		case 44:
			goto st610
		case 115:
			goto st796
		}
		goto st0
	st811:
		if p++; p == pe {
			goto _test_eof811
		}
	st_case_811:
		if data[p] == 101 {
			goto st812
		}
		goto st0
	st812:
		if p++; p == pe {
			goto _test_eof812
		}
	st_case_812:
		if data[p] == 100 {
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
			goto st384
		case 44:
			goto st610
		case 110:
			goto st814
		}
		goto st0
	st814:
		if p++; p == pe {
			goto _test_eof814
		}
	st_case_814:
		if data[p] == 101 {
			goto st808
		}
		goto st0
	st815:
		if p++; p == pe {
			goto _test_eof815
		}
	st_case_815:
		if data[p] == 101 {
			goto st375
		}
		goto st0
	st816:
		if p++; p == pe {
			goto _test_eof816
		}
	st_case_816:
		if data[p] == 97 {
			goto st771
		}
		goto st0
	st817:
		if p++; p == pe {
			goto _test_eof817
		}
	st_case_817:
		if data[p] == 101 {
			goto st797
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
	_test_eof818: cs = 818; goto _test_eof
	_test_eof819: cs = 819; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof820: cs = 820; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof821: cs = 821; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof822: cs = 822; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof823: cs = 823; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof824: cs = 824; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof825: cs = 825; goto _test_eof
	_test_eof826: cs = 826; goto _test_eof
	_test_eof827: cs = 827; goto _test_eof
	_test_eof828: cs = 828; goto _test_eof
	_test_eof829: cs = 829; goto _test_eof
	_test_eof830: cs = 830; goto _test_eof
	_test_eof831: cs = 831; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof832: cs = 832; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof833: cs = 833; goto _test_eof
	_test_eof834: cs = 834; goto _test_eof
	_test_eof835: cs = 835; goto _test_eof
	_test_eof836: cs = 836; goto _test_eof
	_test_eof837: cs = 837; goto _test_eof
	_test_eof838: cs = 838; goto _test_eof
	_test_eof839: cs = 839; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof840: cs = 840; goto _test_eof
	_test_eof841: cs = 841; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof842: cs = 842; goto _test_eof
	_test_eof843: cs = 843; goto _test_eof
	_test_eof844: cs = 844; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof845: cs = 845; goto _test_eof
	_test_eof846: cs = 846; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof847: cs = 847; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof848: cs = 848; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof849: cs = 849; goto _test_eof
	_test_eof850: cs = 850; goto _test_eof
	_test_eof851: cs = 851; goto _test_eof
	_test_eof852: cs = 852; goto _test_eof
	_test_eof853: cs = 853; goto _test_eof
	_test_eof854: cs = 854; goto _test_eof
	_test_eof855: cs = 855; goto _test_eof
	_test_eof856: cs = 856; goto _test_eof
	_test_eof857: cs = 857; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof858: cs = 858; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof859: cs = 859; goto _test_eof
	_test_eof860: cs = 860; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof861: cs = 861; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof862: cs = 862; goto _test_eof
	_test_eof863: cs = 863; goto _test_eof
	_test_eof864: cs = 864; goto _test_eof
	_test_eof865: cs = 865; goto _test_eof
	_test_eof866: cs = 866; goto _test_eof
	_test_eof867: cs = 867; goto _test_eof
	_test_eof868: cs = 868; goto _test_eof
	_test_eof869: cs = 869; goto _test_eof
	_test_eof870: cs = 870; goto _test_eof
	_test_eof871: cs = 871; goto _test_eof
	_test_eof872: cs = 872; goto _test_eof
	_test_eof873: cs = 873; goto _test_eof
	_test_eof874: cs = 874; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof875: cs = 875; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof876: cs = 876; goto _test_eof
	_test_eof877: cs = 877; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof878: cs = 878; goto _test_eof
	_test_eof879: cs = 879; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
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
	_test_eof880: cs = 880; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof881: cs = 881; goto _test_eof
	_test_eof882: cs = 882; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof883: cs = 883; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof884: cs = 884; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
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
	_test_eof153: cs = 153; goto _test_eof
	_test_eof154: cs = 154; goto _test_eof
	_test_eof155: cs = 155; goto _test_eof
	_test_eof156: cs = 156; goto _test_eof
	_test_eof157: cs = 157; goto _test_eof
	_test_eof158: cs = 158; goto _test_eof
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
	_test_eof885: cs = 885; goto _test_eof
	_test_eof280: cs = 280; goto _test_eof
	_test_eof281: cs = 281; goto _test_eof
	_test_eof886: cs = 886; goto _test_eof
	_test_eof282: cs = 282; goto _test_eof
	_test_eof283: cs = 283; goto _test_eof
	_test_eof284: cs = 284; goto _test_eof
	_test_eof887: cs = 887; goto _test_eof
	_test_eof888: cs = 888; goto _test_eof
	_test_eof889: cs = 889; goto _test_eof
	_test_eof890: cs = 890; goto _test_eof
	_test_eof891: cs = 891; goto _test_eof
	_test_eof892: cs = 892; goto _test_eof
	_test_eof893: cs = 893; goto _test_eof
	_test_eof285: cs = 285; goto _test_eof
	_test_eof286: cs = 286; goto _test_eof
	_test_eof287: cs = 287; goto _test_eof
	_test_eof894: cs = 894; goto _test_eof
	_test_eof288: cs = 288; goto _test_eof
	_test_eof895: cs = 895; goto _test_eof
	_test_eof896: cs = 896; goto _test_eof
	_test_eof897: cs = 897; goto _test_eof
	_test_eof898: cs = 898; goto _test_eof
	_test_eof899: cs = 899; goto _test_eof
	_test_eof900: cs = 900; goto _test_eof
	_test_eof901: cs = 901; goto _test_eof
	_test_eof289: cs = 289; goto _test_eof
	_test_eof902: cs = 902; goto _test_eof
	_test_eof290: cs = 290; goto _test_eof
	_test_eof291: cs = 291; goto _test_eof
	_test_eof292: cs = 292; goto _test_eof
	_test_eof293: cs = 293; goto _test_eof
	_test_eof294: cs = 294; goto _test_eof
	_test_eof295: cs = 295; goto _test_eof
	_test_eof903: cs = 903; goto _test_eof
	_test_eof904: cs = 904; goto _test_eof
	_test_eof905: cs = 905; goto _test_eof
	_test_eof906: cs = 906; goto _test_eof
	_test_eof296: cs = 296; goto _test_eof
	_test_eof907: cs = 907; goto _test_eof
	_test_eof908: cs = 908; goto _test_eof
	_test_eof297: cs = 297; goto _test_eof
	_test_eof909: cs = 909; goto _test_eof
	_test_eof910: cs = 910; goto _test_eof
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
	_test_eof310: cs = 310; goto _test_eof
	_test_eof311: cs = 311; goto _test_eof
	_test_eof312: cs = 312; goto _test_eof
	_test_eof313: cs = 313; goto _test_eof
	_test_eof314: cs = 314; goto _test_eof
	_test_eof315: cs = 315; goto _test_eof
	_test_eof316: cs = 316; goto _test_eof
	_test_eof317: cs = 317; goto _test_eof
	_test_eof318: cs = 318; goto _test_eof
	_test_eof319: cs = 319; goto _test_eof
	_test_eof320: cs = 320; goto _test_eof
	_test_eof321: cs = 321; goto _test_eof
	_test_eof322: cs = 322; goto _test_eof
	_test_eof323: cs = 323; goto _test_eof
	_test_eof324: cs = 324; goto _test_eof
	_test_eof325: cs = 325; goto _test_eof
	_test_eof326: cs = 326; goto _test_eof
	_test_eof327: cs = 327; goto _test_eof
	_test_eof328: cs = 328; goto _test_eof
	_test_eof329: cs = 329; goto _test_eof
	_test_eof330: cs = 330; goto _test_eof
	_test_eof331: cs = 331; goto _test_eof
	_test_eof332: cs = 332; goto _test_eof
	_test_eof333: cs = 333; goto _test_eof
	_test_eof334: cs = 334; goto _test_eof
	_test_eof335: cs = 335; goto _test_eof
	_test_eof911: cs = 911; goto _test_eof
	_test_eof336: cs = 336; goto _test_eof
	_test_eof912: cs = 912; goto _test_eof
	_test_eof913: cs = 913; goto _test_eof
	_test_eof337: cs = 337; goto _test_eof
	_test_eof914: cs = 914; goto _test_eof
	_test_eof915: cs = 915; goto _test_eof
	_test_eof338: cs = 338; goto _test_eof
	_test_eof916: cs = 916; goto _test_eof
	_test_eof339: cs = 339; goto _test_eof
	_test_eof917: cs = 917; goto _test_eof
	_test_eof340: cs = 340; goto _test_eof
	_test_eof918: cs = 918; goto _test_eof
	_test_eof919: cs = 919; goto _test_eof
	_test_eof920: cs = 920; goto _test_eof
	_test_eof921: cs = 921; goto _test_eof
	_test_eof922: cs = 922; goto _test_eof
	_test_eof923: cs = 923; goto _test_eof
	_test_eof924: cs = 924; goto _test_eof
	_test_eof925: cs = 925; goto _test_eof
	_test_eof926: cs = 926; goto _test_eof
	_test_eof341: cs = 341; goto _test_eof
	_test_eof927: cs = 927; goto _test_eof
	_test_eof342: cs = 342; goto _test_eof
	_test_eof928: cs = 928; goto _test_eof
	_test_eof929: cs = 929; goto _test_eof
	_test_eof343: cs = 343; goto _test_eof
	_test_eof930: cs = 930; goto _test_eof
	_test_eof344: cs = 344; goto _test_eof
	_test_eof931: cs = 931; goto _test_eof
	_test_eof932: cs = 932; goto _test_eof
	_test_eof933: cs = 933; goto _test_eof
	_test_eof934: cs = 934; goto _test_eof
	_test_eof935: cs = 935; goto _test_eof
	_test_eof936: cs = 936; goto _test_eof
	_test_eof937: cs = 937; goto _test_eof
	_test_eof938: cs = 938; goto _test_eof
	_test_eof939: cs = 939; goto _test_eof
	_test_eof940: cs = 940; goto _test_eof
	_test_eof941: cs = 941; goto _test_eof
	_test_eof942: cs = 942; goto _test_eof
	_test_eof943: cs = 943; goto _test_eof
	_test_eof345: cs = 345; goto _test_eof
	_test_eof944: cs = 944; goto _test_eof
	_test_eof945: cs = 945; goto _test_eof
	_test_eof346: cs = 346; goto _test_eof
	_test_eof347: cs = 347; goto _test_eof
	_test_eof348: cs = 348; goto _test_eof
	_test_eof349: cs = 349; goto _test_eof
	_test_eof350: cs = 350; goto _test_eof
	_test_eof946: cs = 946; goto _test_eof
	_test_eof947: cs = 947; goto _test_eof
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
	_test_eof385: cs = 385; goto _test_eof
	_test_eof386: cs = 386; goto _test_eof
	_test_eof387: cs = 387; goto _test_eof
	_test_eof388: cs = 388; goto _test_eof
	_test_eof389: cs = 389; goto _test_eof
	_test_eof390: cs = 390; goto _test_eof
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
	_test_eof948: cs = 948; goto _test_eof
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
	_test_eof949: cs = 949; goto _test_eof
	_test_eof950: cs = 950; goto _test_eof
	_test_eof482: cs = 482; goto _test_eof
	_test_eof483: cs = 483; goto _test_eof
	_test_eof951: cs = 951; goto _test_eof
	_test_eof484: cs = 484; goto _test_eof
	_test_eof485: cs = 485; goto _test_eof
	_test_eof486: cs = 486; goto _test_eof
	_test_eof952: cs = 952; goto _test_eof
	_test_eof487: cs = 487; goto _test_eof
	_test_eof953: cs = 953; goto _test_eof
	_test_eof488: cs = 488; goto _test_eof
	_test_eof954: cs = 954; goto _test_eof
	_test_eof955: cs = 955; goto _test_eof
	_test_eof956: cs = 956; goto _test_eof
	_test_eof957: cs = 957; goto _test_eof
	_test_eof958: cs = 958; goto _test_eof
	_test_eof959: cs = 959; goto _test_eof
	_test_eof960: cs = 960; goto _test_eof
	_test_eof489: cs = 489; goto _test_eof
	_test_eof490: cs = 490; goto _test_eof
	_test_eof491: cs = 491; goto _test_eof
	_test_eof961: cs = 961; goto _test_eof
	_test_eof492: cs = 492; goto _test_eof
	_test_eof962: cs = 962; goto _test_eof
	_test_eof963: cs = 963; goto _test_eof
	_test_eof964: cs = 964; goto _test_eof
	_test_eof965: cs = 965; goto _test_eof
	_test_eof966: cs = 966; goto _test_eof
	_test_eof967: cs = 967; goto _test_eof
	_test_eof968: cs = 968; goto _test_eof
	_test_eof493: cs = 493; goto _test_eof
	_test_eof494: cs = 494; goto _test_eof
	_test_eof969: cs = 969; goto _test_eof
	_test_eof495: cs = 495; goto _test_eof
	_test_eof970: cs = 970; goto _test_eof
	_test_eof971: cs = 971; goto _test_eof
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
	_test_eof972: cs = 972; goto _test_eof
	_test_eof973: cs = 973; goto _test_eof
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
	_test_eof526: cs = 526; goto _test_eof
	_test_eof527: cs = 527; goto _test_eof
	_test_eof528: cs = 528; goto _test_eof
	_test_eof529: cs = 529; goto _test_eof
	_test_eof530: cs = 530; goto _test_eof
	_test_eof531: cs = 531; goto _test_eof
	_test_eof532: cs = 532; goto _test_eof
	_test_eof533: cs = 533; goto _test_eof
	_test_eof534: cs = 534; goto _test_eof
	_test_eof535: cs = 535; goto _test_eof
	_test_eof974: cs = 974; goto _test_eof
	_test_eof975: cs = 975; goto _test_eof
	_test_eof976: cs = 976; goto _test_eof
	_test_eof977: cs = 977; goto _test_eof
	_test_eof536: cs = 536; goto _test_eof
	_test_eof978: cs = 978; goto _test_eof
	_test_eof979: cs = 979; goto _test_eof
	_test_eof537: cs = 537; goto _test_eof
	_test_eof980: cs = 980; goto _test_eof
	_test_eof538: cs = 538; goto _test_eof
	_test_eof539: cs = 539; goto _test_eof
	_test_eof540: cs = 540; goto _test_eof
	_test_eof541: cs = 541; goto _test_eof
	_test_eof542: cs = 542; goto _test_eof
	_test_eof543: cs = 543; goto _test_eof
	_test_eof544: cs = 544; goto _test_eof
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

	_test_eof: {}
	if p == eof {
		switch cs {
		case 828, 836, 846, 877, 890, 898, 902, 957, 965, 971, 973:
//line ragel/datetime.rl:7
 st.Zoned = true 
		case 882:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

		case 884, 885, 903, 911, 946, 948, 949, 974:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

		case 883:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

		case 880, 881:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

		case 848, 858, 917, 927:
//line ragel/datetime.rl:35

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

		case 821, 825, 845, 954, 970:
//line ragel/datetime.rl:50

    st.Ad_bc = ADBC_BC;

		case 843, 914:
//line ragel/datetime.rl:54

    if st.Hour > 12 {
        err = errors.New("hour out of range")
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

		case 818, 878, 879:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

		case 847, 906, 916, 977:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

		case 840, 874, 875, 905, 908, 909, 912, 943, 944, 976, 979, 980:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

		case 860, 929:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

		case 859, 873, 928, 942:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

		case 871, 940:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

		case 861, 872, 930, 941:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

		case 849, 850, 851, 852, 853, 854, 855, 856, 857, 862, 863, 864, 865, 866, 867, 868, 869, 870, 918, 919, 920, 921, 922, 923, 924, 925, 926, 931, 932, 933, 934, 935, 936, 937, 938, 939:
//line ragel/datetime.rl:117

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

		case 907, 978:
//line ragel/datetime.rl:35

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

		case 829, 830, 837, 838, 891, 892, 899, 900, 958, 959, 966, 967:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
		case 822, 826, 827, 831, 833, 834, 835, 839, 886, 888, 889, 893, 895, 896, 897, 901, 951, 955, 956, 960, 962, 963, 964, 968:
//line ragel/datetime.rl:166

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
		case 823, 832, 887, 894, 952, 961:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
//line ragel/parse_datetime.go:26411
		}
	}

	_out: {}
	}

//line ragel/parse_datetime.go.rl:34

    if p != eof {  // input date not fully consumed
        err = errors.New("input data not fully consumed");
        return
    }

    if cs < datetime_parser_first_final {
        err = fmt.Errorf("time parse error: %s", data)
    }
    return
}

