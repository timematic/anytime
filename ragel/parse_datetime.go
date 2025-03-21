
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
const datetime_parser_first_final int = 767
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
	case 767:
		goto st_case_767
	case 768:
		goto st_case_768
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 769:
		goto st_case_769
	case 13:
		goto st_case_13
	case 770:
		goto st_case_770
	case 14:
		goto st_case_14
	case 771:
		goto st_case_771
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 772:
		goto st_case_772
	case 18:
		goto st_case_18
	case 773:
		goto st_case_773
	case 19:
		goto st_case_19
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
	case 20:
		goto st_case_20
	case 781:
		goto st_case_781
	case 782:
		goto st_case_782
	case 21:
		goto st_case_21
	case 783:
		goto st_case_783
	case 784:
		goto st_case_784
	case 785:
		goto st_case_785
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 786:
		goto st_case_786
	case 24:
		goto st_case_24
	case 787:
		goto st_case_787
	case 25:
		goto st_case_25
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
	case 26:
		goto st_case_26
	case 797:
		goto st_case_797
	case 27:
		goto st_case_27
	case 798:
		goto st_case_798
	case 799:
		goto st_case_799
	case 28:
		goto st_case_28
	case 800:
		goto st_case_800
	case 29:
		goto st_case_29
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
	case 30:
		goto st_case_30
	case 814:
		goto st_case_814
	case 815:
		goto st_case_815
	case 816:
		goto st_case_816
	case 31:
		goto st_case_31
	case 817:
		goto st_case_817
	case 818:
		goto st_case_818
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
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
	case 819:
		goto st_case_819
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
	case 820:
		goto st_case_820
	case 821:
		goto st_case_821
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
	case 822:
		goto st_case_822
	case 117:
		goto st_case_117
	case 823:
		goto st_case_823
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
	case 824:
		goto st_case_824
	case 272:
		goto st_case_272
	case 273:
		goto st_case_273
	case 825:
		goto st_case_825
	case 274:
		goto st_case_274
	case 275:
		goto st_case_275
	case 276:
		goto st_case_276
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
	case 277:
		goto st_case_277
	case 833:
		goto st_case_833
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
	case 834:
		goto st_case_834
	case 835:
		goto st_case_835
	case 836:
		goto st_case_836
	case 837:
		goto st_case_837
	case 284:
		goto st_case_284
	case 838:
		goto st_case_838
	case 839:
		goto st_case_839
	case 285:
		goto st_case_285
	case 840:
		goto st_case_840
	case 841:
		goto st_case_841
	case 286:
		goto st_case_286
	case 842:
		goto st_case_842
	case 843:
		goto st_case_843
	case 287:
		goto st_case_287
	case 844:
		goto st_case_844
	case 845:
		goto st_case_845
	case 288:
		goto st_case_288
	case 846:
		goto st_case_846
	case 289:
		goto st_case_289
	case 847:
		goto st_case_847
	case 290:
		goto st_case_290
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
	case 291:
		goto st_case_291
	case 857:
		goto st_case_857
	case 292:
		goto st_case_292
	case 858:
		goto st_case_858
	case 859:
		goto st_case_859
	case 293:
		goto st_case_293
	case 860:
		goto st_case_860
	case 294:
		goto st_case_294
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
	case 295:
		goto st_case_295
	case 874:
		goto st_case_874
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
	case 875:
		goto st_case_875
	case 876:
		goto st_case_876
	case 334:
		goto st_case_334
	case 335:
		goto st_case_335
	case 336:
		goto st_case_336
	case 337:
		goto st_case_337
	case 338:
		goto st_case_338
	case 877:
		goto st_case_877
	case 339:
		goto st_case_339
	case 340:
		goto st_case_340
	case 341:
		goto st_case_341
	case 342:
		goto st_case_342
	case 343:
		goto st_case_343
	case 344:
		goto st_case_344
	case 345:
		goto st_case_345
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
	case 385:
		goto st_case_385
	case 386:
		goto st_case_386
	case 878:
		goto st_case_878
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
	case 536:
		goto st_case_536
	case 537:
		goto st_case_537
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
	case 879:
		goto st_case_879
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
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 48:
			goto tr0
		case 51:
			goto tr3
		case 65:
			goto st260
		case 68:
			goto st354
		case 70:
			goto st362
		case 74:
			goto st707
		case 77:
			goto st719
		case 78:
			goto st726
		case 79:
			goto st734
		case 83:
			goto st741
		case 84:
			goto st754
		case 87:
			goto st760
		case 97:
			goto st260
		case 100:
			goto st354
		case 102:
			goto st764
		case 106:
			goto st707
		case 109:
			goto st765
		case 110:
			goto st726
		case 111:
			goto st734
		case 115:
			goto st766
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
//line ragel/parse_datetime.go:1885
		if data[p] == 48 {
			goto st3
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st110
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
//line ragel/parse_datetime.go:1939
		switch data[p] {
		case 48:
			goto tr25
		case 49:
			goto tr26
		case 65:
			goto st33
		case 68:
			goto st44
		case 70:
			goto st52
		case 74:
			goto st60
		case 77:
			goto st72
		case 78:
			goto st78
		case 79:
			goto st86
		case 83:
			goto st93
		case 97:
			goto st33
		case 100:
			goto st44
		case 102:
			goto st52
		case 106:
			goto st60
		case 109:
			goto st72
		case 110:
			goto st78
		case 111:
			goto st86
		case 115:
			goto st93
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
//line ragel/parse_datetime.go:1991
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
//line ragel/parse_datetime.go:2005
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
tr76:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st9
tr82:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st9
tr89:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st9
tr98:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st9
tr108:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st9
tr116:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st9
tr119:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st9
tr125:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st9
tr129:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st9
tr133:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st9
tr142:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st9
tr150:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line ragel/parse_datetime.go:2072
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
//line ragel/parse_datetime.go:2097
		if 49 <= data[p] && data[p] <= 57 {
			goto st767
		}
		goto st0
tr41:
//line ragel/datetime.rl:5
 pb = p 
	goto st767
	st767:
		if p++; p == pe {
			goto _test_eof767
		}
	st_case_767:
//line ragel/parse_datetime.go:2111
		switch data[p] {
		case 32:
			goto tr1083
		case 43:
			goto tr1084
		case 45:
			goto tr1085
		case 47:
			goto tr1086
		case 84:
			goto tr1087
		case 90:
			goto tr1088
		case 95:
			goto tr1089
		case 116:
			goto tr1089
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1086
			}
		case data[p] >= 65:
			goto tr1086
		}
		goto st0
tr1219:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st768
tr1083:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st768
tr1197:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st768
tr1205:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st768
tr1212:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st768
	st768:
		if p++; p == pe {
			goto _test_eof768
		}
	st_case_768:
//line ragel/parse_datetime.go:2180
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr50
		case 50:
			goto tr71
		case 65:
			goto tr51
		case 66:
			goto tr52
		case 90:
			goto tr1093
		case 95:
			goto tr50
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr70
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr50
				}
			case data[p] >= 67:
				goto tr50
			}
		default:
			goto tr72
		}
		goto st0
tr1097:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st11
tr1120:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line ragel/parse_datetime.go:2237
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
			goto st769
		}
		goto st0
	st769:
		if p++; p == pe {
			goto _test_eof769
		}
	st_case_769:
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 67 {
			goto st770
		}
		goto st0
	st770:
		if p++; p == pe {
			goto _test_eof770
		}
	st_case_770:
		goto st0
tr1220:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st14
tr1084:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st14
tr1105:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st14
tr1117:
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
tr1122:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st14
tr1137:
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
tr1130:
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
tr1150:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st14
tr1159:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st14
tr1167:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st14
tr1187:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st14
tr1198:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st14
tr1206:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st14
tr1213:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line ragel/parse_datetime.go:2425
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
	goto st771
tr57:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st771
	st771:
		if p++; p == pe {
			goto _test_eof771
		}
	st_case_771:
//line ragel/parse_datetime.go:2453
		switch data[p] {
		case 32:
			goto tr1094
		case 58:
			goto tr1096
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st775
		}
		goto st0
tr1094:
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
tr1099:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st15
tr1102:
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
//line ragel/parse_datetime.go:2518
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
tr1222:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1107:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1124:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1152:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1161:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1170:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1139:
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
	goto st16
tr1190:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1133:
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
	goto st16
tr1086:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1200:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1208:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
tr1215:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line ragel/parse_datetime.go:2693
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
tr1195:
//line ragel/datetime.rl:5
 pb = p 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line ragel/parse_datetime.go:2718
		switch data[p] {
		case 47:
			goto st772
		case 95:
			goto st772
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st772
			}
		case data[p] >= 65:
			goto st772
		}
		goto st0
tr1119:
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
	goto st772
	st772:
		if p++; p == pe {
			goto _test_eof772
		}
	st_case_772:
//line ragel/parse_datetime.go:2766
		switch data[p] {
		case 32:
			goto tr1097
		case 47:
			goto st772
		case 95:
			goto st772
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st772
			}
		case data[p] >= 65:
			goto st772
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
//line ragel/parse_datetime.go:2793
		switch data[p] {
		case 47:
			goto st17
		case 68:
			goto st773
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
	st773:
		if p++; p == pe {
			goto _test_eof773
		}
	st_case_773:
		switch data[p] {
		case 47:
			goto st772
		case 95:
			goto st772
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st772
			}
		case data[p] >= 65:
			goto st772
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
//line ragel/parse_datetime.go:2840
		switch data[p] {
		case 47:
			goto st17
		case 67:
			goto st774
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
	st774:
		if p++; p == pe {
			goto _test_eof774
		}
	st_case_774:
		switch data[p] {
		case 47:
			goto st772
		case 95:
			goto st772
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st772
			}
		case data[p] >= 65:
			goto st772
		}
		goto st0
tr49:
//line ragel/datetime.rl:5
 pb = p 
	goto st775
tr59:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st775
	st775:
		if p++; p == pe {
			goto _test_eof775
		}
	st_case_775:
//line ragel/parse_datetime.go:2893
		switch data[p] {
		case 32:
			goto tr1094
		case 58:
			goto tr1096
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st776
		}
		goto st0
	st776:
		if p++; p == pe {
			goto _test_eof776
		}
	st_case_776:
		if data[p] == 32 {
			goto tr1094
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st776
		}
		goto st0
tr1096:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st777
	st777:
		if p++; p == pe {
			goto _test_eof777
		}
	st_case_777:
//line ragel/parse_datetime.go:2932
		if data[p] == 32 {
			goto tr1099
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1101
			}
		case data[p] >= 48:
			goto tr1100
		}
		goto st0
tr1100:
//line ragel/datetime.rl:5
 pb = p 
	goto st778
	st778:
		if p++; p == pe {
			goto _test_eof778
		}
	st_case_778:
//line ragel/parse_datetime.go:2954
		if data[p] == 32 {
			goto tr1102
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st779
		}
		goto st0
tr1101:
//line ragel/datetime.rl:5
 pb = p 
	goto st779
	st779:
		if p++; p == pe {
			goto _test_eof779
		}
	st_case_779:
//line ragel/parse_datetime.go:2971
		if data[p] == 32 {
			goto tr1102
		}
		goto st0
tr48:
//line ragel/datetime.rl:5
 pb = p 
	goto st780
tr58:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st780
	st780:
		if p++; p == pe {
			goto _test_eof780
		}
	st_case_780:
//line ragel/parse_datetime.go:2991
		switch data[p] {
		case 32:
			goto tr1094
		case 58:
			goto tr1096
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st776
			}
		case data[p] >= 48:
			goto st775
		}
		goto st0
tr1221:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st20
tr1085:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st20
tr1106:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st20
tr1118:
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
tr1123:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st20
tr1138:
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
tr1132:
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
tr1151:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st20
tr1160:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st20
tr1169:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st20
tr1189:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st20
tr1199:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st20
tr1207:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st20
tr1214:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line ragel/parse_datetime.go:3157
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
tr70:
//line ragel/datetime.rl:5
 pb = p 
	goto st781
	st781:
		if p++; p == pe {
			goto _test_eof781
		}
	st_case_781:
//line ragel/parse_datetime.go:3179
		switch data[p] {
		case 32:
			goto tr1104
		case 43:
			goto tr1105
		case 45:
			goto tr1106
		case 47:
			goto tr1107
		case 58:
			goto tr1109
		case 65:
			goto tr1110
		case 80:
			goto tr1110
		case 90:
			goto tr1111
		case 95:
			goto tr1107
		case 97:
			goto tr1112
		case 112:
			goto tr1112
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st786
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1107
			}
		default:
			goto tr1107
		}
		goto st0
tr1104:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st782
tr1121:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st782
tr1175:
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

	goto st782
tr1149:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st782
tr1158:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st782
tr1166:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st782
tr1186:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st782
	st782:
		if p++; p == pe {
			goto _test_eof782
		}
	st_case_782:
//line ragel/parse_datetime.go:3289
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr50
		case 65:
			goto tr1113
		case 66:
			goto tr52
		case 80:
			goto tr1114
		case 90:
			goto tr1093
		case 95:
			goto tr50
		case 97:
			goto tr1115
		case 112:
			goto tr1115
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr50
			}
		case data[p] >= 67:
			goto tr50
		}
		goto st0
tr1113:
//line ragel/datetime.rl:5
 pb = p 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line ragel/parse_datetime.go:3332
		switch data[p] {
		case 47:
			goto st17
		case 68:
			goto st773
		case 77:
			goto st783
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
	st783:
		if p++; p == pe {
			goto _test_eof783
		}
	st_case_783:
		switch data[p] {
		case 32:
			goto tr1116
		case 43:
			goto tr1117
		case 45:
			goto tr1118
		case 47:
			goto tr1119
		case 95:
			goto tr1119
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1119
			}
		case data[p] >= 65:
			goto tr1119
		}
		goto st0
tr1116:
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

	goto st784
tr1136:
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

	goto st784
tr1129:
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

	goto st784
tr1243:
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

	goto st784
	st784:
		if p++; p == pe {
			goto _test_eof784
		}
	st_case_784:
//line ragel/parse_datetime.go:3477
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr50
		case 65:
			goto tr51
		case 66:
			goto tr52
		case 90:
			goto tr1093
		case 95:
			goto tr50
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr50
			}
		case data[p] >= 67:
			goto tr50
		}
		goto st0
tr1093:
//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1224:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1111:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1127:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1156:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1164:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1173:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1141:
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
	goto st785
tr1192:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1135:
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
	goto st785
tr1088:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1202:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1210:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
tr1217:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st785
	st785:
		if p++; p == pe {
			goto _test_eof785
		}
	st_case_785:
//line ragel/parse_datetime.go:3660
		switch data[p] {
		case 32:
			goto tr1120
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
tr1114:
//line ragel/datetime.rl:5
 pb = p 
	goto st22
tr1110:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st22
tr1126:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st22
tr1155:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st22
tr1163:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st22
tr1172:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st22
tr1177:
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
	goto st22
tr1191:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line ragel/parse_datetime.go:3768
		switch data[p] {
		case 47:
			goto st17
		case 77:
			goto st783
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
tr1115:
//line ragel/datetime.rl:5
 pb = p 
	goto st23
tr1112:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st23
tr1128:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st23
tr1157:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st23
tr1165:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st23
tr1174:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st23
tr1178:
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
	goto st23
tr1193:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line ragel/parse_datetime.go:3876
		switch data[p] {
		case 47:
			goto st17
		case 95:
			goto st17
		case 109:
			goto st783
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
	st786:
		if p++; p == pe {
			goto _test_eof786
		}
	st_case_786:
		switch data[p] {
		case 32:
			goto tr1121
		case 43:
			goto tr1122
		case 45:
			goto tr1123
		case 47:
			goto tr1124
		case 58:
			goto tr1125
		case 65:
			goto tr1126
		case 80:
			goto tr1126
		case 90:
			goto tr1127
		case 95:
			goto tr1124
		case 97:
			goto tr1128
		case 112:
			goto tr1128
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st24
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1124
			}
		default:
			goto tr1124
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr1129
		case 43:
			goto tr1130
		case 45:
			goto tr1132
		case 47:
			goto tr1133
		case 90:
			goto tr1135
		case 95:
			goto tr1133
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1131
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1133
				}
			case data[p] >= 65:
				goto tr1133
			}
		default:
			goto st26
		}
		goto st0
tr1131:
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

	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line ragel/parse_datetime.go:4004
		if 48 <= data[p] && data[p] <= 57 {
			goto tr62
		}
		goto st0
tr62:
//line ragel/datetime.rl:5
 pb = p 
	goto st788
	st788:
		if p++; p == pe {
			goto _test_eof788
		}
	st_case_788:
//line ragel/parse_datetime.go:4018
		switch data[p] {
		case 32:
			goto tr1136
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st789
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st789:
		if p++; p == pe {
			goto _test_eof789
		}
	st_case_789:
		switch data[p] {
		case 32:
			goto tr1136
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st790
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st790:
		if p++; p == pe {
			goto _test_eof790
		}
	st_case_790:
		switch data[p] {
		case 32:
			goto tr1136
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st791
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st791:
		if p++; p == pe {
			goto _test_eof791
		}
	st_case_791:
		switch data[p] {
		case 32:
			goto tr1136
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st792
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st792:
		if p++; p == pe {
			goto _test_eof792
		}
	st_case_792:
		switch data[p] {
		case 32:
			goto tr1136
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st793
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st793:
		if p++; p == pe {
			goto _test_eof793
		}
	st_case_793:
		switch data[p] {
		case 32:
			goto tr1136
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st794
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st794:
		if p++; p == pe {
			goto _test_eof794
		}
	st_case_794:
		switch data[p] {
		case 32:
			goto tr1136
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st795
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st795:
		if p++; p == pe {
			goto _test_eof795
		}
	st_case_795:
		switch data[p] {
		case 32:
			goto tr1136
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st796
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st796:
		if p++; p == pe {
			goto _test_eof796
		}
	st_case_796:
		switch data[p] {
		case 32:
			goto tr1136
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		case data[p] >= 65:
			goto tr1139
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr1129
		case 43:
			goto tr1130
		case 45:
			goto tr1132
		case 47:
			goto tr1133
		case 90:
			goto tr1135
		case 95:
			goto tr1133
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1131
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1133
			}
		default:
			goto tr1133
		}
		goto st0
tr1109:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st27
tr1125:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line ragel/parse_datetime.go:4356
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr65
			}
		case data[p] >= 48:
			goto tr64
		}
		goto st0
tr64:
//line ragel/datetime.rl:5
 pb = p 
	goto st798
	st798:
		if p++; p == pe {
			goto _test_eof798
		}
	st_case_798:
//line ragel/parse_datetime.go:4375
		switch data[p] {
		case 32:
			goto tr1149
		case 43:
			goto tr1150
		case 45:
			goto tr1151
		case 47:
			goto tr1152
		case 58:
			goto tr1154
		case 65:
			goto tr1155
		case 80:
			goto tr1155
		case 90:
			goto tr1156
		case 95:
			goto tr1152
		case 97:
			goto tr1157
		case 112:
			goto tr1157
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st799
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1152
			}
		default:
			goto tr1152
		}
		goto st0
	st799:
		if p++; p == pe {
			goto _test_eof799
		}
	st_case_799:
		switch data[p] {
		case 32:
			goto tr1158
		case 43:
			goto tr1159
		case 45:
			goto tr1160
		case 47:
			goto tr1161
		case 58:
			goto tr1162
		case 65:
			goto tr1163
		case 80:
			goto tr1163
		case 90:
			goto tr1164
		case 95:
			goto tr1161
		case 97:
			goto tr1165
		case 112:
			goto tr1165
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1161
			}
		case data[p] >= 66:
			goto tr1161
		}
		goto st0
tr1154:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st28
tr1162:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line ragel/parse_datetime.go:4468
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr67
			}
		case data[p] >= 48:
			goto tr66
		}
		goto st0
tr66:
//line ragel/datetime.rl:5
 pb = p 
	goto st800
	st800:
		if p++; p == pe {
			goto _test_eof800
		}
	st_case_800:
//line ragel/parse_datetime.go:4487
		switch data[p] {
		case 32:
			goto tr1166
		case 43:
			goto tr1167
		case 45:
			goto tr1169
		case 47:
			goto tr1170
		case 65:
			goto tr1172
		case 80:
			goto tr1172
		case 90:
			goto tr1173
		case 95:
			goto tr1170
		case 97:
			goto tr1174
		case 112:
			goto tr1174
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1168
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1170
				}
			case data[p] >= 66:
				goto tr1170
			}
		default:
			goto st810
		}
		goto st0
