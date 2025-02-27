
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
const datetime_parser_first_final int = 413
const datetime_parser_error int = 0

const datetime_parser_en_main int = 1


//line ragel/parse_datetime.go.rl:23



//line ragel/parse_datetime.go:31
const start int = 1

const en_main int = 1


//line ragel/parse_datetime.go.rl:26

func Parse(data string) (st Datetime, err error) {
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
	case 413:
		goto st_case_413
	case 414:
		goto st_case_414
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 415:
		goto st_case_415
	case 11:
		goto st_case_11
	case 416:
		goto st_case_416
	case 12:
		goto st_case_12
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
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 424:
		goto st_case_424
	case 425:
		goto st_case_425
	case 426:
		goto st_case_426
	case 16:
		goto st_case_16
	case 427:
		goto st_case_427
	case 428:
		goto st_case_428
	case 429:
		goto st_case_429
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 430:
		goto st_case_430
	case 431:
		goto st_case_431
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 432:
		goto st_case_432
	case 21:
		goto st_case_21
	case 433:
		goto st_case_433
	case 22:
		goto st_case_22
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
	case 23:
		goto st_case_23
	case 444:
		goto st_case_444
	case 24:
		goto st_case_24
	case 445:
		goto st_case_445
	case 446:
		goto st_case_446
	case 25:
		goto st_case_25
	case 447:
		goto st_case_447
	case 26:
		goto st_case_26
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
	case 27:
		goto st_case_27
	case 461:
		goto st_case_461
	case 462:
		goto st_case_462
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 463:
		goto st_case_463
	case 464:
		goto st_case_464
	case 465:
		goto st_case_465
	case 31:
		goto st_case_31
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
	case 466:
		goto st_case_466
	case 467:
		goto st_case_467
	case 468:
		goto st_case_468
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 469:
		goto st_case_469
	case 470:
		goto st_case_470
	case 471:
		goto st_case_471
	case 82:
		goto st_case_82
	case 472:
		goto st_case_472
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
	case 473:
		goto st_case_473
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
	case 474:
		goto st_case_474
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
	case 475:
		goto st_case_475
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
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 48:
			goto tr1
		case 51:
			goto tr3
		case 70:
			goto st182
		case 77:
			goto st396
		case 83:
			goto st398
		case 84:
			goto st403
		case 87:
			goto st409
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
//line ragel/parse_datetime.go:1042
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
//line ragel/parse_datetime.go:1056
		if data[p] == 48 {
			goto st3
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st132
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
		switch data[p] {
		case 46:
			goto tr15
		case 48:
			goto tr16
		case 49:
			goto tr17
		case 65:
			goto tr19
		case 68:
			goto tr20
		case 70:
			goto tr21
		case 74:
			goto tr22
		case 77:
			goto tr23
		case 78:
			goto tr24
		case 79:
			goto tr25
		case 83:
			goto tr26
		}
		switch {
		case data[p] > 47:
			if 50 <= data[p] && data[p] <= 57 {
				goto tr18
			}
		case data[p] >= 45:
			goto tr14
		}
		goto tr0
tr14:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line ragel/parse_datetime.go:1131
		switch data[p] {
		case 48:
			goto tr27
		case 49:
			goto tr28
		case 65:
			goto st33
		case 68:
			goto st40
		case 70:
			goto st43
		case 74:
			goto st51
		case 77:
			goto st61
		case 78:
			goto st67
		case 79:
			goto st70
		case 83:
			goto st73
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr29
		}
		goto tr0
tr27:
//line ragel/datetime.rl:7
 pb = p 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line ragel/parse_datetime.go:1167
		if data[p] == 48 {
			goto st8
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st28
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if 48 <= data[p] && data[p] <= 57 {
			goto st413
		}
		goto tr0
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
		switch data[p] {
		case 32:
			goto tr582
		case 43:
			goto tr584
		case 45:
			goto tr585
		case 47:
			goto tr586
		case 84:
			goto tr587
		case 90:
			goto tr588
		case 95:
			goto tr586
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr586
			}
		case data[p] >= 65:
			goto tr586
		}
		goto st0
tr719:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st414
tr701:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st414
tr582:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

	goto st414
tr707:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st414
tr713:
//line ragel/datetime.rl:19

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st414
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
//line ragel/parse_datetime.go:1265
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr592
		case 50:
			goto tr594
		case 65:
			goto tr596
		case 66:
			goto tr597
		case 95:
			goto tr592
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr593
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr592
				}
			case data[p] >= 67:
				goto tr592
			}
		default:
			goto tr595
		}
		goto st0
tr598:
//line ragel/datetime.rl:155
 
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
        case 1:{st.Offset_hour, _ = strconv.Atoi(data[pb:pb+1])}
        case 2:{st.Offset_hour, _ = strconv.Atoi(data[pb:pb+2])}
        case 3:{ 
            num , _ := strconv.Atoi(data[pb:pb+3])
            st.Offset_hour = num/100
            st.Offset_minute = num%100
            if st.Offset_minute >=60 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        case 4:{ 
            num := parse_digits(data[pb:pb+4])
            st.Offset_hour = num/100
            st.Offset_minute = num%100
            if st.Offset_minute >=60 || st.Offset_hour>=15 {
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
tr602:
//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st9
tr605:
//line ragel/datetime.rl:145

    switch p - pb {
        case 1: st.Offset_minute, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Offset_minute, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st9
tr607:
//line ragel/datetime.rl:193

    st.Zone_name_or_abbrev = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line ragel/parse_datetime.go:1376
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
			goto st415
		}
		goto tr0
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 67 {
			goto st416
		}
		goto tr0
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		goto st0
tr720:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st12
tr702:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st12
tr609:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st12
tr621:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st12
tr627:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st12
tr635:
//line ragel/datetime.rl:33

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
tr642:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

	goto st12
tr656:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st12
tr665:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st12
tr673:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st12
tr693:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st12
tr584:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

	goto st12
tr708:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st12
tr714:
//line ragel/datetime.rl:19

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line ragel/parse_datetime.go:1562
		if data[p] == 50 {
			goto tr46
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr47
			}
		case data[p] >= 48:
			goto tr45
		}
		goto tr0
tr45:
//line ragel/datetime.rl:7
 pb = p 
	goto st417
tr48:
//line ragel/datetime.rl:133
 st.Negative_offset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st417
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
//line ragel/parse_datetime.go:1590
		switch data[p] {
		case 32:
			goto tr598
		case 58:
			goto tr600
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st418
		}
		goto st0
tr47:
//line ragel/datetime.rl:7
 pb = p 
	goto st418
tr50:
//line ragel/datetime.rl:133
 st.Negative_offset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st418
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
//line ragel/parse_datetime.go:1616
		switch data[p] {
		case 32:
			goto tr598
		case 58:
			goto tr600
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st419
		}
		goto st0
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		if data[p] == 32 {
			goto tr598
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st419
		}
		goto st0
tr600:
//line ragel/datetime.rl:135

    switch p - pb {
        case 1: st.Offset_hour, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Offset_hour, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st420
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
//line ragel/parse_datetime.go:1656
		if data[p] == 32 {
			goto tr602
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr604
			}
		case data[p] >= 48:
			goto tr603
		}
		goto st0
tr603:
//line ragel/datetime.rl:7
 pb = p 
	goto st421
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
//line ragel/parse_datetime.go:1678
		if data[p] == 32 {
			goto tr605
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st422
		}
		goto st0
tr604:
//line ragel/datetime.rl:7
 pb = p 
	goto st422
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
//line ragel/parse_datetime.go:1695
		if data[p] == 32 {
			goto tr605
		}
		goto st0
tr46:
//line ragel/datetime.rl:7
 pb = p 
	goto st423
tr49:
//line ragel/datetime.rl:133
 st.Negative_offset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st423
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
//line ragel/parse_datetime.go:1715
		switch data[p] {
		case 32:
			goto tr598
		case 58:
			goto tr600
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st419
			}
		case data[p] >= 48:
			goto st418
		}
		goto st0
tr721:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st13
tr703:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st13
tr610:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st13
tr622:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st13
tr628:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st13
tr636:
//line ragel/datetime.rl:33

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
tr643:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

	goto st13
tr657:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st13
tr666:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st13
tr674:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st13
tr694:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st13
tr585:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

	goto st13
tr709:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st13
tr715:
//line ragel/datetime.rl:19

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line ragel/parse_datetime.go:1879
		if data[p] == 50 {
			goto tr49
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr50
			}
		case data[p] >= 48:
			goto tr48
		}
		goto tr0
tr592:
//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr722:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr611:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr629:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr638:
//line ragel/datetime.rl:33

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
tr644:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr658:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr667:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr676:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr696:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr586:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr704:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr710:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr716:
//line ragel/datetime.rl:19

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line ragel/parse_datetime.go:2045
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
tr654:
//line ragel/datetime.rl:7
 pb = p 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line ragel/parse_datetime.go:2070
		switch data[p] {
		case 47:
			goto st424
		case 95:
			goto st424
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st424
			}
		case data[p] >= 65:
			goto st424
		}
		goto tr0
tr625:
//line ragel/datetime.rl:7
 pb = p 
	goto st424
tr623:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st424
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
//line ragel/parse_datetime.go:2122
		switch data[p] {
		case 32:
			goto tr607
		case 47:
			goto st424
		case 95:
			goto st424
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st424
			}
		case data[p] >= 65:
			goto st424
		}
		goto st0
tr593:
//line ragel/datetime.rl:7
 pb = p 
	goto st425
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
//line ragel/parse_datetime.go:2149
		switch data[p] {
		case 32:
			goto tr608
		case 43:
			goto tr609
		case 45:
			goto tr610
		case 47:
			goto tr611
		case 58:
			goto tr613
		case 65:
			goto tr614
		case 80:
			goto tr614
		case 90:
			goto tr615
		case 95:
			goto tr611
		case 97:
			goto tr616
		case 112:
			goto tr616
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st432
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr611
			}
		default:
			goto tr611
		}
		goto st0
tr608:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st426
tr626:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st426
tr681:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

	goto st426
tr655:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st426
tr664:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st426
tr672:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st426
tr692:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
//line ragel/parse_datetime.go:2247
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr592
		case 65:
			goto tr617
		case 66:
			goto tr597
		case 80:
			goto tr618
		case 95:
			goto tr592
		case 97:
			goto tr619
		case 112:
			goto tr619
		}
		switch {
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr592
			}
		case data[p] >= 67:
			goto tr592
		}
		goto st0
tr617:
//line ragel/datetime.rl:7
 pb = p 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line ragel/parse_datetime.go:2288
		switch data[p] {
		case 47:
			goto st15
		case 68:
			goto st427
		case 77:
			goto st428
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
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 47:
			goto st424
		case 95:
			goto st424
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st424
			}
		case data[p] >= 65:
			goto st424
		}
		goto st0
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		switch data[p] {
		case 32:
			goto tr620
		case 43:
			goto tr621
		case 45:
			goto tr622
		case 47:
			goto tr623
		case 90:
			goto tr624
		case 95:
			goto tr623
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr623
			}
		case data[p] >= 65:
			goto tr623
		}
		goto st0
tr620:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st429
tr634:
//line ragel/datetime.rl:33

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

	goto st429
