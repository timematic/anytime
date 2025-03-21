
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
const datetime_parser_first_final int = 732
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
	case 732:
		goto st_case_732
	case 733:
		goto st_case_733
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 734:
		goto st_case_734
	case 11:
		goto st_case_11
	case 735:
		goto st_case_735
	case 12:
		goto st_case_12
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
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 743:
		goto st_case_743
	case 744:
		goto st_case_744
	case 745:
		goto st_case_745
	case 16:
		goto st_case_16
	case 746:
		goto st_case_746
	case 747:
		goto st_case_747
	case 748:
		goto st_case_748
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 749:
		goto st_case_749
	case 750:
		goto st_case_750
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 751:
		goto st_case_751
	case 21:
		goto st_case_21
	case 752:
		goto st_case_752
	case 22:
		goto st_case_22
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
	case 23:
		goto st_case_23
	case 762:
		goto st_case_762
	case 24:
		goto st_case_24
	case 763:
		goto st_case_763
	case 764:
		goto st_case_764
	case 25:
		goto st_case_25
	case 765:
		goto st_case_765
	case 26:
		goto st_case_26
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
	case 27:
		goto st_case_27
	case 779:
		goto st_case_779
	case 780:
		goto st_case_780
	case 781:
		goto st_case_781
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 782:
		goto st_case_782
	case 783:
		goto st_case_783
	case 784:
		goto st_case_784
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
	case 105:
		goto st_case_105
	case 785:
		goto st_case_785
	case 786:
		goto st_case_786
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
	case 787:
		goto st_case_787
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
	case 788:
		goto st_case_788
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
	case 789:
		goto st_case_789
	case 266:
		goto st_case_266
	case 267:
		goto st_case_267
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
	case 268:
		goto st_case_268
	case 269:
		goto st_case_269
	case 270:
		goto st_case_270
	case 797:
		goto st_case_797
	case 798:
		goto st_case_798
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
	case 336:
		goto st_case_336
	case 337:
		goto st_case_337
	case 338:
		goto st_case_338
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
	case 799:
		goto st_case_799
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
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 48:
			goto tr1
		case 51:
			goto tr3
		case 65:
			goto st254
		case 68:
			goto st325
		case 70:
			goto st333
		case 74:
			goto st672
		case 77:
			goto st684
		case 78:
			goto st691
		case 79:
			goto st699
		case 83:
			goto st706
		case 84:
			goto st719
		case 87:
			goto st725
		case 97:
			goto st254
		case 100:
			goto st325
		case 102:
			goto st729
		case 106:
			goto st672
		case 109:
			goto st730
		case 110:
			goto st691
		case 111:
			goto st699
		case 115:
			goto st731
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr4
			}
		case data[p] >= 49:
			goto tr2
		}
		goto tr0
tr0:
//line ragel/datetime.rl:5
 return st, err 
	goto st0
//line ragel/parse_datetime.go:1716
st_case_0:
	st0:
		cs = 0
		goto _out
tr1:
//line ragel/datetime.rl:7
 pb = p 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line ragel/parse_datetime.go:1730
		if data[p] == 48 {
			goto st3
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st106
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto tr0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 46 {
			goto tr23
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr24
			}
		case data[p] >= 45:
			goto tr22
		}
		goto tr0
tr22:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line ragel/parse_datetime.go:1784
		switch data[p] {
		case 48:
			goto tr25
		case 49:
			goto tr26
		case 65:
			goto st34
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
			goto st34
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
		goto tr0
tr25:
//line ragel/datetime.rl:7
 pb = p 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line ragel/parse_datetime.go:1836
		if data[p] == 48 {
			goto st8
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st29
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if 48 <= data[p] && data[p] <= 57 {
			goto st732
		}
		goto tr0
	st732:
		if p++; p == pe {
			goto _test_eof732
		}
	st_case_732:
		switch data[p] {
		case 32:
			goto tr945
		case 43:
			goto tr947
		case 45:
			goto tr948
		case 47:
			goto tr949
		case 84:
			goto tr950
		case 90:
			goto tr951
		case 95:
			goto tr952
		case 116:
			goto tr952
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr949
			}
		case data[p] >= 65:
			goto tr949
		}
		goto st0
tr1084:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st733
tr1062:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st733
tr945:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st733
tr1070:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st733
tr1077:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st733
	st733:
		if p++; p == pe {
			goto _test_eof733
		}
	st_case_733:
//line ragel/parse_datetime.go:1926
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr956
		case 50:
			goto tr64
		case 65:
			goto tr957
		case 66:
			goto tr958
		case 90:
			goto tr959
		case 95:
			goto tr956
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr63
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr956
				}
			case data[p] >= 67:
				goto tr956
			}
		default:
			goto tr65
		}
		goto st0
tr960:
//line ragel/datetime.rl:168
 
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

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st9
tr964:
//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st9
tr967:
//line ragel/datetime.rl:159

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st9
tr969:
//line ragel/datetime.rl:196

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line ragel/parse_datetime.go:2028
		switch data[p] {
		case 65:
			goto st10
		case 66:
			goto st11
		}
		goto tr0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 68 {
			goto st734
		}
		goto tr0
	st734:
		if p++; p == pe {
			goto _test_eof734
		}
	st_case_734:
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 67 {
			goto st735
		}
		goto tr0
	st735:
		if p++; p == pe {
			goto _test_eof735
		}
	st_case_735:
		goto st0
tr1085:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st12
tr1063:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st12
tr971:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st12
tr983:
//line ragel/datetime.rl:56

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

	goto st12
tr987:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st12
tr995:
//line ragel/datetime.rl:37

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

	goto st12
tr1002:
//line ragel/datetime.rl:119

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

	goto st12
tr1015:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st12
tr1024:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st12
tr1032:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st12
tr1052:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st12
tr947:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st12
tr1071:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st12
tr1078:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line ragel/parse_datetime.go:2216
		if data[p] == 50 {
			goto tr44
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr45
			}
		case data[p] >= 48:
			goto tr43
		}
		goto tr0
tr43:
//line ragel/datetime.rl:7
 pb = p 
	goto st736
tr46:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st736
	st736:
		if p++; p == pe {
			goto _test_eof736
		}
	st_case_736:
//line ragel/parse_datetime.go:2244
		switch data[p] {
		case 32:
			goto tr960
		case 58:
			goto tr962
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st737
		}
		goto st0
tr45:
//line ragel/datetime.rl:7
 pb = p 
	goto st737
tr48:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st737
	st737:
		if p++; p == pe {
			goto _test_eof737
		}
	st_case_737:
//line ragel/parse_datetime.go:2270
		switch data[p] {
		case 32:
			goto tr960
		case 58:
			goto tr962
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st738
		}
		goto st0
	st738:
		if p++; p == pe {
			goto _test_eof738
		}
	st_case_738:
		if data[p] == 32 {
			goto tr960
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st738
		}
		goto st0
tr962:
//line ragel/datetime.rl:150

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st739
	st739:
		if p++; p == pe {
			goto _test_eof739
		}
	st_case_739:
//line ragel/parse_datetime.go:2309
		if data[p] == 32 {
			goto tr964
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr966
			}
		case data[p] >= 48:
			goto tr965
		}
		goto st0
tr965:
//line ragel/datetime.rl:7
 pb = p 
	goto st740
	st740:
		if p++; p == pe {
			goto _test_eof740
		}
	st_case_740:
//line ragel/parse_datetime.go:2331
		if data[p] == 32 {
			goto tr967
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st741
		}
		goto st0
tr966:
//line ragel/datetime.rl:7
 pb = p 
	goto st741
	st741:
		if p++; p == pe {
			goto _test_eof741
		}
	st_case_741:
//line ragel/parse_datetime.go:2348
		if data[p] == 32 {
			goto tr967
		}
		goto st0
tr44:
//line ragel/datetime.rl:7
 pb = p 
	goto st742
tr47:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st742
	st742:
		if p++; p == pe {
			goto _test_eof742
		}
	st_case_742:
//line ragel/parse_datetime.go:2368
		switch data[p] {
		case 32:
			goto tr960
		case 58:
			goto tr962
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st738
			}
		case data[p] >= 48:
			goto st737
		}
		goto st0
tr1086:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st13
tr1064:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st13
tr972:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st13
tr984:
//line ragel/datetime.rl:56

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

	goto st13
tr988:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st13
tr996:
//line ragel/datetime.rl:37

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

	goto st13
tr1003:
//line ragel/datetime.rl:119

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

	goto st13
tr1016:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st13
tr1025:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st13
tr1033:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st13
tr1053:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st13
tr948:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st13
tr1072:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st13
tr1079:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line ragel/parse_datetime.go:2534
		if data[p] == 50 {
			goto tr47
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr48
			}
		case data[p] >= 48:
			goto tr46
		}
		goto tr0
tr956:
//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1087:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr973:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr989:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1017:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1026:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1035:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1004:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1055:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr998:
//line ragel/datetime.rl:37

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

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr949:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1065:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1073:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1080:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line ragel/parse_datetime.go:2702
		switch data[p] {
		case 47:
			goto st15
		case 95:
			goto st15
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		case data[p] >= 65:
			goto st15
		}
		goto tr0
tr1060:
//line ragel/datetime.rl:7
 pb = p 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line ragel/parse_datetime.go:2727
		switch data[p] {
		case 47:
			goto st743
		case 95:
			goto st743
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st743
			}
		case data[p] >= 65:
			goto st743
		}
		goto tr0
tr985:
//line ragel/datetime.rl:56

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

//line ragel/datetime.rl:7
 pb = p 
	goto st743
	st743:
		if p++; p == pe {
			goto _test_eof743
		}
	st_case_743:
//line ragel/parse_datetime.go:2775
		switch data[p] {
		case 32:
			goto tr969
		case 47:
			goto st743
		case 95:
			goto st743
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st743
			}
		case data[p] >= 65:
			goto st743
		}
		goto st0
tr63:
//line ragel/datetime.rl:7
 pb = p 
	goto st744
	st744:
		if p++; p == pe {
			goto _test_eof744
		}
	st_case_744:
//line ragel/parse_datetime.go:2802
		switch data[p] {
		case 32:
			goto tr970
		case 43:
			goto tr971
		case 45:
			goto tr972
		case 47:
			goto tr973
		case 58:
			goto tr975
		case 65:
			goto tr976
		case 80:
			goto tr976
		case 90:
			goto tr977
		case 95:
			goto tr973
		case 97:
			goto tr978
		case 112:
			goto tr978
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st751
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr973
			}
		default:
			goto tr973
		}
		goto st0
tr970:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st745
tr986:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st745
tr1040:
//line ragel/datetime.rl:119

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

	goto st745
tr1014:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st745
tr1023:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st745
tr1031:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st745
tr1051:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st745
	st745:
		if p++; p == pe {
			goto _test_eof745
		}
	st_case_745:
//line ragel/parse_datetime.go:2912
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr956
		case 65:
			goto tr979
		case 66:
			goto tr958
		case 80:
			goto tr980
		case 90:
			goto tr959
		case 95:
			goto tr956
		case 97:
			goto tr981
		case 112:
			goto tr981
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr956
			}
		case data[p] >= 67:
			goto tr956
		}
		goto st0
tr979:
//line ragel/datetime.rl:7
 pb = p 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line ragel/parse_datetime.go:2955
		switch data[p] {
		case 47:
			goto st15
		case 68:
			goto st746
		case 77:
			goto st747
		case 95:
			goto st15
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		case data[p] >= 65:
			goto st15
		}
		goto tr0
	st746:
		if p++; p == pe {
			goto _test_eof746
		}
	st_case_746:
		switch data[p] {
		case 47:
			goto st743
		case 95:
			goto st743
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st743
			}
		case data[p] >= 65:
			goto st743
		}
		goto st0
	st747:
		if p++; p == pe {
			goto _test_eof747
		}
	st_case_747:
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr983
		case 45:
			goto tr984
		case 47:
			goto tr985
		case 95:
			goto tr985
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr985
			}
		case data[p] >= 65:
			goto tr985
		}
		goto st0
tr982:
//line ragel/datetime.rl:56

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

	goto st748
tr994:
//line ragel/datetime.rl:37

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

	goto st748
tr1001:
//line ragel/datetime.rl:119

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

	goto st748
	st748:
		if p++; p == pe {
			goto _test_eof748
		}
	st_case_748:
//line ragel/parse_datetime.go:3099
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr956
		case 65:
			goto tr957
		case 66:
			goto tr958
		case 90:
			goto tr959
		case 95:
			goto tr956
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr956
			}
		case data[p] >= 67:
			goto tr956
		}
		goto st0
tr957:
//line ragel/datetime.rl:7
 pb = p 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line ragel/parse_datetime.go:3136
		switch data[p] {
		case 47:
			goto st15
		case 68:
			goto st746
		case 95:
			goto st15
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		case data[p] >= 65:
			goto st15
		}
		goto tr0
tr958:
//line ragel/datetime.rl:7
 pb = p 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line ragel/parse_datetime.go:3163
		switch data[p] {
		case 47:
			goto st15
		case 67:
			goto st749
		case 95:
			goto st15
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		case data[p] >= 65:
			goto st15
		}
		goto tr0
	st749:
		if p++; p == pe {
			goto _test_eof749
		}
	st_case_749:
		switch data[p] {
		case 47:
			goto st743
		case 95:
			goto st743
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st743
			}
		case data[p] >= 65:
			goto st743
		}
		goto st0
tr959:
//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1089:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr977:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr992:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1021:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1029:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1038:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1006:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1057:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1000:
//line ragel/datetime.rl:37

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

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr951:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1067:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1075:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
tr1082:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st750
	st750:
		if p++; p == pe {
			goto _test_eof750
		}
	st_case_750:
//line ragel/parse_datetime.go:3356
		switch data[p] {
		case 32:
			goto tr964
		case 47:
			goto st15
		case 95:
			goto st15
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		case data[p] >= 65:
			goto st15
		}
		goto st0
tr980:
//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr976:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr991:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr1020:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr1028:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr1037:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr1042:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr1056:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line ragel/parse_datetime.go:3464
		switch data[p] {
		case 47:
			goto st15
		case 77:
			goto st747
		case 95:
			goto st15
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		case data[p] >= 65:
			goto st15
		}
		goto tr0
tr981:
//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr978:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr993:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr1022:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr1030:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr1039:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr1043:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr1058:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line ragel/parse_datetime.go:3572
		switch data[p] {
		case 47:
			goto st15
		case 95:
			goto st15
		case 109:
			goto st747
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		case data[p] >= 65:
			goto st15
		}
		goto tr0
	st751:
		if p++; p == pe {
			goto _test_eof751
		}
	st_case_751:
		switch data[p] {
		case 32:
			goto tr986
		case 43:
			goto tr987
		case 45:
			goto tr988
		case 47:
			goto tr989
		case 58:
			goto tr990
		case 65:
			goto tr991
		case 80:
			goto tr991
		case 90:
			goto tr992
		case 95:
			goto tr989
		case 97:
			goto tr993
		case 112:
			goto tr993
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st21
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr989
			}
		default:
			goto tr989
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= data[p] && data[p] <= 57 {
			goto st752
		}
		goto tr0
	st752:
		if p++; p == pe {
			goto _test_eof752
		}
	st_case_752:
		switch data[p] {
		case 32:
			goto tr994
		case 43:
			goto tr995
		case 45:
			goto tr996
		case 46:
			goto tr997
		case 47:
			goto tr998
		case 90:
			goto tr1000
		case 95:
			goto tr998
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr998
			}
		default:
			goto tr998
		}
		goto st0
tr997:
//line ragel/datetime.rl:37

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

	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line ragel/parse_datetime.go:3697
		if 48 <= data[p] && data[p] <= 57 {
			goto tr55
		}
		goto tr0
tr55:
//line ragel/datetime.rl:7
 pb = p 
	goto st753
	st753:
		if p++; p == pe {
			goto _test_eof753
		}
	st_case_753:
//line ragel/parse_datetime.go:3711
		switch data[p] {
		case 32:
			goto tr1001
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st754
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st754:
		if p++; p == pe {
			goto _test_eof754
		}
	st_case_754:
		switch data[p] {
		case 32:
			goto tr1001
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st755
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st755:
		if p++; p == pe {
			goto _test_eof755
		}
	st_case_755:
		switch data[p] {
		case 32:
			goto tr1001
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st756
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st756:
		if p++; p == pe {
			goto _test_eof756
		}
	st_case_756:
		switch data[p] {
		case 32:
			goto tr1001
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st757
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st757:
		if p++; p == pe {
			goto _test_eof757
		}
	st_case_757:
		switch data[p] {
		case 32:
			goto tr1001
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st758
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st758:
		if p++; p == pe {
			goto _test_eof758
		}
	st_case_758:
		switch data[p] {
		case 32:
			goto tr1001
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st759
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st759:
		if p++; p == pe {
			goto _test_eof759
		}
	st_case_759:
		switch data[p] {
		case 32:
			goto tr1001
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st760
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st760:
		if p++; p == pe {
			goto _test_eof760
		}
	st_case_760:
		switch data[p] {
		case 32:
			goto tr1001
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st761
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st761:
		if p++; p == pe {
			goto _test_eof761
		}
	st_case_761:
		switch data[p] {
		case 32:
			goto tr1001
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		case data[p] >= 65:
			goto tr1004
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if 48 <= data[p] && data[p] <= 57 {
			goto st762
		}
		goto tr0
	st762:
		if p++; p == pe {
			goto _test_eof762
		}
	st_case_762:
		switch data[p] {
		case 32:
			goto tr994
		case 43:
			goto tr995
		case 45:
			goto tr996
		case 46:
			goto tr997
		case 47:
			goto tr998
		case 90:
			goto tr1000
		case 95:
			goto tr998
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr998
			}
		case data[p] >= 65:
			goto tr998
		}
		goto st0
tr975:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st24
tr990:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line ragel/parse_datetime.go:4047
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr58
			}
		case data[p] >= 48:
			goto tr57
		}
		goto tr0
tr57:
//line ragel/datetime.rl:7
 pb = p 
	goto st763
	st763:
		if p++; p == pe {
			goto _test_eof763
		}
	st_case_763:
//line ragel/parse_datetime.go:4066
		switch data[p] {
		case 32:
			goto tr1014
		case 43:
			goto tr1015
		case 45:
			goto tr1016
		case 47:
			goto tr1017
		case 58:
			goto tr1019
		case 65:
			goto tr1020
		case 80:
			goto tr1020
		case 90:
			goto tr1021
		case 95:
			goto tr1017
		case 97:
			goto tr1022
		case 112:
			goto tr1022
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st764
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1017
			}
		default:
			goto tr1017
		}
		goto st0
	st764:
		if p++; p == pe {
			goto _test_eof764
		}
	st_case_764:
		switch data[p] {
		case 32:
			goto tr1023
		case 43:
			goto tr1024
		case 45:
			goto tr1025
		case 47:
			goto tr1026
		case 58:
			goto tr1027
		case 65:
			goto tr1028
		case 80:
			goto tr1028
		case 90:
			goto tr1029
		case 95:
			goto tr1026
		case 97:
			goto tr1030
		case 112:
			goto tr1030
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1026
			}
		case data[p] >= 66:
			goto tr1026
		}
		goto st0
tr1019:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st25
tr1027:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line ragel/parse_datetime.go:4159
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr60
			}
		case data[p] >= 48:
			goto tr59
		}
		goto tr0
tr59:
//line ragel/datetime.rl:7
 pb = p 
	goto st765
	st765:
		if p++; p == pe {
			goto _test_eof765
		}
	st_case_765:
//line ragel/parse_datetime.go:4178
		switch data[p] {
		case 32:
			goto tr1031
		case 43:
			goto tr1032
		case 45:
			goto tr1033
		case 46:
			goto tr1034
		case 47:
			goto tr1035
		case 65:
			goto tr1037
		case 80:
			goto tr1037
		case 90:
			goto tr1038
		case 95:
			goto tr1035
		case 97:
			goto tr1039
		case 112:
			goto tr1039
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st775
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1035
			}
		default:
			goto tr1035
		}
		goto st0
tr1034:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st26
tr1054:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line ragel/parse_datetime.go:4233
		if 48 <= data[p] && data[p] <= 57 {
			goto tr61
		}
		goto tr0
tr61:
//line ragel/datetime.rl:7
 pb = p 
	goto st766
	st766:
		if p++; p == pe {
			goto _test_eof766
		}
	st_case_766:
//line ragel/parse_datetime.go:4247
		switch data[p] {
		case 32:
			goto tr1040
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 65:
			goto tr1042
		case 80:
			goto tr1042
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		case 97:
			goto tr1043
		case 112:
			goto tr1043
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st767
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st767:
		if p++; p == pe {
			goto _test_eof767
		}
	st_case_767:
		switch data[p] {
		case 32:
			goto tr1040
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 65:
			goto tr1042
		case 80:
			goto tr1042
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		case 97:
			goto tr1043
		case 112:
			goto tr1043
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st768
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st768:
		if p++; p == pe {
			goto _test_eof768
		}
	st_case_768:
		switch data[p] {
		case 32:
			goto tr1040
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 65:
			goto tr1042
		case 80:
			goto tr1042
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		case 97:
			goto tr1043
		case 112:
			goto tr1043
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st769
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st769:
		if p++; p == pe {
			goto _test_eof769
		}
	st_case_769:
		switch data[p] {
		case 32:
			goto tr1040
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 65:
			goto tr1042
		case 80:
			goto tr1042
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		case 97:
			goto tr1043
		case 112:
			goto tr1043
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st770
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st770:
		if p++; p == pe {
			goto _test_eof770
		}
	st_case_770:
		switch data[p] {
		case 32:
			goto tr1040
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 65:
			goto tr1042
		case 80:
			goto tr1042
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		case 97:
			goto tr1043
		case 112:
			goto tr1043
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st771
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st771:
		if p++; p == pe {
			goto _test_eof771
		}
	st_case_771:
		switch data[p] {
		case 32:
			goto tr1040
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 65:
			goto tr1042
		case 80:
			goto tr1042
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		case 97:
			goto tr1043
		case 112:
			goto tr1043
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st772
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st772:
		if p++; p == pe {
			goto _test_eof772
		}
	st_case_772:
		switch data[p] {
		case 32:
			goto tr1040
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 65:
			goto tr1042
		case 80:
			goto tr1042
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		case 97:
			goto tr1043
		case 112:
			goto tr1043
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st773
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st773:
		if p++; p == pe {
			goto _test_eof773
		}
	st_case_773:
		switch data[p] {
		case 32:
			goto tr1040
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 65:
			goto tr1042
		case 80:
			goto tr1042
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		case 97:
			goto tr1043
		case 112:
			goto tr1043
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st774
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		default:
			goto tr1004
		}
		goto st0
	st774:
		if p++; p == pe {
			goto _test_eof774
		}
	st_case_774:
		switch data[p] {
		case 32:
			goto tr1040
		case 43:
			goto tr1002
		case 45:
			goto tr1003
		case 47:
			goto tr1004
		case 65:
			goto tr1042
		case 80:
			goto tr1042
		case 90:
			goto tr1006
		case 95:
			goto tr1004
		case 97:
			goto tr1043
		case 112:
			goto tr1043
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1004
			}
		case data[p] >= 66:
			goto tr1004
		}
		goto st0
	st775:
		if p++; p == pe {
			goto _test_eof775
		}
	st_case_775:
		switch data[p] {
		case 32:
			goto tr1051
		case 43:
			goto tr1052
		case 45:
			goto tr1053
		case 46:
			goto tr1054
		case 47:
			goto tr1055
		case 65:
			goto tr1056
		case 80:
			goto tr1056
		case 90:
			goto tr1057
		case 95:
			goto tr1055
		case 97:
			goto tr1058
		case 112:
			goto tr1058
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1055
			}
		case data[p] >= 66:
			goto tr1055
		}
		goto st0
tr60:
//line ragel/datetime.rl:7
 pb = p 
	goto st776
	st776:
		if p++; p == pe {
			goto _test_eof776
		}
	st_case_776:
//line ragel/parse_datetime.go:4646
		switch data[p] {
		case 32:
			goto tr1031
		case 43:
			goto tr1032
		case 45:
			goto tr1033
		case 46:
			goto tr1034
		case 47:
			goto tr1035
		case 65:
			goto tr1037
		case 80:
			goto tr1037
		case 90:
			goto tr1038
		case 95:
			goto tr1035
		case 97:
			goto tr1039
		case 112:
			goto tr1039
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1035
			}
		case data[p] >= 66:
			goto tr1035
		}
		goto st0
tr58:
//line ragel/datetime.rl:7
 pb = p 
	goto st777
	st777:
		if p++; p == pe {
			goto _test_eof777
		}
	st_case_777:
//line ragel/parse_datetime.go:4689
		switch data[p] {
		case 32:
			goto tr1014
		case 43:
			goto tr1015
		case 45:
			goto tr1016
		case 47:
			goto tr1017
		case 58:
			goto tr1019
		case 65:
			goto tr1020
		case 80:
			goto tr1020
		case 90:
			goto tr1021
		case 95:
			goto tr1017
		case 97:
			goto tr1022
		case 112:
			goto tr1022
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1017
			}
		case data[p] >= 66:
			goto tr1017
		}
		goto st0
tr64:
//line ragel/datetime.rl:7
 pb = p 
	goto st778
	st778:
		if p++; p == pe {
			goto _test_eof778
		}
	st_case_778:
//line ragel/parse_datetime.go:4732
		switch data[p] {
		case 32:
			goto tr970
		case 43:
			goto tr971
		case 45:
			goto tr972
		case 47:
			goto tr973
		case 58:
			goto tr975
		case 65:
			goto tr976
		case 80:
			goto tr976
		case 90:
			goto tr977
		case 95:
			goto tr973
		case 97:
			goto tr978
		case 112:
			goto tr978
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st751
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr973
				}
			case data[p] >= 66:
				goto tr973
			}
		default:
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if 48 <= data[p] && data[p] <= 57 {
			goto st21
		}
		goto tr0
tr65:
//line ragel/datetime.rl:7
 pb = p 
	goto st779
	st779:
		if p++; p == pe {
			goto _test_eof779
		}
	st_case_779:
//line ragel/parse_datetime.go:4793
		switch data[p] {
		case 32:
			goto tr970
		case 43:
			goto tr971
		case 45:
			goto tr972
		case 47:
			goto tr973
		case 58:
			goto tr975
		case 65:
			goto tr976
		case 80:
			goto tr976
		case 90:
			goto tr977
		case 95:
			goto tr973
		case 97:
			goto tr978
		case 112:
			goto tr978
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st27
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr973
			}
		default:
			goto tr973
		}
		goto st0
tr1088:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st780
tr950:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st780
tr1066:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st780
tr1074:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st780
tr1081:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st780
	st780:
		if p++; p == pe {
			goto _test_eof780
		}
	st_case_780:
//line ragel/parse_datetime.go:4882
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr1060
		case 50:
			goto tr64
		case 90:
			goto tr1061
		case 95:
			goto tr1060
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr63
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1060
				}
			case data[p] >= 65:
				goto tr1060
			}
		default:
			goto tr65
		}
		goto st0
tr1061:
//line ragel/datetime.rl:7
 pb = p 
	goto st781
	st781:
		if p++; p == pe {
			goto _test_eof781
		}
	st_case_781:
//line ragel/parse_datetime.go:4926
		switch data[p] {
		case 32:
			goto tr964
		case 47:
			goto st743
		case 95:
			goto st743
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st743
			}
		case data[p] >= 65:
			goto st743
		}
		goto st0
tr1090:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st28
tr952:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st28
tr1068:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st28
tr1076:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st28
tr1083:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line ragel/parse_datetime.go:4995
		switch data[p] {
		case 47:
			goto st15
		case 50:
			goto tr64
		case 95:
			goto st15
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr63
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st15
				}
			case data[p] >= 65:
				goto st15
			}
		default:
			goto tr65
		}
		goto tr0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 45:
			goto tr66
		case 47:
			goto tr66
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st732
		}
		goto tr0
tr66:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st30
tr75:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st30
tr79:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st30
tr85:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st30
tr93:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st30
tr102:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st30
tr109:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st30
tr111:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st30
tr116:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st30
tr119:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st30
tr122:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st30
tr130:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st30
tr137:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line ragel/parse_datetime.go:5096
		switch data[p] {
		case 48:
			goto tr67
		case 51:
			goto tr69
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr70
			}
		case data[p] >= 49:
			goto tr68
		}
		goto tr0
tr67:
//line ragel/datetime.rl:7
 pb = p 
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line ragel/parse_datetime.go:5121
		if 49 <= data[p] && data[p] <= 57 {
			goto st782
		}
		goto tr0
tr70:
//line ragel/datetime.rl:7
 pb = p 
	goto st782
	st782:
		if p++; p == pe {
			goto _test_eof782
		}
	st_case_782:
//line ragel/parse_datetime.go:5135
		switch data[p] {
		case 32:
			goto tr1062
		case 43:
			goto tr1063
		case 45:
			goto tr1064
		case 47:
			goto tr1065
		case 84:
			goto tr1066
		case 90:
			goto tr1067
		case 95:
			goto tr1068
		case 116:
			goto tr1068
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1065
			}
		case data[p] >= 65:
			goto tr1065
		}
		goto st0
tr68:
//line ragel/datetime.rl:7
 pb = p 
	goto st783
	st783:
		if p++; p == pe {
			goto _test_eof783
		}
	st_case_783:
//line ragel/parse_datetime.go:5172
		switch data[p] {
		case 32:
			goto tr1062
		case 43:
			goto tr1063
		case 45:
			goto tr1064
		case 47:
			goto tr1065
		case 84:
			goto tr1066
		case 90:
			goto tr1067
		case 95:
			goto tr1068
		case 116:
			goto tr1068
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st782
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1065
			}
		default:
			goto tr1065
		}
		goto st0
tr69:
//line ragel/datetime.rl:7
 pb = p 
	goto st784
	st784:
		if p++; p == pe {
			goto _test_eof784
		}
	st_case_784:
//line ragel/parse_datetime.go:5213
		switch data[p] {
		case 32:
			goto tr1062
		case 43:
			goto tr1063
		case 45:
			goto tr1064
		case 47:
			goto tr1065
		case 84:
			goto tr1066
		case 90:
			goto tr1067
		case 95:
			goto tr1068
		case 116:
			goto tr1068
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 49 {
				goto st782
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1065
			}
		default:
			goto tr1065
		}
		goto st0
tr26:
//line ragel/datetime.rl:7
 pb = p 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line ragel/parse_datetime.go:5254
		switch data[p] {
		case 45:
			goto tr66
		case 47:
			goto tr66
		}
		switch {
		case data[p] > 50:
			if 51 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] >= 48:
			goto st29
		}
		goto tr0
tr27:
//line ragel/datetime.rl:7
 pb = p 
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line ragel/parse_datetime.go:5279
		switch data[p] {
		case 45:
			goto tr66
		case 47:
			goto tr66
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto tr0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 112:
			goto st35
		case 117:
			goto st39
		}
		goto tr0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 114 {
			goto st36
		}
		goto tr0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 45:
			goto tr75
		case 47:
			goto tr75
		case 105:
			goto st37
		}
		goto tr0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 108 {
			goto st38
		}
		goto tr0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 45:
			goto tr75
		case 47:
			goto tr75
		}
		goto tr0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 103 {
			goto st40
		}
		goto tr0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 45:
			goto tr79
		case 47:
			goto tr79
		case 117:
			goto st41
		}
		goto tr0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 115 {
			goto st42
		}
		goto tr0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 116 {
			goto st43
		}
		goto tr0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 45:
			goto tr79
		case 47:
			goto tr79
		}
		goto tr0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		if data[p] == 101 {
			goto st45
		}
		goto tr0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if data[p] == 99 {
			goto st46
		}
		goto tr0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 45:
			goto tr85
		case 47:
			goto tr85
		case 101:
			goto st47
		}
		goto tr0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 109 {
			goto st48
		}
		goto tr0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 98 {
			goto st49
		}
		goto tr0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 101 {
			goto st50
		}
		goto tr0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		if data[p] == 114 {
			goto st51
		}
		goto tr0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 45:
			goto tr85
		case 47:
			goto tr85
		}
		goto tr0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 101 {
			goto st53
		}
		goto tr0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 98 {
			goto st54
		}
		goto tr0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 45:
			goto tr93
		case 47:
			goto tr93
		case 114:
			goto st55
		}
		goto tr0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 117 {
			goto st56
		}
		goto tr0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 97 {
			goto st57
		}
		goto tr0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 114 {
			goto st58
		}
		goto tr0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 121 {
			goto st59
		}
		goto tr0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 45:
			goto tr93
		case 47:
			goto tr93
		}
		goto tr0
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
		goto tr0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 110 {
			goto st62
		}
		goto tr0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 45:
			goto tr102
		case 47:
			goto tr102
		case 117:
			goto st63
		}
		goto tr0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 97 {
			goto st64
		}
		goto tr0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 114 {
			goto st65
		}
		goto tr0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 121 {
			goto st66
		}
		goto tr0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 45:
			goto tr102
		case 47:
			goto tr102
		}
		goto tr0
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
		goto tr0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 45:
			goto tr109
		case 47:
			goto tr109
		case 121:
			goto st69
		}
		goto tr0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 45:
			goto tr109
		case 47:
			goto tr109
		}
		goto tr0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 45:
			goto tr111
		case 47:
			goto tr111
		case 101:
			goto st71
		}
		goto tr0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 45:
			goto tr111
		case 47:
			goto tr111
		}
		goto tr0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 97 {
			goto st73
		}
		goto tr0
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
		goto tr0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 45:
			goto tr116
		case 47:
			goto tr116
		case 99:
			goto st75
		}
		goto tr0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if data[p] == 104 {
			goto st76
		}
		goto tr0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 45:
			goto tr116
		case 47:
			goto tr116
		}
		goto tr0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 45:
			goto tr119
		case 47:
			goto tr119
		}
		goto tr0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		if data[p] == 111 {
			goto st79
		}
		goto tr0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		if data[p] == 118 {
			goto st80
		}
		goto tr0
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 45:
			goto tr122
		case 47:
			goto tr122
		case 101:
			goto st81
		}
		goto tr0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		if data[p] == 109 {
			goto st82
		}
		goto tr0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		if data[p] == 98 {
			goto st83
		}
		goto tr0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		if data[p] == 101 {
			goto st84
		}
		goto tr0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 114 {
			goto st85
		}
		goto tr0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 45:
			goto tr122
		case 47:
			goto tr122
		}
		goto tr0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		if data[p] == 99 {
			goto st87
		}
		goto tr0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		if data[p] == 116 {
			goto st88
		}
		goto tr0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 45:
			goto tr130
		case 47:
			goto tr130
		case 111:
			goto st89
		}
		goto tr0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		if data[p] == 98 {
			goto st90
		}
		goto tr0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		if data[p] == 101 {
			goto st91
		}
		goto tr0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		if data[p] == 114 {
			goto st92
		}
		goto tr0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 45:
			goto tr130
		case 47:
			goto tr130
		}
		goto tr0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		if data[p] == 101 {
			goto st94
		}
		goto tr0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		if data[p] == 112 {
			goto st95
		}
		goto tr0
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 45:
			goto tr137
		case 47:
			goto tr137
		case 116:
			goto st96
		}
		goto tr0
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		if data[p] == 101 {
			goto st97
		}
		goto tr0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		if data[p] == 109 {
			goto st98
		}
		goto tr0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		if data[p] == 98 {
			goto st99
		}
		goto tr0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		if data[p] == 101 {
			goto st100
		}
		goto tr0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		if data[p] == 114 {
			goto st101
		}
		goto tr0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 45:
			goto tr137
		case 47:
			goto tr137
		}
		goto tr0