tr1168:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st29
tr1188:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line ragel/parse_datetime.go:4545
		if 48 <= data[p] && data[p] <= 57 {
			goto tr68
		}
		goto st0
tr68:
//line ragel/datetime.rl:5
 pb = p 
	goto st801
	st801:
		if p++; p == pe {
			goto _test_eof801
		}
	st_case_801:
//line ragel/parse_datetime.go:4559
		switch data[p] {
		case 32:
			goto tr1175
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 65:
			goto tr1177
		case 80:
			goto tr1177
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		case 97:
			goto tr1178
		case 112:
			goto tr1178
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st802
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st802:
		if p++; p == pe {
			goto _test_eof802
		}
	st_case_802:
		switch data[p] {
		case 32:
			goto tr1175
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 65:
			goto tr1177
		case 80:
			goto tr1177
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		case 97:
			goto tr1178
		case 112:
			goto tr1178
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st803
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st803:
		if p++; p == pe {
			goto _test_eof803
		}
	st_case_803:
		switch data[p] {
		case 32:
			goto tr1175
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 65:
			goto tr1177
		case 80:
			goto tr1177
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		case 97:
			goto tr1178
		case 112:
			goto tr1178
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st804
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st804:
		if p++; p == pe {
			goto _test_eof804
		}
	st_case_804:
		switch data[p] {
		case 32:
			goto tr1175
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 65:
			goto tr1177
		case 80:
			goto tr1177
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		case 97:
			goto tr1178
		case 112:
			goto tr1178
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st805
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st805:
		if p++; p == pe {
			goto _test_eof805
		}
	st_case_805:
		switch data[p] {
		case 32:
			goto tr1175
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 65:
			goto tr1177
		case 80:
			goto tr1177
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		case 97:
			goto tr1178
		case 112:
			goto tr1178
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st806
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st806:
		if p++; p == pe {
			goto _test_eof806
		}
	st_case_806:
		switch data[p] {
		case 32:
			goto tr1175
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 65:
			goto tr1177
		case 80:
			goto tr1177
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		case 97:
			goto tr1178
		case 112:
			goto tr1178
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st807
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st807:
		if p++; p == pe {
			goto _test_eof807
		}
	st_case_807:
		switch data[p] {
		case 32:
			goto tr1175
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 65:
			goto tr1177
		case 80:
			goto tr1177
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		case 97:
			goto tr1178
		case 112:
			goto tr1178
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st808
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st808:
		if p++; p == pe {
			goto _test_eof808
		}
	st_case_808:
		switch data[p] {
		case 32:
			goto tr1175
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 65:
			goto tr1177
		case 80:
			goto tr1177
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		case 97:
			goto tr1178
		case 112:
			goto tr1178
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st809
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		default:
			goto tr1139
		}
		goto st0
	st809:
		if p++; p == pe {
			goto _test_eof809
		}
	st_case_809:
		switch data[p] {
		case 32:
			goto tr1175
		case 43:
			goto tr1137
		case 45:
			goto tr1138
		case 47:
			goto tr1139
		case 65:
			goto tr1177
		case 80:
			goto tr1177
		case 90:
			goto tr1141
		case 95:
			goto tr1139
		case 97:
			goto tr1178
		case 112:
			goto tr1178
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1139
			}
		case data[p] >= 66:
			goto tr1139
		}
		goto st0
	st810:
		if p++; p == pe {
			goto _test_eof810
		}
	st_case_810:
		switch data[p] {
		case 32:
			goto tr1186
		case 43:
			goto tr1187
		case 45:
			goto tr1189
		case 47:
			goto tr1190
		case 65:
			goto tr1191
		case 80:
			goto tr1191
		case 90:
			goto tr1192
		case 95:
			goto tr1190
		case 97:
			goto tr1193
		case 112:
			goto tr1193
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1188
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1190
			}
		default:
			goto tr1190
		}
		goto st0
tr67:
//line ragel/datetime.rl:5
 pb = p 
	goto st811
	st811:
		if p++; p == pe {
			goto _test_eof811
		}
	st_case_811:
//line ragel/parse_datetime.go:4960
		switch data[p] {
		case 32:
			goto tr1166
		case 43:
			goto tr1167
		case 45:
			goto tr1169
		case 47:
			goto tr1170
		case 65:
			goto tr1172
		case 80:
			goto tr1172
		case 90:
			goto tr1173
		case 95:
			goto tr1170
		case 97:
			goto tr1174
		case 112:
			goto tr1174
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1168
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1170
			}
		default:
			goto tr1170
		}
		goto st0
tr65:
//line ragel/datetime.rl:5
 pb = p 
	goto st812
	st812:
		if p++; p == pe {
			goto _test_eof812
		}
	st_case_812:
//line ragel/parse_datetime.go:5005
		switch data[p] {
		case 32:
			goto tr1149
		case 43:
			goto tr1150
		case 45:
			goto tr1151
		case 47:
			goto tr1152
		case 58:
			goto tr1154
		case 65:
			goto tr1155
		case 80:
			goto tr1155
		case 90:
			goto tr1156
		case 95:
			goto tr1152
		case 97:
			goto tr1157
		case 112:
			goto tr1157
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1152
			}
		case data[p] >= 66:
			goto tr1152
		}
		goto st0
tr71:
//line ragel/datetime.rl:5
 pb = p 
	goto st813
	st813:
		if p++; p == pe {
			goto _test_eof813
		}
	st_case_813:
//line ragel/parse_datetime.go:5048
		switch data[p] {
		case 32:
			goto tr1104
		case 43:
			goto tr1105
		case 45:
			goto tr1106
		case 47:
			goto tr1107
		case 58:
			goto tr1109
		case 65:
			goto tr1110
		case 80:
			goto tr1110
		case 90:
			goto tr1111
		case 95:
			goto tr1107
		case 97:
			goto tr1112
		case 112:
			goto tr1112
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st786
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1107
				}
			case data[p] >= 66:
				goto tr1107
			}
		default:
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if 48 <= data[p] && data[p] <= 57 {
			goto st24
		}
		goto st0
tr72:
//line ragel/datetime.rl:5
 pb = p 
	goto st814
	st814:
		if p++; p == pe {
			goto _test_eof814
		}
	st_case_814:
//line ragel/parse_datetime.go:5109
		switch data[p] {
		case 32:
			goto tr1104
		case 43:
			goto tr1105
		case 45:
			goto tr1106
		case 47:
			goto tr1107
		case 58:
			goto tr1109
		case 65:
			goto tr1110
		case 80:
			goto tr1110
		case 90:
			goto tr1111
		case 95:
			goto tr1107
		case 97:
			goto tr1112
		case 112:
			goto tr1112
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st30
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1107
			}
		default:
			goto tr1107
		}
		goto st0
tr1223:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st815
tr1087:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st815
tr1201:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st815
tr1209:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st815
tr1216:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st815
	st815:
		if p++; p == pe {
			goto _test_eof815
		}
	st_case_815:
//line ragel/parse_datetime.go:5198
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr1195
		case 50:
			goto tr71
		case 90:
			goto tr1196
		case 95:
			goto tr1195
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr70
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1195
				}
			case data[p] >= 65:
				goto tr1195
			}
		default:
			goto tr72
		}
		goto st0
tr1196:
//line ragel/datetime.rl:5
 pb = p 
	goto st816
	st816:
		if p++; p == pe {
			goto _test_eof816
		}
	st_case_816:
//line ragel/parse_datetime.go:5242
		switch data[p] {
		case 32:
			goto tr1120
		case 47:
			goto st772
		case 95:
			goto st772
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st772
			}
		case data[p] >= 65:
			goto st772
		}
		goto st0
tr1225:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1089:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1203:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1211:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1218:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line ragel/parse_datetime.go:5311
		switch data[p] {
		case 47:
			goto st17
		case 50:
			goto tr71
		case 95:
			goto st17
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr70
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st17
				}
			case data[p] >= 65:
				goto st17
			}
		default:
			goto tr72
		}
		goto st0
tr39:
//line ragel/datetime.rl:5
 pb = p 
	goto st817
	st817:
		if p++; p == pe {
			goto _test_eof817
		}
	st_case_817:
//line ragel/parse_datetime.go:5347
		switch data[p] {
		case 32:
			goto tr1083
		case 43:
			goto tr1084
		case 45:
			goto tr1085
		case 47:
			goto tr1086
		case 84:
			goto tr1087
		case 90:
			goto tr1088
		case 95:
			goto tr1089
		case 116:
			goto tr1089
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st767
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1086
			}
		default:
			goto tr1086
		}
		goto st0
tr40:
//line ragel/datetime.rl:5
 pb = p 
	goto st818
	st818:
		if p++; p == pe {
			goto _test_eof818
		}
	st_case_818:
//line ragel/parse_datetime.go:5388
		switch data[p] {
		case 32:
			goto tr1083
		case 43:
			goto tr1084
		case 45:
			goto tr1085
		case 47:
			goto tr1086
		case 84:
			goto tr1087
		case 90:
			goto tr1088
		case 95:
			goto tr1089
		case 116:
			goto tr1089
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 49 {
				goto st767
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1086
			}
		default:
			goto tr1086
		}
		goto st0
tr26:
//line ragel/datetime.rl:5
 pb = p 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line ragel/parse_datetime.go:5429
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
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 112:
			goto st34
		case 117:
			goto st39
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 114 {
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 32:
			goto tr76
		case 46:
			goto tr77
		case 105:
			goto st37
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr76
		}
		goto st0
tr77:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st36
tr83:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st36
tr90:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st36
tr99:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st36
tr109:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st36
tr117:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st36
tr120:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st36
tr126:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st36
tr130:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st36
tr134:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st36
tr143:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st36
tr151:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line ragel/parse_datetime.go:5533
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
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 108 {
			goto st38
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 32:
			goto tr76
		case 46:
			goto tr77
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr76
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 103 {
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 32:
			goto tr82
		case 46:
			goto tr83
		case 117:
			goto st41
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr82
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 115 {
			goto st42
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 116 {
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
			goto tr82
		case 46:
			goto tr83
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr82
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		if data[p] == 101 {
			goto st45
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if data[p] == 99 {
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
			goto tr89
		case 46:
			goto tr90
		case 101:
			goto st47
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr89
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 109 {
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 98 {
			goto st49
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 101 {
			goto st50
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		if data[p] == 114 {
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
			goto tr89
		case 46:
			goto tr90
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr89
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
		if data[p] == 98 {
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
			goto tr98
		case 46:
			goto tr99
		case 114:
			goto st55
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr98
		}
		goto st0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 117 {
			goto st56
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 97 {
			goto st57
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 114 {
			goto st58
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 121 {
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
			goto tr98
		case 46:
			goto tr99
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr98
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 97:
			goto st61
		case 117:
			goto st67
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 110 {
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
			goto tr108
		case 46:
			goto tr109
		case 117:
			goto st63
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr108
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 97 {
			goto st64
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 114 {
			goto st65
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 121 {
			goto st66
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 32:
			goto tr108
		case 46:
			goto tr109
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr108
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 108:
			goto st68
		case 110:
			goto st70
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 32:
			goto tr116
		case 46:
			goto tr117
		case 121:
			goto st69
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr116
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 32:
			goto tr116
		case 46:
			goto tr117
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr116
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
		case 101:
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
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 97 {
			goto st73
		}
		goto st0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 114:
			goto st74
		case 121:
			goto st77
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 32:
			goto tr125
		case 46:
			goto tr126
		case 99:
			goto st75
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr125
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if data[p] == 104 {
			goto st76
		}
		goto st0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 32:
			goto tr125
		case 46:
			goto tr126
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr125
		}
		goto st0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 32:
			goto tr129
		case 46:
			goto tr130
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr129
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		if data[p] == 111 {
			goto st79
		}
		goto st0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		if data[p] == 118 {
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
			goto tr133
		case 46:
			goto tr134
		case 101:
			goto st81
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr133
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		if data[p] == 109 {
			goto st82
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		if data[p] == 98 {
			goto st83
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		if data[p] == 101 {
			goto st84
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 114 {
			goto st85
		}
		goto st0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 32:
			goto tr133
		case 46:
			goto tr134
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr133
		}
		goto st0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		if data[p] == 99 {
			goto st87
		}
		goto st0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		if data[p] == 116 {
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
			goto tr142
		case 46:
			goto tr143
		case 111:
			goto st89
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr142
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		if data[p] == 98 {
			goto st90
		}
		goto st0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		if data[p] == 101 {
			goto st91
		}
		goto st0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		if data[p] == 114 {
			goto st92
		}
		goto st0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 32:
			goto tr142
		case 46:
			goto tr143
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr142
		}
		goto st0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		if data[p] == 101 {
			goto st94
		}
		goto st0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		if data[p] == 112 {
			goto st95
		}
		goto st0
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 32:
			goto tr150
		case 46:
			goto tr151
		case 116:
			goto st96
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr150
		}
		goto st0
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 32:
			goto tr150
		case 46:
			goto tr151
		case 101:
			goto st97
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr150
		}
		goto st0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		if data[p] == 109 {
			goto st98
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		if data[p] == 98 {
			goto st99
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		if data[p] == 101 {
			goto st100
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		if data[p] == 114 {
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
			goto tr150
		case 46:
			goto tr151
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr150
		}
		goto st0
tr23:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line ragel/parse_datetime.go:6320
		switch data[p] {
		case 48:
			goto tr158
		case 49:
			goto tr159
		case 65:
			goto st33
		case 68:
			goto st44
		case 70:
			goto st52
		case 74:
			goto st60
		case 77:
			goto st72
		case 78:
			goto st78
		case 79:
			goto st86
		case 83:
			goto st93
		case 97:
			goto st33
		case 100:
			goto st44
		case 102:
			goto st52
		case 106:
			goto st60
		case 109:
			goto st72
		case 110:
			goto st78
		case 111:
			goto st86
		case 115:
			goto st93
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr160
		}
		goto st0
tr158:
//line ragel/datetime.rl:5
 pb = p 
	goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
//line ragel/parse_datetime.go:6372
		if data[p] == 48 {
			goto st104
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st105
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		if 48 <= data[p] && data[p] <= 57 {
			goto st819
		}
		goto st0
	st819:
		if p++; p == pe {
			goto _test_eof819
		}
	st_case_819:
		switch data[p] {
		case 32:
			goto tr1197
		case 43:
			goto tr1198
		case 45:
			goto tr1199
		case 47:
			goto tr1200
		case 84:
			goto tr1201
		case 90:
			goto tr1202
		case 95:
			goto tr1203
		case 116:
			goto tr1203
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1200
			}
		case data[p] >= 65:
			goto tr1200
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		if data[p] == 32 {
			goto tr37
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st819
			}
		case data[p] >= 45:
			goto tr37
		}
		goto st0
tr159:
//line ragel/datetime.rl:5
 pb = p 
	goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line ragel/parse_datetime.go:6447
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
				goto st104
			}
		default:
			goto st105
		}
		goto st0
tr160:
//line ragel/datetime.rl:5
 pb = p 
	goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line ragel/parse_datetime.go:6473
		if data[p] == 32 {
			goto tr37
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st104
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
	goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line ragel/parse_datetime.go:6499
		if 48 <= data[p] && data[p] <= 57 {
			goto st109
		}
		goto st0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		if 48 <= data[p] && data[p] <= 57 {
			goto st820
		}
		goto st0
	st820:
		if p++; p == pe {
			goto _test_eof820
		}
	st_case_820:
		switch data[p] {
		case 32:
			goto tr1197
		case 43:
			goto tr1198
		case 45:
			goto tr1199
		case 47:
			goto tr1200
		case 84:
			goto tr1201
		case 90:
			goto tr1202
		case 95:
			goto tr1203
		case 116:
			goto tr1203
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st821
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1200
			}
		default:
			goto tr1200
		}
		goto st0
	st821:
		if p++; p == pe {
			goto _test_eof821
		}
	st_case_821:
		switch data[p] {
		case 32:
			goto tr1205
		case 43:
			goto tr1206
		case 45:
			goto tr1207
		case 47:
			goto tr1208
		case 84:
			goto tr1209
		case 90:
			goto tr1210
		case 95:
			goto tr1211
		case 116:
			goto tr1211
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1208
			}
		case data[p] >= 65:
			goto tr1208
		}
		goto st0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		if data[p] == 32 {
			goto tr166
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] >= 45:
			goto tr167
		}
		goto st0
tr166:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line ragel/parse_datetime.go:6614
		switch data[p] {
		case 65:
			goto st112
		case 68:
			goto st129
		case 70:
			goto st137
		case 74:
			goto st145
		case 77:
			goto st157
		case 78:
			goto st163
		case 79:
			goto st171
		case 83:
			goto st178
		case 97:
			goto st112
		case 100:
			goto st129
		case 102:
			goto st137
		case 106:
			goto st145
		case 109:
			goto st157
		case 110:
			goto st163
		case 111:
			goto st171
		case 115:
			goto st178
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 112:
			goto st113
		case 117:
			goto st124
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		if data[p] == 114 {
			goto st114
		}
		goto st0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 32:
			goto tr179
		case 46:
			goto tr181
		case 105:
			goto st122
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr180
		}
		goto st0
tr179:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st115
tr193:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st115
tr201:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st115
tr211:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st115
tr222:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st115
tr231:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st115
tr235:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st115
tr242:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st115
tr247:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st115
tr252:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st115
tr262:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st115
tr271:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
//line ragel/parse_datetime.go:6741
		if 48 <= data[p] && data[p] <= 57 {
			goto tr183
		}
		goto st0
tr183:
//line ragel/datetime.rl:5
 pb = p 
	goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line ragel/parse_datetime.go:6755
		if 48 <= data[p] && data[p] <= 57 {
			goto st822
		}
		goto st0
	st822:
		if p++; p == pe {
			goto _test_eof822
		}
	st_case_822:
		switch data[p] {
		case 32:
			goto tr1212
		case 43:
			goto tr1213
		case 45:
			goto tr1214
		case 47:
			goto tr1215
		case 84:
			goto tr1216
		case 90:
			goto tr1217
		case 95:
			goto tr1218
		case 116:
			goto tr1218
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st117
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1215
			}
		default:
			goto tr1215
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr1219
		case 43:
			goto tr1220
		case 45:
			goto tr1221
		case 47:
			goto tr1222
		case 84:
			goto tr1223
		case 90:
			goto tr1224
		case 95:
			goto tr1225
		case 116:
			goto tr1225
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1222
			}
		case data[p] >= 65:
			goto tr1222
		}
		goto st0
tr180:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st118
tr194:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st118
tr202:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st118
tr212:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st118
tr223:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st118
tr232:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st118
tr236:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st118
tr243:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st118
tr248:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st118
tr253:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st118
tr263:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st118
tr272:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st118
tr374:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line ragel/parse_datetime.go:6901
		if 48 <= data[p] && data[p] <= 57 {
			goto tr186
		}
		goto st0
tr186:
//line ragel/datetime.rl:5
 pb = p 
	goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line ragel/parse_datetime.go:6915
		if 48 <= data[p] && data[p] <= 57 {
			goto st120
		}
		goto st0
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		if 48 <= data[p] && data[p] <= 57 {
			goto st117
		}
		goto st0
tr181:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st121
tr195:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st121
tr203:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st121
tr213:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st121
tr224:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st121
tr233:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st121
tr237:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st121
tr244:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st121
tr249:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st121
tr254:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st121
tr264:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st121
tr273:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
//line ragel/parse_datetime.go:6982
		if data[p] == 32 {
			goto st115
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr186
			}
		case data[p] >= 45:
			goto st118
		}
		goto st0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		if data[p] == 108 {
			goto st123
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 32:
			goto tr179
		case 46:
			goto tr181
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr180
		}
		goto st0
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		if data[p] == 103 {
			goto st125
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 32:
			goto tr193
		case 46:
			goto tr195
		case 117:
			goto st126
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr194
		}
		goto st0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		if data[p] == 115 {
			goto st127
		}
		goto st0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		if data[p] == 116 {
			goto st128
		}
		goto st0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 32:
			goto tr193
		case 46:
			goto tr195
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr194
		}
		goto st0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		if data[p] == 101 {
			goto st130
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		if data[p] == 99 {
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
			goto tr201
		case 46:
			goto tr203
		case 101:
			goto st132
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr202
		}
		goto st0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		if data[p] == 109 {
			goto st133
		}
		goto st0
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		if data[p] == 98 {
			goto st134
		}
		goto st0
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		if data[p] == 101 {
			goto st135
		}
		goto st0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		if data[p] == 114 {
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
			goto tr201
		case 46:
			goto tr203
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr202
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
		if data[p] == 98 {
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
			goto tr211
		case 46:
			goto tr213
		case 114:
			goto st140
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr212
		}
		goto st0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		if data[p] == 117 {
			goto st141
		}
		goto st0
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		if data[p] == 97 {
			goto st142
		}
		goto st0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		if data[p] == 114 {
			goto st143
		}
		goto st0
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		if data[p] == 121 {
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
			goto tr211
		case 46:
			goto tr213
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr212
		}
		goto st0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 97:
			goto st146
		case 117:
			goto st152
		}
		goto st0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		if data[p] == 110 {
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
		case 117:
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
		if data[p] == 97 {
			goto st149
		}
		goto st0
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		if data[p] == 114 {
			goto st150
		}
		goto st0
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		if data[p] == 121 {
			goto st151
		}
		goto st0
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
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
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 108:
			goto st153
		case 110:
			goto st155
		}
		goto st0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 32:
			goto tr231
		case 46:
			goto tr233
		case 121:
			goto st154
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr232
		}
		goto st0
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 32:
			goto tr231
		case 46:
			goto tr233
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr232
		}
		goto st0
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 32:
			goto tr235
		case 46:
			goto tr237
		case 101:
			goto st156
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr236
		}
		goto st0
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 32:
			goto tr235
		case 46:
			goto tr237
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr236
		}
		goto st0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		if data[p] == 97 {
			goto st158
		}
		goto st0
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 114:
			goto st159
		case 121:
			goto st162
		}
		goto st0
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 32:
			goto tr242
		case 46:
			goto tr244
		case 99:
			goto st160
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr243
		}
		goto st0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		if data[p] == 104 {
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
			goto tr242
		case 46:
			goto tr244
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
			goto tr247
		case 46:
			goto tr249
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr248
		}
		goto st0
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		if data[p] == 111 {
			goto st164
		}
		goto st0
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		if data[p] == 118 {
			goto st165
		}
		goto st0
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch data[p] {
		case 32:
			goto tr252
		case 46:
			goto tr254
		case 101:
			goto st166
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr253
		}
		goto st0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		if data[p] == 109 {
			goto st167
		}
		goto st0
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		if data[p] == 98 {
			goto st168
		}
		goto st0
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		if data[p] == 101 {
			goto st169
		}
		goto st0
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		if data[p] == 114 {
			goto st170
		}
		goto st0
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 32:
			goto tr252
		case 46:
			goto tr254
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr253
		}
		goto st0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		if data[p] == 99 {
			goto st172
		}
		goto st0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		if data[p] == 116 {
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
			goto tr262
		case 46:
			goto tr264
		case 111:
			goto st174
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr263
		}
		goto st0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		if data[p] == 98 {
			goto st175
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
		if data[p] == 114 {
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
			goto tr262
		case 46:
			goto tr264
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr263
		}
		goto st0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		if data[p] == 101 {
			goto st179
		}
		goto st0
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		if data[p] == 112 {
			goto st180
		}
		goto st0
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		switch data[p] {
		case 32:
			goto tr271
		case 46:
			goto tr273
		case 116:
			goto st181
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr272
		}
		goto st0
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 32:
			goto tr271
		case 46:
			goto tr273
		case 101:
			goto st182
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr272
		}
		goto st0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		if data[p] == 109 {
			goto st183
		}
		goto st0
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		if data[p] == 98 {
			goto st184
		}
		goto st0
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		if data[p] == 101 {
			goto st185
		}
		goto st0
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		if data[p] == 114 {
			goto st186
		}
		goto st0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		switch data[p] {
		case 32:
			goto tr271
		case 46:
			goto tr273
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr272
		}
		goto st0
tr167:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st187
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
//line ragel/parse_datetime.go:7765
		switch data[p] {
		case 65:
			goto st188
		case 68:
			goto st199
		case 70:
			goto st207
		case 74:
			goto st215
		case 77:
			goto st227
		case 78:
			goto st233
		case 79:
			goto st241
		case 83:
			goto st248
		case 97:
			goto st188
		case 100:
			goto st199
		case 102:
			goto st207
		case 106:
			goto st215
		case 109:
			goto st227
		case 110:
			goto st233
		case 111:
			goto st241
		case 115:
			goto st248
		}
		goto st0
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch data[p] {
		case 112:
			goto st189
		case 117:
			goto st194
		}
		goto st0
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		if data[p] == 114 {
			goto st190
		}
		goto st0
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		switch data[p] {
		case 32:
			goto tr180
		case 46:
			goto tr291
		case 105:
			goto st192
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr180
		}
		goto st0
tr291:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st191
tr295:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st191
tr301:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st191
tr309:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st191
tr318:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st191
tr325:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st191
tr327:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st191
tr332:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st191
tr335:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st191
tr338:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st191
tr346:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st191
tr353:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st191
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
//line ragel/parse_datetime.go:7892
		if data[p] == 32 {
			goto st118
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr186
			}
		case data[p] >= 45:
			goto st118
		}
		goto st0
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		if data[p] == 108 {
			goto st193
		}
		goto st0
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		switch data[p] {
		case 32:
			goto tr180
		case 46:
			goto tr291
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr180
		}
		goto st0
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		if data[p] == 103 {
			goto st195
		}
		goto st0
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch data[p] {
		case 32:
			goto tr194
		case 46:
			goto tr295
		case 117:
			goto st196
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr194
		}
		goto st0
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		if data[p] == 115 {
			goto st197
		}
		goto st0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		if data[p] == 116 {
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
			goto tr194
		case 46:
			goto tr295
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr194
		}
		goto st0
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		if data[p] == 101 {
			goto st200
		}
		goto st0
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		if data[p] == 99 {
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
			goto tr202
		case 46:
			goto tr301
		case 101:
			goto st202
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr202
		}
		goto st0
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		if data[p] == 109 {
			goto st203
		}
		goto st0
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		if data[p] == 98 {
			goto st204
		}
		goto st0
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		if data[p] == 101 {
			goto st205
		}
		goto st0
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		if data[p] == 114 {
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
			goto tr202
		case 46:
			goto tr301
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr202
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
		if data[p] == 98 {
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
			goto tr212
		case 46:
			goto tr309
		case 114:
			goto st210
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr212
		}
		goto st0
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		if data[p] == 117 {
			goto st211
		}
		goto st0
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		if data[p] == 97 {
			goto st212
		}
		goto st0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		if data[p] == 114 {
			goto st213
		}
		goto st0
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		if data[p] == 121 {
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
			goto tr212
		case 46:
			goto tr309
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr212
		}
		goto st0
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		switch data[p] {
		case 97:
			goto st216
		case 117:
			goto st222
		}
		goto st0
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		if data[p] == 110 {
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
			goto tr318
		case 117:
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
		if data[p] == 97 {
			goto st219
		}
		goto st0
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		if data[p] == 114 {
			goto st220
		}
		goto st0
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		if data[p] == 121 {
			goto st221
		}
		goto st0
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		switch data[p] {
		case 32:
			goto tr223
		case 46:
			goto tr318
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr223
		}
		goto st0
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 108:
			goto st223
		case 110:
			goto st225
		}
		goto st0
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 32:
			goto tr232
		case 46:
			goto tr325
		case 121:
			goto st224
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr232
		}
		goto st0
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 32:
			goto tr232
		case 46:
			goto tr325
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr232
		}
		goto st0
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 32:
			goto tr236
		case 46:
			goto tr327
		case 101:
			goto st226
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr236
		}
		goto st0
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 32:
			goto tr236
		case 46:
			goto tr327
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr236
		}
		goto st0
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		if data[p] == 97 {
			goto st228
		}
		goto st0
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		switch data[p] {
		case 114:
			goto st229
		case 121:
			goto st232
		}
		goto st0
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
		switch data[p] {
		case 32:
			goto tr243
		case 46:
			goto tr332
		case 99:
			goto st230
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr243
		}
		goto st0
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		if data[p] == 104 {
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
			goto tr243
		case 46:
			goto tr332
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
			goto tr248
		case 46:
			goto tr335
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr248
		}
		goto st0
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		if data[p] == 111 {
			goto st234
		}
		goto st0
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		if data[p] == 118 {
			goto st235
		}
		goto st0
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		switch data[p] {
		case 32:
			goto tr253
		case 46:
			goto tr338
		case 101:
			goto st236
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr253
		}
		goto st0
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		if data[p] == 109 {
			goto st237
		}
		goto st0
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		if data[p] == 98 {
			goto st238
		}
		goto st0
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		if data[p] == 101 {
			goto st239
		}
		goto st0
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		if data[p] == 114 {
			goto st240
		}
		goto st0
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		switch data[p] {
		case 32:
			goto tr253
		case 46:
			goto tr338
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr253
		}
		goto st0
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		if data[p] == 99 {
			goto st242
		}
		goto st0
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		if data[p] == 116 {
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
			goto tr263
		case 46:
			goto tr346
		case 111:
			goto st244
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr263
		}
		goto st0
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		if data[p] == 98 {
			goto st245
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
		if data[p] == 114 {
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
			goto tr263
		case 46:
			goto tr346
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr263
		}
		goto st0
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		if data[p] == 101 {
			goto st249
		}
		goto st0
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		if data[p] == 112 {
			goto st250
		}
		goto st0
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 32:
			goto tr272
		case 46:
			goto tr353
		case 116:
			goto st251
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr272
		}
		goto st0
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		switch data[p] {
		case 32:
			goto tr272
		case 46:
			goto tr353
		case 101:
			goto st252
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr272
		}
		goto st0
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		if data[p] == 109 {
			goto st253
		}
		goto st0
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		if data[p] == 98 {
			goto st254
		}
		goto st0
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		if data[p] == 101 {
			goto st255
		}
		goto st0
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		if data[p] == 114 {
			goto st256
		}
		goto st0
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		switch data[p] {
		case 32:
			goto tr272
		case 46:
			goto tr353
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr272
		}
		goto st0
tr2:
//line ragel/datetime.rl:5
 pb = p 
	goto st257
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
//line ragel/parse_datetime.go:8668
		if data[p] == 32 {
			goto tr166
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st110
			}
		case data[p] >= 45:
			goto tr167
		}
		goto st0
tr3:
//line ragel/datetime.rl:5
 pb = p 
	goto st258
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
//line ragel/parse_datetime.go:8690
		if data[p] == 32 {
			goto tr166
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr167
			}
		case data[p] > 49:
			if 50 <= data[p] && data[p] <= 57 {
				goto st3
			}
		default:
			goto st110
		}
		goto st0
tr4:
//line ragel/datetime.rl:5
 pb = p 
	goto st259
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
//line ragel/parse_datetime.go:8716
		if data[p] == 32 {
			goto tr166
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] >= 45:
			goto tr167
		}
		goto st0
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 112:
			goto st261
		case 117:
			goto st349
		}
		goto st0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		if data[p] == 114 {
			goto st262
		}
		goto st0
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		switch data[p] {
		case 32:
			goto tr363
		case 46:
			goto tr365
		case 105:
			goto st347
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr364
		}
		goto st0