tr641:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

	goto st429
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
//line ragel/parse_datetime.go:2422
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr592
		case 65:
			goto tr596
		case 66:
			goto tr597
		case 95:
			goto tr592
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr592
			}
		case data[p] >= 67:
			goto tr592
		}
		goto st0
tr596:
//line ragel/datetime.rl:7
 pb = p 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line ragel/parse_datetime.go:2457
		switch data[p] {
		case 47:
			goto st15
		case 68:
			goto st427
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
tr597:
//line ragel/datetime.rl:7
 pb = p 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line ragel/parse_datetime.go:2484
		switch data[p] {
		case 47:
			goto st15
		case 67:
			goto st430
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
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
		switch data[p] {
		case 47:
			goto st424
		case 95:
			goto st424
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st424
			}
		case data[p] >= 65:
			goto st424
		}
		goto st0
tr624:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st431
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
//line ragel/parse_datetime.go:2554
		switch data[p] {
		case 32:
			goto tr607
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr625
		case 95:
			goto tr625
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr625
			}
		case data[p] >= 65:
			goto tr625
		}
		goto st0
tr618:
//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr614:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr631:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr683:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr661:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr669:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr678:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr697:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line ragel/parse_datetime.go:2654
		switch data[p] {
		case 47:
			goto st15
		case 77:
			goto st428
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
tr619:
//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr616:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr633:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr684:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr663:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr671:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr680:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr699:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line ragel/parse_datetime.go:2750
		switch data[p] {
		case 47:
			goto st15
		case 95:
			goto st15
		case 109:
			goto st428
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
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
		switch data[p] {
		case 32:
			goto tr626
		case 43:
			goto tr627
		case 45:
			goto tr628
		case 47:
			goto tr629
		case 58:
			goto tr630
		case 65:
			goto tr631
		case 80:
			goto tr631
		case 90:
			goto tr632
		case 95:
			goto tr629
		case 97:
			goto tr633
		case 112:
			goto tr633
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st21
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr629
			}
		default:
			goto tr629
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= data[p] && data[p] <= 57 {
			goto st433
		}
		goto tr0
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 32:
			goto tr634
		case 43:
			goto tr635
		case 45:
			goto tr636
		case 46:
			goto tr637
		case 47:
			goto tr638
		case 90:
			goto tr640
		case 95:
			goto tr638
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr638
			}
		default:
			goto tr638
		}
		goto st0
tr637:
//line ragel/datetime.rl:33

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
//line ragel/parse_datetime.go:2875
		if 48 <= data[p] && data[p] <= 57 {
			goto tr57
		}
		goto tr0
tr57:
//line ragel/datetime.rl:7
 pb = p 
	goto st434
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
//line ragel/parse_datetime.go:2889
		switch data[p] {
		case 32:
			goto tr641
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 90:
			goto tr646
		case 95:
			goto tr644
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st435
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		switch data[p] {
		case 32:
			goto tr641
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 90:
			goto tr646
		case 95:
			goto tr644
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st436
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		switch data[p] {
		case 32:
			goto tr641
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 90:
			goto tr646
		case 95:
			goto tr644
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st437
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		switch data[p] {
		case 32:
			goto tr641
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 90:
			goto tr646
		case 95:
			goto tr644
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st438
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		switch data[p] {
		case 32:
			goto tr641
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 90:
			goto tr646
		case 95:
			goto tr644
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st439
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		switch data[p] {
		case 32:
			goto tr641
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 90:
			goto tr646
		case 95:
			goto tr644
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st440
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
		switch data[p] {
		case 32:
			goto tr641
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 90:
			goto tr646
		case 95:
			goto tr644
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st441
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		switch data[p] {
		case 32:
			goto tr641
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 90:
			goto tr646
		case 95:
			goto tr644
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st442
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		switch data[p] {
		case 32:
			goto tr641
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 90:
			goto tr646
		case 95:
			goto tr644
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		case data[p] >= 65:
			goto tr644
		}
		goto st0
tr724:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr615:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr632:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr640:
//line ragel/datetime.rl:33

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
	goto st443
tr646:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr662:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr670:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr679:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr698:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr588:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr706:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr712:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st443
tr718:
//line ragel/datetime.rl:19

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
//line ragel/parse_datetime.go:3318
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr654
		case 95:
			goto tr654
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr654
			}
		case data[p] >= 65:
			goto tr654
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if 48 <= data[p] && data[p] <= 57 {
			goto st444
		}
		goto tr0
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		switch data[p] {
		case 32:
			goto tr634
		case 43:
			goto tr635
		case 45:
			goto tr636
		case 46:
			goto tr637
		case 47:
			goto tr638
		case 90:
			goto tr640
		case 95:
			goto tr638
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr638
			}
		case data[p] >= 65:
			goto tr638
		}
		goto st0
tr613:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st24
tr630:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line ragel/parse_datetime.go:3396
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
	goto st445
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
//line ragel/parse_datetime.go:3415
		switch data[p] {
		case 32:
			goto tr655
		case 43:
			goto tr656
		case 45:
			goto tr657
		case 47:
			goto tr658
		case 58:
			goto tr660
		case 65:
			goto tr661
		case 80:
			goto tr661
		case 90:
			goto tr662
		case 95:
			goto tr658
		case 97:
			goto tr663
		case 112:
			goto tr663
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st446
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr658
			}
		default:
			goto tr658
		}
		goto st0
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		switch data[p] {
		case 32:
			goto tr664
		case 43:
			goto tr665
		case 45:
			goto tr666
		case 47:
			goto tr667
		case 58:
			goto tr668
		case 65:
			goto tr669
		case 80:
			goto tr669
		case 90:
			goto tr670
		case 95:
			goto tr667
		case 97:
			goto tr671
		case 112:
			goto tr671
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr667
			}
		case data[p] >= 66:
			goto tr667
		}
		goto st0
tr660:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st25
tr668:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line ragel/parse_datetime.go:3508
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr62
			}
		case data[p] >= 48:
			goto tr61
		}
		goto tr0
tr61:
//line ragel/datetime.rl:7
 pb = p 
	goto st447
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
//line ragel/parse_datetime.go:3527
		switch data[p] {
		case 32:
			goto tr672
		case 43:
			goto tr673
		case 45:
			goto tr674
		case 46:
			goto tr675
		case 47:
			goto tr676
		case 65:
			goto tr678
		case 80:
			goto tr678
		case 90:
			goto tr679
		case 95:
			goto tr676
		case 97:
			goto tr680
		case 112:
			goto tr680
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st457
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr676
			}
		default:
			goto tr676
		}
		goto st0
tr675:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st26
tr695:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line ragel/parse_datetime.go:3582
		if 48 <= data[p] && data[p] <= 57 {
			goto tr63
		}
		goto tr0
tr63:
//line ragel/datetime.rl:7
 pb = p 
	goto st448
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
//line ragel/parse_datetime.go:3596
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 65:
			goto tr683
		case 80:
			goto tr683
		case 90:
			goto tr646
		case 95:
			goto tr644
		case 97:
			goto tr684
		case 112:
			goto tr684
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st449
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 65:
			goto tr683
		case 80:
			goto tr683
		case 90:
			goto tr646
		case 95:
			goto tr644
		case 97:
			goto tr684
		case 112:
			goto tr684
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st450
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
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
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 65:
			goto tr683
		case 80:
			goto tr683
		case 90:
			goto tr646
		case 95:
			goto tr644
		case 97:
			goto tr684
		case 112:
			goto tr684
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st451
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 65:
			goto tr683
		case 80:
			goto tr683
		case 90:
			goto tr646
		case 95:
			goto tr644
		case 97:
			goto tr684
		case 112:
			goto tr684
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st452
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 65:
			goto tr683
		case 80:
			goto tr683
		case 90:
			goto tr646
		case 95:
			goto tr644
		case 97:
			goto tr684
		case 112:
			goto tr684
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st453
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 65:
			goto tr683
		case 80:
			goto tr683
		case 90:
			goto tr646
		case 95:
			goto tr644
		case 97:
			goto tr684
		case 112:
			goto tr684
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st454
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 65:
			goto tr683
		case 80:
			goto tr683
		case 90:
			goto tr646
		case 95:
			goto tr644
		case 97:
			goto tr684
		case 112:
			goto tr684
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st455
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 65:
			goto tr683
		case 80:
			goto tr683
		case 90:
			goto tr646
		case 95:
			goto tr644
		case 97:
			goto tr684
		case 112:
			goto tr684
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st456
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		default:
			goto tr644
		}
		goto st0
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		switch data[p] {
		case 32:
			goto tr681
		case 43:
			goto tr642
		case 45:
			goto tr643
		case 47:
			goto tr644
		case 65:
			goto tr683
		case 80:
			goto tr683
		case 90:
			goto tr646
		case 95:
			goto tr644
		case 97:
			goto tr684
		case 112:
			goto tr684
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr644
			}
		case data[p] >= 66:
			goto tr644
		}
		goto st0
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 32:
			goto tr692
		case 43:
			goto tr693
		case 45:
			goto tr694
		case 46:
			goto tr695
		case 47:
			goto tr696
		case 65:
			goto tr697
		case 80:
			goto tr697
		case 90:
			goto tr698
		case 95:
			goto tr696
		case 97:
			goto tr699
		case 112:
			goto tr699
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr696
			}
		case data[p] >= 66:
			goto tr696
		}
		goto st0
tr62:
//line ragel/datetime.rl:7
 pb = p 
	goto st458
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
//line ragel/parse_datetime.go:3995
		switch data[p] {
		case 32:
			goto tr672
		case 43:
			goto tr673
		case 45:
			goto tr674
		case 46:
			goto tr675
		case 47:
			goto tr676
		case 65:
			goto tr678
		case 80:
			goto tr678
		case 90:
			goto tr679
		case 95:
			goto tr676
		case 97:
			goto tr680
		case 112:
			goto tr680
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr676
			}
		case data[p] >= 66:
			goto tr676
		}
		goto st0
tr60:
//line ragel/datetime.rl:7
 pb = p 
	goto st459
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
//line ragel/parse_datetime.go:4038
		switch data[p] {
		case 32:
			goto tr655
		case 43:
			goto tr656
		case 45:
			goto tr657
		case 47:
			goto tr658
		case 58:
			goto tr660
		case 65:
			goto tr661
		case 80:
			goto tr661
		case 90:
			goto tr662
		case 95:
			goto tr658
		case 97:
			goto tr663
		case 112:
			goto tr663
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr658
			}
		case data[p] >= 66:
			goto tr658
		}
		goto st0
tr594:
//line ragel/datetime.rl:7
 pb = p 
	goto st460
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
//line ragel/parse_datetime.go:4081
		switch data[p] {
		case 32:
			goto tr608
		case 43:
			goto tr609
		case 45:
			goto tr610
		case 47:
			goto tr611
		case 58:
			goto tr613
		case 65:
			goto tr614
		case 80:
			goto tr614
		case 90:
			goto tr615
		case 95:
			goto tr611
		case 97:
			goto tr616
		case 112:
			goto tr616
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st432
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr611
				}
			case data[p] >= 66:
				goto tr611
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
tr595:
//line ragel/datetime.rl:7
 pb = p 
	goto st461
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
//line ragel/parse_datetime.go:4142
		switch data[p] {
		case 32:
			goto tr608
		case 43:
			goto tr609
		case 45:
			goto tr610
		case 47:
			goto tr611
		case 58:
			goto tr613
		case 65:
			goto tr614
		case 80:
			goto tr614
		case 90:
			goto tr615
		case 95:
			goto tr611
		case 97:
			goto tr616
		case 112:
			goto tr616
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st27
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr611
			}
		default:
			goto tr611
		}
		goto st0