tr23:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line ragel/parse_datetime.go:6016
		if 48 <= data[p] && data[p] <= 57 {
			goto tr144
		}
		goto tr0
tr144:
//line ragel/datetime.rl:7
 pb = p 
	goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
//line ragel/parse_datetime.go:6030
		if 48 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto tr0
tr24:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line ragel/parse_datetime.go:6048
		if 48 <= data[p] && data[p] <= 57 {
			goto st105
		}
		goto tr0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		if 48 <= data[p] && data[p] <= 57 {
			goto st785
		}
		goto tr0
	st785:
		if p++; p == pe {
			goto _test_eof785
		}
	st_case_785:
		switch data[p] {
		case 32:
			goto tr945
		case 43:
			goto tr947
		case 45:
			goto tr948
		case 47:
			goto tr949
		case 84:
			goto tr950
		case 90:
			goto tr951
		case 95:
			goto tr952
		case 116:
			goto tr952
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st786
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr949
			}
		default:
			goto tr949
		}
		goto st0
	st786:
		if p++; p == pe {
			goto _test_eof786
		}
	st_case_786:
		switch data[p] {
		case 32:
			goto tr1070
		case 43:
			goto tr1071
		case 45:
			goto tr1072
		case 47:
			goto tr1073
		case 84:
			goto tr1074
		case 90:
			goto tr1075
		case 95:
			goto tr1076
		case 116:
			goto tr1076
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1073
			}
		case data[p] >= 65:
			goto tr1073
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 32:
			goto tr147
		case 45:
			goto tr148
		case 47:
			goto tr148
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto tr0
tr147:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line ragel/parse_datetime.go:6163
		switch data[p] {
		case 65:
			goto st108
		case 68:
			goto st120
		case 70:
			goto st128
		case 74:
			goto st136
		case 77:
			goto st148
		case 78:
			goto st154
		case 79:
			goto st162
		case 83:
			goto st169
		case 97:
			goto st108
		case 100:
			goto st120
		case 102:
			goto st128
		case 106:
			goto st136
		case 109:
			goto st148
		case 110:
			goto st154
		case 111:
			goto st162
		case 115:
			goto st169
		}
		goto tr0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 112:
			goto st109
		case 117:
			goto st115
		}
		goto tr0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		if data[p] == 114 {
			goto st110
		}
		goto tr0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 32:
			goto tr160
		case 105:
			goto st113
		}
		goto tr0
tr160:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st111
tr166:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st111
tr172:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st111
tr180:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st111
tr189:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st111
tr196:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st111
tr198:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st111
tr203:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st111
tr206:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st111
tr209:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st111
tr217:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st111
tr224:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line ragel/parse_datetime.go:6285
		if 48 <= data[p] && data[p] <= 57 {
			goto tr162
		}
		goto tr0
tr162:
//line ragel/datetime.rl:7
 pb = p 
	goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line ragel/parse_datetime.go:6299
		if 48 <= data[p] && data[p] <= 57 {
			goto st787
		}
		goto tr0
	st787:
		if p++; p == pe {
			goto _test_eof787
		}
	st_case_787:
		switch data[p] {
		case 32:
			goto tr1077
		case 43:
			goto tr1078
		case 45:
			goto tr1079
		case 47:
			goto tr1080
		case 84:
			goto tr1081
		case 90:
			goto tr1082
		case 95:
			goto tr1083
		case 116:
			goto tr1083
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1080
			}
		case data[p] >= 65:
			goto tr1080
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		if data[p] == 108 {
			goto st114
		}
		goto tr0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		if data[p] == 32 {
			goto tr160
		}
		goto tr0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		if data[p] == 103 {
			goto st116
		}
		goto tr0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 32:
			goto tr166
		case 117:
			goto st117
		}
		goto tr0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		if data[p] == 115 {
			goto st118
		}
		goto tr0
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		if data[p] == 116 {
			goto st119
		}
		goto tr0
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		if data[p] == 32 {
			goto tr166
		}
		goto tr0
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		if data[p] == 101 {
			goto st121
		}
		goto tr0
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		if data[p] == 99 {
			goto st122
		}
		goto tr0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 32:
			goto tr172
		case 101:
			goto st123
		}
		goto tr0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		if data[p] == 109 {
			goto st124
		}
		goto tr0
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		if data[p] == 98 {
			goto st125
		}
		goto tr0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		if data[p] == 101 {
			goto st126
		}
		goto tr0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		if data[p] == 114 {
			goto st127
		}
		goto tr0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		if data[p] == 32 {
			goto tr172
		}
		goto tr0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		if data[p] == 101 {
			goto st129
		}
		goto tr0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		if data[p] == 98 {
			goto st130
		}
		goto tr0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 32:
			goto tr180
		case 114:
			goto st131
		}
		goto tr0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		if data[p] == 117 {
			goto st132
		}
		goto tr0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		if data[p] == 97 {
			goto st133
		}
		goto tr0
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		if data[p] == 114 {
			goto st134
		}
		goto tr0
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		if data[p] == 121 {
			goto st135
		}
		goto tr0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		if data[p] == 32 {
			goto tr180
		}
		goto tr0
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 97:
			goto st137
		case 117:
			goto st143
		}
		goto tr0
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		if data[p] == 110 {
			goto st138
		}
		goto tr0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 32:
			goto tr189
		case 117:
			goto st139
		}
		goto tr0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		if data[p] == 97 {
			goto st140
		}
		goto tr0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		if data[p] == 114 {
			goto st141
		}
		goto tr0
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		if data[p] == 121 {
			goto st142
		}
		goto tr0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		if data[p] == 32 {
			goto tr189
		}
		goto tr0
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 108:
			goto st144
		case 110:
			goto st146
		}
		goto tr0
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 32:
			goto tr196
		case 121:
			goto st145
		}
		goto tr0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		if data[p] == 32 {
			goto tr196
		}
		goto tr0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 32:
			goto tr198
		case 101:
			goto st147
		}
		goto tr0
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		if data[p] == 32 {
			goto tr198
		}
		goto tr0
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		if data[p] == 97 {
			goto st149
		}
		goto tr0
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 114:
			goto st150
		case 121:
			goto st153
		}
		goto tr0
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		switch data[p] {
		case 32:
			goto tr203
		case 99:
			goto st151
		}
		goto tr0
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		if data[p] == 104 {
			goto st152
		}
		goto tr0
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		if data[p] == 32 {
			goto tr203
		}
		goto tr0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		if data[p] == 32 {
			goto tr206
		}
		goto tr0
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		if data[p] == 111 {
			goto st155
		}
		goto tr0
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		if data[p] == 118 {
			goto st156
		}
		goto tr0
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 32:
			goto tr209
		case 101:
			goto st157
		}
		goto tr0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		if data[p] == 109 {
			goto st158
		}
		goto tr0
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		if data[p] == 98 {
			goto st159
		}
		goto tr0
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		if data[p] == 101 {
			goto st160
		}
		goto tr0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		if data[p] == 114 {
			goto st161
		}
		goto tr0
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		if data[p] == 32 {
			goto tr209
		}
		goto tr0
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		if data[p] == 99 {
			goto st163
		}
		goto tr0
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		if data[p] == 116 {
			goto st164
		}
		goto tr0
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch data[p] {
		case 32:
			goto tr217
		case 111:
			goto st165
		}
		goto tr0
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		if data[p] == 98 {
			goto st166
		}
		goto tr0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		if data[p] == 101 {
			goto st167
		}
		goto tr0
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		if data[p] == 114 {
			goto st168
		}
		goto tr0
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		if data[p] == 32 {
			goto tr217
		}
		goto tr0
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		if data[p] == 101 {
			goto st170
		}
		goto tr0
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		if data[p] == 112 {
			goto st171
		}
		goto tr0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 32:
			goto tr224
		case 116:
			goto st172
		}
		goto tr0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		if data[p] == 101 {
			goto st173
		}
		goto tr0
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		if data[p] == 109 {
			goto st174
		}
		goto tr0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		if data[p] == 98 {
			goto st175
		}
		goto tr0
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		if data[p] == 101 {
			goto st176
		}
		goto tr0
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		if data[p] == 114 {
			goto st177
		}
		goto tr0
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		if data[p] == 32 {
			goto tr224
		}
		goto tr0
tr148:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st178
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
//line ragel/parse_datetime.go:6976
		switch data[p] {
		case 65:
			goto st179
		case 68:
			goto st193
		case 70:
			goto st201
		case 74:
			goto st209
		case 77:
			goto st221
		case 78:
			goto st227
		case 79:
			goto st235
		case 83:
			goto st242
		case 97:
			goto st179
		case 100:
			goto st193
		case 102:
			goto st201
		case 106:
			goto st209
		case 109:
			goto st221
		case 110:
			goto st227
		case 111:
			goto st235
		case 115:
			goto st242
		}
		goto tr0
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		switch data[p] {
		case 112:
			goto st180
		case 117:
			goto st188
		}
		goto tr0
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		if data[p] == 114 {
			goto st181
		}
		goto tr0
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 45:
			goto tr242
		case 47:
			goto tr242
		case 105:
			goto st186
		}
		goto tr0
tr242:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st182
tr250:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st182
tr256:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st182
tr264:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st182
tr273:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st182
tr280:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st182
tr282:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st182
tr287:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st182
tr290:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st182
tr293:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st182
tr301:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st182
tr308:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st182
tr414:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st182
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
//line ragel/parse_datetime.go:7111
		if 48 <= data[p] && data[p] <= 57 {
			goto tr244
		}
		goto tr0
tr244:
//line ragel/datetime.rl:7
 pb = p 
	goto st183
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
//line ragel/parse_datetime.go:7125
		if 48 <= data[p] && data[p] <= 57 {
			goto st184
		}
		goto tr0
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		if 48 <= data[p] && data[p] <= 57 {
			goto st185
		}
		goto tr0
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		if 48 <= data[p] && data[p] <= 57 {
			goto st788
		}
		goto tr0
	st788:
		if p++; p == pe {
			goto _test_eof788
		}
	st_case_788:
		switch data[p] {
		case 32:
			goto tr1084
		case 43:
			goto tr1085
		case 45:
			goto tr1086
		case 47:
			goto tr1087
		case 84:
			goto tr1088
		case 90:
			goto tr1089
		case 95:
			goto tr1090
		case 116:
			goto tr1090
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1087
			}
		case data[p] >= 65:
			goto tr1087
		}
		goto st0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		if data[p] == 108 {
			goto st187
		}
		goto tr0
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		switch data[p] {
		case 45:
			goto tr242
		case 47:
			goto tr242
		}
		goto tr0
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		if data[p] == 103 {
			goto st189
		}
		goto tr0
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 45:
			goto tr250
		case 47:
			goto tr250
		case 117:
			goto st190
		}
		goto tr0
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		if data[p] == 115 {
			goto st191
		}
		goto tr0
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		if data[p] == 116 {
			goto st192
		}
		goto tr0
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		switch data[p] {
		case 45:
			goto tr250
		case 47:
			goto tr250
		}
		goto tr0
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		if data[p] == 101 {
			goto st194
		}
		goto tr0
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		if data[p] == 99 {
			goto st195
		}
		goto tr0
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch data[p] {
		case 45:
			goto tr256
		case 47:
			goto tr256
		case 101:
			goto st196
		}
		goto tr0
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		if data[p] == 109 {
			goto st197
		}
		goto tr0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		if data[p] == 98 {
			goto st198
		}
		goto tr0
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		if data[p] == 101 {
			goto st199
		}
		goto tr0
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		if data[p] == 114 {
			goto st200
		}
		goto tr0
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		switch data[p] {
		case 45:
			goto tr256
		case 47:
			goto tr256
		}
		goto tr0
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		if data[p] == 101 {
			goto st202
		}
		goto tr0
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		if data[p] == 98 {
			goto st203
		}
		goto tr0
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		switch data[p] {
		case 45:
			goto tr264
		case 47:
			goto tr264
		case 114:
			goto st204
		}
		goto tr0
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		if data[p] == 117 {
			goto st205
		}
		goto tr0
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		if data[p] == 97 {
			goto st206
		}
		goto tr0
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		if data[p] == 114 {
			goto st207
		}
		goto tr0
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		if data[p] == 121 {
			goto st208
		}
		goto tr0
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		switch data[p] {
		case 45:
			goto tr264
		case 47:
			goto tr264
		}
		goto tr0
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		switch data[p] {
		case 97:
			goto st210
		case 117:
			goto st216
		}
		goto tr0
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		if data[p] == 110 {
			goto st211
		}
		goto tr0
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		switch data[p] {
		case 45:
			goto tr273
		case 47:
			goto tr273
		case 117:
			goto st212
		}
		goto tr0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		if data[p] == 97 {
			goto st213
		}
		goto tr0
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		if data[p] == 114 {
			goto st214
		}
		goto tr0
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		if data[p] == 121 {
			goto st215
		}
		goto tr0
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		switch data[p] {
		case 45:
			goto tr273
		case 47:
			goto tr273
		}
		goto tr0
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 108:
			goto st217
		case 110:
			goto st219
		}
		goto tr0
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		switch data[p] {
		case 45:
			goto tr280
		case 47:
			goto tr280
		case 121:
			goto st218
		}
		goto tr0
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 45:
			goto tr280
		case 47:
			goto tr280
		}
		goto tr0
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		switch data[p] {
		case 45:
			goto tr282
		case 47:
			goto tr282
		case 101:
			goto st220
		}
		goto tr0
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		switch data[p] {
		case 45:
			goto tr282
		case 47:
			goto tr282
		}
		goto tr0
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		if data[p] == 97 {
			goto st222
		}
		goto tr0
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 114:
			goto st223
		case 121:
			goto st226
		}
		goto tr0
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 45:
			goto tr287
		case 47:
			goto tr287
		case 99:
			goto st224
		}
		goto tr0
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		if data[p] == 104 {
			goto st225
		}
		goto tr0
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 45:
			goto tr287
		case 47:
			goto tr287
		}
		goto tr0
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 45:
			goto tr290
		case 47:
			goto tr290
		}
		goto tr0
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		if data[p] == 111 {
			goto st228
		}
		goto tr0
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		if data[p] == 118 {
			goto st229
		}
		goto tr0
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
		switch data[p] {
		case 45:
			goto tr293
		case 47:
			goto tr293
		case 101:
			goto st230
		}
		goto tr0
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		if data[p] == 109 {
			goto st231
		}
		goto tr0
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		if data[p] == 98 {
			goto st232
		}
		goto tr0
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		if data[p] == 101 {
			goto st233
		}
		goto tr0
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		if data[p] == 114 {
			goto st234
		}
		goto tr0
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		switch data[p] {
		case 45:
			goto tr293
		case 47:
			goto tr293
		}
		goto tr0
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		if data[p] == 99 {
			goto st236
		}
		goto tr0
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		if data[p] == 116 {
			goto st237
		}
		goto tr0
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		switch data[p] {
		case 45:
			goto tr301
		case 47:
			goto tr301
		case 111:
			goto st238
		}
		goto tr0
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		if data[p] == 98 {
			goto st239
		}
		goto tr0
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		if data[p] == 101 {
			goto st240
		}
		goto tr0
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		if data[p] == 114 {
			goto st241
		}
		goto tr0
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		switch data[p] {
		case 45:
			goto tr301
		case 47:
			goto tr301
		}
		goto tr0
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		if data[p] == 101 {
			goto st243
		}
		goto tr0
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		if data[p] == 112 {
			goto st244
		}
		goto tr0
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		switch data[p] {
		case 45:
			goto tr308
		case 47:
			goto tr308
		case 116:
			goto st245
		}
		goto tr0
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		if data[p] == 101 {
			goto st246
		}
		goto tr0
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		if data[p] == 109 {
			goto st247
		}
		goto tr0
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		if data[p] == 98 {
			goto st248
		}
		goto tr0
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		if data[p] == 101 {
			goto st249
		}
		goto tr0
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		if data[p] == 114 {
			goto st250
		}
		goto tr0
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 45:
			goto tr308
		case 47:
			goto tr308
		}
		goto tr0