tr363:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st263
tr491:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st263
tr499:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st263
tr510:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st263
tr1012:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st263
tr1021:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st263
tr1025:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st263
tr1033:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st263
tr1038:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st263
tr1043:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st263
tr1053:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st263
tr1066:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st263
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
//line ragel/parse_datetime.go:8820
		switch data[p] {
		case 48:
			goto tr367
		case 51:
			goto tr369
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr370
			}
		case data[p] >= 49:
			goto tr368
		}
		goto st0
tr367:
//line ragel/datetime.rl:5
 pb = p 
	goto st264
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
//line ragel/parse_datetime.go:8845
		if 49 <= data[p] && data[p] <= 57 {
			goto st265
		}
		goto st0
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		switch data[p] {
		case 32:
			goto tr372
		case 44:
			goto tr373
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr374
		}
		goto st0
tr372:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st266
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
//line ragel/parse_datetime.go:8881
		if data[p] == 50 {
			goto tr376
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr377
			}
		case data[p] >= 48:
			goto tr375
		}
		goto st0
tr375:
//line ragel/datetime.rl:5
 pb = p 
	goto st267
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
//line ragel/parse_datetime.go:8903
		switch data[p] {
		case 32:
			goto tr378
		case 58:
			goto tr380
		case 65:
			goto tr381
		case 80:
			goto tr381
		case 97:
			goto tr382
		case 112:
			goto tr382
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st282
		}
		goto st0
tr378:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st268
tr403:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st268
tr456:
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

	goto st268
tr439:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st268
tr444:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st268
tr450:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st268
tr467:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st268
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
//line ragel/parse_datetime.go:8994
		switch data[p] {
		case 65:
			goto tr384
		case 80:
			goto tr384
		case 97:
			goto tr385
		case 112:
			goto tr385
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr383
		}
		goto st0
tr383:
//line ragel/datetime.rl:5
 pb = p 
	goto st269
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
//line ragel/parse_datetime.go:9018
		if 48 <= data[p] && data[p] <= 57 {
			goto st270
		}
		goto st0
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		if 48 <= data[p] && data[p] <= 57 {
			goto st271
		}
		goto st0
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		if 48 <= data[p] && data[p] <= 57 {
			goto st824
		}
		goto st0
	st824:
		if p++; p == pe {
			goto _test_eof824
		}
	st_case_824:
		if data[p] == 32 {
			goto tr1226
		}
		goto st0
tr1226:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st272
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
//line ragel/parse_datetime.go:9061
		switch data[p] {
		case 43:
			goto st273
		case 45:
			goto st277
		case 47:
			goto tr391
		case 90:
			goto tr392
		case 95:
			goto tr391
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr391
			}
		case data[p] >= 65:
			goto tr391
		}
		goto st0
tr1247:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st273
tr1258:
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

	goto st273
tr1262:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st273
tr1277:
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

	goto st273
tr1270:
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

	goto st273
tr1290:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st273
tr1299:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st273
tr1307:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st273
tr1327:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st273
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
//line ragel/parse_datetime.go:9197
		if data[p] == 50 {
			goto tr394
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr395
			}
		case data[p] >= 48:
			goto tr393
		}
		goto st0
tr393:
//line ragel/datetime.rl:5
 pb = p 
	goto st825
tr398:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st825
	st825:
		if p++; p == pe {
			goto _test_eof825
		}
	st_case_825:
//line ragel/parse_datetime.go:9225
		switch data[p] {
		case 32:
			goto tr1227
		case 58:
			goto tr1229
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st827
		}
		goto st0
tr1227:
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

	goto st274
tr1234:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

	goto st274
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
//line ragel/parse_datetime.go:9282
		switch data[p] {
		case 47:
			goto tr391
		case 95:
			goto tr391
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr391
			}
		case data[p] >= 65:
			goto tr391
		}
		goto st0
tr391:
//line ragel/datetime.rl:5
 pb = p 
	goto st275
tr1249:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st275
tr1264:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st275
tr1292:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st275
tr1301:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st275
tr1310:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st275
tr1279:
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
	goto st275
tr1330:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st275
tr1273:
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
	goto st275
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
//line ragel/parse_datetime.go:9407
		switch data[p] {
		case 47:
			goto st276
		case 95:
			goto st276
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st276
			}
		case data[p] >= 65:
			goto st276
		}
		goto st0
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		switch data[p] {
		case 47:
			goto st826
		case 95:
			goto st826
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st826
			}
		case data[p] >= 65:
			goto st826
		}
		goto st0
tr1260:
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
	goto st826
	st826:
		if p++; p == pe {
			goto _test_eof826
		}
	st_case_826:
//line ragel/parse_datetime.go:9475
		switch data[p] {
		case 47:
			goto st826
		case 95:
			goto st826
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st826
			}
		case data[p] >= 65:
			goto st826
		}
		goto st0
tr395:
//line ragel/datetime.rl:5
 pb = p 
	goto st827
tr400:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st827
	st827:
		if p++; p == pe {
			goto _test_eof827
		}
	st_case_827:
//line ragel/parse_datetime.go:9506
		switch data[p] {
		case 32:
			goto tr1227
		case 58:
			goto tr1229
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st828
		}
		goto st0
	st828:
		if p++; p == pe {
			goto _test_eof828
		}
	st_case_828:
		if data[p] == 32 {
			goto tr1227
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st828
		}
		goto st0
tr1229:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st829
	st829:
		if p++; p == pe {
			goto _test_eof829
		}
	st_case_829:
//line ragel/parse_datetime.go:9545
		if data[p] == 32 {
			goto st274
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1233
			}
		case data[p] >= 48:
			goto tr1232
		}
		goto st0
tr1232:
//line ragel/datetime.rl:5
 pb = p 
	goto st830
	st830:
		if p++; p == pe {
			goto _test_eof830
		}
	st_case_830:
//line ragel/parse_datetime.go:9567
		if data[p] == 32 {
			goto tr1234
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st831
		}
		goto st0
tr1233:
//line ragel/datetime.rl:5
 pb = p 
	goto st831
	st831:
		if p++; p == pe {
			goto _test_eof831
		}
	st_case_831:
//line ragel/parse_datetime.go:9584
		if data[p] == 32 {
			goto tr1234
		}
		goto st0
tr394:
//line ragel/datetime.rl:5
 pb = p 
	goto st832
tr399:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st832
	st832:
		if p++; p == pe {
			goto _test_eof832
		}
	st_case_832:
//line ragel/parse_datetime.go:9604
		switch data[p] {
		case 32:
			goto tr1227
		case 58:
			goto tr1229
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st828
			}
		case data[p] >= 48:
			goto st827
		}
		goto st0
tr1248:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st277
tr1259:
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

	goto st277
tr1263:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st277
tr1278:
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

	goto st277
tr1272:
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

	goto st277
tr1291:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st277
tr1300:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st277
tr1309:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st277
tr1329:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st277
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
//line ragel/parse_datetime.go:9734
		if data[p] == 50 {
			goto tr399
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr400
			}
		case data[p] >= 48:
			goto tr398
		}
		goto st0
tr392:
//line ragel/datetime.rl:5
 pb = p 
	goto st833
tr1253:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st833
tr1267:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st833
tr1296:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st833
tr1304:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st833
tr1313:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st833
tr1281:
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
	goto st833
tr1332:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st833
tr1275:
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
	goto st833
	st833:
		if p++; p == pe {
			goto _test_eof833
		}
	st_case_833:
//line ragel/parse_datetime.go:9856
		switch data[p] {
		case 47:
			goto st276
		case 95:
			goto st276
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st276
			}
		case data[p] >= 65:
			goto st276
		}
		goto st0
tr384:
//line ragel/datetime.rl:5
 pb = p 
	goto st278
tr381:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st278
tr406:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st278
tr442:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st278
tr446:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st278
tr453:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st278
tr458:
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
	goto st278
tr469:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st278
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
//line ragel/parse_datetime.go:9962
		if data[p] == 77 {
			goto st279
		}
		goto st0
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		if data[p] == 32 {
			goto tr402
		}
		goto st0
tr402:
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

	goto st280
tr424:
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

	goto st280
tr435:
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

	goto st280
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
//line ragel/parse_datetime.go:10054
		if 48 <= data[p] && data[p] <= 57 {
			goto tr383
		}
		goto st0
tr385:
//line ragel/datetime.rl:5
 pb = p 
	goto st281
tr382:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st281
tr407:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st281
tr443:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st281
tr447:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st281
tr454:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st281
tr459:
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
	goto st281
tr470:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st281
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
//line ragel/parse_datetime.go:10149
		if data[p] == 109 {
			goto st279
		}
		goto st0
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		switch data[p] {
		case 32:
			goto tr403
		case 58:
			goto tr405
		case 65:
			goto tr406
		case 80:
			goto tr406
		case 97:
			goto tr407
		case 112:
			goto tr407
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st283
		}
		goto st0
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		if 48 <= data[p] && data[p] <= 57 {
			goto st834
		}
		goto st0
	st834:
		if p++; p == pe {
			goto _test_eof834
		}
	st_case_834:
		switch data[p] {
		case 32:
			goto tr1236
		case 43:
			goto tr1220
		case 44:
			goto tr1237
		case 45:
			goto tr1221
		case 46:
			goto tr436
		case 47:
			goto tr1222
		case 84:
			goto tr1223
		case 90:
			goto tr1224
		case 95:
			goto tr1225
		case 116:
			goto tr1225
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st306
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1222
			}
		default:
			goto tr1222
		}
		goto st0
tr1236:
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

	goto st835
	st835:
		if p++; p == pe {
			goto _test_eof835
		}
	st_case_835:
//line ragel/parse_datetime.go:10252
		switch data[p] {
		case 32:
			goto st11
		case 43:
			goto st14
		case 45:
			goto st20
		case 47:
			goto tr50
		case 50:
			goto tr1240
		case 65:
			goto tr51
		case 66:
			goto tr52
		case 90:
			goto tr1093
		case 95:
			goto tr50
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr1239
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr50
				}
			case data[p] >= 67:
				goto tr50
			}
		default:
			goto tr1241
		}
		goto st0