tr723:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st462
tr587:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st462
tr705:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st462
tr711:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st462
tr717:
//line ragel/datetime.rl:19

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st462
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
//line ragel/parse_datetime.go:4241
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr654
		case 50:
			goto tr594
		case 95:
			goto tr654
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr593
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr654
				}
			case data[p] >= 65:
				goto tr654
			}
		default:
			goto tr595
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 45:
			goto tr65
		case 47:
			goto tr65
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st413
		}
		goto tr0
tr65:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st29
tr74:
//line ragel/datetime.rl:78
 st.Month = 4 ;
	goto st29
tr78:
//line ragel/datetime.rl:82
 st.Month = 8 ;
	goto st29
tr81:
//line ragel/datetime.rl:86
 st.Month = 12 ;
	goto st29
tr84:
//line ragel/datetime.rl:76
 st.Month = 2 ;
	goto st29
tr93:
//line ragel/datetime.rl:75
 st.Month = 1 ;
	goto st29
tr100:
//line ragel/datetime.rl:81
 st.Month = 7 ;
	goto st29
tr101:
//line ragel/datetime.rl:80
 st.Month = 6 ;
	goto st29
tr105:
//line ragel/datetime.rl:77
 st.Month = 3 ;
	goto st29
tr108:
//line ragel/datetime.rl:79
 st.Month = 5 ;
	goto st29
tr111:
//line ragel/datetime.rl:85
 st.Month = 11 ;
	goto st29
tr114:
//line ragel/datetime.rl:84
 st.Month = 10 ;
	goto st29
tr117:
//line ragel/datetime.rl:83
 st.Month = 9 ;
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line ragel/parse_datetime.go:4348
		switch data[p] {
		case 48:
			goto tr66
		case 51:
			goto tr68
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr69
			}
		case data[p] >= 49:
			goto tr67
		}
		goto tr0
tr66:
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr135:
//line ragel/datetime.rl:78
 st.Month = 4 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr141:
//line ragel/datetime.rl:82
 st.Month = 8 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr146:
//line ragel/datetime.rl:86
 st.Month = 12 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr151:
//line ragel/datetime.rl:76
 st.Month = 2 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr162:
//line ragel/datetime.rl:75
 st.Month = 1 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr171:
//line ragel/datetime.rl:81
 st.Month = 7 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr174:
//line ragel/datetime.rl:80
 st.Month = 6 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr180:
//line ragel/datetime.rl:77
 st.Month = 3 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr185:
//line ragel/datetime.rl:79
 st.Month = 5 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr190:
//line ragel/datetime.rl:85
 st.Month = 11 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr195:
//line ragel/datetime.rl:84
 st.Month = 10 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
tr200:
//line ragel/datetime.rl:83
 st.Month = 9 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line ragel/parse_datetime.go:4445
		if 49 <= data[p] && data[p] <= 57 {
			goto st463
		}
		goto tr0
tr69:
//line ragel/datetime.rl:7
 pb = p 
	goto st463
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
//line ragel/parse_datetime.go:4459
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr702
		case 45:
			goto tr703
		case 47:
			goto tr704
		case 84:
			goto tr705
		case 90:
			goto tr706
		case 95:
			goto tr704
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr704
			}
		case data[p] >= 65:
			goto tr704
		}
		goto st0
tr67:
//line ragel/datetime.rl:7
 pb = p 
	goto st464
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
//line ragel/parse_datetime.go:4494
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr702
		case 45:
			goto tr703
		case 47:
			goto tr704
		case 84:
			goto tr705
		case 90:
			goto tr706
		case 95:
			goto tr704
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st463
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr704
			}
		default:
			goto tr704
		}
		goto st0
tr68:
//line ragel/datetime.rl:7
 pb = p 
	goto st465
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
//line ragel/parse_datetime.go:4533
		switch data[p] {
		case 32:
			goto tr701
		case 43:
			goto tr702
		case 45:
			goto tr703
		case 47:
			goto tr704
		case 84:
			goto tr705
		case 90:
			goto tr706
		case 95:
			goto tr704
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 49 {
				goto st463
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr704
			}
		default:
			goto tr704
		}
		goto st0
tr28:
//line ragel/datetime.rl:7
 pb = p 
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line ragel/parse_datetime.go:4572
		switch data[p] {
		case 45:
			goto tr65
		case 47:
			goto tr65
		}
		switch {
		case data[p] > 50:
			if 51 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] >= 48:
			goto st28
		}
		goto tr0
tr29:
//line ragel/datetime.rl:7
 pb = p 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line ragel/parse_datetime.go:4597
		switch data[p] {
		case 45:
			goto tr65
		case 47:
			goto tr65
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto tr0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 112:
			goto st34
		case 117:
			goto st38
		}
		goto tr0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 114 {
			goto st35
		}
		goto tr0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 45:
			goto tr74
		case 47:
			goto tr74
		case 105:
			goto st36
		}
		goto tr0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 108 {
			goto st37
		}
		goto tr0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 45:
			goto tr74
		case 47:
			goto tr74
		}
		goto tr0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 103 {
			goto st39
		}
		goto tr0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 45:
			goto tr78
		case 47:
			goto tr78
		}
		goto tr0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 101 {
			goto st41
		}
		goto tr0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 99 {
			goto st42
		}
		goto tr0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 45:
			goto tr81
		case 47:
			goto tr81
		}
		goto tr0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 101 {
			goto st44
		}
		goto tr0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		if data[p] == 98 {
			goto st45
		}
		goto tr0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 45:
			goto tr84
		case 47:
			goto tr84
		case 114:
			goto st46
		}
		goto tr0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 117 {
			goto st47
		}
		goto tr0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 97 {
			goto st48
		}
		goto tr0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 114 {
			goto st49
		}
		goto tr0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 121 {
			goto st50
		}
		goto tr0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 45:
			goto tr84
		case 47:
			goto tr84
		}
		goto tr0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 97:
			goto st52
		case 117:
			goto st58
		}
		goto tr0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 110 {
			goto st53
		}
		goto tr0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 45:
			goto tr93
		case 47:
			goto tr93
		case 117:
			goto st54
		}
		goto tr0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if data[p] == 97 {
			goto st55
		}
		goto tr0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 114 {
			goto st56
		}
		goto tr0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 121 {
			goto st57
		}
		goto tr0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 45:
			goto tr93
		case 47:
			goto tr93
		}
		goto tr0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 108:
			goto st59
		case 110:
			goto st60
		}
		goto tr0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 45:
			goto tr100
		case 47:
			goto tr100
		}
		goto tr0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 45:
			goto tr101
		case 47:
			goto tr101
		}
		goto tr0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 97 {
			goto st62
		}
		goto tr0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 114:
			goto st63
		case 121:
			goto st66
		}
		goto tr0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 45:
			goto tr105
		case 47:
			goto tr105
		case 99:
			goto st64
		}
		goto tr0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 104 {
			goto st65
		}
		goto tr0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 45:
			goto tr105
		case 47:
			goto tr105
		}
		goto tr0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 45:
			goto tr108
		case 47:
			goto tr108
		}
		goto tr0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		if data[p] == 111 {
			goto st68
		}
		goto tr0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if data[p] == 118 {
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
			goto tr111
		case 47:
			goto tr111
		}
		goto tr0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		if data[p] == 99 {
			goto st71
		}
		goto tr0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		if data[p] == 116 {
			goto st72
		}
		goto tr0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 45:
			goto tr114
		case 47:
			goto tr114
		}
		goto tr0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 101 {
			goto st74
		}
		goto tr0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		if data[p] == 112 {
			goto st75
		}
		goto tr0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 45:
			goto tr117
		case 47:
			goto tr117
		}
		goto tr0
tr15:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line ragel/parse_datetime.go:5074
		if 48 <= data[p] && data[p] <= 57 {
			goto tr118
		}
		goto tr0
tr118:
//line ragel/datetime.rl:7
 pb = p 
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line ragel/parse_datetime.go:5088
		if 48 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto tr0
tr16:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line ragel/parse_datetime.go:5106
		if data[p] == 48 {
			goto st8
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st79
		}
		goto tr0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 48:
			goto tr120
		case 51:
			goto tr122
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto st413
			}
		case data[p] >= 49:
			goto tr121
		}
		goto tr0
tr120:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st466
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
//line ragel/parse_datetime.go:5147
		switch data[p] {
		case 32:
			goto tr582
		case 43:
			goto tr584
		case 45:
			goto tr585
		case 47:
			goto tr586
		case 84:
			goto tr587
		case 90:
			goto tr588
		case 95:
			goto tr586
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto st463
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr586
			}
		default:
			goto tr586
		}
		goto st0
tr121:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st467
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
//line ragel/parse_datetime.go:5190
		switch data[p] {
		case 32:
			goto tr582
		case 43:
			goto tr584
		case 45:
			goto tr585
		case 47:
			goto tr586
		case 84:
			goto tr587
		case 90:
			goto tr588
		case 95:
			goto tr586
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st463
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr586
			}
		default:
			goto tr586
		}
		goto st0
tr122:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
//line ragel/parse_datetime.go:5233
		switch data[p] {
		case 32:
			goto tr582
		case 43:
			goto tr584
		case 45:
			goto tr585
		case 47:
			goto tr586
		case 84:
			goto tr587
		case 90:
			goto tr588
		case 95:
			goto tr586
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 49 {
				goto st463
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr586
			}
		default:
			goto tr586
		}
		goto st0
tr17:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line ragel/parse_datetime.go:5276
		switch data[p] {
		case 48:
			goto tr123
		case 51:
			goto tr125
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] >= 49:
			goto tr124
		}
		goto tr0
tr123:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line ragel/parse_datetime.go:5305
		switch data[p] {
		case 48:
			goto tr120
		case 51:
			goto tr127
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto st471
			}
		case data[p] >= 49:
			goto tr126
		}
		goto tr0
tr126:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st469
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
//line ragel/parse_datetime.go:5334
		switch data[p] {
		case 32:
			goto tr707
		case 43:
			goto tr708
		case 45:
			goto tr709
		case 47:
			goto tr710
		case 84:
			goto tr711
		case 90:
			goto tr712
		case 95:
			goto tr710
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st463
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr710
			}
		default:
			goto tr710
		}
		goto st0
tr127:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st470
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
//line ragel/parse_datetime.go:5377
		switch data[p] {
		case 32:
			goto tr707
		case 43:
			goto tr708
		case 45:
			goto tr709
		case 47:
			goto tr710
		case 84:
			goto tr711
		case 90:
			goto tr712
		case 95:
			goto tr710
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 49 {
				goto st463
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr710
			}
		default:
			goto tr710
		}
		goto st0
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		switch data[p] {
		case 32:
			goto tr707
		case 43:
			goto tr708
		case 45:
			goto tr709
		case 47:
			goto tr710
		case 84:
			goto tr711
		case 90:
			goto tr712
		case 95:
			goto tr710
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr710
			}
		case data[p] >= 65:
			goto tr710
		}
		goto st0