tr2:
//line ragel/datetime.rl:7
 pb = p 
	goto st251
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
//line ragel/parse_datetime.go:7869
		switch data[p] {
		case 32:
			goto tr147
		case 45:
			goto tr148
		case 47:
			goto tr148
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st106
		}
		goto tr0
tr3:
//line ragel/datetime.rl:7
 pb = p 
	goto st252
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
//line ragel/parse_datetime.go:7891
		switch data[p] {
		case 32:
			goto tr147
		case 45:
			goto tr148
		case 47:
			goto tr148
		}
		switch {
		case data[p] > 49:
			if 50 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] >= 48:
			goto st106
		}
		goto tr0
tr4:
//line ragel/datetime.rl:7
 pb = p 
	goto st253
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
//line ragel/parse_datetime.go:7918
		switch data[p] {
		case 32:
			goto tr147
		case 45:
			goto tr148
		case 47:
			goto tr148
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto tr0
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		switch data[p] {
		case 112:
			goto st255
		case 117:
			goto st320
		}
		goto tr0
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		if data[p] == 114 {
			goto st256
		}
		goto tr0
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		switch data[p] {
		case 32:
			goto tr318
		case 45:
			goto tr319
		case 47:
			goto tr319
		case 105:
			goto st318
		}
		goto tr0
tr318:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st257
tr417:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st257
tr424:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st257
tr434:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st257
tr882:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st257
tr890:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st257
tr893:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st257
tr900:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st257
tr904:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st257
tr908:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st257
tr917:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st257
tr929:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st257
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
//line ragel/parse_datetime.go:8021
		switch data[p] {
		case 48:
			goto tr321
		case 51:
			goto tr323
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr322
		}
		goto tr0
tr321:
//line ragel/datetime.rl:7
 pb = p 
	goto st258
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
//line ragel/parse_datetime.go:8041
		if 49 <= data[p] && data[p] <= 57 {
			goto st259
		}
		goto tr0
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		if data[p] == 32 {
			goto tr325
		}
		goto tr0
tr325:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st260
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
//line ragel/parse_datetime.go:8071
		if data[p] == 50 {
			goto tr327
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr328
			}
		case data[p] >= 48:
			goto tr326
		}
		goto tr0
tr326:
//line ragel/datetime.rl:7
 pb = p 
	goto st261
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
//line ragel/parse_datetime.go:8093
		switch data[p] {
		case 32:
			goto tr329
		case 58:
			goto tr331
		case 65:
			goto tr332
		case 80:
			goto tr332
		case 97:
			goto tr333
		case 112:
			goto tr333
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st275
		}
		goto tr0
tr329:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st262
tr354:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st262
tr393:
//line ragel/datetime.rl:119

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

	goto st262
tr376:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st262
tr381:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st262
tr387:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st262
tr404:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st262
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
//line ragel/parse_datetime.go:8184
		switch data[p] {
		case 65:
			goto tr335
		case 80:
			goto tr335
		case 97:
			goto tr336
		case 112:
			goto tr336
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr334
		}
		goto tr0
tr334:
//line ragel/datetime.rl:7
 pb = p 
	goto st263
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
//line ragel/parse_datetime.go:8208
		if 48 <= data[p] && data[p] <= 57 {
			goto st264
		}
		goto tr0
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		if 48 <= data[p] && data[p] <= 57 {
			goto st265
		}
		goto tr0
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		if 48 <= data[p] && data[p] <= 57 {
			goto st789
		}
		goto tr0
	st789:
		if p++; p == pe {
			goto _test_eof789
		}
	st_case_789:
		if data[p] == 32 {
			goto tr1091
		}
		goto st0
tr1091:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st266
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
//line ragel/parse_datetime.go:8251
		switch data[p] {
		case 43:
			goto st267
		case 45:
			goto st268
		case 47:
			goto tr342
		case 90:
			goto tr343
		case 95:
			goto tr342
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr342
			}
		case data[p] >= 65:
			goto tr342
		}
		goto tr0
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		if data[p] == 50 {
			goto tr345
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr346
			}
		case data[p] >= 48:
			goto tr344
		}
		goto tr0
tr344:
//line ragel/datetime.rl:7
 pb = p 
	goto st790
tr347:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st790
	st790:
		if p++; p == pe {
			goto _test_eof790
		}
	st_case_790:
//line ragel/parse_datetime.go:8305
		if data[p] == 58 {
			goto tr1093
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st791
		}
		goto st0
tr346:
//line ragel/datetime.rl:7
 pb = p 
	goto st791
tr349:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st791
	st791:
		if p++; p == pe {
			goto _test_eof791
		}
	st_case_791:
//line ragel/parse_datetime.go:8328
		if data[p] == 58 {
			goto tr1093
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st792
		}
		goto st0
	st792:
		if p++; p == pe {
			goto _test_eof792
		}
	st_case_792:
		if 48 <= data[p] && data[p] <= 57 {
			goto st792
		}
		goto st0
tr1093:
//line ragel/datetime.rl:150

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st793
	st793:
		if p++; p == pe {
			goto _test_eof793
		}
	st_case_793:
//line ragel/parse_datetime.go:8361
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1096
			}
		case data[p] >= 48:
			goto tr1095
		}
		goto st0
tr1095:
//line ragel/datetime.rl:7
 pb = p 
	goto st794
	st794:
		if p++; p == pe {
			goto _test_eof794
		}
	st_case_794:
//line ragel/parse_datetime.go:8380
		if 48 <= data[p] && data[p] <= 57 {
			goto st795
		}
		goto st0
tr1096:
//line ragel/datetime.rl:7
 pb = p 
	goto st795
	st795:
		if p++; p == pe {
			goto _test_eof795
		}
	st_case_795:
//line ragel/parse_datetime.go:8394
		goto st0
tr345:
//line ragel/datetime.rl:7
 pb = p 
	goto st796
tr348:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st796
	st796:
		if p++; p == pe {
			goto _test_eof796
		}
	st_case_796:
//line ragel/parse_datetime.go:8411
		if data[p] == 58 {
			goto tr1093
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st792
			}
		case data[p] >= 48:
			goto st791
		}
		goto st0
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		if data[p] == 50 {
			goto tr348
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr349
			}
		case data[p] >= 48:
			goto tr347
		}
		goto tr0
tr342:
//line ragel/datetime.rl:7
 pb = p 
	goto st269
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
//line ragel/parse_datetime.go:8450
		switch data[p] {
		case 47:
			goto st270
		case 95:
			goto st270
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st270
			}
		case data[p] >= 65:
			goto st270
		}
		goto tr0
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		switch data[p] {
		case 47:
			goto st797
		case 95:
			goto st797
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st797
			}
		case data[p] >= 65:
			goto st797
		}
		goto tr0
	st797:
		if p++; p == pe {
			goto _test_eof797
		}
	st_case_797:
		switch data[p] {
		case 47:
			goto st797
		case 95:
			goto st797
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st797
			}
		case data[p] >= 65:
			goto st797
		}
		goto st0
tr343:
//line ragel/datetime.rl:7
 pb = p 
	goto st798
	st798:
		if p++; p == pe {
			goto _test_eof798
		}
	st_case_798:
//line ragel/parse_datetime.go:8515
		switch data[p] {
		case 47:
			goto st270
		case 95:
			goto st270
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st270
			}
		case data[p] >= 65:
			goto st270
		}
		goto st0
tr335:
//line ragel/datetime.rl:7
 pb = p 
	goto st271
tr332:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st271
tr357:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st271
tr379:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st271
tr383:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st271
tr390:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st271
tr395:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st271
tr406:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st271
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
//line ragel/parse_datetime.go:8621
		if data[p] == 77 {
			goto st272
		}
		goto tr0
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		if data[p] == 32 {
			goto tr353
		}
		goto tr0
tr353:
//line ragel/datetime.rl:56

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
tr360:
//line ragel/datetime.rl:37

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
tr364:
//line ragel/datetime.rl:119

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
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
//line ragel/parse_datetime.go:8713
		if 48 <= data[p] && data[p] <= 57 {
			goto tr334
		}
		goto tr0
tr336:
//line ragel/datetime.rl:7
 pb = p 
	goto st274
tr333:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st274
tr358:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st274
tr380:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st274
tr384:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st274
tr391:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st274
tr396:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st274
tr407:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st274
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
//line ragel/parse_datetime.go:8808
		if data[p] == 109 {
			goto st272
		}
		goto tr0
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		switch data[p] {
		case 32:
			goto tr354
		case 58:
			goto tr356
		case 65:
			goto tr357
		case 80:
			goto tr357
		case 97:
			goto tr358
		case 112:
			goto tr358
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st276
		}
		goto tr0
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		if 48 <= data[p] && data[p] <= 57 {
			goto st277
		}
		goto tr0
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		switch data[p] {
		case 32:
			goto tr360
		case 46:
			goto tr361
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st288
		}
		goto tr0
tr361:
//line ragel/datetime.rl:37

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

	goto st278
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
//line ragel/parse_datetime.go:8882
		if 48 <= data[p] && data[p] <= 57 {
			goto tr363
		}
		goto tr0
tr363:
//line ragel/datetime.rl:7
 pb = p 
	goto st279
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
//line ragel/parse_datetime.go:8896
		if data[p] == 32 {
			goto tr364
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st280
		}
		goto tr0
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		if data[p] == 32 {
			goto tr364
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st281
		}
		goto tr0
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		if data[p] == 32 {
			goto tr364
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st282
		}
		goto tr0
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		if data[p] == 32 {
			goto tr364
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st283
		}
		goto tr0
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		if data[p] == 32 {
			goto tr364
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st284
		}
		goto tr0
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		if data[p] == 32 {
			goto tr364
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st285
		}
		goto tr0
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		if data[p] == 32 {
			goto tr364
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st286
		}
		goto tr0
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		if data[p] == 32 {
			goto tr364
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st287
		}
		goto tr0
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		if data[p] == 32 {
			goto tr364
		}
		goto tr0
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		if 48 <= data[p] && data[p] <= 57 {
			goto st289
		}
		goto tr0
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		switch data[p] {
		case 32:
			goto tr360
		case 46:
			goto tr361
		}
		goto tr0
tr331:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st290
tr356:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st290
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
//line ragel/parse_datetime.go:9035
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr375
			}
		case data[p] >= 48:
			goto tr374
		}
		goto tr0
tr374:
//line ragel/datetime.rl:7
 pb = p 
	goto st291
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
//line ragel/parse_datetime.go:9054
		switch data[p] {
		case 32:
			goto tr376
		case 58:
			goto tr378
		case 65:
			goto tr379
		case 80:
			goto tr379
		case 97:
			goto tr380
		case 112:
			goto tr380
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st292
		}
		goto tr0
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		switch data[p] {
		case 32:
			goto tr381
		case 58:
			goto tr382
		case 65:
			goto tr383
		case 80:
			goto tr383
		case 97:
			goto tr384
		case 112:
			goto tr384
		}
		goto tr0
tr378:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st293
tr382:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st293
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
//line ragel/parse_datetime.go:9110
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr386
			}
		case data[p] >= 48:
			goto tr385
		}
		goto tr0
tr385:
//line ragel/datetime.rl:7
 pb = p 
	goto st294
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
//line ragel/parse_datetime.go:9129
		switch data[p] {
		case 32:
			goto tr387
		case 46:
			goto tr388
		case 65:
			goto tr390
		case 80:
			goto tr390
		case 97:
			goto tr391
		case 112:
			goto tr391
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st305
		}
		goto tr0
tr388:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st295
tr405:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st295
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
//line ragel/parse_datetime.go:9165
		if 48 <= data[p] && data[p] <= 57 {
			goto tr392
		}
		goto tr0
tr392:
//line ragel/datetime.rl:7
 pb = p 
	goto st296
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
//line ragel/parse_datetime.go:9179
		switch data[p] {
		case 32:
			goto tr393
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
			goto st297
		}
		goto tr0
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		switch data[p] {
		case 32:
			goto tr393
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
			goto st298
		}
		goto tr0
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		switch data[p] {
		case 32:
			goto tr393
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
			goto st299
		}
		goto tr0
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 32:
			goto tr393
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
			goto st300
		}
		goto tr0
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		switch data[p] {
		case 32:
			goto tr393
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
			goto st301
		}
		goto tr0
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		switch data[p] {
		case 32:
			goto tr393
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
			goto st302
		}
		goto tr0
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		switch data[p] {
		case 32:
			goto tr393
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
			goto st303
		}
		goto tr0
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		switch data[p] {
		case 32:
			goto tr393
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
			goto st304
		}
		goto tr0
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		switch data[p] {
		case 32:
			goto tr393
		case 65:
			goto tr395
		case 80:
			goto tr395
		case 97:
			goto tr396
		case 112:
			goto tr396
		}
		goto tr0
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		switch data[p] {
		case 32:
			goto tr404
		case 46:
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
		goto tr0
tr386:
//line ragel/datetime.rl:7
 pb = p 
	goto st306
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
//line ragel/parse_datetime.go:9390
		switch data[p] {
		case 32:
			goto tr387
		case 46:
			goto tr388
		case 65:
			goto tr390
		case 80:
			goto tr390
		case 97:
			goto tr391
		case 112:
			goto tr391
		}
		goto tr0
tr375:
//line ragel/datetime.rl:7
 pb = p 
	goto st307
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
//line ragel/parse_datetime.go:9415
		switch data[p] {
		case 32:
			goto tr376
		case 58:
			goto tr378
		case 65:
			goto tr379
		case 80:
			goto tr379
		case 97:
			goto tr380
		case 112:
			goto tr380
		}
		goto tr0
tr327:
//line ragel/datetime.rl:7
 pb = p 
	goto st308
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
//line ragel/parse_datetime.go:9440
		switch data[p] {
		case 32:
			goto tr329
		case 58:
			goto tr331
		case 65:
			goto tr332
		case 80:
			goto tr332
		case 97:
			goto tr333
		case 112:
			goto tr333
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st309
			}
		case data[p] >= 48:
			goto st275
		}
		goto tr0
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		if 48 <= data[p] && data[p] <= 57 {
			goto st276
		}
		goto tr0
tr328:
//line ragel/datetime.rl:7
 pb = p 
	goto st310
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
//line ragel/parse_datetime.go:9482
		switch data[p] {
		case 32:
			goto tr329
		case 58:
			goto tr331
		case 65:
			goto tr332
		case 80:
			goto tr332
		case 97:
			goto tr333
		case 112:
			goto tr333
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st309
		}
		goto tr0
tr322:
//line ragel/datetime.rl:7
 pb = p 
	goto st311
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
//line ragel/parse_datetime.go:9510
		if 48 <= data[p] && data[p] <= 57 {
			goto st259
		}
		goto tr0
tr323:
//line ragel/datetime.rl:7
 pb = p 
	goto st312
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
//line ragel/parse_datetime.go:9524
		if 48 <= data[p] && data[p] <= 49 {
			goto st259
		}
		goto tr0
tr319:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st313
tr418:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st313
tr425:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st313
tr435:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st313
tr883:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st313
tr891:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st313
tr894:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st313
tr901:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st313
tr905:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st313
tr909:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st313
tr918:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st313
tr930:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st313
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
//line ragel/parse_datetime.go:9582
		switch data[p] {
		case 48:
			goto tr409
		case 51:
			goto tr411
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr412
			}
		case data[p] >= 49:
			goto tr410
		}
		goto tr0
tr409:
//line ragel/datetime.rl:7
 pb = p 
	goto st314
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
//line ragel/parse_datetime.go:9607
		if 49 <= data[p] && data[p] <= 57 {
			goto st315
		}
		goto tr0
tr412:
//line ragel/datetime.rl:7
 pb = p 
	goto st315
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
//line ragel/parse_datetime.go:9621
		switch data[p] {
		case 45:
			goto tr414
		case 47:
			goto tr414
		}
		goto tr0
tr410:
//line ragel/datetime.rl:7
 pb = p 
	goto st316
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
//line ragel/parse_datetime.go:9638
		switch data[p] {
		case 45:
			goto tr414
		case 47:
			goto tr414
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st315
		}
		goto tr0