tr1239:
//line ragel/datetime.rl:5
 pb = p 
	goto st836
	st836:
		if p++; p == pe {
			goto _test_eof836
		}
	st_case_836:
//line ragel/parse_datetime.go:10300
		switch data[p] {
		case 32:
			goto tr1104
		case 43:
			goto tr1105
		case 45:
			goto tr1106
		case 47:
			goto tr1107
		case 58:
			goto tr1109
		case 65:
			goto tr1110
		case 80:
			goto tr1110
		case 90:
			goto tr1111
		case 95:
			goto tr1107
		case 97:
			goto tr1112
		case 112:
			goto tr1112
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st837
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1107
			}
		default:
			goto tr1107
		}
		goto st0
	st837:
		if p++; p == pe {
			goto _test_eof837
		}
	st_case_837:
		switch data[p] {
		case 32:
			goto tr1121
		case 43:
			goto tr1122
		case 45:
			goto tr1123
		case 47:
			goto tr1124
		case 58:
			goto tr1125
		case 65:
			goto tr1126
		case 80:
			goto tr1126
		case 90:
			goto tr1127
		case 95:
			goto tr1124
		case 97:
			goto tr1128
		case 112:
			goto tr1128
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st284
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1124
			}
		default:
			goto tr1124
		}
		goto st0
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		if 48 <= data[p] && data[p] <= 57 {
			goto st838
		}
		goto st0
	st838:
		if p++; p == pe {
			goto _test_eof838
		}
	st_case_838:
		switch data[p] {
		case 32:
			goto tr1243
		case 43:
			goto tr1130
		case 45:
			goto tr1132
		case 47:
			goto tr1133
		case 90:
			goto tr1135
		case 95:
			goto tr1133
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1131
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1133
				}
			case data[p] >= 65:
				goto tr1133
			}
		default:
			goto st26
		}
		goto st0
tr1240:
//line ragel/datetime.rl:5
 pb = p 
	goto st839
	st839:
		if p++; p == pe {
			goto _test_eof839
		}
	st_case_839:
//line ragel/parse_datetime.go:10435
		switch data[p] {
		case 32:
			goto tr1104
		case 43:
			goto tr1105
		case 45:
			goto tr1106
		case 47:
			goto tr1107
		case 58:
			goto tr1109
		case 65:
			goto tr1110
		case 80:
			goto tr1110
		case 90:
			goto tr1111
		case 95:
			goto tr1107
		case 97:
			goto tr1112
		case 112:
			goto tr1112
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st837
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1107
				}
			case data[p] >= 66:
				goto tr1107
			}
		default:
			goto st285
		}
		goto st0
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		if 48 <= data[p] && data[p] <= 57 {
			goto st284
		}
		goto st0
tr1241:
//line ragel/datetime.rl:5
 pb = p 
	goto st840
	st840:
		if p++; p == pe {
			goto _test_eof840
		}
	st_case_840:
//line ragel/parse_datetime.go:10496
		switch data[p] {
		case 32:
			goto tr1104
		case 43:
			goto tr1105
		case 45:
			goto tr1106
		case 47:
			goto tr1107
		case 58:
			goto tr1109
		case 65:
			goto tr1110
		case 80:
			goto tr1110
		case 90:
			goto tr1111
		case 95:
			goto tr1107
		case 97:
			goto tr1112
		case 112:
			goto tr1112
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st285
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1107
			}
		default:
			goto tr1107
		}
		goto st0
tr1237:
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

	goto st841
	st841:
		if p++; p == pe {
			goto _test_eof841
		}
	st_case_841:
//line ragel/parse_datetime.go:10560
		if data[p] == 32 {
			goto st286
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr433
		}
		goto st0
tr1335:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st286
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
//line ragel/parse_datetime.go:10579
		if data[p] == 50 {
			goto tr412
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr413
			}
		case data[p] >= 48:
			goto tr411
		}
		goto st0
tr411:
//line ragel/datetime.rl:5
 pb = p 
	goto st842
	st842:
		if p++; p == pe {
			goto _test_eof842
		}
	st_case_842:
//line ragel/parse_datetime.go:10601
		switch data[p] {
		case 32:
			goto tr1246
		case 43:
			goto tr1247
		case 45:
			goto tr1248
		case 47:
			goto tr1249
		case 58:
			goto tr1251
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1253
		case 95:
			goto tr1249
		case 97:
			goto tr1254
		case 112:
			goto tr1254
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st846
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1249
			}
		default:
			goto tr1249
		}
		goto st0
tr1246:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st843
tr1261:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st843
tr1315:
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

	goto st843
tr1289:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st843
tr1298:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st843
tr1306:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st843
tr1326:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st843
	st843:
		if p++; p == pe {
			goto _test_eof843
		}
	st_case_843:
//line ragel/parse_datetime.go:10711
		switch data[p] {
		case 43:
			goto st273
		case 45:
			goto st277
		case 47:
			goto tr391
		case 65:
			goto tr1255
		case 80:
			goto tr1255
		case 90:
			goto tr392
		case 95:
			goto tr391
		case 97:
			goto tr1256
		case 112:
			goto tr1256
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr391
			}
		case data[p] >= 66:
			goto tr391
		}
		goto st0
tr1255:
//line ragel/datetime.rl:5
 pb = p 
	goto st287
tr1252:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st287
tr1266:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st287
tr1295:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st287
tr1303:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st287
tr1312:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st287
tr1317:
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
	goto st287
tr1331:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st287
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
//line ragel/parse_datetime.go:10831
		switch data[p] {
		case 47:
			goto st276
		case 77:
			goto st844
		case 95:
			goto st276
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st276
			}
		case data[p] >= 65:
			goto st276
		}
		goto st0
	st844:
		if p++; p == pe {
			goto _test_eof844
		}
	st_case_844:
		switch data[p] {
		case 32:
			goto tr1257
		case 43:
			goto tr1258
		case 45:
			goto tr1259
		case 47:
			goto tr1260
		case 95:
			goto tr1260
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1260
			}
		case data[p] >= 65:
			goto tr1260
		}
		goto st0
tr1257:
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

	goto st845
tr1276:
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

	goto st845
tr1269:
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

	goto st845
	st845:
		if p++; p == pe {
			goto _test_eof845
		}
	st_case_845:
//line ragel/parse_datetime.go:10953
		switch data[p] {
		case 43:
			goto st273
		case 45:
			goto st277
		case 47:
			goto tr391
		case 90:
			goto tr392
		case 95:
			goto tr391
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr391
			}
		case data[p] >= 65:
			goto tr391
		}
		goto st0
tr1256:
//line ragel/datetime.rl:5
 pb = p 
	goto st288
tr1254:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st288
tr1268:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st288
tr1297:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st288
tr1305:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st288
tr1314:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st288
tr1318:
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
	goto st288
tr1333:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st288
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
//line ragel/parse_datetime.go:11065
		switch data[p] {
		case 47:
			goto st276
		case 95:
			goto st276
		case 109:
			goto st844
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st276
			}
		case data[p] >= 65:
			goto st276
		}
		goto st0
	st846:
		if p++; p == pe {
			goto _test_eof846
		}
	st_case_846:
		switch data[p] {
		case 32:
			goto tr1261
		case 43:
			goto tr1262
		case 45:
			goto tr1263
		case 47:
			goto tr1264
		case 58:
			goto tr1265
		case 65:
			goto tr1266
		case 80:
			goto tr1266
		case 90:
			goto tr1267
		case 95:
			goto tr1264
		case 97:
			goto tr1268
		case 112:
			goto tr1268
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st289
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1264
			}
		default:
			goto tr1264
		}
		goto st0
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr1269
		case 43:
			goto tr1270
		case 45:
			goto tr1272
		case 47:
			goto tr1273
		case 90:
			goto tr1275
		case 95:
			goto tr1273
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1271
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1273
				}
			case data[p] >= 65:
				goto tr1273
			}
		default:
			goto st291
		}
		goto st0
tr1271:
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

	goto st290
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
//line ragel/parse_datetime.go:11193
		if 48 <= data[p] && data[p] <= 57 {
			goto tr416
		}
		goto st0
tr416:
//line ragel/datetime.rl:5
 pb = p 
	goto st848
	st848:
		if p++; p == pe {
			goto _test_eof848
		}
	st_case_848:
//line ragel/parse_datetime.go:11207
		switch data[p] {
		case 32:
			goto tr1276
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st849
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st849:
		if p++; p == pe {
			goto _test_eof849
		}
	st_case_849:
		switch data[p] {
		case 32:
			goto tr1276
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st850
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st850:
		if p++; p == pe {
			goto _test_eof850
		}
	st_case_850:
		switch data[p] {
		case 32:
			goto tr1276
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st851
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st851:
		if p++; p == pe {
			goto _test_eof851
		}
	st_case_851:
		switch data[p] {
		case 32:
			goto tr1276
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st852
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st852:
		if p++; p == pe {
			goto _test_eof852
		}
	st_case_852:
		switch data[p] {
		case 32:
			goto tr1276
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st853
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st853:
		if p++; p == pe {
			goto _test_eof853
		}
	st_case_853:
		switch data[p] {
		case 32:
			goto tr1276
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st854
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st854:
		if p++; p == pe {
			goto _test_eof854
		}
	st_case_854:
		switch data[p] {
		case 32:
			goto tr1276
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st855
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st855:
		if p++; p == pe {
			goto _test_eof855
		}
	st_case_855:
		switch data[p] {
		case 32:
			goto tr1276
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st856
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st856:
		if p++; p == pe {
			goto _test_eof856
		}
	st_case_856:
		switch data[p] {
		case 32:
			goto tr1276
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		case data[p] >= 65:
			goto tr1279
		}
		goto st0
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		if 48 <= data[p] && data[p] <= 57 {
			goto st857
		}
		goto st0
	st857:
		if p++; p == pe {
			goto _test_eof857
		}
	st_case_857:
		switch data[p] {
		case 32:
			goto tr1269
		case 43:
			goto tr1270
		case 45:
			goto tr1272
		case 47:
			goto tr1273
		case 90:
			goto tr1275
		case 95:
			goto tr1273
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1271
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1273
			}
		default:
			goto tr1273
		}
		goto st0
tr1251:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st292
tr1265:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st292
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
//line ragel/parse_datetime.go:11545
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr419
			}
		case data[p] >= 48:
			goto tr418
		}
		goto st0
tr418:
//line ragel/datetime.rl:5
 pb = p 
	goto st858
	st858:
		if p++; p == pe {
			goto _test_eof858
		}
	st_case_858:
//line ragel/parse_datetime.go:11564
		switch data[p] {
		case 32:
			goto tr1289
		case 43:
			goto tr1290
		case 45:
			goto tr1291
		case 47:
			goto tr1292
		case 58:
			goto tr1294
		case 65:
			goto tr1295
		case 80:
			goto tr1295
		case 90:
			goto tr1296
		case 95:
			goto tr1292
		case 97:
			goto tr1297
		case 112:
			goto tr1297
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st859
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1292
			}
		default:
			goto tr1292
		}
		goto st0
	st859:
		if p++; p == pe {
			goto _test_eof859
		}
	st_case_859:
		switch data[p] {
		case 32:
			goto tr1298
		case 43:
			goto tr1299
		case 45:
			goto tr1300
		case 47:
			goto tr1301
		case 58:
			goto tr1302
		case 65:
			goto tr1303
		case 80:
			goto tr1303
		case 90:
			goto tr1304
		case 95:
			goto tr1301
		case 97:
			goto tr1305
		case 112:
			goto tr1305
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1301
			}
		case data[p] >= 66:
			goto tr1301
		}
		goto st0
tr1294:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st293
tr1302:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st293
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
//line ragel/parse_datetime.go:11657
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr421
			}
		case data[p] >= 48:
			goto tr420
		}
		goto st0
tr420:
//line ragel/datetime.rl:5
 pb = p 
	goto st860
	st860:
		if p++; p == pe {
			goto _test_eof860
		}
	st_case_860:
//line ragel/parse_datetime.go:11676
		switch data[p] {
		case 32:
			goto tr1306
		case 43:
			goto tr1307
		case 45:
			goto tr1309
		case 47:
			goto tr1310
		case 65:
			goto tr1312
		case 80:
			goto tr1312
		case 90:
			goto tr1313
		case 95:
			goto tr1310
		case 97:
			goto tr1314
		case 112:
			goto tr1314
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1308
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1310
				}
			case data[p] >= 66:
				goto tr1310
			}
		default:
			goto st870
		}
		goto st0
tr1308:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st294
tr1328:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st294
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
//line ragel/parse_datetime.go:11734
		if 48 <= data[p] && data[p] <= 57 {
			goto tr422
		}
		goto st0
tr422:
//line ragel/datetime.rl:5
 pb = p 
	goto st861
	st861:
		if p++; p == pe {
			goto _test_eof861
		}
	st_case_861:
//line ragel/parse_datetime.go:11748
		switch data[p] {
		case 32:
			goto tr1315
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 65:
			goto tr1317
		case 80:
			goto tr1317
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		case 97:
			goto tr1318
		case 112:
			goto tr1318
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st862
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st862:
		if p++; p == pe {
			goto _test_eof862
		}
	st_case_862:
		switch data[p] {
		case 32:
			goto tr1315
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 65:
			goto tr1317
		case 80:
			goto tr1317
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		case 97:
			goto tr1318
		case 112:
			goto tr1318
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st863
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st863:
		if p++; p == pe {
			goto _test_eof863
		}
	st_case_863:
		switch data[p] {
		case 32:
			goto tr1315
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 65:
			goto tr1317
		case 80:
			goto tr1317
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		case 97:
			goto tr1318
		case 112:
			goto tr1318
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st864
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st864:
		if p++; p == pe {
			goto _test_eof864
		}
	st_case_864:
		switch data[p] {
		case 32:
			goto tr1315
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 65:
			goto tr1317
		case 80:
			goto tr1317
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		case 97:
			goto tr1318
		case 112:
			goto tr1318
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st865
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st865:
		if p++; p == pe {
			goto _test_eof865
		}
	st_case_865:
		switch data[p] {
		case 32:
			goto tr1315
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 65:
			goto tr1317
		case 80:
			goto tr1317
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		case 97:
			goto tr1318
		case 112:
			goto tr1318
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st866
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st866:
		if p++; p == pe {
			goto _test_eof866
		}
	st_case_866:
		switch data[p] {
		case 32:
			goto tr1315
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 65:
			goto tr1317
		case 80:
			goto tr1317
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		case 97:
			goto tr1318
		case 112:
			goto tr1318
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st867
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st867:
		if p++; p == pe {
			goto _test_eof867
		}
	st_case_867:
		switch data[p] {
		case 32:
			goto tr1315
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 65:
			goto tr1317
		case 80:
			goto tr1317
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		case 97:
			goto tr1318
		case 112:
			goto tr1318
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st868
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st868:
		if p++; p == pe {
			goto _test_eof868
		}
	st_case_868:
		switch data[p] {
		case 32:
			goto tr1315
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 65:
			goto tr1317
		case 80:
			goto tr1317
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		case 97:
			goto tr1318
		case 112:
			goto tr1318
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st869
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		default:
			goto tr1279
		}
		goto st0
	st869:
		if p++; p == pe {
			goto _test_eof869
		}
	st_case_869:
		switch data[p] {
		case 32:
			goto tr1315
		case 43:
			goto tr1277
		case 45:
			goto tr1278
		case 47:
			goto tr1279
		case 65:
			goto tr1317
		case 80:
			goto tr1317
		case 90:
			goto tr1281
		case 95:
			goto tr1279
		case 97:
			goto tr1318
		case 112:
			goto tr1318
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1279
			}
		case data[p] >= 66:
			goto tr1279
		}
		goto st0
	st870:
		if p++; p == pe {
			goto _test_eof870
		}
	st_case_870:
		switch data[p] {
		case 32:
			goto tr1326
		case 43:
			goto tr1327
		case 45:
			goto tr1329
		case 47:
			goto tr1330
		case 65:
			goto tr1331
		case 80:
			goto tr1331
		case 90:
			goto tr1332
		case 95:
			goto tr1330
		case 97:
			goto tr1333
		case 112:
			goto tr1333
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1328
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1330
			}
		default:
			goto tr1330
		}
		goto st0
tr421:
//line ragel/datetime.rl:5
 pb = p 
	goto st871
	st871:
		if p++; p == pe {
			goto _test_eof871
		}
	st_case_871:
//line ragel/parse_datetime.go:12149
		switch data[p] {
		case 32:
			goto tr1306
		case 43:
			goto tr1307
		case 45:
			goto tr1309
		case 47:
			goto tr1310
		case 65:
			goto tr1312
		case 80:
			goto tr1312
		case 90:
			goto tr1313
		case 95:
			goto tr1310
		case 97:
			goto tr1314
		case 112:
			goto tr1314
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1308
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1310
			}
		default:
			goto tr1310
		}
		goto st0
tr419:
//line ragel/datetime.rl:5
 pb = p 
	goto st872
	st872:
		if p++; p == pe {
			goto _test_eof872
		}
	st_case_872:
//line ragel/parse_datetime.go:12194
		switch data[p] {
		case 32:
			goto tr1289
		case 43:
			goto tr1290
		case 45:
			goto tr1291
		case 47:
			goto tr1292
		case 58:
			goto tr1294
		case 65:
			goto tr1295
		case 80:
			goto tr1295
		case 90:
			goto tr1296
		case 95:
			goto tr1292
		case 97:
			goto tr1297
		case 112:
			goto tr1297
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1292
			}
		case data[p] >= 66:
			goto tr1292
		}
		goto st0
tr412:
//line ragel/datetime.rl:5
 pb = p 
	goto st873
	st873:
		if p++; p == pe {
			goto _test_eof873
		}
	st_case_873:
//line ragel/parse_datetime.go:12237
		switch data[p] {
		case 32:
			goto tr1246
		case 43:
			goto tr1247
		case 45:
			goto tr1248
		case 47:
			goto tr1249
		case 58:
			goto tr1251
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1253
		case 95:
			goto tr1249
		case 97:
			goto tr1254
		case 112:
			goto tr1254
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st846
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1249
				}
			case data[p] >= 66:
				goto tr1249
			}
		default:
			goto st295
		}
		goto st0
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		if 48 <= data[p] && data[p] <= 57 {
			goto st289
		}
		goto st0
tr413:
//line ragel/datetime.rl:5
 pb = p 
	goto st874
	st874:
		if p++; p == pe {
			goto _test_eof874
		}
	st_case_874:
//line ragel/parse_datetime.go:12298
		switch data[p] {
		case 32:
			goto tr1246
		case 43:
			goto tr1247
		case 45:
			goto tr1248
		case 47:
			goto tr1249
		case 58:
			goto tr1251
		case 65:
			goto tr1252
		case 80:
			goto tr1252
		case 90:
			goto tr1253
		case 95:
			goto tr1249
		case 97:
			goto tr1254
		case 112:
			goto tr1254
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st295
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1249
			}
		default:
			goto tr1249
		}
		goto st0
tr433:
//line ragel/datetime.rl:5
 pb = p 
	goto st296
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
//line ragel/parse_datetime.go:12345
		if data[p] == 32 {
			goto tr424
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st297
		}
		goto st0
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		if data[p] == 32 {
			goto tr424
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st298
		}
		goto st0
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		if data[p] == 32 {
			goto tr424
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
			goto tr424
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
			goto tr424
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
			goto tr424
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
			goto tr424
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
			goto tr424
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
			goto tr424
		}
		goto st0
tr436:
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

	goto st305
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
//line ragel/parse_datetime.go:12468
		if 48 <= data[p] && data[p] <= 57 {
			goto tr433
		}
		goto st0
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		if 48 <= data[p] && data[p] <= 57 {
			goto st307
		}
		goto st0
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		switch data[p] {
		case 32:
			goto tr435
		case 44:
			goto tr436
		case 46:
			goto tr436
		}
		goto st0
tr380:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st308
tr405:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st308
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
//line ragel/parse_datetime.go:12513
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr438
			}
		case data[p] >= 48:
			goto tr437
		}
		goto st0
tr437:
//line ragel/datetime.rl:5
 pb = p 
	goto st309
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
//line ragel/parse_datetime.go:12532
		switch data[p] {
		case 32:
			goto tr439
		case 58:
			goto tr441
		case 65:
			goto tr442
		case 80:
			goto tr442
		case 97:
			goto tr443
		case 112:
			goto tr443
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st310
		}
		goto st0
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		switch data[p] {
		case 32:
			goto tr444
		case 58:
			goto tr445
		case 65:
			goto tr446
		case 80:
			goto tr446
		case 97:
			goto tr447
		case 112:
			goto tr447
		}
		goto st0