tr124:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line ragel/parse_datetime.go:5450
		switch data[p] {
		case 48:
			goto tr129
		case 51:
			goto tr127
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto st471
			}
		case data[p] >= 49:
			goto tr126
		}
		goto tr0
tr129:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st472
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
//line ragel/parse_datetime.go:5479
		switch data[p] {
		case 32:
			goto tr707
		case 43:
			goto tr708
		case 45:
			goto tr709
		case 47:
			goto tr710
		case 84:
			goto tr711
		case 90:
			goto tr712
		case 95:
			goto tr710
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto st463
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr710
			}
		default:
			goto tr710
		}
		goto st0
tr125:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line ragel/parse_datetime.go:5522
		switch {
		case data[p] > 49:
			if 50 <= data[p] && data[p] <= 57 {
				goto st413
			}
		case data[p] >= 48:
			goto st471
		}
		goto tr0
tr18:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line ragel/parse_datetime.go:5545
		switch data[p] {
		case 48:
			goto tr130
		case 51:
			goto tr125
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] >= 49:
			goto tr131
		}
		goto tr0
tr130:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line ragel/parse_datetime.go:5574
		if data[p] == 48 {
			goto st413
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st471
		}
		goto tr0
tr131:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

//line ragel/datetime.rl:7
 pb = p 
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line ragel/parse_datetime.go:5595
		if 48 <= data[p] && data[p] <= 57 {
			goto st471
		}
		goto tr0
tr19:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line ragel/parse_datetime.go:5611
		switch data[p] {
		case 112:
			goto st88
		case 117:
			goto st94
		}
		goto tr0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		if data[p] == 114 {
			goto st89
		}
		goto tr0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 48:
			goto tr135
		case 51:
			goto tr137
		case 105:
			goto st92
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr136
		}
		goto tr0
tr136:
//line ragel/datetime.rl:78
 st.Month = 4 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr142:
//line ragel/datetime.rl:82
 st.Month = 8 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr147:
//line ragel/datetime.rl:86
 st.Month = 12 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr152:
//line ragel/datetime.rl:76
 st.Month = 2 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr163:
//line ragel/datetime.rl:75
 st.Month = 1 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr172:
//line ragel/datetime.rl:81
 st.Month = 7 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr175:
//line ragel/datetime.rl:80
 st.Month = 6 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr181:
//line ragel/datetime.rl:77
 st.Month = 3 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr186:
//line ragel/datetime.rl:79
 st.Month = 5 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr191:
//line ragel/datetime.rl:85
 st.Month = 11 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr196:
//line ragel/datetime.rl:84
 st.Month = 10 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
tr201:
//line ragel/datetime.rl:83
 st.Month = 9 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line ragel/parse_datetime.go:5722
		if 48 <= data[p] && data[p] <= 57 {
			goto st463
		}
		goto tr0
tr137:
//line ragel/datetime.rl:78
 st.Month = 4 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr143:
//line ragel/datetime.rl:82
 st.Month = 8 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr148:
//line ragel/datetime.rl:86
 st.Month = 12 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr153:
//line ragel/datetime.rl:76
 st.Month = 2 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr164:
//line ragel/datetime.rl:75
 st.Month = 1 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr173:
//line ragel/datetime.rl:81
 st.Month = 7 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr176:
//line ragel/datetime.rl:80
 st.Month = 6 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr182:
//line ragel/datetime.rl:77
 st.Month = 3 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr187:
//line ragel/datetime.rl:79
 st.Month = 5 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr192:
//line ragel/datetime.rl:85
 st.Month = 11 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr197:
//line ragel/datetime.rl:84
 st.Month = 10 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
tr202:
//line ragel/datetime.rl:83
 st.Month = 9 ;
//line ragel/datetime.rl:7
 pb = p 
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line ragel/parse_datetime.go:5804
		if 48 <= data[p] && data[p] <= 49 {
			goto st463
		}
		goto tr0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		if data[p] == 108 {
			goto st93
		}
		goto tr0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 48:
			goto tr135
		case 51:
			goto tr137
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr136
		}
		goto tr0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		if data[p] == 103 {
			goto st95
		}
		goto tr0
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 48:
			goto tr141
		case 51:
			goto tr143
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr142
		}
		goto tr0
tr20:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line ragel/parse_datetime.go:5868
		if data[p] == 101 {
			goto st97
		}
		goto tr0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		if data[p] == 99 {
			goto st98
		}
		goto tr0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 48:
			goto tr146
		case 51:
			goto tr148
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr147
		}
		goto tr0
tr21:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line ragel/parse_datetime.go:5908
		if data[p] == 101 {
			goto st100
		}
		goto tr0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		if data[p] == 98 {
			goto st101
		}
		goto tr0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 48:
			goto tr151
		case 51:
			goto tr153
		case 114:
			goto st102
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr152
		}
		goto tr0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		if data[p] == 117 {
			goto st103
		}
		goto tr0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		if data[p] == 97 {
			goto st104
		}
		goto tr0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		if data[p] == 114 {
			goto st105
		}
		goto tr0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		if data[p] == 121 {
			goto st106
		}
		goto tr0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 48:
			goto tr151
		case 51:
			goto tr153
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr152
		}
		goto tr0
tr22:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line ragel/parse_datetime.go:6001
		switch data[p] {
		case 97:
			goto st108
		case 117:
			goto st114
		}
		goto tr0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		if data[p] == 110 {
			goto st109
		}
		goto tr0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 48:
			goto tr162
		case 51:
			goto tr164
		case 117:
			goto st110
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr163
		}
		goto tr0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		if data[p] == 97 {
			goto st111
		}
		goto tr0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		if data[p] == 114 {
			goto st112
		}
		goto tr0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		if data[p] == 121 {
			goto st113
		}
		goto tr0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 48:
			goto tr162
		case 51:
			goto tr164
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr163
		}
		goto tr0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 108:
			goto st115
		case 110:
			goto st116
		}
		goto tr0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 48:
			goto tr171
		case 51:
			goto tr173
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr172
		}
		goto tr0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 48:
			goto tr174
		case 51:
			goto tr176
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr175
		}
		goto tr0
tr23:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st117
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
//line ragel/parse_datetime.go:6130
		if data[p] == 97 {
			goto st118
		}
		goto tr0
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 114:
			goto st119
		case 121:
			goto st122
		}
		goto tr0
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 48:
			goto tr180
		case 51:
			goto tr182
		case 99:
			goto st120
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr181
		}
		goto tr0
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		if data[p] == 104 {
			goto st121
		}
		goto tr0
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 48:
			goto tr180
		case 51:
			goto tr182
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr181
		}
		goto tr0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 48:
			goto tr185
		case 51:
			goto tr187
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr186
		}
		goto tr0
tr24:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
//line ragel/parse_datetime.go:6214
		if data[p] == 111 {
			goto st124
		}
		goto tr0
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		if data[p] == 118 {
			goto st125
		}
		goto tr0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 48:
			goto tr190
		case 51:
			goto tr192
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr191
		}
		goto tr0
tr25:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line ragel/parse_datetime.go:6254
		if data[p] == 99 {
			goto st127
		}
		goto tr0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		if data[p] == 116 {
			goto st128
		}
		goto tr0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 48:
			goto tr195
		case 51:
			goto tr197
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr196
		}
		goto tr0
tr26:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st129
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
//line ragel/parse_datetime.go:6294
		if data[p] == 101 {
			goto st130
		}
		goto tr0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		if data[p] == 112 {
			goto st131
		}
		goto tr0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 48:
			goto tr200
		case 51:
			goto tr202
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr201
		}
		goto tr0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		if data[p] == 32 {
			goto tr203
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto tr0
tr203:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line ragel/parse_datetime.go:6352
		switch data[p] {
		case 65:
			goto st134
		case 68:
			goto st143
		case 70:
			goto st146
		case 74:
			goto st154
		case 77:
			goto st164
		case 78:
			goto st170
		case 79:
			goto st173
		case 83:
			goto st176
		}
		goto tr0
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 112:
			goto st135
		case 117:
			goto st141
		}
		goto tr0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		if data[p] == 114 {
			goto st136
		}
		goto tr0
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 32:
			goto tr215
		case 105:
			goto st139
		}
		goto tr0
tr215:
//line ragel/datetime.rl:78
 st.Month = 4 ;
	goto st137
tr221:
//line ragel/datetime.rl:82
 st.Month = 8 ;
	goto st137
tr224:
//line ragel/datetime.rl:86
 st.Month = 12 ;
	goto st137
tr227:
//line ragel/datetime.rl:76
 st.Month = 2 ;
	goto st137
tr236:
//line ragel/datetime.rl:75
 st.Month = 1 ;
	goto st137
tr243:
//line ragel/datetime.rl:81
 st.Month = 7 ;
	goto st137
tr244:
//line ragel/datetime.rl:80
 st.Month = 6 ;
	goto st137
tr248:
//line ragel/datetime.rl:77
 st.Month = 3 ;
	goto st137
tr251:
//line ragel/datetime.rl:79
 st.Month = 5 ;
	goto st137
tr254:
//line ragel/datetime.rl:85
 st.Month = 11 ;
	goto st137
tr257:
//line ragel/datetime.rl:84
 st.Month = 10 ;
	goto st137
tr260:
//line ragel/datetime.rl:83
 st.Month = 9 ;
	goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
//line ragel/parse_datetime.go:6458
		if 48 <= data[p] && data[p] <= 57 {
			goto tr217
		}
		goto tr0
tr217:
//line ragel/datetime.rl:7
 pb = p 
	goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
//line ragel/parse_datetime.go:6472
		if 48 <= data[p] && data[p] <= 57 {
			goto st473
		}
		goto tr0
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		switch data[p] {
		case 32:
			goto tr713
		case 43:
			goto tr714
		case 45:
			goto tr715
		case 47:
			goto tr716
		case 84:
			goto tr717
		case 90:
			goto tr718
		case 95:
			goto tr716
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr716
			}
		case data[p] >= 65:
			goto tr716
		}
		goto st0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		if data[p] == 108 {
			goto st140
		}
		goto tr0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		if data[p] == 32 {
			goto tr215
		}
		goto tr0
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		if data[p] == 103 {
			goto st142
		}
		goto tr0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		if data[p] == 32 {
			goto tr221
		}
		goto tr0
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		if data[p] == 101 {
			goto st144
		}
		goto tr0
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		if data[p] == 99 {
			goto st145
		}
		goto tr0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		if data[p] == 32 {
			goto tr224
		}
		goto tr0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		if data[p] == 101 {
			goto st147
		}
		goto tr0
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		if data[p] == 98 {
			goto st148
		}
		goto tr0
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 32:
			goto tr227
		case 114:
			goto st149
		}
		goto tr0
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		if data[p] == 117 {
			goto st150
		}
		goto tr0
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		if data[p] == 97 {
			goto st151
		}
		goto tr0
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		if data[p] == 114 {
			goto st152
		}
		goto tr0
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		if data[p] == 121 {
			goto st153
		}
		goto tr0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		if data[p] == 32 {
			goto tr227
		}
		goto tr0
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 97:
			goto st155
		case 117:
			goto st161
		}
		goto tr0
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		if data[p] == 110 {
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
			goto tr236
		case 117:
			goto st157
		}
		goto tr0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		if data[p] == 97 {
			goto st158
		}
		goto tr0
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		if data[p] == 114 {
			goto st159
		}
		goto tr0
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		if data[p] == 121 {
			goto st160
		}
		goto tr0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		if data[p] == 32 {
			goto tr236
		}
		goto tr0
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 108:
			goto st162
		case 110:
			goto st163
		}
		goto tr0
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		if data[p] == 32 {
			goto tr243
		}
		goto tr0
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		if data[p] == 32 {
			goto tr244
		}
		goto tr0
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		if data[p] == 97 {
			goto st165
		}
		goto tr0
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch data[p] {
		case 114:
			goto st166
		case 121:
			goto st169
		}
		goto tr0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 32:
			goto tr248
		case 99:
			goto st167
		}
		goto tr0
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		if data[p] == 104 {
			goto st168
		}
		goto tr0
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		if data[p] == 32 {
			goto tr248
		}
		goto tr0
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		if data[p] == 32 {
			goto tr251
		}
		goto tr0
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		if data[p] == 111 {
			goto st171
		}
		goto tr0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		if data[p] == 118 {
			goto st172
		}
		goto tr0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		if data[p] == 32 {
			goto tr254
		}
		goto tr0
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		if data[p] == 99 {
			goto st174
		}
		goto tr0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		if data[p] == 116 {
			goto st175
		}
		goto tr0
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		if data[p] == 32 {
			goto tr257
		}
		goto tr0
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		if data[p] == 101 {
			goto st177
		}
		goto tr0
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		if data[p] == 112 {
			goto st178
		}
		goto tr0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		if data[p] == 32 {
			goto tr260
		}
		goto tr0
