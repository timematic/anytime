package ragel

import (
"fmt"
"strconv"
"errors"
)

var datetime_parser_start int = 1
var _ = datetime_parser_start
var datetime_parser_first_final int = 406
var _ = datetime_parser_first_final
var datetime_parser_error int = 0
var _ = datetime_parser_error
var datetime_parser_en_main int = 1
var _ = datetime_parser_en_main
var start int = 1
var _ = start
var en_main int = 1
var _ = en_main
func Parse(data string) (st ParsedDatetime, err error) {
	cs, p, pe := 0, 0, len(data)
	eof := pe
	pb := p
	
	{
		cs=int( start );
		
	}
	{
		switch cs {
			case 1:
			goto st_case_1;
			case 0:
			goto st_case_0;
			case 2:
			goto st_case_2;
			case 3:
			goto st_case_3;
			case 4:
			goto st_case_4;
			case 5:
			goto st_case_5;
			case 6:
			goto st_case_6;
			case 7:
			goto st_case_7;
			case 8:
			goto st_case_8;
			case 406:
			goto st_case_406;
			case 407:
			goto st_case_407;
			case 9:
			goto st_case_9;
			case 10:
			goto st_case_10;
			case 408:
			goto st_case_408;
			case 11:
			goto st_case_11;
			case 409:
			goto st_case_409;
			case 12:
			goto st_case_12;
			case 410:
			goto st_case_410;
			case 411:
			goto st_case_411;
			case 412:
			goto st_case_412;
			case 413:
			goto st_case_413;
			case 414:
			goto st_case_414;
			case 415:
			goto st_case_415;
			case 416:
			goto st_case_416;
			case 13:
			goto st_case_13;
			case 14:
			goto st_case_14;
			case 15:
			goto st_case_15;
			case 417:
			goto st_case_417;
			case 418:
			goto st_case_418;
			case 419:
			goto st_case_419;
			case 16:
			goto st_case_16;
			case 420:
			goto st_case_420;
			case 421:
			goto st_case_421;
			case 422:
			goto st_case_422;
			case 17:
			goto st_case_17;
			case 18:
			goto st_case_18;
			case 423:
			goto st_case_423;
			case 424:
			goto st_case_424;
			case 19:
			goto st_case_19;
			case 20:
			goto st_case_20;
			case 425:
			goto st_case_425;
			case 21:
			goto st_case_21;
			case 426:
			goto st_case_426;
			case 22:
			goto st_case_22;
			case 427:
			goto st_case_427;
			case 428:
			goto st_case_428;
			case 429:
			goto st_case_429;
			case 430:
			goto st_case_430;
			case 431:
			goto st_case_431;
			case 432:
			goto st_case_432;
			case 433:
			goto st_case_433;
			case 434:
			goto st_case_434;
			case 435:
			goto st_case_435;
			case 23:
			goto st_case_23;
			case 436:
			goto st_case_436;
			case 24:
			goto st_case_24;
			case 437:
			goto st_case_437;
			case 438:
			goto st_case_438;
			case 25:
			goto st_case_25;
			case 439:
			goto st_case_439;
			case 26:
			goto st_case_26;
			case 440:
			goto st_case_440;
			case 441:
			goto st_case_441;
			case 442:
			goto st_case_442;
			case 443:
			goto st_case_443;
			case 444:
			goto st_case_444;
			case 445:
			goto st_case_445;
			case 446:
			goto st_case_446;
			case 447:
			goto st_case_447;
			case 448:
			goto st_case_448;
			case 449:
			goto st_case_449;
			case 450:
			goto st_case_450;
			case 451:
			goto st_case_451;
			case 452:
			goto st_case_452;
			case 27:
			goto st_case_27;
			case 453:
			goto st_case_453;
			case 454:
			goto st_case_454;
			case 455:
			goto st_case_455;
			case 28:
			goto st_case_28;
			case 29:
			goto st_case_29;
			case 30:
			goto st_case_30;
			case 31:
			goto st_case_31;
			case 456:
			goto st_case_456;
			case 457:
			goto st_case_457;
			case 458:
			goto st_case_458;
			case 32:
			goto st_case_32;
			case 33:
			goto st_case_33;
			case 34:
			goto st_case_34;
			case 35:
			goto st_case_35;
			case 36:
			goto st_case_36;
			case 37:
			goto st_case_37;
			case 38:
			goto st_case_38;
			case 39:
			goto st_case_39;
			case 40:
			goto st_case_40;
			case 41:
			goto st_case_41;
			case 42:
			goto st_case_42;
			case 43:
			goto st_case_43;
			case 44:
			goto st_case_44;
			case 45:
			goto st_case_45;
			case 46:
			goto st_case_46;
			case 47:
			goto st_case_47;
			case 48:
			goto st_case_48;
			case 49:
			goto st_case_49;
			case 50:
			goto st_case_50;
			case 51:
			goto st_case_51;
			case 52:
			goto st_case_52;
			case 53:
			goto st_case_53;
			case 54:
			goto st_case_54;
			case 55:
			goto st_case_55;
			case 56:
			goto st_case_56;
			case 57:
			goto st_case_57;
			case 58:
			goto st_case_58;
			case 59:
			goto st_case_59;
			case 60:
			goto st_case_60;
			case 61:
			goto st_case_61;
			case 62:
			goto st_case_62;
			case 63:
			goto st_case_63;
			case 64:
			goto st_case_64;
			case 65:
			goto st_case_65;
			case 66:
			goto st_case_66;
			case 67:
			goto st_case_67;
			case 68:
			goto st_case_68;
			case 69:
			goto st_case_69;
			case 70:
			goto st_case_70;
			case 71:
			goto st_case_71;
			case 72:
			goto st_case_72;
			case 73:
			goto st_case_73;
			case 74:
			goto st_case_74;
			case 75:
			goto st_case_75;
			case 76:
			goto st_case_76;
			case 77:
			goto st_case_77;
			case 78:
			goto st_case_78;
			case 79:
			goto st_case_79;
			case 80:
			goto st_case_80;
			case 459:
			goto st_case_459;
			case 460:
			goto st_case_460;
			case 81:
			goto st_case_81;
			case 82:
			goto st_case_82;
			case 83:
			goto st_case_83;
			case 84:
			goto st_case_84;
			case 85:
			goto st_case_85;
			case 86:
			goto st_case_86;
			case 87:
			goto st_case_87;
			case 461:
			goto st_case_461;
			case 88:
			goto st_case_88;
			case 89:
			goto st_case_89;
			case 90:
			goto st_case_90;
			case 91:
			goto st_case_91;
			case 92:
			goto st_case_92;
			case 93:
			goto st_case_93;
			case 94:
			goto st_case_94;
			case 95:
			goto st_case_95;
			case 96:
			goto st_case_96;
			case 97:
			goto st_case_97;
			case 98:
			goto st_case_98;
			case 99:
			goto st_case_99;
			case 100:
			goto st_case_100;
			case 101:
			goto st_case_101;
			case 102:
			goto st_case_102;
			case 103:
			goto st_case_103;
			case 104:
			goto st_case_104;
			case 105:
			goto st_case_105;
			case 106:
			goto st_case_106;
			case 107:
			goto st_case_107;
			case 108:
			goto st_case_108;
			case 109:
			goto st_case_109;
			case 110:
			goto st_case_110;
			case 111:
			goto st_case_111;
			case 112:
			goto st_case_112;
			case 113:
			goto st_case_113;
			case 114:
			goto st_case_114;
			case 115:
			goto st_case_115;
			case 116:
			goto st_case_116;
			case 117:
			goto st_case_117;
			case 118:
			goto st_case_118;
			case 119:
			goto st_case_119;
			case 120:
			goto st_case_120;
			case 121:
			goto st_case_121;
			case 122:
			goto st_case_122;
			case 123:
			goto st_case_123;
			case 124:
			goto st_case_124;
			case 125:
			goto st_case_125;
			case 126:
			goto st_case_126;
			case 127:
			goto st_case_127;
			case 128:
			goto st_case_128;
			case 129:
			goto st_case_129;
			case 130:
			goto st_case_130;
			case 131:
			goto st_case_131;
			case 132:
			goto st_case_132;
			case 133:
			goto st_case_133;
			case 134:
			goto st_case_134;
			case 135:
			goto st_case_135;
			case 136:
			goto st_case_136;
			case 137:
			goto st_case_137;
			case 138:
			goto st_case_138;
			case 139:
			goto st_case_139;
			case 140:
			goto st_case_140;
			case 462:
			goto st_case_462;
			case 141:
			goto st_case_141;
			case 142:
			goto st_case_142;
			case 143:
			goto st_case_143;
			case 144:
			goto st_case_144;
			case 145:
			goto st_case_145;
			case 146:
			goto st_case_146;
			case 147:
			goto st_case_147;
			case 148:
			goto st_case_148;
			case 149:
			goto st_case_149;
			case 150:
			goto st_case_150;
			case 151:
			goto st_case_151;
			case 152:
			goto st_case_152;
			case 153:
			goto st_case_153;
			case 154:
			goto st_case_154;
			case 155:
			goto st_case_155;
			case 156:
			goto st_case_156;
			case 157:
			goto st_case_157;
			case 158:
			goto st_case_158;
			case 159:
			goto st_case_159;
			case 160:
			goto st_case_160;
			case 161:
			goto st_case_161;
			case 162:
			goto st_case_162;
			case 163:
			goto st_case_163;
			case 164:
			goto st_case_164;
			case 165:
			goto st_case_165;
			case 166:
			goto st_case_166;
			case 167:
			goto st_case_167;
			case 168:
			goto st_case_168;
			case 169:
			goto st_case_169;
			case 170:
			goto st_case_170;
			case 171:
			goto st_case_171;
			case 172:
			goto st_case_172;
			case 173:
			goto st_case_173;
			case 463:
			goto st_case_463;
			case 174:
			goto st_case_174;
			case 175:
			goto st_case_175;
			case 176:
			goto st_case_176;
			case 177:
			goto st_case_177;
			case 178:
			goto st_case_178;
			case 179:
			goto st_case_179;
			case 180:
			goto st_case_180;
			case 181:
			goto st_case_181;
			case 182:
			goto st_case_182;
			case 183:
			goto st_case_183;
			case 184:
			goto st_case_184;
			case 185:
			goto st_case_185;
			case 186:
			goto st_case_186;
			case 187:
			goto st_case_187;
			case 188:
			goto st_case_188;
			case 189:
			goto st_case_189;
			case 190:
			goto st_case_190;
			case 191:
			goto st_case_191;
			case 192:
			goto st_case_192;
			case 193:
			goto st_case_193;
			case 194:
			goto st_case_194;
			case 195:
			goto st_case_195;
			case 196:
			goto st_case_196;
			case 197:
			goto st_case_197;
			case 198:
			goto st_case_198;
			case 199:
			goto st_case_199;
			case 200:
			goto st_case_200;
			case 201:
			goto st_case_201;
			case 202:
			goto st_case_202;
			case 203:
			goto st_case_203;
			case 204:
			goto st_case_204;
			case 205:
			goto st_case_205;
			case 206:
			goto st_case_206;
			case 207:
			goto st_case_207;
			case 208:
			goto st_case_208;
			case 209:
			goto st_case_209;
			case 210:
			goto st_case_210;
			case 211:
			goto st_case_211;
			case 212:
			goto st_case_212;
			case 213:
			goto st_case_213;
			case 214:
			goto st_case_214;
			case 215:
			goto st_case_215;
			case 216:
			goto st_case_216;
			case 217:
			goto st_case_217;
			case 218:
			goto st_case_218;
			case 219:
			goto st_case_219;
			case 220:
			goto st_case_220;
			case 221:
			goto st_case_221;
			case 222:
			goto st_case_222;
			case 223:
			goto st_case_223;
			case 224:
			goto st_case_224;
			case 225:
			goto st_case_225;
			case 226:
			goto st_case_226;
			case 227:
			goto st_case_227;
			case 228:
			goto st_case_228;
			case 229:
			goto st_case_229;
			case 230:
			goto st_case_230;
			case 231:
			goto st_case_231;
			case 232:
			goto st_case_232;
			case 233:
			goto st_case_233;
			case 234:
			goto st_case_234;
			case 235:
			goto st_case_235;
			case 236:
			goto st_case_236;
			case 237:
			goto st_case_237;
			case 238:
			goto st_case_238;
			case 239:
			goto st_case_239;
			case 240:
			goto st_case_240;
			case 241:
			goto st_case_241;
			case 242:
			goto st_case_242;
			case 243:
			goto st_case_243;
			case 244:
			goto st_case_244;
			case 245:
			goto st_case_245;
			case 246:
			goto st_case_246;
			case 247:
			goto st_case_247;
			case 248:
			goto st_case_248;
			case 249:
			goto st_case_249;
			case 250:
			goto st_case_250;
			case 251:
			goto st_case_251;
			case 252:
			goto st_case_252;
			case 253:
			goto st_case_253;
			case 254:
			goto st_case_254;
			case 255:
			goto st_case_255;
			case 256:
			goto st_case_256;
			case 257:
			goto st_case_257;
			case 258:
			goto st_case_258;
			case 259:
			goto st_case_259;
			case 260:
			goto st_case_260;
			case 261:
			goto st_case_261;
			case 262:
			goto st_case_262;
			case 263:
			goto st_case_263;
			case 264:
			goto st_case_264;
			case 265:
			goto st_case_265;
			case 266:
			goto st_case_266;
			case 267:
			goto st_case_267;
			case 268:
			goto st_case_268;
			case 269:
			goto st_case_269;
			case 270:
			goto st_case_270;
			case 271:
			goto st_case_271;
			case 272:
			goto st_case_272;
			case 273:
			goto st_case_273;
			case 274:
			goto st_case_274;
			case 275:
			goto st_case_275;
			case 276:
			goto st_case_276;
			case 277:
			goto st_case_277;
			case 278:
			goto st_case_278;
			case 279:
			goto st_case_279;
			case 280:
			goto st_case_280;
			case 281:
			goto st_case_281;
			case 282:
			goto st_case_282;
			case 283:
			goto st_case_283;
			case 284:
			goto st_case_284;
			case 285:
			goto st_case_285;
			case 286:
			goto st_case_286;
			case 287:
			goto st_case_287;
			case 288:
			goto st_case_288;
			case 289:
			goto st_case_289;
			case 290:
			goto st_case_290;
			case 291:
			goto st_case_291;
			case 292:
			goto st_case_292;
			case 293:
			goto st_case_293;
			case 294:
			goto st_case_294;
			case 295:
			goto st_case_295;
			case 296:
			goto st_case_296;
			case 297:
			goto st_case_297;
			case 298:
			goto st_case_298;
			case 299:
			goto st_case_299;
			case 300:
			goto st_case_300;
			case 301:
			goto st_case_301;
			case 302:
			goto st_case_302;
			case 303:
			goto st_case_303;
			case 304:
			goto st_case_304;
			case 305:
			goto st_case_305;
			case 306:
			goto st_case_306;
			case 307:
			goto st_case_307;
			case 308:
			goto st_case_308;
			case 309:
			goto st_case_309;
			case 310:
			goto st_case_310;
			case 311:
			goto st_case_311;
			case 312:
			goto st_case_312;
			case 313:
			goto st_case_313;
			case 314:
			goto st_case_314;
			case 315:
			goto st_case_315;
			case 316:
			goto st_case_316;
			case 317:
			goto st_case_317;
			case 318:
			goto st_case_318;
			case 319:
			goto st_case_319;
			case 320:
			goto st_case_320;
			case 321:
			goto st_case_321;
			case 322:
			goto st_case_322;
			case 323:
			goto st_case_323;
			case 324:
			goto st_case_324;
			case 325:
			goto st_case_325;
			case 326:
			goto st_case_326;
			case 327:
			goto st_case_327;
			case 328:
			goto st_case_328;
			case 329:
			goto st_case_329;
			case 330:
			goto st_case_330;
			case 331:
			goto st_case_331;
			case 332:
			goto st_case_332;
			case 333:
			goto st_case_333;
			case 334:
			goto st_case_334;
			case 335:
			goto st_case_335;
			case 336:
			goto st_case_336;
			case 337:
			goto st_case_337;
			case 338:
			goto st_case_338;
			case 339:
			goto st_case_339;
			case 340:
			goto st_case_340;
			case 341:
			goto st_case_341;
			case 342:
			goto st_case_342;
			case 343:
			goto st_case_343;
			case 344:
			goto st_case_344;
			case 345:
			goto st_case_345;
			case 346:
			goto st_case_346;
			case 347:
			goto st_case_347;
			case 348:
			goto st_case_348;
			case 349:
			goto st_case_349;
			case 350:
			goto st_case_350;
			case 351:
			goto st_case_351;
			case 352:
			goto st_case_352;
			case 353:
			goto st_case_353;
			case 354:
			goto st_case_354;
			case 355:
			goto st_case_355;
			case 356:
			goto st_case_356;
			case 357:
			goto st_case_357;
			case 358:
			goto st_case_358;
			case 359:
			goto st_case_359;
			case 360:
			goto st_case_360;
			case 361:
			goto st_case_361;
			case 362:
			goto st_case_362;
			case 363:
			goto st_case_363;
			case 364:
			goto st_case_364;
			case 365:
			goto st_case_365;
			case 366:
			goto st_case_366;
			case 367:
			goto st_case_367;
			case 368:
			goto st_case_368;
			case 369:
			goto st_case_369;
			case 370:
			goto st_case_370;
			case 371:
			goto st_case_371;
			case 372:
			goto st_case_372;
			case 373:
			goto st_case_373;
			case 374:
			goto st_case_374;
			case 375:
			goto st_case_375;
			case 376:
			goto st_case_376;
			case 377:
			goto st_case_377;
			case 378:
			goto st_case_378;
			case 379:
			goto st_case_379;
			case 380:
			goto st_case_380;
			case 381:
			goto st_case_381;
			case 382:
			goto st_case_382;
			case 383:
			goto st_case_383;
			case 384:
			goto st_case_384;
			case 385:
			goto st_case_385;
			case 386:
			goto st_case_386;
			case 387:
			goto st_case_387;
			case 388:
			goto st_case_388;
			case 389:
			goto st_case_389;
			case 390:
			goto st_case_390;
			case 391:
			goto st_case_391;
			case 392:
			goto st_case_392;
			case 393:
			goto st_case_393;
			case 394:
			goto st_case_394;
			case 395:
			goto st_case_395;
			case 396:
			goto st_case_396;
			case 397:
			goto st_case_397;
			case 398:
			goto st_case_398;
			case 399:
			goto st_case_399;
			case 400:
			goto st_case_400;
			case 401:
			goto st_case_401;
			case 402:
			goto st_case_402;
			case 403:
			goto st_case_403;
			case 404:
			goto st_case_404;
			case 405:
			goto st_case_405;
			
		}
		_ctr1:
		{return st, err }
		goto _st1;
		_st1:
		if p == eof {
			goto _out1;
			
		}
		p+=1;
		st_case_1:
		if p == pe && p != eof {
			goto _out1;
			
		}
		if p == eof {
			goto _ctr1;
			
		} else {
			switch ( data[ p ] ) {
				case 48:
				{
					goto _ctr3;
					
				}
				case 51:
				{
					goto _ctr5;
					
				}
				case 65:
				{
					goto _st131;
					
				}
				case 68:
				{
					goto _st147;
					
				}
				case 70:
				{
					goto _st150;
					
				}
				case 74:
				{
					goto _st366;
					
				}
				case 77:
				{
					goto _st376;
					
				}
				case 78:
				{
					goto _st383;
					
				}
				case 79:
				{
					goto _st386;
					
				}
				case 83:
				{
					goto _st389;
					
				}
				case 84:
				{
					goto _st396;
					
				}
				case 87:
				{
					goto _st402;
					
				}
				
			}
			if ( data[ p ] ) > 50 {
				if 52 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr6;
					
				}
				
			} else if ( data[ p ] ) >= 49 {
				goto _ctr4;
				
			}
			goto _ctr2;
			
		}
		_ctr2:
		{return st, err }
		goto _st0;
		_st0:
		if p == eof {
			goto _out0;
			
		}
		st_case_0:
		goto _out0;
		_ctr17:
		{return st, err }
		goto _st2;
		_ctr3:
		{pb = p }
		goto _st2;
		_st2:
		if p == eof {
			goto _out2;
			
		}
		p+=1;
		st_case_2:
		if p == pe && p != eof {
			goto _out2;
			
		}
		if p == eof {
			goto _ctr17;
			
		} else {
			if ( data[ p ] ) == 48 {
				goto _st3;
				
			}
			if 49 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st81;
				
			}
			goto _ctr2;
			
		}
		_ctr20:
		{return st, err }
		goto _st3;
		_st3:
		if p == eof {
			goto _out3;
			
		}
		p+=1;
		st_case_3:
		if p == pe && p != eof {
			goto _out3;
			
		}
		if p == eof {
			goto _ctr20;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st4;
				
			}
			goto _ctr2;
			
		}
		_ctr22:
		{return st, err }
		goto _st4;
		_st4:
		if p == eof {
			goto _out4;
			
		}
		p+=1;
		st_case_4:
		if p == pe && p != eof {
			goto _out4;
			
		}
		if p == eof {
			goto _ctr22;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st5;
				
			}
			goto _ctr2;
			
		}
		_ctr24:
		{return st, err }
		goto _st5;
		_st5:
		if p == eof {
			goto _out5;
			
		}
		p+=1;
		st_case_5:
		if p == pe && p != eof {
			goto _out5;
			
		}
		if p == eof {
			goto _ctr24;
			
		} else {
			if ( data[ p ] ) == 46 {
				goto _ctr26;
				
			}
			if ( data[ p ] ) > 47 {
				if ( data[ p ] ) <= 57 {
					goto _ctr27;
					
				}
				
			} else if ( data[ p ] ) >= 45 {
				goto _ctr25;
				
			}
			goto _ctr2;
			
		}
		_ctr28:
		{return st, err }
		goto _st6;
		_ctr25:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		goto _st6;
		_st6:
		if p == eof {
			goto _out6;
			
		}
		p+=1;
		st_case_6:
		if p == pe && p != eof {
			goto _out6;
			
		}
		if p == eof {
			goto _ctr28;
			
		} else {
			switch ( data[ p ] ) {
				case 48:
				{
					goto _ctr29;
					
				}
				case 49:
				{
					goto _ctr30;
					
				}
				case 65:
				{
					goto _st34;
					
				}
				case 68:
				{
					goto _st41;
					
				}
				case 70:
				{
					goto _st44;
					
				}
				case 74:
				{
					goto _st52;
					
				}
				case 77:
				{
					goto _st62;
					
				}
				case 78:
				{
					goto _st68;
					
				}
				case 79:
				{
					goto _st71;
					
				}
				case 83:
				{
					goto _st74;
					
				}
				
			}
			if 50 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _ctr31;
				
			}
			goto _ctr2;
			
		}
		_ctr40:
		{return st, err }
		goto _st7;
		_ctr29:
		{pb = p }
		goto _st7;
		_st7:
		if p == eof {
			goto _out7;
			
		}
		p+=1;
		st_case_7:
		if p == pe && p != eof {
			goto _out7;
			
		}
		if p == eof {
			goto _ctr40;
			
		} else {
			if ( data[ p ] ) == 48 {
				goto _st8;
				
			}
			if 49 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st29;
				
			}
			goto _ctr2;
			
		}
		_ctr43:
		{return st, err }
		goto _st8;
		_st8:
		if p == eof {
			goto _out8;
			
		}
		p+=1;
		st_case_8:
		if p == pe && p != eof {
			goto _out8;
			
		}
		if p == eof {
			goto _ctr43;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st406;
				
			}
			goto _ctr2;
			
		}
		_ctr955:
		{st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
		}
		goto _st406;
		_st406:
		if p == eof {
			goto _out406;
			
		}
		p+=1;
		st_case_406:
		if p == pe && p != eof {
			goto _out406;
			
		}
		if p == eof {
			goto _ctr955;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr956;
					
				}
				case 43:
				{
					goto _ctr957;
					
				}
				case 45:
				{
					goto _ctr958;
					
				}
				case 47:
				{
					goto _ctr959;
					
				}
				case 84:
				{
					goto _ctr960;
					
				}
				case 90:
				{
					goto _ctr961;
					
				}
				case 95:
				{
					goto _ctr962;
					
				}
				case 116:
				{
					goto _ctr962;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr959;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr959;
				
			}
			goto _st0;
			
		}
		_ctr1148:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		goto _st407;
		_ctr1120:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st407;
		_ctr956:
		{st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
		}
		goto _st407;
		_ctr1132:
		{st.Month, _ = strconv.Atoi(data[pb:pb+2])
			st.Day, _ = strconv.Atoi(data[pb+2:pb+4])
		}
		goto _st407;
		_ctr1140:
		{st.Year = parse_year_2_digits(data[pb:pb+2])
		}
		goto _st407;
		_st407:
		if p == eof {
			goto _out407;
			
		}
		p+=1;
		st_case_407:
		if p == pe && p != eof {
			goto _out407;
			
		}
		if p == eof {
			goto _st407;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st9;
					
				}
				case 43:
				{
					goto _st12;
					
				}
				case 45:
				{
					goto _st13;
					
				}
				case 47:
				{
					goto _ctr967;
					
				}
				case 50:
				{
					goto _ctr90;
					
				}
				case 65:
				{
					goto _ctr968;
					
				}
				case 66:
				{
					goto _ctr969;
					
				}
				case 90:
				{
					goto _ctr970;
					
				}
				case 95:
				{
					goto _ctr967;
					
				}
				
			}
			if ( data[ p ] ) < 51 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 49 {
					goto _ctr89;
					
				}
				
			} else if ( data[ p ] ) > 57 {
				if ( data[ p ] ) > 89 {
					if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
						goto _ctr967;
						
					}
					
				} else if ( data[ p ] ) >= 67 {
					goto _ctr967;
					
				}
				
			} else {
				goto _ctr91;
				
			}
			goto _st0;
			
		}
		_ctr45:
		{return st, err }
		goto _st9;
		_ctr973:
		{for p - pb > 4 &&  data[pb] =='0' {
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
		}
		{st.Zoned = true }
		goto _st9;
		_ctr980:
		{st.Zoned = true }
		goto _st9;
		_ctr984:
		{switch p - pb {
				case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
				default:
				err = errors.New("invalid offset minute")
				return
			}
		}
		{st.Zoned = true }
		goto _st9;
		_ctr989:
		{st.ZoneName = data[pb:p]
			st.Zoned = true
		}
		{st.Zoned = true }
		goto _st9;
		_st9:
		if p == eof {
			goto _out9;
			
		}
		p+=1;
		st_case_9:
		if p == pe && p != eof {
			goto _out9;
			
		}
		if p == eof {
			goto _ctr45;
			
		} else {
			switch ( data[ p ] ) {
				case 65:
				{
					goto _st10;
					
				}
				case 66:
				{
					goto _st11;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr48:
		{return st, err }
		goto _st10;
		_st10:
		if p == eof {
			goto _out10;
			
		}
		p+=1;
		st_case_10:
		if p == pe && p != eof {
			goto _out10;
			
		}
		if p == eof {
			goto _ctr48;
			
		} else {
			if ( data[ p ] ) == 68 {
				goto _st408;
				
			}
			goto _ctr2;
			
		}
		_st408:
		if p == eof {
			goto _out408;
			
		}
		p+=1;
		st_case_408:
		if p == pe && p != eof {
			goto _out408;
			
		}
		if p == eof {
			goto _st408;
			
		} else {
			goto _st0;
			
		}
		_ctr50:
		{return st, err }
		goto _st11;
		_st11:
		if p == eof {
			goto _out11;
			
		}
		p+=1;
		st_case_11:
		if p == pe && p != eof {
			goto _out11;
			
		}
		if p == eof {
			goto _ctr50;
			
		} else {
			if ( data[ p ] ) == 67 {
				goto _st409;
				
			}
			goto _ctr2;
			
		}
		_ctr971:
		{st.Ad_bc = ADBC_BC;
		}
		goto _st409;
		_st409:
		if p == eof {
			goto _out409;
			
		}
		p+=1;
		st_case_409:
		if p == pe && p != eof {
			goto _out409;
			
		}
		if p == eof {
			goto _ctr971;
			
		} else {
			goto _st0;
			
		}
		_ctr52:
		{return st, err }
		goto _st12;
		_ctr1149:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		goto _st12;
		_ctr1121:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st12;
		_ctr992:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st12;
		_ctr1006:
		{if st.Hour > 12 {
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
		}
		goto _st12;
		_ctr1014:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st12;
		_ctr1023:
		{switch p - pb {
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
		}
		goto _st12;
		_ctr1031:
		{switch p - pb {
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
		}
		goto _st12;
		_ctr1054:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st12;
		_ctr1064:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st12;
		_ctr1073:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st12;
		_ctr1103:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st12;
		_ctr957:
		{st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
		}
		goto _st12;
		_ctr1133:
		{st.Month, _ = strconv.Atoi(data[pb:pb+2])
			st.Day, _ = strconv.Atoi(data[pb+2:pb+4])
		}
		goto _st12;
		_ctr1141:
		{st.Year = parse_year_2_digits(data[pb:pb+2])
		}
		goto _st12;
		_st12:
		if p == eof {
			goto _out12;
			
		}
		p+=1;
		st_case_12:
		if p == pe && p != eof {
			goto _out12;
			
		}
		if p == eof {
			goto _ctr52;
			
		} else {
			if ( data[ p ] ) == 50 {
				goto _ctr54;
				
			}
			if ( data[ p ] ) > 49 {
				if 51 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr55;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr53;
				
			}
			goto _ctr2;
			
		}
		_ctr53:
		{pb = p }
		goto _st410;
		_ctr57:
		{st.NegtiveZoneOffset = true }
		{pb = p }
		goto _st410;
		_ctr972:
		{for p - pb > 4 &&  data[pb] =='0' {
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
		}
		{st.Zoned = true }
		goto _st410;
		_st410:
		if p == eof {
			goto _out410;
			
		}
		p+=1;
		st_case_410:
		if p == pe && p != eof {
			goto _out410;
			
		}
		if p == eof {
			goto _ctr972;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr973;
					
				}
				case 58:
				{
					goto _ctr975;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st411;
				
			}
			goto _st0;
			
		}
		_ctr55:
		{pb = p }
		goto _st411;
		_ctr59:
		{st.NegtiveZoneOffset = true }
		{pb = p }
		goto _st411;
		_ctr976:
		{for p - pb > 4 &&  data[pb] =='0' {
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
		}
		{st.Zoned = true }
		goto _st411;
		_st411:
		if p == eof {
			goto _out411;
			
		}
		p+=1;
		st_case_411:
		if p == pe && p != eof {
			goto _out411;
			
		}
		if p == eof {
			goto _ctr976;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr973;
					
				}
				case 58:
				{
					goto _ctr975;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st412;
				
			}
			goto _st0;
			
		}
		_ctr978:
		{for p - pb > 4 &&  data[pb] =='0' {
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
		}
		{st.Zoned = true }
		goto _st412;
		_st412:
		if p == eof {
			goto _out412;
			
		}
		p+=1;
		st_case_412:
		if p == pe && p != eof {
			goto _out412;
			
		}
		if p == eof {
			goto _ctr978;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr973;
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st412;
				
			}
			goto _st0;
			
		}
		_ctr975:
		{switch p - pb {
				case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
				default:
				err = errors.New("invalid offset hour")
				return
			}
		}
		goto _st413;
		_ctr979:
		{st.Zoned = true }
		goto _st413;
		_st413:
		if p == eof {
			goto _out413;
			
		}
		p+=1;
		st_case_413:
		if p == pe && p != eof {
			goto _out413;
			
		}
		if p == eof {
			goto _ctr979;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr980;
				
			}
			if ( data[ p ] ) > 53 {
				if ( data[ p ] ) <= 57 {
					goto _ctr982;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr981;
				
			}
			goto _st0;
			
		}
		_ctr981:
		{pb = p }
		goto _st414;
		_ctr983:
		{switch p - pb {
				case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
				default:
				err = errors.New("invalid offset minute")
				return
			}
		}
		{st.Zoned = true }
		goto _st414;
		_st414:
		if p == eof {
			goto _out414;
			
		}
		p+=1;
		st_case_414:
		if p == pe && p != eof {
			goto _out414;
			
		}
		if p == eof {
			goto _ctr983;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr984;
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st415;
				
			}
			goto _st0;
			
		}
		_ctr982:
		{pb = p }
		goto _st415;
		_ctr986:
		{switch p - pb {
				case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
				default:
				err = errors.New("invalid offset minute")
				return
			}
		}
		{st.Zoned = true }
		goto _st415;
		_st415:
		if p == eof {
			goto _out415;
			
		}
		p+=1;
		st_case_415:
		if p == pe && p != eof {
			goto _out415;
			
		}
		if p == eof {
			goto _ctr986;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr984;
				
			}
			goto _st0;
			
		}
		_ctr54:
		{pb = p }
		goto _st416;
		_ctr58:
		{st.NegtiveZoneOffset = true }
		{pb = p }
		goto _st416;
		_ctr987:
		{for p - pb > 4 &&  data[pb] =='0' {
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
		}
		{st.Zoned = true }
		goto _st416;
		_st416:
		if p == eof {
			goto _out416;
			
		}
		p+=1;
		st_case_416:
		if p == pe && p != eof {
			goto _out416;
			
		}
		if p == eof {
			goto _ctr987;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr973;
					
				}
				case 58:
				{
					goto _ctr975;
					
				}
				
			}
			if ( data[ p ] ) > 51 {
				if ( data[ p ] ) <= 57 {
					goto _st412;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _st411;
				
			}
			goto _st0;
			
		}
		_ctr56:
		{return st, err }
		goto _st13;
		_ctr1150:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		goto _st13;
		_ctr1122:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st13;
		_ctr993:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st13;
		_ctr1007:
		{if st.Hour > 12 {
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
		}
		goto _st13;
		_ctr1015:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st13;
		_ctr1024:
		{switch p - pb {
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
		}
		goto _st13;
		_ctr1032:
		{switch p - pb {
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
		}
		goto _st13;
		_ctr1055:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st13;
		_ctr1065:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st13;
		_ctr1074:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st13;
		_ctr1104:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st13;
		_ctr958:
		{st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
		}
		goto _st13;
		_ctr1134:
		{st.Month, _ = strconv.Atoi(data[pb:pb+2])
			st.Day, _ = strconv.Atoi(data[pb+2:pb+4])
		}
		goto _st13;
		_ctr1142:
		{st.Year = parse_year_2_digits(data[pb:pb+2])
		}
		goto _st13;
		_st13:
		if p == eof {
			goto _out13;
			
		}
		p+=1;
		st_case_13:
		if p == pe && p != eof {
			goto _out13;
			
		}
		if p == eof {
			goto _ctr56;
			
		} else {
			if ( data[ p ] ) == 50 {
				goto _ctr58;
				
			}
			if ( data[ p ] ) > 49 {
				if 51 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr59;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr57;
				
			}
			goto _ctr2;
			
		}
		_ctr60:
		{return st, err }
		goto _st14;
		_ctr967:
		{pb = p }
		goto _st14;
		_ctr1151:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		{pb = p }
		goto _st14;
		_ctr994:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st14;
		_ctr1016:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st14;
		_ctr1026:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st14;
		_ctr1033:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st14;
		_ctr1056:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st14;
		_ctr1066:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st14;
		_ctr1076:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st14;
		_ctr1106:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st14;
		_ctr959:
		{st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
		}
		{pb = p }
		goto _st14;
		_ctr1123:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		{pb = p }
		goto _st14;
		_ctr1135:
		{st.Month, _ = strconv.Atoi(data[pb:pb+2])
			st.Day, _ = strconv.Atoi(data[pb+2:pb+4])
		}
		{pb = p }
		goto _st14;
		_ctr1143:
		{st.Year = parse_year_2_digits(data[pb:pb+2])
		}
		{pb = p }
		goto _st14;
		_st14:
		if p == eof {
			goto _out14;
			
		}
		p+=1;
		st_case_14:
		if p == pe && p != eof {
			goto _out14;
			
		}
		if p == eof {
			goto _ctr60;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st15;
					
				}
				case 95:
				{
					goto _st15;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st15;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st15;
				
			}
			goto _ctr2;
			
		}
		_ctr62:
		{return st, err }
		goto _st15;
		_ctr1116:
		{pb = p }
		goto _st15;
		_st15:
		if p == eof {
			goto _out15;
			
		}
		p+=1;
		st_case_15:
		if p == pe && p != eof {
			goto _out15;
			
		}
		if p == eof {
			goto _ctr62;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st417;
					
				}
				case 95:
				{
					goto _st417;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st417;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st417;
				
			}
			goto _ctr2;
			
		}
		_ctr988:
		{st.ZoneName = data[pb:p]
			st.Zoned = true
		}
		{st.Zoned = true }
		goto _st417;
		_ctr1008:
		{if st.Hour > 12 {
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
		}
		{pb = p }
		goto _st417;
		_st417:
		if p == eof {
			goto _out417;
			
		}
		p+=1;
		st_case_417:
		if p == pe && p != eof {
			goto _out417;
			
		}
		if p == eof {
			goto _ctr988;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr989;
					
				}
				case 47:
				{
					goto _st417;
					
				}
				case 95:
				{
					goto _st417;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st417;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st417;
				
			}
			goto _st0;
			
		}
		_ctr89:
		{pb = p }
		goto _st418;
		_ctr990:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st418;
		_st418:
		if p == eof {
			goto _out418;
			
		}
		p+=1;
		st_case_418:
		if p == pe && p != eof {
			goto _out418;
			
		}
		if p == eof {
			goto _ctr990;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr991;
					
				}
				case 43:
				{
					goto _ctr992;
					
				}
				case 45:
				{
					goto _ctr993;
					
				}
				case 47:
				{
					goto _ctr994;
					
				}
				case 58:
				{
					goto _ctr996;
					
				}
				case 65:
				{
					goto _ctr997;
					
				}
				case 80:
				{
					goto _ctr997;
					
				}
				case 90:
				{
					goto _ctr998;
					
				}
				case 95:
				{
					goto _ctr994;
					
				}
				case 97:
				{
					goto _ctr999;
					
				}
				case 112:
				{
					goto _ctr999;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st425;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr994;
					
				}
				
			} else {
				goto _ctr994;
				
			}
			goto _st0;
			
		}
		_ctr991:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st419;
		_ctr1013:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st419;
		_ctr1082:
		{switch p - pb {
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
		}
		goto _st419;
		_ctr1053:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st419;
		_ctr1063:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st419;
		_ctr1072:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st419;
		_ctr1102:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st419;
		_st419:
		if p == eof {
			goto _out419;
			
		}
		p+=1;
		st_case_419:
		if p == pe && p != eof {
			goto _out419;
			
		}
		if p == eof {
			goto _st419;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st9;
					
				}
				case 43:
				{
					goto _st12;
					
				}
				case 45:
				{
					goto _st13;
					
				}
				case 47:
				{
					goto _ctr967;
					
				}
				case 65:
				{
					goto _ctr1001;
					
				}
				case 66:
				{
					goto _ctr969;
					
				}
				case 80:
				{
					goto _ctr1002;
					
				}
				case 90:
				{
					goto _ctr970;
					
				}
				case 95:
				{
					goto _ctr967;
					
				}
				case 97:
				{
					goto _ctr1003;
					
				}
				case 112:
				{
					goto _ctr1003;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr967;
					
				}
				
			} else if ( data[ p ] ) >= 67 {
				goto _ctr967;
				
			}
			goto _st0;
			
		}
		_ctr64:
		{return st, err }
		goto _st16;
		_ctr1001:
		{pb = p }
		goto _st16;
		_st16:
		if p == eof {
			goto _out16;
			
		}
		p+=1;
		st_case_16:
		if p == pe && p != eof {
			goto _out16;
			
		}
		if p == eof {
			goto _ctr64;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st15;
					
				}
				case 68:
				{
					goto _st420;
					
				}
				case 77:
				{
					goto _st421;
					
				}
				case 95:
				{
					goto _st15;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st15;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st15;
				
			}
			goto _ctr2;
			
		}
		_st420:
		if p == eof {
			goto _out420;
			
		}
		p+=1;
		st_case_420:
		if p == pe && p != eof {
			goto _out420;
			
		}
		if p == eof {
			goto _st420;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st417;
					
				}
				case 95:
				{
					goto _st417;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st417;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st417;
				
			}
			goto _st0;
			
		}
		_ctr1004:
		{if st.Hour > 12 {
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
		}
		goto _st421;
		_st421:
		if p == eof {
			goto _out421;
			
		}
		p+=1;
		st_case_421:
		if p == pe && p != eof {
			goto _out421;
			
		}
		if p == eof {
			goto _ctr1004;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1005;
					
				}
				case 43:
				{
					goto _ctr1006;
					
				}
				case 45:
				{
					goto _ctr1007;
					
				}
				case 47:
				{
					goto _ctr1008;
					
				}
				case 95:
				{
					goto _ctr1008;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1008;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr1008;
				
			}
			goto _st0;
			
		}
		_ctr1005:
		{if st.Hour > 12 {
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
		}
		goto _st422;
		_ctr1022:
		{switch p - pb {
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
		}
		goto _st422;
		_ctr1030:
		{switch p - pb {
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
		}
		goto _st422;
		_st422:
		if p == eof {
			goto _out422;
			
		}
		p+=1;
		st_case_422:
		if p == pe && p != eof {
			goto _out422;
			
		}
		if p == eof {
			goto _st422;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st9;
					
				}
				case 43:
				{
					goto _st12;
					
				}
				case 45:
				{
					goto _st13;
					
				}
				case 47:
				{
					goto _ctr967;
					
				}
				case 65:
				{
					goto _ctr968;
					
				}
				case 66:
				{
					goto _ctr969;
					
				}
				case 90:
				{
					goto _ctr970;
					
				}
				case 95:
				{
					goto _ctr967;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr967;
					
				}
				
			} else if ( data[ p ] ) >= 67 {
				goto _ctr967;
				
			}
			goto _st0;
			
		}
		_ctr67:
		{return st, err }
		goto _st17;
		_ctr968:
		{pb = p }
		goto _st17;
		_st17:
		if p == eof {
			goto _out17;
			
		}
		p+=1;
		st_case_17:
		if p == pe && p != eof {
			goto _out17;
			
		}
		if p == eof {
			goto _ctr67;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st15;
					
				}
				case 68:
				{
					goto _st420;
					
				}
				case 95:
				{
					goto _st15;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st15;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st15;
				
			}
			goto _ctr2;
			
		}
		_ctr68:
		{return st, err }
		goto _st18;
		_ctr969:
		{pb = p }
		goto _st18;
		_st18:
		if p == eof {
			goto _out18;
			
		}
		p+=1;
		st_case_18:
		if p == pe && p != eof {
			goto _out18;
			
		}
		if p == eof {
			goto _ctr68;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st15;
					
				}
				case 67:
				{
					goto _st423;
					
				}
				case 95:
				{
					goto _st15;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st15;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st15;
				
			}
			goto _ctr2;
			
		}
		_ctr1010:
		{st.Ad_bc = ADBC_BC;
		}
		goto _st423;
		_st423:
		if p == eof {
			goto _out423;
			
		}
		p+=1;
		st_case_423:
		if p == pe && p != eof {
			goto _out423;
			
		}
		if p == eof {
			goto _ctr1010;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st417;
					
				}
				case 95:
				{
					goto _st417;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st417;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st417;
				
			}
			goto _st0;
			
		}
		_ctr970:
		{pb = p }
		goto _st424;
		_ctr1153:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		{pb = p }
		goto _st424;
		_ctr998:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st424;
		_ctr1011:
		{st.Zoned = true }
		goto _st424;
		_ctr1019:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st424;
		_ctr1028:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st424;
		_ctr1035:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st424;
		_ctr1060:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st424;
		_ctr1069:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st424;
		_ctr1079:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st424;
		_ctr1108:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st424;
		_ctr961:
		{st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
		}
		{pb = p }
		goto _st424;
		_ctr1125:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		{pb = p }
		goto _st424;
		_ctr1137:
		{st.Month, _ = strconv.Atoi(data[pb:pb+2])
			st.Day, _ = strconv.Atoi(data[pb+2:pb+4])
		}
		{pb = p }
		goto _st424;
		_ctr1145:
		{st.Year = parse_year_2_digits(data[pb:pb+2])
		}
		{pb = p }
		goto _st424;
		_st424:
		if p == eof {
			goto _out424;
			
		}
		p+=1;
		st_case_424:
		if p == pe && p != eof {
			goto _out424;
			
		}
		if p == eof {
			goto _ctr1011;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr980;
					
				}
				case 47:
				{
					goto _st15;
					
				}
				case 95:
				{
					goto _st15;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st15;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st15;
				
			}
			goto _st0;
			
		}
		_ctr70:
		{return st, err }
		goto _st19;
		_ctr1002:
		{pb = p }
		goto _st19;
		_ctr997:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st19;
		_ctr1018:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st19;
		_ctr1084:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st19;
		_ctr1059:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st19;
		_ctr1068:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st19;
		_ctr1078:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st19;
		_ctr1107:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st19;
		_st19:
		if p == eof {
			goto _out19;
			
		}
		p+=1;
		st_case_19:
		if p == pe && p != eof {
			goto _out19;
			
		}
		if p == eof {
			goto _ctr70;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st15;
					
				}
				case 77:
				{
					goto _st421;
					
				}
				case 95:
				{
					goto _st15;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st15;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st15;
				
			}
			goto _ctr2;
			
		}
		_ctr71:
		{return st, err }
		goto _st20;
		_ctr1003:
		{pb = p }
		goto _st20;
		_ctr999:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st20;
		_ctr1020:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st20;
		_ctr1085:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st20;
		_ctr1061:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st20;
		_ctr1070:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st20;
		_ctr1080:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st20;
		_ctr1109:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st20;
		_st20:
		if p == eof {
			goto _out20;
			
		}
		p+=1;
		st_case_20:
		if p == pe && p != eof {
			goto _out20;
			
		}
		if p == eof {
			goto _ctr71;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st15;
					
				}
				case 95:
				{
					goto _st15;
					
				}
				case 109:
				{
					goto _st421;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st15;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st15;
				
			}
			goto _ctr2;
			
		}
		_ctr1012:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st425;
		_st425:
		if p == eof {
			goto _out425;
			
		}
		p+=1;
		st_case_425:
		if p == pe && p != eof {
			goto _out425;
			
		}
		if p == eof {
			goto _ctr1012;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1013;
					
				}
				case 43:
				{
					goto _ctr1014;
					
				}
				case 45:
				{
					goto _ctr1015;
					
				}
				case 47:
				{
					goto _ctr1016;
					
				}
				case 58:
				{
					goto _ctr1017;
					
				}
				case 65:
				{
					goto _ctr1018;
					
				}
				case 80:
				{
					goto _ctr1018;
					
				}
				case 90:
				{
					goto _ctr1019;
					
				}
				case 95:
				{
					goto _ctr1016;
					
				}
				case 97:
				{
					goto _ctr1020;
					
				}
				case 112:
				{
					goto _ctr1020;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st21;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1016;
					
				}
				
			} else {
				goto _ctr1016;
				
			}
			goto _st0;
			
		}
		_ctr72:
		{return st, err }
		goto _st21;
		_st21:
		if p == eof {
			goto _out21;
			
		}
		p+=1;
		st_case_21:
		if p == pe && p != eof {
			goto _out21;
			
		}
		if p == eof {
			goto _ctr72;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st426;
				
			}
			goto _ctr2;
			
		}
		_ctr1021:
		{switch p - pb {
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
		}
		goto _st426;
		_st426:
		if p == eof {
			goto _out426;
			
		}
		p+=1;
		st_case_426:
		if p == pe && p != eof {
			goto _out426;
			
		}
		if p == eof {
			goto _ctr1021;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1022;
					
				}
				case 43:
				{
					goto _ctr1023;
					
				}
				case 45:
				{
					goto _ctr1024;
					
				}
				case 46:
				{
					goto _ctr1025;
					
				}
				case 47:
				{
					goto _ctr1026;
					
				}
				case 90:
				{
					goto _ctr1028;
					
				}
				case 95:
				{
					goto _ctr1026;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st23;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1026;
					
				}
				
			} else {
				goto _ctr1026;
				
			}
			goto _st0;
			
		}
		_ctr74:
		{return st, err }
		goto _st22;
		_ctr1025:
		{switch p - pb {
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
		}
		goto _st22;
		_st22:
		if p == eof {
			goto _out22;
			
		}
		p+=1;
		st_case_22:
		if p == pe && p != eof {
			goto _out22;
			
		}
		if p == eof {
			goto _ctr74;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _ctr75;
				
			}
			goto _ctr2;
			
		}
		_ctr75:
		{pb = p }
		goto _st427;
		_ctr1029:
		{switch p - pb {
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
		}
		goto _st427;
		_st427:
		if p == eof {
			goto _out427;
			
		}
		p+=1;
		st_case_427:
		if p == pe && p != eof {
			goto _out427;
			
		}
		if p == eof {
			goto _ctr1029;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1030;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st428;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1036:
		{switch p - pb {
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
		}
		goto _st428;
		_st428:
		if p == eof {
			goto _out428;
			
		}
		p+=1;
		st_case_428:
		if p == pe && p != eof {
			goto _out428;
			
		}
		if p == eof {
			goto _ctr1036;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1030;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st429;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1038:
		{switch p - pb {
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
		}
		goto _st429;
		_st429:
		if p == eof {
			goto _out429;
			
		}
		p+=1;
		st_case_429:
		if p == pe && p != eof {
			goto _out429;
			
		}
		if p == eof {
			goto _ctr1038;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1030;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st430;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1040:
		{switch p - pb {
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
		}
		goto _st430;
		_st430:
		if p == eof {
			goto _out430;
			
		}
		p+=1;
		st_case_430:
		if p == pe && p != eof {
			goto _out430;
			
		}
		if p == eof {
			goto _ctr1040;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1030;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st431;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1042:
		{switch p - pb {
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
		}
		goto _st431;
		_st431:
		if p == eof {
			goto _out431;
			
		}
		p+=1;
		st_case_431:
		if p == pe && p != eof {
			goto _out431;
			
		}
		if p == eof {
			goto _ctr1042;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1030;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st432;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1044:
		{switch p - pb {
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
		}
		goto _st432;
		_st432:
		if p == eof {
			goto _out432;
			
		}
		p+=1;
		st_case_432:
		if p == pe && p != eof {
			goto _out432;
			
		}
		if p == eof {
			goto _ctr1044;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1030;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st433;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1046:
		{switch p - pb {
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
		}
		goto _st433;
		_st433:
		if p == eof {
			goto _out433;
			
		}
		p+=1;
		st_case_433:
		if p == pe && p != eof {
			goto _out433;
			
		}
		if p == eof {
			goto _ctr1046;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1030;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st434;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1048:
		{switch p - pb {
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
		}
		goto _st434;
		_st434:
		if p == eof {
			goto _out434;
			
		}
		p+=1;
		st_case_434:
		if p == pe && p != eof {
			goto _out434;
			
		}
		if p == eof {
			goto _ctr1048;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1030;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st435;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1050:
		{switch p - pb {
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
		}
		goto _st435;
		_st435:
		if p == eof {
			goto _out435;
			
		}
		p+=1;
		st_case_435:
		if p == pe && p != eof {
			goto _out435;
			
		}
		if p == eof {
			goto _ctr1050;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1030;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr76:
		{return st, err }
		goto _st23;
		_st23:
		if p == eof {
			goto _out23;
			
		}
		p+=1;
		st_case_23:
		if p == pe && p != eof {
			goto _out23;
			
		}
		if p == eof {
			goto _ctr76;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st436;
				
			}
			goto _ctr2;
			
		}
		_ctr1051:
		{switch p - pb {
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
		}
		goto _st436;
		_st436:
		if p == eof {
			goto _out436;
			
		}
		p+=1;
		st_case_436:
		if p == pe && p != eof {
			goto _out436;
			
		}
		if p == eof {
			goto _ctr1051;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1022;
					
				}
				case 43:
				{
					goto _ctr1023;
					
				}
				case 45:
				{
					goto _ctr1024;
					
				}
				case 46:
				{
					goto _ctr1025;
					
				}
				case 47:
				{
					goto _ctr1026;
					
				}
				case 90:
				{
					goto _ctr1028;
					
				}
				case 95:
				{
					goto _ctr1026;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1026;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr1026;
				
			}
			goto _st0;
			
		}
		_ctr78:
		{return st, err }
		goto _st24;
		_ctr996:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st24;
		_ctr1017:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st24;
		_st24:
		if p == eof {
			goto _out24;
			
		}
		p+=1;
		st_case_24:
		if p == pe && p != eof {
			goto _out24;
			
		}
		if p == eof {
			goto _ctr78;
			
		} else {
			if ( data[ p ] ) > 53 {
				if ( data[ p ] ) <= 57 {
					goto _ctr80;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr79;
				
			}
			goto _ctr2;
			
		}
		_ctr79:
		{pb = p }
		goto _st437;
		_ctr1052:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st437;
		_st437:
		if p == eof {
			goto _out437;
			
		}
		p+=1;
		st_case_437:
		if p == pe && p != eof {
			goto _out437;
			
		}
		if p == eof {
			goto _ctr1052;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1053;
					
				}
				case 43:
				{
					goto _ctr1054;
					
				}
				case 45:
				{
					goto _ctr1055;
					
				}
				case 47:
				{
					goto _ctr1056;
					
				}
				case 58:
				{
					goto _ctr1058;
					
				}
				case 65:
				{
					goto _ctr1059;
					
				}
				case 80:
				{
					goto _ctr1059;
					
				}
				case 90:
				{
					goto _ctr1060;
					
				}
				case 95:
				{
					goto _ctr1056;
					
				}
				case 97:
				{
					goto _ctr1061;
					
				}
				case 112:
				{
					goto _ctr1061;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st438;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1056;
					
				}
				
			} else {
				goto _ctr1056;
				
			}
			goto _st0;
			
		}
		_ctr1062:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st438;
		_st438:
		if p == eof {
			goto _out438;
			
		}
		p+=1;
		st_case_438:
		if p == pe && p != eof {
			goto _out438;
			
		}
		if p == eof {
			goto _ctr1062;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1063;
					
				}
				case 43:
				{
					goto _ctr1064;
					
				}
				case 45:
				{
					goto _ctr1065;
					
				}
				case 47:
				{
					goto _ctr1066;
					
				}
				case 58:
				{
					goto _ctr1067;
					
				}
				case 65:
				{
					goto _ctr1068;
					
				}
				case 80:
				{
					goto _ctr1068;
					
				}
				case 90:
				{
					goto _ctr1069;
					
				}
				case 95:
				{
					goto _ctr1066;
					
				}
				case 97:
				{
					goto _ctr1070;
					
				}
				case 112:
				{
					goto _ctr1070;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1066;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr1066;
				
			}
			goto _st0;
			
		}
		_ctr81:
		{return st, err }
		goto _st25;
		_ctr1058:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st25;
		_ctr1067:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st25;
		_st25:
		if p == eof {
			goto _out25;
			
		}
		p+=1;
		st_case_25:
		if p == pe && p != eof {
			goto _out25;
			
		}
		if p == eof {
			goto _ctr81;
			
		} else {
			if ( data[ p ] ) > 53 {
				if ( data[ p ] ) <= 57 {
					goto _ctr83;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr82;
				
			}
			goto _ctr2;
			
		}
		_ctr82:
		{pb = p }
		goto _st439;
		_ctr1071:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st439;
		_st439:
		if p == eof {
			goto _out439;
			
		}
		p+=1;
		st_case_439:
		if p == pe && p != eof {
			goto _out439;
			
		}
		if p == eof {
			goto _ctr1071;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1072;
					
				}
				case 43:
				{
					goto _ctr1073;
					
				}
				case 45:
				{
					goto _ctr1074;
					
				}
				case 46:
				{
					goto _ctr1075;
					
				}
				case 47:
				{
					goto _ctr1076;
					
				}
				case 65:
				{
					goto _ctr1078;
					
				}
				case 80:
				{
					goto _ctr1078;
					
				}
				case 90:
				{
					goto _ctr1079;
					
				}
				case 95:
				{
					goto _ctr1076;
					
				}
				case 97:
				{
					goto _ctr1080;
					
				}
				case 112:
				{
					goto _ctr1080;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st449;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1076;
					
				}
				
			} else {
				goto _ctr1076;
				
			}
			goto _st0;
			
		}
		_ctr84:
		{return st, err }
		goto _st26;
		_ctr1075:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st26;
		_ctr1105:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st26;
		_st26:
		if p == eof {
			goto _out26;
			
		}
		p+=1;
		st_case_26:
		if p == pe && p != eof {
			goto _out26;
			
		}
		if p == eof {
			goto _ctr84;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _ctr85;
				
			}
			goto _ctr2;
			
		}
		_ctr85:
		{pb = p }
		goto _st440;
		_ctr1081:
		{switch p - pb {
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
		}
		goto _st440;
		_st440:
		if p == eof {
			goto _out440;
			
		}
		p+=1;
		st_case_440:
		if p == pe && p != eof {
			goto _out440;
			
		}
		if p == eof {
			goto _ctr1081;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1082;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 65:
				{
					goto _ctr1084;
					
				}
				case 80:
				{
					goto _ctr1084;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				case 97:
				{
					goto _ctr1085;
					
				}
				case 112:
				{
					goto _ctr1085;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st441;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1086:
		{switch p - pb {
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
		}
		goto _st441;
		_st441:
		if p == eof {
			goto _out441;
			
		}
		p+=1;
		st_case_441:
		if p == pe && p != eof {
			goto _out441;
			
		}
		if p == eof {
			goto _ctr1086;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1082;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 65:
				{
					goto _ctr1084;
					
				}
				case 80:
				{
					goto _ctr1084;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				case 97:
				{
					goto _ctr1085;
					
				}
				case 112:
				{
					goto _ctr1085;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st442;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1088:
		{switch p - pb {
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
		}
		goto _st442;
		_st442:
		if p == eof {
			goto _out442;
			
		}
		p+=1;
		st_case_442:
		if p == pe && p != eof {
			goto _out442;
			
		}
		if p == eof {
			goto _ctr1088;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1082;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 65:
				{
					goto _ctr1084;
					
				}
				case 80:
				{
					goto _ctr1084;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				case 97:
				{
					goto _ctr1085;
					
				}
				case 112:
				{
					goto _ctr1085;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st443;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1090:
		{switch p - pb {
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
		}
		goto _st443;
		_st443:
		if p == eof {
			goto _out443;
			
		}
		p+=1;
		st_case_443:
		if p == pe && p != eof {
			goto _out443;
			
		}
		if p == eof {
			goto _ctr1090;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1082;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 65:
				{
					goto _ctr1084;
					
				}
				case 80:
				{
					goto _ctr1084;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				case 97:
				{
					goto _ctr1085;
					
				}
				case 112:
				{
					goto _ctr1085;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st444;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1092:
		{switch p - pb {
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
		}
		goto _st444;
		_st444:
		if p == eof {
			goto _out444;
			
		}
		p+=1;
		st_case_444:
		if p == pe && p != eof {
			goto _out444;
			
		}
		if p == eof {
			goto _ctr1092;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1082;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 65:
				{
					goto _ctr1084;
					
				}
				case 80:
				{
					goto _ctr1084;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				case 97:
				{
					goto _ctr1085;
					
				}
				case 112:
				{
					goto _ctr1085;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st445;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1094:
		{switch p - pb {
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
		}
		goto _st445;
		_st445:
		if p == eof {
			goto _out445;
			
		}
		p+=1;
		st_case_445:
		if p == pe && p != eof {
			goto _out445;
			
		}
		if p == eof {
			goto _ctr1094;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1082;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 65:
				{
					goto _ctr1084;
					
				}
				case 80:
				{
					goto _ctr1084;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				case 97:
				{
					goto _ctr1085;
					
				}
				case 112:
				{
					goto _ctr1085;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st446;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1096:
		{switch p - pb {
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
		}
		goto _st446;
		_st446:
		if p == eof {
			goto _out446;
			
		}
		p+=1;
		st_case_446:
		if p == pe && p != eof {
			goto _out446;
			
		}
		if p == eof {
			goto _ctr1096;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1082;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 65:
				{
					goto _ctr1084;
					
				}
				case 80:
				{
					goto _ctr1084;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				case 97:
				{
					goto _ctr1085;
					
				}
				case 112:
				{
					goto _ctr1085;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st447;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1098:
		{switch p - pb {
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
		}
		goto _st447;
		_st447:
		if p == eof {
			goto _out447;
			
		}
		p+=1;
		st_case_447:
		if p == pe && p != eof {
			goto _out447;
			
		}
		if p == eof {
			goto _ctr1098;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1082;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 65:
				{
					goto _ctr1084;
					
				}
				case 80:
				{
					goto _ctr1084;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				case 97:
				{
					goto _ctr1085;
					
				}
				case 112:
				{
					goto _ctr1085;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st448;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1100:
		{switch p - pb {
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
		}
		goto _st448;
		_st448:
		if p == eof {
			goto _out448;
			
		}
		p+=1;
		st_case_448:
		if p == pe && p != eof {
			goto _out448;
			
		}
		if p == eof {
			goto _ctr1100;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1082;
					
				}
				case 43:
				{
					goto _ctr1031;
					
				}
				case 45:
				{
					goto _ctr1032;
					
				}
				case 47:
				{
					goto _ctr1033;
					
				}
				case 65:
				{
					goto _ctr1084;
					
				}
				case 80:
				{
					goto _ctr1084;
					
				}
				case 90:
				{
					goto _ctr1035;
					
				}
				case 95:
				{
					goto _ctr1033;
					
				}
				case 97:
				{
					goto _ctr1085;
					
				}
				case 112:
				{
					goto _ctr1085;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1033;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr1033;
				
			}
			goto _st0;
			
		}
		_ctr1101:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st449;
		_st449:
		if p == eof {
			goto _out449;
			
		}
		p+=1;
		st_case_449:
		if p == pe && p != eof {
			goto _out449;
			
		}
		if p == eof {
			goto _ctr1101;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1102;
					
				}
				case 43:
				{
					goto _ctr1103;
					
				}
				case 45:
				{
					goto _ctr1104;
					
				}
				case 46:
				{
					goto _ctr1105;
					
				}
				case 47:
				{
					goto _ctr1106;
					
				}
				case 65:
				{
					goto _ctr1107;
					
				}
				case 80:
				{
					goto _ctr1107;
					
				}
				case 90:
				{
					goto _ctr1108;
					
				}
				case 95:
				{
					goto _ctr1106;
					
				}
				case 97:
				{
					goto _ctr1109;
					
				}
				case 112:
				{
					goto _ctr1109;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1106;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr1106;
				
			}
			goto _st0;
			
		}
		_ctr83:
		{pb = p }
		goto _st450;
		_ctr1110:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st450;
		_st450:
		if p == eof {
			goto _out450;
			
		}
		p+=1;
		st_case_450:
		if p == pe && p != eof {
			goto _out450;
			
		}
		if p == eof {
			goto _ctr1110;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1072;
					
				}
				case 43:
				{
					goto _ctr1073;
					
				}
				case 45:
				{
					goto _ctr1074;
					
				}
				case 46:
				{
					goto _ctr1075;
					
				}
				case 47:
				{
					goto _ctr1076;
					
				}
				case 65:
				{
					goto _ctr1078;
					
				}
				case 80:
				{
					goto _ctr1078;
					
				}
				case 90:
				{
					goto _ctr1079;
					
				}
				case 95:
				{
					goto _ctr1076;
					
				}
				case 97:
				{
					goto _ctr1080;
					
				}
				case 112:
				{
					goto _ctr1080;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1076;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr1076;
				
			}
			goto _st0;
			
		}
		_ctr80:
		{pb = p }
		goto _st451;
		_ctr1111:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st451;
		_st451:
		if p == eof {
			goto _out451;
			
		}
		p+=1;
		st_case_451:
		if p == pe && p != eof {
			goto _out451;
			
		}
		if p == eof {
			goto _ctr1111;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1053;
					
				}
				case 43:
				{
					goto _ctr1054;
					
				}
				case 45:
				{
					goto _ctr1055;
					
				}
				case 47:
				{
					goto _ctr1056;
					
				}
				case 58:
				{
					goto _ctr1058;
					
				}
				case 65:
				{
					goto _ctr1059;
					
				}
				case 80:
				{
					goto _ctr1059;
					
				}
				case 90:
				{
					goto _ctr1060;
					
				}
				case 95:
				{
					goto _ctr1056;
					
				}
				case 97:
				{
					goto _ctr1061;
					
				}
				case 112:
				{
					goto _ctr1061;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1056;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr1056;
				
			}
			goto _st0;
			
		}
		_ctr90:
		{pb = p }
		goto _st452;
		_ctr1112:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st452;
		_st452:
		if p == eof {
			goto _out452;
			
		}
		p+=1;
		st_case_452:
		if p == pe && p != eof {
			goto _out452;
			
		}
		if p == eof {
			goto _ctr1112;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr991;
					
				}
				case 43:
				{
					goto _ctr992;
					
				}
				case 45:
				{
					goto _ctr993;
					
				}
				case 47:
				{
					goto _ctr994;
					
				}
				case 58:
				{
					goto _ctr996;
					
				}
				case 65:
				{
					goto _ctr997;
					
				}
				case 80:
				{
					goto _ctr997;
					
				}
				case 90:
				{
					goto _ctr998;
					
				}
				case 95:
				{
					goto _ctr994;
					
				}
				case 97:
				{
					goto _ctr999;
					
				}
				case 112:
				{
					goto _ctr999;
					
				}
				
			}
			if ( data[ p ] ) < 52 {
				if 48 <= ( data[ p ] ) {
					goto _st425;
					
				}
				
			} else if ( data[ p ] ) > 57 {
				if ( data[ p ] ) > 89 {
					if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
						goto _ctr994;
						
					}
					
				} else if ( data[ p ] ) >= 66 {
					goto _ctr994;
					
				}
				
			} else {
				goto _st27;
				
			}
			goto _st0;
			
		}
		_ctr86:
		{return st, err }
		goto _st27;
		_st27:
		if p == eof {
			goto _out27;
			
		}
		p+=1;
		st_case_27:
		if p == pe && p != eof {
			goto _out27;
			
		}
		if p == eof {
			goto _ctr86;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st21;
				
			}
			goto _ctr2;
			
		}
		_ctr91:
		{pb = p }
		goto _st453;
		_ctr1114:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st453;
		_st453:
		if p == eof {
			goto _out453;
			
		}
		p+=1;
		st_case_453:
		if p == pe && p != eof {
			goto _out453;
			
		}
		if p == eof {
			goto _ctr1114;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr991;
					
				}
				case 43:
				{
					goto _ctr992;
					
				}
				case 45:
				{
					goto _ctr993;
					
				}
				case 47:
				{
					goto _ctr994;
					
				}
				case 58:
				{
					goto _ctr996;
					
				}
				case 65:
				{
					goto _ctr997;
					
				}
				case 80:
				{
					goto _ctr997;
					
				}
				case 90:
				{
					goto _ctr998;
					
				}
				case 95:
				{
					goto _ctr994;
					
				}
				case 97:
				{
					goto _ctr999;
					
				}
				case 112:
				{
					goto _ctr999;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st27;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr994;
					
				}
				
			} else {
				goto _ctr994;
				
			}
			goto _st0;
			
		}
		_ctr1152:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		{pb = p }
		goto _st454;
		_ctr960:
		{st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
		}
		{pb = p }
		goto _st454;
		_ctr1124:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		{pb = p }
		goto _st454;
		_ctr1136:
		{st.Month, _ = strconv.Atoi(data[pb:pb+2])
			st.Day, _ = strconv.Atoi(data[pb+2:pb+4])
		}
		{pb = p }
		goto _st454;
		_ctr1144:
		{st.Year = parse_year_2_digits(data[pb:pb+2])
		}
		{pb = p }
		goto _st454;
		_st454:
		if p == eof {
			goto _out454;
			
		}
		p+=1;
		st_case_454:
		if p == pe && p != eof {
			goto _out454;
			
		}
		if p == eof {
			goto _st454;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st9;
					
				}
				case 43:
				{
					goto _st12;
					
				}
				case 45:
				{
					goto _st13;
					
				}
				case 47:
				{
					goto _ctr1116;
					
				}
				case 50:
				{
					goto _ctr90;
					
				}
				case 90:
				{
					goto _ctr1117;
					
				}
				case 95:
				{
					goto _ctr1116;
					
				}
				
			}
			if ( data[ p ] ) < 51 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 49 {
					goto _ctr89;
					
				}
				
			} else if ( data[ p ] ) > 57 {
				if ( data[ p ] ) > 89 {
					if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
						goto _ctr1116;
						
					}
					
				} else if ( data[ p ] ) >= 65 {
					goto _ctr1116;
					
				}
				
			} else {
				goto _ctr91;
				
			}
			goto _st0;
			
		}
		_ctr1117:
		{pb = p }
		goto _st455;
		_ctr1118:
		{st.Zoned = true }
		goto _st455;
		_st455:
		if p == eof {
			goto _out455;
			
		}
		p+=1;
		st_case_455:
		if p == pe && p != eof {
			goto _out455;
			
		}
		if p == eof {
			goto _ctr1118;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr980;
					
				}
				case 47:
				{
					goto _st417;
					
				}
				case 95:
				{
					goto _st417;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st417;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st417;
				
			}
			goto _st0;
			
		}
		_ctr88:
		{return st, err }
		goto _st28;
		_ctr1154:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		{pb = p }
		goto _st28;
		_ctr962:
		{st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
		}
		{pb = p }
		goto _st28;
		_ctr1126:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		{pb = p }
		goto _st28;
		_ctr1138:
		{st.Month, _ = strconv.Atoi(data[pb:pb+2])
			st.Day, _ = strconv.Atoi(data[pb+2:pb+4])
		}
		{pb = p }
		goto _st28;
		_ctr1146:
		{st.Year = parse_year_2_digits(data[pb:pb+2])
		}
		{pb = p }
		goto _st28;
		_st28:
		if p == eof {
			goto _out28;
			
		}
		p+=1;
		st_case_28:
		if p == pe && p != eof {
			goto _out28;
			
		}
		if p == eof {
			goto _ctr88;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st15;
					
				}
				case 50:
				{
					goto _ctr90;
					
				}
				case 95:
				{
					goto _st15;
					
				}
				
			}
			if ( data[ p ] ) < 51 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 49 {
					goto _ctr89;
					
				}
				
			} else if ( data[ p ] ) > 57 {
				if ( data[ p ] ) > 90 {
					if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
						goto _st15;
						
					}
					
				} else if ( data[ p ] ) >= 65 {
					goto _st15;
					
				}
				
			} else {
				goto _ctr91;
				
			}
			goto _ctr2;
			
		}
		_ctr92:
		{return st, err }
		goto _st29;
		_st29:
		if p == eof {
			goto _out29;
			
		}
		p+=1;
		st_case_29:
		if p == pe && p != eof {
			goto _out29;
			
		}
		if p == eof {
			goto _ctr92;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr93;
					
				}
				case 47:
				{
					goto _ctr93;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st406;
				
			}
			goto _ctr2;
			
		}
		_ctr94:
		{return st, err }
		goto _st30;
		_ctr93:
		{st.Month, _ = strconv.Atoi(data[pb:p])
		}
		goto _st30;
		_ctr109:
		{st.Month = 4 }
		goto _st30;
		_ctr117:
		{st.Month = 8 }
		goto _st30;
		_ctr123:
		{st.Month = 12 }
		goto _st30;
		_ctr129:
		{st.Month = 2 }
		goto _st30;
		_ctr146:
		{st.Month = 1 }
		goto _st30;
		_ctr159:
		{st.Month = 7 }
		goto _st30;
		_ctr161:
		{st.Month = 6 }
		goto _st30;
		_ctr168:
		{st.Month = 3 }
		goto _st30;
		_ctr174:
		{st.Month = 5 }
		goto _st30;
		_ctr180:
		{st.Month = 11 }
		goto _st30;
		_ctr186:
		{st.Month = 10 }
		goto _st30;
		_ctr192:
		{st.Month = 9 }
		goto _st30;
		_st30:
		if p == eof {
			goto _out30;
			
		}
		p+=1;
		st_case_30:
		if p == pe && p != eof {
			goto _out30;
			
		}
		if p == eof {
			goto _ctr94;
			
		} else {
			switch ( data[ p ] ) {
				case 48:
				{
					goto _ctr95;
					
				}
				case 51:
				{
					goto _ctr97;
					
				}
				
			}
			if ( data[ p ] ) > 50 {
				if 52 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr98;
					
				}
				
			} else if ( data[ p ] ) >= 49 {
				goto _ctr96;
				
			}
			goto _ctr2;
			
		}
		_ctr99:
		{return st, err }
		goto _st31;
		_ctr95:
		{pb = p }
		goto _st31;
		_st31:
		if p == eof {
			goto _out31;
			
		}
		p+=1;
		st_case_31:
		if p == pe && p != eof {
			goto _out31;
			
		}
		if p == eof {
			goto _ctr99;
			
		} else {
			if 49 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st456;
				
			}
			goto _ctr2;
			
		}
		_ctr98:
		{pb = p }
		goto _st456;
		_ctr1119:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st456;
		_st456:
		if p == eof {
			goto _out456;
			
		}
		p+=1;
		st_case_456:
		if p == pe && p != eof {
			goto _out456;
			
		}
		if p == eof {
			goto _ctr1119;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1120;
					
				}
				case 43:
				{
					goto _ctr1121;
					
				}
				case 45:
				{
					goto _ctr1122;
					
				}
				case 47:
				{
					goto _ctr1123;
					
				}
				case 84:
				{
					goto _ctr1124;
					
				}
				case 90:
				{
					goto _ctr1125;
					
				}
				case 95:
				{
					goto _ctr1126;
					
				}
				case 116:
				{
					goto _ctr1126;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1123;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr1123;
				
			}
			goto _st0;
			
		}
		_ctr96:
		{pb = p }
		goto _st457;
		_ctr1127:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st457;
		_st457:
		if p == eof {
			goto _out457;
			
		}
		p+=1;
		st_case_457:
		if p == pe && p != eof {
			goto _out457;
			
		}
		if p == eof {
			goto _ctr1127;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1120;
					
				}
				case 43:
				{
					goto _ctr1121;
					
				}
				case 45:
				{
					goto _ctr1122;
					
				}
				case 47:
				{
					goto _ctr1123;
					
				}
				case 84:
				{
					goto _ctr1124;
					
				}
				case 90:
				{
					goto _ctr1125;
					
				}
				case 95:
				{
					goto _ctr1126;
					
				}
				case 116:
				{
					goto _ctr1126;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st456;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1123;
					
				}
				
			} else {
				goto _ctr1123;
				
			}
			goto _st0;
			
		}
		_ctr97:
		{pb = p }
		goto _st458;
		_ctr1128:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st458;
		_st458:
		if p == eof {
			goto _out458;
			
		}
		p+=1;
		st_case_458:
		if p == pe && p != eof {
			goto _out458;
			
		}
		if p == eof {
			goto _ctr1128;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1120;
					
				}
				case 43:
				{
					goto _ctr1121;
					
				}
				case 45:
				{
					goto _ctr1122;
					
				}
				case 47:
				{
					goto _ctr1123;
					
				}
				case 84:
				{
					goto _ctr1124;
					
				}
				case 90:
				{
					goto _ctr1125;
					
				}
				case 95:
				{
					goto _ctr1126;
					
				}
				case 116:
				{
					goto _ctr1126;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 49 {
					goto _st456;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1123;
					
				}
				
			} else {
				goto _ctr1123;
				
			}
			goto _st0;
			
		}
		_ctr101:
		{return st, err }
		goto _st32;
		_ctr30:
		{pb = p }
		goto _st32;
		_st32:
		if p == eof {
			goto _out32;
			
		}
		p+=1;
		st_case_32:
		if p == pe && p != eof {
			goto _out32;
			
		}
		if p == eof {
			goto _ctr101;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr93;
					
				}
				case 47:
				{
					goto _ctr93;
					
				}
				
			}
			if ( data[ p ] ) > 50 {
				if ( data[ p ] ) <= 57 {
					goto _st8;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _st29;
				
			}
			goto _ctr2;
			
		}
		_ctr102:
		{return st, err }
		goto _st33;
		_ctr31:
		{pb = p }
		goto _st33;
		_st33:
		if p == eof {
			goto _out33;
			
		}
		p+=1;
		st_case_33:
		if p == pe && p != eof {
			goto _out33;
			
		}
		if p == eof {
			goto _ctr102;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr93;
					
				}
				case 47:
				{
					goto _ctr93;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st8;
				
			}
			goto _ctr2;
			
		}
		_ctr103:
		{return st, err }
		goto _st34;
		_st34:
		if p == eof {
			goto _out34;
			
		}
		p+=1;
		st_case_34:
		if p == pe && p != eof {
			goto _out34;
			
		}
		if p == eof {
			goto _ctr103;
			
		} else {
			switch ( data[ p ] ) {
				case 112:
				{
					goto _st35;
					
				}
				case 117:
				{
					goto _st39;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr106:
		{return st, err }
		goto _st35;
		_st35:
		if p == eof {
			goto _out35;
			
		}
		p+=1;
		st_case_35:
		if p == pe && p != eof {
			goto _out35;
			
		}
		if p == eof {
			goto _ctr106;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st36;
				
			}
			goto _ctr2;
			
		}
		_ctr108:
		{return st, err }
		goto _st36;
		_st36:
		if p == eof {
			goto _out36;
			
		}
		p+=1;
		st_case_36:
		if p == pe && p != eof {
			goto _out36;
			
		}
		if p == eof {
			goto _ctr108;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr109;
					
				}
				case 47:
				{
					goto _ctr109;
					
				}
				case 105:
				{
					goto _st37;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr111:
		{return st, err }
		goto _st37;
		_st37:
		if p == eof {
			goto _out37;
			
		}
		p+=1;
		st_case_37:
		if p == pe && p != eof {
			goto _out37;
			
		}
		if p == eof {
			goto _ctr111;
			
		} else {
			if ( data[ p ] ) == 108 {
				goto _st38;
				
			}
			goto _ctr2;
			
		}
		_ctr113:
		{return st, err }
		goto _st38;
		_st38:
		if p == eof {
			goto _out38;
			
		}
		p+=1;
		st_case_38:
		if p == pe && p != eof {
			goto _out38;
			
		}
		if p == eof {
			goto _ctr113;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr109;
					
				}
				case 47:
				{
					goto _ctr109;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr114:
		{return st, err }
		goto _st39;
		_st39:
		if p == eof {
			goto _out39;
			
		}
		p+=1;
		st_case_39:
		if p == pe && p != eof {
			goto _out39;
			
		}
		if p == eof {
			goto _ctr114;
			
		} else {
			if ( data[ p ] ) == 103 {
				goto _st40;
				
			}
			goto _ctr2;
			
		}
		_ctr116:
		{return st, err }
		goto _st40;
		_st40:
		if p == eof {
			goto _out40;
			
		}
		p+=1;
		st_case_40:
		if p == pe && p != eof {
			goto _out40;
			
		}
		if p == eof {
			goto _ctr116;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr117;
					
				}
				case 47:
				{
					goto _ctr117;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr118:
		{return st, err }
		goto _st41;
		_st41:
		if p == eof {
			goto _out41;
			
		}
		p+=1;
		st_case_41:
		if p == pe && p != eof {
			goto _out41;
			
		}
		if p == eof {
			goto _ctr118;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st42;
				
			}
			goto _ctr2;
			
		}
		_ctr120:
		{return st, err }
		goto _st42;
		_st42:
		if p == eof {
			goto _out42;
			
		}
		p+=1;
		st_case_42:
		if p == pe && p != eof {
			goto _out42;
			
		}
		if p == eof {
			goto _ctr120;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st43;
				
			}
			goto _ctr2;
			
		}
		_ctr122:
		{return st, err }
		goto _st43;
		_st43:
		if p == eof {
			goto _out43;
			
		}
		p+=1;
		st_case_43:
		if p == pe && p != eof {
			goto _out43;
			
		}
		if p == eof {
			goto _ctr122;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr123;
					
				}
				case 47:
				{
					goto _ctr123;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr124:
		{return st, err }
		goto _st44;
		_st44:
		if p == eof {
			goto _out44;
			
		}
		p+=1;
		st_case_44:
		if p == pe && p != eof {
			goto _out44;
			
		}
		if p == eof {
			goto _ctr124;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st45;
				
			}
			goto _ctr2;
			
		}
		_ctr126:
		{return st, err }
		goto _st45;
		_st45:
		if p == eof {
			goto _out45;
			
		}
		p+=1;
		st_case_45:
		if p == pe && p != eof {
			goto _out45;
			
		}
		if p == eof {
			goto _ctr126;
			
		} else {
			if ( data[ p ] ) == 98 {
				goto _st46;
				
			}
			goto _ctr2;
			
		}
		_ctr128:
		{return st, err }
		goto _st46;
		_st46:
		if p == eof {
			goto _out46;
			
		}
		p+=1;
		st_case_46:
		if p == pe && p != eof {
			goto _out46;
			
		}
		if p == eof {
			goto _ctr128;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr129;
					
				}
				case 47:
				{
					goto _ctr129;
					
				}
				case 114:
				{
					goto _st47;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr131:
		{return st, err }
		goto _st47;
		_st47:
		if p == eof {
			goto _out47;
			
		}
		p+=1;
		st_case_47:
		if p == pe && p != eof {
			goto _out47;
			
		}
		if p == eof {
			goto _ctr131;
			
		} else {
			if ( data[ p ] ) == 117 {
				goto _st48;
				
			}
			goto _ctr2;
			
		}
		_ctr133:
		{return st, err }
		goto _st48;
		_st48:
		if p == eof {
			goto _out48;
			
		}
		p+=1;
		st_case_48:
		if p == pe && p != eof {
			goto _out48;
			
		}
		if p == eof {
			goto _ctr133;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st49;
				
			}
			goto _ctr2;
			
		}
		_ctr135:
		{return st, err }
		goto _st49;
		_st49:
		if p == eof {
			goto _out49;
			
		}
		p+=1;
		st_case_49:
		if p == pe && p != eof {
			goto _out49;
			
		}
		if p == eof {
			goto _ctr135;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st50;
				
			}
			goto _ctr2;
			
		}
		_ctr137:
		{return st, err }
		goto _st50;
		_st50:
		if p == eof {
			goto _out50;
			
		}
		p+=1;
		st_case_50:
		if p == pe && p != eof {
			goto _out50;
			
		}
		if p == eof {
			goto _ctr137;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st51;
				
			}
			goto _ctr2;
			
		}
		_ctr139:
		{return st, err }
		goto _st51;
		_st51:
		if p == eof {
			goto _out51;
			
		}
		p+=1;
		st_case_51:
		if p == pe && p != eof {
			goto _out51;
			
		}
		if p == eof {
			goto _ctr139;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr129;
					
				}
				case 47:
				{
					goto _ctr129;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr140:
		{return st, err }
		goto _st52;
		_st52:
		if p == eof {
			goto _out52;
			
		}
		p+=1;
		st_case_52:
		if p == pe && p != eof {
			goto _out52;
			
		}
		if p == eof {
			goto _ctr140;
			
		} else {
			switch ( data[ p ] ) {
				case 97:
				{
					goto _st53;
					
				}
				case 117:
				{
					goto _st59;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr143:
		{return st, err }
		goto _st53;
		_st53:
		if p == eof {
			goto _out53;
			
		}
		p+=1;
		st_case_53:
		if p == pe && p != eof {
			goto _out53;
			
		}
		if p == eof {
			goto _ctr143;
			
		} else {
			if ( data[ p ] ) == 110 {
				goto _st54;
				
			}
			goto _ctr2;
			
		}
		_ctr145:
		{return st, err }
		goto _st54;
		_st54:
		if p == eof {
			goto _out54;
			
		}
		p+=1;
		st_case_54:
		if p == pe && p != eof {
			goto _out54;
			
		}
		if p == eof {
			goto _ctr145;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr146;
					
				}
				case 47:
				{
					goto _ctr146;
					
				}
				case 117:
				{
					goto _st55;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr148:
		{return st, err }
		goto _st55;
		_st55:
		if p == eof {
			goto _out55;
			
		}
		p+=1;
		st_case_55:
		if p == pe && p != eof {
			goto _out55;
			
		}
		if p == eof {
			goto _ctr148;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st56;
				
			}
			goto _ctr2;
			
		}
		_ctr150:
		{return st, err }
		goto _st56;
		_st56:
		if p == eof {
			goto _out56;
			
		}
		p+=1;
		st_case_56:
		if p == pe && p != eof {
			goto _out56;
			
		}
		if p == eof {
			goto _ctr150;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st57;
				
			}
			goto _ctr2;
			
		}
		_ctr152:
		{return st, err }
		goto _st57;
		_st57:
		if p == eof {
			goto _out57;
			
		}
		p+=1;
		st_case_57:
		if p == pe && p != eof {
			goto _out57;
			
		}
		if p == eof {
			goto _ctr152;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st58;
				
			}
			goto _ctr2;
			
		}
		_ctr154:
		{return st, err }
		goto _st58;
		_st58:
		if p == eof {
			goto _out58;
			
		}
		p+=1;
		st_case_58:
		if p == pe && p != eof {
			goto _out58;
			
		}
		if p == eof {
			goto _ctr154;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr146;
					
				}
				case 47:
				{
					goto _ctr146;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr155:
		{return st, err }
		goto _st59;
		_st59:
		if p == eof {
			goto _out59;
			
		}
		p+=1;
		st_case_59:
		if p == pe && p != eof {
			goto _out59;
			
		}
		if p == eof {
			goto _ctr155;
			
		} else {
			switch ( data[ p ] ) {
				case 108:
				{
					goto _st60;
					
				}
				case 110:
				{
					goto _st61;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr158:
		{return st, err }
		goto _st60;
		_st60:
		if p == eof {
			goto _out60;
			
		}
		p+=1;
		st_case_60:
		if p == pe && p != eof {
			goto _out60;
			
		}
		if p == eof {
			goto _ctr158;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr159;
					
				}
				case 47:
				{
					goto _ctr159;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr160:
		{return st, err }
		goto _st61;
		_st61:
		if p == eof {
			goto _out61;
			
		}
		p+=1;
		st_case_61:
		if p == pe && p != eof {
			goto _out61;
			
		}
		if p == eof {
			goto _ctr160;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr161;
					
				}
				case 47:
				{
					goto _ctr161;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr162:
		{return st, err }
		goto _st62;
		_st62:
		if p == eof {
			goto _out62;
			
		}
		p+=1;
		st_case_62:
		if p == pe && p != eof {
			goto _out62;
			
		}
		if p == eof {
			goto _ctr162;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st63;
				
			}
			goto _ctr2;
			
		}
		_ctr164:
		{return st, err }
		goto _st63;
		_st63:
		if p == eof {
			goto _out63;
			
		}
		p+=1;
		st_case_63:
		if p == pe && p != eof {
			goto _out63;
			
		}
		if p == eof {
			goto _ctr164;
			
		} else {
			switch ( data[ p ] ) {
				case 114:
				{
					goto _st64;
					
				}
				case 121:
				{
					goto _st67;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr167:
		{return st, err }
		goto _st64;
		_st64:
		if p == eof {
			goto _out64;
			
		}
		p+=1;
		st_case_64:
		if p == pe && p != eof {
			goto _out64;
			
		}
		if p == eof {
			goto _ctr167;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr168;
					
				}
				case 47:
				{
					goto _ctr168;
					
				}
				case 99:
				{
					goto _st65;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr170:
		{return st, err }
		goto _st65;
		_st65:
		if p == eof {
			goto _out65;
			
		}
		p+=1;
		st_case_65:
		if p == pe && p != eof {
			goto _out65;
			
		}
		if p == eof {
			goto _ctr170;
			
		} else {
			if ( data[ p ] ) == 104 {
				goto _st66;
				
			}
			goto _ctr2;
			
		}
		_ctr172:
		{return st, err }
		goto _st66;
		_st66:
		if p == eof {
			goto _out66;
			
		}
		p+=1;
		st_case_66:
		if p == pe && p != eof {
			goto _out66;
			
		}
		if p == eof {
			goto _ctr172;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr168;
					
				}
				case 47:
				{
					goto _ctr168;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr173:
		{return st, err }
		goto _st67;
		_st67:
		if p == eof {
			goto _out67;
			
		}
		p+=1;
		st_case_67:
		if p == pe && p != eof {
			goto _out67;
			
		}
		if p == eof {
			goto _ctr173;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr174;
					
				}
				case 47:
				{
					goto _ctr174;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr175:
		{return st, err }
		goto _st68;
		_st68:
		if p == eof {
			goto _out68;
			
		}
		p+=1;
		st_case_68:
		if p == pe && p != eof {
			goto _out68;
			
		}
		if p == eof {
			goto _ctr175;
			
		} else {
			if ( data[ p ] ) == 111 {
				goto _st69;
				
			}
			goto _ctr2;
			
		}
		_ctr177:
		{return st, err }
		goto _st69;
		_st69:
		if p == eof {
			goto _out69;
			
		}
		p+=1;
		st_case_69:
		if p == pe && p != eof {
			goto _out69;
			
		}
		if p == eof {
			goto _ctr177;
			
		} else {
			if ( data[ p ] ) == 118 {
				goto _st70;
				
			}
			goto _ctr2;
			
		}
		_ctr179:
		{return st, err }
		goto _st70;
		_st70:
		if p == eof {
			goto _out70;
			
		}
		p+=1;
		st_case_70:
		if p == pe && p != eof {
			goto _out70;
			
		}
		if p == eof {
			goto _ctr179;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr180;
					
				}
				case 47:
				{
					goto _ctr180;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr181:
		{return st, err }
		goto _st71;
		_st71:
		if p == eof {
			goto _out71;
			
		}
		p+=1;
		st_case_71:
		if p == pe && p != eof {
			goto _out71;
			
		}
		if p == eof {
			goto _ctr181;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st72;
				
			}
			goto _ctr2;
			
		}
		_ctr183:
		{return st, err }
		goto _st72;
		_st72:
		if p == eof {
			goto _out72;
			
		}
		p+=1;
		st_case_72:
		if p == pe && p != eof {
			goto _out72;
			
		}
		if p == eof {
			goto _ctr183;
			
		} else {
			if ( data[ p ] ) == 116 {
				goto _st73;
				
			}
			goto _ctr2;
			
		}
		_ctr185:
		{return st, err }
		goto _st73;
		_st73:
		if p == eof {
			goto _out73;
			
		}
		p+=1;
		st_case_73:
		if p == pe && p != eof {
			goto _out73;
			
		}
		if p == eof {
			goto _ctr185;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr186;
					
				}
				case 47:
				{
					goto _ctr186;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr187:
		{return st, err }
		goto _st74;
		_st74:
		if p == eof {
			goto _out74;
			
		}
		p+=1;
		st_case_74:
		if p == pe && p != eof {
			goto _out74;
			
		}
		if p == eof {
			goto _ctr187;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st75;
				
			}
			goto _ctr2;
			
		}
		_ctr189:
		{return st, err }
		goto _st75;
		_st75:
		if p == eof {
			goto _out75;
			
		}
		p+=1;
		st_case_75:
		if p == pe && p != eof {
			goto _out75;
			
		}
		if p == eof {
			goto _ctr189;
			
		} else {
			if ( data[ p ] ) == 112 {
				goto _st76;
				
			}
			goto _ctr2;
			
		}
		_ctr191:
		{return st, err }
		goto _st76;
		_st76:
		if p == eof {
			goto _out76;
			
		}
		p+=1;
		st_case_76:
		if p == pe && p != eof {
			goto _out76;
			
		}
		if p == eof {
			goto _ctr191;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr192;
					
				}
				case 47:
				{
					goto _ctr192;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr193:
		{return st, err }
		goto _st77;
		_ctr26:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		goto _st77;
		_st77:
		if p == eof {
			goto _out77;
			
		}
		p+=1;
		st_case_77:
		if p == pe && p != eof {
			goto _out77;
			
		}
		if p == eof {
			goto _ctr193;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _ctr194;
				
			}
			goto _ctr2;
			
		}
		_ctr195:
		{return st, err }
		goto _st78;
		_ctr194:
		{pb = p }
		goto _st78;
		_st78:
		if p == eof {
			goto _out78;
			
		}
		p+=1;
		st_case_78:
		if p == pe && p != eof {
			goto _out78;
			
		}
		if p == eof {
			goto _ctr195;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st8;
				
			}
			goto _ctr2;
			
		}
		_ctr196:
		{return st, err }
		goto _st79;
		_ctr27:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		{pb = p }
		goto _st79;
		_st79:
		if p == eof {
			goto _out79;
			
		}
		p+=1;
		st_case_79:
		if p == pe && p != eof {
			goto _out79;
			
		}
		if p == eof {
			goto _ctr196;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st80;
				
			}
			goto _ctr2;
			
		}
		_ctr198:
		{return st, err }
		goto _st80;
		_st80:
		if p == eof {
			goto _out80;
			
		}
		p+=1;
		st_case_80:
		if p == pe && p != eof {
			goto _out80;
			
		}
		if p == eof {
			goto _ctr198;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st459;
				
			}
			goto _ctr2;
			
		}
		_ctr1129:
		{st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
		}
		goto _st459;
		_st459:
		if p == eof {
			goto _out459;
			
		}
		p+=1;
		st_case_459:
		if p == pe && p != eof {
			goto _out459;
			
		}
		if p == eof {
			goto _ctr1129;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr956;
					
				}
				case 43:
				{
					goto _ctr957;
					
				}
				case 45:
				{
					goto _ctr958;
					
				}
				case 47:
				{
					goto _ctr959;
					
				}
				case 84:
				{
					goto _ctr960;
					
				}
				case 90:
				{
					goto _ctr961;
					
				}
				case 95:
				{
					goto _ctr962;
					
				}
				case 116:
				{
					goto _ctr962;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st460;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr959;
					
				}
				
			} else {
				goto _ctr959;
				
			}
			goto _st0;
			
		}
		_ctr1131:
		{st.Month, _ = strconv.Atoi(data[pb:pb+2])
			st.Day, _ = strconv.Atoi(data[pb+2:pb+4])
		}
		goto _st460;
		_st460:
		if p == eof {
			goto _out460;
			
		}
		p+=1;
		st_case_460:
		if p == pe && p != eof {
			goto _out460;
			
		}
		if p == eof {
			goto _ctr1131;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1132;
					
				}
				case 43:
				{
					goto _ctr1133;
					
				}
				case 45:
				{
					goto _ctr1134;
					
				}
				case 47:
				{
					goto _ctr1135;
					
				}
				case 84:
				{
					goto _ctr1136;
					
				}
				case 90:
				{
					goto _ctr1137;
					
				}
				case 95:
				{
					goto _ctr1138;
					
				}
				case 116:
				{
					goto _ctr1138;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1135;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr1135;
				
			}
			goto _st0;
			
		}
		_ctr200:
		{return st, err }
		goto _st81;
		_st81:
		if p == eof {
			goto _out81;
			
		}
		p+=1;
		st_case_81:
		if p == pe && p != eof {
			goto _out81;
			
		}
		if p == eof {
			goto _ctr200;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr201;
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st4;
				
			}
			goto _ctr2;
			
		}
		_ctr202:
		{return st, err }
		goto _st82;
		_ctr201:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st82;
		_st82:
		if p == eof {
			goto _out82;
			
		}
		p+=1;
		st_case_82:
		if p == pe && p != eof {
			goto _out82;
			
		}
		if p == eof {
			goto _ctr202;
			
		} else {
			switch ( data[ p ] ) {
				case 65:
				{
					goto _st83;
					
				}
				case 68:
				{
					goto _st92;
					
				}
				case 70:
				{
					goto _st95;
					
				}
				case 74:
				{
					goto _st103;
					
				}
				case 77:
				{
					goto _st113;
					
				}
				case 78:
				{
					goto _st119;
					
				}
				case 79:
				{
					goto _st122;
					
				}
				case 83:
				{
					goto _st125;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr211:
		{return st, err }
		goto _st83;
		_st83:
		if p == eof {
			goto _out83;
			
		}
		p+=1;
		st_case_83:
		if p == pe && p != eof {
			goto _out83;
			
		}
		if p == eof {
			goto _ctr211;
			
		} else {
			switch ( data[ p ] ) {
				case 112:
				{
					goto _st84;
					
				}
				case 117:
				{
					goto _st90;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr214:
		{return st, err }
		goto _st84;
		_st84:
		if p == eof {
			goto _out84;
			
		}
		p+=1;
		st_case_84:
		if p == pe && p != eof {
			goto _out84;
			
		}
		if p == eof {
			goto _ctr214;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st85;
				
			}
			goto _ctr2;
			
		}
		_ctr216:
		{return st, err }
		goto _st85;
		_st85:
		if p == eof {
			goto _out85;
			
		}
		p+=1;
		st_case_85:
		if p == pe && p != eof {
			goto _out85;
			
		}
		if p == eof {
			goto _ctr216;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr217;
					
				}
				case 105:
				{
					goto _st88;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr219:
		{return st, err }
		goto _st86;
		_ctr217:
		{st.Month = 4 }
		goto _st86;
		_ctr229:
		{st.Month = 8 }
		goto _st86;
		_ctr235:
		{st.Month = 12 }
		goto _st86;
		_ctr241:
		{st.Month = 2 }
		goto _st86;
		_ctr258:
		{st.Month = 1 }
		goto _st86;
		_ctr271:
		{st.Month = 7 }
		goto _st86;
		_ctr273:
		{st.Month = 6 }
		goto _st86;
		_ctr280:
		{st.Month = 3 }
		goto _st86;
		_ctr286:
		{st.Month = 5 }
		goto _st86;
		_ctr292:
		{st.Month = 11 }
		goto _st86;
		_ctr298:
		{st.Month = 10 }
		goto _st86;
		_ctr304:
		{st.Month = 9 }
		goto _st86;
		_st86:
		if p == eof {
			goto _out86;
			
		}
		p+=1;
		st_case_86:
		if p == pe && p != eof {
			goto _out86;
			
		}
		if p == eof {
			goto _ctr219;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _ctr220;
				
			}
			goto _ctr2;
			
		}
		_ctr221:
		{return st, err }
		goto _st87;
		_ctr220:
		{pb = p }
		goto _st87;
		_st87:
		if p == eof {
			goto _out87;
			
		}
		p+=1;
		st_case_87:
		if p == pe && p != eof {
			goto _out87;
			
		}
		if p == eof {
			goto _ctr221;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st461;
				
			}
			goto _ctr2;
			
		}
		_ctr1139:
		{st.Year = parse_year_2_digits(data[pb:pb+2])
		}
		goto _st461;
		_st461:
		if p == eof {
			goto _out461;
			
		}
		p+=1;
		st_case_461:
		if p == pe && p != eof {
			goto _out461;
			
		}
		if p == eof {
			goto _ctr1139;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1140;
					
				}
				case 43:
				{
					goto _ctr1141;
					
				}
				case 45:
				{
					goto _ctr1142;
					
				}
				case 47:
				{
					goto _ctr1143;
					
				}
				case 84:
				{
					goto _ctr1144;
					
				}
				case 90:
				{
					goto _ctr1145;
					
				}
				case 95:
				{
					goto _ctr1146;
					
				}
				case 116:
				{
					goto _ctr1146;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1143;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr1143;
				
			}
			goto _st0;
			
		}
		_ctr223:
		{return st, err }
		goto _st88;
		_st88:
		if p == eof {
			goto _out88;
			
		}
		p+=1;
		st_case_88:
		if p == pe && p != eof {
			goto _out88;
			
		}
		if p == eof {
			goto _ctr223;
			
		} else {
			if ( data[ p ] ) == 108 {
				goto _st89;
				
			}
			goto _ctr2;
			
		}
		_ctr225:
		{return st, err }
		goto _st89;
		_st89:
		if p == eof {
			goto _out89;
			
		}
		p+=1;
		st_case_89:
		if p == pe && p != eof {
			goto _out89;
			
		}
		if p == eof {
			goto _ctr225;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr217;
				
			}
			goto _ctr2;
			
		}
		_ctr226:
		{return st, err }
		goto _st90;
		_st90:
		if p == eof {
			goto _out90;
			
		}
		p+=1;
		st_case_90:
		if p == pe && p != eof {
			goto _out90;
			
		}
		if p == eof {
			goto _ctr226;
			
		} else {
			if ( data[ p ] ) == 103 {
				goto _st91;
				
			}
			goto _ctr2;
			
		}
		_ctr228:
		{return st, err }
		goto _st91;
		_st91:
		if p == eof {
			goto _out91;
			
		}
		p+=1;
		st_case_91:
		if p == pe && p != eof {
			goto _out91;
			
		}
		if p == eof {
			goto _ctr228;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr229;
				
			}
			goto _ctr2;
			
		}
		_ctr230:
		{return st, err }
		goto _st92;
		_st92:
		if p == eof {
			goto _out92;
			
		}
		p+=1;
		st_case_92:
		if p == pe && p != eof {
			goto _out92;
			
		}
		if p == eof {
			goto _ctr230;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st93;
				
			}
			goto _ctr2;
			
		}
		_ctr232:
		{return st, err }
		goto _st93;
		_st93:
		if p == eof {
			goto _out93;
			
		}
		p+=1;
		st_case_93:
		if p == pe && p != eof {
			goto _out93;
			
		}
		if p == eof {
			goto _ctr232;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st94;
				
			}
			goto _ctr2;
			
		}
		_ctr234:
		{return st, err }
		goto _st94;
		_st94:
		if p == eof {
			goto _out94;
			
		}
		p+=1;
		st_case_94:
		if p == pe && p != eof {
			goto _out94;
			
		}
		if p == eof {
			goto _ctr234;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr235;
				
			}
			goto _ctr2;
			
		}
		_ctr236:
		{return st, err }
		goto _st95;
		_st95:
		if p == eof {
			goto _out95;
			
		}
		p+=1;
		st_case_95:
		if p == pe && p != eof {
			goto _out95;
			
		}
		if p == eof {
			goto _ctr236;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st96;
				
			}
			goto _ctr2;
			
		}
		_ctr238:
		{return st, err }
		goto _st96;
		_st96:
		if p == eof {
			goto _out96;
			
		}
		p+=1;
		st_case_96:
		if p == pe && p != eof {
			goto _out96;
			
		}
		if p == eof {
			goto _ctr238;
			
		} else {
			if ( data[ p ] ) == 98 {
				goto _st97;
				
			}
			goto _ctr2;
			
		}
		_ctr240:
		{return st, err }
		goto _st97;
		_st97:
		if p == eof {
			goto _out97;
			
		}
		p+=1;
		st_case_97:
		if p == pe && p != eof {
			goto _out97;
			
		}
		if p == eof {
			goto _ctr240;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr241;
					
				}
				case 114:
				{
					goto _st98;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr243:
		{return st, err }
		goto _st98;
		_st98:
		if p == eof {
			goto _out98;
			
		}
		p+=1;
		st_case_98:
		if p == pe && p != eof {
			goto _out98;
			
		}
		if p == eof {
			goto _ctr243;
			
		} else {
			if ( data[ p ] ) == 117 {
				goto _st99;
				
			}
			goto _ctr2;
			
		}
		_ctr245:
		{return st, err }
		goto _st99;
		_st99:
		if p == eof {
			goto _out99;
			
		}
		p+=1;
		st_case_99:
		if p == pe && p != eof {
			goto _out99;
			
		}
		if p == eof {
			goto _ctr245;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st100;
				
			}
			goto _ctr2;
			
		}
		_ctr247:
		{return st, err }
		goto _st100;
		_st100:
		if p == eof {
			goto _out100;
			
		}
		p+=1;
		st_case_100:
		if p == pe && p != eof {
			goto _out100;
			
		}
		if p == eof {
			goto _ctr247;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st101;
				
			}
			goto _ctr2;
			
		}
		_ctr249:
		{return st, err }
		goto _st101;
		_st101:
		if p == eof {
			goto _out101;
			
		}
		p+=1;
		st_case_101:
		if p == pe && p != eof {
			goto _out101;
			
		}
		if p == eof {
			goto _ctr249;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st102;
				
			}
			goto _ctr2;
			
		}
		_ctr251:
		{return st, err }
		goto _st102;
		_st102:
		if p == eof {
			goto _out102;
			
		}
		p+=1;
		st_case_102:
		if p == pe && p != eof {
			goto _out102;
			
		}
		if p == eof {
			goto _ctr251;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr241;
				
			}
			goto _ctr2;
			
		}
		_ctr252:
		{return st, err }
		goto _st103;
		_st103:
		if p == eof {
			goto _out103;
			
		}
		p+=1;
		st_case_103:
		if p == pe && p != eof {
			goto _out103;
			
		}
		if p == eof {
			goto _ctr252;
			
		} else {
			switch ( data[ p ] ) {
				case 97:
				{
					goto _st104;
					
				}
				case 117:
				{
					goto _st110;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr255:
		{return st, err }
		goto _st104;
		_st104:
		if p == eof {
			goto _out104;
			
		}
		p+=1;
		st_case_104:
		if p == pe && p != eof {
			goto _out104;
			
		}
		if p == eof {
			goto _ctr255;
			
		} else {
			if ( data[ p ] ) == 110 {
				goto _st105;
				
			}
			goto _ctr2;
			
		}
		_ctr257:
		{return st, err }
		goto _st105;
		_st105:
		if p == eof {
			goto _out105;
			
		}
		p+=1;
		st_case_105:
		if p == pe && p != eof {
			goto _out105;
			
		}
		if p == eof {
			goto _ctr257;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr258;
					
				}
				case 117:
				{
					goto _st106;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr260:
		{return st, err }
		goto _st106;
		_st106:
		if p == eof {
			goto _out106;
			
		}
		p+=1;
		st_case_106:
		if p == pe && p != eof {
			goto _out106;
			
		}
		if p == eof {
			goto _ctr260;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st107;
				
			}
			goto _ctr2;
			
		}
		_ctr262:
		{return st, err }
		goto _st107;
		_st107:
		if p == eof {
			goto _out107;
			
		}
		p+=1;
		st_case_107:
		if p == pe && p != eof {
			goto _out107;
			
		}
		if p == eof {
			goto _ctr262;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st108;
				
			}
			goto _ctr2;
			
		}
		_ctr264:
		{return st, err }
		goto _st108;
		_st108:
		if p == eof {
			goto _out108;
			
		}
		p+=1;
		st_case_108:
		if p == pe && p != eof {
			goto _out108;
			
		}
		if p == eof {
			goto _ctr264;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st109;
				
			}
			goto _ctr2;
			
		}
		_ctr266:
		{return st, err }
		goto _st109;
		_st109:
		if p == eof {
			goto _out109;
			
		}
		p+=1;
		st_case_109:
		if p == pe && p != eof {
			goto _out109;
			
		}
		if p == eof {
			goto _ctr266;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr258;
				
			}
			goto _ctr2;
			
		}
		_ctr267:
		{return st, err }
		goto _st110;
		_st110:
		if p == eof {
			goto _out110;
			
		}
		p+=1;
		st_case_110:
		if p == pe && p != eof {
			goto _out110;
			
		}
		if p == eof {
			goto _ctr267;
			
		} else {
			switch ( data[ p ] ) {
				case 108:
				{
					goto _st111;
					
				}
				case 110:
				{
					goto _st112;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr270:
		{return st, err }
		goto _st111;
		_st111:
		if p == eof {
			goto _out111;
			
		}
		p+=1;
		st_case_111:
		if p == pe && p != eof {
			goto _out111;
			
		}
		if p == eof {
			goto _ctr270;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr271;
				
			}
			goto _ctr2;
			
		}
		_ctr272:
		{return st, err }
		goto _st112;
		_st112:
		if p == eof {
			goto _out112;
			
		}
		p+=1;
		st_case_112:
		if p == pe && p != eof {
			goto _out112;
			
		}
		if p == eof {
			goto _ctr272;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr273;
				
			}
			goto _ctr2;
			
		}
		_ctr274:
		{return st, err }
		goto _st113;
		_st113:
		if p == eof {
			goto _out113;
			
		}
		p+=1;
		st_case_113:
		if p == pe && p != eof {
			goto _out113;
			
		}
		if p == eof {
			goto _ctr274;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st114;
				
			}
			goto _ctr2;
			
		}
		_ctr276:
		{return st, err }
		goto _st114;
		_st114:
		if p == eof {
			goto _out114;
			
		}
		p+=1;
		st_case_114:
		if p == pe && p != eof {
			goto _out114;
			
		}
		if p == eof {
			goto _ctr276;
			
		} else {
			switch ( data[ p ] ) {
				case 114:
				{
					goto _st115;
					
				}
				case 121:
				{
					goto _st118;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr279:
		{return st, err }
		goto _st115;
		_st115:
		if p == eof {
			goto _out115;
			
		}
		p+=1;
		st_case_115:
		if p == pe && p != eof {
			goto _out115;
			
		}
		if p == eof {
			goto _ctr279;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr280;
					
				}
				case 99:
				{
					goto _st116;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr282:
		{return st, err }
		goto _st116;
		_st116:
		if p == eof {
			goto _out116;
			
		}
		p+=1;
		st_case_116:
		if p == pe && p != eof {
			goto _out116;
			
		}
		if p == eof {
			goto _ctr282;
			
		} else {
			if ( data[ p ] ) == 104 {
				goto _st117;
				
			}
			goto _ctr2;
			
		}
		_ctr284:
		{return st, err }
		goto _st117;
		_st117:
		if p == eof {
			goto _out117;
			
		}
		p+=1;
		st_case_117:
		if p == pe && p != eof {
			goto _out117;
			
		}
		if p == eof {
			goto _ctr284;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr280;
				
			}
			goto _ctr2;
			
		}
		_ctr285:
		{return st, err }
		goto _st118;
		_st118:
		if p == eof {
			goto _out118;
			
		}
		p+=1;
		st_case_118:
		if p == pe && p != eof {
			goto _out118;
			
		}
		if p == eof {
			goto _ctr285;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr286;
				
			}
			goto _ctr2;
			
		}
		_ctr287:
		{return st, err }
		goto _st119;
		_st119:
		if p == eof {
			goto _out119;
			
		}
		p+=1;
		st_case_119:
		if p == pe && p != eof {
			goto _out119;
			
		}
		if p == eof {
			goto _ctr287;
			
		} else {
			if ( data[ p ] ) == 111 {
				goto _st120;
				
			}
			goto _ctr2;
			
		}
		_ctr289:
		{return st, err }
		goto _st120;
		_st120:
		if p == eof {
			goto _out120;
			
		}
		p+=1;
		st_case_120:
		if p == pe && p != eof {
			goto _out120;
			
		}
		if p == eof {
			goto _ctr289;
			
		} else {
			if ( data[ p ] ) == 118 {
				goto _st121;
				
			}
			goto _ctr2;
			
		}
		_ctr291:
		{return st, err }
		goto _st121;
		_st121:
		if p == eof {
			goto _out121;
			
		}
		p+=1;
		st_case_121:
		if p == pe && p != eof {
			goto _out121;
			
		}
		if p == eof {
			goto _ctr291;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr292;
				
			}
			goto _ctr2;
			
		}
		_ctr293:
		{return st, err }
		goto _st122;
		_st122:
		if p == eof {
			goto _out122;
			
		}
		p+=1;
		st_case_122:
		if p == pe && p != eof {
			goto _out122;
			
		}
		if p == eof {
			goto _ctr293;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st123;
				
			}
			goto _ctr2;
			
		}
		_ctr295:
		{return st, err }
		goto _st123;
		_st123:
		if p == eof {
			goto _out123;
			
		}
		p+=1;
		st_case_123:
		if p == pe && p != eof {
			goto _out123;
			
		}
		if p == eof {
			goto _ctr295;
			
		} else {
			if ( data[ p ] ) == 116 {
				goto _st124;
				
			}
			goto _ctr2;
			
		}
		_ctr297:
		{return st, err }
		goto _st124;
		_st124:
		if p == eof {
			goto _out124;
			
		}
		p+=1;
		st_case_124:
		if p == pe && p != eof {
			goto _out124;
			
		}
		if p == eof {
			goto _ctr297;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr298;
				
			}
			goto _ctr2;
			
		}
		_ctr299:
		{return st, err }
		goto _st125;
		_st125:
		if p == eof {
			goto _out125;
			
		}
		p+=1;
		st_case_125:
		if p == pe && p != eof {
			goto _out125;
			
		}
		if p == eof {
			goto _ctr299;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st126;
				
			}
			goto _ctr2;
			
		}
		_ctr301:
		{return st, err }
		goto _st126;
		_st126:
		if p == eof {
			goto _out126;
			
		}
		p+=1;
		st_case_126:
		if p == pe && p != eof {
			goto _out126;
			
		}
		if p == eof {
			goto _ctr301;
			
		} else {
			if ( data[ p ] ) == 112 {
				goto _st127;
				
			}
			goto _ctr2;
			
		}
		_ctr303:
		{return st, err }
		goto _st127;
		_st127:
		if p == eof {
			goto _out127;
			
		}
		p+=1;
		st_case_127:
		if p == pe && p != eof {
			goto _out127;
			
		}
		if p == eof {
			goto _ctr303;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr304;
				
			}
			goto _ctr2;
			
		}
		_ctr305:
		{return st, err }
		goto _st128;
		_ctr4:
		{pb = p }
		goto _st128;
		_st128:
		if p == eof {
			goto _out128;
			
		}
		p+=1;
		st_case_128:
		if p == pe && p != eof {
			goto _out128;
			
		}
		if p == eof {
			goto _ctr305;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr201;
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st81;
				
			}
			goto _ctr2;
			
		}
		_ctr306:
		{return st, err }
		goto _st129;
		_ctr5:
		{pb = p }
		goto _st129;
		_st129:
		if p == eof {
			goto _out129;
			
		}
		p+=1;
		st_case_129:
		if p == pe && p != eof {
			goto _out129;
			
		}
		if p == eof {
			goto _ctr306;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr201;
				
			}
			if ( data[ p ] ) > 49 {
				if ( data[ p ] ) <= 57 {
					goto _st3;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _st81;
				
			}
			goto _ctr2;
			
		}
		_ctr307:
		{return st, err }
		goto _st130;
		_ctr6:
		{pb = p }
		goto _st130;
		_st130:
		if p == eof {
			goto _out130;
			
		}
		p+=1;
		st_case_130:
		if p == pe && p != eof {
			goto _out130;
			
		}
		if p == eof {
			goto _ctr307;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr201;
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st3;
				
			}
			goto _ctr2;
			
		}
		_ctr308:
		{return st, err }
		goto _st131;
		_st131:
		if p == eof {
			goto _out131;
			
		}
		p+=1;
		st_case_131:
		if p == pe && p != eof {
			goto _out131;
			
		}
		if p == eof {
			goto _ctr308;
			
		} else {
			switch ( data[ p ] ) {
				case 112:
				{
					goto _st132;
					
				}
				case 117:
				{
					goto _st145;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr311:
		{return st, err }
		goto _st132;
		_st132:
		if p == eof {
			goto _out132;
			
		}
		p+=1;
		st_case_132:
		if p == pe && p != eof {
			goto _out132;
			
		}
		if p == eof {
			goto _ctr311;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st133;
				
			}
			goto _ctr2;
			
		}
		_ctr313:
		{return st, err }
		goto _st133;
		_st133:
		if p == eof {
			goto _out133;
			
		}
		p+=1;
		st_case_133:
		if p == pe && p != eof {
			goto _out133;
			
		}
		if p == eof {
			goto _ctr313;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr314;
					
				}
				case 47:
				{
					goto _ctr314;
					
				}
				case 105:
				{
					goto _st143;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr316:
		{return st, err }
		goto _st134;
		_ctr314:
		{st.Month = 4 }
		goto _st134;
		_ctr341:
		{st.Month = 8 }
		goto _st134;
		_ctr347:
		{st.Month = 12 }
		goto _st134;
		_ctr354:
		{st.Month = 2 }
		goto _st134;
		_ctr880:
		{st.Month = 1 }
		goto _st134;
		_ctr893:
		{st.Month = 7 }
		goto _st134;
		_ctr895:
		{st.Month = 6 }
		goto _st134;
		_ctr903:
		{st.Month = 3 }
		goto _st134;
		_ctr909:
		{st.Month = 5 }
		goto _st134;
		_ctr916:
		{st.Month = 11 }
		goto _st134;
		_ctr922:
		{st.Month = 10 }
		goto _st134;
		_ctr936:
		{st.Month = 9 }
		goto _st134;
		_st134:
		if p == eof {
			goto _out134;
			
		}
		p+=1;
		st_case_134:
		if p == pe && p != eof {
			goto _out134;
			
		}
		if p == eof {
			goto _ctr316;
			
		} else {
			switch ( data[ p ] ) {
				case 48:
				{
					goto _ctr317;
					
				}
				case 51:
				{
					goto _ctr319;
					
				}
				
			}
			if ( data[ p ] ) > 50 {
				if 52 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr320;
					
				}
				
			} else if ( data[ p ] ) >= 49 {
				goto _ctr318;
				
			}
			goto _ctr2;
			
		}
		_ctr321:
		{return st, err }
		goto _st135;
		_ctr317:
		{pb = p }
		goto _st135;
		_st135:
		if p == eof {
			goto _out135;
			
		}
		p+=1;
		st_case_135:
		if p == pe && p != eof {
			goto _out135;
			
		}
		if p == eof {
			goto _ctr321;
			
		} else {
			if 49 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st136;
				
			}
			goto _ctr2;
			
		}
		_ctr323:
		{return st, err }
		goto _st136;
		_ctr320:
		{pb = p }
		goto _st136;
		_st136:
		if p == eof {
			goto _out136;
			
		}
		p+=1;
		st_case_136:
		if p == pe && p != eof {
			goto _out136;
			
		}
		if p == eof {
			goto _ctr323;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr324;
					
				}
				case 47:
				{
					goto _ctr324;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr325:
		{return st, err }
		goto _st137;
		_ctr696:
		{st.Month = 4 }
		goto _st137;
		_ctr704:
		{st.Month = 8 }
		goto _st137;
		_ctr710:
		{st.Month = 12 }
		goto _st137;
		_ctr716:
		{st.Month = 2 }
		goto _st137;
		_ctr733:
		{st.Month = 1 }
		goto _st137;
		_ctr746:
		{st.Month = 7 }
		goto _st137;
		_ctr748:
		{st.Month = 6 }
		goto _st137;
		_ctr755:
		{st.Month = 3 }
		goto _st137;
		_ctr761:
		{st.Month = 5 }
		goto _st137;
		_ctr767:
		{st.Month = 11 }
		goto _st137;
		_ctr773:
		{st.Month = 10 }
		goto _st137;
		_ctr779:
		{st.Month = 9 }
		goto _st137;
		_ctr324:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st137;
		_st137:
		if p == eof {
			goto _out137;
			
		}
		p+=1;
		st_case_137:
		if p == pe && p != eof {
			goto _out137;
			
		}
		if p == eof {
			goto _ctr325;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _ctr326;
				
			}
			goto _ctr2;
			
		}
		_ctr327:
		{return st, err }
		goto _st138;
		_ctr326:
		{pb = p }
		goto _st138;
		_st138:
		if p == eof {
			goto _out138;
			
		}
		p+=1;
		st_case_138:
		if p == pe && p != eof {
			goto _out138;
			
		}
		if p == eof {
			goto _ctr327;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st139;
				
			}
			goto _ctr2;
			
		}
		_ctr329:
		{return st, err }
		goto _st139;
		_st139:
		if p == eof {
			goto _out139;
			
		}
		p+=1;
		st_case_139:
		if p == pe && p != eof {
			goto _out139;
			
		}
		if p == eof {
			goto _ctr329;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st140;
				
			}
			goto _ctr2;
			
		}
		_ctr331:
		{return st, err }
		goto _st140;
		_st140:
		if p == eof {
			goto _out140;
			
		}
		p+=1;
		st_case_140:
		if p == pe && p != eof {
			goto _out140;
			
		}
		if p == eof {
			goto _ctr331;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st462;
				
			}
			goto _ctr2;
			
		}
		_ctr1147:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		goto _st462;
		_st462:
		if p == eof {
			goto _out462;
			
		}
		p+=1;
		st_case_462:
		if p == pe && p != eof {
			goto _out462;
			
		}
		if p == eof {
			goto _ctr1147;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr1148;
					
				}
				case 43:
				{
					goto _ctr1149;
					
				}
				case 45:
				{
					goto _ctr1150;
					
				}
				case 47:
				{
					goto _ctr1151;
					
				}
				case 84:
				{
					goto _ctr1152;
					
				}
				case 90:
				{
					goto _ctr1153;
					
				}
				case 95:
				{
					goto _ctr1154;
					
				}
				case 116:
				{
					goto _ctr1154;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr1151;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr1151;
				
			}
			goto _st0;
			
		}
		_ctr333:
		{return st, err }
		goto _st141;
		_ctr318:
		{pb = p }
		goto _st141;
		_st141:
		if p == eof {
			goto _out141;
			
		}
		p+=1;
		st_case_141:
		if p == pe && p != eof {
			goto _out141;
			
		}
		if p == eof {
			goto _ctr333;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr324;
					
				}
				case 47:
				{
					goto _ctr324;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st136;
				
			}
			goto _ctr2;
			
		}
		_ctr334:
		{return st, err }
		goto _st142;
		_ctr319:
		{pb = p }
		goto _st142;
		_st142:
		if p == eof {
			goto _out142;
			
		}
		p+=1;
		st_case_142:
		if p == pe && p != eof {
			goto _out142;
			
		}
		if p == eof {
			goto _ctr334;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr324;
					
				}
				case 47:
				{
					goto _ctr324;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 49 {
				goto _st136;
				
			}
			goto _ctr2;
			
		}
		_ctr335:
		{return st, err }
		goto _st143;
		_st143:
		if p == eof {
			goto _out143;
			
		}
		p+=1;
		st_case_143:
		if p == pe && p != eof {
			goto _out143;
			
		}
		if p == eof {
			goto _ctr335;
			
		} else {
			if ( data[ p ] ) == 108 {
				goto _st144;
				
			}
			goto _ctr2;
			
		}
		_ctr337:
		{return st, err }
		goto _st144;
		_st144:
		if p == eof {
			goto _out144;
			
		}
		p+=1;
		st_case_144:
		if p == pe && p != eof {
			goto _out144;
			
		}
		if p == eof {
			goto _ctr337;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr314;
					
				}
				case 47:
				{
					goto _ctr314;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr338:
		{return st, err }
		goto _st145;
		_st145:
		if p == eof {
			goto _out145;
			
		}
		p+=1;
		st_case_145:
		if p == pe && p != eof {
			goto _out145;
			
		}
		if p == eof {
			goto _ctr338;
			
		} else {
			if ( data[ p ] ) == 103 {
				goto _st146;
				
			}
			goto _ctr2;
			
		}
		_ctr340:
		{return st, err }
		goto _st146;
		_st146:
		if p == eof {
			goto _out146;
			
		}
		p+=1;
		st_case_146:
		if p == pe && p != eof {
			goto _out146;
			
		}
		if p == eof {
			goto _ctr340;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr341;
					
				}
				case 47:
				{
					goto _ctr341;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr342:
		{return st, err }
		goto _st147;
		_st147:
		if p == eof {
			goto _out147;
			
		}
		p+=1;
		st_case_147:
		if p == pe && p != eof {
			goto _out147;
			
		}
		if p == eof {
			goto _ctr342;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st148;
				
			}
			goto _ctr2;
			
		}
		_ctr344:
		{return st, err }
		goto _st148;
		_st148:
		if p == eof {
			goto _out148;
			
		}
		p+=1;
		st_case_148:
		if p == pe && p != eof {
			goto _out148;
			
		}
		if p == eof {
			goto _ctr344;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st149;
				
			}
			goto _ctr2;
			
		}
		_ctr346:
		{return st, err }
		goto _st149;
		_st149:
		if p == eof {
			goto _out149;
			
		}
		p+=1;
		st_case_149:
		if p == pe && p != eof {
			goto _out149;
			
		}
		if p == eof {
			goto _ctr346;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr347;
					
				}
				case 47:
				{
					goto _ctr347;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr348:
		{return st, err }
		goto _st150;
		_st150:
		if p == eof {
			goto _out150;
			
		}
		p+=1;
		st_case_150:
		if p == pe && p != eof {
			goto _out150;
			
		}
		if p == eof {
			goto _ctr348;
			
		} else {
			switch ( data[ p ] ) {
				case 101:
				{
					goto _st151;
					
				}
				case 114:
				{
					goto _st158;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr351:
		{return st, err }
		goto _st151;
		_st151:
		if p == eof {
			goto _out151;
			
		}
		p+=1;
		st_case_151:
		if p == pe && p != eof {
			goto _out151;
			
		}
		if p == eof {
			goto _ctr351;
			
		} else {
			if ( data[ p ] ) == 98 {
				goto _st152;
				
			}
			goto _ctr2;
			
		}
		_ctr353:
		{return st, err }
		goto _st152;
		_st152:
		if p == eof {
			goto _out152;
			
		}
		p+=1;
		st_case_152:
		if p == pe && p != eof {
			goto _out152;
			
		}
		if p == eof {
			goto _ctr353;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr354;
					
				}
				case 47:
				{
					goto _ctr354;
					
				}
				case 114:
				{
					goto _st153;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr356:
		{return st, err }
		goto _st153;
		_st153:
		if p == eof {
			goto _out153;
			
		}
		p+=1;
		st_case_153:
		if p == pe && p != eof {
			goto _out153;
			
		}
		if p == eof {
			goto _ctr356;
			
		} else {
			if ( data[ p ] ) == 117 {
				goto _st154;
				
			}
			goto _ctr2;
			
		}
		_ctr358:
		{return st, err }
		goto _st154;
		_st154:
		if p == eof {
			goto _out154;
			
		}
		p+=1;
		st_case_154:
		if p == pe && p != eof {
			goto _out154;
			
		}
		if p == eof {
			goto _ctr358;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st155;
				
			}
			goto _ctr2;
			
		}
		_ctr360:
		{return st, err }
		goto _st155;
		_st155:
		if p == eof {
			goto _out155;
			
		}
		p+=1;
		st_case_155:
		if p == pe && p != eof {
			goto _out155;
			
		}
		if p == eof {
			goto _ctr360;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st156;
				
			}
			goto _ctr2;
			
		}
		_ctr362:
		{return st, err }
		goto _st156;
		_st156:
		if p == eof {
			goto _out156;
			
		}
		p+=1;
		st_case_156:
		if p == pe && p != eof {
			goto _out156;
			
		}
		if p == eof {
			goto _ctr362;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st157;
				
			}
			goto _ctr2;
			
		}
		_ctr364:
		{return st, err }
		goto _st157;
		_st157:
		if p == eof {
			goto _out157;
			
		}
		p+=1;
		st_case_157:
		if p == pe && p != eof {
			goto _out157;
			
		}
		if p == eof {
			goto _ctr364;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr354;
					
				}
				case 47:
				{
					goto _ctr354;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr365:
		{return st, err }
		goto _st158;
		_st158:
		if p == eof {
			goto _out158;
			
		}
		p+=1;
		st_case_158:
		if p == pe && p != eof {
			goto _out158;
			
		}
		if p == eof {
			goto _ctr365;
			
		} else {
			if ( data[ p ] ) == 105 {
				goto _st159;
				
			}
			goto _ctr2;
			
		}
		_ctr367:
		{return st, err }
		goto _st159;
		_st159:
		if p == eof {
			goto _out159;
			
		}
		p+=1;
		st_case_159:
		if p == pe && p != eof {
			goto _out159;
			
		}
		if p == eof {
			goto _ctr367;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st160;
					
				}
				case 44:
				{
					goto _st269;
					
				}
				case 100:
				{
					goto _st363;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr371:
		{return st, err }
		goto _st160;
		_st160:
		if p == eof {
			goto _out160;
			
		}
		p+=1;
		st_case_160:
		if p == pe && p != eof {
			goto _out160;
			
		}
		if p == eof {
			goto _ctr371;
			
		} else {
			switch ( data[ p ] ) {
				case 65:
				{
					goto _st161;
					
				}
				case 68:
				{
					goto _st233;
					
				}
				case 70:
				{
					goto _st236;
					
				}
				case 74:
				{
					goto _st244;
					
				}
				case 77:
				{
					goto _st254;
					
				}
				case 78:
				{
					goto _st260;
					
				}
				case 79:
				{
					goto _st263;
					
				}
				case 83:
				{
					goto _st266;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr380:
		{return st, err }
		goto _st161;
		_st161:
		if p == eof {
			goto _out161;
			
		}
		p+=1;
		st_case_161:
		if p == pe && p != eof {
			goto _out161;
			
		}
		if p == eof {
			goto _ctr380;
			
		} else {
			switch ( data[ p ] ) {
				case 112:
				{
					goto _st162;
					
				}
				case 117:
				{
					goto _st231;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr383:
		{return st, err }
		goto _st162;
		_st162:
		if p == eof {
			goto _out162;
			
		}
		p+=1;
		st_case_162:
		if p == pe && p != eof {
			goto _out162;
			
		}
		if p == eof {
			goto _ctr383;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st163;
				
			}
			goto _ctr2;
			
		}
		_ctr385:
		{return st, err }
		goto _st163;
		_st163:
		if p == eof {
			goto _out163;
			
		}
		p+=1;
		st_case_163:
		if p == pe && p != eof {
			goto _out163;
			
		}
		if p == eof {
			goto _ctr385;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr386;
					
				}
				case 105:
				{
					goto _st229;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr388:
		{return st, err }
		goto _st164;
		_ctr386:
		{st.Month = 4 }
		goto _st164;
		_ctr593:
		{st.Month = 8 }
		goto _st164;
		_ctr599:
		{st.Month = 12 }
		goto _st164;
		_ctr605:
		{st.Month = 2 }
		goto _st164;
		_ctr622:
		{st.Month = 1 }
		goto _st164;
		_ctr635:
		{st.Month = 7 }
		goto _st164;
		_ctr637:
		{st.Month = 6 }
		goto _st164;
		_ctr644:
		{st.Month = 3 }
		goto _st164;
		_ctr650:
		{st.Month = 5 }
		goto _st164;
		_ctr656:
		{st.Month = 11 }
		goto _st164;
		_ctr662:
		{st.Month = 10 }
		goto _st164;
		_ctr668:
		{st.Month = 9 }
		goto _st164;
		_st164:
		if p == eof {
			goto _out164;
			
		}
		p+=1;
		st_case_164:
		if p == pe && p != eof {
			goto _out164;
			
		}
		if p == eof {
			goto _ctr388;
			
		} else {
			switch ( data[ p ] ) {
				case 48:
				{
					goto _ctr389;
					
				}
				case 51:
				{
					goto _ctr391;
					
				}
				
			}
			if 49 <= ( data[ p ] ) && ( data[ p ] ) <= 50 {
				goto _ctr390;
				
			}
			goto _ctr2;
			
		}
		_ctr392:
		{return st, err }
		goto _st165;
		_ctr389:
		{pb = p }
		goto _st165;
		_st165:
		if p == eof {
			goto _out165;
			
		}
		p+=1;
		st_case_165:
		if p == pe && p != eof {
			goto _out165;
			
		}
		if p == eof {
			goto _ctr392;
			
		} else {
			if 49 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st166;
				
			}
			goto _ctr2;
			
		}
		_ctr394:
		{return st, err }
		goto _st166;
		_st166:
		if p == eof {
			goto _out166;
			
		}
		p+=1;
		st_case_166:
		if p == pe && p != eof {
			goto _out166;
			
		}
		if p == eof {
			goto _ctr394;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr395;
				
			}
			goto _ctr2;
			
		}
		_ctr396:
		{return st, err }
		goto _st167;
		_ctr395:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st167;
		_st167:
		if p == eof {
			goto _out167;
			
		}
		p+=1;
		st_case_167:
		if p == pe && p != eof {
			goto _out167;
			
		}
		if p == eof {
			goto _ctr396;
			
		} else {
			if ( data[ p ] ) == 50 {
				goto _ctr398;
				
			}
			if ( data[ p ] ) > 49 {
				if 51 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr399;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr397;
				
			}
			goto _ctr2;
			
		}
		_ctr400:
		{return st, err }
		goto _st168;
		_ctr397:
		{pb = p }
		goto _st168;
		_st168:
		if p == eof {
			goto _out168;
			
		}
		p+=1;
		st_case_168:
		if p == pe && p != eof {
			goto _out168;
			
		}
		if p == eof {
			goto _ctr400;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr401;
					
				}
				case 43:
				{
					goto _ctr402;
					
				}
				case 45:
				{
					goto _ctr403;
					
				}
				case 47:
				{
					goto _ctr404;
					
				}
				case 58:
				{
					goto _ctr406;
					
				}
				case 65:
				{
					goto _ctr407;
					
				}
				case 80:
				{
					goto _ctr407;
					
				}
				case 90:
				{
					goto _ctr408;
					
				}
				case 95:
				{
					goto _ctr404;
					
				}
				case 97:
				{
					goto _ctr409;
					
				}
				case 112:
				{
					goto _ctr409;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st191;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr404;
					
				}
				
			} else {
				goto _ctr404;
				
			}
			goto _ctr2;
			
		}
		_ctr410:
		{return st, err }
		goto _st169;
		_ctr401:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st169;
		_ctr467:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st169;
		_ctr551:
		{switch p - pb {
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
		}
		goto _st169;
		_ctr517:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st169;
		_ctr527:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st169;
		_ctr539:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st169;
		_ctr571:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st169;
		_st169:
		if p == eof {
			goto _out169;
			
		}
		p+=1;
		st_case_169:
		if p == pe && p != eof {
			goto _out169;
			
		}
		if p == eof {
			goto _ctr410;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st170;
					
				}
				case 43:
				{
					goto _st174;
					
				}
				case 45:
				{
					goto _st182;
					
				}
				case 47:
				{
					goto _ctr414;
					
				}
				case 65:
				{
					goto _ctr416;
					
				}
				case 80:
				{
					goto _ctr416;
					
				}
				case 90:
				{
					goto _ctr417;
					
				}
				case 95:
				{
					goto _ctr414;
					
				}
				case 97:
				{
					goto _ctr418;
					
				}
				case 112:
				{
					goto _ctr418;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr415;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr414;
					
				}
				
			} else {
				goto _ctr414;
				
			}
			goto _ctr2;
			
		}
		_ctr419:
		{return st, err }
		goto _st170;
		_ctr431:
		{for p - pb > 4 &&  data[pb] =='0' {
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
		}
		{st.Zoned = true }
		goto _st170;
		_ctr438:
		{st.Zoned = true }
		goto _st170;
		_ctr442:
		{switch p - pb {
				case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
				default:
				err = errors.New("invalid offset minute")
				return
			}
		}
		{st.Zoned = true }
		goto _st170;
		_ctr455:
		{st.ZoneName = data[pb:p]
			st.Zoned = true
		}
		{st.Zoned = true }
		goto _st170;
		_st170:
		if p == eof {
			goto _out170;
			
		}
		p+=1;
		st_case_170:
		if p == pe && p != eof {
			goto _out170;
			
		}
		if p == eof {
			goto _ctr419;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _ctr415;
				
			}
			goto _ctr2;
			
		}
		_ctr420:
		{return st, err }
		goto _st171;
		_ctr415:
		{pb = p }
		goto _st171;
		_st171:
		if p == eof {
			goto _out171;
			
		}
		p+=1;
		st_case_171:
		if p == pe && p != eof {
			goto _out171;
			
		}
		if p == eof {
			goto _ctr420;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st172;
				
			}
			goto _ctr2;
			
		}
		_ctr422:
		{return st, err }
		goto _st172;
		_st172:
		if p == eof {
			goto _out172;
			
		}
		p+=1;
		st_case_172:
		if p == pe && p != eof {
			goto _out172;
			
		}
		if p == eof {
			goto _ctr422;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st173;
				
			}
			goto _ctr2;
			
		}
		_ctr424:
		{return st, err }
		goto _st173;
		_st173:
		if p == eof {
			goto _out173;
			
		}
		p+=1;
		st_case_173:
		if p == pe && p != eof {
			goto _out173;
			
		}
		if p == eof {
			goto _ctr424;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st463;
				
			}
			goto _ctr2;
			
		}
		_ctr1155:
		{st.Year, _ = strconv.Atoi(data[pb:pb+4])
		}
		goto _st463;
		_st463:
		if p == eof {
			goto _out463;
			
		}
		p+=1;
		st_case_463:
		if p == pe && p != eof {
			goto _out463;
			
		}
		if p == eof {
			goto _ctr1155;
			
		} else {
			goto _st0;
			
		}
		_ctr426:
		{return st, err }
		goto _st174;
		_ctr402:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st174;
		_ctr460:
		{if st.Hour > 12 {
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
		}
		goto _st174;
		_ctr468:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st174;
		_ctr480:
		{switch p - pb {
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
		}
		goto _st174;
		_ctr490:
		{switch p - pb {
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
		}
		goto _st174;
		_ctr518:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st174;
		_ctr528:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st174;
		_ctr540:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st174;
		_ctr572:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st174;
		_st174:
		if p == eof {
			goto _out174;
			
		}
		p+=1;
		st_case_174:
		if p == pe && p != eof {
			goto _out174;
			
		}
		if p == eof {
			goto _ctr426;
			
		} else {
			if ( data[ p ] ) == 50 {
				goto _ctr428;
				
			}
			if ( data[ p ] ) > 49 {
				if 51 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr429;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr427;
				
			}
			goto _ctr2;
			
		}
		_ctr430:
		{return st, err }
		goto _st175;
		_ctr427:
		{pb = p }
		goto _st175;
		_ctr447:
		{st.NegtiveZoneOffset = true }
		{pb = p }
		goto _st175;
		_st175:
		if p == eof {
			goto _out175;
			
		}
		p+=1;
		st_case_175:
		if p == pe && p != eof {
			goto _out175;
			
		}
		if p == eof {
			goto _ctr430;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr431;
					
				}
				case 58:
				{
					goto _ctr433;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st176;
				
			}
			goto _ctr2;
			
		}
		_ctr434:
		{return st, err }
		goto _st176;
		_ctr429:
		{pb = p }
		goto _st176;
		_ctr449:
		{st.NegtiveZoneOffset = true }
		{pb = p }
		goto _st176;
		_st176:
		if p == eof {
			goto _out176;
			
		}
		p+=1;
		st_case_176:
		if p == pe && p != eof {
			goto _out176;
			
		}
		if p == eof {
			goto _ctr434;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr431;
					
				}
				case 58:
				{
					goto _ctr433;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st177;
				
			}
			goto _ctr2;
			
		}
		_ctr436:
		{return st, err }
		goto _st177;
		_st177:
		if p == eof {
			goto _out177;
			
		}
		p+=1;
		st_case_177:
		if p == pe && p != eof {
			goto _out177;
			
		}
		if p == eof {
			goto _ctr436;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr431;
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st177;
				
			}
			goto _ctr2;
			
		}
		_ctr437:
		{return st, err }
		goto _st178;
		_ctr433:
		{switch p - pb {
				case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
				default:
				err = errors.New("invalid offset hour")
				return
			}
		}
		goto _st178;
		_st178:
		if p == eof {
			goto _out178;
			
		}
		p+=1;
		st_case_178:
		if p == pe && p != eof {
			goto _out178;
			
		}
		if p == eof {
			goto _ctr437;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr438;
				
			}
			if ( data[ p ] ) > 53 {
				if ( data[ p ] ) <= 57 {
					goto _ctr440;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr439;
				
			}
			goto _ctr2;
			
		}
		_ctr441:
		{return st, err }
		goto _st179;
		_ctr439:
		{pb = p }
		goto _st179;
		_st179:
		if p == eof {
			goto _out179;
			
		}
		p+=1;
		st_case_179:
		if p == pe && p != eof {
			goto _out179;
			
		}
		if p == eof {
			goto _ctr441;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr442;
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st180;
				
			}
			goto _ctr2;
			
		}
		_ctr444:
		{return st, err }
		goto _st180;
		_ctr440:
		{pb = p }
		goto _st180;
		_st180:
		if p == eof {
			goto _out180;
			
		}
		p+=1;
		st_case_180:
		if p == pe && p != eof {
			goto _out180;
			
		}
		if p == eof {
			goto _ctr444;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr442;
				
			}
			goto _ctr2;
			
		}
		_ctr445:
		{return st, err }
		goto _st181;
		_ctr428:
		{pb = p }
		goto _st181;
		_ctr448:
		{st.NegtiveZoneOffset = true }
		{pb = p }
		goto _st181;
		_st181:
		if p == eof {
			goto _out181;
			
		}
		p+=1;
		st_case_181:
		if p == pe && p != eof {
			goto _out181;
			
		}
		if p == eof {
			goto _ctr445;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr431;
					
				}
				case 58:
				{
					goto _ctr433;
					
				}
				
			}
			if ( data[ p ] ) > 51 {
				if ( data[ p ] ) <= 57 {
					goto _st177;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _st176;
				
			}
			goto _ctr2;
			
		}
		_ctr446:
		{return st, err }
		goto _st182;
		_ctr403:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st182;
		_ctr461:
		{if st.Hour > 12 {
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
		}
		goto _st182;
		_ctr469:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st182;
		_ctr481:
		{switch p - pb {
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
		}
		goto _st182;
		_ctr491:
		{switch p - pb {
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
		}
		goto _st182;
		_ctr519:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st182;
		_ctr529:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st182;
		_ctr541:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st182;
		_ctr573:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st182;
		_st182:
		if p == eof {
			goto _out182;
			
		}
		p+=1;
		st_case_182:
		if p == pe && p != eof {
			goto _out182;
			
		}
		if p == eof {
			goto _ctr446;
			
		} else {
			if ( data[ p ] ) == 50 {
				goto _ctr448;
				
			}
			if ( data[ p ] ) > 49 {
				if 51 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr449;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr447;
				
			}
			goto _ctr2;
			
		}
		_ctr450:
		{return st, err }
		goto _st183;
		_ctr414:
		{pb = p }
		goto _st183;
		_ctr404:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st183;
		_ctr470:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st183;
		_ctr483:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st183;
		_ctr492:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st183;
		_ctr520:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st183;
		_ctr530:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st183;
		_ctr543:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st183;
		_ctr575:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st183;
		_st183:
		if p == eof {
			goto _out183;
			
		}
		p+=1;
		st_case_183:
		if p == pe && p != eof {
			goto _out183;
			
		}
		if p == eof {
			goto _ctr450;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st184;
					
				}
				case 95:
				{
					goto _st184;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st184;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st184;
				
			}
			goto _ctr2;
			
		}
		_ctr452:
		{return st, err }
		goto _st184;
		_st184:
		if p == eof {
			goto _out184;
			
		}
		p+=1;
		st_case_184:
		if p == pe && p != eof {
			goto _out184;
			
		}
		if p == eof {
			goto _ctr452;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st185;
					
				}
				case 95:
				{
					goto _st185;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st185;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st185;
				
			}
			goto _ctr2;
			
		}
		_ctr454:
		{return st, err }
		goto _st185;
		_ctr462:
		{if st.Hour > 12 {
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
		}
		{pb = p }
		goto _st185;
		_st185:
		if p == eof {
			goto _out185;
			
		}
		p+=1;
		st_case_185:
		if p == pe && p != eof {
			goto _out185;
			
		}
		if p == eof {
			goto _ctr454;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr455;
					
				}
				case 47:
				{
					goto _st185;
					
				}
				case 95:
				{
					goto _st185;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st185;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st185;
				
			}
			goto _ctr2;
			
		}
		_ctr456:
		{return st, err }
		goto _st186;
		_ctr416:
		{pb = p }
		goto _st186;
		_ctr407:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st186;
		_ctr473:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st186;
		_ctr553:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st186;
		_ctr523:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st186;
		_ctr532:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st186;
		_ctr545:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st186;
		_ctr576:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st186;
		_st186:
		if p == eof {
			goto _out186;
			
		}
		p+=1;
		st_case_186:
		if p == pe && p != eof {
			goto _out186;
			
		}
		if p == eof {
			goto _ctr456;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st184;
					
				}
				case 77:
				{
					goto _st187;
					
				}
				case 95:
				{
					goto _st184;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st184;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st184;
				
			}
			goto _ctr2;
			
		}
		_ctr458:
		{return st, err }
		goto _st187;
		_st187:
		if p == eof {
			goto _out187;
			
		}
		p+=1;
		st_case_187:
		if p == pe && p != eof {
			goto _out187;
			
		}
		if p == eof {
			goto _ctr458;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr459;
					
				}
				case 43:
				{
					goto _ctr460;
					
				}
				case 45:
				{
					goto _ctr461;
					
				}
				case 47:
				{
					goto _ctr462;
					
				}
				case 95:
				{
					goto _ctr462;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr462;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr462;
				
			}
			goto _ctr2;
			
		}
		_ctr463:
		{return st, err }
		goto _st188;
		_ctr459:
		{if st.Hour > 12 {
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
		}
		goto _st188;
		_ctr479:
		{switch p - pb {
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
		}
		goto _st188;
		_ctr489:
		{switch p - pb {
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
		}
		goto _st188;
		_st188:
		if p == eof {
			goto _out188;
			
		}
		p+=1;
		st_case_188:
		if p == pe && p != eof {
			goto _out188;
			
		}
		if p == eof {
			goto _ctr463;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st170;
					
				}
				case 43:
				{
					goto _st174;
					
				}
				case 45:
				{
					goto _st182;
					
				}
				case 47:
				{
					goto _ctr414;
					
				}
				case 90:
				{
					goto _ctr417;
					
				}
				case 95:
				{
					goto _ctr414;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr415;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr414;
					
				}
				
			} else {
				goto _ctr414;
				
			}
			goto _ctr2;
			
		}
		_ctr464:
		{return st, err }
		goto _st189;
		_ctr417:
		{pb = p }
		goto _st189;
		_ctr408:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st189;
		_ctr474:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st189;
		_ctr485:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st189;
		_ctr494:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st189;
		_ctr524:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st189;
		_ctr533:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st189;
		_ctr546:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st189;
		_ctr577:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st189;
		_st189:
		if p == eof {
			goto _out189;
			
		}
		p+=1;
		st_case_189:
		if p == pe && p != eof {
			goto _out189;
			
		}
		if p == eof {
			goto _ctr464;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr438;
					
				}
				case 47:
				{
					goto _st184;
					
				}
				case 95:
				{
					goto _st184;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st184;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st184;
				
			}
			goto _ctr2;
			
		}
		_ctr465:
		{return st, err }
		goto _st190;
		_ctr418:
		{pb = p }
		goto _st190;
		_ctr409:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st190;
		_ctr475:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st190;
		_ctr554:
		{switch p - pb {
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
		}
		{pb = p }
		goto _st190;
		_ctr525:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st190;
		_ctr534:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st190;
		_ctr547:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		{pb = p }
		goto _st190;
		_ctr578:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		{pb = p }
		goto _st190;
		_st190:
		if p == eof {
			goto _out190;
			
		}
		p+=1;
		st_case_190:
		if p == pe && p != eof {
			goto _out190;
			
		}
		if p == eof {
			goto _ctr465;
			
		} else {
			switch ( data[ p ] ) {
				case 47:
				{
					goto _st184;
					
				}
				case 95:
				{
					goto _st184;
					
				}
				case 109:
				{
					goto _st187;
					
				}
				
			}
			if ( data[ p ] ) > 90 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _st184;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _st184;
				
			}
			goto _ctr2;
			
		}
		_ctr466:
		{return st, err }
		goto _st191;
		_st191:
		if p == eof {
			goto _out191;
			
		}
		p+=1;
		st_case_191:
		if p == pe && p != eof {
			goto _out191;
			
		}
		if p == eof {
			goto _ctr466;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr467;
					
				}
				case 43:
				{
					goto _ctr468;
					
				}
				case 45:
				{
					goto _ctr469;
					
				}
				case 47:
				{
					goto _ctr470;
					
				}
				case 58:
				{
					goto _ctr472;
					
				}
				case 65:
				{
					goto _ctr473;
					
				}
				case 80:
				{
					goto _ctr473;
					
				}
				case 90:
				{
					goto _ctr474;
					
				}
				case 95:
				{
					goto _ctr470;
					
				}
				case 97:
				{
					goto _ctr475;
					
				}
				case 112:
				{
					goto _ctr475;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st192;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr470;
					
				}
				
			} else {
				goto _ctr470;
				
			}
			goto _ctr2;
			
		}
		_ctr476:
		{return st, err }
		goto _st192;
		_st192:
		if p == eof {
			goto _out192;
			
		}
		p+=1;
		st_case_192:
		if p == pe && p != eof {
			goto _out192;
			
		}
		if p == eof {
			goto _ctr476;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st193;
				
			}
			goto _ctr2;
			
		}
		_ctr478:
		{return st, err }
		goto _st193;
		_st193:
		if p == eof {
			goto _out193;
			
		}
		p+=1;
		st_case_193:
		if p == pe && p != eof {
			goto _out193;
			
		}
		if p == eof {
			goto _ctr478;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr479;
					
				}
				case 43:
				{
					goto _ctr480;
					
				}
				case 45:
				{
					goto _ctr481;
					
				}
				case 46:
				{
					goto _ctr482;
					
				}
				case 47:
				{
					goto _ctr483;
					
				}
				case 90:
				{
					goto _ctr485;
					
				}
				case 95:
				{
					goto _ctr483;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st204;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr483;
					
				}
				
			} else {
				goto _ctr483;
				
			}
			goto _ctr2;
			
		}
		_ctr486:
		{return st, err }
		goto _st194;
		_ctr482:
		{switch p - pb {
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
		}
		goto _st194;
		_st194:
		if p == eof {
			goto _out194;
			
		}
		p+=1;
		st_case_194:
		if p == pe && p != eof {
			goto _out194;
			
		}
		if p == eof {
			goto _ctr486;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _ctr487;
				
			}
			goto _ctr2;
			
		}
		_ctr488:
		{return st, err }
		goto _st195;
		_ctr487:
		{pb = p }
		goto _st195;
		_st195:
		if p == eof {
			goto _out195;
			
		}
		p+=1;
		st_case_195:
		if p == pe && p != eof {
			goto _out195;
			
		}
		if p == eof {
			goto _ctr488;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr489;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st196;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr495:
		{return st, err }
		goto _st196;
		_st196:
		if p == eof {
			goto _out196;
			
		}
		p+=1;
		st_case_196:
		if p == pe && p != eof {
			goto _out196;
			
		}
		if p == eof {
			goto _ctr495;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr489;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st197;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr497:
		{return st, err }
		goto _st197;
		_st197:
		if p == eof {
			goto _out197;
			
		}
		p+=1;
		st_case_197:
		if p == pe && p != eof {
			goto _out197;
			
		}
		if p == eof {
			goto _ctr497;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr489;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st198;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr499:
		{return st, err }
		goto _st198;
		_st198:
		if p == eof {
			goto _out198;
			
		}
		p+=1;
		st_case_198:
		if p == pe && p != eof {
			goto _out198;
			
		}
		if p == eof {
			goto _ctr499;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr489;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st199;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr501:
		{return st, err }
		goto _st199;
		_st199:
		if p == eof {
			goto _out199;
			
		}
		p+=1;
		st_case_199:
		if p == pe && p != eof {
			goto _out199;
			
		}
		if p == eof {
			goto _ctr501;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr489;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st200;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr503:
		{return st, err }
		goto _st200;
		_st200:
		if p == eof {
			goto _out200;
			
		}
		p+=1;
		st_case_200:
		if p == pe && p != eof {
			goto _out200;
			
		}
		if p == eof {
			goto _ctr503;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr489;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st201;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr505:
		{return st, err }
		goto _st201;
		_st201:
		if p == eof {
			goto _out201;
			
		}
		p+=1;
		st_case_201:
		if p == pe && p != eof {
			goto _out201;
			
		}
		if p == eof {
			goto _ctr505;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr489;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st202;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr507:
		{return st, err }
		goto _st202;
		_st202:
		if p == eof {
			goto _out202;
			
		}
		p+=1;
		st_case_202:
		if p == pe && p != eof {
			goto _out202;
			
		}
		if p == eof {
			goto _ctr507;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr489;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				
			}
			if ( data[ p ] ) < 65 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st203;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr509:
		{return st, err }
		goto _st203;
		_st203:
		if p == eof {
			goto _out203;
			
		}
		p+=1;
		st_case_203:
		if p == pe && p != eof {
			goto _out203;
			
		}
		if p == eof {
			goto _ctr509;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr489;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr510:
		{return st, err }
		goto _st204;
		_st204:
		if p == eof {
			goto _out204;
			
		}
		p+=1;
		st_case_204:
		if p == pe && p != eof {
			goto _out204;
			
		}
		if p == eof {
			goto _ctr510;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st205;
				
			}
			goto _ctr2;
			
		}
		_ctr512:
		{return st, err }
		goto _st205;
		_st205:
		if p == eof {
			goto _out205;
			
		}
		p+=1;
		st_case_205:
		if p == pe && p != eof {
			goto _out205;
			
		}
		if p == eof {
			goto _ctr512;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr479;
					
				}
				case 43:
				{
					goto _ctr480;
					
				}
				case 45:
				{
					goto _ctr481;
					
				}
				case 46:
				{
					goto _ctr482;
					
				}
				case 47:
				{
					goto _ctr483;
					
				}
				case 90:
				{
					goto _ctr485;
					
				}
				case 95:
				{
					goto _ctr483;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 97 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr483;
					
				}
				
			} else if ( data[ p ] ) >= 65 {
				goto _ctr483;
				
			}
			goto _ctr2;
			
		}
		_ctr513:
		{return st, err }
		goto _st206;
		_ctr406:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st206;
		_ctr472:
		{st.Hour, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st206;
		_st206:
		if p == eof {
			goto _out206;
			
		}
		p+=1;
		st_case_206:
		if p == pe && p != eof {
			goto _out206;
			
		}
		if p == eof {
			goto _ctr513;
			
		} else {
			if ( data[ p ] ) > 53 {
				if ( data[ p ] ) <= 57 {
					goto _ctr515;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr514;
				
			}
			goto _ctr2;
			
		}
		_ctr516:
		{return st, err }
		goto _st207;
		_ctr514:
		{pb = p }
		goto _st207;
		_st207:
		if p == eof {
			goto _out207;
			
		}
		p+=1;
		st_case_207:
		if p == pe && p != eof {
			goto _out207;
			
		}
		if p == eof {
			goto _ctr516;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr517;
					
				}
				case 43:
				{
					goto _ctr518;
					
				}
				case 45:
				{
					goto _ctr519;
					
				}
				case 47:
				{
					goto _ctr520;
					
				}
				case 58:
				{
					goto _ctr522;
					
				}
				case 65:
				{
					goto _ctr523;
					
				}
				case 80:
				{
					goto _ctr523;
					
				}
				case 90:
				{
					goto _ctr524;
					
				}
				case 95:
				{
					goto _ctr520;
					
				}
				case 97:
				{
					goto _ctr525;
					
				}
				case 112:
				{
					goto _ctr525;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st208;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr520;
					
				}
				
			} else {
				goto _ctr520;
				
			}
			goto _ctr2;
			
		}
		_ctr526:
		{return st, err }
		goto _st208;
		_st208:
		if p == eof {
			goto _out208;
			
		}
		p+=1;
		st_case_208:
		if p == pe && p != eof {
			goto _out208;
			
		}
		if p == eof {
			goto _ctr526;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr527;
					
				}
				case 43:
				{
					goto _ctr528;
					
				}
				case 45:
				{
					goto _ctr529;
					
				}
				case 47:
				{
					goto _ctr530;
					
				}
				case 58:
				{
					goto _ctr531;
					
				}
				case 65:
				{
					goto _ctr532;
					
				}
				case 80:
				{
					goto _ctr532;
					
				}
				case 90:
				{
					goto _ctr533;
					
				}
				case 95:
				{
					goto _ctr530;
					
				}
				case 97:
				{
					goto _ctr534;
					
				}
				case 112:
				{
					goto _ctr534;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr530;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr530;
				
			}
			goto _ctr2;
			
		}
		_ctr535:
		{return st, err }
		goto _st209;
		_ctr522:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st209;
		_ctr531:
		{st.Minute, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st209;
		_st209:
		if p == eof {
			goto _out209;
			
		}
		p+=1;
		st_case_209:
		if p == pe && p != eof {
			goto _out209;
			
		}
		if p == eof {
			goto _ctr535;
			
		} else {
			if ( data[ p ] ) > 53 {
				if ( data[ p ] ) <= 57 {
					goto _ctr537;
					
				}
				
			} else if ( data[ p ] ) >= 48 {
				goto _ctr536;
				
			}
			goto _ctr2;
			
		}
		_ctr538:
		{return st, err }
		goto _st210;
		_ctr536:
		{pb = p }
		goto _st210;
		_st210:
		if p == eof {
			goto _out210;
			
		}
		p+=1;
		st_case_210:
		if p == pe && p != eof {
			goto _out210;
			
		}
		if p == eof {
			goto _ctr538;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr539;
					
				}
				case 43:
				{
					goto _ctr540;
					
				}
				case 45:
				{
					goto _ctr541;
					
				}
				case 46:
				{
					goto _ctr542;
					
				}
				case 47:
				{
					goto _ctr543;
					
				}
				case 65:
				{
					goto _ctr545;
					
				}
				case 80:
				{
					goto _ctr545;
					
				}
				case 90:
				{
					goto _ctr546;
					
				}
				case 95:
				{
					goto _ctr543;
					
				}
				case 97:
				{
					goto _ctr547;
					
				}
				case 112:
				{
					goto _ctr547;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st221;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr543;
					
				}
				
			} else {
				goto _ctr543;
				
			}
			goto _ctr2;
			
		}
		_ctr548:
		{return st, err }
		goto _st211;
		_ctr542:
		{st.Second, _ = strconv.Atoi(data[pb:pb+1])
		}
		goto _st211;
		_ctr574:
		{st.Second, _ = strconv.Atoi(data[pb:pb+2])
		}
		goto _st211;
		_st211:
		if p == eof {
			goto _out211;
			
		}
		p+=1;
		st_case_211:
		if p == pe && p != eof {
			goto _out211;
			
		}
		if p == eof {
			goto _ctr548;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _ctr549;
				
			}
			goto _ctr2;
			
		}
		_ctr550:
		{return st, err }
		goto _st212;
		_ctr549:
		{pb = p }
		goto _st212;
		_st212:
		if p == eof {
			goto _out212;
			
		}
		p+=1;
		st_case_212:
		if p == pe && p != eof {
			goto _out212;
			
		}
		if p == eof {
			goto _ctr550;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr551;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 65:
				{
					goto _ctr553;
					
				}
				case 80:
				{
					goto _ctr553;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				case 97:
				{
					goto _ctr554;
					
				}
				case 112:
				{
					goto _ctr554;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st213;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr555:
		{return st, err }
		goto _st213;
		_st213:
		if p == eof {
			goto _out213;
			
		}
		p+=1;
		st_case_213:
		if p == pe && p != eof {
			goto _out213;
			
		}
		if p == eof {
			goto _ctr555;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr551;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 65:
				{
					goto _ctr553;
					
				}
				case 80:
				{
					goto _ctr553;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				case 97:
				{
					goto _ctr554;
					
				}
				case 112:
				{
					goto _ctr554;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st214;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr557:
		{return st, err }
		goto _st214;
		_st214:
		if p == eof {
			goto _out214;
			
		}
		p+=1;
		st_case_214:
		if p == pe && p != eof {
			goto _out214;
			
		}
		if p == eof {
			goto _ctr557;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr551;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 65:
				{
					goto _ctr553;
					
				}
				case 80:
				{
					goto _ctr553;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				case 97:
				{
					goto _ctr554;
					
				}
				case 112:
				{
					goto _ctr554;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st215;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr559:
		{return st, err }
		goto _st215;
		_st215:
		if p == eof {
			goto _out215;
			
		}
		p+=1;
		st_case_215:
		if p == pe && p != eof {
			goto _out215;
			
		}
		if p == eof {
			goto _ctr559;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr551;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 65:
				{
					goto _ctr553;
					
				}
				case 80:
				{
					goto _ctr553;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				case 97:
				{
					goto _ctr554;
					
				}
				case 112:
				{
					goto _ctr554;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st216;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr561:
		{return st, err }
		goto _st216;
		_st216:
		if p == eof {
			goto _out216;
			
		}
		p+=1;
		st_case_216:
		if p == pe && p != eof {
			goto _out216;
			
		}
		if p == eof {
			goto _ctr561;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr551;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 65:
				{
					goto _ctr553;
					
				}
				case 80:
				{
					goto _ctr553;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				case 97:
				{
					goto _ctr554;
					
				}
				case 112:
				{
					goto _ctr554;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st217;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr563:
		{return st, err }
		goto _st217;
		_st217:
		if p == eof {
			goto _out217;
			
		}
		p+=1;
		st_case_217:
		if p == pe && p != eof {
			goto _out217;
			
		}
		if p == eof {
			goto _ctr563;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr551;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 65:
				{
					goto _ctr553;
					
				}
				case 80:
				{
					goto _ctr553;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				case 97:
				{
					goto _ctr554;
					
				}
				case 112:
				{
					goto _ctr554;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st218;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr565:
		{return st, err }
		goto _st218;
		_st218:
		if p == eof {
			goto _out218;
			
		}
		p+=1;
		st_case_218:
		if p == pe && p != eof {
			goto _out218;
			
		}
		if p == eof {
			goto _ctr565;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr551;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 65:
				{
					goto _ctr553;
					
				}
				case 80:
				{
					goto _ctr553;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				case 97:
				{
					goto _ctr554;
					
				}
				case 112:
				{
					goto _ctr554;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st219;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr567:
		{return st, err }
		goto _st219;
		_st219:
		if p == eof {
			goto _out219;
			
		}
		p+=1;
		st_case_219:
		if p == pe && p != eof {
			goto _out219;
			
		}
		if p == eof {
			goto _ctr567;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr551;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 65:
				{
					goto _ctr553;
					
				}
				case 80:
				{
					goto _ctr553;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				case 97:
				{
					goto _ctr554;
					
				}
				case 112:
				{
					goto _ctr554;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st220;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr569:
		{return st, err }
		goto _st220;
		_st220:
		if p == eof {
			goto _out220;
			
		}
		p+=1;
		st_case_220:
		if p == pe && p != eof {
			goto _out220;
			
		}
		if p == eof {
			goto _ctr569;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr551;
					
				}
				case 43:
				{
					goto _ctr490;
					
				}
				case 45:
				{
					goto _ctr491;
					
				}
				case 47:
				{
					goto _ctr492;
					
				}
				case 65:
				{
					goto _ctr553;
					
				}
				case 80:
				{
					goto _ctr553;
					
				}
				case 90:
				{
					goto _ctr494;
					
				}
				case 95:
				{
					goto _ctr492;
					
				}
				case 97:
				{
					goto _ctr554;
					
				}
				case 112:
				{
					goto _ctr554;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr492;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr492;
				
			}
			goto _ctr2;
			
		}
		_ctr570:
		{return st, err }
		goto _st221;
		_st221:
		if p == eof {
			goto _out221;
			
		}
		p+=1;
		st_case_221:
		if p == pe && p != eof {
			goto _out221;
			
		}
		if p == eof {
			goto _ctr570;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr571;
					
				}
				case 43:
				{
					goto _ctr572;
					
				}
				case 45:
				{
					goto _ctr573;
					
				}
				case 46:
				{
					goto _ctr574;
					
				}
				case 47:
				{
					goto _ctr575;
					
				}
				case 65:
				{
					goto _ctr576;
					
				}
				case 80:
				{
					goto _ctr576;
					
				}
				case 90:
				{
					goto _ctr577;
					
				}
				case 95:
				{
					goto _ctr575;
					
				}
				case 97:
				{
					goto _ctr578;
					
				}
				case 112:
				{
					goto _ctr578;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr575;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr575;
				
			}
			goto _ctr2;
			
		}
		_ctr579:
		{return st, err }
		goto _st222;
		_ctr537:
		{pb = p }
		goto _st222;
		_st222:
		if p == eof {
			goto _out222;
			
		}
		p+=1;
		st_case_222:
		if p == pe && p != eof {
			goto _out222;
			
		}
		if p == eof {
			goto _ctr579;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr539;
					
				}
				case 43:
				{
					goto _ctr540;
					
				}
				case 45:
				{
					goto _ctr541;
					
				}
				case 46:
				{
					goto _ctr542;
					
				}
				case 47:
				{
					goto _ctr543;
					
				}
				case 65:
				{
					goto _ctr545;
					
				}
				case 80:
				{
					goto _ctr545;
					
				}
				case 90:
				{
					goto _ctr546;
					
				}
				case 95:
				{
					goto _ctr543;
					
				}
				case 97:
				{
					goto _ctr547;
					
				}
				case 112:
				{
					goto _ctr547;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr543;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr543;
				
			}
			goto _ctr2;
			
		}
		_ctr580:
		{return st, err }
		goto _st223;
		_ctr515:
		{pb = p }
		goto _st223;
		_st223:
		if p == eof {
			goto _out223;
			
		}
		p+=1;
		st_case_223:
		if p == pe && p != eof {
			goto _out223;
			
		}
		if p == eof {
			goto _ctr580;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr517;
					
				}
				case 43:
				{
					goto _ctr518;
					
				}
				case 45:
				{
					goto _ctr519;
					
				}
				case 47:
				{
					goto _ctr520;
					
				}
				case 58:
				{
					goto _ctr522;
					
				}
				case 65:
				{
					goto _ctr523;
					
				}
				case 80:
				{
					goto _ctr523;
					
				}
				case 90:
				{
					goto _ctr524;
					
				}
				case 95:
				{
					goto _ctr520;
					
				}
				case 97:
				{
					goto _ctr525;
					
				}
				case 112:
				{
					goto _ctr525;
					
				}
				
			}
			if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr520;
					
				}
				
			} else if ( data[ p ] ) >= 66 {
				goto _ctr520;
				
			}
			goto _ctr2;
			
		}
		_ctr581:
		{return st, err }
		goto _st224;
		_ctr398:
		{pb = p }
		goto _st224;
		_st224:
		if p == eof {
			goto _out224;
			
		}
		p+=1;
		st_case_224:
		if p == pe && p != eof {
			goto _out224;
			
		}
		if p == eof {
			goto _ctr581;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr401;
					
				}
				case 43:
				{
					goto _ctr402;
					
				}
				case 45:
				{
					goto _ctr403;
					
				}
				case 47:
				{
					goto _ctr404;
					
				}
				case 58:
				{
					goto _ctr406;
					
				}
				case 65:
				{
					goto _ctr407;
					
				}
				case 80:
				{
					goto _ctr407;
					
				}
				case 90:
				{
					goto _ctr408;
					
				}
				case 95:
				{
					goto _ctr404;
					
				}
				case 97:
				{
					goto _ctr409;
					
				}
				case 112:
				{
					goto _ctr409;
					
				}
				
			}
			if ( data[ p ] ) < 52 {
				if 48 <= ( data[ p ] ) {
					goto _st191;
					
				}
				
			} else if ( data[ p ] ) > 57 {
				if ( data[ p ] ) > 89 {
					if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
						goto _ctr404;
						
					}
					
				} else if ( data[ p ] ) >= 66 {
					goto _ctr404;
					
				}
				
			} else {
				goto _st225;
				
			}
			goto _ctr2;
			
		}
		_ctr583:
		{return st, err }
		goto _st225;
		_st225:
		if p == eof {
			goto _out225;
			
		}
		p+=1;
		st_case_225:
		if p == pe && p != eof {
			goto _out225;
			
		}
		if p == eof {
			goto _ctr583;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st192;
				
			}
			goto _ctr2;
			
		}
		_ctr584:
		{return st, err }
		goto _st226;
		_ctr399:
		{pb = p }
		goto _st226;
		_st226:
		if p == eof {
			goto _out226;
			
		}
		p+=1;
		st_case_226:
		if p == pe && p != eof {
			goto _out226;
			
		}
		if p == eof {
			goto _ctr584;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr401;
					
				}
				case 43:
				{
					goto _ctr402;
					
				}
				case 45:
				{
					goto _ctr403;
					
				}
				case 47:
				{
					goto _ctr404;
					
				}
				case 58:
				{
					goto _ctr406;
					
				}
				case 65:
				{
					goto _ctr407;
					
				}
				case 80:
				{
					goto _ctr407;
					
				}
				case 90:
				{
					goto _ctr408;
					
				}
				case 95:
				{
					goto _ctr404;
					
				}
				case 97:
				{
					goto _ctr409;
					
				}
				case 112:
				{
					goto _ctr409;
					
				}
				
			}
			if ( data[ p ] ) < 66 {
				if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _st225;
					
				}
				
			} else if ( data[ p ] ) > 89 {
				if 98 <= ( data[ p ] ) && ( data[ p ] ) <= 122 {
					goto _ctr404;
					
				}
				
			} else {
				goto _ctr404;
				
			}
			goto _ctr2;
			
		}
		_ctr585:
		{return st, err }
		goto _st227;
		_ctr390:
		{pb = p }
		goto _st227;
		_st227:
		if p == eof {
			goto _out227;
			
		}
		p+=1;
		st_case_227:
		if p == pe && p != eof {
			goto _out227;
			
		}
		if p == eof {
			goto _ctr585;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st166;
				
			}
			goto _ctr2;
			
		}
		_ctr586:
		{return st, err }
		goto _st228;
		_ctr391:
		{pb = p }
		goto _st228;
		_st228:
		if p == eof {
			goto _out228;
			
		}
		p+=1;
		st_case_228:
		if p == pe && p != eof {
			goto _out228;
			
		}
		if p == eof {
			goto _ctr586;
			
		} else {
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 49 {
				goto _st166;
				
			}
			goto _ctr2;
			
		}
		_ctr587:
		{return st, err }
		goto _st229;
		_st229:
		if p == eof {
			goto _out229;
			
		}
		p+=1;
		st_case_229:
		if p == pe && p != eof {
			goto _out229;
			
		}
		if p == eof {
			goto _ctr587;
			
		} else {
			if ( data[ p ] ) == 108 {
				goto _st230;
				
			}
			goto _ctr2;
			
		}
		_ctr589:
		{return st, err }
		goto _st230;
		_st230:
		if p == eof {
			goto _out230;
			
		}
		p+=1;
		st_case_230:
		if p == pe && p != eof {
			goto _out230;
			
		}
		if p == eof {
			goto _ctr589;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr386;
				
			}
			goto _ctr2;
			
		}
		_ctr590:
		{return st, err }
		goto _st231;
		_st231:
		if p == eof {
			goto _out231;
			
		}
		p+=1;
		st_case_231:
		if p == pe && p != eof {
			goto _out231;
			
		}
		if p == eof {
			goto _ctr590;
			
		} else {
			if ( data[ p ] ) == 103 {
				goto _st232;
				
			}
			goto _ctr2;
			
		}
		_ctr592:
		{return st, err }
		goto _st232;
		_st232:
		if p == eof {
			goto _out232;
			
		}
		p+=1;
		st_case_232:
		if p == pe && p != eof {
			goto _out232;
			
		}
		if p == eof {
			goto _ctr592;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr593;
				
			}
			goto _ctr2;
			
		}
		_ctr594:
		{return st, err }
		goto _st233;
		_st233:
		if p == eof {
			goto _out233;
			
		}
		p+=1;
		st_case_233:
		if p == pe && p != eof {
			goto _out233;
			
		}
		if p == eof {
			goto _ctr594;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st234;
				
			}
			goto _ctr2;
			
		}
		_ctr596:
		{return st, err }
		goto _st234;
		_st234:
		if p == eof {
			goto _out234;
			
		}
		p+=1;
		st_case_234:
		if p == pe && p != eof {
			goto _out234;
			
		}
		if p == eof {
			goto _ctr596;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st235;
				
			}
			goto _ctr2;
			
		}
		_ctr598:
		{return st, err }
		goto _st235;
		_st235:
		if p == eof {
			goto _out235;
			
		}
		p+=1;
		st_case_235:
		if p == pe && p != eof {
			goto _out235;
			
		}
		if p == eof {
			goto _ctr598;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr599;
				
			}
			goto _ctr2;
			
		}
		_ctr600:
		{return st, err }
		goto _st236;
		_st236:
		if p == eof {
			goto _out236;
			
		}
		p+=1;
		st_case_236:
		if p == pe && p != eof {
			goto _out236;
			
		}
		if p == eof {
			goto _ctr600;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st237;
				
			}
			goto _ctr2;
			
		}
		_ctr602:
		{return st, err }
		goto _st237;
		_st237:
		if p == eof {
			goto _out237;
			
		}
		p+=1;
		st_case_237:
		if p == pe && p != eof {
			goto _out237;
			
		}
		if p == eof {
			goto _ctr602;
			
		} else {
			if ( data[ p ] ) == 98 {
				goto _st238;
				
			}
			goto _ctr2;
			
		}
		_ctr604:
		{return st, err }
		goto _st238;
		_st238:
		if p == eof {
			goto _out238;
			
		}
		p+=1;
		st_case_238:
		if p == pe && p != eof {
			goto _out238;
			
		}
		if p == eof {
			goto _ctr604;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr605;
					
				}
				case 114:
				{
					goto _st239;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr607:
		{return st, err }
		goto _st239;
		_st239:
		if p == eof {
			goto _out239;
			
		}
		p+=1;
		st_case_239:
		if p == pe && p != eof {
			goto _out239;
			
		}
		if p == eof {
			goto _ctr607;
			
		} else {
			if ( data[ p ] ) == 117 {
				goto _st240;
				
			}
			goto _ctr2;
			
		}
		_ctr609:
		{return st, err }
		goto _st240;
		_st240:
		if p == eof {
			goto _out240;
			
		}
		p+=1;
		st_case_240:
		if p == pe && p != eof {
			goto _out240;
			
		}
		if p == eof {
			goto _ctr609;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st241;
				
			}
			goto _ctr2;
			
		}
		_ctr611:
		{return st, err }
		goto _st241;
		_st241:
		if p == eof {
			goto _out241;
			
		}
		p+=1;
		st_case_241:
		if p == pe && p != eof {
			goto _out241;
			
		}
		if p == eof {
			goto _ctr611;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st242;
				
			}
			goto _ctr2;
			
		}
		_ctr613:
		{return st, err }
		goto _st242;
		_st242:
		if p == eof {
			goto _out242;
			
		}
		p+=1;
		st_case_242:
		if p == pe && p != eof {
			goto _out242;
			
		}
		if p == eof {
			goto _ctr613;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st243;
				
			}
			goto _ctr2;
			
		}
		_ctr615:
		{return st, err }
		goto _st243;
		_st243:
		if p == eof {
			goto _out243;
			
		}
		p+=1;
		st_case_243:
		if p == pe && p != eof {
			goto _out243;
			
		}
		if p == eof {
			goto _ctr615;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr605;
				
			}
			goto _ctr2;
			
		}
		_ctr616:
		{return st, err }
		goto _st244;
		_st244:
		if p == eof {
			goto _out244;
			
		}
		p+=1;
		st_case_244:
		if p == pe && p != eof {
			goto _out244;
			
		}
		if p == eof {
			goto _ctr616;
			
		} else {
			switch ( data[ p ] ) {
				case 97:
				{
					goto _st245;
					
				}
				case 117:
				{
					goto _st251;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr619:
		{return st, err }
		goto _st245;
		_st245:
		if p == eof {
			goto _out245;
			
		}
		p+=1;
		st_case_245:
		if p == pe && p != eof {
			goto _out245;
			
		}
		if p == eof {
			goto _ctr619;
			
		} else {
			if ( data[ p ] ) == 110 {
				goto _st246;
				
			}
			goto _ctr2;
			
		}
		_ctr621:
		{return st, err }
		goto _st246;
		_st246:
		if p == eof {
			goto _out246;
			
		}
		p+=1;
		st_case_246:
		if p == pe && p != eof {
			goto _out246;
			
		}
		if p == eof {
			goto _ctr621;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr622;
					
				}
				case 117:
				{
					goto _st247;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr624:
		{return st, err }
		goto _st247;
		_st247:
		if p == eof {
			goto _out247;
			
		}
		p+=1;
		st_case_247:
		if p == pe && p != eof {
			goto _out247;
			
		}
		if p == eof {
			goto _ctr624;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st248;
				
			}
			goto _ctr2;
			
		}
		_ctr626:
		{return st, err }
		goto _st248;
		_st248:
		if p == eof {
			goto _out248;
			
		}
		p+=1;
		st_case_248:
		if p == pe && p != eof {
			goto _out248;
			
		}
		if p == eof {
			goto _ctr626;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st249;
				
			}
			goto _ctr2;
			
		}
		_ctr628:
		{return st, err }
		goto _st249;
		_st249:
		if p == eof {
			goto _out249;
			
		}
		p+=1;
		st_case_249:
		if p == pe && p != eof {
			goto _out249;
			
		}
		if p == eof {
			goto _ctr628;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st250;
				
			}
			goto _ctr2;
			
		}
		_ctr630:
		{return st, err }
		goto _st250;
		_st250:
		if p == eof {
			goto _out250;
			
		}
		p+=1;
		st_case_250:
		if p == pe && p != eof {
			goto _out250;
			
		}
		if p == eof {
			goto _ctr630;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr622;
				
			}
			goto _ctr2;
			
		}
		_ctr631:
		{return st, err }
		goto _st251;
		_st251:
		if p == eof {
			goto _out251;
			
		}
		p+=1;
		st_case_251:
		if p == pe && p != eof {
			goto _out251;
			
		}
		if p == eof {
			goto _ctr631;
			
		} else {
			switch ( data[ p ] ) {
				case 108:
				{
					goto _st252;
					
				}
				case 110:
				{
					goto _st253;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr634:
		{return st, err }
		goto _st252;
		_st252:
		if p == eof {
			goto _out252;
			
		}
		p+=1;
		st_case_252:
		if p == pe && p != eof {
			goto _out252;
			
		}
		if p == eof {
			goto _ctr634;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr635;
				
			}
			goto _ctr2;
			
		}
		_ctr636:
		{return st, err }
		goto _st253;
		_st253:
		if p == eof {
			goto _out253;
			
		}
		p+=1;
		st_case_253:
		if p == pe && p != eof {
			goto _out253;
			
		}
		if p == eof {
			goto _ctr636;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr637;
				
			}
			goto _ctr2;
			
		}
		_ctr638:
		{return st, err }
		goto _st254;
		_st254:
		if p == eof {
			goto _out254;
			
		}
		p+=1;
		st_case_254:
		if p == pe && p != eof {
			goto _out254;
			
		}
		if p == eof {
			goto _ctr638;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st255;
				
			}
			goto _ctr2;
			
		}
		_ctr640:
		{return st, err }
		goto _st255;
		_st255:
		if p == eof {
			goto _out255;
			
		}
		p+=1;
		st_case_255:
		if p == pe && p != eof {
			goto _out255;
			
		}
		if p == eof {
			goto _ctr640;
			
		} else {
			switch ( data[ p ] ) {
				case 114:
				{
					goto _st256;
					
				}
				case 121:
				{
					goto _st259;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr643:
		{return st, err }
		goto _st256;
		_st256:
		if p == eof {
			goto _out256;
			
		}
		p+=1;
		st_case_256:
		if p == pe && p != eof {
			goto _out256;
			
		}
		if p == eof {
			goto _ctr643;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr644;
					
				}
				case 99:
				{
					goto _st257;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr646:
		{return st, err }
		goto _st257;
		_st257:
		if p == eof {
			goto _out257;
			
		}
		p+=1;
		st_case_257:
		if p == pe && p != eof {
			goto _out257;
			
		}
		if p == eof {
			goto _ctr646;
			
		} else {
			if ( data[ p ] ) == 104 {
				goto _st258;
				
			}
			goto _ctr2;
			
		}
		_ctr648:
		{return st, err }
		goto _st258;
		_st258:
		if p == eof {
			goto _out258;
			
		}
		p+=1;
		st_case_258:
		if p == pe && p != eof {
			goto _out258;
			
		}
		if p == eof {
			goto _ctr648;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr644;
				
			}
			goto _ctr2;
			
		}
		_ctr649:
		{return st, err }
		goto _st259;
		_st259:
		if p == eof {
			goto _out259;
			
		}
		p+=1;
		st_case_259:
		if p == pe && p != eof {
			goto _out259;
			
		}
		if p == eof {
			goto _ctr649;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr650;
				
			}
			goto _ctr2;
			
		}
		_ctr651:
		{return st, err }
		goto _st260;
		_st260:
		if p == eof {
			goto _out260;
			
		}
		p+=1;
		st_case_260:
		if p == pe && p != eof {
			goto _out260;
			
		}
		if p == eof {
			goto _ctr651;
			
		} else {
			if ( data[ p ] ) == 111 {
				goto _st261;
				
			}
			goto _ctr2;
			
		}
		_ctr653:
		{return st, err }
		goto _st261;
		_st261:
		if p == eof {
			goto _out261;
			
		}
		p+=1;
		st_case_261:
		if p == pe && p != eof {
			goto _out261;
			
		}
		if p == eof {
			goto _ctr653;
			
		} else {
			if ( data[ p ] ) == 118 {
				goto _st262;
				
			}
			goto _ctr2;
			
		}
		_ctr655:
		{return st, err }
		goto _st262;
		_st262:
		if p == eof {
			goto _out262;
			
		}
		p+=1;
		st_case_262:
		if p == pe && p != eof {
			goto _out262;
			
		}
		if p == eof {
			goto _ctr655;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr656;
				
			}
			goto _ctr2;
			
		}
		_ctr657:
		{return st, err }
		goto _st263;
		_st263:
		if p == eof {
			goto _out263;
			
		}
		p+=1;
		st_case_263:
		if p == pe && p != eof {
			goto _out263;
			
		}
		if p == eof {
			goto _ctr657;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st264;
				
			}
			goto _ctr2;
			
		}
		_ctr659:
		{return st, err }
		goto _st264;
		_st264:
		if p == eof {
			goto _out264;
			
		}
		p+=1;
		st_case_264:
		if p == pe && p != eof {
			goto _out264;
			
		}
		if p == eof {
			goto _ctr659;
			
		} else {
			if ( data[ p ] ) == 116 {
				goto _st265;
				
			}
			goto _ctr2;
			
		}
		_ctr661:
		{return st, err }
		goto _st265;
		_st265:
		if p == eof {
			goto _out265;
			
		}
		p+=1;
		st_case_265:
		if p == pe && p != eof {
			goto _out265;
			
		}
		if p == eof {
			goto _ctr661;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr662;
				
			}
			goto _ctr2;
			
		}
		_ctr663:
		{return st, err }
		goto _st266;
		_st266:
		if p == eof {
			goto _out266;
			
		}
		p+=1;
		st_case_266:
		if p == pe && p != eof {
			goto _out266;
			
		}
		if p == eof {
			goto _ctr663;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st267;
				
			}
			goto _ctr2;
			
		}
		_ctr665:
		{return st, err }
		goto _st267;
		_st267:
		if p == eof {
			goto _out267;
			
		}
		p+=1;
		st_case_267:
		if p == pe && p != eof {
			goto _out267;
			
		}
		if p == eof {
			goto _ctr665;
			
		} else {
			if ( data[ p ] ) == 112 {
				goto _st268;
				
			}
			goto _ctr2;
			
		}
		_ctr667:
		{return st, err }
		goto _st268;
		_st268:
		if p == eof {
			goto _out268;
			
		}
		p+=1;
		st_case_268:
		if p == pe && p != eof {
			goto _out268;
			
		}
		if p == eof {
			goto _ctr667;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr668;
				
			}
			goto _ctr2;
			
		}
		_ctr669:
		{return st, err }
		goto _st269;
		_st269:
		if p == eof {
			goto _out269;
			
		}
		p+=1;
		st_case_269:
		if p == pe && p != eof {
			goto _out269;
			
		}
		if p == eof {
			goto _ctr669;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _st270;
				
			}
			goto _ctr2;
			
		}
		_ctr671:
		{return st, err }
		goto _st270;
		_st270:
		if p == eof {
			goto _out270;
			
		}
		p+=1;
		st_case_270:
		if p == pe && p != eof {
			goto _out270;
			
		}
		if p == eof {
			goto _ctr671;
			
		} else {
			switch ( data[ p ] ) {
				case 48:
				{
					goto _ctr672;
					
				}
				case 51:
				{
					goto _ctr674;
					
				}
				
			}
			if ( data[ p ] ) > 50 {
				if 52 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
					goto _ctr675;
					
				}
				
			} else if ( data[ p ] ) >= 49 {
				goto _ctr673;
				
			}
			goto _ctr2;
			
		}
		_ctr676:
		{return st, err }
		goto _st271;
		_ctr672:
		{pb = p }
		goto _st271;
		_st271:
		if p == eof {
			goto _out271;
			
		}
		p+=1;
		st_case_271:
		if p == pe && p != eof {
			goto _out271;
			
		}
		if p == eof {
			goto _ctr676;
			
		} else {
			if 49 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st272;
				
			}
			goto _ctr2;
			
		}
		_ctr678:
		{return st, err }
		goto _st272;
		_ctr675:
		{pb = p }
		goto _st272;
		_st272:
		if p == eof {
			goto _out272;
			
		}
		p+=1;
		st_case_272:
		if p == pe && p != eof {
			goto _out272;
			
		}
		if p == eof {
			goto _ctr678;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr679;
					
				}
				case 45:
				{
					goto _ctr680;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr681:
		{return st, err }
		goto _st273;
		_ctr679:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st273;
		_st273:
		if p == eof {
			goto _out273;
			
		}
		p+=1;
		st_case_273:
		if p == pe && p != eof {
			goto _out273;
			
		}
		if p == eof {
			goto _ctr681;
			
		} else {
			switch ( data[ p ] ) {
				case 65:
				{
					goto _st274;
					
				}
				case 68:
				{
					goto _st281;
					
				}
				case 70:
				{
					goto _st284;
					
				}
				case 74:
				{
					goto _st292;
					
				}
				case 77:
				{
					goto _st302;
					
				}
				case 78:
				{
					goto _st308;
					
				}
				case 79:
				{
					goto _st311;
					
				}
				case 83:
				{
					goto _st314;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr690:
		{return st, err }
		goto _st274;
		_st274:
		if p == eof {
			goto _out274;
			
		}
		p+=1;
		st_case_274:
		if p == pe && p != eof {
			goto _out274;
			
		}
		if p == eof {
			goto _ctr690;
			
		} else {
			switch ( data[ p ] ) {
				case 112:
				{
					goto _st275;
					
				}
				case 117:
				{
					goto _st279;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr693:
		{return st, err }
		goto _st275;
		_st275:
		if p == eof {
			goto _out275;
			
		}
		p+=1;
		st_case_275:
		if p == pe && p != eof {
			goto _out275;
			
		}
		if p == eof {
			goto _ctr693;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st276;
				
			}
			goto _ctr2;
			
		}
		_ctr695:
		{return st, err }
		goto _st276;
		_st276:
		if p == eof {
			goto _out276;
			
		}
		p+=1;
		st_case_276:
		if p == pe && p != eof {
			goto _out276;
			
		}
		if p == eof {
			goto _ctr695;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr696;
					
				}
				case 105:
				{
					goto _st277;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr698:
		{return st, err }
		goto _st277;
		_st277:
		if p == eof {
			goto _out277;
			
		}
		p+=1;
		st_case_277:
		if p == pe && p != eof {
			goto _out277;
			
		}
		if p == eof {
			goto _ctr698;
			
		} else {
			if ( data[ p ] ) == 108 {
				goto _st278;
				
			}
			goto _ctr2;
			
		}
		_ctr700:
		{return st, err }
		goto _st278;
		_st278:
		if p == eof {
			goto _out278;
			
		}
		p+=1;
		st_case_278:
		if p == pe && p != eof {
			goto _out278;
			
		}
		if p == eof {
			goto _ctr700;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr696;
				
			}
			goto _ctr2;
			
		}
		_ctr701:
		{return st, err }
		goto _st279;
		_st279:
		if p == eof {
			goto _out279;
			
		}
		p+=1;
		st_case_279:
		if p == pe && p != eof {
			goto _out279;
			
		}
		if p == eof {
			goto _ctr701;
			
		} else {
			if ( data[ p ] ) == 103 {
				goto _st280;
				
			}
			goto _ctr2;
			
		}
		_ctr703:
		{return st, err }
		goto _st280;
		_st280:
		if p == eof {
			goto _out280;
			
		}
		p+=1;
		st_case_280:
		if p == pe && p != eof {
			goto _out280;
			
		}
		if p == eof {
			goto _ctr703;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr704;
				
			}
			goto _ctr2;
			
		}
		_ctr705:
		{return st, err }
		goto _st281;
		_st281:
		if p == eof {
			goto _out281;
			
		}
		p+=1;
		st_case_281:
		if p == pe && p != eof {
			goto _out281;
			
		}
		if p == eof {
			goto _ctr705;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st282;
				
			}
			goto _ctr2;
			
		}
		_ctr707:
		{return st, err }
		goto _st282;
		_st282:
		if p == eof {
			goto _out282;
			
		}
		p+=1;
		st_case_282:
		if p == pe && p != eof {
			goto _out282;
			
		}
		if p == eof {
			goto _ctr707;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st283;
				
			}
			goto _ctr2;
			
		}
		_ctr709:
		{return st, err }
		goto _st283;
		_st283:
		if p == eof {
			goto _out283;
			
		}
		p+=1;
		st_case_283:
		if p == pe && p != eof {
			goto _out283;
			
		}
		if p == eof {
			goto _ctr709;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr710;
				
			}
			goto _ctr2;
			
		}
		_ctr711:
		{return st, err }
		goto _st284;
		_st284:
		if p == eof {
			goto _out284;
			
		}
		p+=1;
		st_case_284:
		if p == pe && p != eof {
			goto _out284;
			
		}
		if p == eof {
			goto _ctr711;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st285;
				
			}
			goto _ctr2;
			
		}
		_ctr713:
		{return st, err }
		goto _st285;
		_st285:
		if p == eof {
			goto _out285;
			
		}
		p+=1;
		st_case_285:
		if p == pe && p != eof {
			goto _out285;
			
		}
		if p == eof {
			goto _ctr713;
			
		} else {
			if ( data[ p ] ) == 98 {
				goto _st286;
				
			}
			goto _ctr2;
			
		}
		_ctr715:
		{return st, err }
		goto _st286;
		_st286:
		if p == eof {
			goto _out286;
			
		}
		p+=1;
		st_case_286:
		if p == pe && p != eof {
			goto _out286;
			
		}
		if p == eof {
			goto _ctr715;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr716;
					
				}
				case 114:
				{
					goto _st287;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr718:
		{return st, err }
		goto _st287;
		_st287:
		if p == eof {
			goto _out287;
			
		}
		p+=1;
		st_case_287:
		if p == pe && p != eof {
			goto _out287;
			
		}
		if p == eof {
			goto _ctr718;
			
		} else {
			if ( data[ p ] ) == 117 {
				goto _st288;
				
			}
			goto _ctr2;
			
		}
		_ctr720:
		{return st, err }
		goto _st288;
		_st288:
		if p == eof {
			goto _out288;
			
		}
		p+=1;
		st_case_288:
		if p == pe && p != eof {
			goto _out288;
			
		}
		if p == eof {
			goto _ctr720;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st289;
				
			}
			goto _ctr2;
			
		}
		_ctr722:
		{return st, err }
		goto _st289;
		_st289:
		if p == eof {
			goto _out289;
			
		}
		p+=1;
		st_case_289:
		if p == pe && p != eof {
			goto _out289;
			
		}
		if p == eof {
			goto _ctr722;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st290;
				
			}
			goto _ctr2;
			
		}
		_ctr724:
		{return st, err }
		goto _st290;
		_st290:
		if p == eof {
			goto _out290;
			
		}
		p+=1;
		st_case_290:
		if p == pe && p != eof {
			goto _out290;
			
		}
		if p == eof {
			goto _ctr724;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st291;
				
			}
			goto _ctr2;
			
		}
		_ctr726:
		{return st, err }
		goto _st291;
		_st291:
		if p == eof {
			goto _out291;
			
		}
		p+=1;
		st_case_291:
		if p == pe && p != eof {
			goto _out291;
			
		}
		if p == eof {
			goto _ctr726;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr716;
				
			}
			goto _ctr2;
			
		}
		_ctr727:
		{return st, err }
		goto _st292;
		_st292:
		if p == eof {
			goto _out292;
			
		}
		p+=1;
		st_case_292:
		if p == pe && p != eof {
			goto _out292;
			
		}
		if p == eof {
			goto _ctr727;
			
		} else {
			switch ( data[ p ] ) {
				case 97:
				{
					goto _st293;
					
				}
				case 117:
				{
					goto _st299;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr730:
		{return st, err }
		goto _st293;
		_st293:
		if p == eof {
			goto _out293;
			
		}
		p+=1;
		st_case_293:
		if p == pe && p != eof {
			goto _out293;
			
		}
		if p == eof {
			goto _ctr730;
			
		} else {
			if ( data[ p ] ) == 110 {
				goto _st294;
				
			}
			goto _ctr2;
			
		}
		_ctr732:
		{return st, err }
		goto _st294;
		_st294:
		if p == eof {
			goto _out294;
			
		}
		p+=1;
		st_case_294:
		if p == pe && p != eof {
			goto _out294;
			
		}
		if p == eof {
			goto _ctr732;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr733;
					
				}
				case 117:
				{
					goto _st295;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr735:
		{return st, err }
		goto _st295;
		_st295:
		if p == eof {
			goto _out295;
			
		}
		p+=1;
		st_case_295:
		if p == pe && p != eof {
			goto _out295;
			
		}
		if p == eof {
			goto _ctr735;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st296;
				
			}
			goto _ctr2;
			
		}
		_ctr737:
		{return st, err }
		goto _st296;
		_st296:
		if p == eof {
			goto _out296;
			
		}
		p+=1;
		st_case_296:
		if p == pe && p != eof {
			goto _out296;
			
		}
		if p == eof {
			goto _ctr737;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st297;
				
			}
			goto _ctr2;
			
		}
		_ctr739:
		{return st, err }
		goto _st297;
		_st297:
		if p == eof {
			goto _out297;
			
		}
		p+=1;
		st_case_297:
		if p == pe && p != eof {
			goto _out297;
			
		}
		if p == eof {
			goto _ctr739;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st298;
				
			}
			goto _ctr2;
			
		}
		_ctr741:
		{return st, err }
		goto _st298;
		_st298:
		if p == eof {
			goto _out298;
			
		}
		p+=1;
		st_case_298:
		if p == pe && p != eof {
			goto _out298;
			
		}
		if p == eof {
			goto _ctr741;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr733;
				
			}
			goto _ctr2;
			
		}
		_ctr742:
		{return st, err }
		goto _st299;
		_st299:
		if p == eof {
			goto _out299;
			
		}
		p+=1;
		st_case_299:
		if p == pe && p != eof {
			goto _out299;
			
		}
		if p == eof {
			goto _ctr742;
			
		} else {
			switch ( data[ p ] ) {
				case 108:
				{
					goto _st300;
					
				}
				case 110:
				{
					goto _st301;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr745:
		{return st, err }
		goto _st300;
		_st300:
		if p == eof {
			goto _out300;
			
		}
		p+=1;
		st_case_300:
		if p == pe && p != eof {
			goto _out300;
			
		}
		if p == eof {
			goto _ctr745;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr746;
				
			}
			goto _ctr2;
			
		}
		_ctr747:
		{return st, err }
		goto _st301;
		_st301:
		if p == eof {
			goto _out301;
			
		}
		p+=1;
		st_case_301:
		if p == pe && p != eof {
			goto _out301;
			
		}
		if p == eof {
			goto _ctr747;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr748;
				
			}
			goto _ctr2;
			
		}
		_ctr749:
		{return st, err }
		goto _st302;
		_st302:
		if p == eof {
			goto _out302;
			
		}
		p+=1;
		st_case_302:
		if p == pe && p != eof {
			goto _out302;
			
		}
		if p == eof {
			goto _ctr749;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st303;
				
			}
			goto _ctr2;
			
		}
		_ctr751:
		{return st, err }
		goto _st303;
		_st303:
		if p == eof {
			goto _out303;
			
		}
		p+=1;
		st_case_303:
		if p == pe && p != eof {
			goto _out303;
			
		}
		if p == eof {
			goto _ctr751;
			
		} else {
			switch ( data[ p ] ) {
				case 114:
				{
					goto _st304;
					
				}
				case 121:
				{
					goto _st307;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr754:
		{return st, err }
		goto _st304;
		_st304:
		if p == eof {
			goto _out304;
			
		}
		p+=1;
		st_case_304:
		if p == pe && p != eof {
			goto _out304;
			
		}
		if p == eof {
			goto _ctr754;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr755;
					
				}
				case 99:
				{
					goto _st305;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr757:
		{return st, err }
		goto _st305;
		_st305:
		if p == eof {
			goto _out305;
			
		}
		p+=1;
		st_case_305:
		if p == pe && p != eof {
			goto _out305;
			
		}
		if p == eof {
			goto _ctr757;
			
		} else {
			if ( data[ p ] ) == 104 {
				goto _st306;
				
			}
			goto _ctr2;
			
		}
		_ctr759:
		{return st, err }
		goto _st306;
		_st306:
		if p == eof {
			goto _out306;
			
		}
		p+=1;
		st_case_306:
		if p == pe && p != eof {
			goto _out306;
			
		}
		if p == eof {
			goto _ctr759;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr755;
				
			}
			goto _ctr2;
			
		}
		_ctr760:
		{return st, err }
		goto _st307;
		_st307:
		if p == eof {
			goto _out307;
			
		}
		p+=1;
		st_case_307:
		if p == pe && p != eof {
			goto _out307;
			
		}
		if p == eof {
			goto _ctr760;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr761;
				
			}
			goto _ctr2;
			
		}
		_ctr762:
		{return st, err }
		goto _st308;
		_st308:
		if p == eof {
			goto _out308;
			
		}
		p+=1;
		st_case_308:
		if p == pe && p != eof {
			goto _out308;
			
		}
		if p == eof {
			goto _ctr762;
			
		} else {
			if ( data[ p ] ) == 111 {
				goto _st309;
				
			}
			goto _ctr2;
			
		}
		_ctr764:
		{return st, err }
		goto _st309;
		_st309:
		if p == eof {
			goto _out309;
			
		}
		p+=1;
		st_case_309:
		if p == pe && p != eof {
			goto _out309;
			
		}
		if p == eof {
			goto _ctr764;
			
		} else {
			if ( data[ p ] ) == 118 {
				goto _st310;
				
			}
			goto _ctr2;
			
		}
		_ctr766:
		{return st, err }
		goto _st310;
		_st310:
		if p == eof {
			goto _out310;
			
		}
		p+=1;
		st_case_310:
		if p == pe && p != eof {
			goto _out310;
			
		}
		if p == eof {
			goto _ctr766;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr767;
				
			}
			goto _ctr2;
			
		}
		_ctr768:
		{return st, err }
		goto _st311;
		_st311:
		if p == eof {
			goto _out311;
			
		}
		p+=1;
		st_case_311:
		if p == pe && p != eof {
			goto _out311;
			
		}
		if p == eof {
			goto _ctr768;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st312;
				
			}
			goto _ctr2;
			
		}
		_ctr770:
		{return st, err }
		goto _st312;
		_st312:
		if p == eof {
			goto _out312;
			
		}
		p+=1;
		st_case_312:
		if p == pe && p != eof {
			goto _out312;
			
		}
		if p == eof {
			goto _ctr770;
			
		} else {
			if ( data[ p ] ) == 116 {
				goto _st313;
				
			}
			goto _ctr2;
			
		}
		_ctr772:
		{return st, err }
		goto _st313;
		_st313:
		if p == eof {
			goto _out313;
			
		}
		p+=1;
		st_case_313:
		if p == pe && p != eof {
			goto _out313;
			
		}
		if p == eof {
			goto _ctr772;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr773;
				
			}
			goto _ctr2;
			
		}
		_ctr774:
		{return st, err }
		goto _st314;
		_st314:
		if p == eof {
			goto _out314;
			
		}
		p+=1;
		st_case_314:
		if p == pe && p != eof {
			goto _out314;
			
		}
		if p == eof {
			goto _ctr774;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st315;
				
			}
			goto _ctr2;
			
		}
		_ctr776:
		{return st, err }
		goto _st315;
		_st315:
		if p == eof {
			goto _out315;
			
		}
		p+=1;
		st_case_315:
		if p == pe && p != eof {
			goto _out315;
			
		}
		if p == eof {
			goto _ctr776;
			
		} else {
			if ( data[ p ] ) == 112 {
				goto _st316;
				
			}
			goto _ctr2;
			
		}
		_ctr778:
		{return st, err }
		goto _st316;
		_st316:
		if p == eof {
			goto _out316;
			
		}
		p+=1;
		st_case_316:
		if p == pe && p != eof {
			goto _out316;
			
		}
		if p == eof {
			goto _ctr778;
			
		} else {
			if ( data[ p ] ) == 32 {
				goto _ctr779;
				
			}
			goto _ctr2;
			
		}
		_ctr780:
		{return st, err }
		goto _st317;
		_ctr680:
		{switch p - pb {
				case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
				default:
				err = fmt.Errorf("invalid day digits %s", data[pb:p])
				return
			}
		}
		goto _st317;
		_st317:
		if p == eof {
			goto _out317;
			
		}
		p+=1;
		st_case_317:
		if p == pe && p != eof {
			goto _out317;
			
		}
		if p == eof {
			goto _ctr780;
			
		} else {
			switch ( data[ p ] ) {
				case 65:
				{
					goto _st318;
					
				}
				case 68:
				{
					goto _st325;
					
				}
				case 70:
				{
					goto _st328;
					
				}
				case 74:
				{
					goto _st336;
					
				}
				case 77:
				{
					goto _st346;
					
				}
				case 78:
				{
					goto _st352;
					
				}
				case 79:
				{
					goto _st355;
					
				}
				case 83:
				{
					goto _st358;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr789:
		{return st, err }
		goto _st318;
		_st318:
		if p == eof {
			goto _out318;
			
		}
		p+=1;
		st_case_318:
		if p == pe && p != eof {
			goto _out318;
			
		}
		if p == eof {
			goto _ctr789;
			
		} else {
			switch ( data[ p ] ) {
				case 112:
				{
					goto _st319;
					
				}
				case 117:
				{
					goto _st323;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr792:
		{return st, err }
		goto _st319;
		_st319:
		if p == eof {
			goto _out319;
			
		}
		p+=1;
		st_case_319:
		if p == pe && p != eof {
			goto _out319;
			
		}
		if p == eof {
			goto _ctr792;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st320;
				
			}
			goto _ctr2;
			
		}
		_ctr794:
		{return st, err }
		goto _st320;
		_st320:
		if p == eof {
			goto _out320;
			
		}
		p+=1;
		st_case_320:
		if p == pe && p != eof {
			goto _out320;
			
		}
		if p == eof {
			goto _ctr794;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr217;
					
				}
				case 105:
				{
					goto _st321;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr796:
		{return st, err }
		goto _st321;
		_st321:
		if p == eof {
			goto _out321;
			
		}
		p+=1;
		st_case_321:
		if p == pe && p != eof {
			goto _out321;
			
		}
		if p == eof {
			goto _ctr796;
			
		} else {
			if ( data[ p ] ) == 108 {
				goto _st322;
				
			}
			goto _ctr2;
			
		}
		_ctr798:
		{return st, err }
		goto _st322;
		_st322:
		if p == eof {
			goto _out322;
			
		}
		p+=1;
		st_case_322:
		if p == pe && p != eof {
			goto _out322;
			
		}
		if p == eof {
			goto _ctr798;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr217;
				
			}
			goto _ctr2;
			
		}
		_ctr799:
		{return st, err }
		goto _st323;
		_st323:
		if p == eof {
			goto _out323;
			
		}
		p+=1;
		st_case_323:
		if p == pe && p != eof {
			goto _out323;
			
		}
		if p == eof {
			goto _ctr799;
			
		} else {
			if ( data[ p ] ) == 103 {
				goto _st324;
				
			}
			goto _ctr2;
			
		}
		_ctr801:
		{return st, err }
		goto _st324;
		_st324:
		if p == eof {
			goto _out324;
			
		}
		p+=1;
		st_case_324:
		if p == pe && p != eof {
			goto _out324;
			
		}
		if p == eof {
			goto _ctr801;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr229;
				
			}
			goto _ctr2;
			
		}
		_ctr802:
		{return st, err }
		goto _st325;
		_st325:
		if p == eof {
			goto _out325;
			
		}
		p+=1;
		st_case_325:
		if p == pe && p != eof {
			goto _out325;
			
		}
		if p == eof {
			goto _ctr802;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st326;
				
			}
			goto _ctr2;
			
		}
		_ctr804:
		{return st, err }
		goto _st326;
		_st326:
		if p == eof {
			goto _out326;
			
		}
		p+=1;
		st_case_326:
		if p == pe && p != eof {
			goto _out326;
			
		}
		if p == eof {
			goto _ctr804;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st327;
				
			}
			goto _ctr2;
			
		}
		_ctr806:
		{return st, err }
		goto _st327;
		_st327:
		if p == eof {
			goto _out327;
			
		}
		p+=1;
		st_case_327:
		if p == pe && p != eof {
			goto _out327;
			
		}
		if p == eof {
			goto _ctr806;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr235;
				
			}
			goto _ctr2;
			
		}
		_ctr807:
		{return st, err }
		goto _st328;
		_st328:
		if p == eof {
			goto _out328;
			
		}
		p+=1;
		st_case_328:
		if p == pe && p != eof {
			goto _out328;
			
		}
		if p == eof {
			goto _ctr807;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st329;
				
			}
			goto _ctr2;
			
		}
		_ctr809:
		{return st, err }
		goto _st329;
		_st329:
		if p == eof {
			goto _out329;
			
		}
		p+=1;
		st_case_329:
		if p == pe && p != eof {
			goto _out329;
			
		}
		if p == eof {
			goto _ctr809;
			
		} else {
			if ( data[ p ] ) == 98 {
				goto _st330;
				
			}
			goto _ctr2;
			
		}
		_ctr811:
		{return st, err }
		goto _st330;
		_st330:
		if p == eof {
			goto _out330;
			
		}
		p+=1;
		st_case_330:
		if p == pe && p != eof {
			goto _out330;
			
		}
		if p == eof {
			goto _ctr811;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr241;
					
				}
				case 114:
				{
					goto _st331;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr813:
		{return st, err }
		goto _st331;
		_st331:
		if p == eof {
			goto _out331;
			
		}
		p+=1;
		st_case_331:
		if p == pe && p != eof {
			goto _out331;
			
		}
		if p == eof {
			goto _ctr813;
			
		} else {
			if ( data[ p ] ) == 117 {
				goto _st332;
				
			}
			goto _ctr2;
			
		}
		_ctr815:
		{return st, err }
		goto _st332;
		_st332:
		if p == eof {
			goto _out332;
			
		}
		p+=1;
		st_case_332:
		if p == pe && p != eof {
			goto _out332;
			
		}
		if p == eof {
			goto _ctr815;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st333;
				
			}
			goto _ctr2;
			
		}
		_ctr817:
		{return st, err }
		goto _st333;
		_st333:
		if p == eof {
			goto _out333;
			
		}
		p+=1;
		st_case_333:
		if p == pe && p != eof {
			goto _out333;
			
		}
		if p == eof {
			goto _ctr817;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st334;
				
			}
			goto _ctr2;
			
		}
		_ctr819:
		{return st, err }
		goto _st334;
		_st334:
		if p == eof {
			goto _out334;
			
		}
		p+=1;
		st_case_334:
		if p == pe && p != eof {
			goto _out334;
			
		}
		if p == eof {
			goto _ctr819;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st335;
				
			}
			goto _ctr2;
			
		}
		_ctr821:
		{return st, err }
		goto _st335;
		_st335:
		if p == eof {
			goto _out335;
			
		}
		p+=1;
		st_case_335:
		if p == pe && p != eof {
			goto _out335;
			
		}
		if p == eof {
			goto _ctr821;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr241;
				
			}
			goto _ctr2;
			
		}
		_ctr822:
		{return st, err }
		goto _st336;
		_st336:
		if p == eof {
			goto _out336;
			
		}
		p+=1;
		st_case_336:
		if p == pe && p != eof {
			goto _out336;
			
		}
		if p == eof {
			goto _ctr822;
			
		} else {
			switch ( data[ p ] ) {
				case 97:
				{
					goto _st337;
					
				}
				case 117:
				{
					goto _st343;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr825:
		{return st, err }
		goto _st337;
		_st337:
		if p == eof {
			goto _out337;
			
		}
		p+=1;
		st_case_337:
		if p == pe && p != eof {
			goto _out337;
			
		}
		if p == eof {
			goto _ctr825;
			
		} else {
			if ( data[ p ] ) == 110 {
				goto _st338;
				
			}
			goto _ctr2;
			
		}
		_ctr827:
		{return st, err }
		goto _st338;
		_st338:
		if p == eof {
			goto _out338;
			
		}
		p+=1;
		st_case_338:
		if p == pe && p != eof {
			goto _out338;
			
		}
		if p == eof {
			goto _ctr827;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr258;
					
				}
				case 117:
				{
					goto _st339;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr829:
		{return st, err }
		goto _st339;
		_st339:
		if p == eof {
			goto _out339;
			
		}
		p+=1;
		st_case_339:
		if p == pe && p != eof {
			goto _out339;
			
		}
		if p == eof {
			goto _ctr829;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st340;
				
			}
			goto _ctr2;
			
		}
		_ctr831:
		{return st, err }
		goto _st340;
		_st340:
		if p == eof {
			goto _out340;
			
		}
		p+=1;
		st_case_340:
		if p == pe && p != eof {
			goto _out340;
			
		}
		if p == eof {
			goto _ctr831;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st341;
				
			}
			goto _ctr2;
			
		}
		_ctr833:
		{return st, err }
		goto _st341;
		_st341:
		if p == eof {
			goto _out341;
			
		}
		p+=1;
		st_case_341:
		if p == pe && p != eof {
			goto _out341;
			
		}
		if p == eof {
			goto _ctr833;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st342;
				
			}
			goto _ctr2;
			
		}
		_ctr835:
		{return st, err }
		goto _st342;
		_st342:
		if p == eof {
			goto _out342;
			
		}
		p+=1;
		st_case_342:
		if p == pe && p != eof {
			goto _out342;
			
		}
		if p == eof {
			goto _ctr835;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr258;
				
			}
			goto _ctr2;
			
		}
		_ctr836:
		{return st, err }
		goto _st343;
		_st343:
		if p == eof {
			goto _out343;
			
		}
		p+=1;
		st_case_343:
		if p == pe && p != eof {
			goto _out343;
			
		}
		if p == eof {
			goto _ctr836;
			
		} else {
			switch ( data[ p ] ) {
				case 108:
				{
					goto _st344;
					
				}
				case 110:
				{
					goto _st345;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr839:
		{return st, err }
		goto _st344;
		_st344:
		if p == eof {
			goto _out344;
			
		}
		p+=1;
		st_case_344:
		if p == pe && p != eof {
			goto _out344;
			
		}
		if p == eof {
			goto _ctr839;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr271;
				
			}
			goto _ctr2;
			
		}
		_ctr840:
		{return st, err }
		goto _st345;
		_st345:
		if p == eof {
			goto _out345;
			
		}
		p+=1;
		st_case_345:
		if p == pe && p != eof {
			goto _out345;
			
		}
		if p == eof {
			goto _ctr840;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr273;
				
			}
			goto _ctr2;
			
		}
		_ctr841:
		{return st, err }
		goto _st346;
		_st346:
		if p == eof {
			goto _out346;
			
		}
		p+=1;
		st_case_346:
		if p == pe && p != eof {
			goto _out346;
			
		}
		if p == eof {
			goto _ctr841;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st347;
				
			}
			goto _ctr2;
			
		}
		_ctr843:
		{return st, err }
		goto _st347;
		_st347:
		if p == eof {
			goto _out347;
			
		}
		p+=1;
		st_case_347:
		if p == pe && p != eof {
			goto _out347;
			
		}
		if p == eof {
			goto _ctr843;
			
		} else {
			switch ( data[ p ] ) {
				case 114:
				{
					goto _st348;
					
				}
				case 121:
				{
					goto _st351;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr846:
		{return st, err }
		goto _st348;
		_st348:
		if p == eof {
			goto _out348;
			
		}
		p+=1;
		st_case_348:
		if p == pe && p != eof {
			goto _out348;
			
		}
		if p == eof {
			goto _ctr846;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr280;
					
				}
				case 99:
				{
					goto _st349;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr848:
		{return st, err }
		goto _st349;
		_st349:
		if p == eof {
			goto _out349;
			
		}
		p+=1;
		st_case_349:
		if p == pe && p != eof {
			goto _out349;
			
		}
		if p == eof {
			goto _ctr848;
			
		} else {
			if ( data[ p ] ) == 104 {
				goto _st350;
				
			}
			goto _ctr2;
			
		}
		_ctr850:
		{return st, err }
		goto _st350;
		_st350:
		if p == eof {
			goto _out350;
			
		}
		p+=1;
		st_case_350:
		if p == pe && p != eof {
			goto _out350;
			
		}
		if p == eof {
			goto _ctr850;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr280;
				
			}
			goto _ctr2;
			
		}
		_ctr851:
		{return st, err }
		goto _st351;
		_st351:
		if p == eof {
			goto _out351;
			
		}
		p+=1;
		st_case_351:
		if p == pe && p != eof {
			goto _out351;
			
		}
		if p == eof {
			goto _ctr851;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr286;
				
			}
			goto _ctr2;
			
		}
		_ctr852:
		{return st, err }
		goto _st352;
		_st352:
		if p == eof {
			goto _out352;
			
		}
		p+=1;
		st_case_352:
		if p == pe && p != eof {
			goto _out352;
			
		}
		if p == eof {
			goto _ctr852;
			
		} else {
			if ( data[ p ] ) == 111 {
				goto _st353;
				
			}
			goto _ctr2;
			
		}
		_ctr854:
		{return st, err }
		goto _st353;
		_st353:
		if p == eof {
			goto _out353;
			
		}
		p+=1;
		st_case_353:
		if p == pe && p != eof {
			goto _out353;
			
		}
		if p == eof {
			goto _ctr854;
			
		} else {
			if ( data[ p ] ) == 118 {
				goto _st354;
				
			}
			goto _ctr2;
			
		}
		_ctr856:
		{return st, err }
		goto _st354;
		_st354:
		if p == eof {
			goto _out354;
			
		}
		p+=1;
		st_case_354:
		if p == pe && p != eof {
			goto _out354;
			
		}
		if p == eof {
			goto _ctr856;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr292;
				
			}
			goto _ctr2;
			
		}
		_ctr857:
		{return st, err }
		goto _st355;
		_st355:
		if p == eof {
			goto _out355;
			
		}
		p+=1;
		st_case_355:
		if p == pe && p != eof {
			goto _out355;
			
		}
		if p == eof {
			goto _ctr857;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st356;
				
			}
			goto _ctr2;
			
		}
		_ctr859:
		{return st, err }
		goto _st356;
		_st356:
		if p == eof {
			goto _out356;
			
		}
		p+=1;
		st_case_356:
		if p == pe && p != eof {
			goto _out356;
			
		}
		if p == eof {
			goto _ctr859;
			
		} else {
			if ( data[ p ] ) == 116 {
				goto _st357;
				
			}
			goto _ctr2;
			
		}
		_ctr861:
		{return st, err }
		goto _st357;
		_st357:
		if p == eof {
			goto _out357;
			
		}
		p+=1;
		st_case_357:
		if p == pe && p != eof {
			goto _out357;
			
		}
		if p == eof {
			goto _ctr861;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr298;
				
			}
			goto _ctr2;
			
		}
		_ctr862:
		{return st, err }
		goto _st358;
		_st358:
		if p == eof {
			goto _out358;
			
		}
		p+=1;
		st_case_358:
		if p == pe && p != eof {
			goto _out358;
			
		}
		if p == eof {
			goto _ctr862;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st359;
				
			}
			goto _ctr2;
			
		}
		_ctr864:
		{return st, err }
		goto _st359;
		_st359:
		if p == eof {
			goto _out359;
			
		}
		p+=1;
		st_case_359:
		if p == pe && p != eof {
			goto _out359;
			
		}
		if p == eof {
			goto _ctr864;
			
		} else {
			if ( data[ p ] ) == 112 {
				goto _st360;
				
			}
			goto _ctr2;
			
		}
		_ctr866:
		{return st, err }
		goto _st360;
		_st360:
		if p == eof {
			goto _out360;
			
		}
		p+=1;
		st_case_360:
		if p == pe && p != eof {
			goto _out360;
			
		}
		if p == eof {
			goto _ctr866;
			
		} else {
			if ( data[ p ] ) == 45 {
				goto _ctr304;
				
			}
			goto _ctr2;
			
		}
		_ctr867:
		{return st, err }
		goto _st361;
		_ctr673:
		{pb = p }
		goto _st361;
		_st361:
		if p == eof {
			goto _out361;
			
		}
		p+=1;
		st_case_361:
		if p == pe && p != eof {
			goto _out361;
			
		}
		if p == eof {
			goto _ctr867;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr679;
					
				}
				case 45:
				{
					goto _ctr680;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 57 {
				goto _st272;
				
			}
			goto _ctr2;
			
		}
		_ctr868:
		{return st, err }
		goto _st362;
		_ctr674:
		{pb = p }
		goto _st362;
		_st362:
		if p == eof {
			goto _out362;
			
		}
		p+=1;
		st_case_362:
		if p == pe && p != eof {
			goto _out362;
			
		}
		if p == eof {
			goto _ctr868;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _ctr679;
					
				}
				case 45:
				{
					goto _ctr680;
					
				}
				
			}
			if 48 <= ( data[ p ] ) && ( data[ p ] ) <= 49 {
				goto _st272;
				
			}
			goto _ctr2;
			
		}
		_ctr869:
		{return st, err }
		goto _st363;
		_st363:
		if p == eof {
			goto _out363;
			
		}
		p+=1;
		st_case_363:
		if p == pe && p != eof {
			goto _out363;
			
		}
		if p == eof {
			goto _ctr869;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st364;
				
			}
			goto _ctr2;
			
		}
		_ctr871:
		{return st, err }
		goto _st364;
		_st364:
		if p == eof {
			goto _out364;
			
		}
		p+=1;
		st_case_364:
		if p == pe && p != eof {
			goto _out364;
			
		}
		if p == eof {
			goto _ctr871;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st365;
				
			}
			goto _ctr2;
			
		}
		_ctr873:
		{return st, err }
		goto _st365;
		_st365:
		if p == eof {
			goto _out365;
			
		}
		p+=1;
		st_case_365:
		if p == pe && p != eof {
			goto _out365;
			
		}
		if p == eof {
			goto _ctr873;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st160;
					
				}
				case 44:
				{
					goto _st269;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr874:
		{return st, err }
		goto _st366;
		_st366:
		if p == eof {
			goto _out366;
			
		}
		p+=1;
		st_case_366:
		if p == pe && p != eof {
			goto _out366;
			
		}
		if p == eof {
			goto _ctr874;
			
		} else {
			switch ( data[ p ] ) {
				case 97:
				{
					goto _st367;
					
				}
				case 117:
				{
					goto _st373;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr877:
		{return st, err }
		goto _st367;
		_st367:
		if p == eof {
			goto _out367;
			
		}
		p+=1;
		st_case_367:
		if p == pe && p != eof {
			goto _out367;
			
		}
		if p == eof {
			goto _ctr877;
			
		} else {
			if ( data[ p ] ) == 110 {
				goto _st368;
				
			}
			goto _ctr2;
			
		}
		_ctr879:
		{return st, err }
		goto _st368;
		_st368:
		if p == eof {
			goto _out368;
			
		}
		p+=1;
		st_case_368:
		if p == pe && p != eof {
			goto _out368;
			
		}
		if p == eof {
			goto _ctr879;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr880;
					
				}
				case 47:
				{
					goto _ctr880;
					
				}
				case 117:
				{
					goto _st369;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr882:
		{return st, err }
		goto _st369;
		_st369:
		if p == eof {
			goto _out369;
			
		}
		p+=1;
		st_case_369:
		if p == pe && p != eof {
			goto _out369;
			
		}
		if p == eof {
			goto _ctr882;
			
		} else {
			if ( data[ p ] ) == 97 {
				goto _st370;
				
			}
			goto _ctr2;
			
		}
		_ctr884:
		{return st, err }
		goto _st370;
		_st370:
		if p == eof {
			goto _out370;
			
		}
		p+=1;
		st_case_370:
		if p == pe && p != eof {
			goto _out370;
			
		}
		if p == eof {
			goto _ctr884;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st371;
				
			}
			goto _ctr2;
			
		}
		_ctr886:
		{return st, err }
		goto _st371;
		_st371:
		if p == eof {
			goto _out371;
			
		}
		p+=1;
		st_case_371:
		if p == pe && p != eof {
			goto _out371;
			
		}
		if p == eof {
			goto _ctr886;
			
		} else {
			if ( data[ p ] ) == 121 {
				goto _st372;
				
			}
			goto _ctr2;
			
		}
		_ctr888:
		{return st, err }
		goto _st372;
		_st372:
		if p == eof {
			goto _out372;
			
		}
		p+=1;
		st_case_372:
		if p == pe && p != eof {
			goto _out372;
			
		}
		if p == eof {
			goto _ctr888;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr880;
					
				}
				case 47:
				{
					goto _ctr880;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr889:
		{return st, err }
		goto _st373;
		_st373:
		if p == eof {
			goto _out373;
			
		}
		p+=1;
		st_case_373:
		if p == pe && p != eof {
			goto _out373;
			
		}
		if p == eof {
			goto _ctr889;
			
		} else {
			switch ( data[ p ] ) {
				case 108:
				{
					goto _st374;
					
				}
				case 110:
				{
					goto _st375;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr892:
		{return st, err }
		goto _st374;
		_st374:
		if p == eof {
			goto _out374;
			
		}
		p+=1;
		st_case_374:
		if p == pe && p != eof {
			goto _out374;
			
		}
		if p == eof {
			goto _ctr892;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr893;
					
				}
				case 47:
				{
					goto _ctr893;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr894:
		{return st, err }
		goto _st375;
		_st375:
		if p == eof {
			goto _out375;
			
		}
		p+=1;
		st_case_375:
		if p == pe && p != eof {
			goto _out375;
			
		}
		if p == eof {
			goto _ctr894;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr895;
					
				}
				case 47:
				{
					goto _ctr895;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr896:
		{return st, err }
		goto _st376;
		_st376:
		if p == eof {
			goto _out376;
			
		}
		p+=1;
		st_case_376:
		if p == pe && p != eof {
			goto _out376;
			
		}
		if p == eof {
			goto _ctr896;
			
		} else {
			switch ( data[ p ] ) {
				case 97:
				{
					goto _st377;
					
				}
				case 111:
				{
					goto _st382;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr899:
		{return st, err }
		goto _st377;
		_st377:
		if p == eof {
			goto _out377;
			
		}
		p+=1;
		st_case_377:
		if p == pe && p != eof {
			goto _out377;
			
		}
		if p == eof {
			goto _ctr899;
			
		} else {
			switch ( data[ p ] ) {
				case 114:
				{
					goto _st378;
					
				}
				case 121:
				{
					goto _st381;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr902:
		{return st, err }
		goto _st378;
		_st378:
		if p == eof {
			goto _out378;
			
		}
		p+=1;
		st_case_378:
		if p == pe && p != eof {
			goto _out378;
			
		}
		if p == eof {
			goto _ctr902;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr903;
					
				}
				case 47:
				{
					goto _ctr903;
					
				}
				case 99:
				{
					goto _st379;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr905:
		{return st, err }
		goto _st379;
		_st379:
		if p == eof {
			goto _out379;
			
		}
		p+=1;
		st_case_379:
		if p == pe && p != eof {
			goto _out379;
			
		}
		if p == eof {
			goto _ctr905;
			
		} else {
			if ( data[ p ] ) == 104 {
				goto _st380;
				
			}
			goto _ctr2;
			
		}
		_ctr907:
		{return st, err }
		goto _st380;
		_st380:
		if p == eof {
			goto _out380;
			
		}
		p+=1;
		st_case_380:
		if p == pe && p != eof {
			goto _out380;
			
		}
		if p == eof {
			goto _ctr907;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr903;
					
				}
				case 47:
				{
					goto _ctr903;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr908:
		{return st, err }
		goto _st381;
		_st381:
		if p == eof {
			goto _out381;
			
		}
		p+=1;
		st_case_381:
		if p == pe && p != eof {
			goto _out381;
			
		}
		if p == eof {
			goto _ctr908;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr909;
					
				}
				case 47:
				{
					goto _ctr909;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr910:
		{return st, err }
		goto _st382;
		_st382:
		if p == eof {
			goto _out382;
			
		}
		p+=1;
		st_case_382:
		if p == pe && p != eof {
			goto _out382;
			
		}
		if p == eof {
			goto _ctr910;
			
		} else {
			if ( data[ p ] ) == 110 {
				goto _st159;
				
			}
			goto _ctr2;
			
		}
		_ctr911:
		{return st, err }
		goto _st383;
		_st383:
		if p == eof {
			goto _out383;
			
		}
		p+=1;
		st_case_383:
		if p == pe && p != eof {
			goto _out383;
			
		}
		if p == eof {
			goto _ctr911;
			
		} else {
			if ( data[ p ] ) == 111 {
				goto _st384;
				
			}
			goto _ctr2;
			
		}
		_ctr913:
		{return st, err }
		goto _st384;
		_st384:
		if p == eof {
			goto _out384;
			
		}
		p+=1;
		st_case_384:
		if p == pe && p != eof {
			goto _out384;
			
		}
		if p == eof {
			goto _ctr913;
			
		} else {
			if ( data[ p ] ) == 118 {
				goto _st385;
				
			}
			goto _ctr2;
			
		}
		_ctr915:
		{return st, err }
		goto _st385;
		_st385:
		if p == eof {
			goto _out385;
			
		}
		p+=1;
		st_case_385:
		if p == pe && p != eof {
			goto _out385;
			
		}
		if p == eof {
			goto _ctr915;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr916;
					
				}
				case 47:
				{
					goto _ctr916;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr917:
		{return st, err }
		goto _st386;
		_st386:
		if p == eof {
			goto _out386;
			
		}
		p+=1;
		st_case_386:
		if p == pe && p != eof {
			goto _out386;
			
		}
		if p == eof {
			goto _ctr917;
			
		} else {
			if ( data[ p ] ) == 99 {
				goto _st387;
				
			}
			goto _ctr2;
			
		}
		_ctr919:
		{return st, err }
		goto _st387;
		_st387:
		if p == eof {
			goto _out387;
			
		}
		p+=1;
		st_case_387:
		if p == pe && p != eof {
			goto _out387;
			
		}
		if p == eof {
			goto _ctr919;
			
		} else {
			if ( data[ p ] ) == 116 {
				goto _st388;
				
			}
			goto _ctr2;
			
		}
		_ctr921:
		{return st, err }
		goto _st388;
		_st388:
		if p == eof {
			goto _out388;
			
		}
		p+=1;
		st_case_388:
		if p == pe && p != eof {
			goto _out388;
			
		}
		if p == eof {
			goto _ctr921;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr922;
					
				}
				case 47:
				{
					goto _ctr922;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr923:
		{return st, err }
		goto _st389;
		_st389:
		if p == eof {
			goto _out389;
			
		}
		p+=1;
		st_case_389:
		if p == pe && p != eof {
			goto _out389;
			
		}
		if p == eof {
			goto _ctr923;
			
		} else {
			switch ( data[ p ] ) {
				case 97:
				{
					goto _st390;
					
				}
				case 101:
				{
					goto _st394;
					
				}
				case 117:
				{
					goto _st382;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr926:
		{return st, err }
		goto _st390;
		_st390:
		if p == eof {
			goto _out390;
			
		}
		p+=1;
		st_case_390:
		if p == pe && p != eof {
			goto _out390;
			
		}
		if p == eof {
			goto _ctr926;
			
		} else {
			if ( data[ p ] ) == 116 {
				goto _st391;
				
			}
			goto _ctr2;
			
		}
		_ctr928:
		{return st, err }
		goto _st391;
		_st391:
		if p == eof {
			goto _out391;
			
		}
		p+=1;
		st_case_391:
		if p == pe && p != eof {
			goto _out391;
			
		}
		if p == eof {
			goto _ctr928;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st160;
					
				}
				case 44:
				{
					goto _st269;
					
				}
				case 117:
				{
					goto _st392;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr930:
		{return st, err }
		goto _st392;
		_st392:
		if p == eof {
			goto _out392;
			
		}
		p+=1;
		st_case_392:
		if p == pe && p != eof {
			goto _out392;
			
		}
		if p == eof {
			goto _ctr930;
			
		} else {
			if ( data[ p ] ) == 114 {
				goto _st393;
				
			}
			goto _ctr2;
			
		}
		_ctr932:
		{return st, err }
		goto _st393;
		_st393:
		if p == eof {
			goto _out393;
			
		}
		p+=1;
		st_case_393:
		if p == pe && p != eof {
			goto _out393;
			
		}
		if p == eof {
			goto _ctr932;
			
		} else {
			if ( data[ p ] ) == 100 {
				goto _st363;
				
			}
			goto _ctr2;
			
		}
		_ctr933:
		{return st, err }
		goto _st394;
		_st394:
		if p == eof {
			goto _out394;
			
		}
		p+=1;
		st_case_394:
		if p == pe && p != eof {
			goto _out394;
			
		}
		if p == eof {
			goto _ctr933;
			
		} else {
			if ( data[ p ] ) == 112 {
				goto _st395;
				
			}
			goto _ctr2;
			
		}
		_ctr935:
		{return st, err }
		goto _st395;
		_st395:
		if p == eof {
			goto _out395;
			
		}
		p+=1;
		st_case_395:
		if p == pe && p != eof {
			goto _out395;
			
		}
		if p == eof {
			goto _ctr935;
			
		} else {
			switch ( data[ p ] ) {
				case 45:
				{
					goto _ctr936;
					
				}
				case 47:
				{
					goto _ctr936;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr937:
		{return st, err }
		goto _st396;
		_st396:
		if p == eof {
			goto _out396;
			
		}
		p+=1;
		st_case_396:
		if p == pe && p != eof {
			goto _out396;
			
		}
		if p == eof {
			goto _ctr937;
			
		} else {
			switch ( data[ p ] ) {
				case 104:
				{
					goto _st397;
					
				}
				case 117:
				{
					goto _st400;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr940:
		{return st, err }
		goto _st397;
		_st397:
		if p == eof {
			goto _out397;
			
		}
		p+=1;
		st_case_397:
		if p == pe && p != eof {
			goto _out397;
			
		}
		if p == eof {
			goto _ctr940;
			
		} else {
			if ( data[ p ] ) == 117 {
				goto _st398;
				
			}
			goto _ctr2;
			
		}
		_ctr942:
		{return st, err }
		goto _st398;
		_st398:
		if p == eof {
			goto _out398;
			
		}
		p+=1;
		st_case_398:
		if p == pe && p != eof {
			goto _out398;
			
		}
		if p == eof {
			goto _ctr942;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st160;
					
				}
				case 44:
				{
					goto _st269;
					
				}
				case 114:
				{
					goto _st399;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr944:
		{return st, err }
		goto _st399;
		_st399:
		if p == eof {
			goto _out399;
			
		}
		p+=1;
		st_case_399:
		if p == pe && p != eof {
			goto _out399;
			
		}
		if p == eof {
			goto _ctr944;
			
		} else {
			if ( data[ p ] ) == 115 {
				goto _st393;
				
			}
			goto _ctr2;
			
		}
		_ctr945:
		{return st, err }
		goto _st400;
		_st400:
		if p == eof {
			goto _out400;
			
		}
		p+=1;
		st_case_400:
		if p == pe && p != eof {
			goto _out400;
			
		}
		if p == eof {
			goto _ctr945;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st401;
				
			}
			goto _ctr2;
			
		}
		_ctr947:
		{return st, err }
		goto _st401;
		_st401:
		if p == eof {
			goto _out401;
			
		}
		p+=1;
		st_case_401:
		if p == pe && p != eof {
			goto _out401;
			
		}
		if p == eof {
			goto _ctr947;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st160;
					
				}
				case 44:
				{
					goto _st269;
					
				}
				case 115:
				{
					goto _st393;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr948:
		{return st, err }
		goto _st402;
		_st402:
		if p == eof {
			goto _out402;
			
		}
		p+=1;
		st_case_402:
		if p == pe && p != eof {
			goto _out402;
			
		}
		if p == eof {
			goto _ctr948;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st403;
				
			}
			goto _ctr2;
			
		}
		_ctr950:
		{return st, err }
		goto _st403;
		_st403:
		if p == eof {
			goto _out403;
			
		}
		p+=1;
		st_case_403:
		if p == pe && p != eof {
			goto _out403;
			
		}
		if p == eof {
			goto _ctr950;
			
		} else {
			if ( data[ p ] ) == 100 {
				goto _st404;
				
			}
			goto _ctr2;
			
		}
		_ctr952:
		{return st, err }
		goto _st404;
		_st404:
		if p == eof {
			goto _out404;
			
		}
		p+=1;
		st_case_404:
		if p == pe && p != eof {
			goto _out404;
			
		}
		if p == eof {
			goto _ctr952;
			
		} else {
			switch ( data[ p ] ) {
				case 32:
				{
					goto _st160;
					
				}
				case 44:
				{
					goto _st269;
					
				}
				case 110:
				{
					goto _st405;
					
				}
				
			}
			goto _ctr2;
			
		}
		_ctr954:
		{return st, err }
		goto _st405;
		_st405:
		if p == eof {
			goto _out405;
			
		}
		p+=1;
		st_case_405:
		if p == pe && p != eof {
			goto _out405;
			
		}
		if p == eof {
			goto _ctr954;
			
		} else {
			if ( data[ p ] ) == 101 {
				goto _st399;
				
			}
			goto _ctr2;
			
		}
		_out1:
		cs=1;
		goto _out;
		_out0:
		cs=0;
		goto _out;
		_out2:
		cs=2;
		goto _out;
		_out3:
		cs=3;
		goto _out;
		_out4:
		cs=4;
		goto _out;
		_out5:
		cs=5;
		goto _out;
		_out6:
		cs=6;
		goto _out;
		_out7:
		cs=7;
		goto _out;
		_out8:
		cs=8;
		goto _out;
		_out406:
		cs=406;
		goto _out;
		_out407:
		cs=407;
		goto _out;
		_out9:
		cs=9;
		goto _out;
		_out10:
		cs=10;
		goto _out;
		_out408:
		cs=408;
		goto _out;
		_out11:
		cs=11;
		goto _out;
		_out409:
		cs=409;
		goto _out;
		_out12:
		cs=12;
		goto _out;
		_out410:
		cs=410;
		goto _out;
		_out411:
		cs=411;
		goto _out;
		_out412:
		cs=412;
		goto _out;
		_out413:
		cs=413;
		goto _out;
		_out414:
		cs=414;
		goto _out;
		_out415:
		cs=415;
		goto _out;
		_out416:
		cs=416;
		goto _out;
		_out13:
		cs=13;
		goto _out;
		_out14:
		cs=14;
		goto _out;
		_out15:
		cs=15;
		goto _out;
		_out417:
		cs=417;
		goto _out;
		_out418:
		cs=418;
		goto _out;
		_out419:
		cs=419;
		goto _out;
		_out16:
		cs=16;
		goto _out;
		_out420:
		cs=420;
		goto _out;
		_out421:
		cs=421;
		goto _out;
		_out422:
		cs=422;
		goto _out;
		_out17:
		cs=17;
		goto _out;
		_out18:
		cs=18;
		goto _out;
		_out423:
		cs=423;
		goto _out;
		_out424:
		cs=424;
		goto _out;
		_out19:
		cs=19;
		goto _out;
		_out20:
		cs=20;
		goto _out;
		_out425:
		cs=425;
		goto _out;
		_out21:
		cs=21;
		goto _out;
		_out426:
		cs=426;
		goto _out;
		_out22:
		cs=22;
		goto _out;
		_out427:
		cs=427;
		goto _out;
		_out428:
		cs=428;
		goto _out;
		_out429:
		cs=429;
		goto _out;
		_out430:
		cs=430;
		goto _out;
		_out431:
		cs=431;
		goto _out;
		_out432:
		cs=432;
		goto _out;
		_out433:
		cs=433;
		goto _out;
		_out434:
		cs=434;
		goto _out;
		_out435:
		cs=435;
		goto _out;
		_out23:
		cs=23;
		goto _out;
		_out436:
		cs=436;
		goto _out;
		_out24:
		cs=24;
		goto _out;
		_out437:
		cs=437;
		goto _out;
		_out438:
		cs=438;
		goto _out;
		_out25:
		cs=25;
		goto _out;
		_out439:
		cs=439;
		goto _out;
		_out26:
		cs=26;
		goto _out;
		_out440:
		cs=440;
		goto _out;
		_out441:
		cs=441;
		goto _out;
		_out442:
		cs=442;
		goto _out;
		_out443:
		cs=443;
		goto _out;
		_out444:
		cs=444;
		goto _out;
		_out445:
		cs=445;
		goto _out;
		_out446:
		cs=446;
		goto _out;
		_out447:
		cs=447;
		goto _out;
		_out448:
		cs=448;
		goto _out;
		_out449:
		cs=449;
		goto _out;
		_out450:
		cs=450;
		goto _out;
		_out451:
		cs=451;
		goto _out;
		_out452:
		cs=452;
		goto _out;
		_out27:
		cs=27;
		goto _out;
		_out453:
		cs=453;
		goto _out;
		_out454:
		cs=454;
		goto _out;
		_out455:
		cs=455;
		goto _out;
		_out28:
		cs=28;
		goto _out;
		_out29:
		cs=29;
		goto _out;
		_out30:
		cs=30;
		goto _out;
		_out31:
		cs=31;
		goto _out;
		_out456:
		cs=456;
		goto _out;
		_out457:
		cs=457;
		goto _out;
		_out458:
		cs=458;
		goto _out;
		_out32:
		cs=32;
		goto _out;
		_out33:
		cs=33;
		goto _out;
		_out34:
		cs=34;
		goto _out;
		_out35:
		cs=35;
		goto _out;
		_out36:
		cs=36;
		goto _out;
		_out37:
		cs=37;
		goto _out;
		_out38:
		cs=38;
		goto _out;
		_out39:
		cs=39;
		goto _out;
		_out40:
		cs=40;
		goto _out;
		_out41:
		cs=41;
		goto _out;
		_out42:
		cs=42;
		goto _out;
		_out43:
		cs=43;
		goto _out;
		_out44:
		cs=44;
		goto _out;
		_out45:
		cs=45;
		goto _out;
		_out46:
		cs=46;
		goto _out;
		_out47:
		cs=47;
		goto _out;
		_out48:
		cs=48;
		goto _out;
		_out49:
		cs=49;
		goto _out;
		_out50:
		cs=50;
		goto _out;
		_out51:
		cs=51;
		goto _out;
		_out52:
		cs=52;
		goto _out;
		_out53:
		cs=53;
		goto _out;
		_out54:
		cs=54;
		goto _out;
		_out55:
		cs=55;
		goto _out;
		_out56:
		cs=56;
		goto _out;
		_out57:
		cs=57;
		goto _out;
		_out58:
		cs=58;
		goto _out;
		_out59:
		cs=59;
		goto _out;
		_out60:
		cs=60;
		goto _out;
		_out61:
		cs=61;
		goto _out;
		_out62:
		cs=62;
		goto _out;
		_out63:
		cs=63;
		goto _out;
		_out64:
		cs=64;
		goto _out;
		_out65:
		cs=65;
		goto _out;
		_out66:
		cs=66;
		goto _out;
		_out67:
		cs=67;
		goto _out;
		_out68:
		cs=68;
		goto _out;
		_out69:
		cs=69;
		goto _out;
		_out70:
		cs=70;
		goto _out;
		_out71:
		cs=71;
		goto _out;
		_out72:
		cs=72;
		goto _out;
		_out73:
		cs=73;
		goto _out;
		_out74:
		cs=74;
		goto _out;
		_out75:
		cs=75;
		goto _out;
		_out76:
		cs=76;
		goto _out;
		_out77:
		cs=77;
		goto _out;
		_out78:
		cs=78;
		goto _out;
		_out79:
		cs=79;
		goto _out;
		_out80:
		cs=80;
		goto _out;
		_out459:
		cs=459;
		goto _out;
		_out460:
		cs=460;
		goto _out;
		_out81:
		cs=81;
		goto _out;
		_out82:
		cs=82;
		goto _out;
		_out83:
		cs=83;
		goto _out;
		_out84:
		cs=84;
		goto _out;
		_out85:
		cs=85;
		goto _out;
		_out86:
		cs=86;
		goto _out;
		_out87:
		cs=87;
		goto _out;
		_out461:
		cs=461;
		goto _out;
		_out88:
		cs=88;
		goto _out;
		_out89:
		cs=89;
		goto _out;
		_out90:
		cs=90;
		goto _out;
		_out91:
		cs=91;
		goto _out;
		_out92:
		cs=92;
		goto _out;
		_out93:
		cs=93;
		goto _out;
		_out94:
		cs=94;
		goto _out;
		_out95:
		cs=95;
		goto _out;
		_out96:
		cs=96;
		goto _out;
		_out97:
		cs=97;
		goto _out;
		_out98:
		cs=98;
		goto _out;
		_out99:
		cs=99;
		goto _out;
		_out100:
		cs=100;
		goto _out;
		_out101:
		cs=101;
		goto _out;
		_out102:
		cs=102;
		goto _out;
		_out103:
		cs=103;
		goto _out;
		_out104:
		cs=104;
		goto _out;
		_out105:
		cs=105;
		goto _out;
		_out106:
		cs=106;
		goto _out;
		_out107:
		cs=107;
		goto _out;
		_out108:
		cs=108;
		goto _out;
		_out109:
		cs=109;
		goto _out;
		_out110:
		cs=110;
		goto _out;
		_out111:
		cs=111;
		goto _out;
		_out112:
		cs=112;
		goto _out;
		_out113:
		cs=113;
		goto _out;
		_out114:
		cs=114;
		goto _out;
		_out115:
		cs=115;
		goto _out;
		_out116:
		cs=116;
		goto _out;
		_out117:
		cs=117;
		goto _out;
		_out118:
		cs=118;
		goto _out;
		_out119:
		cs=119;
		goto _out;
		_out120:
		cs=120;
		goto _out;
		_out121:
		cs=121;
		goto _out;
		_out122:
		cs=122;
		goto _out;
		_out123:
		cs=123;
		goto _out;
		_out124:
		cs=124;
		goto _out;
		_out125:
		cs=125;
		goto _out;
		_out126:
		cs=126;
		goto _out;
		_out127:
		cs=127;
		goto _out;
		_out128:
		cs=128;
		goto _out;
		_out129:
		cs=129;
		goto _out;
		_out130:
		cs=130;
		goto _out;
		_out131:
		cs=131;
		goto _out;
		_out132:
		cs=132;
		goto _out;
		_out133:
		cs=133;
		goto _out;
		_out134:
		cs=134;
		goto _out;
		_out135:
		cs=135;
		goto _out;
		_out136:
		cs=136;
		goto _out;
		_out137:
		cs=137;
		goto _out;
		_out138:
		cs=138;
		goto _out;
		_out139:
		cs=139;
		goto _out;
		_out140:
		cs=140;
		goto _out;
		_out462:
		cs=462;
		goto _out;
		_out141:
		cs=141;
		goto _out;
		_out142:
		cs=142;
		goto _out;
		_out143:
		cs=143;
		goto _out;
		_out144:
		cs=144;
		goto _out;
		_out145:
		cs=145;
		goto _out;
		_out146:
		cs=146;
		goto _out;
		_out147:
		cs=147;
		goto _out;
		_out148:
		cs=148;
		goto _out;
		_out149:
		cs=149;
		goto _out;
		_out150:
		cs=150;
		goto _out;
		_out151:
		cs=151;
		goto _out;
		_out152:
		cs=152;
		goto _out;
		_out153:
		cs=153;
		goto _out;
		_out154:
		cs=154;
		goto _out;
		_out155:
		cs=155;
		goto _out;
		_out156:
		cs=156;
		goto _out;
		_out157:
		cs=157;
		goto _out;
		_out158:
		cs=158;
		goto _out;
		_out159:
		cs=159;
		goto _out;
		_out160:
		cs=160;
		goto _out;
		_out161:
		cs=161;
		goto _out;
		_out162:
		cs=162;
		goto _out;
		_out163:
		cs=163;
		goto _out;
		_out164:
		cs=164;
		goto _out;
		_out165:
		cs=165;
		goto _out;
		_out166:
		cs=166;
		goto _out;
		_out167:
		cs=167;
		goto _out;
		_out168:
		cs=168;
		goto _out;
		_out169:
		cs=169;
		goto _out;
		_out170:
		cs=170;
		goto _out;
		_out171:
		cs=171;
		goto _out;
		_out172:
		cs=172;
		goto _out;
		_out173:
		cs=173;
		goto _out;
		_out463:
		cs=463;
		goto _out;
		_out174:
		cs=174;
		goto _out;
		_out175:
		cs=175;
		goto _out;
		_out176:
		cs=176;
		goto _out;
		_out177:
		cs=177;
		goto _out;
		_out178:
		cs=178;
		goto _out;
		_out179:
		cs=179;
		goto _out;
		_out180:
		cs=180;
		goto _out;
		_out181:
		cs=181;
		goto _out;
		_out182:
		cs=182;
		goto _out;
		_out183:
		cs=183;
		goto _out;
		_out184:
		cs=184;
		goto _out;
		_out185:
		cs=185;
		goto _out;
		_out186:
		cs=186;
		goto _out;
		_out187:
		cs=187;
		goto _out;
		_out188:
		cs=188;
		goto _out;
		_out189:
		cs=189;
		goto _out;
		_out190:
		cs=190;
		goto _out;
		_out191:
		cs=191;
		goto _out;
		_out192:
		cs=192;
		goto _out;
		_out193:
		cs=193;
		goto _out;
		_out194:
		cs=194;
		goto _out;
		_out195:
		cs=195;
		goto _out;
		_out196:
		cs=196;
		goto _out;
		_out197:
		cs=197;
		goto _out;
		_out198:
		cs=198;
		goto _out;
		_out199:
		cs=199;
		goto _out;
		_out200:
		cs=200;
		goto _out;
		_out201:
		cs=201;
		goto _out;
		_out202:
		cs=202;
		goto _out;
		_out203:
		cs=203;
		goto _out;
		_out204:
		cs=204;
		goto _out;
		_out205:
		cs=205;
		goto _out;
		_out206:
		cs=206;
		goto _out;
		_out207:
		cs=207;
		goto _out;
		_out208:
		cs=208;
		goto _out;
		_out209:
		cs=209;
		goto _out;
		_out210:
		cs=210;
		goto _out;
		_out211:
		cs=211;
		goto _out;
		_out212:
		cs=212;
		goto _out;
		_out213:
		cs=213;
		goto _out;
		_out214:
		cs=214;
		goto _out;
		_out215:
		cs=215;
		goto _out;
		_out216:
		cs=216;
		goto _out;
		_out217:
		cs=217;
		goto _out;
		_out218:
		cs=218;
		goto _out;
		_out219:
		cs=219;
		goto _out;
		_out220:
		cs=220;
		goto _out;
		_out221:
		cs=221;
		goto _out;
		_out222:
		cs=222;
		goto _out;
		_out223:
		cs=223;
		goto _out;
		_out224:
		cs=224;
		goto _out;
		_out225:
		cs=225;
		goto _out;
		_out226:
		cs=226;
		goto _out;
		_out227:
		cs=227;
		goto _out;
		_out228:
		cs=228;
		goto _out;
		_out229:
		cs=229;
		goto _out;
		_out230:
		cs=230;
		goto _out;
		_out231:
		cs=231;
		goto _out;
		_out232:
		cs=232;
		goto _out;
		_out233:
		cs=233;
		goto _out;
		_out234:
		cs=234;
		goto _out;
		_out235:
		cs=235;
		goto _out;
		_out236:
		cs=236;
		goto _out;
		_out237:
		cs=237;
		goto _out;
		_out238:
		cs=238;
		goto _out;
		_out239:
		cs=239;
		goto _out;
		_out240:
		cs=240;
		goto _out;
		_out241:
		cs=241;
		goto _out;
		_out242:
		cs=242;
		goto _out;
		_out243:
		cs=243;
		goto _out;
		_out244:
		cs=244;
		goto _out;
		_out245:
		cs=245;
		goto _out;
		_out246:
		cs=246;
		goto _out;
		_out247:
		cs=247;
		goto _out;
		_out248:
		cs=248;
		goto _out;
		_out249:
		cs=249;
		goto _out;
		_out250:
		cs=250;
		goto _out;
		_out251:
		cs=251;
		goto _out;
		_out252:
		cs=252;
		goto _out;
		_out253:
		cs=253;
		goto _out;
		_out254:
		cs=254;
		goto _out;
		_out255:
		cs=255;
		goto _out;
		_out256:
		cs=256;
		goto _out;
		_out257:
		cs=257;
		goto _out;
		_out258:
		cs=258;
		goto _out;
		_out259:
		cs=259;
		goto _out;
		_out260:
		cs=260;
		goto _out;
		_out261:
		cs=261;
		goto _out;
		_out262:
		cs=262;
		goto _out;
		_out263:
		cs=263;
		goto _out;
		_out264:
		cs=264;
		goto _out;
		_out265:
		cs=265;
		goto _out;
		_out266:
		cs=266;
		goto _out;
		_out267:
		cs=267;
		goto _out;
		_out268:
		cs=268;
		goto _out;
		_out269:
		cs=269;
		goto _out;
		_out270:
		cs=270;
		goto _out;
		_out271:
		cs=271;
		goto _out;
		_out272:
		cs=272;
		goto _out;
		_out273:
		cs=273;
		goto _out;
		_out274:
		cs=274;
		goto _out;
		_out275:
		cs=275;
		goto _out;
		_out276:
		cs=276;
		goto _out;
		_out277:
		cs=277;
		goto _out;
		_out278:
		cs=278;
		goto _out;
		_out279:
		cs=279;
		goto _out;
		_out280:
		cs=280;
		goto _out;
		_out281:
		cs=281;
		goto _out;
		_out282:
		cs=282;
		goto _out;
		_out283:
		cs=283;
		goto _out;
		_out284:
		cs=284;
		goto _out;
		_out285:
		cs=285;
		goto _out;
		_out286:
		cs=286;
		goto _out;
		_out287:
		cs=287;
		goto _out;
		_out288:
		cs=288;
		goto _out;
		_out289:
		cs=289;
		goto _out;
		_out290:
		cs=290;
		goto _out;
		_out291:
		cs=291;
		goto _out;
		_out292:
		cs=292;
		goto _out;
		_out293:
		cs=293;
		goto _out;
		_out294:
		cs=294;
		goto _out;
		_out295:
		cs=295;
		goto _out;
		_out296:
		cs=296;
		goto _out;
		_out297:
		cs=297;
		goto _out;
		_out298:
		cs=298;
		goto _out;
		_out299:
		cs=299;
		goto _out;
		_out300:
		cs=300;
		goto _out;
		_out301:
		cs=301;
		goto _out;
		_out302:
		cs=302;
		goto _out;
		_out303:
		cs=303;
		goto _out;
		_out304:
		cs=304;
		goto _out;
		_out305:
		cs=305;
		goto _out;
		_out306:
		cs=306;
		goto _out;
		_out307:
		cs=307;
		goto _out;
		_out308:
		cs=308;
		goto _out;
		_out309:
		cs=309;
		goto _out;
		_out310:
		cs=310;
		goto _out;
		_out311:
		cs=311;
		goto _out;
		_out312:
		cs=312;
		goto _out;
		_out313:
		cs=313;
		goto _out;
		_out314:
		cs=314;
		goto _out;
		_out315:
		cs=315;
		goto _out;
		_out316:
		cs=316;
		goto _out;
		_out317:
		cs=317;
		goto _out;
		_out318:
		cs=318;
		goto _out;
		_out319:
		cs=319;
		goto _out;
		_out320:
		cs=320;
		goto _out;
		_out321:
		cs=321;
		goto _out;
		_out322:
		cs=322;
		goto _out;
		_out323:
		cs=323;
		goto _out;
		_out324:
		cs=324;
		goto _out;
		_out325:
		cs=325;
		goto _out;
		_out326:
		cs=326;
		goto _out;
		_out327:
		cs=327;
		goto _out;
		_out328:
		cs=328;
		goto _out;
		_out329:
		cs=329;
		goto _out;
		_out330:
		cs=330;
		goto _out;
		_out331:
		cs=331;
		goto _out;
		_out332:
		cs=332;
		goto _out;
		_out333:
		cs=333;
		goto _out;
		_out334:
		cs=334;
		goto _out;
		_out335:
		cs=335;
		goto _out;
		_out336:
		cs=336;
		goto _out;
		_out337:
		cs=337;
		goto _out;
		_out338:
		cs=338;
		goto _out;
		_out339:
		cs=339;
		goto _out;
		_out340:
		cs=340;
		goto _out;
		_out341:
		cs=341;
		goto _out;
		_out342:
		cs=342;
		goto _out;
		_out343:
		cs=343;
		goto _out;
		_out344:
		cs=344;
		goto _out;
		_out345:
		cs=345;
		goto _out;
		_out346:
		cs=346;
		goto _out;
		_out347:
		cs=347;
		goto _out;
		_out348:
		cs=348;
		goto _out;
		_out349:
		cs=349;
		goto _out;
		_out350:
		cs=350;
		goto _out;
		_out351:
		cs=351;
		goto _out;
		_out352:
		cs=352;
		goto _out;
		_out353:
		cs=353;
		goto _out;
		_out354:
		cs=354;
		goto _out;
		_out355:
		cs=355;
		goto _out;
		_out356:
		cs=356;
		goto _out;
		_out357:
		cs=357;
		goto _out;
		_out358:
		cs=358;
		goto _out;
		_out359:
		cs=359;
		goto _out;
		_out360:
		cs=360;
		goto _out;
		_out361:
		cs=361;
		goto _out;
		_out362:
		cs=362;
		goto _out;
		_out363:
		cs=363;
		goto _out;
		_out364:
		cs=364;
		goto _out;
		_out365:
		cs=365;
		goto _out;
		_out366:
		cs=366;
		goto _out;
		_out367:
		cs=367;
		goto _out;
		_out368:
		cs=368;
		goto _out;
		_out369:
		cs=369;
		goto _out;
		_out370:
		cs=370;
		goto _out;
		_out371:
		cs=371;
		goto _out;
		_out372:
		cs=372;
		goto _out;
		_out373:
		cs=373;
		goto _out;
		_out374:
		cs=374;
		goto _out;
		_out375:
		cs=375;
		goto _out;
		_out376:
		cs=376;
		goto _out;
		_out377:
		cs=377;
		goto _out;
		_out378:
		cs=378;
		goto _out;
		_out379:
		cs=379;
		goto _out;
		_out380:
		cs=380;
		goto _out;
		_out381:
		cs=381;
		goto _out;
		_out382:
		cs=382;
		goto _out;
		_out383:
		cs=383;
		goto _out;
		_out384:
		cs=384;
		goto _out;
		_out385:
		cs=385;
		goto _out;
		_out386:
		cs=386;
		goto _out;
		_out387:
		cs=387;
		goto _out;
		_out388:
		cs=388;
		goto _out;
		_out389:
		cs=389;
		goto _out;
		_out390:
		cs=390;
		goto _out;
		_out391:
		cs=391;
		goto _out;
		_out392:
		cs=392;
		goto _out;
		_out393:
		cs=393;
		goto _out;
		_out394:
		cs=394;
		goto _out;
		_out395:
		cs=395;
		goto _out;
		_out396:
		cs=396;
		goto _out;
		_out397:
		cs=397;
		goto _out;
		_out398:
		cs=398;
		goto _out;
		_out399:
		cs=399;
		goto _out;
		_out400:
		cs=400;
		goto _out;
		_out401:
		cs=401;
		goto _out;
		_out402:
		cs=402;
		goto _out;
		_out403:
		cs=403;
		goto _out;
		_out404:
		cs=404;
		goto _out;
		_out405:
		cs=405;
		goto _out;
		_out:
		{
		
		}
		
	}
	if p != eof {  // input date not fully consumed
		err = errors.New("input data not fully consumed");
		return
	}
	
	if cs < datetime_parser_first_final {
		err = fmt.Errorf("time parse error: %s", data)
	}
	return
}