tr441:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st311
tr445:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st311
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
//line ragel/parse_datetime.go:12588
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr449
			}
		case data[p] >= 48:
			goto tr448
		}
		goto st0
tr448:
//line ragel/datetime.rl:5
 pb = p 
	goto st312
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
//line ragel/parse_datetime.go:12607
		switch data[p] {
		case 32:
			goto tr450
		case 44:
			goto tr451
		case 46:
			goto tr451
		case 65:
			goto tr453
		case 80:
			goto tr453
		case 97:
			goto tr454
		case 112:
			goto tr454
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st323
		}
		goto st0
tr451:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st313
tr468:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st313
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
//line ragel/parse_datetime.go:12645
		if 48 <= data[p] && data[p] <= 57 {
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
//line ragel/parse_datetime.go:12659
		switch data[p] {
		case 32:
			goto tr456
		case 65:
			goto tr458
		case 80:
			goto tr458
		case 97:
			goto tr459
		case 112:
			goto tr459
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st315
		}
		goto st0
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		switch data[p] {
		case 32:
			goto tr456
		case 65:
			goto tr458
		case 80:
			goto tr458
		case 97:
			goto tr459
		case 112:
			goto tr459
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st316
		}
		goto st0
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		switch data[p] {
		case 32:
			goto tr456
		case 65:
			goto tr458
		case 80:
			goto tr458
		case 97:
			goto tr459
		case 112:
			goto tr459
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
			goto tr456
		case 65:
			goto tr458
		case 80:
			goto tr458
		case 97:
			goto tr459
		case 112:
			goto tr459
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
			goto tr456
		case 65:
			goto tr458
		case 80:
			goto tr458
		case 97:
			goto tr459
		case 112:
			goto tr459
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
			goto tr456
		case 65:
			goto tr458
		case 80:
			goto tr458
		case 97:
			goto tr459
		case 112:
			goto tr459
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
			goto tr456
		case 65:
			goto tr458
		case 80:
			goto tr458
		case 97:
			goto tr459
		case 112:
			goto tr459
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
			goto tr456
		case 65:
			goto tr458
		case 80:
			goto tr458
		case 97:
			goto tr459
		case 112:
			goto tr459
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
			goto tr456
		case 65:
			goto tr458
		case 80:
			goto tr458
		case 97:
			goto tr459
		case 112:
			goto tr459
		}
		goto st0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 32:
			goto tr467
		case 44:
			goto tr468
		case 46:
			goto tr468
		case 65:
			goto tr469
		case 80:
			goto tr469
		case 97:
			goto tr470
		case 112:
			goto tr470
		}
		goto st0
tr449:
//line ragel/datetime.rl:5
 pb = p 
	goto st324
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
//line ragel/parse_datetime.go:12872
		switch data[p] {
		case 32:
			goto tr450
		case 44:
			goto tr451
		case 46:
			goto tr451
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
tr438:
//line ragel/datetime.rl:5
 pb = p 
	goto st325
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
//line ragel/parse_datetime.go:12899
		switch data[p] {
		case 32:
			goto tr439
		case 58:
			goto tr441
		case 65:
			goto tr442
		case 80:
			goto tr442
		case 97:
			goto tr443
		case 112:
			goto tr443
		}
		goto st0
tr376:
//line ragel/datetime.rl:5
 pb = p 
	goto st326
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
//line ragel/parse_datetime.go:12924
		switch data[p] {
		case 32:
			goto tr378
		case 58:
			goto tr380
		case 65:
			goto tr381
		case 80:
			goto tr381
		case 97:
			goto tr382
		case 112:
			goto tr382
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st327
			}
		case data[p] >= 48:
			goto st282
		}
		goto st0
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		if 48 <= data[p] && data[p] <= 57 {
			goto st283
		}
		goto st0
tr377:
//line ragel/datetime.rl:5
 pb = p 
	goto st328
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
//line ragel/parse_datetime.go:12966
		switch data[p] {
		case 32:
			goto tr378
		case 58:
			goto tr380
		case 65:
			goto tr381
		case 80:
			goto tr381
		case 97:
			goto tr382
		case 112:
			goto tr382
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st327
		}
		goto st0
tr373:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st329
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
//line ragel/parse_datetime.go:13001
		if data[p] == 32 {
			goto st330
		}
		goto st0
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr473
		}
		goto st0
tr473:
//line ragel/datetime.rl:5
 pb = p 
	goto st331
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
//line ragel/parse_datetime.go:13024
		if 48 <= data[p] && data[p] <= 57 {
			goto st332
		}
		goto st0
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		if 48 <= data[p] && data[p] <= 57 {
			goto st333
		}
		goto st0
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		if 48 <= data[p] && data[p] <= 57 {
			goto st875
		}
		goto st0
	st875:
		if p++; p == pe {
			goto _test_eof875
		}
	st_case_875:
		switch data[p] {
		case 32:
			goto tr1335
		case 44:
			goto tr1336
		}
		goto st0
tr1336:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st876
	st876:
		if p++; p == pe {
			goto _test_eof876
		}
	st_case_876:
//line ragel/parse_datetime.go:13070
		if data[p] == 32 {
			goto st286
		}
		goto st0
tr368:
//line ragel/datetime.rl:5
 pb = p 
	goto st334
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
//line ragel/parse_datetime.go:13084
		switch data[p] {
		case 32:
			goto tr477
		case 44:
			goto tr373
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st265
			}
		case data[p] >= 45:
			goto tr374
		}
		goto st0
tr477:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st335
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
//line ragel/parse_datetime.go:13116
		if 48 <= data[p] && data[p] <= 57 {
			goto tr478
		}
		goto st0
tr478:
//line ragel/datetime.rl:5
 pb = p 
	goto st336
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
//line ragel/parse_datetime.go:13130
		if 48 <= data[p] && data[p] <= 57 {
			goto st337
		}
		goto st0
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		if 48 <= data[p] && data[p] <= 57 {
			goto st338
		}
		goto st0
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		if 48 <= data[p] && data[p] <= 57 {
			goto st877
		}
		goto st0
	st877:
		if p++; p == pe {
			goto _test_eof877
		}
	st_case_877:
		switch data[p] {
		case 32:
			goto tr1219
		case 43:
			goto tr1220
		case 44:
			goto tr1336
		case 45:
			goto tr1221
		case 47:
			goto tr1222
		case 84:
			goto tr1223
		case 90:
			goto tr1224
		case 95:
			goto tr1225
		case 116:
			goto tr1225
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1222
			}
		case data[p] >= 65:
			goto tr1222
		}
		goto st0
tr369:
//line ragel/datetime.rl:5
 pb = p 
	goto st339
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
//line ragel/parse_datetime.go:13196
		switch data[p] {
		case 32:
			goto tr477
		case 44:
			goto tr373
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st265
			}
		case data[p] >= 45:
			goto tr374
		}
		goto st0
tr370:
//line ragel/datetime.rl:5
 pb = p 
	goto st340
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
//line ragel/parse_datetime.go:13221
		switch data[p] {
		case 32:
			goto tr477
		case 44:
			goto tr373
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr374
		}
		goto st0
tr364:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st341
tr492:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st341
tr500:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st341
tr511:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st341
tr1013:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st341
tr1022:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st341
tr1026:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st341
tr1034:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st341
tr1039:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st341
tr1044:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st341
tr1054:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st341
tr1067:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st341
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
//line ragel/parse_datetime.go:13285
		switch data[p] {
		case 48:
			goto tr482
		case 51:
			goto tr484
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr485
			}
		case data[p] >= 49:
			goto tr483
		}
		goto st0
tr482:
//line ragel/datetime.rl:5
 pb = p 
	goto st342
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
//line ragel/parse_datetime.go:13310
		if 49 <= data[p] && data[p] <= 57 {
			goto st343
		}
		goto st0
tr485:
//line ragel/datetime.rl:5
 pb = p 
	goto st343
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
//line ragel/parse_datetime.go:13324
		if data[p] == 32 {
			goto tr374
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr374
		}
		goto st0
tr483:
//line ragel/datetime.rl:5
 pb = p 
	goto st344
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
//line ragel/parse_datetime.go:13341
		if data[p] == 32 {
			goto tr374
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st343
			}
		case data[p] >= 45:
			goto tr374
		}
		goto st0
tr484:
//line ragel/datetime.rl:5
 pb = p 
	goto st345
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
//line ragel/parse_datetime.go:13363
		if data[p] == 32 {
			goto tr374
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st343
			}
		case data[p] >= 45:
			goto tr374
		}
		goto st0
tr365:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st346
tr493:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st346
tr501:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st346
tr512:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st346
tr1014:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st346
tr1023:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st346
tr1027:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st346
tr1035:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st346
tr1040:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st346
tr1045:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st346
tr1055:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st346
tr1068:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st346
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
//line ragel/parse_datetime.go:13429
		switch data[p] {
		case 32:
			goto st263
		case 48:
			goto tr482
		case 51:
			goto tr484
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st341
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr485
			}
		default:
			goto tr483
		}
		goto st0
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		if data[p] == 108 {
			goto st348
		}
		goto st0
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		switch data[p] {
		case 32:
			goto tr363
		case 46:
			goto tr365
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr364
		}
		goto st0
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		if data[p] == 103 {
			goto st350
		}
		goto st0
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		switch data[p] {
		case 32:
			goto tr491
		case 46:
			goto tr493
		case 117:
			goto st351
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr492
		}
		goto st0
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		if data[p] == 115 {
			goto st352
		}
		goto st0
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		if data[p] == 116 {
			goto st353
		}
		goto st0
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		switch data[p] {
		case 32:
			goto tr491
		case 46:
			goto tr493
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr492
		}
		goto st0
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		if data[p] == 101 {
			goto st355
		}
		goto st0
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		if data[p] == 99 {
			goto st356
		}
		goto st0
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		switch data[p] {
		case 32:
			goto tr499
		case 46:
			goto tr501
		case 101:
			goto st357
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr500
		}
		goto st0
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		if data[p] == 109 {
			goto st358
		}
		goto st0
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		if data[p] == 98 {
			goto st359
		}
		goto st0
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		if data[p] == 101 {
			goto st360
		}
		goto st0
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		if data[p] == 114 {
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
			goto tr499
		case 46:
			goto tr501
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr500
		}
		goto st0
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		switch data[p] {
		case 101:
			goto st363
		case 114:
			goto st370
		}
		goto st0
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		if data[p] == 98 {
			goto st364
		}
		goto st0
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		switch data[p] {
		case 32:
			goto tr510
		case 46:
			goto tr512
		case 114:
			goto st365
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr511
		}
		goto st0
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		if data[p] == 117 {
			goto st366
		}
		goto st0
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		if data[p] == 97 {
			goto st367
		}
		goto st0
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		if data[p] == 114 {
			goto st368
		}
		goto st0
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		if data[p] == 121 {
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
			goto tr510
		case 46:
			goto tr512
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr511
		}
		goto st0
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		if data[p] == 105 {
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
			goto st372
		case 44:
			goto st556
		case 100:
			goto st704
		}
		goto st0
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		switch data[p] {
		case 65:
			goto st373
		case 68:
			goto st498
		case 70:
			goto st506
		case 74:
			goto st514
		case 77:
			goto st526
		case 78:
			goto st532
		case 79:
			goto st540
		case 83:
			goto st547
		case 97:
			goto st373
		case 100:
			goto st498
		case 102:
			goto st506
		case 106:
			goto st514
		case 109:
			goto st526
		case 110:
			goto st532
		case 111:
			goto st540
		case 115:
			goto st547
		}
		goto st0
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		switch data[p] {
		case 112:
			goto st374
		case 117:
			goto st493
		}
		goto st0
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		if data[p] == 114 {
			goto st375
		}
		goto st0
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		switch data[p] {
		case 32:
			goto tr533
		case 46:
			goto tr534
		case 105:
			goto st491
		}
		goto st0
tr533:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st376
tr748:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st376
tr755:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st376
tr764:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st376
tr774:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st376
tr782:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st376
tr785:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st376
tr791:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st376
tr795:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st376
tr799:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st376
tr808:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st376
tr816:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st376
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
//line ragel/parse_datetime.go:13860
		switch data[p] {
		case 32:
			goto st377
		case 48:
			goto tr537
		case 51:
			goto tr539
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr540
			}
		case data[p] >= 49:
			goto tr538
		}
		goto st0
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		switch data[p] {
		case 48:
			goto tr541
		case 51:
			goto tr543
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr540
			}
		case data[p] >= 49:
			goto tr542
		}
		goto st0
tr541:
//line ragel/datetime.rl:5
 pb = p 
	goto st378
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
//line ragel/parse_datetime.go:13907
		if 49 <= data[p] && data[p] <= 57 {
			goto st379
		}
		goto st0
tr540:
//line ragel/datetime.rl:5
 pb = p 
	goto st379
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
//line ragel/parse_datetime.go:13921
		if data[p] == 32 {
			goto tr545
		}
		goto st0
tr545:
//line ragel/datetime.rl:90

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
//line ragel/parse_datetime.go:13942
		if data[p] == 50 {
			goto tr547
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr548
			}
		case data[p] >= 48:
			goto tr546
		}
		goto st0
tr546:
//line ragel/datetime.rl:5
 pb = p 
	goto st381
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
//line ragel/parse_datetime.go:13964
		switch data[p] {
		case 32:
			goto tr549
		case 43:
			goto tr550
		case 45:
			goto tr551
		case 47:
			goto tr552
		case 58:
			goto tr554
		case 65:
			goto tr555
		case 80:
			goto tr555
		case 90:
			goto tr556
		case 95:
			goto tr552
		case 97:
			goto tr557
		case 112:
			goto tr557
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st405
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr552
			}
		default:
			goto tr552
		}
		goto st0
tr549:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st382
tr593:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st382
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

	goto st382
tr627:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st382
tr636:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st382
tr646:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st382
tr667:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st382
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
//line ragel/parse_datetime.go:14074
		switch data[p] {
		case 32:
			goto st383
		case 43:
			goto st387
		case 45:
			goto st399
		case 47:
			goto tr561
		case 65:
			goto tr563
		case 80:
			goto tr563
		case 90:
			goto tr564
		case 95:
			goto tr561
		case 97:
			goto tr565
		case 112:
			goto tr565
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr562
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr561
			}
		default:
			goto tr561
		}
		goto st0
tr577:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st383
tr592:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st383
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
//line ragel/parse_datetime.go:14128
		if 48 <= data[p] && data[p] <= 57 {
			goto tr562
		}
		goto st0
tr562:
//line ragel/datetime.rl:5
 pb = p 
	goto st384
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
//line ragel/parse_datetime.go:14142
		if 48 <= data[p] && data[p] <= 57 {
			goto st385
		}
		goto st0
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		if 48 <= data[p] && data[p] <= 57 {
			goto st386
		}
		goto st0
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		if 48 <= data[p] && data[p] <= 57 {
			goto st878
		}
		goto st0
	st878:
		if p++; p == pe {
			goto _test_eof878
		}
	st_case_878:
		goto st0
tr550:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st387
tr589:
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

	goto st387
tr594:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st387
tr612:
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

	goto st387
tr604:
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

	goto st387
tr628:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st387
tr637:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st387
tr647:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st387
tr668:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st387
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
//line ragel/parse_datetime.go:14285
		if data[p] == 50 {
			goto tr570
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr571
			}
		case data[p] >= 48:
			goto tr569
		}
		goto st0
tr569:
//line ragel/datetime.rl:5
 pb = p 
	goto st388
tr584:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st388
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
//line ragel/parse_datetime.go:14313
		switch data[p] {
		case 32:
			goto tr572
		case 58:
			goto tr574
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st393
		}
		goto st0
tr572:
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
	goto st389
tr579:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st389
tr582:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st389
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
//line ragel/parse_datetime.go:14378
		switch data[p] {
		case 47:
			goto tr561
		case 95:
			goto tr561
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr562
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr561
			}
		default:
			goto tr561
		}
		goto st0
tr561:
//line ragel/datetime.rl:5
 pb = p 
	goto st390
tr552:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st390
tr596:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st390
tr630:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st390
tr639:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st390
tr650:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st390
tr614:
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
	goto st390
tr671:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st390
tr607:
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
	goto st390
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
//line ragel/parse_datetime.go:14507
		switch data[p] {
		case 47:
			goto st391
		case 95:
			goto st391
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st391
			}
		case data[p] >= 65:
			goto st391
		}
		goto st0
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
		switch data[p] {
		case 47:
			goto st392
		case 95:
			goto st392
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st392
			}
		case data[p] >= 65:
			goto st392
		}
		goto st0
tr591:
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
	goto st392
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
//line ragel/parse_datetime.go:14575
		switch data[p] {
		case 32:
			goto tr577
		case 47:
			goto st392
		case 95:
			goto st392
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st392
			}
		case data[p] >= 65:
			goto st392
		}
		goto st0
tr571:
//line ragel/datetime.rl:5
 pb = p 
	goto st393
tr586:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st393
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
//line ragel/parse_datetime.go:14608
		switch data[p] {
		case 32:
			goto tr572
		case 58:
			goto tr574
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st394
		}
		goto st0
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		if data[p] == 32 {
			goto tr572
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st394
		}
		goto st0
tr574:
//line ragel/datetime.rl:148

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st395
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
//line ragel/parse_datetime.go:14647
		if data[p] == 32 {
			goto tr579
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr581
			}
		case data[p] >= 48:
			goto tr580
		}
		goto st0
tr580:
//line ragel/datetime.rl:5
 pb = p 
	goto st396
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
//line ragel/parse_datetime.go:14669
		if data[p] == 32 {
			goto tr582
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st397
		}
		goto st0
tr581:
//line ragel/datetime.rl:5
 pb = p 
	goto st397
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
//line ragel/parse_datetime.go:14686
		if data[p] == 32 {
			goto tr582
		}
		goto st0
tr570:
//line ragel/datetime.rl:5
 pb = p 
	goto st398
tr585:
//line ragel/datetime.rl:146
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st398
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
//line ragel/parse_datetime.go:14706
		switch data[p] {
		case 32:
			goto tr572
		case 58:
			goto tr574
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st394
			}
		case data[p] >= 48:
			goto st393
		}
		goto st0
tr551:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st399
tr590:
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

	goto st399
tr595:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st399
tr613:
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

	goto st399
tr606:
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

	goto st399
tr629:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st399
tr638:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st399
tr649:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st399
tr670:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st399
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
//line ragel/parse_datetime.go:14836
		if data[p] == 50 {
			goto tr585
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr586
			}
		case data[p] >= 48:
			goto tr584
		}
		goto st0
tr563:
//line ragel/datetime.rl:5
 pb = p 
	goto st400
tr555:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st400
tr599:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st400
tr633:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st400
tr641:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st400
tr652:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st400
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

//line ragel/datetime.rl:5
 pb = p 
	goto st400
tr672:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st400
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
//line ragel/parse_datetime.go:14939
		switch data[p] {
		case 47:
			goto st391
		case 77:
			goto st401
		case 95:
			goto st391
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st391
			}
		case data[p] >= 65:
			goto st391
		}
		goto st0
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		switch data[p] {
		case 32:
			goto tr588
		case 43:
			goto tr589
		case 45:
			goto tr590
		case 47:
			goto tr591
		case 95:
			goto tr591
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr591
			}
		case data[p] >= 65:
			goto tr591
		}
		goto st0
tr588:
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

	goto st402
tr611:
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

	goto st402
tr603:
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

	goto st402
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
//line ragel/parse_datetime.go:15061
		switch data[p] {
		case 32:
			goto st383
		case 43:
			goto st387
		case 45:
			goto st399
		case 47:
			goto tr561
		case 90:
			goto tr564
		case 95:
			goto tr561
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr562
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr561
			}
		default:
			goto tr561
		}
		goto st0
tr564:
//line ragel/datetime.rl:5
 pb = p 
	goto st403
tr556:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st403
tr600:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st403
tr634:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st403
tr642:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st403
tr653:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st403
tr616:
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
	goto st403
tr673:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st403
tr609:
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
	goto st403
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
//line ragel/parse_datetime.go:15198
		switch data[p] {
		case 32:
			goto tr592
		case 47:
			goto st391
		case 95:
			goto st391
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st391
			}
		case data[p] >= 65:
			goto st391
		}
		goto st0
tr565:
//line ragel/datetime.rl:5
 pb = p 
	goto st404
tr557:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st404
tr601:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st404
tr635:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st404
tr643:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st404
tr654:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st404
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
	goto st404