tr2:
//line ragel/datetime.rl:7
 pb = p 
	goto st179
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
//line ragel/parse_datetime.go:6894
		if data[p] == 32 {
			goto tr203
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st132
		}
		goto tr0
tr3:
//line ragel/datetime.rl:7
 pb = p 
	goto st180
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
//line ragel/parse_datetime.go:6911
		if data[p] == 32 {
			goto tr203
		}
		switch {
		case data[p] > 49:
			if 50 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] >= 48:
			goto st132
		}
		goto tr0
tr4:
//line ragel/datetime.rl:7
 pb = p 
	goto st181
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
//line ragel/parse_datetime.go:6933
		if data[p] == 32 {
			goto tr203
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto tr0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		if data[p] == 114 {
			goto st183
		}
		goto tr0
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		if data[p] == 105 {
			goto st184
		}
		goto tr0
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		switch data[p] {
		case 32:
			goto st185
		case 44:
			goto st295
		case 100:
			goto st393
		}
		goto tr0
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		switch data[p] {
		case 65:
			goto st186
		case 68:
			goto st259
		case 70:
			goto st262
		case 74:
			goto st270
		case 77:
			goto st280
		case 78:
			goto st286
		case 79:
			goto st289
		case 83:
			goto st292
		}
		goto tr0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		switch data[p] {
		case 112:
			goto st187
		case 117:
			goto st257
		}
		goto tr0
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		if data[p] == 114 {
			goto st188
		}
		goto tr0
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch data[p] {
		case 32:
			goto tr277
		case 105:
			goto st255
		}
		goto tr0
tr277:
//line ragel/datetime.rl:78
 st.Month = 4 ;
	goto st189
tr417:
//line ragel/datetime.rl:82
 st.Month = 8 ;
	goto st189
tr420:
//line ragel/datetime.rl:86
 st.Month = 12 ;
	goto st189
tr423:
//line ragel/datetime.rl:76
 st.Month = 2 ;
	goto st189
tr432:
//line ragel/datetime.rl:75
 st.Month = 1 ;
	goto st189
tr439:
//line ragel/datetime.rl:81
 st.Month = 7 ;
	goto st189
tr440:
//line ragel/datetime.rl:80
 st.Month = 6 ;
	goto st189
tr444:
//line ragel/datetime.rl:77
 st.Month = 3 ;
	goto st189
tr447:
//line ragel/datetime.rl:79
 st.Month = 5 ;
	goto st189
tr450:
//line ragel/datetime.rl:85
 st.Month = 11 ;
	goto st189
tr453:
//line ragel/datetime.rl:84
 st.Month = 10 ;
	goto st189
tr456:
//line ragel/datetime.rl:83
 st.Month = 9 ;
	goto st189
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
//line ragel/parse_datetime.go:7083
		switch data[p] {
		case 48:
			goto tr279
		case 51:
			goto tr281
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr280
		}
		goto tr0
tr279:
//line ragel/datetime.rl:7
 pb = p 
	goto st190
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
//line ragel/parse_datetime.go:7103
		if 49 <= data[p] && data[p] <= 57 {
			goto st191
		}
		goto tr0
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		if data[p] == 32 {
			goto tr283
		}
		goto tr0
tr283:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st192
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
//line ragel/parse_datetime.go:7134
		if data[p] == 50 {
			goto tr285
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr286
			}
		case data[p] >= 48:
			goto tr284
		}
		goto tr0
tr284:
//line ragel/datetime.rl:7
 pb = p 
	goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
//line ragel/parse_datetime.go:7156
		switch data[p] {
		case 32:
			goto tr287
		case 43:
			goto tr288
		case 45:
			goto tr289
		case 47:
			goto tr290
		case 58:
			goto tr292
		case 65:
			goto tr293
		case 80:
			goto tr293
		case 90:
			goto tr294
		case 95:
			goto tr290
		case 97:
			goto tr295
		case 112:
			goto tr295
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st216
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr290
			}
		default:
			goto tr290
		}
		goto tr0
tr287:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st194
tr331:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st194
tr395:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

	goto st194
tr366:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st194
tr375:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st194
tr385:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st194
tr406:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
//line ragel/parse_datetime.go:7254
		switch data[p] {
		case 32:
			goto st195
		case 43:
			goto st199
		case 45:
			goto st207
		case 47:
			goto tr299
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 95:
			goto tr299
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr300
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr299
			}
		default:
			goto tr299
		}
		goto tr0
tr309:
//line ragel/datetime.rl:155
 
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
        case 1:{st.Offset_hour, _ = strconv.Atoi(data[pb:pb+1])}
        case 2:{st.Offset_hour, _ = strconv.Atoi(data[pb:pb+2])}
        case 3:{ 
            num , _ := strconv.Atoi(data[pb:pb+3])
            st.Offset_hour = num/100
            st.Offset_minute = num%100
            if st.Offset_minute >=60 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        case 4:{ 
            num := parse_digits(data[pb:pb+4])
            st.Offset_hour = num/100
            st.Offset_minute = num%100
            if st.Offset_minute >=60 || st.Offset_hour>=15 {
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
	goto st195
tr313:
//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st195
tr316:
//line ragel/datetime.rl:145

    switch p - pb {
        case 1: st.Offset_minute, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Offset_minute, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st195
tr323:
//line ragel/datetime.rl:193

    st.Zone_name_or_abbrev = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st195
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
//line ragel/parse_datetime.go:7362
		if 48 <= data[p] && data[p] <= 57 {
			goto tr300
		}
		goto tr0
tr300:
//line ragel/datetime.rl:7
 pb = p 
	goto st196
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
//line ragel/parse_datetime.go:7376
		if 48 <= data[p] && data[p] <= 57 {
			goto st197
		}
		goto tr0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		if 48 <= data[p] && data[p] <= 57 {
			goto st198
		}
		goto tr0
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		if 48 <= data[p] && data[p] <= 57 {
			goto st474
		}
		goto tr0
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		goto st0
tr288:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st199
tr326:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st199
tr332:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st199
tr342:
//line ragel/datetime.rl:33

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

	goto st199
tr350:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

	goto st199
tr367:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st199
tr376:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st199
tr386:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st199
tr407:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
//line ragel/parse_datetime.go:7507
		if data[p] == 50 {
			goto tr307
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr308
			}
		case data[p] >= 48:
			goto tr306
		}
		goto tr0
tr306:
//line ragel/datetime.rl:7
 pb = p 
	goto st200
tr318:
//line ragel/datetime.rl:133
 st.Negative_offset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st200
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
//line ragel/parse_datetime.go:7535
		switch data[p] {
		case 32:
			goto tr309
		case 58:
			goto tr311
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st201
		}
		goto tr0
tr308:
//line ragel/datetime.rl:7
 pb = p 
	goto st201
tr320:
//line ragel/datetime.rl:133
 st.Negative_offset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st201
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
//line ragel/parse_datetime.go:7561
		switch data[p] {
		case 32:
			goto tr309
		case 58:
			goto tr311
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st202
		}
		goto tr0
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		if data[p] == 32 {
			goto tr309
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st202
		}
		goto tr0
tr311:
//line ragel/datetime.rl:135

    switch p - pb {
        case 1: st.Offset_hour, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Offset_hour, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st203
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
//line ragel/parse_datetime.go:7601
		if data[p] == 32 {
			goto tr313
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr315
			}
		case data[p] >= 48:
			goto tr314
		}
		goto tr0
tr314:
//line ragel/datetime.rl:7
 pb = p 
	goto st204
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
//line ragel/parse_datetime.go:7623
		if data[p] == 32 {
			goto tr316
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st205
		}
		goto tr0
tr315:
//line ragel/datetime.rl:7
 pb = p 
	goto st205
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
//line ragel/parse_datetime.go:7640
		if data[p] == 32 {
			goto tr316
		}
		goto tr0
tr307:
//line ragel/datetime.rl:7
 pb = p 
	goto st206
tr319:
//line ragel/datetime.rl:133
 st.Negative_offset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
//line ragel/parse_datetime.go:7660
		switch data[p] {
		case 32:
			goto tr309
		case 58:
			goto tr311
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st202
			}
		case data[p] >= 48:
			goto st201
		}
		goto tr0
tr289:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st207
tr327:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st207
tr333:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st207
tr343:
//line ragel/datetime.rl:33

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

	goto st207
tr351:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

	goto st207
tr368:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st207
tr377:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st207
tr387:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st207
tr408:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st207
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
//line ragel/parse_datetime.go:7778
		if data[p] == 50 {
			goto tr319
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr320
			}
		case data[p] >= 48:
			goto tr318
		}
		goto tr0
tr299:
//line ragel/datetime.rl:7
 pb = p 
	goto st208
tr290:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st208
tr334:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st208
tr345:
//line ragel/datetime.rl:33

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
	goto st208
tr352:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st208
tr369:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st208
tr378:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st208
tr389:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st208
tr410:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st208
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
//line ragel/parse_datetime.go:7888
		switch data[p] {
		case 47:
			goto st209
		case 95:
			goto st209
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st209
			}
		case data[p] >= 65:
			goto st209
		}
		goto tr0
tr362:
//line ragel/datetime.rl:7
 pb = p 
	goto st209
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
//line ragel/parse_datetime.go:7913
		switch data[p] {
		case 47:
			goto st210
		case 95:
			goto st210
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st210
			}
		case data[p] >= 65:
			goto st210
		}
		goto tr0
tr330:
//line ragel/datetime.rl:7
 pb = p 
	goto st210
tr328:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st210
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
//line ragel/parse_datetime.go:7965
		switch data[p] {
		case 32:
			goto tr323
		case 47:
			goto st210
		case 95:
			goto st210
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st210
			}
		case data[p] >= 65:
			goto st210
		}
		goto tr0