tr411:
//line ragel/datetime.rl:7
 pb = p 
	goto st317
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
//line ragel/parse_datetime.go:9658
		switch data[p] {
		case 45:
			goto tr414
		case 47:
			goto tr414
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st315
		}
		goto tr0
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		if data[p] == 108 {
			goto st319
		}
		goto tr0
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		switch data[p] {
		case 32:
			goto tr318
		case 45:
			goto tr319
		case 47:
			goto tr319
		}
		goto tr0
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		if data[p] == 103 {
			goto st321
		}
		goto tr0
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		switch data[p] {
		case 32:
			goto tr417
		case 45:
			goto tr418
		case 47:
			goto tr418
		case 117:
			goto st322
		}
		goto tr0
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		if data[p] == 115 {
			goto st323
		}
		goto tr0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		if data[p] == 116 {
			goto st324
		}
		goto tr0
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		switch data[p] {
		case 32:
			goto tr417
		case 45:
			goto tr418
		case 47:
			goto tr418
		}
		goto tr0
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
		if data[p] == 101 {
			goto st326
		}
		goto tr0
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		if data[p] == 99 {
			goto st327
		}
		goto tr0
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		switch data[p] {
		case 32:
			goto tr424
		case 45:
			goto tr425
		case 47:
			goto tr425
		case 101:
			goto st328
		}
		goto tr0
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		if data[p] == 109 {
			goto st329
		}
		goto tr0
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		if data[p] == 98 {
			goto st330
		}
		goto tr0
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		if data[p] == 101 {
			goto st331
		}
		goto tr0
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		if data[p] == 114 {
			goto st332
		}
		goto tr0
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		switch data[p] {
		case 32:
			goto tr424
		case 45:
			goto tr425
		case 47:
			goto tr425
		}
		goto tr0
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 101:
			goto st334
		case 114:
			goto st341
		}
		goto tr0
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		if data[p] == 98 {
			goto st335
		}
		goto tr0
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		switch data[p] {
		case 32:
			goto tr434
		case 45:
			goto tr435
		case 47:
			goto tr435
		case 114:
			goto st336
		}
		goto tr0
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		if data[p] == 117 {
			goto st337
		}
		goto tr0
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		if data[p] == 97 {
			goto st338
		}
		goto tr0
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		if data[p] == 114 {
			goto st339
		}
		goto tr0
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		if data[p] == 121 {
			goto st340
		}
		goto tr0
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		switch data[p] {
		case 32:
			goto tr434
		case 45:
			goto tr435
		case 47:
			goto tr435
		}
		goto tr0
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		if data[p] == 105 {
			goto st342
		}
		goto tr0
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		switch data[p] {
		case 32:
			goto st343
		case 44:
			goto st525
		case 100:
			goto st669
		}
		goto tr0
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
		switch data[p] {
		case 65:
			goto st344
		case 68:
			goto st467
		case 70:
			goto st475
		case 74:
			goto st483
		case 77:
			goto st495
		case 78:
			goto st501
		case 79:
			goto st509
		case 83:
			goto st516
		case 97:
			goto st344
		case 100:
			goto st467
		case 102:
			goto st475
		case 106:
			goto st483
		case 109:
			goto st495
		case 110:
			goto st501
		case 111:
			goto st509
		case 115:
			goto st516
		}
		goto tr0
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		switch data[p] {
		case 112:
			goto st345
		case 117:
			goto st462
		}
		goto tr0
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		if data[p] == 114 {
			goto st346
		}
		goto tr0
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		switch data[p] {
		case 32:
			goto tr456
		case 105:
			goto st460
		}
		goto tr0
tr456:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st347
tr668:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st347
tr674:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st347
tr682:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st347
tr691:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st347
tr698:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st347
tr700:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st347
tr705:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st347
tr708:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st347
tr711:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st347
tr719:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st347
tr726:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st347
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
//line ragel/parse_datetime.go:10069
		switch data[p] {
		case 32:
			goto st348
		case 48:
			goto tr459
		case 51:
			goto tr461
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr462
			}
		case data[p] >= 49:
			goto tr460
		}
		goto tr0
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		switch data[p] {
		case 48:
			goto tr463
		case 51:
			goto tr465
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr462
			}
		case data[p] >= 49:
			goto tr464
		}
		goto tr0
tr463:
//line ragel/datetime.rl:7
 pb = p 
	goto st349
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
//line ragel/parse_datetime.go:10116
		if 49 <= data[p] && data[p] <= 57 {
			goto st350
		}
		goto tr0
tr462:
//line ragel/datetime.rl:7
 pb = p 
	goto st350
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
//line ragel/parse_datetime.go:10130
		if data[p] == 32 {
			goto tr467
		}
		goto tr0
tr467:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st351
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
//line ragel/parse_datetime.go:10151
		if data[p] == 50 {
			goto tr469
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr470
			}
		case data[p] >= 48:
			goto tr468
		}
		goto tr0
tr468:
//line ragel/datetime.rl:7
 pb = p 
	goto st352
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
//line ragel/parse_datetime.go:10173
		switch data[p] {
		case 32:
			goto tr471
		case 43:
			goto tr472
		case 45:
			goto tr473
		case 47:
			goto tr474
		case 58:
			goto tr476
		case 65:
			goto tr477
		case 80:
			goto tr477
		case 90:
			goto tr478
		case 95:
			goto tr474
		case 97:
			goto tr479
		case 112:
			goto tr479
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st375
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr474
			}
		default:
			goto tr474
		}
		goto tr0
tr471:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st353
tr514:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st353
tr577:
//line ragel/datetime.rl:119

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

	goto st353
tr548:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st353
tr557:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st353
tr567:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st353
tr588:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st353
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
//line ragel/parse_datetime.go:10283
		switch data[p] {
		case 32:
			goto st354
		case 43:
			goto st358
		case 45:
			goto st366
		case 47:
			goto tr483
		case 65:
			goto tr485
		case 80:
			goto tr485
		case 90:
			goto tr486
		case 95:
			goto tr483
		case 97:
			goto tr487
		case 112:
			goto tr487
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr484
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr483
			}
		default:
			goto tr483
		}
		goto tr0
tr494:
//line ragel/datetime.rl:168
 
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

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st354
tr498:
//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st354
tr501:
//line ragel/datetime.rl:159

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st354
tr508:
//line ragel/datetime.rl:196

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st354
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
//line ragel/parse_datetime.go:10382
		if 48 <= data[p] && data[p] <= 57 {
			goto tr484
		}
		goto tr0
tr484:
//line ragel/datetime.rl:7
 pb = p 
	goto st355
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
//line ragel/parse_datetime.go:10396
		if 48 <= data[p] && data[p] <= 57 {
			goto st356
		}
		goto tr0
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		if 48 <= data[p] && data[p] <= 57 {
			goto st357
		}
		goto tr0
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		if 48 <= data[p] && data[p] <= 57 {
			goto st799
		}
		goto tr0
	st799:
		if p++; p == pe {
			goto _test_eof799
		}
	st_case_799:
		goto st0
tr472:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st358
tr511:
//line ragel/datetime.rl:56

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

	goto st358
tr515:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st358
tr525:
//line ragel/datetime.rl:37

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

	goto st358
tr533:
//line ragel/datetime.rl:119

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

	goto st358
tr549:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st358
tr558:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st358
tr568:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st358
tr589:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st358
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
//line ragel/parse_datetime.go:10539
		if data[p] == 50 {
			goto tr492
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr493
			}
		case data[p] >= 48:
			goto tr491
		}
		goto tr0
tr491:
//line ragel/datetime.rl:7
 pb = p 
	goto st359
tr503:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st359
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
//line ragel/parse_datetime.go:10567
		switch data[p] {
		case 32:
			goto tr494
		case 58:
			goto tr496
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st360
		}
		goto tr0
tr493:
//line ragel/datetime.rl:7
 pb = p 
	goto st360
tr505:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st360
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
//line ragel/parse_datetime.go:10593
		switch data[p] {
		case 32:
			goto tr494
		case 58:
			goto tr496
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st361
		}
		goto tr0
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		if data[p] == 32 {
			goto tr494
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st361
		}
		goto tr0
tr496:
//line ragel/datetime.rl:150

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st362
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
//line ragel/parse_datetime.go:10632
		if data[p] == 32 {
			goto tr498
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr500
			}
		case data[p] >= 48:
			goto tr499
		}
		goto tr0
tr499:
//line ragel/datetime.rl:7
 pb = p 
	goto st363
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
//line ragel/parse_datetime.go:10654
		if data[p] == 32 {
			goto tr501
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st364
		}
		goto tr0
tr500:
//line ragel/datetime.rl:7
 pb = p 
	goto st364
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
//line ragel/parse_datetime.go:10671
		if data[p] == 32 {
			goto tr501
		}
		goto tr0
tr492:
//line ragel/datetime.rl:7
 pb = p 
	goto st365
tr504:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st365
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
//line ragel/parse_datetime.go:10691
		switch data[p] {
		case 32:
			goto tr494
		case 58:
			goto tr496
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st361
			}
		case data[p] >= 48:
			goto st360
		}
		goto tr0
tr473:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st366
tr512:
//line ragel/datetime.rl:56

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

	goto st366
tr516:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st366
tr526:
//line ragel/datetime.rl:37

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

	goto st366
tr534:
//line ragel/datetime.rl:119

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

	goto st366
tr550:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st366
tr559:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st366
tr569:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st366
tr590:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st366
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
//line ragel/parse_datetime.go:10821
		if data[p] == 50 {
			goto tr504
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr505
			}
		case data[p] >= 48:
			goto tr503
		}
		goto tr0
tr483:
//line ragel/datetime.rl:7
 pb = p 
	goto st367
tr474:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st367
tr517:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st367
tr551:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st367
tr560:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st367
tr571:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st367
tr535:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st367
tr592:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st367
tr528:
//line ragel/datetime.rl:37

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

//line ragel/datetime.rl:7
 pb = p 
	goto st367
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
//line ragel/parse_datetime.go:10943
		switch data[p] {
		case 47:
			goto st368
		case 95:
			goto st368
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st368
			}
		case data[p] >= 65:
			goto st368
		}
		goto tr0
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		switch data[p] {
		case 47:
			goto st369
		case 95:
			goto st369
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st369
			}
		case data[p] >= 65:
			goto st369
		}
		goto tr0
tr513:
//line ragel/datetime.rl:56

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

//line ragel/datetime.rl:7
 pb = p 
	goto st369
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
//line ragel/parse_datetime.go:11011
		switch data[p] {
		case 32:
			goto tr508
		case 47:
			goto st369
		case 95:
			goto st369
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st369
			}
		case data[p] >= 65:
			goto st369
		}
		goto tr0
tr485:
//line ragel/datetime.rl:7
 pb = p 
	goto st370
tr477:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st370
tr520:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st370
tr554:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st370
tr562:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st370
tr573:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st370
tr579:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st370
tr593:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st370
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
//line ragel/parse_datetime.go:11119
		switch data[p] {
		case 47:
			goto st368
		case 77:
			goto st371
		case 95:
			goto st368
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st368
			}
		case data[p] >= 65:
			goto st368
		}
		goto tr0
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		switch data[p] {
		case 32:
			goto tr510
		case 43:
			goto tr511
		case 45:
			goto tr512
		case 47:
			goto tr513
		case 95:
			goto tr513
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr513
			}
		case data[p] >= 65:
			goto tr513
		}
		goto tr0
tr510:
//line ragel/datetime.rl:56

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

	goto st372
tr524:
//line ragel/datetime.rl:37

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

	goto st372
tr532:
//line ragel/datetime.rl:119

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

	goto st372
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
//line ragel/parse_datetime.go:11241
		switch data[p] {
		case 32:
			goto st354
		case 43:
			goto st358
		case 45:
			goto st366
		case 47:
			goto tr483
		case 90:
			goto tr486
		case 95:
			goto tr483
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr484
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr483
			}
		default:
			goto tr483
		}
		goto tr0
tr486:
//line ragel/datetime.rl:7
 pb = p 
	goto st373
tr478:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st373
tr521:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st373
tr555:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st373
tr563:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st373
tr574:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st373
tr537:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st373
tr594:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st373
tr530:
//line ragel/datetime.rl:37

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

//line ragel/datetime.rl:7
 pb = p 
	goto st373
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
//line ragel/parse_datetime.go:11378
		switch data[p] {
		case 32:
			goto tr498
		case 47:
			goto st368
		case 95:
			goto st368
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st368
			}
		case data[p] >= 65:
			goto st368
		}
		goto tr0
tr487:
//line ragel/datetime.rl:7
 pb = p 
	goto st374
tr479:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st374
tr522:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st374
tr556:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st374
tr564:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st374
tr575:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st374
tr580:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st374
tr595:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st374
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
//line ragel/parse_datetime.go:11486
		switch data[p] {
		case 47:
			goto st368
		case 95:
			goto st368
		case 109:
			goto st371
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st368
			}
		case data[p] >= 65:
			goto st368
		}
		goto tr0
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		switch data[p] {
		case 32:
			goto tr514
		case 43:
			goto tr515
		case 45:
			goto tr516
		case 47:
			goto tr517
		case 58:
			goto tr519
		case 65:
			goto tr520
		case 80:
			goto tr520
		case 90:
			goto tr521
		case 95:
			goto tr517
		case 97:
			goto tr522
		case 112:
			goto tr522
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st376
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr517
			}
		default:
			goto tr517
		}
		goto tr0
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		if 48 <= data[p] && data[p] <= 57 {
			goto st377
		}
		goto tr0
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		switch data[p] {
		case 32:
			goto tr524
		case 43:
			goto tr525
		case 45:
			goto tr526
		case 46:
			goto tr527
		case 47:
			goto tr528
		case 90:
			goto tr530
		case 95:
			goto tr528
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st388
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr528
			}
		default:
			goto tr528
		}
		goto tr0
tr527:
//line ragel/datetime.rl:37

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

	goto st378
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
//line ragel/parse_datetime.go:11611
		if 48 <= data[p] && data[p] <= 57 {
			goto tr531
		}
		goto tr0
tr531:
//line ragel/datetime.rl:7
 pb = p 
	goto st379
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
//line ragel/parse_datetime.go:11625
		switch data[p] {
		case 32:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st380
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		switch data[p] {
		case 32:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st381
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		switch data[p] {
		case 32:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st382
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		switch data[p] {
		case 32:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st383
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 32:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st384
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		switch data[p] {
		case 32:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st385
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		switch data[p] {
		case 32:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st386
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		switch data[p] {
		case 32:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st387
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		switch data[p] {
		case 32:
			goto tr532
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		case data[p] >= 65:
			goto tr535
		}
		goto tr0
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
		if 48 <= data[p] && data[p] <= 57 {
			goto st389
		}
		goto tr0
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 32:
			goto tr524
		case 43:
			goto tr525
		case 45:
			goto tr526
		case 46:
			goto tr527
		case 47:
			goto tr528
		case 90:
			goto tr530
		case 95:
			goto tr528
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr528
			}
		case data[p] >= 65:
			goto tr528
		}
		goto tr0
tr476:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st390
tr519:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st390
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
//line ragel/parse_datetime.go:11961
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr547
			}
		case data[p] >= 48:
			goto tr546
		}
		goto tr0
tr546:
//line ragel/datetime.rl:7
 pb = p 
	goto st391
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
//line ragel/parse_datetime.go:11980
		switch data[p] {
		case 32:
			goto tr548
		case 43:
			goto tr549
		case 45:
			goto tr550
		case 47:
			goto tr551
		case 58:
			goto tr553
		case 65:
			goto tr554
		case 80:
			goto tr554
		case 90:
			goto tr555
		case 95:
			goto tr551
		case 97:
			goto tr556
		case 112:
			goto tr556
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st392
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr551
			}
		default:
			goto tr551
		}
		goto tr0
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		switch data[p] {
		case 32:
			goto tr557
		case 43:
			goto tr558
		case 45:
			goto tr559
		case 47:
			goto tr560
		case 58:
			goto tr561
		case 65:
			goto tr562
		case 80:
			goto tr562
		case 90:
			goto tr563
		case 95:
			goto tr560
		case 97:
			goto tr564
		case 112:
			goto tr564
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr560
			}
		case data[p] >= 66:
			goto tr560
		}
		goto tr0
tr553:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st393
tr561:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st393
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
//line ragel/parse_datetime.go:12073
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr566
			}
		case data[p] >= 48:
			goto tr565
		}
		goto tr0
tr565:
//line ragel/datetime.rl:7
 pb = p 
	goto st394
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
//line ragel/parse_datetime.go:12092
		switch data[p] {
		case 32:
			goto tr567
		case 43:
			goto tr568
		case 45:
			goto tr569
		case 46:
			goto tr570
		case 47:
			goto tr571
		case 65:
			goto tr573
		case 80:
			goto tr573
		case 90:
			goto tr574
		case 95:
			goto tr571
		case 97:
			goto tr575
		case 112:
			goto tr575
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st405
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr571
			}
		default:
			goto tr571
		}
		goto tr0
tr570:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st395
tr591:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st395
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
//line ragel/parse_datetime.go:12147
		if 48 <= data[p] && data[p] <= 57 {
			goto tr576
		}
		goto tr0
tr576:
//line ragel/datetime.rl:7
 pb = p 
	goto st396
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
//line ragel/parse_datetime.go:12161
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st397
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st398
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st399
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st400
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st401
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st402
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st403
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st404
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		case data[p] >= 66:
			goto tr535
		}
		goto tr0
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 32:
			goto tr588
		case 43:
			goto tr589
		case 45:
			goto tr590
		case 46:
			goto tr591
		case 47:
			goto tr592
		case 65:
			goto tr593
		case 80:
			goto tr593
		case 90:
			goto tr594
		case 95:
			goto tr592
		case 97:
			goto tr595
		case 112:
			goto tr595
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr592
			}
		case data[p] >= 66:
			goto tr592
		}
		goto tr0