tr674:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st404
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
//line ragel/parse_datetime.go:15306
		switch data[p] {
		case 47:
			goto st391
		case 95:
			goto st391
		case 109:
			goto st401
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st391
			}
		case data[p] >= 65:
			goto st391
		}
		goto st0
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 32:
			goto tr593
		case 43:
			goto tr594
		case 45:
			goto tr595
		case 47:
			goto tr596
		case 58:
			goto tr598
		case 65:
			goto tr599
		case 80:
			goto tr599
		case 90:
			goto tr600
		case 95:
			goto tr596
		case 97:
			goto tr601
		case 112:
			goto tr601
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st406
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr596
			}
		default:
			goto tr596
		}
		goto st0
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		if 48 <= data[p] && data[p] <= 57 {
			goto st407
		}
		goto st0
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 32:
			goto tr603
		case 43:
			goto tr604
		case 45:
			goto tr606
		case 47:
			goto tr607
		case 90:
			goto tr609
		case 95:
			goto tr607
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr605
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr607
				}
			case data[p] >= 65:
				goto tr607
			}
		default:
			goto st418
		}
		goto st0
tr605:
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

	goto st408
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
//line ragel/parse_datetime.go:15434
		if 48 <= data[p] && data[p] <= 57 {
			goto tr610
		}
		goto st0
tr610:
//line ragel/datetime.rl:5
 pb = p 
	goto st409
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
//line ragel/parse_datetime.go:15448
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st410
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st411
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st412
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st413
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st414
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st415
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st416
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st417
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		case data[p] >= 65:
			goto tr614
		}
		goto st0
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		if 48 <= data[p] && data[p] <= 57 {
			goto st419
		}
		goto st0
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		switch data[p] {
		case 32:
			goto tr603
		case 43:
			goto tr604
		case 45:
			goto tr606
		case 47:
			goto tr607
		case 90:
			goto tr609
		case 95:
			goto tr607
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr605
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr607
			}
		default:
			goto tr607
		}
		goto st0
tr554:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st420
tr598:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st420
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
//line ragel/parse_datetime.go:15786
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr626
			}
		case data[p] >= 48:
			goto tr625
		}
		goto st0
tr625:
//line ragel/datetime.rl:5
 pb = p 
	goto st421
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
//line ragel/parse_datetime.go:15805
		switch data[p] {
		case 32:
			goto tr627
		case 43:
			goto tr628
		case 45:
			goto tr629
		case 47:
			goto tr630
		case 58:
			goto tr632
		case 65:
			goto tr633
		case 80:
			goto tr633
		case 90:
			goto tr634
		case 95:
			goto tr630
		case 97:
			goto tr635
		case 112:
			goto tr635
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st422
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr630
			}
		default:
			goto tr630
		}
		goto st0
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		switch data[p] {
		case 32:
			goto tr636
		case 43:
			goto tr637
		case 45:
			goto tr638
		case 47:
			goto tr639
		case 58:
			goto tr640
		case 65:
			goto tr641
		case 80:
			goto tr641
		case 90:
			goto tr642
		case 95:
			goto tr639
		case 97:
			goto tr643
		case 112:
			goto tr643
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr639
			}
		case data[p] >= 66:
			goto tr639
		}
		goto st0
tr632:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st423
tr640:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st423
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
//line ragel/parse_datetime.go:15898
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr645
			}
		case data[p] >= 48:
			goto tr644
		}
		goto st0
tr644:
//line ragel/datetime.rl:5
 pb = p 
	goto st424
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
//line ragel/parse_datetime.go:15917
		switch data[p] {
		case 32:
			goto tr646
		case 43:
			goto tr647
		case 45:
			goto tr649
		case 47:
			goto tr650
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr653
		case 95:
			goto tr650
		case 97:
			goto tr654
		case 112:
			goto tr654
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr648
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr650
				}
			case data[p] >= 66:
				goto tr650
			}
		default:
			goto st435
		}
		goto st0
tr648:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st425
tr669:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st425
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
//line ragel/parse_datetime.go:15975
		if 48 <= data[p] && data[p] <= 57 {
			goto tr655
		}
		goto st0
tr655:
//line ragel/datetime.rl:5
 pb = p 
	goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
//line ragel/parse_datetime.go:15989
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr658
		case 80:
			goto tr658
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr659
		case 112:
			goto tr659
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st427
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr658
		case 80:
			goto tr658
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr659
		case 112:
			goto tr659
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st428
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr658
		case 80:
			goto tr658
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr659
		case 112:
			goto tr659
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st429
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr658
		case 80:
			goto tr658
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr659
		case 112:
			goto tr659
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st430
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr658
		case 80:
			goto tr658
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr659
		case 112:
			goto tr659
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st431
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr658
		case 80:
			goto tr658
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr659
		case 112:
			goto tr659
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st432
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr658
		case 80:
			goto tr658
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr659
		case 112:
			goto tr659
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st433
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr658
		case 80:
			goto tr658
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr659
		case 112:
			goto tr659
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st434
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		switch data[p] {
		case 32:
			goto tr656
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr658
		case 80:
			goto tr658
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr659
		case 112:
			goto tr659
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		case data[p] >= 66:
			goto tr614
		}
		goto st0
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		switch data[p] {
		case 32:
			goto tr667
		case 43:
			goto tr668
		case 45:
			goto tr670
		case 47:
			goto tr671
		case 65:
			goto tr672
		case 80:
			goto tr672
		case 90:
			goto tr673
		case 95:
			goto tr671
		case 97:
			goto tr674
		case 112:
			goto tr674
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr669
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr671
			}
		default:
			goto tr671
		}
		goto st0
tr645:
//line ragel/datetime.rl:5
 pb = p 
	goto st436
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
//line ragel/parse_datetime.go:16390
		switch data[p] {
		case 32:
			goto tr646
		case 43:
			goto tr647
		case 45:
			goto tr649
		case 47:
			goto tr650
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr653
		case 95:
			goto tr650
		case 97:
			goto tr654
		case 112:
			goto tr654
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr648
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr650
			}
		default:
			goto tr650
		}
		goto st0
tr626:
//line ragel/datetime.rl:5
 pb = p 
	goto st437
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
//line ragel/parse_datetime.go:16435
		switch data[p] {
		case 32:
			goto tr627
		case 43:
			goto tr628
		case 45:
			goto tr629
		case 47:
			goto tr630
		case 58:
			goto tr632
		case 65:
			goto tr633
		case 80:
			goto tr633
		case 90:
			goto tr634
		case 95:
			goto tr630
		case 97:
			goto tr635
		case 112:
			goto tr635
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr630
			}
		case data[p] >= 66:
			goto tr630
		}
		goto st0
tr547:
//line ragel/datetime.rl:5
 pb = p 
	goto st438
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
//line ragel/parse_datetime.go:16478
		switch data[p] {
		case 32:
			goto tr549
		case 43:
			goto tr550
		case 45:
			goto tr551
		case 47:
			goto tr552
		case 58:
			goto tr554
		case 65:
			goto tr555
		case 80:
			goto tr555
		case 90:
			goto tr556
		case 95:
			goto tr552
		case 97:
			goto tr557
		case 112:
			goto tr557
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st405
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr552
				}
			case data[p] >= 66:
				goto tr552
			}
		default:
			goto st439
		}
		goto st0
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		if 48 <= data[p] && data[p] <= 57 {
			goto st406
		}
		goto st0
tr548:
//line ragel/datetime.rl:5
 pb = p 
	goto st440
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
//line ragel/parse_datetime.go:16539
		switch data[p] {
		case 32:
			goto tr549
		case 43:
			goto tr550
		case 45:
			goto tr551
		case 47:
			goto tr552
		case 58:
			goto tr554
		case 65:
			goto tr555
		case 80:
			goto tr555
		case 90:
			goto tr556
		case 95:
			goto tr552
		case 97:
			goto tr557
		case 112:
			goto tr557
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st439
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr552
			}
		default:
			goto tr552
		}
		goto st0
tr542:
//line ragel/datetime.rl:5
 pb = p 
	goto st441
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
//line ragel/parse_datetime.go:16586
		if data[p] == 32 {
			goto tr545
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st379
		}
		goto st0
tr543:
//line ragel/datetime.rl:5
 pb = p 
	goto st442
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
//line ragel/parse_datetime.go:16603
		if data[p] == 32 {
			goto tr545
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st379
		}
		goto st0
tr537:
//line ragel/datetime.rl:5
 pb = p 
	goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
//line ragel/parse_datetime.go:16620
		if 49 <= data[p] && data[p] <= 57 {
			goto st444
		}
		goto st0
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		if data[p] == 32 {
			goto tr677
		}
		goto st0
tr677:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st445
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
//line ragel/parse_datetime.go:16650
		if data[p] == 50 {
			goto tr679
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr680
			}
		case data[p] >= 48:
			goto tr678
		}
		goto st0
tr678:
//line ragel/datetime.rl:5
 pb = p 
	goto st446
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
//line ragel/parse_datetime.go:16672
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr550
		case 45:
			goto tr551
		case 47:
			goto tr552
		case 58:
			goto tr683
		case 65:
			goto tr684
		case 80:
			goto tr684
		case 90:
			goto tr556
		case 95:
			goto tr552
		case 97:
			goto tr685
		case 112:
			goto tr685
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st452
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr552
			}
		default:
			goto tr552
		}
		goto st0
tr681:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st447
tr690:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st447
tr729:
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

	goto st447
tr712:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st447
tr717:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st447
tr723:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st447
tr740:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st447
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
//line ragel/parse_datetime.go:16782
		switch data[p] {
		case 32:
			goto st383
		case 43:
			goto st387
		case 45:
			goto st399
		case 47:
			goto tr561
		case 65:
			goto tr686
		case 80:
			goto tr686
		case 90:
			goto tr564
		case 95:
			goto tr561
		case 97:
			goto tr687
		case 112:
			goto tr687
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr383
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr561
			}
		default:
			goto tr561
		}
		goto st0
tr686:
//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr684:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr693:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr715:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr719:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr726:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr731:
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
	goto st448
tr742:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
//line ragel/parse_datetime.go:16908
		switch data[p] {
		case 47:
			goto st391
		case 77:
			goto st449
		case 95:
			goto st391
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st391
			}
		case data[p] >= 65:
			goto st391
		}
		goto st0
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		switch data[p] {
		case 32:
			goto tr689
		case 43:
			goto tr589
		case 45:
			goto tr590
		case 47:
			goto tr591
		case 95:
			goto tr591
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr591
			}
		case data[p] >= 65:
			goto tr591
		}
		goto st0
tr689:
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

	goto st450
tr700:
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

	goto st450
tr696:
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

	goto st450
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
//line ragel/parse_datetime.go:17030
		switch data[p] {
		case 32:
			goto st383
		case 43:
			goto st387
		case 45:
			goto st399
		case 47:
			goto tr561
		case 90:
			goto tr564
		case 95:
			goto tr561
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr383
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr561
			}
		default:
			goto tr561
		}
		goto st0
tr687:
//line ragel/datetime.rl:5
 pb = p 
	goto st451
tr685:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st451
tr694:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st451
tr716:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st451
tr720:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st451
tr727:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st451
tr732:
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
	goto st451
tr743:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st451
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
//line ragel/parse_datetime.go:17148
		switch data[p] {
		case 47:
			goto st391
		case 95:
			goto st391
		case 109:
			goto st449
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st391
			}
		case data[p] >= 65:
			goto st391
		}
		goto st0
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		switch data[p] {
		case 32:
			goto tr690
		case 43:
			goto tr594
		case 45:
			goto tr595
		case 47:
			goto tr596
		case 58:
			goto tr692
		case 65:
			goto tr693
		case 80:
			goto tr693
		case 90:
			goto tr600
		case 95:
			goto tr596
		case 97:
			goto tr694
		case 112:
			goto tr694
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st453
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr596
			}
		default:
			goto tr596
		}
		goto st0
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		if 48 <= data[p] && data[p] <= 57 {
			goto st454
		}
		goto st0
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		switch data[p] {
		case 32:
			goto tr696
		case 43:
			goto tr604
		case 45:
			goto tr606
		case 47:
			goto tr607
		case 90:
			goto tr609
		case 95:
			goto tr607
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr697
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr607
				}
			case data[p] >= 65:
				goto tr607
			}
		default:
			goto st465
		}
		goto st0
tr697:
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

	goto st455
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
//line ragel/parse_datetime.go:17276
		if 48 <= data[p] && data[p] <= 57 {
			goto tr699
		}
		goto st0
tr699:
//line ragel/datetime.rl:5
 pb = p 
	goto st456
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
//line ragel/parse_datetime.go:17290
		switch data[p] {
		case 32:
			goto tr700
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st457
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 32:
			goto tr700
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st458
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
		switch data[p] {
		case 32:
			goto tr700
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st459
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
		switch data[p] {
		case 32:
			goto tr700
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st460
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
		switch data[p] {
		case 32:
			goto tr700
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st461
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		switch data[p] {
		case 32:
			goto tr700
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st462
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
		switch data[p] {
		case 32:
			goto tr700
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st463
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
		switch data[p] {
		case 32:
			goto tr700
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st464
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
		switch data[p] {
		case 32:
			goto tr700
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 90:
			goto tr616
		case 95:
			goto tr614
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		case data[p] >= 65:
			goto tr614
		}
		goto st0
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
		if 48 <= data[p] && data[p] <= 57 {
			goto st466
		}
		goto st0
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		switch data[p] {
		case 32:
			goto tr696
		case 43:
			goto tr604
		case 45:
			goto tr606
		case 47:
			goto tr607
		case 90:
			goto tr609
		case 95:
			goto tr607
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr697
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr607
			}
		default:
			goto tr607
		}
		goto st0
tr683:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st467
tr692:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st467
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
//line ragel/parse_datetime.go:17628
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr711
			}
		case data[p] >= 48:
			goto tr710
		}
		goto st0
tr710:
//line ragel/datetime.rl:5
 pb = p 
	goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
//line ragel/parse_datetime.go:17647
		switch data[p] {
		case 32:
			goto tr712
		case 43:
			goto tr628
		case 45:
			goto tr629
		case 47:
			goto tr630
		case 58:
			goto tr714
		case 65:
			goto tr715
		case 80:
			goto tr715
		case 90:
			goto tr634
		case 95:
			goto tr630
		case 97:
			goto tr716
		case 112:
			goto tr716
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st469
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr630
			}
		default:
			goto tr630
		}
		goto st0
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		switch data[p] {
		case 32:
			goto tr717
		case 43:
			goto tr637
		case 45:
			goto tr638
		case 47:
			goto tr639
		case 58:
			goto tr718
		case 65:
			goto tr719
		case 80:
			goto tr719
		case 90:
			goto tr642
		case 95:
			goto tr639
		case 97:
			goto tr720
		case 112:
			goto tr720
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr639
			}
		case data[p] >= 66:
			goto tr639
		}
		goto st0
tr714:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st470
tr718:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st470
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
//line ragel/parse_datetime.go:17740
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr722
			}
		case data[p] >= 48:
			goto tr721
		}
		goto st0
tr721:
//line ragel/datetime.rl:5
 pb = p 
	goto st471
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
//line ragel/parse_datetime.go:17759
		switch data[p] {
		case 32:
			goto tr723
		case 43:
			goto tr647
		case 45:
			goto tr649
		case 47:
			goto tr650
		case 65:
			goto tr726
		case 80:
			goto tr726
		case 90:
			goto tr653
		case 95:
			goto tr650
		case 97:
			goto tr727
		case 112:
			goto tr727
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr724
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr650
				}
			case data[p] >= 66:
				goto tr650
			}
		default:
			goto st482
		}
		goto st0
tr724:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st472
tr741:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st472
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
//line ragel/parse_datetime.go:17817
		if 48 <= data[p] && data[p] <= 57 {
			goto tr728
		}
		goto st0
tr728:
//line ragel/datetime.rl:5
 pb = p 
	goto st473
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
//line ragel/parse_datetime.go:17831
		switch data[p] {
		case 32:
			goto tr729
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st474
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		switch data[p] {
		case 32:
			goto tr729
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st475
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		switch data[p] {
		case 32:
			goto tr729
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st476
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
		switch data[p] {
		case 32:
			goto tr729
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st477
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		switch data[p] {
		case 32:
			goto tr729
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st478
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		switch data[p] {
		case 32:
			goto tr729
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st479
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		switch data[p] {
		case 32:
			goto tr729
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st480
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		switch data[p] {
		case 32:
			goto tr729
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st481
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		default:
			goto tr614
		}
		goto st0
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		switch data[p] {
		case 32:
			goto tr729
		case 43:
			goto tr612
		case 45:
			goto tr613
		case 47:
			goto tr614
		case 65:
			goto tr731
		case 80:
			goto tr731
		case 90:
			goto tr616
		case 95:
			goto tr614
		case 97:
			goto tr732
		case 112:
			goto tr732
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr614
			}
		case data[p] >= 66:
			goto tr614
		}
		goto st0
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
		switch data[p] {
		case 32:
			goto tr740
		case 43:
			goto tr668
		case 45:
			goto tr670
		case 47:
			goto tr671
		case 65:
			goto tr742
		case 80:
			goto tr742
		case 90:
			goto tr673
		case 95:
			goto tr671
		case 97:
			goto tr743
		case 112:
			goto tr743
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr741
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr671
			}
		default:
			goto tr671
		}
		goto st0
tr722:
//line ragel/datetime.rl:5
 pb = p 
	goto st483
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
//line ragel/parse_datetime.go:18232
		switch data[p] {
		case 32:
			goto tr723
		case 43:
			goto tr647
		case 45:
			goto tr649
		case 47:
			goto tr650
		case 65:
			goto tr726
		case 80:
			goto tr726
		case 90:
			goto tr653
		case 95:
			goto tr650
		case 97:
			goto tr727
		case 112:
			goto tr727
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr724
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr650
			}
		default:
			goto tr650
		}
		goto st0
tr711:
//line ragel/datetime.rl:5
 pb = p 
	goto st484
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
//line ragel/parse_datetime.go:18277
		switch data[p] {
		case 32:
			goto tr712
		case 43:
			goto tr628
		case 45:
			goto tr629
		case 47:
			goto tr630
		case 58:
			goto tr714
		case 65:
			goto tr715
		case 80:
			goto tr715
		case 90:
			goto tr634
		case 95:
			goto tr630
		case 97:
			goto tr716
		case 112:
			goto tr716
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr630
			}
		case data[p] >= 66:
			goto tr630
		}
		goto st0
tr679:
//line ragel/datetime.rl:5
 pb = p 
	goto st485
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
//line ragel/parse_datetime.go:18320
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr550
		case 45:
			goto tr551
		case 47:
			goto tr552
		case 58:
			goto tr683
		case 65:
			goto tr684
		case 80:
			goto tr684
		case 90:
			goto tr556
		case 95:
			goto tr552
		case 97:
			goto tr685
		case 112:
			goto tr685
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st452
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr552
				}
			case data[p] >= 66:
				goto tr552
			}
		default:
			goto st486
		}
		goto st0
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		if 48 <= data[p] && data[p] <= 57 {
			goto st453
		}
		goto st0
tr680:
//line ragel/datetime.rl:5
 pb = p 
	goto st487
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
//line ragel/parse_datetime.go:18381
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr550
		case 45:
			goto tr551
		case 47:
			goto tr552
		case 58:
			goto tr683
		case 65:
			goto tr684
		case 80:
			goto tr684
		case 90:
			goto tr556
		case 95:
			goto tr552
		case 97:
			goto tr685
		case 112:
			goto tr685
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st486
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr552
			}
		default:
			goto tr552
		}
		goto st0
tr538:
//line ragel/datetime.rl:5
 pb = p 
	goto st488
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
//line ragel/parse_datetime.go:18428
		if data[p] == 32 {
			goto tr545
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st444
		}
		goto st0
tr539:
//line ragel/datetime.rl:5
 pb = p 
	goto st489
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
//line ragel/parse_datetime.go:18445
		if data[p] == 32 {
			goto tr545
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st444
		}
		goto st0