tr301:
//line ragel/datetime.rl:7
 pb = p 
	goto st211
tr293:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st211
tr337:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st211
tr397:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st211
tr372:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st211
tr380:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st211
tr391:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st211
tr411:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st211
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
//line ragel/parse_datetime.go:8061
		switch data[p] {
		case 47:
			goto st209
		case 77:
			goto st212
		case 95:
			goto st209
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st209
			}
		case data[p] >= 65:
			goto st209
		}
		goto tr0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		switch data[p] {
		case 32:
			goto tr325
		case 43:
			goto tr326
		case 45:
			goto tr327
		case 47:
			goto tr328
		case 90:
			goto tr329
		case 95:
			goto tr328
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr328
			}
		case data[p] >= 65:
			goto tr328
		}
		goto tr0
tr325:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

	goto st213
tr341:
//line ragel/datetime.rl:33

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

	goto st213
tr349:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

	goto st213
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
//line ragel/parse_datetime.go:8173
		switch data[p] {
		case 32:
			goto st195
		case 43:
			goto st199
		case 45:
			goto st207
		case 47:
			goto tr299
		case 95:
			goto tr299
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr300
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr299
			}
		default:
			goto tr299
		}
		goto tr0
tr329:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st214
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
//line ragel/parse_datetime.go:8231
		switch data[p] {
		case 32:
			goto tr323
		case 43:
			goto st199
		case 45:
			goto st207
		case 47:
			goto tr330
		case 95:
			goto tr330
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr330
			}
		case data[p] >= 65:
			goto tr330
		}
		goto tr0
tr302:
//line ragel/datetime.rl:7
 pb = p 
	goto st215
tr295:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st215
tr339:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st215
tr398:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st215
tr374:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st215
tr382:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st215
tr393:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st215
tr413:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st215
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
//line ragel/parse_datetime.go:8331
		switch data[p] {
		case 47:
			goto st209
		case 95:
			goto st209
		case 109:
			goto st212
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st209
			}
		case data[p] >= 65:
			goto st209
		}
		goto tr0
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 32:
			goto tr331
		case 43:
			goto tr332
		case 45:
			goto tr333
		case 47:
			goto tr334
		case 58:
			goto tr336
		case 65:
			goto tr337
		case 80:
			goto tr337
		case 90:
			goto tr338
		case 95:
			goto tr334
		case 97:
			goto tr339
		case 112:
			goto tr339
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st217
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr334
			}
		default:
			goto tr334
		}
		goto tr0
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		if 48 <= data[p] && data[p] <= 57 {
			goto st218
		}
		goto tr0
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 32:
			goto tr341
		case 43:
			goto tr342
		case 45:
			goto tr343
		case 46:
			goto tr344
		case 47:
			goto tr345
		case 90:
			goto tr347
		case 95:
			goto tr345
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st230
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr345
			}
		default:
			goto tr345
		}
		goto tr0
tr344:
//line ragel/datetime.rl:33

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

	goto st219
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
//line ragel/parse_datetime.go:8456
		if 48 <= data[p] && data[p] <= 57 {
			goto tr348
		}
		goto tr0
tr348:
//line ragel/datetime.rl:7
 pb = p 
	goto st220
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
//line ragel/parse_datetime.go:8470
		switch data[p] {
		case 32:
			goto tr349
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 90:
			goto tr354
		case 95:
			goto tr352
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st221
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		switch data[p] {
		case 32:
			goto tr349
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 90:
			goto tr354
		case 95:
			goto tr352
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st222
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 32:
			goto tr349
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 90:
			goto tr354
		case 95:
			goto tr352
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st223
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 32:
			goto tr349
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 90:
			goto tr354
		case 95:
			goto tr352
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st224
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 32:
			goto tr349
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 90:
			goto tr354
		case 95:
			goto tr352
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st225
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 32:
			goto tr349
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 90:
			goto tr354
		case 95:
			goto tr352
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st226
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 32:
			goto tr349
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 90:
			goto tr354
		case 95:
			goto tr352
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st227
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		switch data[p] {
		case 32:
			goto tr349
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 90:
			goto tr354
		case 95:
			goto tr352
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st228
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		switch data[p] {
		case 32:
			goto tr349
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 90:
			goto tr354
		case 95:
			goto tr352
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		case data[p] >= 65:
			goto tr352
		}
		goto tr0
tr294:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st229
tr338:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st229
tr347:
//line ragel/datetime.rl:33

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
	goto st229
tr354:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st229
tr373:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st229
tr381:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st229
tr392:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st229
tr412:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line ragel/parse_datetime.go:8843
		switch data[p] {
		case 32:
			goto st195
		case 43:
			goto st199
		case 45:
			goto st207
		case 47:
			goto tr362
		case 95:
			goto tr362
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr362
			}
		case data[p] >= 65:
			goto tr362
		}
		goto tr0
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		if 48 <= data[p] && data[p] <= 57 {
			goto st231
		}
		goto tr0
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		switch data[p] {
		case 32:
			goto tr341
		case 43:
			goto tr342
		case 45:
			goto tr343
		case 46:
			goto tr344
		case 47:
			goto tr345
		case 90:
			goto tr347
		case 95:
			goto tr345
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr345
			}
		case data[p] >= 65:
			goto tr345
		}
		goto tr0
tr292:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st232
tr336:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st232
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
//line ragel/parse_datetime.go:8921
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr365
			}
		case data[p] >= 48:
			goto tr364
		}
		goto tr0
tr364:
//line ragel/datetime.rl:7
 pb = p 
	goto st233
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
//line ragel/parse_datetime.go:8940
		switch data[p] {
		case 32:
			goto tr366
		case 43:
			goto tr367
		case 45:
			goto tr368
		case 47:
			goto tr369
		case 58:
			goto tr371
		case 65:
			goto tr372
		case 80:
			goto tr372
		case 90:
			goto tr373
		case 95:
			goto tr369
		case 97:
			goto tr374
		case 112:
			goto tr374
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st234
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr369
			}
		default:
			goto tr369
		}
		goto tr0
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		switch data[p] {
		case 32:
			goto tr375
		case 43:
			goto tr376
		case 45:
			goto tr377
		case 47:
			goto tr378
		case 58:
			goto tr379
		case 65:
			goto tr380
		case 80:
			goto tr380
		case 90:
			goto tr381
		case 95:
			goto tr378
		case 97:
			goto tr382
		case 112:
			goto tr382
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr378
			}
		case data[p] >= 66:
			goto tr378
		}
		goto tr0
tr371:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st235
tr379:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st235
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
//line ragel/parse_datetime.go:9033
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr384
			}
		case data[p] >= 48:
			goto tr383
		}
		goto tr0
tr383:
//line ragel/datetime.rl:7
 pb = p 
	goto st236
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
//line ragel/parse_datetime.go:9052
		switch data[p] {
		case 32:
			goto tr385
		case 43:
			goto tr386
		case 45:
			goto tr387
		case 46:
			goto tr388
		case 47:
			goto tr389
		case 65:
			goto tr391
		case 80:
			goto tr391
		case 90:
			goto tr392
		case 95:
			goto tr389
		case 97:
			goto tr393
		case 112:
			goto tr393
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st247
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr389
			}
		default:
			goto tr389
		}
		goto tr0
tr388:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st237
tr409:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st237
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
//line ragel/parse_datetime.go:9107
		if 48 <= data[p] && data[p] <= 57 {
			goto tr394
		}
		goto tr0
tr394:
//line ragel/datetime.rl:7
 pb = p 
	goto st238
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
//line ragel/parse_datetime.go:9121
		switch data[p] {
		case 32:
			goto tr395
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 65:
			goto tr397
		case 80:
			goto tr397
		case 90:
			goto tr354
		case 95:
			goto tr352
		case 97:
			goto tr398
		case 112:
			goto tr398
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st239
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		switch data[p] {
		case 32:
			goto tr395
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 65:
			goto tr397
		case 80:
			goto tr397
		case 90:
			goto tr354
		case 95:
			goto tr352
		case 97:
			goto tr398
		case 112:
			goto tr398
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st240
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		switch data[p] {
		case 32:
			goto tr395
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 65:
			goto tr397
		case 80:
			goto tr397
		case 90:
			goto tr354
		case 95:
			goto tr352
		case 97:
			goto tr398
		case 112:
			goto tr398
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st241
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		switch data[p] {
		case 32:
			goto tr395
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 65:
			goto tr397
		case 80:
			goto tr397
		case 90:
			goto tr354
		case 95:
			goto tr352
		case 97:
			goto tr398
		case 112:
			goto tr398
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st242
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		switch data[p] {
		case 32:
			goto tr395
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 65:
			goto tr397
		case 80:
			goto tr397
		case 90:
			goto tr354
		case 95:
			goto tr352
		case 97:
			goto tr398
		case 112:
			goto tr398
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st243
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		switch data[p] {
		case 32:
			goto tr395
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 65:
			goto tr397
		case 80:
			goto tr397
		case 90:
			goto tr354
		case 95:
			goto tr352
		case 97:
			goto tr398
		case 112:
			goto tr398
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st244
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		switch data[p] {
		case 32:
			goto tr395
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 65:
			goto tr397
		case 80:
			goto tr397
		case 90:
			goto tr354
		case 95:
			goto tr352
		case 97:
			goto tr398
		case 112:
			goto tr398
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st245
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		switch data[p] {
		case 32:
			goto tr395
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 65:
			goto tr397
		case 80:
			goto tr397
		case 90:
			goto tr354
		case 95:
			goto tr352
		case 97:
			goto tr398
		case 112:
			goto tr398
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st246
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		default:
			goto tr352
		}
		goto tr0
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		switch data[p] {
		case 32:
			goto tr395
		case 43:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr352
		case 65:
			goto tr397
		case 80:
			goto tr397
		case 90:
			goto tr354
		case 95:
			goto tr352
		case 97:
			goto tr398
		case 112:
			goto tr398
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr352
			}
		case data[p] >= 66:
			goto tr352
		}
		goto tr0
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		switch data[p] {
		case 32:
			goto tr406
		case 43:
			goto tr407
		case 45:
			goto tr408
		case 46:
			goto tr409
		case 47:
			goto tr410
		case 65:
			goto tr411
		case 80:
			goto tr411
		case 90:
			goto tr412
		case 95:
			goto tr410
		case 97:
			goto tr413
		case 112:
			goto tr413
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr410
			}
		case data[p] >= 66:
			goto tr410
		}
		goto tr0
tr384:
//line ragel/datetime.rl:7
 pb = p 
	goto st248
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
//line ragel/parse_datetime.go:9520
		switch data[p] {
		case 32:
			goto tr385
		case 43:
			goto tr386
		case 45:
			goto tr387
		case 46:
			goto tr388
		case 47:
			goto tr389
		case 65:
			goto tr391
		case 80:
			goto tr391
		case 90:
			goto tr392
		case 95:
			goto tr389
		case 97:
			goto tr393
		case 112:
			goto tr393
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr389
			}
		case data[p] >= 66:
			goto tr389
		}
		goto tr0
tr365:
//line ragel/datetime.rl:7
 pb = p 
	goto st249
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
//line ragel/parse_datetime.go:9563
		switch data[p] {
		case 32:
			goto tr366
		case 43:
			goto tr367
		case 45:
			goto tr368
		case 47:
			goto tr369
		case 58:
			goto tr371
		case 65:
			goto tr372
		case 80:
			goto tr372
		case 90:
			goto tr373
		case 95:
			goto tr369
		case 97:
			goto tr374
		case 112:
			goto tr374
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr369
			}
		case data[p] >= 66:
			goto tr369
		}
		goto tr0
tr285:
//line ragel/datetime.rl:7
 pb = p 
	goto st250
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
//line ragel/parse_datetime.go:9606
		switch data[p] {
		case 32:
			goto tr287
		case 43:
			goto tr288
		case 45:
			goto tr289
		case 47:
			goto tr290
		case 58:
			goto tr292
		case 65:
			goto tr293
		case 80:
			goto tr293
		case 90:
			goto tr294
		case 95:
			goto tr290
		case 97:
			goto tr295
		case 112:
			goto tr295
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st216
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr290
				}
			case data[p] >= 66:
				goto tr290
			}
		default:
			goto st251
		}
		goto tr0
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		if 48 <= data[p] && data[p] <= 57 {
			goto st217
		}
		goto tr0
tr286:
//line ragel/datetime.rl:7
 pb = p 
	goto st252
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
//line ragel/parse_datetime.go:9667
		switch data[p] {
		case 32:
			goto tr287
		case 43:
			goto tr288
		case 45:
			goto tr289
		case 47:
			goto tr290
		case 58:
			goto tr292
		case 65:
			goto tr293
		case 80:
			goto tr293
		case 90:
			goto tr294
		case 95:
			goto tr290
		case 97:
			goto tr295
		case 112:
			goto tr295
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st251
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr290
			}
		default:
			goto tr290
		}
		goto tr0