tr566:
//line ragel/datetime.rl:7
 pb = p 
	goto st406
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
//line ragel/parse_datetime.go:12560
		switch data[p] {
		case 32:
			goto tr567
		case 43:
			goto tr568
		case 45:
			goto tr569
		case 46:
			goto tr570
		case 47:
			goto tr571
		case 65:
			goto tr573
		case 80:
			goto tr573
		case 90:
			goto tr574
		case 95:
			goto tr571
		case 97:
			goto tr575
		case 112:
			goto tr575
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr571
			}
		case data[p] >= 66:
			goto tr571
		}
		goto tr0
tr547:
//line ragel/datetime.rl:7
 pb = p 
	goto st407
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
//line ragel/parse_datetime.go:12603
		switch data[p] {
		case 32:
			goto tr548
		case 43:
			goto tr549
		case 45:
			goto tr550
		case 47:
			goto tr551
		case 58:
			goto tr553
		case 65:
			goto tr554
		case 80:
			goto tr554
		case 90:
			goto tr555
		case 95:
			goto tr551
		case 97:
			goto tr556
		case 112:
			goto tr556
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr551
			}
		case data[p] >= 66:
			goto tr551
		}
		goto tr0
tr469:
//line ragel/datetime.rl:7
 pb = p 
	goto st408
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
//line ragel/parse_datetime.go:12646
		switch data[p] {
		case 32:
			goto tr471
		case 43:
			goto tr472
		case 45:
			goto tr473
		case 47:
			goto tr474
		case 58:
			goto tr476
		case 65:
			goto tr477
		case 80:
			goto tr477
		case 90:
			goto tr478
		case 95:
			goto tr474
		case 97:
			goto tr479
		case 112:
			goto tr479
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st375
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr474
				}
			case data[p] >= 66:
				goto tr474
			}
		default:
			goto st409
		}
		goto tr0
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		if 48 <= data[p] && data[p] <= 57 {
			goto st376
		}
		goto tr0
tr470:
//line ragel/datetime.rl:7
 pb = p 
	goto st410
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
//line ragel/parse_datetime.go:12707
		switch data[p] {
		case 32:
			goto tr471
		case 43:
			goto tr472
		case 45:
			goto tr473
		case 47:
			goto tr474
		case 58:
			goto tr476
		case 65:
			goto tr477
		case 80:
			goto tr477
		case 90:
			goto tr478
		case 95:
			goto tr474
		case 97:
			goto tr479
		case 112:
			goto tr479
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st409
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr474
			}
		default:
			goto tr474
		}
		goto tr0
tr464:
//line ragel/datetime.rl:7
 pb = p 
	goto st411
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
//line ragel/parse_datetime.go:12754
		if data[p] == 32 {
			goto tr467
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st350
		}
		goto tr0
tr465:
//line ragel/datetime.rl:7
 pb = p 
	goto st412
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
//line ragel/parse_datetime.go:12771
		if data[p] == 32 {
			goto tr467
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st350
		}
		goto tr0
tr459:
//line ragel/datetime.rl:7
 pb = p 
	goto st413
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
//line ragel/parse_datetime.go:12788
		if 49 <= data[p] && data[p] <= 57 {
			goto st414
		}
		goto tr0
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		if data[p] == 32 {
			goto tr598
		}
		goto tr0
tr598:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st415
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
//line ragel/parse_datetime.go:12818
		if data[p] == 50 {
			goto tr600
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr601
			}
		case data[p] >= 48:
			goto tr599
		}
		goto tr0
tr599:
//line ragel/datetime.rl:7
 pb = p 
	goto st416
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
//line ragel/parse_datetime.go:12840
		switch data[p] {
		case 32:
			goto tr602
		case 43:
			goto tr472
		case 45:
			goto tr473
		case 47:
			goto tr474
		case 58:
			goto tr604
		case 65:
			goto tr605
		case 80:
			goto tr605
		case 90:
			goto tr478
		case 95:
			goto tr474
		case 97:
			goto tr606
		case 112:
			goto tr606
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st422
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr474
			}
		default:
			goto tr474
		}
		goto tr0
tr602:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st417
tr611:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st417
tr650:
//line ragel/datetime.rl:119

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

	goto st417
tr633:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st417
tr638:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st417
tr644:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st417
tr661:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st417
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
//line ragel/parse_datetime.go:12950
		switch data[p] {
		case 32:
			goto st354
		case 43:
			goto st358
		case 45:
			goto st366
		case 47:
			goto tr483
		case 65:
			goto tr607
		case 80:
			goto tr607
		case 90:
			goto tr486
		case 95:
			goto tr483
		case 97:
			goto tr608
		case 112:
			goto tr608
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr334
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr483
			}
		default:
			goto tr483
		}
		goto tr0
tr607:
//line ragel/datetime.rl:7
 pb = p 
	goto st418
tr605:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st418
tr614:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st418
tr636:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st418
tr640:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st418
tr647:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st418
tr652:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st418
tr663:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st418
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
//line ragel/parse_datetime.go:13076
		switch data[p] {
		case 47:
			goto st368
		case 77:
			goto st419
		case 95:
			goto st368
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st368
			}
		case data[p] >= 65:
			goto st368
		}
		goto tr0
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		switch data[p] {
		case 32:
			goto tr610
		case 43:
			goto tr511
		case 45:
			goto tr512
		case 47:
			goto tr513
		case 95:
			goto tr513
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr513
			}
		case data[p] >= 65:
			goto tr513
		}
		goto tr0
tr610:
//line ragel/datetime.rl:56

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

	goto st420
tr617:
//line ragel/datetime.rl:37

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

	goto st420
tr621:
//line ragel/datetime.rl:119

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

	goto st420
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
//line ragel/parse_datetime.go:13198
		switch data[p] {
		case 32:
			goto st354
		case 43:
			goto st358
		case 45:
			goto st366
		case 47:
			goto tr483
		case 90:
			goto tr486
		case 95:
			goto tr483
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr334
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr483
			}
		default:
			goto tr483
		}
		goto tr0
tr608:
//line ragel/datetime.rl:7
 pb = p 
	goto st421
tr606:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st421
tr615:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st421
tr637:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st421
tr641:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st421
tr648:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st421
tr653:
//line ragel/datetime.rl:119

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

//line ragel/datetime.rl:7
 pb = p 
	goto st421
tr664:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st421
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
//line ragel/parse_datetime.go:13316
		switch data[p] {
		case 47:
			goto st368
		case 95:
			goto st368
		case 109:
			goto st419
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st368
			}
		case data[p] >= 65:
			goto st368
		}
		goto tr0
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		switch data[p] {
		case 32:
			goto tr611
		case 43:
			goto tr515
		case 45:
			goto tr516
		case 47:
			goto tr517
		case 58:
			goto tr613
		case 65:
			goto tr614
		case 80:
			goto tr614
		case 90:
			goto tr521
		case 95:
			goto tr517
		case 97:
			goto tr615
		case 112:
			goto tr615
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st423
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr517
			}
		default:
			goto tr517
		}
		goto tr0
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		if 48 <= data[p] && data[p] <= 57 {
			goto st424
		}
		goto tr0
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 32:
			goto tr617
		case 43:
			goto tr525
		case 45:
			goto tr526
		case 46:
			goto tr618
		case 47:
			goto tr528
		case 90:
			goto tr530
		case 95:
			goto tr528
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st435
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr528
			}
		default:
			goto tr528
		}
		goto tr0
tr618:
//line ragel/datetime.rl:37

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

	goto st425
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
//line ragel/parse_datetime.go:13441
		if 48 <= data[p] && data[p] <= 57 {
			goto tr620
		}
		goto tr0
tr620:
//line ragel/datetime.rl:7
 pb = p 
	goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
//line ragel/parse_datetime.go:13455
		switch data[p] {
		case 32:
			goto tr621
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st427
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 32:
			goto tr621
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st428
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		switch data[p] {
		case 32:
			goto tr621
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st429
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
		switch data[p] {
		case 32:
			goto tr621
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st430
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
		switch data[p] {
		case 32:
			goto tr621
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st431
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		switch data[p] {
		case 32:
			goto tr621
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st432
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
		switch data[p] {
		case 32:
			goto tr621
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st433
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 32:
			goto tr621
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st434
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		switch data[p] {
		case 32:
			goto tr621
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 90:
			goto tr537
		case 95:
			goto tr535
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		case data[p] >= 65:
			goto tr535
		}
		goto tr0
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		if 48 <= data[p] && data[p] <= 57 {
			goto st436
		}
		goto tr0
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		switch data[p] {
		case 32:
			goto tr617
		case 43:
			goto tr525
		case 45:
			goto tr526
		case 46:
			goto tr618
		case 47:
			goto tr528
		case 90:
			goto tr530
		case 95:
			goto tr528
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr528
			}
		case data[p] >= 65:
			goto tr528
		}
		goto tr0
tr604:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st437
tr613:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st437
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
//line ragel/parse_datetime.go:13791
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr632
			}
		case data[p] >= 48:
			goto tr631
		}
		goto tr0
tr631:
//line ragel/datetime.rl:7
 pb = p 
	goto st438
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
//line ragel/parse_datetime.go:13810
		switch data[p] {
		case 32:
			goto tr633
		case 43:
			goto tr549
		case 45:
			goto tr550
		case 47:
			goto tr551
		case 58:
			goto tr635
		case 65:
			goto tr636
		case 80:
			goto tr636
		case 90:
			goto tr555
		case 95:
			goto tr551
		case 97:
			goto tr637
		case 112:
			goto tr637
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st439
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr551
			}
		default:
			goto tr551
		}
		goto tr0
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		switch data[p] {
		case 32:
			goto tr638
		case 43:
			goto tr558
		case 45:
			goto tr559
		case 47:
			goto tr560
		case 58:
			goto tr639
		case 65:
			goto tr640
		case 80:
			goto tr640
		case 90:
			goto tr563
		case 95:
			goto tr560
		case 97:
			goto tr641
		case 112:
			goto tr641
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr560
			}
		case data[p] >= 66:
			goto tr560
		}
		goto tr0
tr635:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st440
tr639:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st440
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
//line ragel/parse_datetime.go:13903
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr643
			}
		case data[p] >= 48:
			goto tr642
		}
		goto tr0
tr642:
//line ragel/datetime.rl:7
 pb = p 
	goto st441
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
//line ragel/parse_datetime.go:13922
		switch data[p] {
		case 32:
			goto tr644
		case 43:
			goto tr568
		case 45:
			goto tr569
		case 46:
			goto tr645
		case 47:
			goto tr571
		case 65:
			goto tr647
		case 80:
			goto tr647
		case 90:
			goto tr574
		case 95:
			goto tr571
		case 97:
			goto tr648
		case 112:
			goto tr648
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st452
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr571
			}
		default:
			goto tr571
		}
		goto tr0
tr645:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st442
tr662:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st442
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
//line ragel/parse_datetime.go:13977
		if 48 <= data[p] && data[p] <= 57 {
			goto tr649
		}
		goto tr0
tr649:
//line ragel/datetime.rl:7
 pb = p 
	goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
//line ragel/parse_datetime.go:13991
		switch data[p] {
		case 32:
			goto tr650
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr653
		case 112:
			goto tr653
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st444
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		switch data[p] {
		case 32:
			goto tr650
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr653
		case 112:
			goto tr653
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st445
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
		switch data[p] {
		case 32:
			goto tr650
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr653
		case 112:
			goto tr653
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st446
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		switch data[p] {
		case 32:
			goto tr650
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr653
		case 112:
			goto tr653
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st447
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
		switch data[p] {
		case 32:
			goto tr650
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr653
		case 112:
			goto tr653
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st448
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		switch data[p] {
		case 32:
			goto tr650
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr653
		case 112:
			goto tr653
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st449
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		switch data[p] {
		case 32:
			goto tr650
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr653
		case 112:
			goto tr653
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st450
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
		switch data[p] {
		case 32:
			goto tr650
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr653
		case 112:
			goto tr653
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st451
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		default:
			goto tr535
		}
		goto tr0
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		switch data[p] {
		case 32:
			goto tr650
		case 43:
			goto tr533
		case 45:
			goto tr534
		case 47:
			goto tr535
		case 65:
			goto tr652
		case 80:
			goto tr652
		case 90:
			goto tr537
		case 95:
			goto tr535
		case 97:
			goto tr653
		case 112:
			goto tr653
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr535
			}
		case data[p] >= 66:
			goto tr535
		}
		goto tr0
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		switch data[p] {
		case 32:
			goto tr661
		case 43:
			goto tr589
		case 45:
			goto tr590
		case 46:
			goto tr662
		case 47:
			goto tr592
		case 65:
			goto tr663
		case 80:
			goto tr663
		case 90:
			goto tr594
		case 95:
			goto tr592
		case 97:
			goto tr664
		case 112:
			goto tr664
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr592
			}
		case data[p] >= 66:
			goto tr592
		}
		goto tr0
tr643:
//line ragel/datetime.rl:7
 pb = p 
	goto st453
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
//line ragel/parse_datetime.go:14390
		switch data[p] {
		case 32:
			goto tr644
		case 43:
			goto tr568
		case 45:
			goto tr569
		case 46:
			goto tr645
		case 47:
			goto tr571
		case 65:
			goto tr647
		case 80:
			goto tr647
		case 90:
			goto tr574
		case 95:
			goto tr571
		case 97:
			goto tr648
		case 112:
			goto tr648
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr571
			}
		case data[p] >= 66:
			goto tr571
		}
		goto tr0
tr632:
//line ragel/datetime.rl:7
 pb = p 
	goto st454
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
//line ragel/parse_datetime.go:14433
		switch data[p] {
		case 32:
			goto tr633
		case 43:
			goto tr549
		case 45:
			goto tr550
		case 47:
			goto tr551
		case 58:
			goto tr635
		case 65:
			goto tr636
		case 80:
			goto tr636
		case 90:
			goto tr555
		case 95:
			goto tr551
		case 97:
			goto tr637
		case 112:
			goto tr637
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr551
			}
		case data[p] >= 66:
			goto tr551
		}
		goto tr0
tr600:
//line ragel/datetime.rl:7
 pb = p 
	goto st455
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
//line ragel/parse_datetime.go:14476
		switch data[p] {
		case 32:
			goto tr602
		case 43:
			goto tr472
		case 45:
			goto tr473
		case 47:
			goto tr474
		case 58:
			goto tr604
		case 65:
			goto tr605
		case 80:
			goto tr605
		case 90:
			goto tr478
		case 95:
			goto tr474
		case 97:
			goto tr606
		case 112:
			goto tr606
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st422
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr474
				}
			case data[p] >= 66:
				goto tr474
			}
		default:
			goto st456
		}
		goto tr0
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		if 48 <= data[p] && data[p] <= 57 {
			goto st423
		}
		goto tr0
tr601:
//line ragel/datetime.rl:7
 pb = p 
	goto st457
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
//line ragel/parse_datetime.go:14537
		switch data[p] {
		case 32:
			goto tr602
		case 43:
			goto tr472
		case 45:
			goto tr473
		case 47:
			goto tr474
		case 58:
			goto tr604
		case 65:
			goto tr605
		case 80:
			goto tr605
		case 90:
			goto tr478
		case 95:
			goto tr474
		case 97:
			goto tr606
		case 112:
			goto tr606
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st456
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr474
			}
		default:
			goto tr474
		}
		goto tr0
tr460:
//line ragel/datetime.rl:7
 pb = p 
	goto st458
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
//line ragel/parse_datetime.go:14584
		if data[p] == 32 {
			goto tr467
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st414
		}
		goto tr0