tr534:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st490
tr749:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st490
tr756:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st490
tr765:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st490
tr775:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st490
tr783:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st490
tr786:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st490
tr792:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st490
tr796:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st490
tr800:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st490
tr809:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st490
tr817:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st490
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
//line ragel/parse_datetime.go:18506
		if data[p] == 32 {
			goto st376
		}
		goto st0
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		if data[p] == 108 {
			goto st492
		}
		goto st0
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
		switch data[p] {
		case 32:
			goto tr533
		case 46:
			goto tr534
		}
		goto st0
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
		if data[p] == 103 {
			goto st494
		}
		goto st0
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		switch data[p] {
		case 32:
			goto tr748
		case 46:
			goto tr749
		case 117:
			goto st495
		}
		goto st0
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		if data[p] == 115 {
			goto st496
		}
		goto st0
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
		if data[p] == 116 {
			goto st497
		}
		goto st0
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		switch data[p] {
		case 32:
			goto tr748
		case 46:
			goto tr749
		}
		goto st0
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		if data[p] == 101 {
			goto st499
		}
		goto st0
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		if data[p] == 99 {
			goto st500
		}
		goto st0
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		switch data[p] {
		case 32:
			goto tr755
		case 46:
			goto tr756
		case 101:
			goto st501
		}
		goto st0
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		if data[p] == 109 {
			goto st502
		}
		goto st0
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		if data[p] == 98 {
			goto st503
		}
		goto st0
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		if data[p] == 101 {
			goto st504
		}
		goto st0
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		if data[p] == 114 {
			goto st505
		}
		goto st0
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		switch data[p] {
		case 32:
			goto tr755
		case 46:
			goto tr756
		}
		goto st0
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		if data[p] == 101 {
			goto st507
		}
		goto st0
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		if data[p] == 98 {
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
			goto tr764
		case 46:
			goto tr765
		case 114:
			goto st509
		}
		goto st0
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		if data[p] == 117 {
			goto st510
		}
		goto st0
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
		if data[p] == 97 {
			goto st511
		}
		goto st0
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		if data[p] == 114 {
			goto st512
		}
		goto st0
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		if data[p] == 121 {
			goto st513
		}
		goto st0
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
		switch data[p] {
		case 32:
			goto tr764
		case 46:
			goto tr765
		}
		goto st0
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		switch data[p] {
		case 97:
			goto st515
		case 117:
			goto st521
		}
		goto st0
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
		if data[p] == 110 {
			goto st516
		}
		goto st0
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		switch data[p] {
		case 32:
			goto tr774
		case 46:
			goto tr775
		case 117:
			goto st517
		}
		goto st0
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		if data[p] == 97 {
			goto st518
		}
		goto st0
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		if data[p] == 114 {
			goto st519
		}
		goto st0
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		if data[p] == 121 {
			goto st520
		}
		goto st0
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
		switch data[p] {
		case 32:
			goto tr774
		case 46:
			goto tr775
		}
		goto st0
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		switch data[p] {
		case 108:
			goto st522
		case 110:
			goto st524
		}
		goto st0
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		switch data[p] {
		case 32:
			goto tr782
		case 46:
			goto tr783
		case 121:
			goto st523
		}
		goto st0
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
		switch data[p] {
		case 32:
			goto tr782
		case 46:
			goto tr783
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
		case 46:
			goto tr786
		case 101:
			goto st525
		}
		goto st0
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		switch data[p] {
		case 32:
			goto tr785
		case 46:
			goto tr786
		}
		goto st0
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
		if data[p] == 97 {
			goto st527
		}
		goto st0
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
		switch data[p] {
		case 114:
			goto st528
		case 121:
			goto st531
		}
		goto st0
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
		switch data[p] {
		case 32:
			goto tr791
		case 46:
			goto tr792
		case 99:
			goto st529
		}
		goto st0
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
		if data[p] == 104 {
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
			goto tr791
		case 46:
			goto tr792
		}
		goto st0
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		switch data[p] {
		case 32:
			goto tr795
		case 46:
			goto tr796
		}
		goto st0
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
		if data[p] == 111 {
			goto st533
		}
		goto st0
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		if data[p] == 118 {
			goto st534
		}
		goto st0
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
		switch data[p] {
		case 32:
			goto tr799
		case 46:
			goto tr800
		case 101:
			goto st535
		}
		goto st0
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		if data[p] == 109 {
			goto st536
		}
		goto st0
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		if data[p] == 98 {
			goto st537
		}
		goto st0
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		if data[p] == 101 {
			goto st538
		}
		goto st0
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
		if data[p] == 114 {
			goto st539
		}
		goto st0
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
		switch data[p] {
		case 32:
			goto tr799
		case 46:
			goto tr800
		}
		goto st0
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
		if data[p] == 99 {
			goto st541
		}
		goto st0
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
		if data[p] == 116 {
			goto st542
		}
		goto st0
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
		switch data[p] {
		case 32:
			goto tr808
		case 46:
			goto tr809
		case 111:
			goto st543
		}
		goto st0
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
		if data[p] == 98 {
			goto st544
		}
		goto st0
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
		if data[p] == 101 {
			goto st545
		}
		goto st0
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
		if data[p] == 114 {
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
			goto tr808
		case 46:
			goto tr809
		}
		goto st0
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
		if data[p] == 101 {
			goto st548
		}
		goto st0
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
		if data[p] == 112 {
			goto st549
		}
		goto st0
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
		switch data[p] {
		case 32:
			goto tr816
		case 46:
			goto tr817
		case 116:
			goto st550
		}
		goto st0
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
		switch data[p] {
		case 32:
			goto tr816
		case 46:
			goto tr817
		case 101:
			goto st551
		}
		goto st0
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
		if data[p] == 109 {
			goto st552
		}
		goto st0
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
		if data[p] == 98 {
			goto st553
		}
		goto st0
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
		if data[p] == 101 {
			goto st554
		}
		goto st0
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
		if data[p] == 114 {
			goto st555
		}
		goto st0
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
		switch data[p] {
		case 32:
			goto tr816
		case 46:
			goto tr817
		}
		goto st0
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
		if data[p] == 32 {
			goto st557
		}
		goto st0
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
		switch data[p] {
		case 48:
			goto tr825
		case 51:
			goto tr827
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr828
			}
		case data[p] >= 49:
			goto tr826
		}
		goto st0
tr825:
//line ragel/datetime.rl:5
 pb = p 
	goto st558
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
//line ragel/parse_datetime.go:19234
		if 49 <= data[p] && data[p] <= 57 {
			goto st559
		}
		goto st0
tr828:
//line ragel/datetime.rl:5
 pb = p 
	goto st559
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
//line ragel/parse_datetime.go:19248
		switch data[p] {
		case 32:
			goto tr830
		case 45:
			goto tr831
		}
		goto st0
tr830:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st560
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
//line ragel/parse_datetime.go:19272
		switch data[p] {
		case 65:
			goto st561
		case 68:
			goto st572
		case 70:
			goto st580
		case 74:
			goto st588
		case 77:
			goto st600
		case 78:
			goto st606
		case 79:
			goto st614
		case 83:
			goto st621
		case 97:
			goto st561
		case 100:
			goto st572
		case 102:
			goto st580
		case 106:
			goto st588
		case 109:
			goto st600
		case 110:
			goto st606
		case 111:
			goto st614
		case 115:
			goto st621
		}
		goto st0
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
		switch data[p] {
		case 112:
			goto st562
		case 117:
			goto st567
		}
		goto st0
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
		if data[p] == 114 {
			goto st563
		}
		goto st0
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
		switch data[p] {
		case 32:
			goto tr180
		case 46:
			goto tr843
		case 105:
			goto st565
		}
		goto st0