tr280:
//line ragel/datetime.rl:7
 pb = p 
	goto st253
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
//line ragel/parse_datetime.go:9714
		if 48 <= data[p] && data[p] <= 57 {
			goto st191
		}
		goto tr0
tr281:
//line ragel/datetime.rl:7
 pb = p 
	goto st254
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
//line ragel/parse_datetime.go:9728
		if 48 <= data[p] && data[p] <= 49 {
			goto st191
		}
		goto tr0
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		if data[p] == 108 {
			goto st256
		}
		goto tr0
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		if data[p] == 32 {
			goto tr277
		}
		goto tr0
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		if data[p] == 103 {
			goto st258
		}
		goto tr0
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		if data[p] == 32 {
			goto tr417
		}
		goto tr0
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		if data[p] == 101 {
			goto st260
		}
		goto tr0
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		if data[p] == 99 {
			goto st261
		}
		goto tr0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		if data[p] == 32 {
			goto tr420
		}
		goto tr0
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		if data[p] == 101 {
			goto st263
		}
		goto tr0
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		if data[p] == 98 {
			goto st264
		}
		goto tr0
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 32:
			goto tr423
		case 114:
			goto st265
		}
		goto tr0
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		if data[p] == 117 {
			goto st266
		}
		goto tr0
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
		if data[p] == 97 {
			goto st267
		}
		goto tr0
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		if data[p] == 114 {
			goto st268
		}
		goto tr0
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		if data[p] == 121 {
			goto st269
		}
		goto tr0
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		if data[p] == 32 {
			goto tr423
		}
		goto tr0
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		switch data[p] {
		case 97:
			goto st271
		case 117:
			goto st277
		}
		goto tr0
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		if data[p] == 110 {
			goto st272
		}
		goto tr0
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		switch data[p] {
		case 32:
			goto tr432
		case 117:
			goto st273
		}
		goto tr0
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
		if data[p] == 97 {
			goto st274
		}
		goto tr0
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		if data[p] == 114 {
			goto st275
		}
		goto tr0
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		if data[p] == 121 {
			goto st276
		}
		goto tr0
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		if data[p] == 32 {
			goto tr432
		}
		goto tr0
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		switch data[p] {
		case 108:
			goto st278
		case 110:
			goto st279
		}
		goto tr0
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		if data[p] == 32 {
			goto tr439
		}
		goto tr0
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		if data[p] == 32 {
			goto tr440
		}
		goto tr0
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		if data[p] == 97 {
			goto st281
		}
		goto tr0
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		switch data[p] {
		case 114:
			goto st282
		case 121:
			goto st285
		}
		goto tr0
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		switch data[p] {
		case 32:
			goto tr444
		case 99:
			goto st283
		}
		goto tr0
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		if data[p] == 104 {
			goto st284
		}
		goto tr0
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		if data[p] == 32 {
			goto tr444
		}
		goto tr0
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		if data[p] == 32 {
			goto tr447
		}
		goto tr0
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		if data[p] == 111 {
			goto st287
		}
		goto tr0
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		if data[p] == 118 {
			goto st288
		}
		goto tr0
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		if data[p] == 32 {
			goto tr450
		}
		goto tr0
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		if data[p] == 99 {
			goto st290
		}
		goto tr0
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
		if data[p] == 116 {
			goto st291
		}
		goto tr0
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		if data[p] == 32 {
			goto tr453
		}
		goto tr0
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		if data[p] == 101 {
			goto st293
		}
		goto tr0
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
		if data[p] == 112 {
			goto st294
		}
		goto tr0
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		if data[p] == 32 {
			goto tr456
		}
		goto tr0
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		if data[p] == 32 {
			goto st296
		}
		goto tr0
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		switch data[p] {
		case 48:
			goto tr458
		case 51:
			goto tr460
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr461
			}
		case data[p] >= 49:
			goto tr459
		}
		goto tr0
tr458:
//line ragel/datetime.rl:7
 pb = p 
	goto st297
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
//line ragel/parse_datetime.go:10149
		if 49 <= data[p] && data[p] <= 57 {
			goto st298
		}
		goto tr0
tr461:
//line ragel/datetime.rl:7
 pb = p 
	goto st298
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
//line ragel/parse_datetime.go:10163
		switch data[p] {
		case 32:
			goto tr463
		case 45:
			goto tr464
		}
		goto tr0
tr463:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st299
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
//line ragel/parse_datetime.go:10188
		switch data[p] {
		case 65:
			goto st300
		case 68:
			goto st311
		case 70:
			goto st314
		case 74:
			goto st322
		case 77:
			goto st332
		case 78:
			goto st338
		case 79:
			goto st341
		case 83:
			goto st344
		}
		goto tr0
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		switch data[p] {
		case 112:
			goto st301
		case 117:
			goto st309
		}
		goto tr0
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		if data[p] == 114 {
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
			goto tr476
		case 105:
			goto st307
		}
		goto tr0
tr476:
//line ragel/datetime.rl:78
 st.Month = 4 ;
	goto st303
tr484:
//line ragel/datetime.rl:82
 st.Month = 8 ;
	goto st303
tr487:
//line ragel/datetime.rl:86
 st.Month = 12 ;
	goto st303
tr490:
//line ragel/datetime.rl:76
 st.Month = 2 ;
	goto st303
tr499:
//line ragel/datetime.rl:75
 st.Month = 1 ;
	goto st303
tr506:
//line ragel/datetime.rl:81
 st.Month = 7 ;
	goto st303
tr507:
//line ragel/datetime.rl:80
 st.Month = 6 ;
	goto st303
tr511:
//line ragel/datetime.rl:77
 st.Month = 3 ;
	goto st303
tr514:
//line ragel/datetime.rl:79
 st.Month = 5 ;
	goto st303
tr517:
//line ragel/datetime.rl:85
 st.Month = 11 ;
	goto st303
tr520:
//line ragel/datetime.rl:84
 st.Month = 10 ;
	goto st303
tr523:
//line ragel/datetime.rl:83
 st.Month = 9 ;
	goto st303
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
//line ragel/parse_datetime.go:10294
		if 48 <= data[p] && data[p] <= 57 {
			goto tr478
		}
		goto tr0
tr478:
//line ragel/datetime.rl:7
 pb = p 
	goto st304
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
//line ragel/parse_datetime.go:10308
		if 48 <= data[p] && data[p] <= 57 {
			goto st305
		}
		goto tr0
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		if 48 <= data[p] && data[p] <= 57 {
			goto st306
		}
		goto tr0
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		if 48 <= data[p] && data[p] <= 57 {
			goto st475
		}
		goto tr0
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		switch data[p] {
		case 32:
			goto tr719
		case 43:
			goto tr720
		case 45:
			goto tr721
		case 47:
			goto tr722
		case 84:
			goto tr723
		case 90:
			goto tr724
		case 95:
			goto tr722
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr722
			}
		case data[p] >= 65:
			goto tr722
		}
		goto st0
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		if data[p] == 108 {
			goto st308
		}
		goto tr0
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		if data[p] == 32 {
			goto tr476
		}
		goto tr0
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		if data[p] == 103 {
			goto st310
		}
		goto tr0
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		if data[p] == 32 {
			goto tr484
		}
		goto tr0
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
		if data[p] == 101 {
			goto st312
		}
		goto tr0
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		if data[p] == 99 {
			goto st313
		}
		goto tr0
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		if data[p] == 32 {
			goto tr487
		}
		goto tr0
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		if data[p] == 101 {
			goto st315
		}
		goto tr0
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		if data[p] == 98 {
			goto st316
		}
		goto tr0
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		switch data[p] {
		case 32:
			goto tr490
		case 114:
			goto st317
		}
		goto tr0
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		if data[p] == 117 {
			goto st318
		}
		goto tr0
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		if data[p] == 97 {
			goto st319
		}
		goto tr0
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		if data[p] == 114 {
			goto st320
		}
		goto tr0
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		if data[p] == 121 {
			goto st321
		}
		goto tr0
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		if data[p] == 32 {
			goto tr490
		}
		goto tr0
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		switch data[p] {
		case 97:
			goto st323
		case 117:
			goto st329
		}
		goto tr0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		if data[p] == 110 {
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
			goto tr499
		case 117:
			goto st325
		}
		goto tr0
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
		if data[p] == 97 {
			goto st326
		}
		goto tr0
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		if data[p] == 114 {
			goto st327
		}
		goto tr0
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		if data[p] == 121 {
			goto st328
		}
		goto tr0
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		if data[p] == 32 {
			goto tr499
		}
		goto tr0
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		switch data[p] {
		case 108:
			goto st330
		case 110:
			goto st331
		}
		goto tr0
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		if data[p] == 32 {
			goto tr506
		}
		goto tr0
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		if data[p] == 32 {
			goto tr507
		}
		goto tr0
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		if data[p] == 97 {
			goto st333
		}
		goto tr0
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 114:
			goto st334
		case 121:
			goto st337
		}
		goto tr0
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		switch data[p] {
		case 32:
			goto tr511
		case 99:
			goto st335
		}
		goto tr0
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		if data[p] == 104 {
			goto st336
		}
		goto tr0
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		if data[p] == 32 {
			goto tr511
		}
		goto tr0
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		if data[p] == 32 {
			goto tr514
		}
		goto tr0
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		if data[p] == 111 {
			goto st339
		}
		goto tr0
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		if data[p] == 118 {
			goto st340
		}
		goto tr0
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		if data[p] == 32 {
			goto tr517
		}
		goto tr0
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		if data[p] == 99 {
			goto st342
		}
		goto tr0
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		if data[p] == 116 {
			goto st343
		}
		goto tr0
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
		if data[p] == 32 {
			goto tr520
		}
		goto tr0
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		if data[p] == 101 {
			goto st345
		}
		goto tr0
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		if data[p] == 112 {
			goto st346
		}
		goto tr0
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		if data[p] == 32 {
			goto tr523
		}
		goto tr0