tr461:
//line ragel/datetime.rl:7
 pb = p 
	goto st459
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
//line ragel/parse_datetime.go:14601
		if data[p] == 32 {
			goto tr467
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st414
		}
		goto tr0
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
		if data[p] == 108 {
			goto st461
		}
		goto tr0
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		if data[p] == 32 {
			goto tr456
		}
		goto tr0
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
		if data[p] == 103 {
			goto st463
		}
		goto tr0
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
		switch data[p] {
		case 32:
			goto tr668
		case 117:
			goto st464
		}
		goto tr0
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
		if data[p] == 115 {
			goto st465
		}
		goto tr0
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
		if data[p] == 116 {
			goto st466
		}
		goto tr0
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		if data[p] == 32 {
			goto tr668
		}
		goto tr0
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
		if data[p] == 101 {
			goto st468
		}
		goto tr0
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
		if data[p] == 99 {
			goto st469
		}
		goto tr0
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		switch data[p] {
		case 32:
			goto tr674
		case 101:
			goto st470
		}
		goto tr0
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		if data[p] == 109 {
			goto st471
		}
		goto tr0
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		if data[p] == 98 {
			goto st472
		}
		goto tr0
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
		if data[p] == 101 {
			goto st473
		}
		goto tr0
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		if data[p] == 114 {
			goto st474
		}
		goto tr0
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		if data[p] == 32 {
			goto tr674
		}
		goto tr0
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		if data[p] == 101 {
			goto st476
		}
		goto tr0
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
		if data[p] == 98 {
			goto st477
		}
		goto tr0
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		switch data[p] {
		case 32:
			goto tr682
		case 114:
			goto st478
		}
		goto tr0
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		if data[p] == 117 {
			goto st479
		}
		goto tr0
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		if data[p] == 97 {
			goto st480
		}
		goto tr0
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		if data[p] == 114 {
			goto st481
		}
		goto tr0
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		if data[p] == 121 {
			goto st482
		}
		goto tr0
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
		if data[p] == 32 {
			goto tr682
		}
		goto tr0
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
		switch data[p] {
		case 97:
			goto st484
		case 117:
			goto st490
		}
		goto tr0
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
		if data[p] == 110 {
			goto st485
		}
		goto tr0
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
		switch data[p] {
		case 32:
			goto tr691
		case 117:
			goto st486
		}
		goto tr0
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		if data[p] == 97 {
			goto st487
		}
		goto tr0
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
		if data[p] == 114 {
			goto st488
		}
		goto tr0
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
		if data[p] == 121 {
			goto st489
		}
		goto tr0
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
		if data[p] == 32 {
			goto tr691
		}
		goto tr0
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		switch data[p] {
		case 108:
			goto st491
		case 110:
			goto st493
		}
		goto tr0
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		switch data[p] {
		case 32:
			goto tr698
		case 121:
			goto st492
		}
		goto tr0
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
		if data[p] == 32 {
			goto tr698
		}
		goto tr0
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
		switch data[p] {
		case 32:
			goto tr700
		case 101:
			goto st494
		}
		goto tr0
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		if data[p] == 32 {
			goto tr700
		}
		goto tr0
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		if data[p] == 97 {
			goto st496
		}
		goto tr0
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
		switch data[p] {
		case 114:
			goto st497
		case 121:
			goto st500
		}
		goto tr0
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		switch data[p] {
		case 32:
			goto tr705
		case 99:
			goto st498
		}
		goto tr0
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		if data[p] == 104 {
			goto st499
		}
		goto tr0
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		if data[p] == 32 {
			goto tr705
		}
		goto tr0
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		if data[p] == 32 {
			goto tr708
		}
		goto tr0
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		if data[p] == 111 {
			goto st502
		}
		goto tr0
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		if data[p] == 118 {
			goto st503
		}
		goto tr0
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		switch data[p] {
		case 32:
			goto tr711
		case 101:
			goto st504
		}
		goto tr0
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		if data[p] == 109 {
			goto st505
		}
		goto tr0
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		if data[p] == 98 {
			goto st506
		}
		goto tr0
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		if data[p] == 101 {
			goto st507
		}
		goto tr0
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		if data[p] == 114 {
			goto st508
		}
		goto tr0
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		if data[p] == 32 {
			goto tr711
		}
		goto tr0
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		if data[p] == 99 {
			goto st510
		}
		goto tr0
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
		if data[p] == 116 {
			goto st511
		}
		goto tr0
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		switch data[p] {
		case 32:
			goto tr719
		case 111:
			goto st512
		}
		goto tr0
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		if data[p] == 98 {
			goto st513
		}
		goto tr0
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
		if data[p] == 101 {
			goto st514
		}
		goto tr0
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		if data[p] == 114 {
			goto st515
		}
		goto tr0
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
		if data[p] == 32 {
			goto tr719
		}
		goto tr0
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		if data[p] == 101 {
			goto st517
		}
		goto tr0
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		if data[p] == 112 {
			goto st518
		}
		goto tr0
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		switch data[p] {
		case 32:
			goto tr726
		case 116:
			goto st519
		}
		goto tr0
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		if data[p] == 101 {
			goto st520
		}
		goto tr0
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
		if data[p] == 109 {
			goto st521
		}
		goto tr0
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		if data[p] == 98 {
			goto st522
		}
		goto tr0
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		if data[p] == 101 {
			goto st523
		}
		goto tr0
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
		if data[p] == 114 {
			goto st524
		}
		goto tr0
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
		if data[p] == 32 {
			goto tr726
		}
		goto tr0
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		if data[p] == 32 {
			goto st526
		}
		goto tr0
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
		switch data[p] {
		case 48:
			goto tr734
		case 51:
			goto tr736
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr737
			}
		case data[p] >= 49:
			goto tr735
		}
		goto tr0
tr734:
//line ragel/datetime.rl:7
 pb = p 
	goto st527
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
//line ragel/parse_datetime.go:15271
		if 49 <= data[p] && data[p] <= 57 {
			goto st528
		}
		goto tr0
tr737:
//line ragel/datetime.rl:7
 pb = p 
	goto st528
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
//line ragel/parse_datetime.go:15285
		switch data[p] {
		case 32:
			goto tr739
		case 45:
			goto tr740
		}
		goto tr0
tr739:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st529
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
//line ragel/parse_datetime.go:15309
		switch data[p] {
		case 65:
			goto st530
		case 68:
			goto st540
		case 70:
			goto st548
		case 74:
			goto st556
		case 77:
			goto st568
		case 78:
			goto st574
		case 79:
			goto st582
		case 83:
			goto st589
		case 97:
			goto st530
		case 100:
			goto st540
		case 102:
			goto st548
		case 106:
			goto st556
		case 109:
			goto st568
		case 110:
			goto st574
		case 111:
			goto st582
		case 115:
			goto st589
		}
		goto tr0
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
		switch data[p] {
		case 112:
			goto st531
		case 117:
			goto st535
		}
		goto tr0
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		if data[p] == 114 {
			goto st532
		}
		goto tr0
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
		switch data[p] {
		case 32:
			goto tr242
		case 105:
			goto st533
		}
		goto tr0
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		if data[p] == 108 {
			goto st534
		}
		goto tr0
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
		if data[p] == 32 {
			goto tr242
		}
		goto tr0
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		if data[p] == 103 {
			goto st536
		}
		goto tr0
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		switch data[p] {
		case 32:
			goto tr250
		case 117:
			goto st537
		}
		goto tr0
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		if data[p] == 115 {
			goto st538
		}
		goto tr0
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
		if data[p] == 116 {
			goto st539
		}
		goto tr0
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
		if data[p] == 32 {
			goto tr250
		}
		goto tr0
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
		if data[p] == 101 {
			goto st541
		}
		goto tr0
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
		if data[p] == 99 {
			goto st542
		}
		goto tr0
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
		switch data[p] {
		case 32:
			goto tr256
		case 101:
			goto st543
		}
		goto tr0
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
		if data[p] == 109 {
			goto st544
		}
		goto tr0
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
		if data[p] == 98 {
			goto st545
		}
		goto tr0
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
		if data[p] == 101 {
			goto st546
		}
		goto tr0
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
		if data[p] == 114 {
			goto st547
		}
		goto tr0
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
		if data[p] == 32 {
			goto tr256
		}
		goto tr0
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
		if data[p] == 101 {
			goto st549
		}
		goto tr0
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
		if data[p] == 98 {
			goto st550
		}
		goto tr0
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
		switch data[p] {
		case 32:
			goto tr264
		case 114:
			goto st551
		}
		goto tr0
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
		if data[p] == 117 {
			goto st552
		}
		goto tr0
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
		if data[p] == 97 {
			goto st553
		}
		goto tr0
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
		if data[p] == 114 {
			goto st554
		}
		goto tr0
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
		if data[p] == 121 {
			goto st555
		}
		goto tr0
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
		if data[p] == 32 {
			goto tr264
		}
		goto tr0
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
		switch data[p] {
		case 97:
			goto st557
		case 117:
			goto st563
		}
		goto tr0
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
		if data[p] == 110 {
			goto st558
		}
		goto tr0
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
		switch data[p] {
		case 32:
			goto tr273
		case 117:
			goto st559
		}
		goto tr0
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
		if data[p] == 97 {
			goto st560
		}
		goto tr0
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
		if data[p] == 114 {
			goto st561
		}
		goto tr0
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
		if data[p] == 121 {
			goto st562
		}
		goto tr0
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
		if data[p] == 32 {
			goto tr273
		}
		goto tr0
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
		switch data[p] {
		case 108:
			goto st564
		case 110:
			goto st566
		}
		goto tr0
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
		switch data[p] {
		case 32:
			goto tr280
		case 121:
			goto st565
		}
		goto tr0
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
		if data[p] == 32 {
			goto tr280
		}
		goto tr0
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
		switch data[p] {
		case 32:
			goto tr282
		case 101:
			goto st567
		}
		goto tr0
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		if data[p] == 32 {
			goto tr282
		}
		goto tr0
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
		if data[p] == 97 {
			goto st569
		}
		goto tr0
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		switch data[p] {
		case 114:
			goto st570
		case 121:
			goto st573
		}
		goto tr0
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		switch data[p] {
		case 32:
			goto tr287
		case 99:
			goto st571
		}
		goto tr0
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
		if data[p] == 104 {
			goto st572
		}
		goto tr0
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
		if data[p] == 32 {
			goto tr287
		}
		goto tr0
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		if data[p] == 32 {
			goto tr290
		}
		goto tr0
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
		if data[p] == 111 {
			goto st575
		}
		goto tr0
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
		if data[p] == 118 {
			goto st576
		}
		goto tr0
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
		switch data[p] {
		case 32:
			goto tr293
		case 101:
			goto st577
		}
		goto tr0
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
		if data[p] == 109 {
			goto st578
		}
		goto tr0
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
		if data[p] == 98 {
			goto st579
		}
		goto tr0
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
		if data[p] == 101 {
			goto st580
		}
		goto tr0
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		if data[p] == 114 {
			goto st581
		}
		goto tr0
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
		if data[p] == 32 {
			goto tr293
		}
		goto tr0
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		if data[p] == 99 {
			goto st583
		}
		goto tr0
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		if data[p] == 116 {
			goto st584
		}
		goto tr0
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		switch data[p] {
		case 32:
			goto tr301
		case 111:
			goto st585
		}
		goto tr0
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		if data[p] == 98 {
			goto st586
		}
		goto tr0
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		if data[p] == 101 {
			goto st587
		}
		goto tr0
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
		if data[p] == 114 {
			goto st588
		}
		goto tr0
	st588:
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
		if data[p] == 32 {
			goto tr301
		}
		goto tr0
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		if data[p] == 101 {
			goto st590
		}
		goto tr0
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
		if data[p] == 112 {
			goto st591
		}
		goto tr0
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		switch data[p] {
		case 32:
			goto tr308
		case 116:
			goto st592
		}
		goto tr0
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
		if data[p] == 101 {
			goto st593
		}
		goto tr0
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
	st_case_593:
		if data[p] == 109 {
			goto st594
		}
		goto tr0
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		if data[p] == 98 {
			goto st595
		}
		goto tr0
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		if data[p] == 101 {
			goto st596
		}
		goto tr0
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		if data[p] == 114 {
			goto st597
		}
		goto tr0
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
	st_case_597:
		if data[p] == 32 {
			goto tr308
		}
		goto tr0
tr740:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st598
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
//line ragel/parse_datetime.go:16018
		switch data[p] {
		case 65:
			goto st599
		case 68:
			goto st609
		case 70:
			goto st617
		case 74:
			goto st625
		case 77:
			goto st637
		case 78:
			goto st643
		case 79:
			goto st651
		case 83:
			goto st658
		case 97:
			goto st599
		case 100:
			goto st609
		case 102:
			goto st617
		case 106:
			goto st625
		case 109:
			goto st637
		case 110:
			goto st643
		case 111:
			goto st651
		case 115:
			goto st658
		}
		goto tr0
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		switch data[p] {
		case 112:
			goto st600
		case 117:
			goto st604
		}
		goto tr0
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		if data[p] == 114 {
			goto st601
		}
		goto tr0
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
		switch data[p] {
		case 45:
			goto tr160
		case 105:
			goto st602
		}
		goto tr0
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
		if data[p] == 108 {
			goto st603
		}
		goto tr0
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
		if data[p] == 45 {
			goto tr160
		}
		goto tr0
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		if data[p] == 103 {
			goto st605
		}
		goto tr0
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
		switch data[p] {
		case 45:
			goto tr166
		case 117:
			goto st606
		}
		goto tr0
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
		if data[p] == 115 {
			goto st607
		}
		goto tr0
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
		if data[p] == 116 {
			goto st608
		}
		goto tr0
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		if data[p] == 45 {
			goto tr166
		}
		goto tr0
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
		if data[p] == 101 {
			goto st610
		}
		goto tr0
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		if data[p] == 99 {
			goto st611
		}
		goto tr0
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		switch data[p] {
		case 45:
			goto tr172
		case 101:
			goto st612
		}
		goto tr0
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
		if data[p] == 109 {
			goto st613
		}
		goto tr0
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		if data[p] == 98 {
			goto st614
		}
		goto tr0
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
		if data[p] == 101 {
			goto st615
		}
		goto tr0
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
		if data[p] == 114 {
			goto st616
		}
		goto tr0
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
		if data[p] == 45 {
			goto tr172
		}
		goto tr0
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
		if data[p] == 101 {
			goto st618
		}
		goto tr0
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
		if data[p] == 98 {
			goto st619
		}
		goto tr0
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		switch data[p] {
		case 45:
			goto tr180
		case 114:
			goto st620
		}
		goto tr0
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		if data[p] == 117 {
			goto st621
		}
		goto tr0
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		if data[p] == 97 {
			goto st622
		}
		goto tr0
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		if data[p] == 114 {
			goto st623
		}
		goto tr0
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
		if data[p] == 121 {
			goto st624
		}
		goto tr0
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
		if data[p] == 45 {
			goto tr180
		}
		goto tr0
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
		switch data[p] {
		case 97:
			goto st626
		case 117:
			goto st632
		}
		goto tr0
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
		if data[p] == 110 {
			goto st627
		}
		goto tr0
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
		switch data[p] {
		case 45:
			goto tr189
		case 117:
			goto st628
		}
		goto tr0
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
		if data[p] == 97 {
			goto st629
		}
		goto tr0
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
		if data[p] == 114 {
			goto st630
		}
		goto tr0
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
		if data[p] == 121 {
			goto st631
		}
		goto tr0
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		if data[p] == 45 {
			goto tr189
		}
		goto tr0
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
		switch data[p] {
		case 108:
			goto st633
		case 110:
			goto st635
		}
		goto tr0
	st633:
		if p++; p == pe {
			goto _test_eof633
		}
	st_case_633:
		switch data[p] {
		case 45:
			goto tr196
		case 121:
			goto st634
		}
		goto tr0
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
		if data[p] == 45 {
			goto tr196
		}
		goto tr0
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
		switch data[p] {
		case 45:
			goto tr198
		case 101:
			goto st636
		}
		goto tr0
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
		if data[p] == 45 {
			goto tr198
		}
		goto tr0
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
		if data[p] == 97 {
			goto st638
		}
		goto tr0
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
		switch data[p] {
		case 114:
			goto st639
		case 121:
			goto st642
		}
		goto tr0
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
		switch data[p] {
		case 45:
			goto tr203
		case 99:
			goto st640
		}
		goto tr0
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
		if data[p] == 104 {
			goto st641
		}
		goto tr0
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
		if data[p] == 45 {
			goto tr203
		}
		goto tr0
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
		if data[p] == 45 {
			goto tr206
		}
		goto tr0
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
		if data[p] == 111 {
			goto st644
		}
		goto tr0
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
		if data[p] == 118 {
			goto st645
		}
		goto tr0
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
		switch data[p] {
		case 45:
			goto tr209
		case 101:
			goto st646
		}
		goto tr0
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
		if data[p] == 109 {
			goto st647
		}
		goto tr0
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
		if data[p] == 98 {
			goto st648
		}
		goto tr0
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
		if data[p] == 101 {
			goto st649
		}
		goto tr0
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
		if data[p] == 114 {
			goto st650
		}
		goto tr0
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
		if data[p] == 45 {
			goto tr209
		}
		goto tr0
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		if data[p] == 99 {
			goto st652
		}
		goto tr0
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
		if data[p] == 116 {
			goto st653
		}
		goto tr0
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		switch data[p] {
		case 45:
			goto tr217
		case 111:
			goto st654
		}
		goto tr0
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
		if data[p] == 98 {
			goto st655
		}
		goto tr0
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
		if data[p] == 101 {
			goto st656
		}
		goto tr0
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
		if data[p] == 114 {
			goto st657
		}
		goto tr0
	st657:
		if p++; p == pe {
			goto _test_eof657
		}
	st_case_657:
		if data[p] == 45 {
			goto tr217
		}
		goto tr0
	st658:
		if p++; p == pe {
			goto _test_eof658
		}
	st_case_658:
		if data[p] == 101 {
			goto st659
		}
		goto tr0
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
		if data[p] == 112 {
			goto st660
		}
		goto tr0
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
		switch data[p] {
		case 45:
			goto tr224
		case 116:
			goto st661
		}
		goto tr0
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
	st_case_661:
		if data[p] == 101 {
			goto st662
		}
		goto tr0
	st662:
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
		if data[p] == 109 {
			goto st663
		}
		goto tr0
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
		if data[p] == 98 {
			goto st664
		}
		goto tr0
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
		if data[p] == 101 {
			goto st665
		}
		goto tr0
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
		if data[p] == 114 {
			goto st666
		}
		goto tr0
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
		if data[p] == 45 {
			goto tr224
		}
		goto tr0
tr735:
//line ragel/datetime.rl:7
 pb = p 
	goto st667
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
	st_case_667:
//line ragel/parse_datetime.go:16720
		switch data[p] {
		case 32:
			goto tr739
		case 45:
			goto tr740
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st528
		}
		goto tr0
tr736:
//line ragel/datetime.rl:7
 pb = p 
	goto st668
	st668:
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
//line ragel/parse_datetime.go:16740
		switch data[p] {
		case 32:
			goto tr739
		case 45:
			goto tr740
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st528
		}
		goto tr0
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
		if data[p] == 97 {
			goto st670
		}
		goto tr0
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
		if data[p] == 121 {
			goto st671
		}
		goto tr0
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
		switch data[p] {
		case 32:
			goto st343
		case 44:
			goto st525
		}
		goto tr0
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
	st_case_672:
		switch data[p] {
		case 97:
			goto st673
		case 117:
			goto st679
		}
		goto tr0
	st673:
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
		if data[p] == 110 {
			goto st674
		}
		goto tr0
	st674:
		if p++; p == pe {
			goto _test_eof674
		}
	st_case_674:
		switch data[p] {
		case 32:
			goto tr882
		case 45:
			goto tr883
		case 47:
			goto tr883
		case 117:
			goto st675
		}
		goto tr0
	st675:
		if p++; p == pe {
			goto _test_eof675
		}
	st_case_675:
		if data[p] == 97 {
			goto st676
		}
		goto tr0
	st676:
		if p++; p == pe {
			goto _test_eof676
		}
	st_case_676:
		if data[p] == 114 {
			goto st677
		}
		goto tr0
	st677:
		if p++; p == pe {
			goto _test_eof677
		}
	st_case_677:
		if data[p] == 121 {
			goto st678
		}
		goto tr0
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
		switch data[p] {
		case 32:
			goto tr882
		case 45:
			goto tr883
		case 47:
			goto tr883
		}
		goto tr0
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
		switch data[p] {
		case 108:
			goto st680
		case 110:
			goto st682
		}
		goto tr0
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
		switch data[p] {
		case 32:
			goto tr890
		case 45:
			goto tr891
		case 47:
			goto tr891
		case 121:
			goto st681
		}
		goto tr0
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
	st_case_681:
		switch data[p] {
		case 32:
			goto tr890
		case 45:
			goto tr891
		case 47:
			goto tr891
		}
		goto tr0
	st682:
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
		switch data[p] {
		case 32:
			goto tr893
		case 45:
			goto tr894
		case 47:
			goto tr894
		case 101:
			goto st683
		}
		goto tr0
	st683:
		if p++; p == pe {
			goto _test_eof683
		}
	st_case_683:
		switch data[p] {
		case 32:
			goto tr893
		case 45:
			goto tr894
		case 47:
			goto tr894
		}
		goto tr0
	st684:
		if p++; p == pe {
			goto _test_eof684
		}
	st_case_684:
		switch data[p] {
		case 97:
			goto st685
		case 111:
			goto st690
		}
		goto tr0
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
	st_case_685:
		switch data[p] {
		case 114:
			goto st686
		case 121:
			goto st689
		}
		goto tr0
	st686:
		if p++; p == pe {
			goto _test_eof686
		}
	st_case_686:
		switch data[p] {
		case 32:
			goto tr900
		case 45:
			goto tr901
		case 47:
			goto tr901
		case 99:
			goto st687
		}
		goto tr0
	st687:
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
		if data[p] == 104 {
			goto st688
		}
		goto tr0
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
		switch data[p] {
		case 32:
			goto tr900
		case 45:
			goto tr901
		case 47:
			goto tr901
		}
		goto tr0
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
		switch data[p] {
		case 32:
			goto tr904
		case 45:
			goto tr905
		case 47:
			goto tr905
		}
		goto tr0
	st690:
		if p++; p == pe {
			goto _test_eof690
		}
	st_case_690:
		if data[p] == 110 {
			goto st342
		}
		goto tr0
	st691:
		if p++; p == pe {
			goto _test_eof691
		}
	st_case_691:
		if data[p] == 111 {
			goto st692
		}
		goto tr0
	st692:
		if p++; p == pe {
			goto _test_eof692
		}
	st_case_692:
		if data[p] == 118 {
			goto st693
		}
		goto tr0
	st693:
		if p++; p == pe {
			goto _test_eof693
		}
	st_case_693:
		switch data[p] {
		case 32:
			goto tr908
		case 45:
			goto tr909
		case 47:
			goto tr909
		case 101:
			goto st694
		}
		goto tr0
	st694:
		if p++; p == pe {
			goto _test_eof694
		}
	st_case_694:
		if data[p] == 109 {
			goto st695
		}
		goto tr0
	st695:
		if p++; p == pe {
			goto _test_eof695
		}
	st_case_695:
		if data[p] == 98 {
			goto st696
		}
		goto tr0
	st696:
		if p++; p == pe {
			goto _test_eof696
		}
	st_case_696:
		if data[p] == 101 {
			goto st697
		}
		goto tr0
	st697:
		if p++; p == pe {
			goto _test_eof697
		}
	st_case_697:
		if data[p] == 114 {
			goto st698
		}
		goto tr0
	st698:
		if p++; p == pe {
			goto _test_eof698
		}
	st_case_698:
		switch data[p] {
		case 32:
			goto tr908
		case 45:
			goto tr909
		case 47:
			goto tr909
		}
		goto tr0
	st699:
		if p++; p == pe {
			goto _test_eof699
		}
	st_case_699:
		if data[p] == 99 {
			goto st700
		}
		goto tr0
	st700:
		if p++; p == pe {
			goto _test_eof700
		}
	st_case_700:
		if data[p] == 116 {
			goto st701
		}
		goto tr0
	st701:
		if p++; p == pe {
			goto _test_eof701
		}
	st_case_701:
		switch data[p] {
		case 32:
			goto tr917
		case 45:
			goto tr918
		case 47:
			goto tr918
		case 111:
			goto st702
		}
		goto tr0
	st702:
		if p++; p == pe {
			goto _test_eof702
		}
	st_case_702:
		if data[p] == 98 {
			goto st703
		}
		goto tr0
	st703:
		if p++; p == pe {
			goto _test_eof703
		}
	st_case_703:
		if data[p] == 101 {
			goto st704
		}
		goto tr0
	st704:
		if p++; p == pe {
			goto _test_eof704
		}
	st_case_704:
		if data[p] == 114 {
			goto st705
		}
		goto tr0
	st705:
		if p++; p == pe {
			goto _test_eof705
		}
	st_case_705:
		switch data[p] {
		case 32:
			goto tr917
		case 45:
			goto tr918
		case 47:
			goto tr918
		}
		goto tr0
	st706:
		if p++; p == pe {
			goto _test_eof706
		}
	st_case_706:
		switch data[p] {
		case 97:
			goto st707
		case 101:
			goto st711
		case 117:
			goto st690
		}
		goto tr0
	st707:
		if p++; p == pe {
			goto _test_eof707
		}
	st_case_707:
		if data[p] == 116 {
			goto st708
		}
		goto tr0
	st708:
		if p++; p == pe {
			goto _test_eof708
		}
	st_case_708:
		switch data[p] {
		case 32:
			goto st343
		case 44:
			goto st525
		case 117:
			goto st709
		}
		goto tr0
	st709:
		if p++; p == pe {
			goto _test_eof709
		}
	st_case_709:
		if data[p] == 114 {
			goto st710
		}
		goto tr0
	st710:
		if p++; p == pe {
			goto _test_eof710
		}
	st_case_710:
		if data[p] == 100 {
			goto st669
		}
		goto tr0
	st711:
		if p++; p == pe {
			goto _test_eof711
		}
	st_case_711:
		if data[p] == 112 {
			goto st712
		}
		goto tr0
	st712:
		if p++; p == pe {
			goto _test_eof712
		}
	st_case_712:
		switch data[p] {
		case 32:
			goto tr929
		case 45:
			goto tr930
		case 47:
			goto tr930
		case 116:
			goto st713
		}
		goto tr0
	st713:
		if p++; p == pe {
			goto _test_eof713
		}
	st_case_713:
		if data[p] == 101 {
			goto st714
		}
		goto tr0
	st714:
		if p++; p == pe {
			goto _test_eof714
		}
	st_case_714:
		if data[p] == 109 {
			goto st715
		}
		goto tr0
	st715:
		if p++; p == pe {
			goto _test_eof715
		}
	st_case_715:
		if data[p] == 98 {
			goto st716
		}
		goto tr0
	st716:
		if p++; p == pe {
			goto _test_eof716
		}
	st_case_716:
		if data[p] == 101 {
			goto st717
		}
		goto tr0
	st717:
		if p++; p == pe {
			goto _test_eof717
		}
	st_case_717:
		if data[p] == 114 {
			goto st718
		}
		goto tr0
	st718:
		if p++; p == pe {
			goto _test_eof718
		}
	st_case_718:
		switch data[p] {
		case 32:
			goto tr929
		case 45:
			goto tr930
		case 47:
			goto tr930
		}
		goto tr0
	st719:
		if p++; p == pe {
			goto _test_eof719
		}
	st_case_719:
		switch data[p] {
		case 104:
			goto st720
		case 117:
			goto st723
		}
		goto tr0
	st720:
		if p++; p == pe {
			goto _test_eof720
		}
	st_case_720:
		if data[p] == 117 {
			goto st721
		}
		goto tr0
	st721:
		if p++; p == pe {
			goto _test_eof721
		}
	st_case_721:
		switch data[p] {
		case 32:
			goto st343
		case 44:
			goto st525
		case 114:
			goto st722
		}
		goto tr0
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
		if data[p] == 115 {
			goto st710
		}
		goto tr0
	st723:
		if p++; p == pe {
			goto _test_eof723
		}
	st_case_723:
		if data[p] == 101 {
			goto st724
		}
		goto tr0
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
	st_case_724:
		switch data[p] {
		case 32:
			goto st343
		case 44:
			goto st525
		case 115:
			goto st710
		}
		goto tr0
	st725:
		if p++; p == pe {
			goto _test_eof725
		}
	st_case_725:
		if data[p] == 101 {
			goto st726
		}
		goto tr0
	st726:
		if p++; p == pe {
			goto _test_eof726
		}
	st_case_726:
		if data[p] == 100 {
			goto st727
		}
		goto tr0
	st727:
		if p++; p == pe {
			goto _test_eof727
		}
	st_case_727:
		switch data[p] {
		case 32:
			goto st343
		case 44:
			goto st525
		case 110:
			goto st728
		}
		goto tr0
	st728:
		if p++; p == pe {
			goto _test_eof728
		}
	st_case_728:
		if data[p] == 101 {
			goto st722
		}
		goto tr0
	st729:
		if p++; p == pe {
			goto _test_eof729
		}
	st_case_729:
		if data[p] == 101 {
			goto st334
		}
		goto tr0
	st730:
		if p++; p == pe {
			goto _test_eof730
		}
	st_case_730:
		if data[p] == 97 {
			goto st685
		}
		goto tr0
	st731:
		if p++; p == pe {
			goto _test_eof731
		}
	st_case_731:
		if data[p] == 101 {
			goto st711
		}
		goto tr0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof732: cs = 732; goto _test_eof
	_test_eof733: cs = 733; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof734: cs = 734; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof735: cs = 735; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof736: cs = 736; goto _test_eof
	_test_eof737: cs = 737; goto _test_eof
	_test_eof738: cs = 738; goto _test_eof
	_test_eof739: cs = 739; goto _test_eof
	_test_eof740: cs = 740; goto _test_eof
	_test_eof741: cs = 741; goto _test_eof
	_test_eof742: cs = 742; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof743: cs = 743; goto _test_eof
	_test_eof744: cs = 744; goto _test_eof
	_test_eof745: cs = 745; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof746: cs = 746; goto _test_eof
	_test_eof747: cs = 747; goto _test_eof
	_test_eof748: cs = 748; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof749: cs = 749; goto _test_eof
	_test_eof750: cs = 750; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof751: cs = 751; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof752: cs = 752; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof753: cs = 753; goto _test_eof
	_test_eof754: cs = 754; goto _test_eof
	_test_eof755: cs = 755; goto _test_eof
	_test_eof756: cs = 756; goto _test_eof
	_test_eof757: cs = 757; goto _test_eof
	_test_eof758: cs = 758; goto _test_eof
	_test_eof759: cs = 759; goto _test_eof
	_test_eof760: cs = 760; goto _test_eof
	_test_eof761: cs = 761; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof762: cs = 762; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof763: cs = 763; goto _test_eof
	_test_eof764: cs = 764; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof765: cs = 765; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
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
	_test_eof27: cs = 27; goto _test_eof
	_test_eof779: cs = 779; goto _test_eof
	_test_eof780: cs = 780; goto _test_eof
	_test_eof781: cs = 781; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof782: cs = 782; goto _test_eof
	_test_eof783: cs = 783; goto _test_eof
	_test_eof784: cs = 784; goto _test_eof
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
	_test_eof105: cs = 105; goto _test_eof
	_test_eof785: cs = 785; goto _test_eof
	_test_eof786: cs = 786; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof787: cs = 787; goto _test_eof
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
	_test_eof788: cs = 788; goto _test_eof
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
	_test_eof789: cs = 789; goto _test_eof
	_test_eof266: cs = 266; goto _test_eof
	_test_eof267: cs = 267; goto _test_eof
	_test_eof790: cs = 790; goto _test_eof
	_test_eof791: cs = 791; goto _test_eof
	_test_eof792: cs = 792; goto _test_eof
	_test_eof793: cs = 793; goto _test_eof
	_test_eof794: cs = 794; goto _test_eof
	_test_eof795: cs = 795; goto _test_eof
	_test_eof796: cs = 796; goto _test_eof
	_test_eof268: cs = 268; goto _test_eof
	_test_eof269: cs = 269; goto _test_eof
	_test_eof270: cs = 270; goto _test_eof
	_test_eof797: cs = 797; goto _test_eof
	_test_eof798: cs = 798; goto _test_eof
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
	_test_eof336: cs = 336; goto _test_eof
	_test_eof337: cs = 337; goto _test_eof
	_test_eof338: cs = 338; goto _test_eof
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
	_test_eof799: cs = 799; goto _test_eof
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

	_test_eof: {}
	if p == eof {
		switch cs {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279, 280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299, 300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339, 340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 383, 384, 385, 386, 387, 388, 389, 390, 391, 392, 393, 394, 395, 396, 397, 398, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422, 423, 424, 425, 426, 427, 428, 429, 430, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 446, 447, 448, 449, 450, 451, 452, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 495, 496, 497, 498, 499, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 510, 511, 512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 524, 525, 526, 527, 528, 529, 530, 531, 532, 533, 534, 535, 536, 537, 538, 539, 540, 541, 542, 543, 544, 545, 546, 547, 548, 549, 550, 551, 552, 553, 554, 555, 556, 557, 558, 559, 560, 561, 562, 563, 564, 565, 566, 567, 568, 569, 570, 571, 572, 573, 574, 575, 576, 577, 578, 579, 580, 581, 582, 583, 584, 585, 586, 587, 588, 589, 590, 591, 592, 593, 594, 595, 596, 597, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 608, 609, 610, 611, 612, 613, 614, 615, 616, 617, 618, 619, 620, 621, 622, 623, 624, 625, 626, 627, 628, 629, 630, 631, 632, 633, 634, 635, 636, 637, 638, 639, 640, 641, 642, 643, 644, 645, 646, 647, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 658, 659, 660, 661, 662, 663, 664, 665, 666, 667, 668, 669, 670, 671, 672, 673, 674, 675, 676, 677, 678, 679, 680, 681, 682, 683, 684, 685, 686, 687, 688, 689, 690, 691, 692, 693, 694, 695, 696, 697, 698, 699, 700, 701, 702, 703, 704, 705, 706, 707, 708, 709, 710, 711, 712, 713, 714, 715, 716, 717, 718, 719, 720, 721, 722, 723, 724, 725, 726, 727, 728, 729, 730, 731:
//line ragel/datetime.rl:5
 return st, err 
		case 739, 750, 781, 793, 798:
//line ragel/datetime.rl:9
 st.Zoned = true 
		case 786:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

		case 788, 789, 799:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

		case 787:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

		case 732, 785:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

		case 752, 762:
//line ragel/datetime.rl:37

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

		case 735, 749:
//line ragel/datetime.rl:52

    st.Ad_bc = ADBC_BC;

		case 747:
//line ragel/datetime.rl:56

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

		case 782, 783, 784:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

		case 751:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

		case 744, 778, 779:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

		case 764:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

		case 763, 777:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

		case 775:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

		case 765, 776:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

		case 753, 754, 755, 756, 757, 758, 759, 760, 761, 766, 767, 768, 769, 770, 771, 772, 773, 774:
//line ragel/datetime.rl:119

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

		case 740, 741, 794, 795:
//line ragel/datetime.rl:159

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:9
 st.Zoned = true 
		case 736, 737, 738, 742, 790, 791, 792, 796:
//line ragel/datetime.rl:168
 
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

//line ragel/datetime.rl:9
 st.Zoned = true 
		case 743, 797:
//line ragel/datetime.rl:196

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:9
 st.Zoned = true 
//line ragel/parse_datetime.go:18446
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