tr843:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st564
tr847:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st564
tr853:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st564
tr861:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st564
tr870:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st564
tr877:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st564
tr879:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st564
tr884:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st564
tr887:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st564
tr890:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st564
tr898:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st564
tr905:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st564
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
//line ragel/parse_datetime.go:19396
		if data[p] == 32 {
			goto st118
		}
		goto st0
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
		if data[p] == 108 {
			goto st566
		}
		goto st0
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
		switch data[p] {
		case 32:
			goto tr180
		case 46:
			goto tr843
		}
		goto st0
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		if data[p] == 103 {
			goto st568
		}
		goto st0
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
		switch data[p] {
		case 32:
			goto tr194
		case 46:
			goto tr847
		case 117:
			goto st569
		}
		goto st0
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		if data[p] == 115 {
			goto st570
		}
		goto st0
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		if data[p] == 116 {
			goto st571
		}
		goto st0
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
		switch data[p] {
		case 32:
			goto tr194
		case 46:
			goto tr847
		}
		goto st0
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
		if data[p] == 101 {
			goto st573
		}
		goto st0
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		if data[p] == 99 {
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
			goto tr202
		case 46:
			goto tr853
		case 101:
			goto st575
		}
		goto st0
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
		if data[p] == 109 {
			goto st576
		}
		goto st0
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
		if data[p] == 98 {
			goto st577
		}
		goto st0
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
		if data[p] == 101 {
			goto st578
		}
		goto st0
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
		if data[p] == 114 {
			goto st579
		}
		goto st0
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
		switch data[p] {
		case 32:
			goto tr202
		case 46:
			goto tr853
		}
		goto st0
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		if data[p] == 101 {
			goto st581
		}
		goto st0
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
		if data[p] == 98 {
			goto st582
		}
		goto st0
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		switch data[p] {
		case 32:
			goto tr212
		case 46:
			goto tr861
		case 114:
			goto st583
		}
		goto st0
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		if data[p] == 117 {
			goto st584
		}
		goto st0
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		if data[p] == 97 {
			goto st585
		}
		goto st0
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		if data[p] == 114 {
			goto st586
		}
		goto st0
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		if data[p] == 121 {
			goto st587
		}
		goto st0
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
		switch data[p] {
		case 32:
			goto tr212
		case 46:
			goto tr861
		}
		goto st0
	st588:
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
		switch data[p] {
		case 97:
			goto st589
		case 117:
			goto st595
		}
		goto st0
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		if data[p] == 110 {
			goto st590
		}
		goto st0
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
		switch data[p] {
		case 32:
			goto tr223
		case 46:
			goto tr870
		case 117:
			goto st591
		}
		goto st0
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		if data[p] == 97 {
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
		if data[p] == 121 {
			goto st594
		}
		goto st0
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		switch data[p] {
		case 32:
			goto tr223
		case 46:
			goto tr870
		}
		goto st0
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		switch data[p] {
		case 108:
			goto st596
		case 110:
			goto st598
		}
		goto st0
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		switch data[p] {
		case 32:
			goto tr232
		case 46:
			goto tr877
		case 121:
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
			goto tr232
		case 46:
			goto tr877
		}
		goto st0
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		switch data[p] {
		case 32:
			goto tr236
		case 46:
			goto tr879
		case 101:
			goto st599
		}
		goto st0
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		switch data[p] {
		case 32:
			goto tr236
		case 46:
			goto tr879
		}
		goto st0
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		if data[p] == 97 {
			goto st601
		}
		goto st0
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
		switch data[p] {
		case 114:
			goto st602
		case 121:
			goto st605
		}
		goto st0
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
		switch data[p] {
		case 32:
			goto tr243
		case 46:
			goto tr884
		case 99:
			goto st603
		}
		goto st0
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
		if data[p] == 104 {
			goto st604
		}
		goto st0
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		switch data[p] {
		case 32:
			goto tr243
		case 46:
			goto tr884
		}
		goto st0
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
		switch data[p] {
		case 32:
			goto tr248
		case 46:
			goto tr887
		}
		goto st0
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
		if data[p] == 111 {
			goto st607
		}
		goto st0
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
		if data[p] == 118 {
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
			goto tr253
		case 46:
			goto tr890
		case 101:
			goto st609
		}
		goto st0
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
		if data[p] == 109 {
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
		if data[p] == 101 {
			goto st612
		}
		goto st0
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
		if data[p] == 114 {
			goto st613
		}
		goto st0
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		switch data[p] {
		case 32:
			goto tr253
		case 46:
			goto tr890
		}
		goto st0
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
		if data[p] == 99 {
			goto st615
		}
		goto st0
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
		if data[p] == 116 {
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
			goto tr263
		case 46:
			goto tr898
		case 111:
			goto st617
		}
		goto st0
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
		if data[p] == 98 {
			goto st618
		}
		goto st0
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
		if data[p] == 101 {
			goto st619
		}
		goto st0
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		if data[p] == 114 {
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
			goto tr263
		case 46:
			goto tr898
		}
		goto st0
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		if data[p] == 101 {
			goto st622
		}
		goto st0
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		if data[p] == 112 {
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
			goto tr272
		case 46:
			goto tr905
		case 116:
			goto st624
		}
		goto st0
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
		switch data[p] {
		case 32:
			goto tr272
		case 46:
			goto tr905
		case 101:
			goto st625
		}
		goto st0
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
		if data[p] == 109 {
			goto st626
		}
		goto st0
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
		if data[p] == 98 {
			goto st627
		}
		goto st0
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
		if data[p] == 101 {
			goto st628
		}
		goto st0
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
		if data[p] == 114 {
			goto st629
		}
		goto st0
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
		switch data[p] {
		case 32:
			goto tr272
		case 46:
			goto tr905
		}
		goto st0
tr831:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st630
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
//line ragel/parse_datetime.go:20102
		switch data[p] {
		case 65:
			goto st631
		case 68:
			goto st644
		case 70:
			goto st652
		case 74:
			goto st660
		case 77:
			goto st672
		case 78:
			goto st678
		case 79:
			goto st686
		case 83:
			goto st693
		case 97:
			goto st631
		case 100:
			goto st644
		case 102:
			goto st652
		case 106:
			goto st660
		case 109:
			goto st672
		case 110:
			goto st678
		case 111:
			goto st686
		case 115:
			goto st693
		}
		goto st0
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		switch data[p] {
		case 112:
			goto st632
		case 117:
			goto st639
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
		case 45:
			goto tr923
		case 46:
			goto tr924
		case 105:
			goto st637
		}
		goto st0
tr923:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st634
tr931:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st634
tr938:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st634
tr947:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st634
tr957:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st634
tr965:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st634
tr968:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st634
tr974:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st634
tr978:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st634
tr982:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st634
tr991:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st634
tr999:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st634
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
//line ragel/parse_datetime.go:20226
		if 48 <= data[p] && data[p] <= 57 {
			goto tr926
		}
		goto st0
tr926:
//line ragel/datetime.rl:5
 pb = p 
	goto st635
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
//line ragel/parse_datetime.go:20240
		if 48 <= data[p] && data[p] <= 57 {
			goto st879
		}
		goto st0
	st879:
		if p++; p == pe {
			goto _test_eof879
		}
	st_case_879:
		switch data[p] {
		case 32:
			goto tr1212
		case 43:
			goto tr1213
		case 45:
			goto tr1214
		case 47:
			goto tr1215
		case 84:
			goto tr1216
		case 90:
			goto tr1217
		case 95:
			goto tr1218
		case 116:
			goto tr1218
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1215
			}
		case data[p] >= 65:
			goto tr1215
		}
		goto st0
tr924:
//line ragel/datetime.rl:80
 st.Month = 4 
	goto st636
tr932:
//line ragel/datetime.rl:84
 st.Month = 8 
	goto st636
tr939:
//line ragel/datetime.rl:88
 st.Month = 12 
	goto st636
tr948:
//line ragel/datetime.rl:78
 st.Month = 2 
	goto st636
tr958:
//line ragel/datetime.rl:77
 st.Month = 1 
	goto st636
tr966:
//line ragel/datetime.rl:83
 st.Month = 7 
	goto st636
tr969:
//line ragel/datetime.rl:82
 st.Month = 6 
	goto st636
tr975:
//line ragel/datetime.rl:79
 st.Month = 3 
	goto st636
tr979:
//line ragel/datetime.rl:81
 st.Month = 5 
	goto st636
tr983:
//line ragel/datetime.rl:87
 st.Month = 11 
	goto st636
tr992:
//line ragel/datetime.rl:86
 st.Month = 10 
	goto st636
tr1000:
//line ragel/datetime.rl:85
 st.Month = 9 
	goto st636
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
//line ragel/parse_datetime.go:20330
		if data[p] == 45 {
			goto st634
		}
		goto st0
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
		if data[p] == 108 {
			goto st638
		}
		goto st0
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
		switch data[p] {
		case 45:
			goto tr923
		case 46:
			goto tr924
		}
		goto st0
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
		if data[p] == 103 {
			goto st640
		}
		goto st0
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
		switch data[p] {
		case 45:
			goto tr931
		case 46:
			goto tr932
		case 117:
			goto st641
		}
		goto st0
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
		if data[p] == 115 {
			goto st642
		}
		goto st0
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
		if data[p] == 116 {
			goto st643
		}
		goto st0
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
		switch data[p] {
		case 45:
			goto tr931
		case 46:
			goto tr932
		}
		goto st0
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
		if data[p] == 101 {
			goto st645
		}
		goto st0
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
		if data[p] == 99 {
			goto st646
		}
		goto st0
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
		switch data[p] {
		case 45:
			goto tr938
		case 46:
			goto tr939
		case 101:
			goto st647
		}
		goto st0
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
		if data[p] == 109 {
			goto st648
		}
		goto st0
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
		if data[p] == 98 {
			goto st649
		}
		goto st0
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
		if data[p] == 101 {
			goto st650
		}
		goto st0
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
		if data[p] == 114 {
			goto st651
		}
		goto st0
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		switch data[p] {
		case 45:
			goto tr938
		case 46:
			goto tr939
		}
		goto st0
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
		if data[p] == 101 {
			goto st653
		}
		goto st0
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		if data[p] == 98 {
			goto st654
		}
		goto st0
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
		switch data[p] {
		case 45:
			goto tr947
		case 46:
			goto tr948
		case 114:
			goto st655
		}
		goto st0
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
		if data[p] == 117 {
			goto st656
		}
		goto st0
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
		if data[p] == 97 {
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
		if data[p] == 121 {
			goto st659
		}
		goto st0
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
		switch data[p] {
		case 45:
			goto tr947
		case 46:
			goto tr948
		}
		goto st0
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
		switch data[p] {
		case 97:
			goto st661
		case 117:
			goto st667
		}
		goto st0
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
	st_case_661:
		if data[p] == 110 {
			goto st662
		}
		goto st0
	st662:
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
		switch data[p] {
		case 45:
			goto tr957
		case 46:
			goto tr958
		case 117:
			goto st663
		}
		goto st0
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
		if data[p] == 97 {
			goto st664
		}
		goto st0
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
		if data[p] == 114 {
			goto st665
		}
		goto st0
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
		if data[p] == 121 {
			goto st666
		}
		goto st0
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
		switch data[p] {
		case 45:
			goto tr957
		case 46:
			goto tr958
		}
		goto st0
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
	st_case_667:
		switch data[p] {
		case 108:
			goto st668
		case 110:
			goto st670
		}
		goto st0
	st668:
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
		switch data[p] {
		case 45:
			goto tr965
		case 46:
			goto tr966
		case 121:
			goto st669
		}
		goto st0
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
		switch data[p] {
		case 45:
			goto tr965
		case 46:
			goto tr966
		}
		goto st0
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
		switch data[p] {
		case 45:
			goto tr968
		case 46:
			goto tr969
		case 101:
			goto st671
		}
		goto st0
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
		switch data[p] {
		case 45:
			goto tr968
		case 46:
			goto tr969
		}
		goto st0
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
	st_case_672:
		if data[p] == 97 {
			goto st673
		}
		goto st0
	st673:
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
		switch data[p] {
		case 114:
			goto st674
		case 121:
			goto st677
		}
		goto st0
	st674:
		if p++; p == pe {
			goto _test_eof674
		}
	st_case_674:
		switch data[p] {
		case 45:
			goto tr974
		case 46:
			goto tr975
		case 99:
			goto st675
		}
		goto st0
	st675:
		if p++; p == pe {
			goto _test_eof675
		}
	st_case_675:
		if data[p] == 104 {
			goto st676
		}
		goto st0
	st676:
		if p++; p == pe {
			goto _test_eof676
		}
	st_case_676:
		switch data[p] {
		case 45:
			goto tr974
		case 46:
			goto tr975
		}
		goto st0
	st677:
		if p++; p == pe {
			goto _test_eof677
		}
	st_case_677:
		switch data[p] {
		case 45:
			goto tr978
		case 46:
			goto tr979
		}
		goto st0
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
		if data[p] == 111 {
			goto st679
		}
		goto st0
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
		if data[p] == 118 {
			goto st680
		}
		goto st0
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
		switch data[p] {
		case 45:
			goto tr982
		case 46:
			goto tr983
		case 101:
			goto st681
		}
		goto st0
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
	st_case_681:
		if data[p] == 109 {
			goto st682
		}
		goto st0
	st682:
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
		if data[p] == 98 {
			goto st683
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
		if data[p] == 114 {
			goto st685
		}
		goto st0
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
	st_case_685:
		switch data[p] {
		case 45:
			goto tr982
		case 46:
			goto tr983
		}
		goto st0
	st686:
		if p++; p == pe {
			goto _test_eof686
		}
	st_case_686:
		if data[p] == 99 {
			goto st687
		}
		goto st0
	st687:
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
		if data[p] == 116 {
			goto st688
		}
		goto st0
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
		switch data[p] {
		case 45:
			goto tr991
		case 46:
			goto tr992
		case 111:
			goto st689
		}
		goto st0
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
		if data[p] == 98 {
			goto st690
		}
		goto st0
	st690:
		if p++; p == pe {
			goto _test_eof690
		}
	st_case_690:
		if data[p] == 101 {
			goto st691
		}
		goto st0
	st691:
		if p++; p == pe {
			goto _test_eof691
		}
	st_case_691:
		if data[p] == 114 {
			goto st692
		}
		goto st0
	st692:
		if p++; p == pe {
			goto _test_eof692
		}
	st_case_692:
		switch data[p] {
		case 45:
			goto tr991
		case 46:
			goto tr992
		}
		goto st0
	st693:
		if p++; p == pe {
			goto _test_eof693
		}
	st_case_693:
		if data[p] == 101 {
			goto st694
		}
		goto st0
	st694:
		if p++; p == pe {
			goto _test_eof694
		}
	st_case_694:
		if data[p] == 112 {
			goto st695
		}
		goto st0
	st695:
		if p++; p == pe {
			goto _test_eof695
		}
	st_case_695:
		switch data[p] {
		case 45:
			goto tr999
		case 46:
			goto tr1000
		case 116:
			goto st696
		}
		goto st0
	st696:
		if p++; p == pe {
			goto _test_eof696
		}
	st_case_696:
		switch data[p] {
		case 45:
			goto tr999
		case 46:
			goto tr1000
		case 101:
			goto st697
		}
		goto st0
	st697:
		if p++; p == pe {
			goto _test_eof697
		}
	st_case_697:
		if data[p] == 109 {
			goto st698
		}
		goto st0
	st698:
		if p++; p == pe {
			goto _test_eof698
		}
	st_case_698:
		if data[p] == 98 {
			goto st699
		}
		goto st0
	st699:
		if p++; p == pe {
			goto _test_eof699
		}
	st_case_699:
		if data[p] == 101 {
			goto st700
		}
		goto st0
	st700:
		if p++; p == pe {
			goto _test_eof700
		}
	st_case_700:
		if data[p] == 114 {
			goto st701
		}
		goto st0
	st701:
		if p++; p == pe {
			goto _test_eof701
		}
	st_case_701:
		switch data[p] {
		case 45:
			goto tr999
		case 46:
			goto tr1000
		}
		goto st0
tr826:
//line ragel/datetime.rl:5
 pb = p 
	goto st702
	st702:
		if p++; p == pe {
			goto _test_eof702
		}
	st_case_702:
//line ragel/parse_datetime.go:21029
		switch data[p] {
		case 32:
			goto tr830
		case 45:
			goto tr831
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st559
		}
		goto st0
tr827:
//line ragel/datetime.rl:5
 pb = p 
	goto st703
	st703:
		if p++; p == pe {
			goto _test_eof703
		}
	st_case_703:
//line ragel/parse_datetime.go:21049
		switch data[p] {
		case 32:
			goto tr830
		case 45:
			goto tr831
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st559
		}
		goto st0
	st704:
		if p++; p == pe {
			goto _test_eof704
		}
	st_case_704:
		if data[p] == 97 {
			goto st705
		}
		goto st0
	st705:
		if p++; p == pe {
			goto _test_eof705
		}
	st_case_705:
		if data[p] == 121 {
			goto st706
		}
		goto st0
	st706:
		if p++; p == pe {
			goto _test_eof706
		}
	st_case_706:
		switch data[p] {
		case 32:
			goto st372
		case 44:
			goto st556
		}
		goto st0
	st707:
		if p++; p == pe {
			goto _test_eof707
		}
	st_case_707:
		switch data[p] {
		case 97:
			goto st708
		case 117:
			goto st714
		}
		goto st0
	st708:
		if p++; p == pe {
			goto _test_eof708
		}
	st_case_708:
		if data[p] == 110 {
			goto st709
		}
		goto st0
	st709:
		if p++; p == pe {
			goto _test_eof709
		}
	st_case_709:
		switch data[p] {
		case 32:
			goto tr1012
		case 46:
			goto tr1014
		case 117:
			goto st710
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1013
		}
		goto st0
	st710:
		if p++; p == pe {
			goto _test_eof710
		}
	st_case_710:
		if data[p] == 97 {
			goto st711
		}
		goto st0
	st711:
		if p++; p == pe {
			goto _test_eof711
		}
	st_case_711:
		if data[p] == 114 {
			goto st712
		}
		goto st0
	st712:
		if p++; p == pe {
			goto _test_eof712
		}
	st_case_712:
		if data[p] == 121 {
			goto st713
		}
		goto st0
	st713:
		if p++; p == pe {
			goto _test_eof713
		}
	st_case_713:
		switch data[p] {
		case 32:
			goto tr1012
		case 46:
			goto tr1014
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1013
		}
		goto st0
	st714:
		if p++; p == pe {
			goto _test_eof714
		}
	st_case_714:
		switch data[p] {
		case 108:
			goto st715
		case 110:
			goto st717
		}
		goto st0
	st715:
		if p++; p == pe {
			goto _test_eof715
		}
	st_case_715:
		switch data[p] {
		case 32:
			goto tr1021
		case 46:
			goto tr1023
		case 121:
			goto st716
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1022
		}
		goto st0
	st716:
		if p++; p == pe {
			goto _test_eof716
		}
	st_case_716:
		switch data[p] {
		case 32:
			goto tr1021
		case 46:
			goto tr1023
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1022
		}
		goto st0
	st717:
		if p++; p == pe {
			goto _test_eof717
		}
	st_case_717:
		switch data[p] {
		case 32:
			goto tr1025
		case 46:
			goto tr1027
		case 101:
			goto st718
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1026
		}
		goto st0
	st718:
		if p++; p == pe {
			goto _test_eof718
		}
	st_case_718:
		switch data[p] {
		case 32:
			goto tr1025
		case 46:
			goto tr1027
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1026
		}
		goto st0
	st719:
		if p++; p == pe {
			goto _test_eof719
		}
	st_case_719:
		switch data[p] {
		case 97:
			goto st720
		case 111:
			goto st725
		}
		goto st0
	st720:
		if p++; p == pe {
			goto _test_eof720
		}
	st_case_720:
		switch data[p] {
		case 114:
			goto st721
		case 121:
			goto st724
		}
		goto st0
	st721:
		if p++; p == pe {
			goto _test_eof721
		}
	st_case_721:
		switch data[p] {
		case 32:
			goto tr1033
		case 46:
			goto tr1035
		case 99:
			goto st722
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1034
		}
		goto st0
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
		if data[p] == 104 {
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
			goto tr1033
		case 46:
			goto tr1035
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1034
		}
		goto st0
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
	st_case_724:
		switch data[p] {
		case 32:
			goto tr1038
		case 46:
			goto tr1040
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1039
		}
		goto st0
	st725:
		if p++; p == pe {
			goto _test_eof725
		}
	st_case_725:
		if data[p] == 110 {
			goto st371
		}
		goto st0
	st726:
		if p++; p == pe {
			goto _test_eof726
		}
	st_case_726:
		if data[p] == 111 {
			goto st727
		}
		goto st0
	st727:
		if p++; p == pe {
			goto _test_eof727
		}
	st_case_727:
		if data[p] == 118 {
			goto st728
		}
		goto st0
	st728:
		if p++; p == pe {
			goto _test_eof728
		}
	st_case_728:
		switch data[p] {
		case 32:
			goto tr1043
		case 46:
			goto tr1045
		case 101:
			goto st729
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1044
		}
		goto st0
	st729:
		if p++; p == pe {
			goto _test_eof729
		}
	st_case_729:
		if data[p] == 109 {
			goto st730
		}
		goto st0
	st730:
		if p++; p == pe {
			goto _test_eof730
		}
	st_case_730:
		if data[p] == 98 {
			goto st731
		}
		goto st0
	st731:
		if p++; p == pe {
			goto _test_eof731
		}
	st_case_731:
		if data[p] == 101 {
			goto st732
		}
		goto st0
	st732:
		if p++; p == pe {
			goto _test_eof732
		}
	st_case_732:
		if data[p] == 114 {
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
			goto tr1043
		case 46:
			goto tr1045
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1044
		}
		goto st0
	st734:
		if p++; p == pe {
			goto _test_eof734
		}
	st_case_734:
		if data[p] == 99 {
			goto st735
		}
		goto st0
	st735:
		if p++; p == pe {
			goto _test_eof735
		}
	st_case_735:
		if data[p] == 116 {
			goto st736
		}
		goto st0
	st736:
		if p++; p == pe {
			goto _test_eof736
		}
	st_case_736:
		switch data[p] {
		case 32:
			goto tr1053
		case 46:
			goto tr1055
		case 111:
			goto st737
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1054
		}
		goto st0
	st737:
		if p++; p == pe {
			goto _test_eof737
		}
	st_case_737:
		if data[p] == 98 {
			goto st738
		}
		goto st0
	st738:
		if p++; p == pe {
			goto _test_eof738
		}
	st_case_738:
		if data[p] == 101 {
			goto st739
		}
		goto st0
	st739:
		if p++; p == pe {
			goto _test_eof739
		}
	st_case_739:
		if data[p] == 114 {
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
			goto tr1053
		case 46:
			goto tr1055
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1054
		}
		goto st0
	st741:
		if p++; p == pe {
			goto _test_eof741
		}
	st_case_741:
		switch data[p] {
		case 97:
			goto st742
		case 101:
			goto st746
		case 117:
			goto st725
		}
		goto st0
	st742:
		if p++; p == pe {
			goto _test_eof742
		}
	st_case_742:
		if data[p] == 116 {
			goto st743
		}
		goto st0
	st743:
		if p++; p == pe {
			goto _test_eof743
		}
	st_case_743:
		switch data[p] {
		case 32:
			goto st372
		case 44:
			goto st556
		case 117:
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
		if data[p] == 100 {
			goto st704
		}
		goto st0
	st746:
		if p++; p == pe {
			goto _test_eof746
		}
	st_case_746:
		if data[p] == 112 {
			goto st747
		}
		goto st0
	st747:
		if p++; p == pe {
			goto _test_eof747
		}
	st_case_747:
		switch data[p] {
		case 32:
			goto tr1066
		case 46:
			goto tr1068
		case 116:
			goto st748
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1067
		}
		goto st0
	st748:
		if p++; p == pe {
			goto _test_eof748
		}
	st_case_748:
		switch data[p] {
		case 32:
			goto tr1066
		case 46:
			goto tr1068
		case 101:
			goto st749
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1067
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
			goto tr1066
		case 46:
			goto tr1068
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1067
		}
		goto st0
	st754:
		if p++; p == pe {
			goto _test_eof754
		}
	st_case_754:
		switch data[p] {
		case 104:
			goto st755
		case 117:
			goto st758
		}
		goto st0
	st755:
		if p++; p == pe {
			goto _test_eof755
		}
	st_case_755:
		if data[p] == 117 {
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
			goto st372
		case 44:
			goto st556
		case 114:
			goto st757
		}
		goto st0
	st757:
		if p++; p == pe {
			goto _test_eof757
		}
	st_case_757:
		if data[p] == 115 {
			goto st745
		}
		goto st0
	st758:
		if p++; p == pe {
			goto _test_eof758
		}
	st_case_758:
		if data[p] == 101 {
			goto st759
		}
		goto st0
	st759:
		if p++; p == pe {
			goto _test_eof759
		}
	st_case_759:
		switch data[p] {
		case 32:
			goto st372
		case 44:
			goto st556
		case 115:
			goto st745
		}
		goto st0
	st760:
		if p++; p == pe {
			goto _test_eof760
		}
	st_case_760:
		if data[p] == 101 {
			goto st761
		}
		goto st0
	st761:
		if p++; p == pe {
			goto _test_eof761
		}
	st_case_761:
		if data[p] == 100 {
			goto st762
		}
		goto st0
	st762:
		if p++; p == pe {
			goto _test_eof762
		}
	st_case_762:
		switch data[p] {
		case 32:
			goto st372
		case 44:
			goto st556
		case 110:
			goto st763
		}
		goto st0
	st763:
		if p++; p == pe {
			goto _test_eof763
		}
	st_case_763:
		if data[p] == 101 {
			goto st757
		}
		goto st0
	st764:
		if p++; p == pe {
			goto _test_eof764
		}
	st_case_764:
		if data[p] == 101 {
			goto st363
		}
		goto st0
	st765:
		if p++; p == pe {
			goto _test_eof765
		}
	st_case_765:
		if data[p] == 97 {
			goto st720
		}
		goto st0
	st766:
		if p++; p == pe {
			goto _test_eof766
		}
	st_case_766:
		if data[p] == 101 {
			goto st746
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
	_test_eof767: cs = 767; goto _test_eof
	_test_eof768: cs = 768; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof769: cs = 769; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof770: cs = 770; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof771: cs = 771; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof772: cs = 772; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof773: cs = 773; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof774: cs = 774; goto _test_eof
	_test_eof775: cs = 775; goto _test_eof
	_test_eof776: cs = 776; goto _test_eof
	_test_eof777: cs = 777; goto _test_eof
	_test_eof778: cs = 778; goto _test_eof
	_test_eof779: cs = 779; goto _test_eof
	_test_eof780: cs = 780; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof781: cs = 781; goto _test_eof
	_test_eof782: cs = 782; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof783: cs = 783; goto _test_eof
	_test_eof784: cs = 784; goto _test_eof
	_test_eof785: cs = 785; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof786: cs = 786; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof787: cs = 787; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof788: cs = 788; goto _test_eof
	_test_eof789: cs = 789; goto _test_eof
	_test_eof790: cs = 790; goto _test_eof
	_test_eof791: cs = 791; goto _test_eof
	_test_eof792: cs = 792; goto _test_eof
	_test_eof793: cs = 793; goto _test_eof
	_test_eof794: cs = 794; goto _test_eof
	_test_eof795: cs = 795; goto _test_eof
	_test_eof796: cs = 796; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof797: cs = 797; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof798: cs = 798; goto _test_eof
	_test_eof799: cs = 799; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof800: cs = 800; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
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
	_test_eof30: cs = 30; goto _test_eof
	_test_eof814: cs = 814; goto _test_eof
	_test_eof815: cs = 815; goto _test_eof
	_test_eof816: cs = 816; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof817: cs = 817; goto _test_eof
	_test_eof818: cs = 818; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
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
	_test_eof819: cs = 819; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof820: cs = 820; goto _test_eof
	_test_eof821: cs = 821; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof822: cs = 822; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof823: cs = 823; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
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
	_test_eof824: cs = 824; goto _test_eof
	_test_eof272: cs = 272; goto _test_eof
	_test_eof273: cs = 273; goto _test_eof
	_test_eof825: cs = 825; goto _test_eof
	_test_eof274: cs = 274; goto _test_eof
	_test_eof275: cs = 275; goto _test_eof
	_test_eof276: cs = 276; goto _test_eof
	_test_eof826: cs = 826; goto _test_eof
	_test_eof827: cs = 827; goto _test_eof
	_test_eof828: cs = 828; goto _test_eof
	_test_eof829: cs = 829; goto _test_eof
	_test_eof830: cs = 830; goto _test_eof
	_test_eof831: cs = 831; goto _test_eof
	_test_eof832: cs = 832; goto _test_eof
	_test_eof277: cs = 277; goto _test_eof
	_test_eof833: cs = 833; goto _test_eof
	_test_eof278: cs = 278; goto _test_eof
	_test_eof279: cs = 279; goto _test_eof
	_test_eof280: cs = 280; goto _test_eof
	_test_eof281: cs = 281; goto _test_eof
	_test_eof282: cs = 282; goto _test_eof
	_test_eof283: cs = 283; goto _test_eof
	_test_eof834: cs = 834; goto _test_eof
	_test_eof835: cs = 835; goto _test_eof
	_test_eof836: cs = 836; goto _test_eof
	_test_eof837: cs = 837; goto _test_eof
	_test_eof284: cs = 284; goto _test_eof
	_test_eof838: cs = 838; goto _test_eof
	_test_eof839: cs = 839; goto _test_eof
	_test_eof285: cs = 285; goto _test_eof
	_test_eof840: cs = 840; goto _test_eof
	_test_eof841: cs = 841; goto _test_eof
	_test_eof286: cs = 286; goto _test_eof
	_test_eof842: cs = 842; goto _test_eof
	_test_eof843: cs = 843; goto _test_eof
	_test_eof287: cs = 287; goto _test_eof
	_test_eof844: cs = 844; goto _test_eof
	_test_eof845: cs = 845; goto _test_eof
	_test_eof288: cs = 288; goto _test_eof
	_test_eof846: cs = 846; goto _test_eof
	_test_eof289: cs = 289; goto _test_eof
	_test_eof847: cs = 847; goto _test_eof
	_test_eof290: cs = 290; goto _test_eof
	_test_eof848: cs = 848; goto _test_eof
	_test_eof849: cs = 849; goto _test_eof
	_test_eof850: cs = 850; goto _test_eof
	_test_eof851: cs = 851; goto _test_eof
	_test_eof852: cs = 852; goto _test_eof
	_test_eof853: cs = 853; goto _test_eof
	_test_eof854: cs = 854; goto _test_eof
	_test_eof855: cs = 855; goto _test_eof
	_test_eof856: cs = 856; goto _test_eof
	_test_eof291: cs = 291; goto _test_eof
	_test_eof857: cs = 857; goto _test_eof
	_test_eof292: cs = 292; goto _test_eof
	_test_eof858: cs = 858; goto _test_eof
	_test_eof859: cs = 859; goto _test_eof
	_test_eof293: cs = 293; goto _test_eof
	_test_eof860: cs = 860; goto _test_eof
	_test_eof294: cs = 294; goto _test_eof
	_test_eof861: cs = 861; goto _test_eof
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
	_test_eof295: cs = 295; goto _test_eof
	_test_eof874: cs = 874; goto _test_eof
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
	_test_eof875: cs = 875; goto _test_eof
	_test_eof876: cs = 876; goto _test_eof
	_test_eof334: cs = 334; goto _test_eof
	_test_eof335: cs = 335; goto _test_eof
	_test_eof336: cs = 336; goto _test_eof
	_test_eof337: cs = 337; goto _test_eof
	_test_eof338: cs = 338; goto _test_eof
	_test_eof877: cs = 877; goto _test_eof
	_test_eof339: cs = 339; goto _test_eof
	_test_eof340: cs = 340; goto _test_eof
	_test_eof341: cs = 341; goto _test_eof
	_test_eof342: cs = 342; goto _test_eof
	_test_eof343: cs = 343; goto _test_eof
	_test_eof344: cs = 344; goto _test_eof
	_test_eof345: cs = 345; goto _test_eof
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
	_test_eof385: cs = 385; goto _test_eof
	_test_eof386: cs = 386; goto _test_eof
	_test_eof878: cs = 878; goto _test_eof
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
	_test_eof536: cs = 536; goto _test_eof
	_test_eof537: cs = 537; goto _test_eof
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
	_test_eof879: cs = 879; goto _test_eof
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

	_test_eof: {}
	if p == eof {
		switch cs {
		case 777, 785, 816, 829, 833:
//line ragel/datetime.rl:7
 st.Zoned = true 
		case 821:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

		case 823, 824, 834, 875, 877, 878:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

		case 822, 879:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

		case 819, 820:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

		case 787, 797, 847, 857:
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

		case 770, 774:
//line ragel/datetime.rl:50

    st.Ad_bc = ADBC_BC;

		case 783, 844:
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

		case 767, 817, 818:
//line ragel/datetime.rl:90

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

		case 786, 837, 846:
//line ragel/datetime.rl:99

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

		case 781, 813, 814, 836, 839, 840, 842, 873, 874:
//line ragel/datetime.rl:102

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

		case 799, 859:
//line ragel/datetime.rl:105

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

		case 798, 812, 858, 872:
//line ragel/datetime.rl:108

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

		case 810, 870:
//line ragel/datetime.rl:111

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

		case 800, 811, 860, 871:
//line ragel/datetime.rl:114

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

		case 788, 789, 790, 791, 792, 793, 794, 795, 796, 801, 802, 803, 804, 805, 806, 807, 808, 809, 848, 849, 850, 851, 852, 853, 854, 855, 856, 861, 862, 863, 864, 865, 866, 867, 868, 869:
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

		case 838:
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

		case 778, 779, 830, 831:
//line ragel/datetime.rl:157

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:7
 st.Zoned = true 
		case 771, 775, 776, 780, 825, 827, 828, 832:
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
		case 772, 826:
//line ragel/datetime.rl:194

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
//line ragel/parse_datetime.go:22875
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