tr464:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
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
//line ragel/parse_datetime.go:10756
		switch data[p] {
		case 65:
			goto st348
		case 68:
			goto st355
		case 70:
			goto st358
		case 74:
			goto st366
		case 77:
			goto st376
		case 78:
			goto st382
		case 79:
			goto st385
		case 83:
			goto st388
		}
		goto tr0
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		switch data[p] {
		case 112:
			goto st349
		case 117:
			goto st353
		}
		goto tr0
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		if data[p] == 114 {
			goto st350
		}
		goto tr0
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		switch data[p] {
		case 45:
			goto tr215
		case 105:
			goto st351
		}
		goto tr0
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		if data[p] == 108 {
			goto st352
		}
		goto tr0
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		if data[p] == 45 {
			goto tr215
		}
		goto tr0
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		if data[p] == 103 {
			goto st354
		}
		goto tr0
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		if data[p] == 45 {
			goto tr221
		}
		goto tr0
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		if data[p] == 101 {
			goto st356
		}
		goto tr0
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		if data[p] == 99 {
			goto st357
		}
		goto tr0
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		if data[p] == 45 {
			goto tr224
		}
		goto tr0
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		if data[p] == 101 {
			goto st359
		}
		goto tr0
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		if data[p] == 98 {
			goto st360
		}
		goto tr0
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		switch data[p] {
		case 45:
			goto tr227
		case 114:
			goto st361
		}
		goto tr0
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		if data[p] == 117 {
			goto st362
		}
		goto tr0
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		if data[p] == 97 {
			goto st363
		}
		goto tr0
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		if data[p] == 114 {
			goto st364
		}
		goto tr0
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		if data[p] == 121 {
			goto st365
		}
		goto tr0
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		if data[p] == 45 {
			goto tr227
		}
		goto tr0
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		switch data[p] {
		case 97:
			goto st367
		case 117:
			goto st373
		}
		goto tr0
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		if data[p] == 110 {
			goto st368
		}
		goto tr0
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		switch data[p] {
		case 45:
			goto tr236
		case 117:
			goto st369
		}
		goto tr0
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
		if data[p] == 97 {
			goto st370
		}
		goto tr0
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		if data[p] == 114 {
			goto st371
		}
		goto tr0
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		if data[p] == 121 {
			goto st372
		}
		goto tr0
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		if data[p] == 45 {
			goto tr236
		}
		goto tr0
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		switch data[p] {
		case 108:
			goto st374
		case 110:
			goto st375
		}
		goto tr0
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		if data[p] == 45 {
			goto tr243
		}
		goto tr0
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		if data[p] == 45 {
			goto tr244
		}
		goto tr0
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		if data[p] == 97 {
			goto st377
		}
		goto tr0
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		switch data[p] {
		case 114:
			goto st378
		case 121:
			goto st381
		}
		goto tr0
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		switch data[p] {
		case 45:
			goto tr248
		case 99:
			goto st379
		}
		goto tr0
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		if data[p] == 104 {
			goto st380
		}
		goto tr0
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		if data[p] == 45 {
			goto tr248
		}
		goto tr0
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		if data[p] == 45 {
			goto tr251
		}
		goto tr0
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		if data[p] == 111 {
			goto st383
		}
		goto tr0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		if data[p] == 118 {
			goto st384
		}
		goto tr0
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		if data[p] == 45 {
			goto tr254
		}
		goto tr0
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		if data[p] == 99 {
			goto st386
		}
		goto tr0
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		if data[p] == 116 {
			goto st387
		}
		goto tr0
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		if data[p] == 45 {
			goto tr257
		}
		goto tr0
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
		if data[p] == 101 {
			goto st389
		}
		goto tr0
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		if data[p] == 112 {
			goto st390
		}
		goto tr0
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		if data[p] == 45 {
			goto tr260
		}
		goto tr0
tr459:
//line ragel/datetime.rl:7
 pb = p 
	goto st391
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
//line ragel/parse_datetime.go:11196
		switch data[p] {
		case 32:
			goto tr463
		case 45:
			goto tr464
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st298
		}
		goto tr0
tr460:
//line ragel/datetime.rl:7
 pb = p 
	goto st392
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
//line ragel/parse_datetime.go:11216
		switch data[p] {
		case 32:
			goto tr463
		case 45:
			goto tr464
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st298
		}
		goto tr0
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		if data[p] == 97 {
			goto st394
		}
		goto tr0
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		if data[p] == 121 {
			goto st395
		}
		goto tr0
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		switch data[p] {
		case 32:
			goto st185
		case 44:
			goto st295
		}
		goto tr0
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		if data[p] == 111 {
			goto st397
		}
		goto tr0
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		if data[p] == 110 {
			goto st184
		}
		goto tr0
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		switch data[p] {
		case 97:
			goto st399
		case 117:
			goto st397
		}
		goto tr0
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		if data[p] == 116 {
			goto st400
		}
		goto tr0
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		switch data[p] {
		case 32:
			goto st185
		case 44:
			goto st295
		case 117:
			goto st401
		}
		goto tr0
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		if data[p] == 114 {
			goto st402
		}
		goto tr0
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
		if data[p] == 100 {
			goto st393
		}
		goto tr0
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		switch data[p] {
		case 104:
			goto st404
		case 117:
			goto st407
		}
		goto tr0
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		if data[p] == 117 {
			goto st405
		}
		goto tr0
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 32:
			goto st185
		case 44:
			goto st295
		case 114:
			goto st406
		}
		goto tr0
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		if data[p] == 115 {
			goto st402
		}
		goto tr0
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		if data[p] == 101 {
			goto st408
		}
		goto tr0
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
		switch data[p] {
		case 32:
			goto st185
		case 44:
			goto st295
		case 115:
			goto st402
		}
		goto tr0
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		if data[p] == 101 {
			goto st410
		}
		goto tr0
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		if data[p] == 100 {
			goto st411
		}
		goto tr0
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		switch data[p] {
		case 32:
			goto st185
		case 44:
			goto st295
		case 110:
			goto st412
		}
		goto tr0
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		if data[p] == 101 {
			goto st406
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
	_test_eof413: cs = 413; goto _test_eof
	_test_eof414: cs = 414; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof415: cs = 415; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof416: cs = 416; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof417: cs = 417; goto _test_eof
	_test_eof418: cs = 418; goto _test_eof
	_test_eof419: cs = 419; goto _test_eof
	_test_eof420: cs = 420; goto _test_eof
	_test_eof421: cs = 421; goto _test_eof
	_test_eof422: cs = 422; goto _test_eof
	_test_eof423: cs = 423; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof424: cs = 424; goto _test_eof
	_test_eof425: cs = 425; goto _test_eof
	_test_eof426: cs = 426; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof427: cs = 427; goto _test_eof
	_test_eof428: cs = 428; goto _test_eof
	_test_eof429: cs = 429; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof430: cs = 430; goto _test_eof
	_test_eof431: cs = 431; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof432: cs = 432; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof433: cs = 433; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
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
	_test_eof23: cs = 23; goto _test_eof
	_test_eof444: cs = 444; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof445: cs = 445; goto _test_eof
	_test_eof446: cs = 446; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof447: cs = 447; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
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
	_test_eof27: cs = 27; goto _test_eof
	_test_eof461: cs = 461; goto _test_eof
	_test_eof462: cs = 462; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof463: cs = 463; goto _test_eof
	_test_eof464: cs = 464; goto _test_eof
	_test_eof465: cs = 465; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
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
	_test_eof466: cs = 466; goto _test_eof
	_test_eof467: cs = 467; goto _test_eof
	_test_eof468: cs = 468; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof469: cs = 469; goto _test_eof
	_test_eof470: cs = 470; goto _test_eof
	_test_eof471: cs = 471; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof472: cs = 472; goto _test_eof
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
	_test_eof473: cs = 473; goto _test_eof
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
	_test_eof474: cs = 474; goto _test_eof
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
	_test_eof475: cs = 475; goto _test_eof
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

	_test_eof: {}
	if p == eof {
		switch cs {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279, 280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299, 300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339, 340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 383, 384, 385, 386, 387, 388, 389, 390, 391, 392, 393, 394, 395, 396, 397, 398, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412:
//line ragel/datetime.rl:5
 return st, err 
		case 420:
//line ragel/datetime.rl:9
 st.Zoned = true 
		case 474, 475:
//line ragel/datetime.rl:15

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

		case 473:
//line ragel/datetime.rl:19

    st.Year = parse_year_2_digits(data[pb:pb+2])

		case 413, 466, 467, 468:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

		case 433, 444:
//line ragel/datetime.rl:33

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

		case 416, 430:
//line ragel/datetime.rl:48

    st.Ad_bc = ADBC_BC;

		case 428:
//line ragel/datetime.rl:52

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
                    st.Hour += 12;
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }

		case 463, 464, 465:
//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

		case 432:
//line ragel/datetime.rl:98

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

		case 425, 460, 461:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

		case 446:
//line ragel/datetime.rl:104

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

		case 445, 459:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

		case 457:
//line ragel/datetime.rl:110

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

		case 447, 458:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

		case 434, 435, 436, 437, 438, 439, 440, 441, 442, 448, 449, 450, 451, 452, 453, 454, 455, 456:
//line ragel/datetime.rl:116

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
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+6]) 
    }

		case 469, 470, 471, 472:
//line ragel/datetime.rl:23

    st.Day_of_year, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:88

    switch p - pb {
        case 1: st.Day, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Day, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

		case 421, 422:
//line ragel/datetime.rl:145

    switch p - pb {
        case 1: st.Offset_minute, _ = strconv.Atoi(data[pb:pb+1])
        case 2: st.Offset_minute, _ = strconv.Atoi(data[pb:pb+2])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:9
 st.Zoned = true 
		case 417, 418, 419, 423:
//line ragel/datetime.rl:155
 
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
        case 1:{st.Offset_hour, _ = strconv.Atoi(data[pb:pb+1])}
        case 2:{st.Offset_hour, _ = strconv.Atoi(data[pb:pb+2])}
        case 3:{ 
            num , _ := strconv.Atoi(data[pb:pb+3])
            st.Offset_hour = num/100
            st.Offset_minute = num%100
            if st.Offset_minute >=60 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        case 4:{ 
            num := parse_digits(data[pb:pb+4])
            st.Offset_hour = num/100
            st.Offset_minute = num%100
            if st.Offset_minute >=60 || st.Offset_hour>=15 {
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
		case 424, 431:
//line ragel/datetime.rl:193

    st.Zone_name_or_abbrev = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:9
 st.Zoned = true 
//line ragel/parse_datetime.go:12117
		}
	}

	_out: {}
	}

//line ragel/parse_datetime.go.rl:34

    if p != eof {  // input date not fully consumed
        err = errors.New("input data not fully consumed");
        return
    }

    if cs < atoi_first_final {
        err = fmt.Errorf("time parse error: %s", data)
    }
    return
}

