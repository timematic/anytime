
//line ragel/parse_datetime.go.rl:1
// DO NOT EDIT!!! GENERETED BY ragel AS:
//  ragel -Z -G2 -e -o ragel_parse_datetime.go parse_datetime.go.rl
// it might be helpful to generate the FSM graph in debug:
//   ragel -Vp parse_datetime.go.rl -o parse_datetime.dot
//   dot parse_datetime.dot -Tpng -o parse_datetime.png

package anytime

import (
    "fmt"
    "strconv"
    "errors"
    "strings"
)


//line ragel_parse_datetime.go:20
const datetime_parser_start int = 1
const datetime_parser_first_final int = 976
const datetime_parser_error int = 0

const datetime_parser_en_main int = 1


//line ragel/parse_datetime.go.rl:24



//line ragel_parse_datetime.go:32
const start int = 1

const en_main int = 1


//line ragel/parse_datetime.go.rl:27

func ragelParse(data string) (st parsedTime, err error) {
    cs, p, pe := 0, 0, len(data)
    eof := pe
    pb := p

    
//line ragel_parse_datetime.go:46
	{
	cs = start
	}

//line ragel/parse_datetime.go.rl:34
    
//line ragel_parse_datetime.go:53
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
	case 976:
		goto st_case_976
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
	case 977:
		goto st_case_977
	case 978:
		goto st_case_978
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 979:
		goto st_case_979
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 980:
		goto st_case_980
	case 16:
		goto st_case_16
	case 981:
		goto st_case_981
	case 17:
		goto st_case_17
	case 982:
		goto st_case_982
	case 18:
		goto st_case_18
	case 983:
		goto st_case_983
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 984:
		goto st_case_984
	case 22:
		goto st_case_22
	case 985:
		goto st_case_985
	case 23:
		goto st_case_23
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
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 992:
		goto st_case_992
	case 27:
		goto st_case_27
	case 993:
		goto st_case_993
	case 28:
		goto st_case_28
	case 994:
		goto st_case_994
	case 29:
		goto st_case_29
	case 995:
		goto st_case_995
	case 996:
		goto st_case_996
	case 997:
		goto st_case_997
	case 998:
		goto st_case_998
	case 999:
		goto st_case_999
	case 1000:
		goto st_case_1000
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 1001:
		goto st_case_1001
	case 1002:
		goto st_case_1002
	case 1003:
		goto st_case_1003
	case 33:
		goto st_case_33
	case 1004:
		goto st_case_1004
	case 1005:
		goto st_case_1005
	case 1006:
		goto st_case_1006
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 1007:
		goto st_case_1007
	case 1008:
		goto st_case_1008
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 1009:
		goto st_case_1009
	case 39:
		goto st_case_39
	case 1010:
		goto st_case_1010
	case 40:
		goto st_case_40
	case 1011:
		goto st_case_1011
	case 1012:
		goto st_case_1012
	case 1013:
		goto st_case_1013
	case 1014:
		goto st_case_1014
	case 1015:
		goto st_case_1015
	case 1016:
		goto st_case_1016
	case 1017:
		goto st_case_1017
	case 1018:
		goto st_case_1018
	case 1019:
		goto st_case_1019
	case 41:
		goto st_case_41
	case 1020:
		goto st_case_1020
	case 42:
		goto st_case_42
	case 1021:
		goto st_case_1021
	case 1022:
		goto st_case_1022
	case 43:
		goto st_case_43
	case 1023:
		goto st_case_1023
	case 44:
		goto st_case_44
	case 1024:
		goto st_case_1024
	case 1025:
		goto st_case_1025
	case 1026:
		goto st_case_1026
	case 1027:
		goto st_case_1027
	case 1028:
		goto st_case_1028
	case 1029:
		goto st_case_1029
	case 1030:
		goto st_case_1030
	case 1031:
		goto st_case_1031
	case 1032:
		goto st_case_1032
	case 1033:
		goto st_case_1033
	case 1034:
		goto st_case_1034
	case 1035:
		goto st_case_1035
	case 1036:
		goto st_case_1036
	case 45:
		goto st_case_45
	case 1037:
		goto st_case_1037
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
	case 1038:
		goto st_case_1038
	case 1039:
		goto st_case_1039
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 1040:
		goto st_case_1040
	case 1041:
		goto st_case_1041
	case 1042:
		goto st_case_1042
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 1043:
		goto st_case_1043
	case 1044:
		goto st_case_1044
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
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 1045:
		goto st_case_1045
	case 1046:
		goto st_case_1046
	case 129:
		goto st_case_129
	case 1047:
		goto st_case_1047
	case 1048:
		goto st_case_1048
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 1049:
		goto st_case_1049
	case 1050:
		goto st_case_1050
	case 132:
		goto st_case_132
	case 1051:
		goto st_case_1051
	case 133:
		goto st_case_133
	case 1052:
		goto st_case_1052
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 1053:
		goto st_case_1053
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 1054:
		goto st_case_1054
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 1055:
		goto st_case_1055
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	case 1056:
		goto st_case_1056
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 1057:
		goto st_case_1057
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 1058:
		goto st_case_1058
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 1059:
		goto st_case_1059
	case 153:
		goto st_case_153
	case 1060:
		goto st_case_1060
	case 1061:
		goto st_case_1061
	case 1062:
		goto st_case_1062
	case 1063:
		goto st_case_1063
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 1064:
		goto st_case_1064
	case 156:
		goto st_case_156
	case 1065:
		goto st_case_1065
	case 1066:
		goto st_case_1066
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 1067:
		goto st_case_1067
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 1068:
		goto st_case_1068
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 1069:
		goto st_case_1069
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 1070:
		goto st_case_1070
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 1071:
		goto st_case_1071
	case 1072:
		goto st_case_1072
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 1073:
		goto st_case_1073
	case 174:
		goto st_case_174
	case 175:
		goto st_case_175
	case 1074:
		goto st_case_1074
	case 1075:
		goto st_case_1075
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
	case 1076:
		goto st_case_1076
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
	case 1077:
		goto st_case_1077
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
	case 1078:
		goto st_case_1078
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	case 1079:
		goto st_case_1079
	case 211:
		goto st_case_211
	case 212:
		goto st_case_212
	case 213:
		goto st_case_213
	case 1080:
		goto st_case_1080
	case 214:
		goto st_case_214
	case 1081:
		goto st_case_1081
	case 215:
		goto st_case_215
	case 1082:
		goto st_case_1082
	case 1083:
		goto st_case_1083
	case 1084:
		goto st_case_1084
	case 1085:
		goto st_case_1085
	case 1086:
		goto st_case_1086
	case 1087:
		goto st_case_1087
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 218:
		goto st_case_218
	case 1088:
		goto st_case_1088
	case 219:
		goto st_case_219
	case 1089:
		goto st_case_1089
	case 1090:
		goto st_case_1090
	case 1091:
		goto st_case_1091
	case 1092:
		goto st_case_1092
	case 1093:
		goto st_case_1093
	case 1094:
		goto st_case_1094
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
	case 222:
		goto st_case_222
	case 1095:
		goto st_case_1095
	case 1096:
		goto st_case_1096
	case 223:
		goto st_case_223
	case 224:
		goto st_case_224
	case 1097:
		goto st_case_1097
	case 225:
		goto st_case_225
	case 1098:
		goto st_case_1098
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
	case 1099:
		goto st_case_1099
	case 1100:
		goto st_case_1100
	case 1101:
		goto st_case_1101
	case 1102:
		goto st_case_1102
	case 232:
		goto st_case_232
	case 1103:
		goto st_case_1103
	case 1104:
		goto st_case_1104
	case 1105:
		goto st_case_1105
	case 1106:
		goto st_case_1106
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
	case 1107:
		goto st_case_1107
	case 288:
		goto st_case_288
	case 1108:
		goto st_case_1108
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
	case 1109:
		goto st_case_1109
	case 446:
		goto st_case_446
	case 1110:
		goto st_case_1110
	case 1111:
		goto st_case_1111
	case 447:
		goto st_case_447
	case 1112:
		goto st_case_1112
	case 1113:
		goto st_case_1113
	case 448:
		goto st_case_448
	case 1114:
		goto st_case_1114
	case 449:
		goto st_case_449
	case 1115:
		goto st_case_1115
	case 450:
		goto st_case_450
	case 1116:
		goto st_case_1116
	case 1117:
		goto st_case_1117
	case 1118:
		goto st_case_1118
	case 1119:
		goto st_case_1119
	case 1120:
		goto st_case_1120
	case 1121:
		goto st_case_1121
	case 1122:
		goto st_case_1122
	case 1123:
		goto st_case_1123
	case 1124:
		goto st_case_1124
	case 451:
		goto st_case_451
	case 1125:
		goto st_case_1125
	case 452:
		goto st_case_452
	case 1126:
		goto st_case_1126
	case 1127:
		goto st_case_1127
	case 453:
		goto st_case_453
	case 1128:
		goto st_case_1128
	case 454:
		goto st_case_454
	case 1129:
		goto st_case_1129
	case 1130:
		goto st_case_1130
	case 1131:
		goto st_case_1131
	case 1132:
		goto st_case_1132
	case 1133:
		goto st_case_1133
	case 1134:
		goto st_case_1134
	case 1135:
		goto st_case_1135
	case 1136:
		goto st_case_1136
	case 1137:
		goto st_case_1137
	case 1138:
		goto st_case_1138
	case 1139:
		goto st_case_1139
	case 1140:
		goto st_case_1140
	case 1141:
		goto st_case_1141
	case 455:
		goto st_case_455
	case 1142:
		goto st_case_1142
	case 456:
		goto st_case_456
	case 457:
		goto st_case_457
	case 458:
		goto st_case_458
	case 459:
		goto st_case_459
	case 1143:
		goto st_case_1143
	case 460:
		goto st_case_460
	case 1144:
		goto st_case_1144
	case 461:
		goto st_case_461
	case 1145:
		goto st_case_1145
	case 1146:
		goto st_case_1146
	case 462:
		goto st_case_462
	case 1147:
		goto st_case_1147
	case 1148:
		goto st_case_1148
	case 1149:
		goto st_case_1149
	case 1150:
		goto st_case_1150
	case 463:
		goto st_case_463
	case 464:
		goto st_case_464
	case 465:
		goto st_case_465
	case 1151:
		goto st_case_1151
	case 1152:
		goto st_case_1152
	case 466:
		goto st_case_466
	case 467:
		goto st_case_467
	case 1153:
		goto st_case_1153
	case 468:
		goto st_case_468
	case 469:
		goto st_case_469
	case 470:
		goto st_case_470
	case 471:
		goto st_case_471
	case 1154:
		goto st_case_1154
	case 472:
		goto st_case_472
	case 1155:
		goto st_case_1155
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
	case 1156:
		goto st_case_1156
	case 538:
		goto st_case_538
	case 1157:
		goto st_case_1157
	case 539:
		goto st_case_539
	case 1158:
		goto st_case_1158
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
	case 1159:
		goto st_case_1159
	case 1160:
		goto st_case_1160
	case 627:
		goto st_case_627
	case 628:
		goto st_case_628
	case 1161:
		goto st_case_1161
	case 629:
		goto st_case_629
	case 630:
		goto st_case_630
	case 631:
		goto st_case_631
	case 1162:
		goto st_case_1162
	case 632:
		goto st_case_632
	case 1163:
		goto st_case_1163
	case 633:
		goto st_case_633
	case 1164:
		goto st_case_1164
	case 1165:
		goto st_case_1165
	case 1166:
		goto st_case_1166
	case 1167:
		goto st_case_1167
	case 1168:
		goto st_case_1168
	case 1169:
		goto st_case_1169
	case 634:
		goto st_case_634
	case 635:
		goto st_case_635
	case 636:
		goto st_case_636
	case 1170:
		goto st_case_1170
	case 637:
		goto st_case_637
	case 1171:
		goto st_case_1171
	case 638:
		goto st_case_638
	case 1172:
		goto st_case_1172
	case 639:
		goto st_case_639
	case 1173:
		goto st_case_1173
	case 1174:
		goto st_case_1174
	case 1175:
		goto st_case_1175
	case 1176:
		goto st_case_1176
	case 1177:
		goto st_case_1177
	case 1178:
		goto st_case_1178
	case 640:
		goto st_case_640
	case 641:
		goto st_case_641
	case 642:
		goto st_case_642
	case 1179:
		goto st_case_1179
	case 643:
		goto st_case_643
	case 1180:
		goto st_case_1180
	case 644:
		goto st_case_644
	case 645:
		goto st_case_645
	case 1181:
		goto st_case_1181
	case 1182:
		goto st_case_1182
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
	case 1183:
		goto st_case_1183
	case 1184:
		goto st_case_1184
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
	case 875:
		goto st_case_875
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
	case 882:
		goto st_case_882
	case 883:
		goto st_case_883
	case 884:
		goto st_case_884
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
	case 891:
		goto st_case_891
	case 892:
		goto st_case_892
	case 893:
		goto st_case_893
	case 894:
		goto st_case_894
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
	case 910:
		goto st_case_910
	case 911:
		goto st_case_911
	case 912:
		goto st_case_912
	case 913:
		goto st_case_913
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
	case 927:
		goto st_case_927
	case 928:
		goto st_case_928
	case 929:
		goto st_case_929
	case 930:
		goto st_case_930
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
	case 944:
		goto st_case_944
	case 945:
		goto st_case_945
	case 946:
		goto st_case_946
	case 947:
		goto st_case_947
	case 948:
		goto st_case_948
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
	case 961:
		goto st_case_961
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
	case 969:
		goto st_case_969
	case 970:
		goto st_case_970
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
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 48:
			goto tr0
		case 49:
			goto tr2
		case 50:
			goto tr3
		case 51:
			goto tr4
		case 65:
			goto st437
		case 68:
			goto st496
		case 70:
			goto st504
		case 74:
			goto st916
		case 77:
			goto st928
		case 78:
			goto st935
		case 79:
			goto st943
		case 83:
			goto st950
		case 84:
			goto st963
		case 87:
			goto st969
		case 97:
			goto st437
		case 100:
			goto st496
		case 102:
			goto st973
		case 106:
			goto st916
		case 109:
			goto st974
		case 110:
			goto st935
		case 111:
			goto st943
		case 115:
			goto st975
		}
		if 52 <= data[p] && data[p] <= 57 {
			goto tr5
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
//line ragel_parse_datetime.go:2495
		if data[p] == 48 {
			goto st3
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st195
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
			goto st976
		}
		goto st0
	st976:
		if p++; p == pe {
			goto _test_eof976
		}
	st_case_976:
		switch data[p] {
		case 32:
			goto tr1349
		case 229:
			goto tr1352
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr1351
			}
		case data[p] >= 45:
			goto tr1350
		}
		goto st0
tr1349:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line ragel_parse_datetime.go:2552
		switch data[p] {
		case 48:
			goto tr23
		case 49:
			goto tr24
		case 65:
			goto st56
		case 68:
			goto st67
		case 70:
			goto st75
		case 74:
			goto st83
		case 77:
			goto st95
		case 78:
			goto st101
		case 79:
			goto st109
		case 83:
			goto st116
		case 97:
			goto st56
		case 100:
			goto st67
		case 102:
			goto st75
		case 106:
			goto st83
		case 109:
			goto st125
		case 110:
			goto st101
		case 111:
			goto st109
		case 115:
			goto st116
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr25
		}
		goto st0
tr23:
//line ragel/datetime.rl:5
 pb = p 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line ragel_parse_datetime.go:2604
		if 49 <= data[p] && data[p] <= 57 {
			goto st7
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
//line ragel_parse_datetime.go:2618
		if data[p] == 32 {
			goto tr36
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr36
		}
		goto st0
tr36:
//line ragel/datetime.rl:9

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st8
tr99:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st8
tr105:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st8
tr112:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st8
tr121:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st8
tr131:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st8
tr139:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st8
tr142:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st8
tr148:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st8
tr152:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st8
tr156:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st8
tr165:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st8
tr173:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line ragel_parse_datetime.go:2685
		switch data[p] {
		case 48:
			goto tr37
		case 51:
			goto tr39
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr40
			}
		case data[p] >= 49:
			goto tr38
		}
		goto st0
tr37:
//line ragel/datetime.rl:5
 pb = p 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line ragel_parse_datetime.go:2710
		if 49 <= data[p] && data[p] <= 57 {
			goto st977
		}
		goto st0
tr40:
//line ragel/datetime.rl:5
 pb = p 
	goto st977
	st977:
		if p++; p == pe {
			goto _test_eof977
		}
	st_case_977:
//line ragel_parse_datetime.go:2724
		switch data[p] {
		case 32:
			goto tr1353
		case 43:
			goto tr1354
		case 44:
			goto tr1355
		case 45:
			goto tr1356
		case 47:
			goto tr1357
		case 84:
			goto tr1358
		case 90:
			goto tr1359
		case 95:
			goto tr1360
		case 110:
			goto tr1361
		case 114:
			goto tr1361
		case 115:
			goto tr1362
		case 116:
			goto tr1363
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1357
			}
		case data[p] >= 65:
			goto tr1357
		}
		goto st0
tr1353:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st978
tr1556:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st978
tr1499:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st978
tr1546:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st978
tr1599:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st978
	st978:
		if p++; p == pe {
			goto _test_eof978
		}
	st_case_978:
//line ragel_parse_datetime.go:2801
		switch data[p] {
		case 32:
			goto st10
		case 43:
			goto st18
		case 45:
			goto st30
		case 47:
			goto tr1367
		case 50:
			goto tr93
		case 65:
			goto tr1368
		case 66:
			goto tr1369
		case 90:
			goto tr1370
		case 95:
			goto tr1367
		case 97:
			goto tr1371
		case 109:
			goto tr1372
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr92
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1367
				}
			case data[p] >= 67:
				goto tr1367
			}
		default:
			goto tr94
		}
		goto st0
tr1380:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line ragel_parse_datetime.go:2853
		switch data[p] {
		case 65:
			goto st11
		case 66:
			goto st17
		case 109:
			goto st13
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 68 {
			goto st979
		}
		goto st0
	st979:
		if p++; p == pe {
			goto _test_eof979
		}
	st_case_979:
		if data[p] == 32 {
			goto st12
		}
		goto st0
tr1568:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st12
tr1715:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st12
tr1376:
//line ragel/datetime.rl:67

    st.Ad_bc = ADBC_BC;

	goto st12
tr1713:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line ragel_parse_datetime.go:2908
		if data[p] == 109 {
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 61 {
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 43:
			goto st15
		case 45:
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr48
		}
		goto st0
tr48:
//line ragel/datetime.rl:5
 pb = p 
	goto st980
	st980:
		if p++; p == pe {
			goto _test_eof980
		}
	st_case_980:
//line ragel_parse_datetime.go:2952
		if data[p] == 46 {
			goto st16
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st980
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= data[p] && data[p] <= 57 {
			goto st981
		}
		goto st0
	st981:
		if p++; p == pe {
			goto _test_eof981
		}
	st_case_981:
		if 48 <= data[p] && data[p] <= 57 {
			goto st981
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 67 {
			goto st982
		}
		goto st0
	st982:
		if p++; p == pe {
			goto _test_eof982
		}
	st_case_982:
		if data[p] == 32 {
			goto tr1376
		}
		goto st0
tr1354:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st18
tr1400:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st18
tr1412:
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

	goto st18
tr1417:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st18
tr1432:
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

	goto st18
tr1425:
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

	goto st18
tr1445:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st18
tr1454:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st18
tr1462:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st18
tr1482:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st18
tr1497:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st18
tr1557:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st18
tr1500:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st18
tr1547:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st18
tr1600:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line ragel_parse_datetime.go:3153
		if data[p] == 50 {
			goto tr52
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr53
			}
		case data[p] >= 48:
			goto tr51
		}
		goto st0
tr51:
//line ragel/datetime.rl:5
 pb = p 
	goto st983
tr73:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st983
	st983:
		if p++; p == pe {
			goto _test_eof983
		}
	st_case_983:
//line ragel_parse_datetime.go:3181
		switch data[p] {
		case 32:
			goto tr1377
		case 58:
			goto tr1379
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st995
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
	goto st19
tr1392:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st19
tr1395:
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
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line ragel_parse_datetime.go:3249
		switch data[p] {
		case 40:
			goto st20
		case 43:
			goto st22
		case 45:
			goto st24
		case 47:
			goto tr57
		case 65:
			goto tr58
		case 66:
			goto tr59
		case 95:
			goto tr57
		case 109:
			goto tr60
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr57
			}
		case data[p] >= 67:
			goto tr57
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 32:
			goto st21
		case 43:
			goto st21
		case 45:
			goto st21
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st21
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st21
			}
		default:
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 32:
			goto st21
		case 41:
			goto st984
		case 43:
			goto st21
		case 45:
			goto st21
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st21
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st21
			}
		default:
			goto st21
		}
		goto st0
	st984:
		if p++; p == pe {
			goto _test_eof984
		}
	st_case_984:
		if data[p] == 32 {
			goto tr1380
		}
		goto st0
tr1397:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line ragel_parse_datetime.go:3352
		if data[p] == 50 {
			goto tr64
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr65
			}
		case data[p] >= 48:
			goto tr63
		}
		goto st0
tr63:
//line ragel/datetime.rl:5
 pb = p 
	goto st985
tr66:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st985
	st985:
		if p++; p == pe {
			goto _test_eof985
		}
	st_case_985:
//line ragel_parse_datetime.go:3380
		switch data[p] {
		case 32:
			goto tr1381
		case 58:
			goto tr1383
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st986
		}
		goto st0
tr1381:
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
	goto st23
tr1385:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st23
tr1388:
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
	goto st23
tr1390:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line ragel_parse_datetime.go:3457
		switch data[p] {
		case 40:
			goto st20
		case 65:
			goto st11
		case 66:
			goto st17
		case 109:
			goto st13
		}
		goto st0
tr65:
//line ragel/datetime.rl:5
 pb = p 
	goto st986
tr68:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st986
	st986:
		if p++; p == pe {
			goto _test_eof986
		}
	st_case_986:
//line ragel_parse_datetime.go:3484
		switch data[p] {
		case 32:
			goto tr1381
		case 58:
			goto tr1383
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st987
		}
		goto st0
	st987:
		if p++; p == pe {
			goto _test_eof987
		}
	st_case_987:
		if data[p] == 32 {
			goto tr1381
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st987
		}
		goto st0
tr1383:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st988
	st988:
		if p++; p == pe {
			goto _test_eof988
		}
	st_case_988:
//line ragel_parse_datetime.go:3524
		if data[p] == 32 {
			goto tr1385
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1387
			}
		case data[p] >= 48:
			goto tr1386
		}
		goto st0
tr1386:
//line ragel/datetime.rl:5
 pb = p 
	goto st989
	st989:
		if p++; p == pe {
			goto _test_eof989
		}
	st_case_989:
//line ragel_parse_datetime.go:3546
		if data[p] == 32 {
			goto tr1388
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st990
		}
		goto st0
tr1387:
//line ragel/datetime.rl:5
 pb = p 
	goto st990
	st990:
		if p++; p == pe {
			goto _test_eof990
		}
	st_case_990:
//line ragel_parse_datetime.go:3563
		if data[p] == 32 {
			goto tr1388
		}
		goto st0
tr64:
//line ragel/datetime.rl:5
 pb = p 
	goto st991
tr67:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st991
	st991:
		if p++; p == pe {
			goto _test_eof991
		}
	st_case_991:
//line ragel_parse_datetime.go:3583
		switch data[p] {
		case 32:
			goto tr1381
		case 58:
			goto tr1383
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st987
			}
		case data[p] >= 48:
			goto st986
		}
		goto st0
tr1398:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line ragel_parse_datetime.go:3611
		if data[p] == 50 {
			goto tr67
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr68
			}
		case data[p] >= 48:
			goto tr66
		}
		goto st0
tr57:
//line ragel/datetime.rl:5
 pb = p 
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line ragel_parse_datetime.go:3633
		switch data[p] {
		case 47:
			goto st26
		case 95:
			goto st26
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		case data[p] >= 65:
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 47:
			goto st992
		case 95:
			goto st992
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st992
			}
		case data[p] >= 65:
			goto st992
		}
		goto st0
	st992:
		if p++; p == pe {
			goto _test_eof992
		}
	st_case_992:
		switch data[p] {
		case 32:
			goto tr1390
		case 47:
			goto st992
		case 95:
			goto st992
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st992
			}
		case data[p] >= 65:
			goto st992
		}
		goto st0
tr58:
//line ragel/datetime.rl:5
 pb = p 
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line ragel_parse_datetime.go:3700
		switch data[p] {
		case 47:
			goto st26
		case 68:
			goto st993
		case 95:
			goto st26
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		case data[p] >= 65:
			goto st26
		}
		goto st0
	st993:
		if p++; p == pe {
			goto _test_eof993
		}
	st_case_993:
		switch data[p] {
		case 32:
			goto st12
		case 47:
			goto st992
		case 95:
			goto st992
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st992
			}
		case data[p] >= 65:
			goto st992
		}
		goto st0
tr59:
//line ragel/datetime.rl:5
 pb = p 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line ragel_parse_datetime.go:3749
		switch data[p] {
		case 47:
			goto st26
		case 67:
			goto st994
		case 95:
			goto st26
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		case data[p] >= 65:
			goto st26
		}
		goto st0
	st994:
		if p++; p == pe {
			goto _test_eof994
		}
	st_case_994:
		switch data[p] {
		case 32:
			goto tr1376
		case 47:
			goto st992
		case 95:
			goto st992
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st992
			}
		case data[p] >= 65:
			goto st992
		}
		goto st0
tr60:
//line ragel/datetime.rl:5
 pb = p 
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line ragel_parse_datetime.go:3798
		switch data[p] {
		case 47:
			goto st26
		case 61:
			goto st14
		case 95:
			goto st26
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		case data[p] >= 65:
			goto st26
		}
		goto st0
tr53:
//line ragel/datetime.rl:5
 pb = p 
	goto st995
tr75:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st995
	st995:
		if p++; p == pe {
			goto _test_eof995
		}
	st_case_995:
//line ragel_parse_datetime.go:3831
		switch data[p] {
		case 32:
			goto tr1377
		case 58:
			goto tr1379
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st996
		}
		goto st0
	st996:
		if p++; p == pe {
			goto _test_eof996
		}
	st_case_996:
		if data[p] == 32 {
			goto tr1377
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st996
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

	goto st997
	st997:
		if p++; p == pe {
			goto _test_eof997
		}
	st_case_997:
//line ragel_parse_datetime.go:3871
		if data[p] == 32 {
			goto tr1392
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1394
			}
		case data[p] >= 48:
			goto tr1393
		}
		goto st0
tr1393:
//line ragel/datetime.rl:5
 pb = p 
	goto st998
	st998:
		if p++; p == pe {
			goto _test_eof998
		}
	st_case_998:
//line ragel_parse_datetime.go:3893
		if data[p] == 32 {
			goto tr1395
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st999
		}
		goto st0
tr1394:
//line ragel/datetime.rl:5
 pb = p 
	goto st999
	st999:
		if p++; p == pe {
			goto _test_eof999
		}
	st_case_999:
//line ragel_parse_datetime.go:3910
		if data[p] == 32 {
			goto tr1395
		}
		goto st0
tr52:
//line ragel/datetime.rl:5
 pb = p 
	goto st1000
tr74:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1000
	st1000:
		if p++; p == pe {
			goto _test_eof1000
		}
	st_case_1000:
//line ragel_parse_datetime.go:3930
		switch data[p] {
		case 32:
			goto tr1377
		case 58:
			goto tr1379
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st996
			}
		case data[p] >= 48:
			goto st995
		}
		goto st0
tr1356:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st30
tr1401:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st30
tr1413:
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

	goto st30
tr1418:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st30
tr1433:
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

	goto st30
tr1427:
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

	goto st30
tr1446:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st30
tr1455:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st30
tr1464:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st30
tr1484:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st30
tr1498:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st30
tr1559:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st30
tr1502:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st30
tr1549:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st30
tr1602:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line ragel_parse_datetime.go:4103
		if data[p] == 50 {
			goto tr74
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr75
			}
		case data[p] >= 48:
			goto tr73
		}
		goto st0
tr1367:
//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1402:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1419:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1447:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1456:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1465:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1434:
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
	goto st31
tr1485:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
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

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1560:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1357:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1503:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1550:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st31
tr1603:
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
//line ragel_parse_datetime.go:4271
		switch data[p] {
		case 47:
			goto st32
		case 95:
			goto st32
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
tr1490:
//line ragel/datetime.rl:5
 pb = p 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line ragel_parse_datetime.go:4296
		switch data[p] {
		case 47:
			goto st1001
		case 95:
			goto st1001
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1001
			}
		case data[p] >= 65:
			goto st1001
		}
		goto st0
tr1494:
//line ragel/datetime.rl:5
 pb = p 
	goto st1001
tr1414:
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
	goto st1001
	st1001:
		if p++; p == pe {
			goto _test_eof1001
		}
	st_case_1001:
//line ragel_parse_datetime.go:4348
		switch data[p] {
		case 32:
			goto tr1390
		case 43:
			goto tr1397
		case 45:
			goto tr1398
		case 47:
			goto st1001
		case 95:
			goto st1001
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1001
			}
		case data[p] >= 65:
			goto st1001
		}
		goto st0
tr92:
//line ragel/datetime.rl:5
 pb = p 
	goto st1002
	st1002:
		if p++; p == pe {
			goto _test_eof1002
		}
	st_case_1002:
//line ragel_parse_datetime.go:4379
		switch data[p] {
		case 32:
			goto tr1399
		case 43:
			goto tr1400
		case 45:
			goto tr1401
		case 47:
			goto tr1402
		case 58:
			goto tr1404
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1406
		case 95:
			goto tr1402
		case 97:
			goto tr1407
		case 112:
			goto tr1407
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1009
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1402
			}
		default:
			goto tr1402
		}
		goto st0
tr1399:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st1003
tr1416:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st1003
tr1470:
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

	goto st1003
tr1444:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st1003
tr1453:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st1003
tr1461:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st1003
tr1481:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st1003
tr1595:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st1003
	st1003:
		if p++; p == pe {
			goto _test_eof1003
		}
	st_case_1003:
//line ragel_parse_datetime.go:4499
		switch data[p] {
		case 32:
			goto st10
		case 43:
			goto st18
		case 45:
			goto st30
		case 47:
			goto tr1367
		case 65:
			goto tr1408
		case 66:
			goto tr1369
		case 80:
			goto tr1409
		case 90:
			goto tr1370
		case 95:
			goto tr1367
		case 97:
			goto tr1410
		case 109:
			goto tr1372
		case 112:
			goto tr1410
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		case data[p] >= 67:
			goto tr1367
		}
		goto st0
tr1408:
//line ragel/datetime.rl:5
 pb = p 
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line ragel_parse_datetime.go:4544
		switch data[p] {
		case 47:
			goto st32
		case 68:
			goto st1004
		case 77:
			goto st1005
		case 95:
			goto st32
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
	st1004:
		if p++; p == pe {
			goto _test_eof1004
		}
	st_case_1004:
		switch data[p] {
		case 32:
			goto st12
		case 47:
			goto st1001
		case 95:
			goto st1001
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1001
			}
		case data[p] >= 65:
			goto st1001
		}
		goto st0
	st1005:
		if p++; p == pe {
			goto _test_eof1005
		}
	st_case_1005:
		switch data[p] {
		case 32:
			goto tr1411
		case 43:
			goto tr1412
		case 45:
			goto tr1413
		case 47:
			goto tr1414
		case 95:
			goto tr1414
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1414
			}
		case data[p] >= 65:
			goto tr1414
		}
		goto st0
tr1411:
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

	goto st1006
tr1431:
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

	goto st1006
tr1424:
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

	goto st1006
tr1597:
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

	goto st1006
	st1006:
		if p++; p == pe {
			goto _test_eof1006
		}
	st_case_1006:
//line ragel_parse_datetime.go:4711
		switch data[p] {
		case 32:
			goto st10
		case 43:
			goto st18
		case 45:
			goto st30
		case 47:
			goto tr1367
		case 65:
			goto tr1415
		case 66:
			goto tr1369
		case 90:
			goto tr1370
		case 95:
			goto tr1367
		case 109:
			goto tr1372
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1367
			}
		case data[p] >= 67:
			goto tr1367
		}
		goto st0
tr1415:
//line ragel/datetime.rl:5
 pb = p 
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line ragel_parse_datetime.go:4750
		switch data[p] {
		case 47:
			goto st32
		case 68:
			goto st1004
		case 95:
			goto st32
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
tr1369:
//line ragel/datetime.rl:5
 pb = p 
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line ragel_parse_datetime.go:4777
		switch data[p] {
		case 47:
			goto st32
		case 67:
			goto st1007
		case 95:
			goto st32
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
	st1007:
		if p++; p == pe {
			goto _test_eof1007
		}
	st_case_1007:
		switch data[p] {
		case 32:
			goto tr1376
		case 47:
			goto st1001
		case 95:
			goto st1001
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1001
			}
		case data[p] >= 65:
			goto st1001
		}
		goto st0
tr1370:
//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1406:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1422:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1451:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1459:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1468:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
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

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1487:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
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

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1562:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1359:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1505:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1552:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
tr1605:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st1008
	st1008:
		if p++; p == pe {
			goto _test_eof1008
		}
	st_case_1008:
//line ragel_parse_datetime.go:4972
		switch data[p] {
		case 32:
			goto tr1385
		case 47:
			goto st32
		case 95:
			goto st32
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
tr1372:
//line ragel/datetime.rl:5
 pb = p 
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line ragel_parse_datetime.go:4999
		switch data[p] {
		case 47:
			goto st32
		case 61:
			goto st14
		case 95:
			goto st32
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
tr1409:
//line ragel/datetime.rl:5
 pb = p 
	goto st37
tr1405:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st37
tr1421:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st37
tr1450:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st37
tr1458:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st37
tr1467:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st37
tr1472:
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
	goto st37
tr1486:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line ragel_parse_datetime.go:5107
		switch data[p] {
		case 47:
			goto st32
		case 77:
			goto st1005
		case 95:
			goto st32
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
tr1410:
//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1407:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1423:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1452:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1460:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1469:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st38
tr1473:
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
tr1488:
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
//line ragel_parse_datetime.go:5215
		switch data[p] {
		case 47:
			goto st32
		case 95:
			goto st32
		case 109:
			goto st1005
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
	st1009:
		if p++; p == pe {
			goto _test_eof1009
		}
	st_case_1009:
		switch data[p] {
		case 32:
			goto tr1416
		case 43:
			goto tr1417
		case 45:
			goto tr1418
		case 47:
			goto tr1419
		case 58:
			goto tr1420
		case 65:
			goto tr1421
		case 80:
			goto tr1421
		case 90:
			goto tr1422
		case 95:
			goto tr1419
		case 97:
			goto tr1423
		case 112:
			goto tr1423
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st39
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1419
			}
		default:
			goto tr1419
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1010
		}
		goto st0
	st1010:
		if p++; p == pe {
			goto _test_eof1010
		}
	st_case_1010:
		switch data[p] {
		case 32:
			goto tr1424
		case 43:
			goto tr1425
		case 45:
			goto tr1427
		case 47:
			goto tr1428
		case 90:
			goto tr1430
		case 95:
			goto tr1428
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1426
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1428
				}
			case data[p] >= 65:
				goto tr1428
			}
		default:
			goto st41
		}
		goto st0
tr1426:
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

	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line ragel_parse_datetime.go:5343
		if 48 <= data[p] && data[p] <= 57 {
			goto tr82
		}
		goto st0
tr82:
//line ragel/datetime.rl:5
 pb = p 
	goto st1011
	st1011:
		if p++; p == pe {
			goto _test_eof1011
		}
	st_case_1011:
//line ragel_parse_datetime.go:5357
		switch data[p] {
		case 32:
			goto tr1431
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1012
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1012:
		if p++; p == pe {
			goto _test_eof1012
		}
	st_case_1012:
		switch data[p] {
		case 32:
			goto tr1431
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1013
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1013:
		if p++; p == pe {
			goto _test_eof1013
		}
	st_case_1013:
		switch data[p] {
		case 32:
			goto tr1431
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1014
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1014:
		if p++; p == pe {
			goto _test_eof1014
		}
	st_case_1014:
		switch data[p] {
		case 32:
			goto tr1431
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1015
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1015:
		if p++; p == pe {
			goto _test_eof1015
		}
	st_case_1015:
		switch data[p] {
		case 32:
			goto tr1431
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1016
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1016:
		if p++; p == pe {
			goto _test_eof1016
		}
	st_case_1016:
		switch data[p] {
		case 32:
			goto tr1431
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1017
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1017:
		if p++; p == pe {
			goto _test_eof1017
		}
	st_case_1017:
		switch data[p] {
		case 32:
			goto tr1431
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1018
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1018:
		if p++; p == pe {
			goto _test_eof1018
		}
	st_case_1018:
		switch data[p] {
		case 32:
			goto tr1431
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1019
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1019:
		if p++; p == pe {
			goto _test_eof1019
		}
	st_case_1019:
		switch data[p] {
		case 32:
			goto tr1431
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		case data[p] >= 65:
			goto tr1434
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1020
		}
		goto st0
	st1020:
		if p++; p == pe {
			goto _test_eof1020
		}
	st_case_1020:
		switch data[p] {
		case 32:
			goto tr1424
		case 43:
			goto tr1425
		case 45:
			goto tr1427
		case 47:
			goto tr1428
		case 90:
			goto tr1430
		case 95:
			goto tr1428
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1426
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1428
			}
		default:
			goto tr1428
		}
		goto st0
tr1404:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st42
tr1420:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line ragel_parse_datetime.go:5695
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr85
			}
		case data[p] >= 48:
			goto tr84
		}
		goto st0
tr84:
//line ragel/datetime.rl:5
 pb = p 
	goto st1021
	st1021:
		if p++; p == pe {
			goto _test_eof1021
		}
	st_case_1021:
//line ragel_parse_datetime.go:5714
		switch data[p] {
		case 32:
			goto tr1444
		case 43:
			goto tr1445
		case 45:
			goto tr1446
		case 47:
			goto tr1447
		case 58:
			goto tr1449
		case 65:
			goto tr1450
		case 80:
			goto tr1450
		case 90:
			goto tr1451
		case 95:
			goto tr1447
		case 97:
			goto tr1452
		case 112:
			goto tr1452
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1022
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1447
			}
		default:
			goto tr1447
		}
		goto st0
	st1022:
		if p++; p == pe {
			goto _test_eof1022
		}
	st_case_1022:
		switch data[p] {
		case 32:
			goto tr1453
		case 43:
			goto tr1454
		case 45:
			goto tr1455
		case 47:
			goto tr1456
		case 58:
			goto tr1457
		case 65:
			goto tr1458
		case 80:
			goto tr1458
		case 90:
			goto tr1459
		case 95:
			goto tr1456
		case 97:
			goto tr1460
		case 112:
			goto tr1460
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1456
			}
		case data[p] >= 66:
			goto tr1456
		}
		goto st0
tr1449:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st43
tr1457:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line ragel_parse_datetime.go:5807
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
	goto st1023
	st1023:
		if p++; p == pe {
			goto _test_eof1023
		}
	st_case_1023:
//line ragel_parse_datetime.go:5826
		switch data[p] {
		case 32:
			goto tr1461
		case 43:
			goto tr1462
		case 45:
			goto tr1464
		case 47:
			goto tr1465
		case 65:
			goto tr1467
		case 80:
			goto tr1467
		case 90:
			goto tr1468
		case 95:
			goto tr1465
		case 97:
			goto tr1469
		case 112:
			goto tr1469
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1463
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1465
				}
			case data[p] >= 66:
				goto tr1465
			}
		default:
			goto st1033
		}
		goto st0
tr1463:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st44
tr1483:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line ragel_parse_datetime.go:5884
		if 48 <= data[p] && data[p] <= 57 {
			goto tr88
		}
		goto st0
tr88:
//line ragel/datetime.rl:5
 pb = p 
	goto st1024
	st1024:
		if p++; p == pe {
			goto _test_eof1024
		}
	st_case_1024:
//line ragel_parse_datetime.go:5898
		switch data[p] {
		case 32:
			goto tr1470
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 65:
			goto tr1472
		case 80:
			goto tr1472
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1025
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1025:
		if p++; p == pe {
			goto _test_eof1025
		}
	st_case_1025:
		switch data[p] {
		case 32:
			goto tr1470
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 65:
			goto tr1472
		case 80:
			goto tr1472
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1026
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1026:
		if p++; p == pe {
			goto _test_eof1026
		}
	st_case_1026:
		switch data[p] {
		case 32:
			goto tr1470
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 65:
			goto tr1472
		case 80:
			goto tr1472
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1027
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1027:
		if p++; p == pe {
			goto _test_eof1027
		}
	st_case_1027:
		switch data[p] {
		case 32:
			goto tr1470
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 65:
			goto tr1472
		case 80:
			goto tr1472
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1028
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1028:
		if p++; p == pe {
			goto _test_eof1028
		}
	st_case_1028:
		switch data[p] {
		case 32:
			goto tr1470
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 65:
			goto tr1472
		case 80:
			goto tr1472
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1029
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1029:
		if p++; p == pe {
			goto _test_eof1029
		}
	st_case_1029:
		switch data[p] {
		case 32:
			goto tr1470
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 65:
			goto tr1472
		case 80:
			goto tr1472
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1030
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1030:
		if p++; p == pe {
			goto _test_eof1030
		}
	st_case_1030:
		switch data[p] {
		case 32:
			goto tr1470
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 65:
			goto tr1472
		case 80:
			goto tr1472
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1031
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1031:
		if p++; p == pe {
			goto _test_eof1031
		}
	st_case_1031:
		switch data[p] {
		case 32:
			goto tr1470
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 65:
			goto tr1472
		case 80:
			goto tr1472
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1032
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		default:
			goto tr1434
		}
		goto st0
	st1032:
		if p++; p == pe {
			goto _test_eof1032
		}
	st_case_1032:
		switch data[p] {
		case 32:
			goto tr1470
		case 43:
			goto tr1432
		case 45:
			goto tr1433
		case 47:
			goto tr1434
		case 65:
			goto tr1472
		case 80:
			goto tr1472
		case 90:
			goto tr1436
		case 95:
			goto tr1434
		case 97:
			goto tr1473
		case 112:
			goto tr1473
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1434
			}
		case data[p] >= 66:
			goto tr1434
		}
		goto st0
	st1033:
		if p++; p == pe {
			goto _test_eof1033
		}
	st_case_1033:
		switch data[p] {
		case 32:
			goto tr1481
		case 43:
			goto tr1482
		case 45:
			goto tr1484
		case 47:
			goto tr1485
		case 65:
			goto tr1486
		case 80:
			goto tr1486
		case 90:
			goto tr1487
		case 95:
			goto tr1485
		case 97:
			goto tr1488
		case 112:
			goto tr1488
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1483
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1485
			}
		default:
			goto tr1485
		}
		goto st0
tr87:
//line ragel/datetime.rl:5
 pb = p 
	goto st1034
	st1034:
		if p++; p == pe {
			goto _test_eof1034
		}
	st_case_1034:
//line ragel_parse_datetime.go:6299
		switch data[p] {
		case 32:
			goto tr1461
		case 43:
			goto tr1462
		case 45:
			goto tr1464
		case 47:
			goto tr1465
		case 65:
			goto tr1467
		case 80:
			goto tr1467
		case 90:
			goto tr1468
		case 95:
			goto tr1465
		case 97:
			goto tr1469
		case 112:
			goto tr1469
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1463
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1465
			}
		default:
			goto tr1465
		}
		goto st0
tr85:
//line ragel/datetime.rl:5
 pb = p 
	goto st1035
	st1035:
		if p++; p == pe {
			goto _test_eof1035
		}
	st_case_1035:
//line ragel_parse_datetime.go:6344
		switch data[p] {
		case 32:
			goto tr1444
		case 43:
			goto tr1445
		case 45:
			goto tr1446
		case 47:
			goto tr1447
		case 58:
			goto tr1449
		case 65:
			goto tr1450
		case 80:
			goto tr1450
		case 90:
			goto tr1451
		case 95:
			goto tr1447
		case 97:
			goto tr1452
		case 112:
			goto tr1452
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1447
			}
		case data[p] >= 66:
			goto tr1447
		}
		goto st0
tr93:
//line ragel/datetime.rl:5
 pb = p 
	goto st1036
	st1036:
		if p++; p == pe {
			goto _test_eof1036
		}
	st_case_1036:
//line ragel_parse_datetime.go:6387
		switch data[p] {
		case 32:
			goto tr1399
		case 43:
			goto tr1400
		case 45:
			goto tr1401
		case 47:
			goto tr1402
		case 58:
			goto tr1404
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1406
		case 95:
			goto tr1402
		case 97:
			goto tr1407
		case 112:
			goto tr1407
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st1009
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1402
				}
			case data[p] >= 66:
				goto tr1402
			}
		default:
			goto st45
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if 48 <= data[p] && data[p] <= 57 {
			goto st39
		}
		goto st0
tr94:
//line ragel/datetime.rl:5
 pb = p 
	goto st1037
	st1037:
		if p++; p == pe {
			goto _test_eof1037
		}
	st_case_1037:
//line ragel_parse_datetime.go:6448
		switch data[p] {
		case 32:
			goto tr1399
		case 43:
			goto tr1400
		case 45:
			goto tr1401
		case 47:
			goto tr1402
		case 58:
			goto tr1404
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1406
		case 95:
			goto tr1402
		case 97:
			goto tr1407
		case 112:
			goto tr1407
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st45
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1402
			}
		default:
			goto tr1402
		}
		goto st0
tr1368:
//line ragel/datetime.rl:5
 pb = p 
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line ragel_parse_datetime.go:6495
		switch data[p] {
		case 47:
			goto st32
		case 68:
			goto st1004
		case 84:
			goto st47
		case 95:
			goto st32
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 32:
			goto st48
		case 47:
			goto st1001
		case 95:
			goto st1001
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1001
			}
		case data[p] >= 65:
			goto st1001
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 50 {
			goto tr93
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr94
			}
		case data[p] >= 48:
			goto tr92
		}
		goto st0
tr1371:
//line ragel/datetime.rl:5
 pb = p 
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line ragel_parse_datetime.go:6563
		switch data[p] {
		case 47:
			goto st32
		case 95:
			goto st32
		case 116:
			goto st47
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
tr1355:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st50
tr1558:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st50
tr1501:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st50
tr1548:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st50
tr1601:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line ragel_parse_datetime.go:6622
		if data[p] == 32 {
			goto st48
		}
		goto st0
tr1554:
//line ragel/datetime.rl:5
 pb = p 
	goto st1038
tr1561:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st1038
tr1358:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st1038
tr1504:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st1038
tr1551:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st1038
tr1604:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st1038
	st1038:
		if p++; p == pe {
			goto _test_eof1038
		}
	st_case_1038:
//line ragel_parse_datetime.go:6682
		switch data[p] {
		case 32:
			goto st10
		case 43:
			goto st18
		case 45:
			goto st30
		case 47:
			goto tr1490
		case 50:
			goto tr93
		case 90:
			goto tr1491
		case 95:
			goto tr1490
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr92
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1490
				}
			case data[p] >= 65:
				goto tr1490
			}
		default:
			goto tr94
		}
		goto st0
tr1491:
//line ragel/datetime.rl:5
 pb = p 
	goto st1039
	st1039:
		if p++; p == pe {
			goto _test_eof1039
		}
	st_case_1039:
//line ragel_parse_datetime.go:6726
		switch data[p] {
		case 32:
			goto tr1385
		case 47:
			goto st1001
		case 95:
			goto st1001
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1001
			}
		case data[p] >= 65:
			goto st1001
		}
		goto st0
tr1555:
//line ragel/datetime.rl:5
 pb = p 
	goto st51
tr1563:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st51
tr1360:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st51
tr1506:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:5
 pb = p 
	goto st51
tr1553:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st51
tr1606:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line ragel_parse_datetime.go:6799
		switch data[p] {
		case 47:
			goto st32
		case 50:
			goto tr93
		case 95:
			goto st32
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr92
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st32
				}
			case data[p] >= 65:
				goto st32
			}
		default:
			goto tr94
		}
		goto st0
tr1361:
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
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line ragel_parse_datetime.go:6844
		switch data[p] {
		case 47:
			goto st32
		case 95:
			goto st32
		case 100:
			goto st1040
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
	st1040:
		if p++; p == pe {
			goto _test_eof1040
		}
	st_case_1040:
		switch data[p] {
		case 32:
			goto st978
		case 43:
			goto st18
		case 44:
			goto st50
		case 45:
			goto st30
		case 47:
			goto tr1494
		case 84:
			goto tr1495
		case 95:
			goto tr1496
		case 116:
			goto tr1496
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1494
			}
		case data[p] >= 65:
			goto tr1494
		}
		goto st0
tr1495:
//line ragel/datetime.rl:5
 pb = p 
	goto st1041
	st1041:
		if p++; p == pe {
			goto _test_eof1041
		}
	st_case_1041:
//line ragel_parse_datetime.go:6903
		switch data[p] {
		case 32:
			goto tr1390
		case 43:
			goto tr1497
		case 45:
			goto tr1498
		case 47:
			goto tr1494
		case 50:
			goto tr93
		case 95:
			goto tr1494
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr92
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1494
				}
			case data[p] >= 65:
				goto tr1494
			}
		default:
			goto tr94
		}
		goto st0
tr1496:
//line ragel/datetime.rl:5
 pb = p 
	goto st1042
	st1042:
		if p++; p == pe {
			goto _test_eof1042
		}
	st_case_1042:
//line ragel_parse_datetime.go:6945
		switch data[p] {
		case 32:
			goto tr1390
		case 43:
			goto tr1397
		case 45:
			goto tr1398
		case 47:
			goto st1001
		case 50:
			goto tr93
		case 95:
			goto st1001
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr92
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st1001
				}
			case data[p] >= 65:
				goto st1001
			}
		default:
			goto tr94
		}
		goto st0
tr1362:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line ragel_parse_datetime.go:6996
		switch data[p] {
		case 47:
			goto st32
		case 95:
			goto st32
		case 116:
			goto st1040
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
tr1363:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:5
 pb = p 
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line ragel_parse_datetime.go:7032
		switch data[p] {
		case 47:
			goto st32
		case 50:
			goto tr93
		case 95:
			goto st32
		case 104:
			goto st1040
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr92
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st32
				}
			case data[p] >= 65:
				goto st32
			}
		default:
			goto tr94
		}
		goto st0
tr38:
//line ragel/datetime.rl:5
 pb = p 
	goto st1043
	st1043:
		if p++; p == pe {
			goto _test_eof1043
		}
	st_case_1043:
//line ragel_parse_datetime.go:7070
		switch data[p] {
		case 32:
			goto tr1353
		case 43:
			goto tr1354
		case 44:
			goto tr1355
		case 45:
			goto tr1356
		case 47:
			goto tr1357
		case 84:
			goto tr1358
		case 90:
			goto tr1359
		case 95:
			goto tr1360
		case 110:
			goto tr1361
		case 114:
			goto tr1361
		case 115:
			goto tr1362
		case 116:
			goto tr1363
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st977
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1357
			}
		default:
			goto tr1357
		}
		goto st0
tr39:
//line ragel/datetime.rl:5
 pb = p 
	goto st1044
	st1044:
		if p++; p == pe {
			goto _test_eof1044
		}
	st_case_1044:
//line ragel_parse_datetime.go:7119
		switch data[p] {
		case 32:
			goto tr1353
		case 43:
			goto tr1354
		case 44:
			goto tr1355
		case 45:
			goto tr1356
		case 47:
			goto tr1357
		case 84:
			goto tr1358
		case 90:
			goto tr1359
		case 95:
			goto tr1360
		case 110:
			goto tr1361
		case 114:
			goto tr1361
		case 115:
			goto tr1362
		case 116:
			goto tr1363
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 49 {
				goto st977
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1357
			}
		default:
			goto tr1357
		}
		goto st0
tr24:
//line ragel/datetime.rl:5
 pb = p 
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line ragel_parse_datetime.go:7168
		if data[p] == 32 {
			goto tr36
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 50 {
				goto st7
			}
		case data[p] >= 45:
			goto tr36
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 112:
			goto st57
		case 117:
			goto st62
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
		switch data[p] {
		case 32:
			goto tr99
		case 46:
			goto tr100
		case 105:
			goto st60
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr99
		}
		goto st0
tr100:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st59
tr106:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st59
tr113:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st59
tr122:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st59
tr132:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st59
tr140:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st59
tr143:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st59
tr149:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st59
tr153:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st59
tr157:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st59
tr166:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st59
tr174:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line ragel_parse_datetime.go:7272
		switch data[p] {
		case 32:
			goto st8
		case 48:
			goto tr37
		case 51:
			goto tr39
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st8
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr40
			}
		default:
			goto tr38
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if data[p] == 108 {
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
			goto tr99
		case 46:
			goto tr100
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr99
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if data[p] == 103 {
			goto st63
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 32:
			goto tr105
		case 46:
			goto tr106
		case 117:
			goto st64
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr105
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 115 {
			goto st65
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 116 {
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
			goto tr105
		case 46:
			goto tr106
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr105
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		if data[p] == 101 {
			goto st68
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if data[p] == 99 {
			goto st69
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 32:
			goto tr112
		case 46:
			goto tr113
		case 101:
			goto st70
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr112
		}
		goto st0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		if data[p] == 109 {
			goto st71
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		if data[p] == 98 {
			goto st72
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 101 {
			goto st73
		}
		goto st0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 114 {
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
			goto tr112
		case 46:
			goto tr113
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr112
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if data[p] == 101 {
			goto st76
		}
		goto st0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		if data[p] == 98 {
			goto st77
		}
		goto st0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 32:
			goto tr121
		case 46:
			goto tr122
		case 114:
			goto st78
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr121
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		if data[p] == 117 {
			goto st79
		}
		goto st0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		if data[p] == 97 {
			goto st80
		}
		goto st0
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		if data[p] == 114 {
			goto st81
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		if data[p] == 121 {
			goto st82
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 32:
			goto tr121
		case 46:
			goto tr122
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr121
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 97:
			goto st84
		case 117:
			goto st90
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 110 {
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
			goto tr131
		case 46:
			goto tr132
		case 117:
			goto st86
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr131
		}
		goto st0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		if data[p] == 97 {
			goto st87
		}
		goto st0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		if data[p] == 114 {
			goto st88
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		if data[p] == 121 {
			goto st89
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 32:
			goto tr131
		case 46:
			goto tr132
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr131
		}
		goto st0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 108:
			goto st91
		case 110:
			goto st93
		}
		goto st0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 32:
			goto tr139
		case 46:
			goto tr140
		case 121:
			goto st92
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr139
		}
		goto st0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 32:
			goto tr139
		case 46:
			goto tr140
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr139
		}
		goto st0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 32:
			goto tr142
		case 46:
			goto tr143
		case 101:
			goto st94
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr142
		}
		goto st0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
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
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		if data[p] == 97 {
			goto st96
		}
		goto st0
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 114:
			goto st97
		case 121:
			goto st100
		}
		goto st0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 32:
			goto tr148
		case 46:
			goto tr149
		case 99:
			goto st98
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr148
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		if data[p] == 104 {
			goto st99
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 32:
			goto tr148
		case 46:
			goto tr149
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr148
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 32:
			goto tr152
		case 46:
			goto tr153
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr152
		}
		goto st0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		if data[p] == 111 {
			goto st102
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		if data[p] == 118 {
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
			goto tr156
		case 46:
			goto tr157
		case 101:
			goto st104
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr156
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		if data[p] == 109 {
			goto st105
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		if data[p] == 98 {
			goto st106
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		if data[p] == 101 {
			goto st107
		}
		goto st0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		if data[p] == 114 {
			goto st108
		}
		goto st0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 32:
			goto tr156
		case 46:
			goto tr157
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr156
		}
		goto st0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		if data[p] == 99 {
			goto st110
		}
		goto st0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		if data[p] == 116 {
			goto st111
		}
		goto st0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 32:
			goto tr165
		case 46:
			goto tr166
		case 111:
			goto st112
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr165
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		if data[p] == 98 {
			goto st113
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		if data[p] == 101 {
			goto st114
		}
		goto st0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		if data[p] == 114 {
			goto st115
		}
		goto st0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 32:
			goto tr165
		case 46:
			goto tr166
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr165
		}
		goto st0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		if data[p] == 101 {
			goto st117
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		if data[p] == 112 {
			goto st118
		}
		goto st0
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 32:
			goto tr173
		case 46:
			goto tr174
		case 116:
			goto st119
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr173
		}
		goto st0
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 32:
			goto tr173
		case 46:
			goto tr174
		case 101:
			goto st120
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr173
		}
		goto st0
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		if data[p] == 109 {
			goto st121
		}
		goto st0
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		if data[p] == 98 {
			goto st122
		}
		goto st0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		if data[p] == 101 {
			goto st123
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		if data[p] == 114 {
			goto st124
		}
		goto st0
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 32:
			goto tr173
		case 46:
			goto tr174
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr173
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 61:
			goto st14
		case 97:
			goto st96
		}
		goto st0
tr1350:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line ragel_parse_datetime.go:8071
		switch data[p] {
		case 48:
			goto tr181
		case 49:
			goto tr182
		case 65:
			goto st130
		case 68:
			goto st136
		case 70:
			goto st142
		case 74:
			goto st148
		case 77:
			goto st154
		case 78:
			goto st157
		case 79:
			goto st163
		case 83:
			goto st168
		case 97:
			goto st130
		case 100:
			goto st136
		case 102:
			goto st142
		case 106:
			goto st148
		case 109:
			goto st154
		case 110:
			goto st157
		case 111:
			goto st163
		case 115:
			goto st168
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr183
		}
		goto st0
tr181:
//line ragel/datetime.rl:5
 pb = p 
	goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line ragel_parse_datetime.go:8123
		if data[p] == 48 {
			goto st128
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st1046
		}
		goto st0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1045
		}
		goto st0
	st1045:
		if p++; p == pe {
			goto _test_eof1045
		}
	st_case_1045:
		switch data[p] {
		case 32:
			goto tr1499
		case 43:
			goto tr1500
		case 44:
			goto tr1501
		case 45:
			goto tr1502
		case 47:
			goto tr1503
		case 84:
			goto tr1504
		case 90:
			goto tr1505
		case 95:
			goto tr1506
		case 116:
			goto tr1506
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1503
			}
		case data[p] >= 65:
			goto tr1503
		}
		goto st0
	st1046:
		if p++; p == pe {
			goto _test_eof1046
		}
	st_case_1046:
		if data[p] == 32 {
			goto tr1507
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1045
			}
		case data[p] >= 45:
			goto tr36
		}
		goto st0
tr1507:
//line ragel/datetime.rl:9

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st129
tr1508:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st129
tr1512:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st129
tr1515:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st129
tr1518:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st129
tr1521:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st129
tr1524:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st129
tr1527:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st129
tr1530:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st129
tr1533:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st129
tr1535:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st129
tr1538:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st129
tr1541:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st129
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
//line ragel_parse_datetime.go:8250
		switch data[p] {
		case 48:
			goto tr37
		case 51:
			goto tr39
		case 109:
			goto st13
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr40
			}
		case data[p] >= 49:
			goto tr38
		}
		goto st0
tr182:
//line ragel/datetime.rl:5
 pb = p 
	goto st1047
	st1047:
		if p++; p == pe {
			goto _test_eof1047
		}
	st_case_1047:
//line ragel_parse_datetime.go:8277
		if data[p] == 32 {
			goto tr1507
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr36
			}
		case data[p] > 50:
			if 51 <= data[p] && data[p] <= 57 {
				goto st128
			}
		default:
			goto st1046
		}
		goto st0
tr183:
//line ragel/datetime.rl:5
 pb = p 
	goto st1048
	st1048:
		if p++; p == pe {
			goto _test_eof1048
		}
	st_case_1048:
//line ragel_parse_datetime.go:8303
		if data[p] == 32 {
			goto tr1507
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st128
			}
		case data[p] >= 45:
			goto tr36
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 112:
			goto st131
		case 117:
			goto st133
		}
		goto st0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		if data[p] == 114 {
			goto st1049
		}
		goto st0
	st1049:
		if p++; p == pe {
			goto _test_eof1049
		}
	st_case_1049:
		switch data[p] {
		case 32:
			goto tr1508
		case 46:
			goto tr1509
		case 105:
			goto st132
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr99
		}
		goto st0
tr1509:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st1050
tr1513:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st1050
tr1516:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st1050
tr1519:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st1050
tr1522:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st1050
tr1525:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st1050
tr1528:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st1050
tr1531:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st1050
tr1534:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st1050
tr1536:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st1050
tr1539:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st1050
tr1542:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st1050
	st1050:
		if p++; p == pe {
			goto _test_eof1050
		}
	st_case_1050:
//line ragel_parse_datetime.go:8407
		switch data[p] {
		case 32:
			goto st129
		case 48:
			goto tr37
		case 51:
			goto tr39
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st8
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr40
			}
		default:
			goto tr38
		}
		goto st0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		if data[p] == 108 {
			goto st1051
		}
		goto st0
	st1051:
		if p++; p == pe {
			goto _test_eof1051
		}
	st_case_1051:
		switch data[p] {
		case 32:
			goto tr1508
		case 46:
			goto tr1509
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr99
		}
		goto st0
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		if data[p] == 103 {
			goto st1052
		}
		goto st0
	st1052:
		if p++; p == pe {
			goto _test_eof1052
		}
	st_case_1052:
		switch data[p] {
		case 32:
			goto tr1512
		case 46:
			goto tr1513
		case 117:
			goto st134
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr105
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
			goto st1053
		}
		goto st0
	st1053:
		if p++; p == pe {
			goto _test_eof1053
		}
	st_case_1053:
		switch data[p] {
		case 32:
			goto tr1512
		case 46:
			goto tr1513
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr105
		}
		goto st0
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		if data[p] == 101 {
			goto st137
		}
		goto st0
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		if data[p] == 99 {
			goto st1054
		}
		goto st0
	st1054:
		if p++; p == pe {
			goto _test_eof1054
		}
	st_case_1054:
		switch data[p] {
		case 32:
			goto tr1515
		case 46:
			goto tr1516
		case 101:
			goto st138
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr112
		}
		goto st0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		if data[p] == 109 {
			goto st139
		}
		goto st0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		if data[p] == 98 {
			goto st140
		}
		goto st0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		if data[p] == 101 {
			goto st141
		}
		goto st0
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		if data[p] == 114 {
			goto st1055
		}
		goto st0
	st1055:
		if p++; p == pe {
			goto _test_eof1055
		}
	st_case_1055:
		switch data[p] {
		case 32:
			goto tr1515
		case 46:
			goto tr1516
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr112
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
		if data[p] == 98 {
			goto st1056
		}
		goto st0
	st1056:
		if p++; p == pe {
			goto _test_eof1056
		}
	st_case_1056:
		switch data[p] {
		case 32:
			goto tr1518
		case 46:
			goto tr1519
		case 114:
			goto st144
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr121
		}
		goto st0
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		if data[p] == 117 {
			goto st145
		}
		goto st0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		if data[p] == 97 {
			goto st146
		}
		goto st0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		if data[p] == 114 {
			goto st147
		}
		goto st0
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		if data[p] == 121 {
			goto st1057
		}
		goto st0
	st1057:
		if p++; p == pe {
			goto _test_eof1057
		}
	st_case_1057:
		switch data[p] {
		case 32:
			goto tr1518
		case 46:
			goto tr1519
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr121
		}
		goto st0
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 97:
			goto st149
		case 117:
			goto st153
		}
		goto st0
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		if data[p] == 110 {
			goto st1058
		}
		goto st0
	st1058:
		if p++; p == pe {
			goto _test_eof1058
		}
	st_case_1058:
		switch data[p] {
		case 32:
			goto tr1521
		case 46:
			goto tr1522
		case 117:
			goto st150
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr131
		}
		goto st0
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		if data[p] == 97 {
			goto st151
		}
		goto st0
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		if data[p] == 114 {
			goto st152
		}
		goto st0
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		if data[p] == 121 {
			goto st1059
		}
		goto st0
	st1059:
		if p++; p == pe {
			goto _test_eof1059
		}
	st_case_1059:
		switch data[p] {
		case 32:
			goto tr1521
		case 46:
			goto tr1522
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr131
		}
		goto st0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 108:
			goto st1060
		case 110:
			goto st1062
		}
		goto st0
	st1060:
		if p++; p == pe {
			goto _test_eof1060
		}
	st_case_1060:
		switch data[p] {
		case 32:
			goto tr1524
		case 46:
			goto tr1525
		case 121:
			goto st1061
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr139
		}
		goto st0
	st1061:
		if p++; p == pe {
			goto _test_eof1061
		}
	st_case_1061:
		switch data[p] {
		case 32:
			goto tr1524
		case 46:
			goto tr1525
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr139
		}
		goto st0
	st1062:
		if p++; p == pe {
			goto _test_eof1062
		}
	st_case_1062:
		switch data[p] {
		case 32:
			goto tr1527
		case 46:
			goto tr1528
		case 101:
			goto st1063
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr142
		}
		goto st0
	st1063:
		if p++; p == pe {
			goto _test_eof1063
		}
	st_case_1063:
		switch data[p] {
		case 32:
			goto tr1527
		case 46:
			goto tr1528
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr142
		}
		goto st0
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		if data[p] == 97 {
			goto st155
		}
		goto st0
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 114:
			goto st1064
		case 121:
			goto st1066
		}
		goto st0
	st1064:
		if p++; p == pe {
			goto _test_eof1064
		}
	st_case_1064:
		switch data[p] {
		case 32:
			goto tr1530
		case 46:
			goto tr1531
		case 99:
			goto st156
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr148
		}
		goto st0
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		if data[p] == 104 {
			goto st1065
		}
		goto st0
	st1065:
		if p++; p == pe {
			goto _test_eof1065
		}
	st_case_1065:
		switch data[p] {
		case 32:
			goto tr1530
		case 46:
			goto tr1531
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr148
		}
		goto st0
	st1066:
		if p++; p == pe {
			goto _test_eof1066
		}
	st_case_1066:
		switch data[p] {
		case 32:
			goto tr1533
		case 46:
			goto tr1534
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr152
		}
		goto st0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		if data[p] == 111 {
			goto st158
		}
		goto st0
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		if data[p] == 118 {
			goto st1067
		}
		goto st0
	st1067:
		if p++; p == pe {
			goto _test_eof1067
		}
	st_case_1067:
		switch data[p] {
		case 32:
			goto tr1535
		case 46:
			goto tr1536
		case 101:
			goto st159
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr156
		}
		goto st0
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		if data[p] == 109 {
			goto st160
		}
		goto st0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		if data[p] == 98 {
			goto st161
		}
		goto st0
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		if data[p] == 101 {
			goto st162
		}
		goto st0
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		if data[p] == 114 {
			goto st1068
		}
		goto st0
	st1068:
		if p++; p == pe {
			goto _test_eof1068
		}
	st_case_1068:
		switch data[p] {
		case 32:
			goto tr1535
		case 46:
			goto tr1536
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr156
		}
		goto st0
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		if data[p] == 99 {
			goto st164
		}
		goto st0
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		if data[p] == 116 {
			goto st1069
		}
		goto st0
	st1069:
		if p++; p == pe {
			goto _test_eof1069
		}
	st_case_1069:
		switch data[p] {
		case 32:
			goto tr1538
		case 46:
			goto tr1539
		case 111:
			goto st165
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr165
		}
		goto st0
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		if data[p] == 98 {
			goto st166
		}
		goto st0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		if data[p] == 101 {
			goto st167
		}
		goto st0
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		if data[p] == 114 {
			goto st1070
		}
		goto st0
	st1070:
		if p++; p == pe {
			goto _test_eof1070
		}
	st_case_1070:
		switch data[p] {
		case 32:
			goto tr1538
		case 46:
			goto tr1539
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr165
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
		if data[p] == 112 {
			goto st1071
		}
		goto st0
	st1071:
		if p++; p == pe {
			goto _test_eof1071
		}
	st_case_1071:
		switch data[p] {
		case 32:
			goto tr1541
		case 46:
			goto tr1542
		case 116:
			goto st1072
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr173
		}
		goto st0
	st1072:
		if p++; p == pe {
			goto _test_eof1072
		}
	st_case_1072:
		switch data[p] {
		case 32:
			goto tr1541
		case 46:
			goto tr1542
		case 101:
			goto st170
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr173
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
			goto st1073
		}
		goto st0
	st1073:
		if p++; p == pe {
			goto _test_eof1073
		}
	st_case_1073:
		switch data[p] {
		case 32:
			goto tr1541
		case 46:
			goto tr1542
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr173
		}
		goto st0
tr1351:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:5
 pb = p 
	goto st174
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
//line ragel_parse_datetime.go:9196
		if 48 <= data[p] && data[p] <= 57 {
			goto st175
		}
		goto st0
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1074
		}
		goto st0
	st1074:
		if p++; p == pe {
			goto _test_eof1074
		}
	st_case_1074:
		switch data[p] {
		case 32:
			goto tr1499
		case 43:
			goto tr1500
		case 44:
			goto tr1501
		case 45:
			goto tr1502
		case 47:
			goto tr1503
		case 84:
			goto tr1504
		case 90:
			goto tr1505
		case 95:
			goto tr1506
		case 116:
			goto tr1506
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1075
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1503
			}
		default:
			goto tr1503
		}
		goto st0
	st1075:
		if p++; p == pe {
			goto _test_eof1075
		}
	st_case_1075:
		switch data[p] {
		case 32:
			goto tr1546
		case 43:
			goto tr1547
		case 44:
			goto tr1548
		case 45:
			goto tr1549
		case 47:
			goto tr1550
		case 84:
			goto tr1551
		case 90:
			goto tr1552
		case 95:
			goto tr1553
		case 116:
			goto tr1553
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1550
			}
		case data[p] >= 65:
			goto tr1550
		}
		goto st0
tr1352:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st176
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
//line ragel_parse_datetime.go:9293
		if data[p] == 185 {
			goto st177
		}
		goto st0
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		if data[p] == 180 {
			goto st178
		}
		goto st0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 48:
			goto tr247
		case 49:
			goto tr248
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr249
		}
		goto st0
tr247:
//line ragel/datetime.rl:5
 pb = p 
	goto st179
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
//line ragel_parse_datetime.go:9331
		if 49 <= data[p] && data[p] <= 57 {
			goto st180
		}
		goto st0
tr249:
//line ragel/datetime.rl:5
 pb = p 
	goto st180
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
//line ragel_parse_datetime.go:9345
		if data[p] == 230 {
			goto tr251
		}
		goto st0
tr251:
//line ragel/datetime.rl:9

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st181
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
//line ragel_parse_datetime.go:9361
		if data[p] == 156 {
			goto st182
		}
		goto st0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		if data[p] == 136 {
			goto st183
		}
		goto st0
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		switch data[p] {
		case 48:
			goto tr254
		case 51:
			goto tr256
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr257
			}
		case data[p] >= 49:
			goto tr255
		}
		goto st0
tr254:
//line ragel/datetime.rl:5
 pb = p 
	goto st184
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
//line ragel_parse_datetime.go:9404
		if 49 <= data[p] && data[p] <= 57 {
			goto st185
		}
		goto st0
tr257:
//line ragel/datetime.rl:5
 pb = p 
	goto st185
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
//line ragel_parse_datetime.go:9418
		switch data[p] {
		case 110:
			goto tr259
		case 114:
			goto tr259
		case 115:
			goto tr260
		case 116:
			goto tr261
		case 230:
			goto tr262
		}
		goto st0
tr259:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st186
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
//line ragel_parse_datetime.go:9448
		if data[p] == 100 {
			goto st187
		}
		goto st0
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		if data[p] == 230 {
			goto st188
		}
		goto st0
tr262:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st188
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
//line ragel_parse_datetime.go:9478
		if data[p] == 151 {
			goto st189
		}
		goto st0
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		if data[p] == 165 {
			goto st1076
		}
		goto st0
	st1076:
		if p++; p == pe {
			goto _test_eof1076
		}
	st_case_1076:
		switch data[p] {
		case 32:
			goto st978
		case 43:
			goto st18
		case 44:
			goto st50
		case 45:
			goto st30
		case 47:
			goto tr1367
		case 84:
			goto tr1554
		case 90:
			goto tr1370
		case 95:
			goto tr1555
		case 116:
			goto tr1555
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
tr260:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st190
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
//line ragel_parse_datetime.go:9542
		if data[p] == 116 {
			goto st187
		}
		goto st0
tr261:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st191
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
//line ragel_parse_datetime.go:9563
		if data[p] == 104 {
			goto st187
		}
		goto st0
tr255:
//line ragel/datetime.rl:5
 pb = p 
	goto st192
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
//line ragel_parse_datetime.go:9577
		switch data[p] {
		case 110:
			goto tr259
		case 114:
			goto tr259
		case 115:
			goto tr260
		case 116:
			goto tr261
		case 230:
			goto tr262
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st185
		}
		goto st0
tr256:
//line ragel/datetime.rl:5
 pb = p 
	goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
//line ragel_parse_datetime.go:9603
		switch data[p] {
		case 110:
			goto tr259
		case 114:
			goto tr259
		case 115:
			goto tr260
		case 116:
			goto tr261
		case 230:
			goto tr262
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st185
		}
		goto st0
tr248:
//line ragel/datetime.rl:5
 pb = p 
	goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
//line ragel_parse_datetime.go:9629
		if data[p] == 230 {
			goto tr251
		}
		if 48 <= data[p] && data[p] <= 50 {
			goto st180
		}
		goto st0
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch data[p] {
		case 32:
			goto tr267
		case 110:
			goto tr269
		case 114:
			goto tr269
		case 115:
			goto tr270
		case 116:
			goto tr271
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] >= 45:
			goto tr268
		}
		goto st0
tr267:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:9

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st196
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
//line ragel_parse_datetime.go:9683
		switch data[p] {
		case 48:
			goto tr272
		case 51:
			goto tr274
		case 65:
			goto st282
		case 68:
			goto st297
		case 70:
			goto st305
		case 74:
			goto st313
		case 77:
			goto st325
		case 78:
			goto st331
		case 79:
			goto st339
		case 83:
			goto st346
		case 97:
			goto st282
		case 100:
			goto st297
		case 102:
			goto st305
		case 106:
			goto st313
		case 109:
			goto st325
		case 110:
			goto st331
		case 111:
			goto st339
		case 115:
			goto st346
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr275
			}
		case data[p] >= 49:
			goto tr273
		}
		goto st0
tr272:
//line ragel/datetime.rl:5
 pb = p 
	goto st197
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
//line ragel_parse_datetime.go:9740
		switch data[p] {
		case 32:
			goto tr284
		case 48:
			goto st202
		}
		switch {
		case data[p] > 47:
			if 49 <= data[p] && data[p] <= 57 {
				goto st203
			}
		case data[p] >= 45:
			goto tr284
		}
		goto st0
tr410:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st198
tr423:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st198
tr431:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st198
tr441:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st198
tr452:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st198
tr461:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st198
tr465:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st198
tr472:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st198
tr477:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st198
tr482:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st198
tr492:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st198
tr501:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st198
tr610:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st198
tr284:
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

	goto st198
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
//line ragel_parse_datetime.go:9839
		if 48 <= data[p] && data[p] <= 57 {
			goto tr287
		}
		goto st0
tr287:
//line ragel/datetime.rl:5
 pb = p 
	goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
//line ragel_parse_datetime.go:9853
		if 48 <= data[p] && data[p] <= 57 {
			goto st200
		}
		goto st0
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		if 48 <= data[p] && data[p] <= 57 {
			goto st201
		}
		goto st0
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1077
		}
		goto st0
	st1077:
		if p++; p == pe {
			goto _test_eof1077
		}
	st_case_1077:
		switch data[p] {
		case 32:
			goto tr1556
		case 43:
			goto tr1557
		case 44:
			goto tr1558
		case 45:
			goto tr1559
		case 47:
			goto tr1560
		case 84:
			goto tr1561
		case 90:
			goto tr1562
		case 95:
			goto tr1563
		case 116:
			goto tr1563
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1560
			}
		case data[p] >= 65:
			goto tr1560
		}
		goto st0
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		if data[p] == 32 {
			goto tr284
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr284
		}
		goto st0
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		switch data[p] {
		case 32:
			goto tr291
		case 110:
			goto tr292
		case 114:
			goto tr292
		case 115:
			goto tr293
		case 116:
			goto tr294
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr284
		}
		goto st0
tr658:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st204
tr291:
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

//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st204
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
//line ragel_parse_datetime.go:9987
		if data[p] == 50 {
			goto tr296
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr297
			}
		case data[p] >= 48:
			goto tr295
		}
		goto st0
tr295:
//line ragel/datetime.rl:5
 pb = p 
	goto st205
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
//line ragel_parse_datetime.go:10009
		switch data[p] {
		case 32:
			goto tr298
		case 58:
			goto tr300
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st230
		}
		goto st0
tr298:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st206
tr341:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st206
tr380:
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

	goto st206
tr363:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st206
tr368:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st206
tr374:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st206
tr391:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
//line ragel_parse_datetime.go:10100
		switch data[p] {
		case 39:
			goto st207
		case 65:
			goto tr305
		case 80:
			goto tr305
		case 97:
			goto tr306
		case 112:
			goto tr306
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr304
		}
		goto st0
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr307
		}
		goto st0
tr307:
//line ragel/datetime.rl:5
 pb = p 
	goto st208
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
//line ragel_parse_datetime.go:10135
		if 48 <= data[p] && data[p] <= 57 {
			goto st1078
		}
		goto st0
	st1078:
		if p++; p == pe {
			goto _test_eof1078
		}
	st_case_1078:
		if data[p] == 32 {
			goto tr1564
		}
		goto st0
tr1588:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st209
tr1564:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st209
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
//line ragel_parse_datetime.go:10166
		switch data[p] {
		case 43:
			goto st210
		case 45:
			goto st220
		case 47:
			goto tr311
		case 90:
			goto tr312
		case 95:
			goto tr311
		case 109:
			goto tr313
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr311
			}
		case data[p] >= 65:
			goto tr311
		}
		goto st0
tr1611:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st210
tr1622:
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

	goto st210
tr1626:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st210
tr1641:
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

	goto st210
tr1634:
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

	goto st210
tr1654:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st210
tr1663:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st210
tr1671:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st210
tr1691:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st210
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
//line ragel_parse_datetime.go:10304
		if data[p] == 50 {
			goto tr315
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr316
			}
		case data[p] >= 48:
			goto tr314
		}
		goto st0
tr314:
//line ragel/datetime.rl:5
 pb = p 
	goto st1079
tr332:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1079
	st1079:
		if p++; p == pe {
			goto _test_eof1079
		}
	st_case_1079:
//line ragel_parse_datetime.go:10332
		switch data[p] {
		case 32:
			goto tr1565
		case 58:
			goto tr1567
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1089
		}
		goto st0
tr1565:
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
	goto st211
tr1580:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st211
tr1583:
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
	goto st211
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
//line ragel_parse_datetime.go:10400
		switch data[p] {
		case 40:
			goto st212
		case 43:
			goto st214
		case 45:
			goto st216
		case 47:
			goto tr320
		case 95:
			goto tr320
		case 109:
			goto tr321
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr320
			}
		case data[p] >= 65:
			goto tr320
		}
		goto st0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		switch data[p] {
		case 32:
			goto st213
		case 43:
			goto st213
		case 45:
			goto st213
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st213
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st213
			}
		default:
			goto st213
		}
		goto st0
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		switch data[p] {
		case 32:
			goto st213
		case 41:
			goto st1080
		case 43:
			goto st213
		case 45:
			goto st213
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st213
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st213
			}
		default:
			goto st213
		}
		goto st0
	st1080:
		if p++; p == pe {
			goto _test_eof1080
		}
	st_case_1080:
		if data[p] == 32 {
			goto tr1568
		}
		goto st0
tr1585:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st214
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
//line ragel_parse_datetime.go:10499
		if data[p] == 50 {
			goto tr325
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr326
			}
		case data[p] >= 48:
			goto tr324
		}
		goto st0
tr324:
//line ragel/datetime.rl:5
 pb = p 
	goto st1081
tr327:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1081
	st1081:
		if p++; p == pe {
			goto _test_eof1081
		}
	st_case_1081:
//line ragel_parse_datetime.go:10527
		switch data[p] {
		case 32:
			goto tr1569
		case 58:
			goto tr1571
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1082
		}
		goto st0
tr1569:
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
	goto st215
tr1573:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st215
tr1576:
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
	goto st215
tr1578:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st215
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
//line ragel_parse_datetime.go:10604
		switch data[p] {
		case 40:
			goto st212
		case 109:
			goto st13
		}
		goto st0
tr326:
//line ragel/datetime.rl:5
 pb = p 
	goto st1082
tr329:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1082
	st1082:
		if p++; p == pe {
			goto _test_eof1082
		}
	st_case_1082:
//line ragel_parse_datetime.go:10627
		switch data[p] {
		case 32:
			goto tr1569
		case 58:
			goto tr1571
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1083
		}
		goto st0
	st1083:
		if p++; p == pe {
			goto _test_eof1083
		}
	st_case_1083:
		if data[p] == 32 {
			goto tr1569
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1083
		}
		goto st0
tr1571:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st1084
	st1084:
		if p++; p == pe {
			goto _test_eof1084
		}
	st_case_1084:
//line ragel_parse_datetime.go:10667
		if data[p] == 32 {
			goto tr1573
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1575
			}
		case data[p] >= 48:
			goto tr1574
		}
		goto st0
tr1574:
//line ragel/datetime.rl:5
 pb = p 
	goto st1085
	st1085:
		if p++; p == pe {
			goto _test_eof1085
		}
	st_case_1085:
//line ragel_parse_datetime.go:10689
		if data[p] == 32 {
			goto tr1576
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1086
		}
		goto st0
tr1575:
//line ragel/datetime.rl:5
 pb = p 
	goto st1086
	st1086:
		if p++; p == pe {
			goto _test_eof1086
		}
	st_case_1086:
//line ragel_parse_datetime.go:10706
		if data[p] == 32 {
			goto tr1576
		}
		goto st0
tr325:
//line ragel/datetime.rl:5
 pb = p 
	goto st1087
tr328:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1087
	st1087:
		if p++; p == pe {
			goto _test_eof1087
		}
	st_case_1087:
//line ragel_parse_datetime.go:10726
		switch data[p] {
		case 32:
			goto tr1569
		case 58:
			goto tr1571
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st1083
			}
		case data[p] >= 48:
			goto st1082
		}
		goto st0
tr1586:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st216
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
//line ragel_parse_datetime.go:10754
		if data[p] == 50 {
			goto tr328
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr329
			}
		case data[p] >= 48:
			goto tr327
		}
		goto st0
tr320:
//line ragel/datetime.rl:5
 pb = p 
	goto st217
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
//line ragel_parse_datetime.go:10776
		switch data[p] {
		case 47:
			goto st218
		case 95:
			goto st218
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st218
			}
		case data[p] >= 65:
			goto st218
		}
		goto st0
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 47:
			goto st1088
		case 95:
			goto st1088
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1088
			}
		case data[p] >= 65:
			goto st1088
		}
		goto st0
	st1088:
		if p++; p == pe {
			goto _test_eof1088
		}
	st_case_1088:
		switch data[p] {
		case 32:
			goto tr1578
		case 47:
			goto st1088
		case 95:
			goto st1088
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1088
			}
		case data[p] >= 65:
			goto st1088
		}
		goto st0
tr321:
//line ragel/datetime.rl:5
 pb = p 
	goto st219
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
//line ragel_parse_datetime.go:10843
		switch data[p] {
		case 47:
			goto st218
		case 61:
			goto st14
		case 95:
			goto st218
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st218
			}
		case data[p] >= 65:
			goto st218
		}
		goto st0
tr316:
//line ragel/datetime.rl:5
 pb = p 
	goto st1089
tr334:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1089
	st1089:
		if p++; p == pe {
			goto _test_eof1089
		}
	st_case_1089:
//line ragel_parse_datetime.go:10876
		switch data[p] {
		case 32:
			goto tr1565
		case 58:
			goto tr1567
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1090
		}
		goto st0
	st1090:
		if p++; p == pe {
			goto _test_eof1090
		}
	st_case_1090:
		if data[p] == 32 {
			goto tr1565
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1090
		}
		goto st0
tr1567:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st1091
	st1091:
		if p++; p == pe {
			goto _test_eof1091
		}
	st_case_1091:
//line ragel_parse_datetime.go:10916
		if data[p] == 32 {
			goto tr1580
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1582
			}
		case data[p] >= 48:
			goto tr1581
		}
		goto st0
tr1581:
//line ragel/datetime.rl:5
 pb = p 
	goto st1092
	st1092:
		if p++; p == pe {
			goto _test_eof1092
		}
	st_case_1092:
//line ragel_parse_datetime.go:10938
		if data[p] == 32 {
			goto tr1583
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1093
		}
		goto st0
tr1582:
//line ragel/datetime.rl:5
 pb = p 
	goto st1093
	st1093:
		if p++; p == pe {
			goto _test_eof1093
		}
	st_case_1093:
//line ragel_parse_datetime.go:10955
		if data[p] == 32 {
			goto tr1583
		}
		goto st0
tr315:
//line ragel/datetime.rl:5
 pb = p 
	goto st1094
tr333:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1094
	st1094:
		if p++; p == pe {
			goto _test_eof1094
		}
	st_case_1094:
//line ragel_parse_datetime.go:10975
		switch data[p] {
		case 32:
			goto tr1565
		case 58:
			goto tr1567
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st1090
			}
		case data[p] >= 48:
			goto st1089
		}
		goto st0
tr1612:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st220
tr1623:
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

	goto st220
tr1627:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st220
tr1642:
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

	goto st220
tr1636:
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

	goto st220
tr1655:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st220
tr1664:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st220
tr1673:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st220
tr1693:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st220
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
//line ragel_parse_datetime.go:11105
		if data[p] == 50 {
			goto tr333
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr334
			}
		case data[p] >= 48:
			goto tr332
		}
		goto st0
tr311:
//line ragel/datetime.rl:5
 pb = p 
	goto st221
tr1613:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st221
tr1628:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st221
tr1656:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st221
tr1665:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st221
tr1674:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st221
tr1643:
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
	goto st221
tr1694:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st221
tr1637:
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
	goto st221
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
//line ragel_parse_datetime.go:11227
		switch data[p] {
		case 47:
			goto st222
		case 95:
			goto st222
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st222
			}
		case data[p] >= 65:
			goto st222
		}
		goto st0
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 47:
			goto st1095
		case 95:
			goto st1095
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1095
			}
		case data[p] >= 65:
			goto st1095
		}
		goto st0
tr1624:
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
	goto st1095
	st1095:
		if p++; p == pe {
			goto _test_eof1095
		}
	st_case_1095:
//line ragel_parse_datetime.go:11295
		switch data[p] {
		case 32:
			goto tr1578
		case 43:
			goto tr1585
		case 45:
			goto tr1586
		case 47:
			goto st1095
		case 95:
			goto st1095
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1095
			}
		case data[p] >= 65:
			goto st1095
		}
		goto st0
tr312:
//line ragel/datetime.rl:5
 pb = p 
	goto st1096
tr1617:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st1096
tr1631:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st1096
tr1660:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st1096
tr1668:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st1096
tr1677:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st1096
tr1645:
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
	goto st1096
tr1696:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st1096
tr1639:
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
	goto st1096
	st1096:
		if p++; p == pe {
			goto _test_eof1096
		}
	st_case_1096:
//line ragel_parse_datetime.go:11426
		switch data[p] {
		case 32:
			goto tr1573
		case 47:
			goto st222
		case 95:
			goto st222
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st222
			}
		case data[p] >= 65:
			goto st222
		}
		goto st0
tr313:
//line ragel/datetime.rl:5
 pb = p 
	goto st223
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
//line ragel_parse_datetime.go:11453
		switch data[p] {
		case 47:
			goto st222
		case 61:
			goto st14
		case 95:
			goto st222
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st222
			}
		case data[p] >= 65:
			goto st222
		}
		goto st0
tr304:
//line ragel/datetime.rl:5
 pb = p 
	goto st224
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
//line ragel_parse_datetime.go:11480
		if 48 <= data[p] && data[p] <= 57 {
			goto st1097
		}
		goto st0
	st1097:
		if p++; p == pe {
			goto _test_eof1097
		}
	st_case_1097:
		if data[p] == 32 {
			goto tr1564
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st225
		}
		goto st0
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1098
		}
		goto st0
	st1098:
		if p++; p == pe {
			goto _test_eof1098
		}
	st_case_1098:
		if data[p] == 32 {
			goto tr1588
		}
		goto st0
tr305:
//line ragel/datetime.rl:5
 pb = p 
	goto st226
tr301:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st226
tr344:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st226
tr366:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st226
tr370:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st226
tr377:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st226
tr382:
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
	goto st226
tr393:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st226
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
//line ragel_parse_datetime.go:11605
		if data[p] == 77 {
			goto st227
		}
		goto st0
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		if data[p] == 32 {
			goto tr340
		}
		goto st0
tr340:
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

	goto st228
tr349:
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

	goto st228
tr359:
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

	goto st228
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
//line ragel_parse_datetime.go:11697
		if data[p] == 39 {
			goto st207
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr304
		}
		goto st0
tr306:
//line ragel/datetime.rl:5
 pb = p 
	goto st229
tr302:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st229
tr345:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st229
tr367:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st229
tr371:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st229
tr378:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st229
tr383:
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
	goto st229
tr394:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line ragel_parse_datetime.go:11795
		if data[p] == 109 {
			goto st227
		}
		goto st0
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 32:
			goto tr341
		case 58:
			goto tr343
		case 65:
			goto tr344
		case 80:
			goto tr344
		case 97:
			goto tr345
		case 112:
			goto tr345
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st231
		}
		goto st0
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1099
		}
		goto st0
	st1099:
		if p++; p == pe {
			goto _test_eof1099
		}
	st_case_1099:
		switch data[p] {
		case 32:
			goto tr1589
		case 43:
			goto tr1557
		case 44:
			goto tr1590
		case 45:
			goto tr1559
		case 46:
			goto tr360
		case 47:
			goto tr1560
		case 84:
			goto tr1561
		case 90:
			goto tr1562
		case 95:
			goto tr1563
		case 116:
			goto tr1563
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st244
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1560
			}
		default:
			goto tr1560
		}
		goto st0
tr1589:
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

	goto st1100
	st1100:
		if p++; p == pe {
			goto _test_eof1100
		}
	st_case_1100:
//line ragel_parse_datetime.go:11898
		switch data[p] {
		case 32:
			goto st10
		case 39:
			goto st207
		case 43:
			goto st18
		case 45:
			goto st30
		case 47:
			goto tr1367
		case 50:
			goto tr1592
		case 65:
			goto tr1368
		case 66:
			goto tr1369
		case 90:
			goto tr1370
		case 95:
			goto tr1367
		case 97:
			goto tr1371
		case 109:
			goto tr1372
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr1591
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1367
				}
			case data[p] >= 67:
				goto tr1367
			}
		default:
			goto tr1593
		}
		goto st0
tr1591:
//line ragel/datetime.rl:5
 pb = p 
	goto st1101
	st1101:
		if p++; p == pe {
			goto _test_eof1101
		}
	st_case_1101:
//line ragel_parse_datetime.go:11952
		switch data[p] {
		case 32:
			goto tr1399
		case 43:
			goto tr1400
		case 45:
			goto tr1401
		case 47:
			goto tr1402
		case 58:
			goto tr1404
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1406
		case 95:
			goto tr1402
		case 97:
			goto tr1407
		case 112:
			goto tr1407
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1102
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1402
			}
		default:
			goto tr1402
		}
		goto st0
	st1102:
		if p++; p == pe {
			goto _test_eof1102
		}
	st_case_1102:
		switch data[p] {
		case 32:
			goto tr1595
		case 43:
			goto tr1417
		case 45:
			goto tr1418
		case 47:
			goto tr1419
		case 58:
			goto tr1420
		case 65:
			goto tr1421
		case 80:
			goto tr1421
		case 90:
			goto tr1422
		case 95:
			goto tr1419
		case 97:
			goto tr1423
		case 112:
			goto tr1423
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st232
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1419
			}
		default:
			goto tr1419
		}
		goto st0
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1103
		}
		goto st0
	st1103:
		if p++; p == pe {
			goto _test_eof1103
		}
	st_case_1103:
		switch data[p] {
		case 32:
			goto tr1597
		case 43:
			goto tr1425
		case 45:
			goto tr1427
		case 47:
			goto tr1428
		case 90:
			goto tr1430
		case 95:
			goto tr1428
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1426
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1428
				}
			case data[p] >= 65:
				goto tr1428
			}
		default:
			goto st41
		}
		goto st0
tr1592:
//line ragel/datetime.rl:5
 pb = p 
	goto st1104
	st1104:
		if p++; p == pe {
			goto _test_eof1104
		}
	st_case_1104:
//line ragel_parse_datetime.go:12087
		switch data[p] {
		case 32:
			goto tr1399
		case 43:
			goto tr1400
		case 45:
			goto tr1401
		case 47:
			goto tr1402
		case 58:
			goto tr1404
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1406
		case 95:
			goto tr1402
		case 97:
			goto tr1407
		case 112:
			goto tr1407
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st1102
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1402
				}
			case data[p] >= 66:
				goto tr1402
			}
		default:
			goto st1105
		}
		goto st0
	st1105:
		if p++; p == pe {
			goto _test_eof1105
		}
	st_case_1105:
		if data[p] == 32 {
			goto tr1564
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st232
		}
		goto st0
tr1593:
//line ragel/datetime.rl:5
 pb = p 
	goto st1106
	st1106:
		if p++; p == pe {
			goto _test_eof1106
		}
	st_case_1106:
//line ragel_parse_datetime.go:12151
		switch data[p] {
		case 32:
			goto tr1399
		case 43:
			goto tr1400
		case 45:
			goto tr1401
		case 47:
			goto tr1402
		case 58:
			goto tr1404
		case 65:
			goto tr1405
		case 80:
			goto tr1405
		case 90:
			goto tr1406
		case 95:
			goto tr1402
		case 97:
			goto tr1407
		case 112:
			goto tr1407
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1105
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1402
			}
		default:
			goto tr1402
		}
		goto st0
tr1590:
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

	goto st233
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
//line ragel_parse_datetime.go:12215
		if data[p] == 32 {
			goto st48
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr348
		}
		goto st0
tr348:
//line ragel/datetime.rl:5
 pb = p 
	goto st234
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
//line ragel_parse_datetime.go:12232
		if data[p] == 32 {
			goto tr349
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st235
		}
		goto st0
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		if data[p] == 32 {
			goto tr349
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st236
		}
		goto st0
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		if data[p] == 32 {
			goto tr349
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st237
		}
		goto st0
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		if data[p] == 32 {
			goto tr349
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st238
		}
		goto st0
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		if data[p] == 32 {
			goto tr349
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st239
		}
		goto st0
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		if data[p] == 32 {
			goto tr349
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st240
		}
		goto st0
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		if data[p] == 32 {
			goto tr349
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st241
		}
		goto st0
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		if data[p] == 32 {
			goto tr349
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st242
		}
		goto st0
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		if data[p] == 32 {
			goto tr349
		}
		goto st0
tr360:
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

	goto st243
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
//line ragel_parse_datetime.go:12355
		if 48 <= data[p] && data[p] <= 57 {
			goto tr348
		}
		goto st0
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		if 48 <= data[p] && data[p] <= 57 {
			goto st245
		}
		goto st0
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		switch data[p] {
		case 32:
			goto tr359
		case 44:
			goto tr360
		case 46:
			goto tr360
		}
		goto st0
tr300:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st246
tr343:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st246
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
//line ragel_parse_datetime.go:12400
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr362
			}
		case data[p] >= 48:
			goto tr361
		}
		goto st0
tr361:
//line ragel/datetime.rl:5
 pb = p 
	goto st247
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
//line ragel_parse_datetime.go:12419
		switch data[p] {
		case 32:
			goto tr363
		case 58:
			goto tr365
		case 65:
			goto tr366
		case 80:
			goto tr366
		case 97:
			goto tr367
		case 112:
			goto tr367
		}
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr368
		case 58:
			goto tr369
		case 65:
			goto tr370
		case 80:
			goto tr370
		case 97:
			goto tr371
		case 112:
			goto tr371
		}
		goto st0
tr365:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st249
tr369:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st249
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
//line ragel_parse_datetime.go:12475
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr373
			}
		case data[p] >= 48:
			goto tr372
		}
		goto st0
tr372:
//line ragel/datetime.rl:5
 pb = p 
	goto st250
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
//line ragel_parse_datetime.go:12494
		switch data[p] {
		case 32:
			goto tr374
		case 44:
			goto tr375
		case 46:
			goto tr375
		case 65:
			goto tr377
		case 80:
			goto tr377
		case 97:
			goto tr378
		case 112:
			goto tr378
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st261
		}
		goto st0
tr375:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st251
tr392:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st251
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
//line ragel_parse_datetime.go:12532
		if 48 <= data[p] && data[p] <= 57 {
			goto tr379
		}
		goto st0
tr379:
//line ragel/datetime.rl:5
 pb = p 
	goto st252
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
//line ragel_parse_datetime.go:12546
		switch data[p] {
		case 32:
			goto tr380
		case 65:
			goto tr382
		case 80:
			goto tr382
		case 97:
			goto tr383
		case 112:
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st253
		}
		goto st0
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		switch data[p] {
		case 32:
			goto tr380
		case 65:
			goto tr382
		case 80:
			goto tr382
		case 97:
			goto tr383
		case 112:
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st254
		}
		goto st0
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		switch data[p] {
		case 32:
			goto tr380
		case 65:
			goto tr382
		case 80:
			goto tr382
		case 97:
			goto tr383
		case 112:
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr380
		case 65:
			goto tr382
		case 80:
			goto tr382
		case 97:
			goto tr383
		case 112:
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr380
		case 65:
			goto tr382
		case 80:
			goto tr382
		case 97:
			goto tr383
		case 112:
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st257
		}
		goto st0
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		switch data[p] {
		case 32:
			goto tr380
		case 65:
			goto tr382
		case 80:
			goto tr382
		case 97:
			goto tr383
		case 112:
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr380
		case 65:
			goto tr382
		case 80:
			goto tr382
		case 97:
			goto tr383
		case 112:
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr380
		case 65:
			goto tr382
		case 80:
			goto tr382
		case 97:
			goto tr383
		case 112:
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st260
		}
		goto st0
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 32:
			goto tr380
		case 65:
			goto tr382
		case 80:
			goto tr382
		case 97:
			goto tr383
		case 112:
			goto tr383
		}
		goto st0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		switch data[p] {
		case 32:
			goto tr391
		case 44:
			goto tr392
		case 46:
			goto tr392
		case 65:
			goto tr393
		case 80:
			goto tr393
		case 97:
			goto tr394
		case 112:
			goto tr394
		}
		goto st0
tr373:
//line ragel/datetime.rl:5
 pb = p 
	goto st262
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
//line ragel_parse_datetime.go:12759
		switch data[p] {
		case 32:
			goto tr374
		case 44:
			goto tr375
		case 46:
			goto tr375
		case 65:
			goto tr377
		case 80:
			goto tr377
		case 97:
			goto tr378
		case 112:
			goto tr378
		}
		goto st0
tr362:
//line ragel/datetime.rl:5
 pb = p 
	goto st263
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
//line ragel_parse_datetime.go:12786
		switch data[p] {
		case 32:
			goto tr363
		case 58:
			goto tr365
		case 65:
			goto tr366
		case 80:
			goto tr366
		case 97:
			goto tr367
		case 112:
			goto tr367
		}
		goto st0
tr296:
//line ragel/datetime.rl:5
 pb = p 
	goto st264
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
//line ragel_parse_datetime.go:12811
		switch data[p] {
		case 32:
			goto tr298
		case 58:
			goto tr300
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st265
			}
		case data[p] >= 48:
			goto st230
		}
		goto st0
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		if 48 <= data[p] && data[p] <= 57 {
			goto st231
		}
		goto st0
tr297:
//line ragel/datetime.rl:5
 pb = p 
	goto st266
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
//line ragel_parse_datetime.go:12853
		switch data[p] {
		case 32:
			goto tr298
		case 58:
			goto tr300
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st265
		}
		goto st0
tr292:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st267
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
//line ragel_parse_datetime.go:12888
		if data[p] == 100 {
			goto st268
		}
		goto st0
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		if data[p] == 32 {
			goto st269
		}
		goto st0
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
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
tr398:
//line ragel/datetime.rl:5
 pb = p 
	goto st270
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
//line ragel_parse_datetime.go:12928
		switch data[p] {
		case 32:
			goto tr298
		case 58:
			goto tr300
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st271
		}
		goto st0
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		switch data[p] {
		case 32:
			goto tr341
		case 58:
			goto tr343
		case 65:
			goto tr344
		case 80:
			goto tr344
		case 97:
			goto tr345
		case 112:
			goto tr345
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st272
		}
		goto st0
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr359
		case 44:
			goto tr360
		case 46:
			goto tr360
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st244
		}
		goto st0
tr399:
//line ragel/datetime.rl:5
 pb = p 
	goto st274
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
//line ragel_parse_datetime.go:13005
		switch data[p] {
		case 32:
			goto tr298
		case 58:
			goto tr300
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st275
			}
		case data[p] >= 48:
			goto st271
		}
		goto st0
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		if 48 <= data[p] && data[p] <= 57 {
			goto st272
		}
		goto st0
tr400:
//line ragel/datetime.rl:5
 pb = p 
	goto st276
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
//line ragel_parse_datetime.go:13047
		switch data[p] {
		case 32:
			goto tr298
		case 58:
			goto tr300
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st275
		}
		goto st0
tr293:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st277
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
//line ragel_parse_datetime.go:13082
		if data[p] == 116 {
			goto st268
		}
		goto st0
tr294:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st278
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
//line ragel_parse_datetime.go:13103
		if data[p] == 104 {
			goto st268
		}
		goto st0
tr273:
//line ragel/datetime.rl:5
 pb = p 
	goto st279
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
//line ragel_parse_datetime.go:13117
		switch data[p] {
		case 32:
			goto tr291
		case 110:
			goto tr292
		case 114:
			goto tr292
		case 115:
			goto tr293
		case 116:
			goto tr294
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st203
			}
		case data[p] >= 45:
			goto tr284
		}
		goto st0
tr274:
//line ragel/datetime.rl:5
 pb = p 
	goto st280
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
//line ragel_parse_datetime.go:13148
		switch data[p] {
		case 32:
			goto tr291
		case 110:
			goto tr292
		case 114:
			goto tr292
		case 115:
			goto tr293
		case 116:
			goto tr294
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr284
			}
		case data[p] > 49:
			if 50 <= data[p] && data[p] <= 57 {
				goto st202
			}
		default:
			goto st203
		}
		goto st0
tr275:
//line ragel/datetime.rl:5
 pb = p 
	goto st281
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
//line ragel_parse_datetime.go:13183
		switch data[p] {
		case 32:
			goto tr291
		case 110:
			goto tr292
		case 114:
			goto tr292
		case 115:
			goto tr293
		case 116:
			goto tr294
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st202
			}
		case data[p] >= 45:
			goto tr284
		}
		goto st0
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		switch data[p] {
		case 112:
			goto st283
		case 117:
			goto st292
		}
		goto st0
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		if data[p] == 114 {
			goto st284
		}
		goto st0
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		switch data[p] {
		case 32:
			goto tr409
		case 46:
			goto tr411
		case 105:
			goto st290
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr410
		}
		goto st0
tr409:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st285
tr422:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st285
tr430:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st285
tr440:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st285
tr451:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st285
tr460:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st285
tr464:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st285
tr471:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st285
tr476:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st285
tr481:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st285
tr491:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st285
tr500:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st285
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
//line ragel_parse_datetime.go:13296
		if data[p] == 39 {
			goto st286
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr414
		}
		goto st0
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr415
		}
		goto st0
tr415:
//line ragel/datetime.rl:5
 pb = p 
	goto st287
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
//line ragel_parse_datetime.go:13322
		if 48 <= data[p] && data[p] <= 57 {
			goto st1107
		}
		goto st0
	st1107:
		if p++; p == pe {
			goto _test_eof1107
		}
	st_case_1107:
		switch data[p] {
		case 32:
			goto tr1599
		case 43:
			goto tr1600
		case 44:
			goto tr1601
		case 45:
			goto tr1602
		case 47:
			goto tr1603
		case 84:
			goto tr1604
		case 90:
			goto tr1605
		case 95:
			goto tr1606
		case 116:
			goto tr1606
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1603
			}
		case data[p] >= 65:
			goto tr1603
		}
		goto st0
tr414:
//line ragel/datetime.rl:5
 pb = p 
	goto st288
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
//line ragel_parse_datetime.go:13370
		if 48 <= data[p] && data[p] <= 57 {
			goto st1108
		}
		goto st0
	st1108:
		if p++; p == pe {
			goto _test_eof1108
		}
	st_case_1108:
		switch data[p] {
		case 32:
			goto tr1599
		case 43:
			goto tr1600
		case 44:
			goto tr1601
		case 45:
			goto tr1602
		case 47:
			goto tr1603
		case 84:
			goto tr1604
		case 90:
			goto tr1605
		case 95:
			goto tr1606
		case 116:
			goto tr1606
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st201
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1603
			}
		default:
			goto tr1603
		}
		goto st0
tr411:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st289
tr424:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st289
tr432:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st289
tr442:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st289
tr453:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st289
tr462:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st289
tr466:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st289
tr473:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st289
tr478:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st289
tr483:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st289
tr493:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st289
tr502:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st289
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
//line ragel_parse_datetime.go:13466
		if data[p] == 32 {
			goto st285
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr287
			}
		case data[p] >= 45:
			goto st198
		}
		goto st0
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
		if data[p] == 108 {
			goto st291
		}
		goto st0
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		switch data[p] {
		case 32:
			goto tr409
		case 46:
			goto tr411
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr410
		}
		goto st0
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		if data[p] == 103 {
			goto st293
		}
		goto st0
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
		switch data[p] {
		case 32:
			goto tr422
		case 46:
			goto tr424
		case 117:
			goto st294
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr423
		}
		goto st0
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		if data[p] == 115 {
			goto st295
		}
		goto st0
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		if data[p] == 116 {
			goto st296
		}
		goto st0
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		switch data[p] {
		case 32:
			goto tr422
		case 46:
			goto tr424
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr423
		}
		goto st0
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		if data[p] == 101 {
			goto st298
		}
		goto st0
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		if data[p] == 99 {
			goto st299
		}
		goto st0
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 32:
			goto tr430
		case 46:
			goto tr432
		case 101:
			goto st300
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr431
		}
		goto st0
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		if data[p] == 109 {
			goto st301
		}
		goto st0
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		if data[p] == 98 {
			goto st302
		}
		goto st0
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		if data[p] == 101 {
			goto st303
		}
		goto st0
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		if data[p] == 114 {
			goto st304
		}
		goto st0
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		switch data[p] {
		case 32:
			goto tr430
		case 46:
			goto tr432
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr431
		}
		goto st0
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		if data[p] == 101 {
			goto st306
		}
		goto st0
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		if data[p] == 98 {
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
			goto tr440
		case 46:
			goto tr442
		case 114:
			goto st308
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr441
		}
		goto st0
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		if data[p] == 117 {
			goto st309
		}
		goto st0
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		if data[p] == 97 {
			goto st310
		}
		goto st0
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		if data[p] == 114 {
			goto st311
		}
		goto st0
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
		if data[p] == 121 {
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
			goto tr440
		case 46:
			goto tr442
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr441
		}
		goto st0
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 97:
			goto st314
		case 117:
			goto st320
		}
		goto st0
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		if data[p] == 110 {
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
			goto tr451
		case 46:
			goto tr453
		case 117:
			goto st316
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr452
		}
		goto st0
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		if data[p] == 97 {
			goto st317
		}
		goto st0
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		if data[p] == 114 {
			goto st318
		}
		goto st0
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		if data[p] == 121 {
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
			goto tr451
		case 46:
			goto tr453
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr452
		}
		goto st0
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		switch data[p] {
		case 108:
			goto st321
		case 110:
			goto st323
		}
		goto st0
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		switch data[p] {
		case 32:
			goto tr460
		case 46:
			goto tr462
		case 121:
			goto st322
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr461
		}
		goto st0
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		switch data[p] {
		case 32:
			goto tr460
		case 46:
			goto tr462
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr461
		}
		goto st0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 32:
			goto tr464
		case 46:
			goto tr466
		case 101:
			goto st324
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr465
		}
		goto st0
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		switch data[p] {
		case 32:
			goto tr464
		case 46:
			goto tr466
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr465
		}
		goto st0
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
		if data[p] == 97 {
			goto st326
		}
		goto st0
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		switch data[p] {
		case 114:
			goto st327
		case 121:
			goto st330
		}
		goto st0
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		switch data[p] {
		case 32:
			goto tr471
		case 46:
			goto tr473
		case 99:
			goto st328
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr472
		}
		goto st0
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		if data[p] == 104 {
			goto st329
		}
		goto st0
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		switch data[p] {
		case 32:
			goto tr471
		case 46:
			goto tr473
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr472
		}
		goto st0
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		switch data[p] {
		case 32:
			goto tr476
		case 46:
			goto tr478
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr477
		}
		goto st0
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		if data[p] == 111 {
			goto st332
		}
		goto st0
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		if data[p] == 118 {
			goto st333
		}
		goto st0
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 32:
			goto tr481
		case 46:
			goto tr483
		case 101:
			goto st334
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr482
		}
		goto st0
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		if data[p] == 109 {
			goto st335
		}
		goto st0
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		if data[p] == 98 {
			goto st336
		}
		goto st0
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		if data[p] == 101 {
			goto st337
		}
		goto st0
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		if data[p] == 114 {
			goto st338
		}
		goto st0
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		switch data[p] {
		case 32:
			goto tr481
		case 46:
			goto tr483
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr482
		}
		goto st0
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		if data[p] == 99 {
			goto st340
		}
		goto st0
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		if data[p] == 116 {
			goto st341
		}
		goto st0
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		switch data[p] {
		case 32:
			goto tr491
		case 46:
			goto tr493
		case 111:
			goto st342
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr492
		}
		goto st0
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		if data[p] == 98 {
			goto st343
		}
		goto st0
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
		if data[p] == 101 {
			goto st344
		}
		goto st0
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		if data[p] == 114 {
			goto st345
		}
		goto st0
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
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
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		if data[p] == 101 {
			goto st347
		}
		goto st0
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		if data[p] == 112 {
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
			goto tr500
		case 46:
			goto tr502
		case 116:
			goto st349
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr501
		}
		goto st0
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		switch data[p] {
		case 32:
			goto tr500
		case 46:
			goto tr502
		case 101:
			goto st350
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr501
		}
		goto st0
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		if data[p] == 109 {
			goto st351
		}
		goto st0
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		if data[p] == 98 {
			goto st352
		}
		goto st0
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		if data[p] == 101 {
			goto st353
		}
		goto st0
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		if data[p] == 114 {
			goto st354
		}
		goto st0
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		switch data[p] {
		case 32:
			goto tr500
		case 46:
			goto tr502
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr501
		}
		goto st0
tr268:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:9

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st355
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
//line ragel_parse_datetime.go:14253
		switch data[p] {
		case 48:
			goto tr272
		case 51:
			goto tr274
		case 65:
			goto st356
		case 68:
			goto st367
		case 70:
			goto st375
		case 74:
			goto st383
		case 77:
			goto st395
		case 78:
			goto st401
		case 79:
			goto st409
		case 83:
			goto st416
		case 97:
			goto st356
		case 100:
			goto st367
		case 102:
			goto st375
		case 106:
			goto st383
		case 109:
			goto st395
		case 110:
			goto st401
		case 111:
			goto st409
		case 115:
			goto st416
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr275
			}
		case data[p] >= 49:
			goto tr273
		}
		goto st0
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		switch data[p] {
		case 112:
			goto st357
		case 117:
			goto st362
		}
		goto st0
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		if data[p] == 114 {
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
			goto tr410
		case 46:
			goto tr520
		case 105:
			goto st360
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr410
		}
		goto st0
tr520:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st359
tr524:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st359
tr530:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st359
tr538:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st359
tr547:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st359
tr554:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st359
tr556:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st359
tr561:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st359
tr564:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st359
tr567:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st359
tr575:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st359
tr582:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st359
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
//line ragel_parse_datetime.go:14392
		if data[p] == 32 {
			goto st198
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr287
			}
		case data[p] >= 45:
			goto st198
		}
		goto st0
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		if data[p] == 108 {
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
			goto tr410
		case 46:
			goto tr520
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr410
		}
		goto st0
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		if data[p] == 103 {
			goto st363
		}
		goto st0
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		switch data[p] {
		case 32:
			goto tr423
		case 46:
			goto tr524
		case 117:
			goto st364
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr423
		}
		goto st0
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		if data[p] == 115 {
			goto st365
		}
		goto st0
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		if data[p] == 116 {
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
			goto tr423
		case 46:
			goto tr524
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr423
		}
		goto st0
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		if data[p] == 101 {
			goto st368
		}
		goto st0
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		if data[p] == 99 {
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
			goto tr431
		case 46:
			goto tr530
		case 101:
			goto st370
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr431
		}
		goto st0
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		if data[p] == 109 {
			goto st371
		}
		goto st0
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		if data[p] == 98 {
			goto st372
		}
		goto st0
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		if data[p] == 101 {
			goto st373
		}
		goto st0
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		if data[p] == 114 {
			goto st374
		}
		goto st0
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		switch data[p] {
		case 32:
			goto tr431
		case 46:
			goto tr530
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr431
		}
		goto st0
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		if data[p] == 101 {
			goto st376
		}
		goto st0
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		if data[p] == 98 {
			goto st377
		}
		goto st0
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		switch data[p] {
		case 32:
			goto tr441
		case 46:
			goto tr538
		case 114:
			goto st378
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr441
		}
		goto st0
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		if data[p] == 117 {
			goto st379
		}
		goto st0
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		if data[p] == 97 {
			goto st380
		}
		goto st0
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		if data[p] == 114 {
			goto st381
		}
		goto st0
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		if data[p] == 121 {
			goto st382
		}
		goto st0
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		switch data[p] {
		case 32:
			goto tr441
		case 46:
			goto tr538
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr441
		}
		goto st0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 97:
			goto st384
		case 117:
			goto st390
		}
		goto st0
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		if data[p] == 110 {
			goto st385
		}
		goto st0
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		switch data[p] {
		case 32:
			goto tr452
		case 46:
			goto tr547
		case 117:
			goto st386
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr452
		}
		goto st0
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		if data[p] == 97 {
			goto st387
		}
		goto st0
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		if data[p] == 114 {
			goto st388
		}
		goto st0
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
		if data[p] == 121 {
			goto st389
		}
		goto st0
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 32:
			goto tr452
		case 46:
			goto tr547
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr452
		}
		goto st0
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		switch data[p] {
		case 108:
			goto st391
		case 110:
			goto st393
		}
		goto st0
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
		switch data[p] {
		case 32:
			goto tr461
		case 46:
			goto tr554
		case 121:
			goto st392
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr461
		}
		goto st0
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		switch data[p] {
		case 32:
			goto tr461
		case 46:
			goto tr554
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr461
		}
		goto st0
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		switch data[p] {
		case 32:
			goto tr465
		case 46:
			goto tr556
		case 101:
			goto st394
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr465
		}
		goto st0
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		switch data[p] {
		case 32:
			goto tr465
		case 46:
			goto tr556
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr465
		}
		goto st0
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		if data[p] == 97 {
			goto st396
		}
		goto st0
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		switch data[p] {
		case 114:
			goto st397
		case 121:
			goto st400
		}
		goto st0
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		switch data[p] {
		case 32:
			goto tr472
		case 46:
			goto tr561
		case 99:
			goto st398
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr472
		}
		goto st0
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		if data[p] == 104 {
			goto st399
		}
		goto st0
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		switch data[p] {
		case 32:
			goto tr472
		case 46:
			goto tr561
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr472
		}
		goto st0
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		switch data[p] {
		case 32:
			goto tr477
		case 46:
			goto tr564
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr477
		}
		goto st0
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		if data[p] == 111 {
			goto st402
		}
		goto st0
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
		if data[p] == 118 {
			goto st403
		}
		goto st0
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		switch data[p] {
		case 32:
			goto tr482
		case 46:
			goto tr567
		case 101:
			goto st404
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr482
		}
		goto st0
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		if data[p] == 109 {
			goto st405
		}
		goto st0
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		if data[p] == 98 {
			goto st406
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
		if data[p] == 114 {
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
			goto tr482
		case 46:
			goto tr567
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr482
		}
		goto st0
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		if data[p] == 99 {
			goto st410
		}
		goto st0
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		if data[p] == 116 {
			goto st411
		}
		goto st0
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		switch data[p] {
		case 32:
			goto tr492
		case 46:
			goto tr575
		case 111:
			goto st412
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr492
		}
		goto st0
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		if data[p] == 98 {
			goto st413
		}
		goto st0
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
		if data[p] == 101 {
			goto st414
		}
		goto st0
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		if data[p] == 114 {
			goto st415
		}
		goto st0
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		switch data[p] {
		case 32:
			goto tr492
		case 46:
			goto tr575
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr492
		}
		goto st0
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		if data[p] == 101 {
			goto st417
		}
		goto st0
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		if data[p] == 112 {
			goto st418
		}
		goto st0
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		switch data[p] {
		case 32:
			goto tr501
		case 46:
			goto tr582
		case 116:
			goto st419
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr501
		}
		goto st0
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		switch data[p] {
		case 32:
			goto tr501
		case 46:
			goto tr582
		case 101:
			goto st420
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr501
		}
		goto st0
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
		if data[p] == 109 {
			goto st421
		}
		goto st0
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		if data[p] == 98 {
			goto st422
		}
		goto st0
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		if data[p] == 101 {
			goto st423
		}
		goto st0
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		if data[p] == 114 {
			goto st424
		}
		goto st0
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 32:
			goto tr501
		case 46:
			goto tr582
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr501
		}
		goto st0
tr269:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st425
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
//line ragel_parse_datetime.go:15175
		if data[p] == 100 {
			goto st426
		}
		goto st0
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
		if data[p] == 32 {
			goto st427
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto st429
		}
		goto st0
tr594:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st427
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
//line ragel_parse_datetime.go:15208
		switch data[p] {
		case 65:
			goto st282
		case 68:
			goto st297
		case 70:
			goto st305
		case 74:
			goto st313
		case 77:
			goto st325
		case 78:
			goto st331
		case 79:
			goto st339
		case 83:
			goto st346
		case 97:
			goto st282
		case 100:
			goto st297
		case 102:
			goto st305
		case 106:
			goto st313
		case 109:
			goto st325
		case 110:
			goto st331
		case 111:
			goto st339
		case 115:
			goto st346
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr592
		}
		goto st0
tr592:
//line ragel/datetime.rl:5
 pb = p 
	goto st428
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
//line ragel_parse_datetime.go:15256
		if data[p] == 32 {
			goto tr284
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st202
			}
		case data[p] >= 45:
			goto tr284
		}
		goto st0
tr595:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st429
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
//line ragel_parse_datetime.go:15285
		switch data[p] {
		case 65:
			goto st356
		case 68:
			goto st367
		case 70:
			goto st375
		case 74:
			goto st383
		case 77:
			goto st395
		case 78:
			goto st401
		case 79:
			goto st409
		case 83:
			goto st416
		case 97:
			goto st356
		case 100:
			goto st367
		case 102:
			goto st375
		case 106:
			goto st383
		case 109:
			goto st395
		case 110:
			goto st401
		case 111:
			goto st409
		case 115:
			goto st416
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr592
		}
		goto st0
tr270:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st430
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
//line ragel_parse_datetime.go:15340
		if data[p] == 116 {
			goto st426
		}
		goto st0
tr271:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st431
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
//line ragel_parse_datetime.go:15361
		if data[p] == 104 {
			goto st426
		}
		goto st0
tr2:
//line ragel/datetime.rl:5
 pb = p 
	goto st432
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
//line ragel_parse_datetime.go:15375
		switch data[p] {
		case 32:
			goto tr267
		case 110:
			goto tr269
		case 114:
			goto tr269
		case 115:
			goto tr270
		case 116:
			goto tr271
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr268
			}
		case data[p] > 50:
			if 51 <= data[p] && data[p] <= 57 {
				goto st433
			}
		default:
			goto st195
		}
		goto st0
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 32:
			goto tr594
		case 110:
			goto tr269
		case 114:
			goto tr269
		case 115:
			goto tr270
		case 116:
			goto tr271
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] >= 45:
			goto tr595
		}
		goto st0
tr3:
//line ragel/datetime.rl:5
 pb = p 
	goto st434
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
//line ragel_parse_datetime.go:15436
		switch data[p] {
		case 32:
			goto tr267
		case 110:
			goto tr269
		case 114:
			goto tr269
		case 115:
			goto tr270
		case 116:
			goto tr271
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st433
			}
		case data[p] >= 45:
			goto tr268
		}
		goto st0
tr4:
//line ragel/datetime.rl:5
 pb = p 
	goto st435
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
//line ragel_parse_datetime.go:15467
		switch data[p] {
		case 32:
			goto tr267
		case 110:
			goto tr269
		case 114:
			goto tr269
		case 115:
			goto tr270
		case 116:
			goto tr271
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr268
			}
		case data[p] > 49:
			if 50 <= data[p] && data[p] <= 57 {
				goto st3
			}
		default:
			goto st433
		}
		goto st0
tr5:
//line ragel/datetime.rl:5
 pb = p 
	goto st436
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
//line ragel_parse_datetime.go:15502
		switch data[p] {
		case 32:
			goto tr267
		case 110:
			goto tr269
		case 114:
			goto tr269
		case 115:
			goto tr270
		case 116:
			goto tr271
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] >= 45:
			goto tr268
		}
		goto st0
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		switch data[p] {
		case 112:
			goto st438
		case 117:
			goto st491
		}
		goto st0
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		if data[p] == 114 {
			goto st439
		}
		goto st0
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		switch data[p] {
		case 32:
			goto tr599
		case 46:
			goto tr601
		case 105:
			goto st489
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr600
		}
		goto st0
tr599:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st440
tr668:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st440
tr676:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st440
tr687:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st440
tr1286:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st440
tr1294:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st440
tr1297:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st440
tr1304:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st440
tr1308:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st440
tr1312:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st440
tr1321:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st440
tr1333:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st440
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
//line ragel_parse_datetime.go:15615
		switch data[p] {
		case 48:
			goto tr603
		case 51:
			goto tr605
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr606
			}
		case data[p] >= 49:
			goto tr604
		}
		goto st0
tr603:
//line ragel/datetime.rl:5
 pb = p 
	goto st441
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
//line ragel_parse_datetime.go:15640
		if 49 <= data[p] && data[p] <= 57 {
			goto st442
		}
		goto st0
tr606:
//line ragel/datetime.rl:5
 pb = p 
	goto st442
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
//line ragel_parse_datetime.go:15654
		switch data[p] {
		case 32:
			goto tr608
		case 44:
			goto tr609
		case 110:
			goto tr611
		case 114:
			goto tr611
		case 115:
			goto tr612
		case 116:
			goto tr613
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr610
		}
		goto st0
tr608:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
//line ragel_parse_datetime.go:15689
		switch data[p] {
		case 39:
			goto st444
		case 50:
			goto tr616
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr617
			}
		case data[p] >= 48:
			goto tr615
		}
		goto st0
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr618
		}
		goto st0
tr618:
//line ragel/datetime.rl:5
 pb = p 
	goto st445
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
//line ragel_parse_datetime.go:15723
		if 48 <= data[p] && data[p] <= 57 {
			goto st1109
		}
		goto st0
	st1109:
		if p++; p == pe {
			goto _test_eof1109
		}
	st_case_1109:
		switch data[p] {
		case 32:
			goto tr1607
		case 44:
			goto tr1608
		case 84:
			goto tr1609
		case 95:
			goto tr1609
		case 116:
			goto tr1609
		}
		goto st0
tr1710:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st446
tr1607:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st446
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
//line ragel_parse_datetime.go:15763
		switch data[p] {
		case 50:
			goto tr621
		case 65:
			goto st456
		case 97:
			goto st459
		case 109:
			goto st13
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr622
			}
		case data[p] >= 48:
			goto tr620
		}
		goto st0
tr620:
//line ragel/datetime.rl:5
 pb = p 
	goto st1110
	st1110:
		if p++; p == pe {
			goto _test_eof1110
		}
	st_case_1110:
//line ragel_parse_datetime.go:15792
		switch data[p] {
		case 32:
			goto tr1610
		case 43:
			goto tr1611
		case 45:
			goto tr1612
		case 47:
			goto tr1613
		case 58:
			goto tr1615
		case 65:
			goto tr1616
		case 80:
			goto tr1616
		case 90:
			goto tr1617
		case 95:
			goto tr1613
		case 97:
			goto tr1618
		case 112:
			goto tr1618
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1114
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1613
			}
		default:
			goto tr1613
		}
		goto st0
tr1610:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st1111
tr1625:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st1111
tr1679:
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

	goto st1111
tr1653:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st1111
tr1662:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st1111
tr1670:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st1111
tr1690:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st1111
tr1703:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st1111
	st1111:
		if p++; p == pe {
			goto _test_eof1111
		}
	st_case_1111:
//line ragel_parse_datetime.go:15912
		switch data[p] {
		case 32:
			goto st12
		case 43:
			goto st210
		case 45:
			goto st220
		case 47:
			goto tr311
		case 65:
			goto tr1619
		case 80:
			goto tr1619
		case 90:
			goto tr312
		case 95:
			goto tr311
		case 97:
			goto tr1620
		case 109:
			goto tr313
		case 112:
			goto tr1620
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr311
			}
		case data[p] >= 66:
			goto tr311
		}
		goto st0
tr1619:
//line ragel/datetime.rl:5
 pb = p 
	goto st447
tr1616:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st447
tr1630:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st447
tr1659:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st447
tr1667:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st447
tr1676:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st447
tr1681:
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
	goto st447
tr1695:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st447
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
//line ragel_parse_datetime.go:16036
		switch data[p] {
		case 47:
			goto st222
		case 77:
			goto st1112
		case 95:
			goto st222
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st222
			}
		case data[p] >= 65:
			goto st222
		}
		goto st0
	st1112:
		if p++; p == pe {
			goto _test_eof1112
		}
	st_case_1112:
		switch data[p] {
		case 32:
			goto tr1621
		case 43:
			goto tr1622
		case 45:
			goto tr1623
		case 47:
			goto tr1624
		case 95:
			goto tr1624
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1624
			}
		case data[p] >= 65:
			goto tr1624
		}
		goto st0
tr1621:
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

	goto st1113
tr1640:
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

	goto st1113
tr1633:
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

	goto st1113
tr1705:
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

	goto st1113
	st1113:
		if p++; p == pe {
			goto _test_eof1113
		}
	st_case_1113:
//line ragel_parse_datetime.go:16179
		switch data[p] {
		case 32:
			goto st12
		case 43:
			goto st210
		case 45:
			goto st220
		case 47:
			goto tr311
		case 90:
			goto tr312
		case 95:
			goto tr311
		case 109:
			goto tr313
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr311
			}
		case data[p] >= 65:
			goto tr311
		}
		goto st0
tr1620:
//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr1618:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr1632:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr1661:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr1669:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr1678:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
tr1682:
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
	goto st448
tr1697:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st448
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
//line ragel_parse_datetime.go:16295
		switch data[p] {
		case 47:
			goto st222
		case 95:
			goto st222
		case 109:
			goto st1112
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st222
			}
		case data[p] >= 65:
			goto st222
		}
		goto st0
	st1114:
		if p++; p == pe {
			goto _test_eof1114
		}
	st_case_1114:
		switch data[p] {
		case 32:
			goto tr1625
		case 43:
			goto tr1626
		case 45:
			goto tr1627
		case 47:
			goto tr1628
		case 58:
			goto tr1629
		case 65:
			goto tr1630
		case 80:
			goto tr1630
		case 90:
			goto tr1631
		case 95:
			goto tr1628
		case 97:
			goto tr1632
		case 112:
			goto tr1632
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st449
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1628
			}
		default:
			goto tr1628
		}
		goto st0
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1115
		}
		goto st0
	st1115:
		if p++; p == pe {
			goto _test_eof1115
		}
	st_case_1115:
		switch data[p] {
		case 32:
			goto tr1633
		case 43:
			goto tr1634
		case 45:
			goto tr1636
		case 47:
			goto tr1637
		case 90:
			goto tr1639
		case 95:
			goto tr1637
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1635
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1637
				}
			case data[p] >= 65:
				goto tr1637
			}
		default:
			goto st451
		}
		goto st0
tr1635:
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

	goto st450
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
//line ragel_parse_datetime.go:16423
		if 48 <= data[p] && data[p] <= 57 {
			goto tr627
		}
		goto st0
tr627:
//line ragel/datetime.rl:5
 pb = p 
	goto st1116
	st1116:
		if p++; p == pe {
			goto _test_eof1116
		}
	st_case_1116:
//line ragel_parse_datetime.go:16437
		switch data[p] {
		case 32:
			goto tr1640
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1117
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1117:
		if p++; p == pe {
			goto _test_eof1117
		}
	st_case_1117:
		switch data[p] {
		case 32:
			goto tr1640
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1118
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1118:
		if p++; p == pe {
			goto _test_eof1118
		}
	st_case_1118:
		switch data[p] {
		case 32:
			goto tr1640
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1119
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1119:
		if p++; p == pe {
			goto _test_eof1119
		}
	st_case_1119:
		switch data[p] {
		case 32:
			goto tr1640
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1120
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1120:
		if p++; p == pe {
			goto _test_eof1120
		}
	st_case_1120:
		switch data[p] {
		case 32:
			goto tr1640
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1121
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1121:
		if p++; p == pe {
			goto _test_eof1121
		}
	st_case_1121:
		switch data[p] {
		case 32:
			goto tr1640
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1122
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1122:
		if p++; p == pe {
			goto _test_eof1122
		}
	st_case_1122:
		switch data[p] {
		case 32:
			goto tr1640
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1123
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1123:
		if p++; p == pe {
			goto _test_eof1123
		}
	st_case_1123:
		switch data[p] {
		case 32:
			goto tr1640
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1124
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1124:
		if p++; p == pe {
			goto _test_eof1124
		}
	st_case_1124:
		switch data[p] {
		case 32:
			goto tr1640
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		case data[p] >= 65:
			goto tr1643
		}
		goto st0
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1125
		}
		goto st0
	st1125:
		if p++; p == pe {
			goto _test_eof1125
		}
	st_case_1125:
		switch data[p] {
		case 32:
			goto tr1633
		case 43:
			goto tr1634
		case 45:
			goto tr1636
		case 47:
			goto tr1637
		case 90:
			goto tr1639
		case 95:
			goto tr1637
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1635
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1637
			}
		default:
			goto tr1637
		}
		goto st0
tr1615:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st452
tr1629:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st452
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
//line ragel_parse_datetime.go:16775
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr630
			}
		case data[p] >= 48:
			goto tr629
		}
		goto st0
tr629:
//line ragel/datetime.rl:5
 pb = p 
	goto st1126
	st1126:
		if p++; p == pe {
			goto _test_eof1126
		}
	st_case_1126:
//line ragel_parse_datetime.go:16794
		switch data[p] {
		case 32:
			goto tr1653
		case 43:
			goto tr1654
		case 45:
			goto tr1655
		case 47:
			goto tr1656
		case 58:
			goto tr1658
		case 65:
			goto tr1659
		case 80:
			goto tr1659
		case 90:
			goto tr1660
		case 95:
			goto tr1656
		case 97:
			goto tr1661
		case 112:
			goto tr1661
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1127
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1656
			}
		default:
			goto tr1656
		}
		goto st0
	st1127:
		if p++; p == pe {
			goto _test_eof1127
		}
	st_case_1127:
		switch data[p] {
		case 32:
			goto tr1662
		case 43:
			goto tr1663
		case 45:
			goto tr1664
		case 47:
			goto tr1665
		case 58:
			goto tr1666
		case 65:
			goto tr1667
		case 80:
			goto tr1667
		case 90:
			goto tr1668
		case 95:
			goto tr1665
		case 97:
			goto tr1669
		case 112:
			goto tr1669
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1665
			}
		case data[p] >= 66:
			goto tr1665
		}
		goto st0
tr1658:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st453
tr1666:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st453
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
//line ragel_parse_datetime.go:16887
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr632
			}
		case data[p] >= 48:
			goto tr631
		}
		goto st0
tr631:
//line ragel/datetime.rl:5
 pb = p 
	goto st1128
	st1128:
		if p++; p == pe {
			goto _test_eof1128
		}
	st_case_1128:
//line ragel_parse_datetime.go:16906
		switch data[p] {
		case 32:
			goto tr1670
		case 43:
			goto tr1671
		case 45:
			goto tr1673
		case 47:
			goto tr1674
		case 65:
			goto tr1676
		case 80:
			goto tr1676
		case 90:
			goto tr1677
		case 95:
			goto tr1674
		case 97:
			goto tr1678
		case 112:
			goto tr1678
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1672
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1674
				}
			case data[p] >= 66:
				goto tr1674
			}
		default:
			goto st1138
		}
		goto st0
tr1672:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st454
tr1692:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st454
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
//line ragel_parse_datetime.go:16964
		if 48 <= data[p] && data[p] <= 57 {
			goto tr633
		}
		goto st0
tr633:
//line ragel/datetime.rl:5
 pb = p 
	goto st1129
	st1129:
		if p++; p == pe {
			goto _test_eof1129
		}
	st_case_1129:
//line ragel_parse_datetime.go:16978
		switch data[p] {
		case 32:
			goto tr1679
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 65:
			goto tr1681
		case 80:
			goto tr1681
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		case 97:
			goto tr1682
		case 112:
			goto tr1682
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1130
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1130:
		if p++; p == pe {
			goto _test_eof1130
		}
	st_case_1130:
		switch data[p] {
		case 32:
			goto tr1679
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 65:
			goto tr1681
		case 80:
			goto tr1681
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		case 97:
			goto tr1682
		case 112:
			goto tr1682
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1131
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1131:
		if p++; p == pe {
			goto _test_eof1131
		}
	st_case_1131:
		switch data[p] {
		case 32:
			goto tr1679
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 65:
			goto tr1681
		case 80:
			goto tr1681
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		case 97:
			goto tr1682
		case 112:
			goto tr1682
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1132
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1132:
		if p++; p == pe {
			goto _test_eof1132
		}
	st_case_1132:
		switch data[p] {
		case 32:
			goto tr1679
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 65:
			goto tr1681
		case 80:
			goto tr1681
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		case 97:
			goto tr1682
		case 112:
			goto tr1682
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1133
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1133:
		if p++; p == pe {
			goto _test_eof1133
		}
	st_case_1133:
		switch data[p] {
		case 32:
			goto tr1679
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 65:
			goto tr1681
		case 80:
			goto tr1681
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		case 97:
			goto tr1682
		case 112:
			goto tr1682
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1134
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1134:
		if p++; p == pe {
			goto _test_eof1134
		}
	st_case_1134:
		switch data[p] {
		case 32:
			goto tr1679
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 65:
			goto tr1681
		case 80:
			goto tr1681
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		case 97:
			goto tr1682
		case 112:
			goto tr1682
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1135
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1135:
		if p++; p == pe {
			goto _test_eof1135
		}
	st_case_1135:
		switch data[p] {
		case 32:
			goto tr1679
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 65:
			goto tr1681
		case 80:
			goto tr1681
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		case 97:
			goto tr1682
		case 112:
			goto tr1682
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1136
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1136:
		if p++; p == pe {
			goto _test_eof1136
		}
	st_case_1136:
		switch data[p] {
		case 32:
			goto tr1679
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 65:
			goto tr1681
		case 80:
			goto tr1681
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		case 97:
			goto tr1682
		case 112:
			goto tr1682
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1137
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		default:
			goto tr1643
		}
		goto st0
	st1137:
		if p++; p == pe {
			goto _test_eof1137
		}
	st_case_1137:
		switch data[p] {
		case 32:
			goto tr1679
		case 43:
			goto tr1641
		case 45:
			goto tr1642
		case 47:
			goto tr1643
		case 65:
			goto tr1681
		case 80:
			goto tr1681
		case 90:
			goto tr1645
		case 95:
			goto tr1643
		case 97:
			goto tr1682
		case 112:
			goto tr1682
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1643
			}
		case data[p] >= 66:
			goto tr1643
		}
		goto st0
	st1138:
		if p++; p == pe {
			goto _test_eof1138
		}
	st_case_1138:
		switch data[p] {
		case 32:
			goto tr1690
		case 43:
			goto tr1691
		case 45:
			goto tr1693
		case 47:
			goto tr1694
		case 65:
			goto tr1695
		case 80:
			goto tr1695
		case 90:
			goto tr1696
		case 95:
			goto tr1694
		case 97:
			goto tr1697
		case 112:
			goto tr1697
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1692
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1694
			}
		default:
			goto tr1694
		}
		goto st0
tr632:
//line ragel/datetime.rl:5
 pb = p 
	goto st1139
	st1139:
		if p++; p == pe {
			goto _test_eof1139
		}
	st_case_1139:
//line ragel_parse_datetime.go:17379
		switch data[p] {
		case 32:
			goto tr1670
		case 43:
			goto tr1671
		case 45:
			goto tr1673
		case 47:
			goto tr1674
		case 65:
			goto tr1676
		case 80:
			goto tr1676
		case 90:
			goto tr1677
		case 95:
			goto tr1674
		case 97:
			goto tr1678
		case 112:
			goto tr1678
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1672
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1674
			}
		default:
			goto tr1674
		}
		goto st0
tr630:
//line ragel/datetime.rl:5
 pb = p 
	goto st1140
	st1140:
		if p++; p == pe {
			goto _test_eof1140
		}
	st_case_1140:
//line ragel_parse_datetime.go:17424
		switch data[p] {
		case 32:
			goto tr1653
		case 43:
			goto tr1654
		case 45:
			goto tr1655
		case 47:
			goto tr1656
		case 58:
			goto tr1658
		case 65:
			goto tr1659
		case 80:
			goto tr1659
		case 90:
			goto tr1660
		case 95:
			goto tr1656
		case 97:
			goto tr1661
		case 112:
			goto tr1661
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1656
			}
		case data[p] >= 66:
			goto tr1656
		}
		goto st0
tr621:
//line ragel/datetime.rl:5
 pb = p 
	goto st1141
	st1141:
		if p++; p == pe {
			goto _test_eof1141
		}
	st_case_1141:
//line ragel_parse_datetime.go:17467
		switch data[p] {
		case 32:
			goto tr1610
		case 43:
			goto tr1611
		case 45:
			goto tr1612
		case 47:
			goto tr1613
		case 58:
			goto tr1615
		case 65:
			goto tr1616
		case 80:
			goto tr1616
		case 90:
			goto tr1617
		case 95:
			goto tr1613
		case 97:
			goto tr1618
		case 112:
			goto tr1618
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st1114
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1613
				}
			case data[p] >= 66:
				goto tr1613
			}
		default:
			goto st455
		}
		goto st0
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
		if 48 <= data[p] && data[p] <= 57 {
			goto st449
		}
		goto st0
tr622:
//line ragel/datetime.rl:5
 pb = p 
	goto st1142
	st1142:
		if p++; p == pe {
			goto _test_eof1142
		}
	st_case_1142:
//line ragel_parse_datetime.go:17528
		switch data[p] {
		case 32:
			goto tr1610
		case 43:
			goto tr1611
		case 45:
			goto tr1612
		case 47:
			goto tr1613
		case 58:
			goto tr1615
		case 65:
			goto tr1616
		case 80:
			goto tr1616
		case 90:
			goto tr1617
		case 95:
			goto tr1613
		case 97:
			goto tr1618
		case 112:
			goto tr1618
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st455
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1613
			}
		default:
			goto tr1613
		}
		goto st0
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		if data[p] == 84 {
			goto st457
		}
		goto st0
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		if data[p] == 32 {
			goto st458
		}
		goto st0
tr1712:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st458
tr1609:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st458
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
//line ragel_parse_datetime.go:17601
		if data[p] == 50 {
			goto tr621
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr622
			}
		case data[p] >= 48:
			goto tr620
		}
		goto st0
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
		if data[p] == 116 {
			goto st457
		}
		goto st0
tr1711:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st1143
tr1608:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st1143
	st1143:
		if p++; p == pe {
			goto _test_eof1143
		}
	st_case_1143:
//line ragel_parse_datetime.go:17640
		switch data[p] {
		case 32:
			goto st446
		case 44:
			goto st457
		case 84:
			goto st458
		case 95:
			goto st458
		case 116:
			goto st458
		}
		goto st0
tr615:
//line ragel/datetime.rl:5
 pb = p 
	goto st460
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
//line ragel_parse_datetime.go:17663
		switch data[p] {
		case 32:
			goto tr298
		case 58:
			goto tr300
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1144
		}
		goto st0
	st1144:
		if p++; p == pe {
			goto _test_eof1144
		}
	st_case_1144:
		switch data[p] {
		case 32:
			goto tr1700
		case 44:
			goto tr1608
		case 58:
			goto tr343
		case 65:
			goto tr344
		case 80:
			goto tr344
		case 84:
			goto tr1609
		case 95:
			goto tr1609
		case 97:
			goto tr345
		case 112:
			goto tr345
		case 116:
			goto tr1609
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st465
		}
		goto st0
tr1700:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st461
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
//line ragel_parse_datetime.go:17728
		switch data[p] {
		case 39:
			goto st207
		case 50:
			goto tr639
		case 65:
			goto tr641
		case 80:
			goto tr305
		case 97:
			goto tr642
		case 109:
			goto st13
		case 112:
			goto tr306
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr640
			}
		case data[p] >= 48:
			goto tr638
		}
		goto st0
tr638:
//line ragel/datetime.rl:5
 pb = p 
	goto st1145
	st1145:
		if p++; p == pe {
			goto _test_eof1145
		}
	st_case_1145:
//line ragel_parse_datetime.go:17763
		switch data[p] {
		case 32:
			goto tr1610
		case 43:
			goto tr1611
		case 45:
			goto tr1612
		case 47:
			goto tr1613
		case 58:
			goto tr1615
		case 65:
			goto tr1616
		case 80:
			goto tr1616
		case 90:
			goto tr1617
		case 95:
			goto tr1613
		case 97:
			goto tr1618
		case 112:
			goto tr1618
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1146
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1613
			}
		default:
			goto tr1613
		}
		goto st0
	st1146:
		if p++; p == pe {
			goto _test_eof1146
		}
	st_case_1146:
		switch data[p] {
		case 32:
			goto tr1703
		case 43:
			goto tr1626
		case 45:
			goto tr1627
		case 47:
			goto tr1628
		case 58:
			goto tr1629
		case 65:
			goto tr1630
		case 80:
			goto tr1630
		case 90:
			goto tr1631
		case 95:
			goto tr1628
		case 97:
			goto tr1632
		case 112:
			goto tr1632
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st462
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1628
			}
		default:
			goto tr1628
		}
		goto st0
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1147
		}
		goto st0
	st1147:
		if p++; p == pe {
			goto _test_eof1147
		}
	st_case_1147:
		switch data[p] {
		case 32:
			goto tr1705
		case 43:
			goto tr1634
		case 45:
			goto tr1636
		case 47:
			goto tr1637
		case 90:
			goto tr1639
		case 95:
			goto tr1637
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr1635
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1637
				}
			case data[p] >= 65:
				goto tr1637
			}
		default:
			goto st451
		}
		goto st0
tr639:
//line ragel/datetime.rl:5
 pb = p 
	goto st1148
	st1148:
		if p++; p == pe {
			goto _test_eof1148
		}
	st_case_1148:
//line ragel_parse_datetime.go:17898
		switch data[p] {
		case 32:
			goto tr1610
		case 43:
			goto tr1611
		case 45:
			goto tr1612
		case 47:
			goto tr1613
		case 58:
			goto tr1615
		case 65:
			goto tr1616
		case 80:
			goto tr1616
		case 90:
			goto tr1617
		case 95:
			goto tr1613
		case 97:
			goto tr1618
		case 112:
			goto tr1618
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st1146
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1613
				}
			case data[p] >= 66:
				goto tr1613
			}
		default:
			goto st1149
		}
		goto st0
	st1149:
		if p++; p == pe {
			goto _test_eof1149
		}
	st_case_1149:
		if data[p] == 32 {
			goto tr1564
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st462
		}
		goto st0
tr640:
//line ragel/datetime.rl:5
 pb = p 
	goto st1150
	st1150:
		if p++; p == pe {
			goto _test_eof1150
		}
	st_case_1150:
//line ragel_parse_datetime.go:17962
		switch data[p] {
		case 32:
			goto tr1610
		case 43:
			goto tr1611
		case 45:
			goto tr1612
		case 47:
			goto tr1613
		case 58:
			goto tr1615
		case 65:
			goto tr1616
		case 80:
			goto tr1616
		case 90:
			goto tr1617
		case 95:
			goto tr1613
		case 97:
			goto tr1618
		case 112:
			goto tr1618
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1149
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr1613
			}
		default:
			goto tr1613
		}
		goto st0
tr641:
//line ragel/datetime.rl:5
 pb = p 
	goto st463
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
//line ragel_parse_datetime.go:18009
		switch data[p] {
		case 77:
			goto st227
		case 84:
			goto st457
		}
		goto st0
tr642:
//line ragel/datetime.rl:5
 pb = p 
	goto st464
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
//line ragel_parse_datetime.go:18026
		switch data[p] {
		case 109:
			goto st227
		case 116:
			goto st457
		}
		goto st0
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1151
		}
		goto st0
	st1151:
		if p++; p == pe {
			goto _test_eof1151
		}
	st_case_1151:
		switch data[p] {
		case 32:
			goto tr1589
		case 43:
			goto tr1557
		case 44:
			goto tr1707
		case 45:
			goto tr1559
		case 46:
			goto tr360
		case 47:
			goto tr1560
		case 84:
			goto tr1561
		case 90:
			goto tr1562
		case 95:
			goto tr1563
		case 116:
			goto tr1563
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st244
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1560
			}
		default:
			goto tr1560
		}
		goto st0
tr1707:
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

	goto st1152
	st1152:
		if p++; p == pe {
			goto _test_eof1152
		}
	st_case_1152:
//line ragel_parse_datetime.go:18109
		switch data[p] {
		case 32:
			goto st466
		case 44:
			goto st457
		case 84:
			goto st458
		case 95:
			goto st458
		case 116:
			goto st458
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr348
		}
		goto st0
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		switch data[p] {
		case 50:
			goto tr93
		case 65:
			goto st456
		case 97:
			goto st459
		case 109:
			goto st13
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr94
			}
		case data[p] >= 48:
			goto tr92
		}
		goto st0
tr616:
//line ragel/datetime.rl:5
 pb = p 
	goto st467
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
//line ragel_parse_datetime.go:18159
		switch data[p] {
		case 32:
			goto tr298
		case 58:
			goto tr300
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st1153
			}
		case data[p] >= 48:
			goto st1144
		}
		goto st0
	st1153:
		if p++; p == pe {
			goto _test_eof1153
		}
	st_case_1153:
		switch data[p] {
		case 32:
			goto tr1607
		case 44:
			goto tr1608
		case 84:
			goto tr1609
		case 95:
			goto tr1609
		case 116:
			goto tr1609
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st465
		}
		goto st0
tr617:
//line ragel/datetime.rl:5
 pb = p 
	goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
//line ragel_parse_datetime.go:18213
		switch data[p] {
		case 32:
			goto tr298
		case 58:
			goto tr300
		case 65:
			goto tr301
		case 80:
			goto tr301
		case 97:
			goto tr302
		case 112:
			goto tr302
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1153
		}
		goto st0
tr609:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st469
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
//line ragel_parse_datetime.go:18248
		if data[p] == 32 {
			goto st470
		}
		goto st0
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		if data[p] == 39 {
			goto st444
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr647
		}
		goto st0
tr647:
//line ragel/datetime.rl:5
 pb = p 
	goto st471
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
//line ragel_parse_datetime.go:18274
		if 48 <= data[p] && data[p] <= 57 {
			goto st1154
		}
		goto st0
	st1154:
		if p++; p == pe {
			goto _test_eof1154
		}
	st_case_1154:
		switch data[p] {
		case 32:
			goto tr1607
		case 44:
			goto tr1608
		case 84:
			goto tr1609
		case 95:
			goto tr1609
		case 116:
			goto tr1609
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st472
		}
		goto st0
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1155
		}
		goto st0
	st1155:
		if p++; p == pe {
			goto _test_eof1155
		}
	st_case_1155:
		switch data[p] {
		case 32:
			goto tr1710
		case 44:
			goto tr1711
		case 84:
			goto tr1712
		case 95:
			goto tr1712
		case 116:
			goto tr1712
		}
		goto st0
tr611:
//line ragel/datetime.rl:107

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
//line ragel_parse_datetime.go:18343
		if data[p] == 100 {
			goto st474
		}
		goto st0
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		switch data[p] {
		case 32:
			goto st443
		case 44:
			goto st469
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto st198
		}
		goto st0
tr612:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st475
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
//line ragel_parse_datetime.go:18379
		if data[p] == 116 {
			goto st474
		}
		goto st0
tr613:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st476
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
//line ragel_parse_datetime.go:18400
		if data[p] == 104 {
			goto st474
		}
		goto st0
tr604:
//line ragel/datetime.rl:5
 pb = p 
	goto st477
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
//line ragel_parse_datetime.go:18414
		switch data[p] {
		case 32:
			goto tr608
		case 44:
			goto tr609
		case 110:
			goto tr611
		case 114:
			goto tr611
		case 115:
			goto tr612
		case 116:
			goto tr613
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st442
			}
		case data[p] >= 45:
			goto tr610
		}
		goto st0
tr605:
//line ragel/datetime.rl:5
 pb = p 
	goto st478
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
//line ragel_parse_datetime.go:18447
		switch data[p] {
		case 32:
			goto tr608
		case 44:
			goto tr609
		case 110:
			goto tr611
		case 114:
			goto tr611
		case 115:
			goto tr612
		case 116:
			goto tr613
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st442
			}
		case data[p] >= 45:
			goto tr610
		}
		goto st0
tr600:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st479
tr669:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st479
tr677:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st479
tr688:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st479
tr1030:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st479
tr1039:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st479
tr1043:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st479
tr1050:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st479
tr1055:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st479
tr1060:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st479
tr1070:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st479
tr1079:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st479
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
//line ragel_parse_datetime.go:18524
		switch data[p] {
		case 48:
			goto tr653
		case 51:
			goto tr655
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr656
			}
		case data[p] >= 49:
			goto tr654
		}
		goto st0
tr653:
//line ragel/datetime.rl:5
 pb = p 
	goto st480
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
//line ragel_parse_datetime.go:18549
		if 49 <= data[p] && data[p] <= 57 {
			goto st481
		}
		goto st0
tr656:
//line ragel/datetime.rl:5
 pb = p 
	goto st481
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
//line ragel_parse_datetime.go:18563
		switch data[p] {
		case 32:
			goto tr658
		case 110:
			goto tr659
		case 114:
			goto tr659
		case 115:
			goto tr660
		case 116:
			goto tr661
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr610
		}
		goto st0
tr659:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st482
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
//line ragel_parse_datetime.go:18596
		if data[p] == 100 {
			goto st483
		}
		goto st0
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
		if data[p] == 32 {
			goto st204
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto st198
		}
		goto st0
tr660:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st484
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
//line ragel_parse_datetime.go:18629
		if data[p] == 116 {
			goto st483
		}
		goto st0
tr661:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st485
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
//line ragel_parse_datetime.go:18650
		if data[p] == 104 {
			goto st483
		}
		goto st0
tr654:
//line ragel/datetime.rl:5
 pb = p 
	goto st486
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
//line ragel_parse_datetime.go:18664
		switch data[p] {
		case 32:
			goto tr658
		case 110:
			goto tr659
		case 114:
			goto tr659
		case 115:
			goto tr660
		case 116:
			goto tr661
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st481
			}
		case data[p] >= 45:
			goto tr610
		}
		goto st0
tr655:
//line ragel/datetime.rl:5
 pb = p 
	goto st487
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
//line ragel_parse_datetime.go:18695
		switch data[p] {
		case 32:
			goto tr658
		case 110:
			goto tr659
		case 114:
			goto tr659
		case 115:
			goto tr660
		case 116:
			goto tr661
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st481
			}
		case data[p] >= 45:
			goto tr610
		}
		goto st0
tr601:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st488
tr670:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st488
tr678:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st488
tr689:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st488
tr1287:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st488
tr1295:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st488
tr1298:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st488
tr1305:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st488
tr1309:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st488
tr1313:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st488
tr1322:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st488
tr1334:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st488
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
//line ragel_parse_datetime.go:18770
		switch data[p] {
		case 32:
			goto st440
		case 48:
			goto tr653
		case 51:
			goto tr655
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st479
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr656
			}
		default:
			goto tr654
		}
		goto st0
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
		if data[p] == 108 {
			goto st490
		}
		goto st0
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		switch data[p] {
		case 32:
			goto tr599
		case 46:
			goto tr601
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr600
		}
		goto st0
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		if data[p] == 103 {
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
			goto tr668
		case 46:
			goto tr670
		case 117:
			goto st493
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr669
		}
		goto st0
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
		if data[p] == 115 {
			goto st494
		}
		goto st0
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		if data[p] == 116 {
			goto st495
		}
		goto st0
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		switch data[p] {
		case 32:
			goto tr668
		case 46:
			goto tr670
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr669
		}
		goto st0
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
		if data[p] == 101 {
			goto st497
		}
		goto st0
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		if data[p] == 99 {
			goto st498
		}
		goto st0
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		switch data[p] {
		case 32:
			goto tr676
		case 46:
			goto tr678
		case 101:
			goto st499
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr677
		}
		goto st0
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		if data[p] == 109 {
			goto st500
		}
		goto st0
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		if data[p] == 98 {
			goto st501
		}
		goto st0
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		if data[p] == 101 {
			goto st502
		}
		goto st0
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		if data[p] == 114 {
			goto st503
		}
		goto st0
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		switch data[p] {
		case 32:
			goto tr676
		case 46:
			goto tr678
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr677
		}
		goto st0
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		switch data[p] {
		case 101:
			goto st505
		case 114:
			goto st512
		}
		goto st0
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		if data[p] == 98 {
			goto st506
		}
		goto st0
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		switch data[p] {
		case 32:
			goto tr687
		case 46:
			goto tr689
		case 114:
			goto st507
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr688
		}
		goto st0
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		if data[p] == 117 {
			goto st508
		}
		goto st0
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		if data[p] == 97 {
			goto st509
		}
		goto st0
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		if data[p] == 114 {
			goto st510
		}
		goto st0
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
		if data[p] == 121 {
			goto st511
		}
		goto st0
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		switch data[p] {
		case 32:
			goto tr687
		case 46:
			goto tr689
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr688
		}
		goto st0
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		if data[p] == 105 {
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
			goto st514
		case 44:
			goto st755
		case 100:
			goto st913
		}
		goto st0
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		switch data[p] {
		case 48:
			goto tr699
		case 49:
			goto tr700
		case 50:
			goto tr701
		case 51:
			goto tr702
		case 65:
			goto st525
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
			goto st525
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
		if 52 <= data[p] && data[p] <= 57 {
			goto tr703
		}
		goto st0
tr699:
//line ragel/datetime.rl:5
 pb = p 
	goto st515
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
//line ragel_parse_datetime.go:19133
		if 49 <= data[p] && data[p] <= 57 {
			goto st516
		}
		goto st0
tr703:
//line ragel/datetime.rl:5
 pb = p 
	goto st516
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
//line ragel_parse_datetime.go:19147
		switch data[p] {
		case 32:
			goto tr268
		case 110:
			goto tr713
		case 114:
			goto tr713
		case 115:
			goto tr714
		case 116:
			goto tr715
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr268
		}
		goto st0
tr713:
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
//line ragel_parse_datetime.go:19180
		if data[p] == 100 {
			goto st518
		}
		goto st0
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		if data[p] == 32 {
			goto st429
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto st429
		}
		goto st0
tr714:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st519
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
//line ragel_parse_datetime.go:19213
		if data[p] == 116 {
			goto st518
		}
		goto st0
tr715:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st520
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
//line ragel_parse_datetime.go:19234
		if data[p] == 104 {
			goto st518
		}
		goto st0
tr700:
//line ragel/datetime.rl:5
 pb = p 
	goto st521
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
//line ragel_parse_datetime.go:19248
		switch data[p] {
		case 32:
			goto tr268
		case 110:
			goto tr713
		case 114:
			goto tr713
		case 115:
			goto tr714
		case 116:
			goto tr715
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 47 {
				goto tr268
			}
		case data[p] > 50:
			if 51 <= data[p] && data[p] <= 57 {
				goto st522
			}
		default:
			goto st516
		}
		goto st0
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		switch data[p] {
		case 32:
			goto tr595
		case 110:
			goto tr713
		case 114:
			goto tr713
		case 115:
			goto tr714
		case 116:
			goto tr715
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr595
		}
		goto st0
tr701:
//line ragel/datetime.rl:5
 pb = p 
	goto st523
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
//line ragel_parse_datetime.go:19304
		switch data[p] {
		case 32:
			goto tr268
		case 110:
			goto tr713
		case 114:
			goto tr713
		case 115:
			goto tr714
		case 116:
			goto tr715
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st522
			}
		case data[p] >= 45:
			goto tr268
		}
		goto st0
tr702:
//line ragel/datetime.rl:5
 pb = p 
	goto st524
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
//line ragel_parse_datetime.go:19335
		switch data[p] {
		case 32:
			goto tr268
		case 110:
			goto tr713
		case 114:
			goto tr713
		case 115:
			goto tr714
		case 116:
			goto tr715
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st522
			}
		case data[p] >= 45:
			goto tr268
		}
		goto st0
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		switch data[p] {
		case 112:
			goto st526
		case 117:
			goto st692
		}
		goto st0
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
		if data[p] == 114 {
			goto st527
		}
		goto st0
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
		switch data[p] {
		case 32:
			goto tr721
		case 46:
			goto tr722
		case 105:
			goto st690
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr600
		}
		goto st0
tr721:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st528
tr1003:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st528
tr1010:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st528
tr1019:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st528
tr1029:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st528
tr1038:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st528
tr1042:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st528
tr1049:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st528
tr1054:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st528
tr1059:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st528
tr1069:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st528
tr1078:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st528
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
//line ragel_parse_datetime.go:19448
		switch data[p] {
		case 32:
			goto st529
		case 48:
			goto tr725
		case 51:
			goto tr727
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr728
			}
		case data[p] >= 49:
			goto tr726
		}
		goto st0
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
		switch data[p] {
		case 48:
			goto tr729
		case 51:
			goto tr731
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr732
			}
		case data[p] >= 49:
			goto tr730
		}
		goto st0
tr729:
//line ragel/datetime.rl:5
 pb = p 
	goto st530
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
//line ragel_parse_datetime.go:19495
		if 49 <= data[p] && data[p] <= 57 {
			goto st531
		}
		goto st0
tr732:
//line ragel/datetime.rl:5
 pb = p 
	goto st531
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
//line ragel_parse_datetime.go:19509
		switch data[p] {
		case 32:
			goto tr734
		case 110:
			goto tr735
		case 114:
			goto tr735
		case 115:
			goto tr736
		case 116:
			goto tr737
		}
		goto st0
tr734:
//line ragel/datetime.rl:107

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
//line ragel_parse_datetime.go:19539
		if data[p] == 50 {
			goto tr739
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr740
			}
		case data[p] >= 48:
			goto tr738
		}
		goto st0
tr738:
//line ragel/datetime.rl:5
 pb = p 
	goto st533
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
//line ragel_parse_datetime.go:19561
		switch data[p] {
		case 32:
			goto tr741
		case 43:
			goto tr742
		case 45:
			goto tr743
		case 47:
			goto tr744
		case 58:
			goto tr746
		case 65:
			goto tr747
		case 80:
			goto tr747
		case 90:
			goto tr748
		case 95:
			goto tr744
		case 97:
			goto tr749
		case 112:
			goto tr749
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st574
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr744
			}
		default:
			goto tr744
		}
		goto st0
tr741:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st534
tr812:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st534
tr875:
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

	goto st534
tr846:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st534
tr855:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st534
tr865:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st534
tr886:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st534
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
//line ragel_parse_datetime.go:19671
		switch data[p] {
		case 32:
			goto st535
		case 39:
			goto st536
		case 43:
			goto st540
		case 45:
			goto st565
		case 47:
			goto tr754
		case 65:
			goto tr756
		case 80:
			goto tr756
		case 90:
			goto tr757
		case 95:
			goto tr754
		case 97:
			goto tr758
		case 112:
			goto tr758
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr755
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr754
			}
		default:
			goto tr754
		}
		goto st0
tr775:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st535
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
//line ragel_parse_datetime.go:19718
		if data[p] == 39 {
			goto st536
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr755
		}
		goto st0
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr759
		}
		goto st0
tr759:
//line ragel/datetime.rl:5
 pb = p 
	goto st537
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
//line ragel_parse_datetime.go:19744
		if 48 <= data[p] && data[p] <= 57 {
			goto st1156
		}
		goto st0
	st1156:
		if p++; p == pe {
			goto _test_eof1156
		}
	st_case_1156:
		if data[p] == 32 {
			goto tr1713
		}
		goto st0
tr755:
//line ragel/datetime.rl:5
 pb = p 
	goto st538
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
//line ragel_parse_datetime.go:19767
		if 48 <= data[p] && data[p] <= 57 {
			goto st1157
		}
		goto st0
	st1157:
		if p++; p == pe {
			goto _test_eof1157
		}
	st_case_1157:
		if data[p] == 32 {
			goto tr1713
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st539
		}
		goto st0
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1158
		}
		goto st0
	st1158:
		if p++; p == pe {
			goto _test_eof1158
		}
	st_case_1158:
		if data[p] == 32 {
			goto tr1715
		}
		goto st0
tr742:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st540
tr809:
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

	goto st540
tr813:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st540
tr831:
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

	goto st540
tr823:
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
tr847:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st540
tr856:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st540
tr866:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st540
tr887:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st540
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
//line ragel_parse_datetime.go:19916
		if data[p] == 50 {
			goto tr764
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr765
			}
		case data[p] >= 48:
			goto tr763
		}
		goto st0
tr763:
//line ragel/datetime.rl:5
 pb = p 
	goto st541
tr800:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st541
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
//line ragel_parse_datetime.go:19944
		switch data[p] {
		case 32:
			goto tr766
		case 58:
			goto tr768
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st559
		}
		goto st0
tr766:
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
	goto st542
tr795:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st542
tr798:
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
	goto st542
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
//line ragel_parse_datetime.go:20012
		switch data[p] {
		case 39:
			goto st536
		case 40:
			goto st543
		case 43:
			goto st546
		case 45:
			goto st555
		case 47:
			goto tr772
		case 95:
			goto tr772
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr755
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr772
			}
		default:
			goto tr772
		}
		goto st0
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
		switch data[p] {
		case 32:
			goto st544
		case 43:
			goto st544
		case 45:
			goto st544
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st544
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st544
			}
		default:
			goto st544
		}
		goto st0
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
		switch data[p] {
		case 32:
			goto st544
		case 41:
			goto st545
		case 43:
			goto st544
		case 45:
			goto st544
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st544
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st544
			}
		default:
			goto st544
		}
		goto st0
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
		if data[p] == 32 {
			goto tr775
		}
		goto st0
tr805:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st546
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
//line ragel_parse_datetime.go:20115
		if data[p] == 50 {
			goto tr777
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr778
			}
		case data[p] >= 48:
			goto tr776
		}
		goto st0
tr776:
//line ragel/datetime.rl:5
 pb = p 
	goto st547
tr788:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st547
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
//line ragel_parse_datetime.go:20143
		switch data[p] {
		case 32:
			goto tr779
		case 58:
			goto tr781
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st549
		}
		goto st0
tr779:
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
	goto st548
tr783:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st548
tr786:
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
	goto st548
tr793:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st548
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
//line ragel_parse_datetime.go:20220
		switch data[p] {
		case 39:
			goto st536
		case 40:
			goto st543
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr755
		}
		goto st0
tr778:
//line ragel/datetime.rl:5
 pb = p 
	goto st549
tr790:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st549
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
//line ragel_parse_datetime.go:20246
		switch data[p] {
		case 32:
			goto tr779
		case 58:
			goto tr781
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st550
		}
		goto st0
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
		if data[p] == 32 {
			goto tr779
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st550
		}
		goto st0
tr781:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st551
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
//line ragel_parse_datetime.go:20286
		if data[p] == 32 {
			goto tr783
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr785
			}
		case data[p] >= 48:
			goto tr784
		}
		goto st0
tr784:
//line ragel/datetime.rl:5
 pb = p 
	goto st552
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
//line ragel_parse_datetime.go:20308
		if data[p] == 32 {
			goto tr786
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st553
		}
		goto st0
tr785:
//line ragel/datetime.rl:5
 pb = p 
	goto st553
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
//line ragel_parse_datetime.go:20325
		if data[p] == 32 {
			goto tr786
		}
		goto st0
tr777:
//line ragel/datetime.rl:5
 pb = p 
	goto st554
tr789:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st554
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
//line ragel_parse_datetime.go:20345
		switch data[p] {
		case 32:
			goto tr779
		case 58:
			goto tr781
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st550
			}
		case data[p] >= 48:
			goto st549
		}
		goto st0
tr806:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st555
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
//line ragel_parse_datetime.go:20373
		if data[p] == 50 {
			goto tr789
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr790
			}
		case data[p] >= 48:
			goto tr788
		}
		goto st0
tr772:
//line ragel/datetime.rl:5
 pb = p 
	goto st556
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
//line ragel_parse_datetime.go:20395
		switch data[p] {
		case 47:
			goto st557
		case 95:
			goto st557
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st557
			}
		case data[p] >= 65:
			goto st557
		}
		goto st0
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
		switch data[p] {
		case 47:
			goto st558
		case 95:
			goto st558
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st558
			}
		case data[p] >= 65:
			goto st558
		}
		goto st0
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
		switch data[p] {
		case 32:
			goto tr793
		case 47:
			goto st558
		case 95:
			goto st558
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st558
			}
		case data[p] >= 65:
			goto st558
		}
		goto st0
tr765:
//line ragel/datetime.rl:5
 pb = p 
	goto st559
tr802:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st559
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
//line ragel_parse_datetime.go:20468
		switch data[p] {
		case 32:
			goto tr766
		case 58:
			goto tr768
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st560
		}
		goto st0
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
		if data[p] == 32 {
			goto tr766
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st560
		}
		goto st0
tr768:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st561
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
//line ragel_parse_datetime.go:20508
		if data[p] == 32 {
			goto tr795
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr797
			}
		case data[p] >= 48:
			goto tr796
		}
		goto st0
tr796:
//line ragel/datetime.rl:5
 pb = p 
	goto st562
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
//line ragel_parse_datetime.go:20530
		if data[p] == 32 {
			goto tr798
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st563
		}
		goto st0
tr797:
//line ragel/datetime.rl:5
 pb = p 
	goto st563
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
//line ragel_parse_datetime.go:20547
		if data[p] == 32 {
			goto tr798
		}
		goto st0
tr764:
//line ragel/datetime.rl:5
 pb = p 
	goto st564
tr801:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st564
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
//line ragel_parse_datetime.go:20567
		switch data[p] {
		case 32:
			goto tr766
		case 58:
			goto tr768
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st560
			}
		case data[p] >= 48:
			goto st559
		}
		goto st0
tr743:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st565
tr810:
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

	goto st565
tr814:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st565
tr832:
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

	goto st565
tr825:
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

	goto st565
tr848:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st565
tr857:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st565
tr868:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st565
tr889:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st565
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
//line ragel_parse_datetime.go:20697
		if data[p] == 50 {
			goto tr801
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr802
			}
		case data[p] >= 48:
			goto tr800
		}
		goto st0
tr754:
//line ragel/datetime.rl:5
 pb = p 
	goto st566
tr744:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st566
tr815:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st566
tr849:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st566
tr858:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st566
tr869:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st566
tr833:
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
	goto st566
tr890:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st566
tr826:
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
	goto st566
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
//line ragel_parse_datetime.go:20819
		switch data[p] {
		case 47:
			goto st567
		case 95:
			goto st567
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st567
			}
		case data[p] >= 65:
			goto st567
		}
		goto st0
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		switch data[p] {
		case 47:
			goto st568
		case 95:
			goto st568
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st568
			}
		case data[p] >= 65:
			goto st568
		}
		goto st0
tr811:
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
	goto st568
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
//line ragel_parse_datetime.go:20887
		switch data[p] {
		case 32:
			goto tr793
		case 43:
			goto tr805
		case 45:
			goto tr806
		case 47:
			goto st568
		case 95:
			goto st568
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st568
			}
		case data[p] >= 65:
			goto st568
		}
		goto st0
tr756:
//line ragel/datetime.rl:5
 pb = p 
	goto st569
tr747:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st569
tr818:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st569
tr852:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st569
tr860:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st569
tr871:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st569
tr877:
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
	goto st569
tr891:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st569
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
//line ragel_parse_datetime.go:20999
		switch data[p] {
		case 47:
			goto st567
		case 77:
			goto st570
		case 95:
			goto st567
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st567
			}
		case data[p] >= 65:
			goto st567
		}
		goto st0
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		switch data[p] {
		case 32:
			goto tr808
		case 43:
			goto tr809
		case 45:
			goto tr810
		case 47:
			goto tr811
		case 95:
			goto tr811
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr811
			}
		case data[p] >= 65:
			goto tr811
		}
		goto st0
tr808:
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

	goto st571
tr830:
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

	goto st571
tr822:
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

	goto st571
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
//line ragel_parse_datetime.go:21121
		switch data[p] {
		case 32:
			goto st535
		case 39:
			goto st536
		case 43:
			goto st540
		case 45:
			goto st565
		case 47:
			goto tr754
		case 90:
			goto tr757
		case 95:
			goto tr754
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr755
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr754
			}
		default:
			goto tr754
		}
		goto st0
tr757:
//line ragel/datetime.rl:5
 pb = p 
	goto st572
tr748:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st572
tr819:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st572
tr853:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st572
tr861:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st572
tr872:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st572
tr835:
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
	goto st572
tr892:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st572
tr828:
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
	goto st572
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
//line ragel_parse_datetime.go:21260
		switch data[p] {
		case 32:
			goto tr783
		case 47:
			goto st567
		case 95:
			goto st567
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st567
			}
		case data[p] >= 65:
			goto st567
		}
		goto st0
tr758:
//line ragel/datetime.rl:5
 pb = p 
	goto st573
tr749:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st573
tr820:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st573
tr854:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st573
tr862:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st573
tr873:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st573
tr878:
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
	goto st573
tr893:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st573
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
//line ragel_parse_datetime.go:21368
		switch data[p] {
		case 47:
			goto st567
		case 95:
			goto st567
		case 109:
			goto st570
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st567
			}
		case data[p] >= 65:
			goto st567
		}
		goto st0
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
		switch data[p] {
		case 32:
			goto tr812
		case 43:
			goto tr813
		case 45:
			goto tr814
		case 47:
			goto tr815
		case 58:
			goto tr817
		case 65:
			goto tr818
		case 80:
			goto tr818
		case 90:
			goto tr819
		case 95:
			goto tr815
		case 97:
			goto tr820
		case 112:
			goto tr820
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st575
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr815
			}
		default:
			goto tr815
		}
		goto st0
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
		if 48 <= data[p] && data[p] <= 57 {
			goto st576
		}
		goto st0
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
		switch data[p] {
		case 32:
			goto tr822
		case 43:
			goto tr823
		case 45:
			goto tr825
		case 47:
			goto tr826
		case 90:
			goto tr828
		case 95:
			goto tr826
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr824
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr826
				}
			case data[p] >= 65:
				goto tr826
			}
		default:
			goto st587
		}
		goto st0
tr824:
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

	goto st577
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
//line ragel_parse_datetime.go:21496
		if 48 <= data[p] && data[p] <= 57 {
			goto tr829
		}
		goto st0
tr829:
//line ragel/datetime.rl:5
 pb = p 
	goto st578
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
//line ragel_parse_datetime.go:21510
		switch data[p] {
		case 32:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st579
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
		switch data[p] {
		case 32:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st580
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		switch data[p] {
		case 32:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st581
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
		switch data[p] {
		case 32:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st582
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		switch data[p] {
		case 32:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st583
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		switch data[p] {
		case 32:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st584
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		switch data[p] {
		case 32:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st585
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		switch data[p] {
		case 32:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st586
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		switch data[p] {
		case 32:
			goto tr830
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		case data[p] >= 65:
			goto tr833
		}
		goto st0
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
		if 48 <= data[p] && data[p] <= 57 {
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
			goto tr822
		case 43:
			goto tr823
		case 45:
			goto tr825
		case 47:
			goto tr826
		case 90:
			goto tr828
		case 95:
			goto tr826
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr824
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr826
			}
		default:
			goto tr826
		}
		goto st0
tr746:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st589
tr817:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st589
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
//line ragel_parse_datetime.go:21848
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr845
			}
		case data[p] >= 48:
			goto tr844
		}
		goto st0
tr844:
//line ragel/datetime.rl:5
 pb = p 
	goto st590
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
//line ragel_parse_datetime.go:21867
		switch data[p] {
		case 32:
			goto tr846
		case 43:
			goto tr847
		case 45:
			goto tr848
		case 47:
			goto tr849
		case 58:
			goto tr851
		case 65:
			goto tr852
		case 80:
			goto tr852
		case 90:
			goto tr853
		case 95:
			goto tr849
		case 97:
			goto tr854
		case 112:
			goto tr854
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st591
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr849
			}
		default:
			goto tr849
		}
		goto st0
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		switch data[p] {
		case 32:
			goto tr855
		case 43:
			goto tr856
		case 45:
			goto tr857
		case 47:
			goto tr858
		case 58:
			goto tr859
		case 65:
			goto tr860
		case 80:
			goto tr860
		case 90:
			goto tr861
		case 95:
			goto tr858
		case 97:
			goto tr862
		case 112:
			goto tr862
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr858
			}
		case data[p] >= 66:
			goto tr858
		}
		goto st0
tr851:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st592
tr859:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st592
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
//line ragel_parse_datetime.go:21960
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr864
			}
		case data[p] >= 48:
			goto tr863
		}
		goto st0
tr863:
//line ragel/datetime.rl:5
 pb = p 
	goto st593
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
	st_case_593:
//line ragel_parse_datetime.go:21979
		switch data[p] {
		case 32:
			goto tr865
		case 43:
			goto tr866
		case 45:
			goto tr868
		case 47:
			goto tr869
		case 65:
			goto tr871
		case 80:
			goto tr871
		case 90:
			goto tr872
		case 95:
			goto tr869
		case 97:
			goto tr873
		case 112:
			goto tr873
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr867
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr869
				}
			case data[p] >= 66:
				goto tr869
			}
		default:
			goto st604
		}
		goto st0
tr867:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st594
tr888:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st594
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
//line ragel_parse_datetime.go:22037
		if 48 <= data[p] && data[p] <= 57 {
			goto tr874
		}
		goto st0
tr874:
//line ragel/datetime.rl:5
 pb = p 
	goto st595
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
//line ragel_parse_datetime.go:22051
		switch data[p] {
		case 32:
			goto tr875
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr877
		case 80:
			goto tr877
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr878
		case 112:
			goto tr878
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st596
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		switch data[p] {
		case 32:
			goto tr875
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr877
		case 80:
			goto tr877
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr878
		case 112:
			goto tr878
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st597
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
	st_case_597:
		switch data[p] {
		case 32:
			goto tr875
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr877
		case 80:
			goto tr877
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr878
		case 112:
			goto tr878
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st598
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		switch data[p] {
		case 32:
			goto tr875
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr877
		case 80:
			goto tr877
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr878
		case 112:
			goto tr878
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st599
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		switch data[p] {
		case 32:
			goto tr875
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr877
		case 80:
			goto tr877
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr878
		case 112:
			goto tr878
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st600
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		switch data[p] {
		case 32:
			goto tr875
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr877
		case 80:
			goto tr877
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr878
		case 112:
			goto tr878
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st601
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
		switch data[p] {
		case 32:
			goto tr875
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr877
		case 80:
			goto tr877
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr878
		case 112:
			goto tr878
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st602
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
		switch data[p] {
		case 32:
			goto tr875
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr877
		case 80:
			goto tr877
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr878
		case 112:
			goto tr878
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st603
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
		switch data[p] {
		case 32:
			goto tr875
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr877
		case 80:
			goto tr877
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr878
		case 112:
			goto tr878
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		case data[p] >= 66:
			goto tr833
		}
		goto st0
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		switch data[p] {
		case 32:
			goto tr886
		case 43:
			goto tr887
		case 45:
			goto tr889
		case 47:
			goto tr890
		case 65:
			goto tr891
		case 80:
			goto tr891
		case 90:
			goto tr892
		case 95:
			goto tr890
		case 97:
			goto tr893
		case 112:
			goto tr893
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr888
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr890
			}
		default:
			goto tr890
		}
		goto st0
tr864:
//line ragel/datetime.rl:5
 pb = p 
	goto st605
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
//line ragel_parse_datetime.go:22452
		switch data[p] {
		case 32:
			goto tr865
		case 43:
			goto tr866
		case 45:
			goto tr868
		case 47:
			goto tr869
		case 65:
			goto tr871
		case 80:
			goto tr871
		case 90:
			goto tr872
		case 95:
			goto tr869
		case 97:
			goto tr873
		case 112:
			goto tr873
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr867
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr869
			}
		default:
			goto tr869
		}
		goto st0
tr845:
//line ragel/datetime.rl:5
 pb = p 
	goto st606
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
//line ragel_parse_datetime.go:22497
		switch data[p] {
		case 32:
			goto tr846
		case 43:
			goto tr847
		case 45:
			goto tr848
		case 47:
			goto tr849
		case 58:
			goto tr851
		case 65:
			goto tr852
		case 80:
			goto tr852
		case 90:
			goto tr853
		case 95:
			goto tr849
		case 97:
			goto tr854
		case 112:
			goto tr854
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr849
			}
		case data[p] >= 66:
			goto tr849
		}
		goto st0
tr739:
//line ragel/datetime.rl:5
 pb = p 
	goto st607
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
//line ragel_parse_datetime.go:22540
		switch data[p] {
		case 32:
			goto tr741
		case 43:
			goto tr742
		case 45:
			goto tr743
		case 47:
			goto tr744
		case 58:
			goto tr746
		case 65:
			goto tr747
		case 80:
			goto tr747
		case 90:
			goto tr748
		case 95:
			goto tr744
		case 97:
			goto tr749
		case 112:
			goto tr749
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st574
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr744
				}
			case data[p] >= 66:
				goto tr744
			}
		default:
			goto st608
		}
		goto st0
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		if 48 <= data[p] && data[p] <= 57 {
			goto st575
		}
		goto st0
tr740:
//line ragel/datetime.rl:5
 pb = p 
	goto st609
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
//line ragel_parse_datetime.go:22601
		switch data[p] {
		case 32:
			goto tr741
		case 43:
			goto tr742
		case 45:
			goto tr743
		case 47:
			goto tr744
		case 58:
			goto tr746
		case 65:
			goto tr747
		case 80:
			goto tr747
		case 90:
			goto tr748
		case 95:
			goto tr744
		case 97:
			goto tr749
		case 112:
			goto tr749
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st608
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr744
			}
		default:
			goto tr744
		}
		goto st0
tr735:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st610
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
//line ragel_parse_datetime.go:22655
		if data[p] == 100 {
			goto st611
		}
		goto st0
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		if data[p] == 32 {
			goto st532
		}
		goto st0
tr736:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st612
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
//line ragel_parse_datetime.go:22685
		if data[p] == 116 {
			goto st611
		}
		goto st0
tr737:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st613
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
//line ragel_parse_datetime.go:22706
		if data[p] == 104 {
			goto st611
		}
		goto st0
tr730:
//line ragel/datetime.rl:5
 pb = p 
	goto st614
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
//line ragel_parse_datetime.go:22720
		switch data[p] {
		case 32:
			goto tr734
		case 110:
			goto tr735
		case 114:
			goto tr735
		case 115:
			goto tr736
		case 116:
			goto tr737
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st531
		}
		goto st0
tr731:
//line ragel/datetime.rl:5
 pb = p 
	goto st615
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
//line ragel_parse_datetime.go:22746
		switch data[p] {
		case 32:
			goto tr734
		case 110:
			goto tr735
		case 114:
			goto tr735
		case 115:
			goto tr736
		case 116:
			goto tr737
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st531
		}
		goto st0
tr725:
//line ragel/datetime.rl:5
 pb = p 
	goto st616
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
//line ragel_parse_datetime.go:22772
		if 49 <= data[p] && data[p] <= 57 {
			goto st617
		}
		goto st0
tr728:
//line ragel/datetime.rl:5
 pb = p 
	goto st617
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
//line ragel_parse_datetime.go:22786
		switch data[p] {
		case 32:
			goto tr898
		case 110:
			goto tr899
		case 114:
			goto tr899
		case 115:
			goto tr900
		case 116:
			goto tr901
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr610
		}
		goto st0
tr898:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st618
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
//line ragel_parse_datetime.go:22819
		if data[p] == 50 {
			goto tr903
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr904
			}
		case data[p] >= 48:
			goto tr902
		}
		goto st0
tr902:
//line ragel/datetime.rl:5
 pb = p 
	goto st619
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
//line ragel_parse_datetime.go:22841
		switch data[p] {
		case 32:
			goto tr905
		case 43:
			goto tr742
		case 45:
			goto tr743
		case 47:
			goto tr744
		case 58:
			goto tr907
		case 65:
			goto tr908
		case 80:
			goto tr908
		case 90:
			goto tr748
		case 95:
			goto tr744
		case 97:
			goto tr909
		case 112:
			goto tr909
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st625
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr744
			}
		default:
			goto tr744
		}
		goto st0
tr905:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st620
tr914:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st620
tr982:
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

	goto st620
tr965:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st620
tr970:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st620
tr976:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st620
tr993:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st620
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
//line ragel_parse_datetime.go:22951
		switch data[p] {
		case 32:
			goto st535
		case 39:
			goto st207
		case 43:
			goto st540
		case 45:
			goto st565
		case 47:
			goto tr754
		case 65:
			goto tr910
		case 80:
			goto tr910
		case 90:
			goto tr757
		case 95:
			goto tr754
		case 97:
			goto tr911
		case 112:
			goto tr911
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr304
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr754
			}
		default:
			goto tr754
		}
		goto st0
tr910:
//line ragel/datetime.rl:5
 pb = p 
	goto st621
tr908:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st621
tr917:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st621
tr968:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st621
tr972:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st621
tr979:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st621
tr984:
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
	goto st621
tr995:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st621
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
//line ragel_parse_datetime.go:23079
		switch data[p] {
		case 47:
			goto st567
		case 77:
			goto st622
		case 95:
			goto st567
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st567
			}
		case data[p] >= 65:
			goto st567
		}
		goto st0
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		switch data[p] {
		case 32:
			goto tr913
		case 43:
			goto tr809
		case 45:
			goto tr810
		case 47:
			goto tr811
		case 95:
			goto tr811
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr811
			}
		case data[p] >= 65:
			goto tr811
		}
		goto st0
tr913:
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

	goto st623
tr951:
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

	goto st623
tr961:
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

	goto st623
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
//line ragel_parse_datetime.go:23201
		switch data[p] {
		case 32:
			goto st535
		case 39:
			goto st207
		case 43:
			goto st540
		case 45:
			goto st565
		case 47:
			goto tr754
		case 90:
			goto tr757
		case 95:
			goto tr754
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr304
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr754
			}
		default:
			goto tr754
		}
		goto st0
tr911:
//line ragel/datetime.rl:5
 pb = p 
	goto st624
tr909:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st624
tr918:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st624
tr969:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st624
tr973:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st624
tr980:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:5
 pb = p 
	goto st624
tr985:
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
	goto st624
tr996:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:5
 pb = p 
	goto st624
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
//line ragel_parse_datetime.go:23321
		switch data[p] {
		case 47:
			goto st567
		case 95:
			goto st567
		case 109:
			goto st622
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st567
			}
		case data[p] >= 65:
			goto st567
		}
		goto st0
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
		switch data[p] {
		case 32:
			goto tr914
		case 43:
			goto tr813
		case 45:
			goto tr814
		case 47:
			goto tr815
		case 58:
			goto tr916
		case 65:
			goto tr917
		case 80:
			goto tr917
		case 90:
			goto tr819
		case 95:
			goto tr815
		case 97:
			goto tr918
		case 112:
			goto tr918
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st626
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr815
			}
		default:
			goto tr815
		}
		goto st0
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1159
		}
		goto st0
	st1159:
		if p++; p == pe {
			goto _test_eof1159
		}
	st_case_1159:
		switch data[p] {
		case 32:
			goto tr1716
		case 43:
			goto tr1717
		case 44:
			goto tr1718
		case 45:
			goto tr1719
		case 46:
			goto tr962
		case 47:
			goto tr1720
		case 84:
			goto tr1722
		case 90:
			goto tr1723
		case 95:
			goto tr1724
		case 116:
			goto tr1724
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st659
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1720
			}
		default:
			goto tr1720
		}
		goto st0
tr1716:
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

	goto st1160
	st1160:
		if p++; p == pe {
			goto _test_eof1160
		}
	st_case_1160:
//line ragel_parse_datetime.go:23456
		switch data[p] {
		case 32:
			goto st627
		case 39:
			goto st207
		case 43:
			goto st628
		case 45:
			goto st640
		case 47:
			goto tr1728
		case 50:
			goto tr1592
		case 65:
			goto tr1729
		case 66:
			goto tr1730
		case 90:
			goto tr1731
		case 95:
			goto tr1728
		case 97:
			goto tr1732
		case 109:
			goto tr1733
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr1591
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr1728
				}
			case data[p] >= 67:
				goto tr1728
			}
		default:
			goto tr1593
		}
		goto st0
tr1737:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st627
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
//line ragel_parse_datetime.go:23510
		switch data[p] {
		case 39:
			goto st536
		case 65:
			goto st11
		case 66:
			goto st17
		case 109:
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr755
		}
		goto st0
tr1717:
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

	goto st628
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
//line ragel_parse_datetime.go:23551
		if data[p] == 50 {
			goto tr921
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr922
			}
		case data[p] >= 48:
			goto tr920
		}
		goto st0
tr920:
//line ragel/datetime.rl:5
 pb = p 
	goto st1161
tr942:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1161
	st1161:
		if p++; p == pe {
			goto _test_eof1161
		}
	st_case_1161:
//line ragel_parse_datetime.go:23579
		switch data[p] {
		case 32:
			goto tr1734
		case 58:
			goto tr1736
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1173
		}
		goto st0
tr1734:
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
	goto st629
tr1749:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st629
tr1752:
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
	goto st629
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
//line ragel_parse_datetime.go:23647
		switch data[p] {
		case 39:
			goto st536
		case 40:
			goto st630
		case 43:
			goto st632
		case 45:
			goto st634
		case 47:
			goto tr926
		case 65:
			goto tr927
		case 66:
			goto tr928
		case 95:
			goto tr926
		case 109:
			goto tr929
		}
		switch {
		case data[p] < 67:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr755
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr926
			}
		default:
			goto tr926
		}
		goto st0
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
		switch data[p] {
		case 32:
			goto st631
		case 43:
			goto st631
		case 45:
			goto st631
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st631
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st631
			}
		default:
			goto st631
		}
		goto st0
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		switch data[p] {
		case 32:
			goto st631
		case 41:
			goto st1162
		case 43:
			goto st631
		case 45:
			goto st631
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st631
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st631
			}
		default:
			goto st631
		}
		goto st0
	st1162:
		if p++; p == pe {
			goto _test_eof1162
		}
	st_case_1162:
		if data[p] == 32 {
			goto tr1737
		}
		goto st0
tr1754:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st632
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
//line ragel_parse_datetime.go:23756
		if data[p] == 50 {
			goto tr933
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr934
			}
		case data[p] >= 48:
			goto tr932
		}
		goto st0
tr932:
//line ragel/datetime.rl:5
 pb = p 
	goto st1163
tr935:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1163
	st1163:
		if p++; p == pe {
			goto _test_eof1163
		}
	st_case_1163:
//line ragel_parse_datetime.go:23784
		switch data[p] {
		case 32:
			goto tr1738
		case 58:
			goto tr1740
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1164
		}
		goto st0
tr1738:
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
	goto st633
tr1742:
//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st633
tr1745:
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
	goto st633
tr1747:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
	goto st633
	st633:
		if p++; p == pe {
			goto _test_eof633
		}
	st_case_633:
//line ragel_parse_datetime.go:23861
		switch data[p] {
		case 39:
			goto st536
		case 40:
			goto st630
		case 65:
			goto st11
		case 66:
			goto st17
		case 109:
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr755
		}
		goto st0
tr934:
//line ragel/datetime.rl:5
 pb = p 
	goto st1164
tr937:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1164
	st1164:
		if p++; p == pe {
			goto _test_eof1164
		}
	st_case_1164:
//line ragel_parse_datetime.go:23893
		switch data[p] {
		case 32:
			goto tr1738
		case 58:
			goto tr1740
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1165
		}
		goto st0
	st1165:
		if p++; p == pe {
			goto _test_eof1165
		}
	st_case_1165:
		if data[p] == 32 {
			goto tr1738
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1165
		}
		goto st0
tr1740:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st1166
	st1166:
		if p++; p == pe {
			goto _test_eof1166
		}
	st_case_1166:
//line ragel_parse_datetime.go:23933
		if data[p] == 32 {
			goto tr1742
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1744
			}
		case data[p] >= 48:
			goto tr1743
		}
		goto st0
tr1743:
//line ragel/datetime.rl:5
 pb = p 
	goto st1167
	st1167:
		if p++; p == pe {
			goto _test_eof1167
		}
	st_case_1167:
//line ragel_parse_datetime.go:23955
		if data[p] == 32 {
			goto tr1745
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1168
		}
		goto st0
tr1744:
//line ragel/datetime.rl:5
 pb = p 
	goto st1168
	st1168:
		if p++; p == pe {
			goto _test_eof1168
		}
	st_case_1168:
//line ragel_parse_datetime.go:23972
		if data[p] == 32 {
			goto tr1745
		}
		goto st0
tr933:
//line ragel/datetime.rl:5
 pb = p 
	goto st1169
tr936:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1169
	st1169:
		if p++; p == pe {
			goto _test_eof1169
		}
	st_case_1169:
//line ragel_parse_datetime.go:23992
		switch data[p] {
		case 32:
			goto tr1738
		case 58:
			goto tr1740
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st1165
			}
		case data[p] >= 48:
			goto st1164
		}
		goto st0
tr1755:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

	goto st634
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
//line ragel_parse_datetime.go:24020
		if data[p] == 50 {
			goto tr936
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr937
			}
		case data[p] >= 48:
			goto tr935
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
//line ragel_parse_datetime.go:24042
		switch data[p] {
		case 47:
			goto st636
		case 95:
			goto st636
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st636
			}
		case data[p] >= 65:
			goto st636
		}
		goto st0
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
		switch data[p] {
		case 47:
			goto st1170
		case 95:
			goto st1170
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1170
			}
		case data[p] >= 65:
			goto st1170
		}
		goto st0
	st1170:
		if p++; p == pe {
			goto _test_eof1170
		}
	st_case_1170:
		switch data[p] {
		case 32:
			goto tr1747
		case 47:
			goto st1170
		case 95:
			goto st1170
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1170
			}
		case data[p] >= 65:
			goto st1170
		}
		goto st0
tr927:
//line ragel/datetime.rl:5
 pb = p 
	goto st637
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
//line ragel_parse_datetime.go:24109
		switch data[p] {
		case 47:
			goto st636
		case 68:
			goto st1171
		case 95:
			goto st636
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st636
			}
		case data[p] >= 65:
			goto st636
		}
		goto st0
	st1171:
		if p++; p == pe {
			goto _test_eof1171
		}
	st_case_1171:
		switch data[p] {
		case 32:
			goto st12
		case 47:
			goto st1170
		case 95:
			goto st1170
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1170
			}
		case data[p] >= 65:
			goto st1170
		}
		goto st0
tr928:
//line ragel/datetime.rl:5
 pb = p 
	goto st638
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
//line ragel_parse_datetime.go:24158
		switch data[p] {
		case 47:
			goto st636
		case 67:
			goto st1172
		case 95:
			goto st636
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st636
			}
		case data[p] >= 65:
			goto st636
		}
		goto st0
	st1172:
		if p++; p == pe {
			goto _test_eof1172
		}
	st_case_1172:
		switch data[p] {
		case 32:
			goto tr1376
		case 47:
			goto st1170
		case 95:
			goto st1170
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1170
			}
		case data[p] >= 65:
			goto st1170
		}
		goto st0
tr929:
//line ragel/datetime.rl:5
 pb = p 
	goto st639
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
//line ragel_parse_datetime.go:24207
		switch data[p] {
		case 47:
			goto st636
		case 61:
			goto st14
		case 95:
			goto st636
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st636
			}
		case data[p] >= 65:
			goto st636
		}
		goto st0
tr922:
//line ragel/datetime.rl:5
 pb = p 
	goto st1173
tr944:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1173
	st1173:
		if p++; p == pe {
			goto _test_eof1173
		}
	st_case_1173:
//line ragel_parse_datetime.go:24240
		switch data[p] {
		case 32:
			goto tr1734
		case 58:
			goto tr1736
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1174
		}
		goto st0
	st1174:
		if p++; p == pe {
			goto _test_eof1174
		}
	st_case_1174:
		if data[p] == 32 {
			goto tr1734
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1174
		}
		goto st0
tr1736:
//line ragel/datetime.rl:177

    st.ZoneOffsetIsValid = true
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st1175
	st1175:
		if p++; p == pe {
			goto _test_eof1175
		}
	st_case_1175:
//line ragel_parse_datetime.go:24280
		if data[p] == 32 {
			goto tr1749
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1751
			}
		case data[p] >= 48:
			goto tr1750
		}
		goto st0
tr1750:
//line ragel/datetime.rl:5
 pb = p 
	goto st1176
	st1176:
		if p++; p == pe {
			goto _test_eof1176
		}
	st_case_1176:
//line ragel_parse_datetime.go:24302
		if data[p] == 32 {
			goto tr1752
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1177
		}
		goto st0
tr1751:
//line ragel/datetime.rl:5
 pb = p 
	goto st1177
	st1177:
		if p++; p == pe {
			goto _test_eof1177
		}
	st_case_1177:
//line ragel_parse_datetime.go:24319
		if data[p] == 32 {
			goto tr1752
		}
		goto st0
tr921:
//line ragel/datetime.rl:5
 pb = p 
	goto st1178
tr943:
//line ragel/datetime.rl:175
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:5
 pb = p 
	goto st1178
	st1178:
		if p++; p == pe {
			goto _test_eof1178
		}
	st_case_1178:
//line ragel_parse_datetime.go:24339
		switch data[p] {
		case 32:
			goto tr1734
		case 58:
			goto tr1736
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st1174
			}
		case data[p] >= 48:
			goto st1173
		}
		goto st0
tr1719:
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

	goto st640
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
//line ragel_parse_datetime.go:24381
		if data[p] == 50 {
			goto tr943
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr944
			}
		case data[p] >= 48:
			goto tr942
		}
		goto st0
tr1728:
//line ragel/datetime.rl:5
 pb = p 
	goto st641
tr1720:
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

	goto st641
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
//line ragel_parse_datetime.go:24426
		switch data[p] {
		case 47:
			goto st642
		case 95:
			goto st642
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st642
			}
		case data[p] >= 65:
			goto st642
		}
		goto st0
tr1756:
//line ragel/datetime.rl:5
 pb = p 
	goto st642
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
//line ragel_parse_datetime.go:24451
		switch data[p] {
		case 47:
			goto st1179
		case 95:
			goto st1179
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1179
			}
		case data[p] >= 65:
			goto st1179
		}
		goto st0
	st1179:
		if p++; p == pe {
			goto _test_eof1179
		}
	st_case_1179:
		switch data[p] {
		case 32:
			goto tr1747
		case 43:
			goto tr1754
		case 45:
			goto tr1755
		case 47:
			goto st1179
		case 95:
			goto st1179
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1179
			}
		case data[p] >= 65:
			goto st1179
		}
		goto st0
tr1729:
//line ragel/datetime.rl:5
 pb = p 
	goto st643
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
//line ragel_parse_datetime.go:24502
		switch data[p] {
		case 47:
			goto st642
		case 68:
			goto st1180
		case 84:
			goto st644
		case 95:
			goto st642
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st642
			}
		case data[p] >= 65:
			goto st642
		}
		goto st0
	st1180:
		if p++; p == pe {
			goto _test_eof1180
		}
	st_case_1180:
		switch data[p] {
		case 32:
			goto st12
		case 47:
			goto st1179
		case 95:
			goto st1179
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1179
			}
		case data[p] >= 65:
			goto st1179
		}
		goto st0
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
		switch data[p] {
		case 32:
			goto st48
		case 47:
			goto st1179
		case 95:
			goto st1179
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1179
			}
		case data[p] >= 65:
			goto st1179
		}
		goto st0
tr1730:
//line ragel/datetime.rl:5
 pb = p 
	goto st645
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
//line ragel_parse_datetime.go:24575
		switch data[p] {
		case 47:
			goto st642
		case 67:
			goto st1181
		case 95:
			goto st642
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st642
			}
		case data[p] >= 65:
			goto st642
		}
		goto st0
	st1181:
		if p++; p == pe {
			goto _test_eof1181
		}
	st_case_1181:
		switch data[p] {
		case 32:
			goto tr1376
		case 47:
			goto st1179
		case 95:
			goto st1179
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1179
			}
		case data[p] >= 65:
			goto st1179
		}
		goto st0
tr1731:
//line ragel/datetime.rl:5
 pb = p 
	goto st1182
tr1723:
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

	goto st1182
	st1182:
		if p++; p == pe {
			goto _test_eof1182
		}
	st_case_1182:
//line ragel_parse_datetime.go:24647
		switch data[p] {
		case 32:
			goto tr1742
		case 47:
			goto st642
		case 95:
			goto st642
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st642
			}
		case data[p] >= 65:
			goto st642
		}
		goto st0
tr1732:
//line ragel/datetime.rl:5
 pb = p 
	goto st646
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
//line ragel_parse_datetime.go:24674
		switch data[p] {
		case 47:
			goto st642
		case 95:
			goto st642
		case 116:
			goto st644
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st642
			}
		case data[p] >= 65:
			goto st642
		}
		goto st0
tr1733:
//line ragel/datetime.rl:5
 pb = p 
	goto st647
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
//line ragel_parse_datetime.go:24701
		switch data[p] {
		case 47:
			goto st642
		case 61:
			goto st14
		case 95:
			goto st642
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st642
			}
		case data[p] >= 65:
			goto st642
		}
		goto st0
tr1718:
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

	goto st648
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
//line ragel_parse_datetime.go:24745
		if data[p] == 32 {
			goto st48
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr950
		}
		goto st0
tr950:
//line ragel/datetime.rl:5
 pb = p 
	goto st649
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
//line ragel_parse_datetime.go:24762
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st650
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st651
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st652
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st653
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st654
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st655
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st656
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st657
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st657:
		if p++; p == pe {
			goto _test_eof657
		}
	st_case_657:
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 90:
			goto tr835
		case 95:
			goto tr833
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		case data[p] >= 65:
			goto tr833
		}
		goto st0
tr962:
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

	goto st658
	st658:
		if p++; p == pe {
			goto _test_eof658
		}
	st_case_658:
//line ragel_parse_datetime.go:25064
		if 48 <= data[p] && data[p] <= 57 {
			goto tr950
		}
		goto st0
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
		if 48 <= data[p] && data[p] <= 57 {
			goto st660
		}
		goto st0
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
		switch data[p] {
		case 32:
			goto tr961
		case 43:
			goto tr823
		case 45:
			goto tr825
		case 47:
			goto tr826
		case 90:
			goto tr828
		case 95:
			goto tr826
		}
		switch {
		case data[p] < 65:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr962
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr826
			}
		default:
			goto tr826
		}
		goto st0
tr1722:
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

	goto st1183
	st1183:
		if p++; p == pe {
			goto _test_eof1183
		}
	st_case_1183:
//line ragel_parse_datetime.go:25138
		switch data[p] {
		case 32:
			goto st10
		case 43:
			goto st18
		case 45:
			goto st30
		case 47:
			goto tr1756
		case 50:
			goto tr93
		case 90:
			goto tr1757
		case 95:
			goto tr1756
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr92
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr1756
				}
			case data[p] >= 65:
				goto tr1756
			}
		default:
			goto tr94
		}
		goto st0
tr1757:
//line ragel/datetime.rl:5
 pb = p 
	goto st1184
	st1184:
		if p++; p == pe {
			goto _test_eof1184
		}
	st_case_1184:
//line ragel_parse_datetime.go:25182
		switch data[p] {
		case 32:
			goto tr1385
		case 47:
			goto st1179
		case 95:
			goto st1179
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1179
			}
		case data[p] >= 65:
			goto st1179
		}
		goto st0
tr1724:
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

	goto st661
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
	st_case_661:
//line ragel_parse_datetime.go:25228
		switch data[p] {
		case 47:
			goto st642
		case 50:
			goto tr93
		case 95:
			goto st642
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr92
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st642
				}
			case data[p] >= 65:
				goto st642
			}
		default:
			goto tr94
		}
		goto st0
tr907:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st662
tr916:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st662
	st662:
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
//line ragel_parse_datetime.go:25272
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr964
			}
		case data[p] >= 48:
			goto tr963
		}
		goto st0
tr963:
//line ragel/datetime.rl:5
 pb = p 
	goto st663
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
//line ragel_parse_datetime.go:25291
		switch data[p] {
		case 32:
			goto tr965
		case 43:
			goto tr847
		case 45:
			goto tr848
		case 47:
			goto tr849
		case 58:
			goto tr967
		case 65:
			goto tr968
		case 80:
			goto tr968
		case 90:
			goto tr853
		case 95:
			goto tr849
		case 97:
			goto tr969
		case 112:
			goto tr969
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st664
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr849
			}
		default:
			goto tr849
		}
		goto st0
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
		switch data[p] {
		case 32:
			goto tr970
		case 43:
			goto tr856
		case 45:
			goto tr857
		case 47:
			goto tr858
		case 58:
			goto tr971
		case 65:
			goto tr972
		case 80:
			goto tr972
		case 90:
			goto tr861
		case 95:
			goto tr858
		case 97:
			goto tr973
		case 112:
			goto tr973
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr858
			}
		case data[p] >= 66:
			goto tr858
		}
		goto st0
tr967:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st665
tr971:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st665
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
//line ragel_parse_datetime.go:25384
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr975
			}
		case data[p] >= 48:
			goto tr974
		}
		goto st0
tr974:
//line ragel/datetime.rl:5
 pb = p 
	goto st666
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
//line ragel_parse_datetime.go:25403
		switch data[p] {
		case 32:
			goto tr976
		case 43:
			goto tr866
		case 45:
			goto tr868
		case 47:
			goto tr869
		case 65:
			goto tr979
		case 80:
			goto tr979
		case 90:
			goto tr872
		case 95:
			goto tr869
		case 97:
			goto tr980
		case 112:
			goto tr980
		}
		switch {
		case data[p] < 48:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr977
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr869
				}
			case data[p] >= 66:
				goto tr869
			}
		default:
			goto st677
		}
		goto st0
tr977:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st667
tr994:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st667
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
	st_case_667:
//line ragel_parse_datetime.go:25461
		if 48 <= data[p] && data[p] <= 57 {
			goto tr981
		}
		goto st0
tr981:
//line ragel/datetime.rl:5
 pb = p 
	goto st668
	st668:
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
//line ragel_parse_datetime.go:25475
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr985
		case 112:
			goto tr985
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st669
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr985
		case 112:
			goto tr985
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st670
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr985
		case 112:
			goto tr985
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st671
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr985
		case 112:
			goto tr985
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st672
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
	st_case_672:
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr985
		case 112:
			goto tr985
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st673
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st673:
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr985
		case 112:
			goto tr985
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st674
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st674:
		if p++; p == pe {
			goto _test_eof674
		}
	st_case_674:
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr985
		case 112:
			goto tr985
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st675
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st675:
		if p++; p == pe {
			goto _test_eof675
		}
	st_case_675:
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr985
		case 112:
			goto tr985
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st676
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		default:
			goto tr833
		}
		goto st0
	st676:
		if p++; p == pe {
			goto _test_eof676
		}
	st_case_676:
		switch data[p] {
		case 32:
			goto tr982
		case 43:
			goto tr831
		case 45:
			goto tr832
		case 47:
			goto tr833
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr835
		case 95:
			goto tr833
		case 97:
			goto tr985
		case 112:
			goto tr985
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr833
			}
		case data[p] >= 66:
			goto tr833
		}
		goto st0
	st677:
		if p++; p == pe {
			goto _test_eof677
		}
	st_case_677:
		switch data[p] {
		case 32:
			goto tr993
		case 43:
			goto tr887
		case 45:
			goto tr889
		case 47:
			goto tr890
		case 65:
			goto tr995
		case 80:
			goto tr995
		case 90:
			goto tr892
		case 95:
			goto tr890
		case 97:
			goto tr996
		case 112:
			goto tr996
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr994
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr890
			}
		default:
			goto tr890
		}
		goto st0
tr975:
//line ragel/datetime.rl:5
 pb = p 
	goto st678
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
//line ragel_parse_datetime.go:25876
		switch data[p] {
		case 32:
			goto tr976
		case 43:
			goto tr866
		case 45:
			goto tr868
		case 47:
			goto tr869
		case 65:
			goto tr979
		case 80:
			goto tr979
		case 90:
			goto tr872
		case 95:
			goto tr869
		case 97:
			goto tr980
		case 112:
			goto tr980
		}
		switch {
		case data[p] < 66:
			if 44 <= data[p] && data[p] <= 46 {
				goto tr977
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr869
			}
		default:
			goto tr869
		}
		goto st0
tr964:
//line ragel/datetime.rl:5
 pb = p 
	goto st679
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
//line ragel_parse_datetime.go:25921
		switch data[p] {
		case 32:
			goto tr965
		case 43:
			goto tr847
		case 45:
			goto tr848
		case 47:
			goto tr849
		case 58:
			goto tr967
		case 65:
			goto tr968
		case 80:
			goto tr968
		case 90:
			goto tr853
		case 95:
			goto tr849
		case 97:
			goto tr969
		case 112:
			goto tr969
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr849
			}
		case data[p] >= 66:
			goto tr849
		}
		goto st0
tr903:
//line ragel/datetime.rl:5
 pb = p 
	goto st680
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
//line ragel_parse_datetime.go:25964
		switch data[p] {
		case 32:
			goto tr905
		case 43:
			goto tr742
		case 45:
			goto tr743
		case 47:
			goto tr744
		case 58:
			goto tr907
		case 65:
			goto tr908
		case 80:
			goto tr908
		case 90:
			goto tr748
		case 95:
			goto tr744
		case 97:
			goto tr909
		case 112:
			goto tr909
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st625
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr744
				}
			case data[p] >= 66:
				goto tr744
			}
		default:
			goto st681
		}
		goto st0
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
	st_case_681:
		if 48 <= data[p] && data[p] <= 57 {
			goto st626
		}
		goto st0
tr904:
//line ragel/datetime.rl:5
 pb = p 
	goto st682
	st682:
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
//line ragel_parse_datetime.go:26025
		switch data[p] {
		case 32:
			goto tr905
		case 43:
			goto tr742
		case 45:
			goto tr743
		case 47:
			goto tr744
		case 58:
			goto tr907
		case 65:
			goto tr908
		case 80:
			goto tr908
		case 90:
			goto tr748
		case 95:
			goto tr744
		case 97:
			goto tr909
		case 112:
			goto tr909
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st681
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr744
			}
		default:
			goto tr744
		}
		goto st0
tr899:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st683
	st683:
		if p++; p == pe {
			goto _test_eof683
		}
	st_case_683:
//line ragel_parse_datetime.go:26079
		if data[p] == 100 {
			goto st684
		}
		goto st0
	st684:
		if p++; p == pe {
			goto _test_eof684
		}
	st_case_684:
		if data[p] == 32 {
			goto st618
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto st198
		}
		goto st0
tr900:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st685
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
	st_case_685:
//line ragel_parse_datetime.go:26112
		if data[p] == 116 {
			goto st684
		}
		goto st0
tr901:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st686
	st686:
		if p++; p == pe {
			goto _test_eof686
		}
	st_case_686:
//line ragel_parse_datetime.go:26133
		if data[p] == 104 {
			goto st684
		}
		goto st0
tr726:
//line ragel/datetime.rl:5
 pb = p 
	goto st687
	st687:
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
//line ragel_parse_datetime.go:26147
		switch data[p] {
		case 32:
			goto tr898
		case 110:
			goto tr899
		case 114:
			goto tr899
		case 115:
			goto tr900
		case 116:
			goto tr901
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st617
			}
		case data[p] >= 45:
			goto tr610
		}
		goto st0
tr727:
//line ragel/datetime.rl:5
 pb = p 
	goto st688
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
//line ragel_parse_datetime.go:26178
		switch data[p] {
		case 32:
			goto tr898
		case 110:
			goto tr899
		case 114:
			goto tr899
		case 115:
			goto tr900
		case 116:
			goto tr901
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st617
			}
		case data[p] >= 45:
			goto tr610
		}
		goto st0
tr722:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st689
tr1004:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st689
tr1011:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st689
tr1020:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st689
tr1031:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st689
tr1040:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st689
tr1044:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st689
tr1051:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st689
tr1056:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st689
tr1061:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st689
tr1071:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st689
tr1080:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st689
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
//line ragel_parse_datetime.go:26253
		switch data[p] {
		case 32:
			goto st528
		case 48:
			goto tr653
		case 51:
			goto tr655
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st479
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr656
			}
		default:
			goto tr654
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
			goto tr721
		case 46:
			goto tr722
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr600
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
			goto tr1003
		case 46:
			goto tr1004
		case 117:
			goto st694
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr669
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
			goto tr1003
		case 46:
			goto tr1004
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr669
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
			goto tr1010
		case 46:
			goto tr1011
		case 101:
			goto st700
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr677
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
			goto tr1010
		case 46:
			goto tr1011
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr677
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
			goto tr1019
		case 46:
			goto tr1020
		case 114:
			goto st708
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr688
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
			goto tr1019
		case 46:
			goto tr1020
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr688
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
			goto tr1029
		case 46:
			goto tr1031
		case 117:
			goto st716
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1030
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
			goto tr1029
		case 46:
			goto tr1031
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1030
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
			goto tr1038
		case 46:
			goto tr1040
		case 121:
			goto st722
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1039
		}
		goto st0
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
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
	st723:
		if p++; p == pe {
			goto _test_eof723
		}
	st_case_723:
		switch data[p] {
		case 32:
			goto tr1042
		case 46:
			goto tr1044
		case 101:
			goto st724
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1043
		}
		goto st0
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
	st_case_724:
		switch data[p] {
		case 32:
			goto tr1042
		case 46:
			goto tr1044
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1043
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
			goto tr1049
		case 46:
			goto tr1051
		case 99:
			goto st728
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1050
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
			goto tr1049
		case 46:
			goto tr1051
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1050
		}
		goto st0
	st730:
		if p++; p == pe {
			goto _test_eof730
		}
	st_case_730:
		switch data[p] {
		case 32:
			goto tr1054
		case 46:
			goto tr1056
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1055
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
			goto tr1059
		case 46:
			goto tr1061
		case 101:
			goto st734
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1060
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
			goto tr1059
		case 46:
			goto tr1061
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1060
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
			goto tr1069
		case 46:
			goto tr1071
		case 111:
			goto st742
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1070
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
			goto tr1069
		case 46:
			goto tr1071
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1070
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
			goto tr1078
		case 46:
			goto tr1080
		case 116:
			goto st749
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1079
		}
		goto st0
	st749:
		if p++; p == pe {
			goto _test_eof749
		}
	st_case_749:
		switch data[p] {
		case 32:
			goto tr1078
		case 46:
			goto tr1080
		case 101:
			goto st750
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1079
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
			goto tr1078
		case 46:
			goto tr1080
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1079
		}
		goto st0
	st755:
		if p++; p == pe {
			goto _test_eof755
		}
	st_case_755:
		if data[p] == 32 {
			goto st756
		}
		goto st0
	st756:
		if p++; p == pe {
			goto _test_eof756
		}
	st_case_756:
		switch data[p] {
		case 48:
			goto tr1088
		case 51:
			goto tr1090
		case 65:
			goto st835
		case 68:
			goto st855
		case 70:
			goto st863
		case 74:
			goto st871
		case 77:
			goto st883
		case 78:
			goto st889
		case 79:
			goto st897
		case 83:
			goto st904
		case 97:
			goto st835
		case 100:
			goto st855
		case 102:
			goto st863
		case 106:
			goto st871
		case 109:
			goto st883
		case 110:
			goto st889
		case 111:
			goto st897
		case 115:
			goto st904
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr1091
			}
		case data[p] >= 49:
			goto tr1089
		}
		goto st0
tr1088:
//line ragel/datetime.rl:5
 pb = p 
	goto st757
	st757:
		if p++; p == pe {
			goto _test_eof757
		}
	st_case_757:
//line ragel_parse_datetime.go:27099
		if 49 <= data[p] && data[p] <= 57 {
			goto st758
		}
		goto st0
tr1091:
//line ragel/datetime.rl:5
 pb = p 
	goto st758
	st758:
		if p++; p == pe {
			goto _test_eof758
		}
	st_case_758:
//line ragel_parse_datetime.go:27113
		switch data[p] {
		case 32:
			goto tr595
		case 45:
			goto tr1101
		case 110:
			goto tr1102
		case 114:
			goto tr1102
		case 115:
			goto tr1103
		case 116:
			goto tr1104
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto tr595
		}
		goto st0
tr1101:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st759
	st759:
		if p++; p == pe {
			goto _test_eof759
		}
	st_case_759:
//line ragel_parse_datetime.go:27148
		switch data[p] {
		case 65:
			goto st760
		case 68:
			goto st771
		case 70:
			goto st779
		case 74:
			goto st787
		case 77:
			goto st799
		case 78:
			goto st805
		case 79:
			goto st813
		case 83:
			goto st820
		case 97:
			goto st760
		case 100:
			goto st771
		case 102:
			goto st779
		case 106:
			goto st787
		case 109:
			goto st799
		case 110:
			goto st805
		case 111:
			goto st813
		case 115:
			goto st820
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr592
		}
		goto st0
	st760:
		if p++; p == pe {
			goto _test_eof760
		}
	st_case_760:
		switch data[p] {
		case 112:
			goto st761
		case 117:
			goto st766
		}
		goto st0
	st761:
		if p++; p == pe {
			goto _test_eof761
		}
	st_case_761:
		if data[p] == 114 {
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
			goto tr410
		case 45:
			goto tr409
		case 46:
			goto tr1116
		case 47:
			goto tr410
		case 105:
			goto st764
		}
		goto st0
tr1116:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st763
tr1120:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st763
tr1126:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st763
tr1134:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st763
tr1143:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st763
tr1150:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st763
tr1152:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st763
tr1157:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st763
tr1160:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st763
tr1163:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st763
tr1171:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st763
tr1178:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st763
	st763:
		if p++; p == pe {
			goto _test_eof763
		}
	st_case_763:
//line ragel_parse_datetime.go:27279
		switch data[p] {
		case 32:
			goto st198
		case 45:
			goto st285
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr287
			}
		case data[p] >= 46:
			goto st198
		}
		goto st0
	st764:
		if p++; p == pe {
			goto _test_eof764
		}
	st_case_764:
		if data[p] == 108 {
			goto st765
		}
		goto st0
	st765:
		if p++; p == pe {
			goto _test_eof765
		}
	st_case_765:
		switch data[p] {
		case 32:
			goto tr410
		case 45:
			goto tr409
		case 46:
			goto tr1116
		case 47:
			goto tr410
		}
		goto st0
	st766:
		if p++; p == pe {
			goto _test_eof766
		}
	st_case_766:
		if data[p] == 103 {
			goto st767
		}
		goto st0
	st767:
		if p++; p == pe {
			goto _test_eof767
		}
	st_case_767:
		switch data[p] {
		case 32:
			goto tr423
		case 45:
			goto tr422
		case 46:
			goto tr1120
		case 47:
			goto tr423
		case 117:
			goto st768
		}
		goto st0
	st768:
		if p++; p == pe {
			goto _test_eof768
		}
	st_case_768:
		if data[p] == 115 {
			goto st769
		}
		goto st0
	st769:
		if p++; p == pe {
			goto _test_eof769
		}
	st_case_769:
		if data[p] == 116 {
			goto st770
		}
		goto st0
	st770:
		if p++; p == pe {
			goto _test_eof770
		}
	st_case_770:
		switch data[p] {
		case 32:
			goto tr423
		case 45:
			goto tr422
		case 46:
			goto tr1120
		case 47:
			goto tr423
		}
		goto st0
	st771:
		if p++; p == pe {
			goto _test_eof771
		}
	st_case_771:
		if data[p] == 101 {
			goto st772
		}
		goto st0
	st772:
		if p++; p == pe {
			goto _test_eof772
		}
	st_case_772:
		if data[p] == 99 {
			goto st773
		}
		goto st0
	st773:
		if p++; p == pe {
			goto _test_eof773
		}
	st_case_773:
		switch data[p] {
		case 32:
			goto tr431
		case 45:
			goto tr430
		case 46:
			goto tr1126
		case 47:
			goto tr431
		case 101:
			goto st774
		}
		goto st0
	st774:
		if p++; p == pe {
			goto _test_eof774
		}
	st_case_774:
		if data[p] == 109 {
			goto st775
		}
		goto st0
	st775:
		if p++; p == pe {
			goto _test_eof775
		}
	st_case_775:
		if data[p] == 98 {
			goto st776
		}
		goto st0
	st776:
		if p++; p == pe {
			goto _test_eof776
		}
	st_case_776:
		if data[p] == 101 {
			goto st777
		}
		goto st0
	st777:
		if p++; p == pe {
			goto _test_eof777
		}
	st_case_777:
		if data[p] == 114 {
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
			goto tr431
		case 45:
			goto tr430
		case 46:
			goto tr1126
		case 47:
			goto tr431
		}
		goto st0
	st779:
		if p++; p == pe {
			goto _test_eof779
		}
	st_case_779:
		if data[p] == 101 {
			goto st780
		}
		goto st0
	st780:
		if p++; p == pe {
			goto _test_eof780
		}
	st_case_780:
		if data[p] == 98 {
			goto st781
		}
		goto st0
	st781:
		if p++; p == pe {
			goto _test_eof781
		}
	st_case_781:
		switch data[p] {
		case 32:
			goto tr441
		case 45:
			goto tr440
		case 46:
			goto tr1134
		case 47:
			goto tr441
		case 114:
			goto st782
		}
		goto st0
	st782:
		if p++; p == pe {
			goto _test_eof782
		}
	st_case_782:
		if data[p] == 117 {
			goto st783
		}
		goto st0
	st783:
		if p++; p == pe {
			goto _test_eof783
		}
	st_case_783:
		if data[p] == 97 {
			goto st784
		}
		goto st0
	st784:
		if p++; p == pe {
			goto _test_eof784
		}
	st_case_784:
		if data[p] == 114 {
			goto st785
		}
		goto st0
	st785:
		if p++; p == pe {
			goto _test_eof785
		}
	st_case_785:
		if data[p] == 121 {
			goto st786
		}
		goto st0
	st786:
		if p++; p == pe {
			goto _test_eof786
		}
	st_case_786:
		switch data[p] {
		case 32:
			goto tr441
		case 45:
			goto tr440
		case 46:
			goto tr1134
		case 47:
			goto tr441
		}
		goto st0
	st787:
		if p++; p == pe {
			goto _test_eof787
		}
	st_case_787:
		switch data[p] {
		case 97:
			goto st788
		case 117:
			goto st794
		}
		goto st0
	st788:
		if p++; p == pe {
			goto _test_eof788
		}
	st_case_788:
		if data[p] == 110 {
			goto st789
		}
		goto st0
	st789:
		if p++; p == pe {
			goto _test_eof789
		}
	st_case_789:
		switch data[p] {
		case 32:
			goto tr452
		case 45:
			goto tr451
		case 46:
			goto tr1143
		case 47:
			goto tr452
		case 117:
			goto st790
		}
		goto st0
	st790:
		if p++; p == pe {
			goto _test_eof790
		}
	st_case_790:
		if data[p] == 97 {
			goto st791
		}
		goto st0
	st791:
		if p++; p == pe {
			goto _test_eof791
		}
	st_case_791:
		if data[p] == 114 {
			goto st792
		}
		goto st0
	st792:
		if p++; p == pe {
			goto _test_eof792
		}
	st_case_792:
		if data[p] == 121 {
			goto st793
		}
		goto st0
	st793:
		if p++; p == pe {
			goto _test_eof793
		}
	st_case_793:
		switch data[p] {
		case 32:
			goto tr452
		case 45:
			goto tr451
		case 46:
			goto tr1143
		case 47:
			goto tr452
		}
		goto st0
	st794:
		if p++; p == pe {
			goto _test_eof794
		}
	st_case_794:
		switch data[p] {
		case 108:
			goto st795
		case 110:
			goto st797
		}
		goto st0
	st795:
		if p++; p == pe {
			goto _test_eof795
		}
	st_case_795:
		switch data[p] {
		case 32:
			goto tr461
		case 45:
			goto tr460
		case 46:
			goto tr1150
		case 47:
			goto tr461
		case 121:
			goto st796
		}
		goto st0
	st796:
		if p++; p == pe {
			goto _test_eof796
		}
	st_case_796:
		switch data[p] {
		case 32:
			goto tr461
		case 45:
			goto tr460
		case 46:
			goto tr1150
		case 47:
			goto tr461
		}
		goto st0
	st797:
		if p++; p == pe {
			goto _test_eof797
		}
	st_case_797:
		switch data[p] {
		case 32:
			goto tr465
		case 45:
			goto tr464
		case 46:
			goto tr1152
		case 47:
			goto tr465
		case 101:
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
			goto tr465
		case 45:
			goto tr464
		case 46:
			goto tr1152
		case 47:
			goto tr465
		}
		goto st0
	st799:
		if p++; p == pe {
			goto _test_eof799
		}
	st_case_799:
		if data[p] == 97 {
			goto st800
		}
		goto st0
	st800:
		if p++; p == pe {
			goto _test_eof800
		}
	st_case_800:
		switch data[p] {
		case 114:
			goto st801
		case 121:
			goto st804
		}
		goto st0
	st801:
		if p++; p == pe {
			goto _test_eof801
		}
	st_case_801:
		switch data[p] {
		case 32:
			goto tr472
		case 45:
			goto tr471
		case 46:
			goto tr1157
		case 47:
			goto tr472
		case 99:
			goto st802
		}
		goto st0
	st802:
		if p++; p == pe {
			goto _test_eof802
		}
	st_case_802:
		if data[p] == 104 {
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
			goto tr472
		case 45:
			goto tr471
		case 46:
			goto tr1157
		case 47:
			goto tr472
		}
		goto st0
	st804:
		if p++; p == pe {
			goto _test_eof804
		}
	st_case_804:
		switch data[p] {
		case 32:
			goto tr477
		case 45:
			goto tr476
		case 46:
			goto tr1160
		case 47:
			goto tr477
		}
		goto st0
	st805:
		if p++; p == pe {
			goto _test_eof805
		}
	st_case_805:
		if data[p] == 111 {
			goto st806
		}
		goto st0
	st806:
		if p++; p == pe {
			goto _test_eof806
		}
	st_case_806:
		if data[p] == 118 {
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
			goto tr482
		case 45:
			goto tr481
		case 46:
			goto tr1163
		case 47:
			goto tr482
		case 101:
			goto st808
		}
		goto st0
	st808:
		if p++; p == pe {
			goto _test_eof808
		}
	st_case_808:
		if data[p] == 109 {
			goto st809
		}
		goto st0
	st809:
		if p++; p == pe {
			goto _test_eof809
		}
	st_case_809:
		if data[p] == 98 {
			goto st810
		}
		goto st0
	st810:
		if p++; p == pe {
			goto _test_eof810
		}
	st_case_810:
		if data[p] == 101 {
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
		switch data[p] {
		case 32:
			goto tr482
		case 45:
			goto tr481
		case 46:
			goto tr1163
		case 47:
			goto tr482
		}
		goto st0
	st813:
		if p++; p == pe {
			goto _test_eof813
		}
	st_case_813:
		if data[p] == 99 {
			goto st814
		}
		goto st0
	st814:
		if p++; p == pe {
			goto _test_eof814
		}
	st_case_814:
		if data[p] == 116 {
			goto st815
		}
		goto st0
	st815:
		if p++; p == pe {
			goto _test_eof815
		}
	st_case_815:
		switch data[p] {
		case 32:
			goto tr492
		case 45:
			goto tr491
		case 46:
			goto tr1171
		case 47:
			goto tr492
		case 111:
			goto st816
		}
		goto st0
	st816:
		if p++; p == pe {
			goto _test_eof816
		}
	st_case_816:
		if data[p] == 98 {
			goto st817
		}
		goto st0
	st817:
		if p++; p == pe {
			goto _test_eof817
		}
	st_case_817:
		if data[p] == 101 {
			goto st818
		}
		goto st0
	st818:
		if p++; p == pe {
			goto _test_eof818
		}
	st_case_818:
		if data[p] == 114 {
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
			goto tr492
		case 45:
			goto tr491
		case 46:
			goto tr1171
		case 47:
			goto tr492
		}
		goto st0
	st820:
		if p++; p == pe {
			goto _test_eof820
		}
	st_case_820:
		if data[p] == 101 {
			goto st821
		}
		goto st0
	st821:
		if p++; p == pe {
			goto _test_eof821
		}
	st_case_821:
		if data[p] == 112 {
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
			goto tr501
		case 45:
			goto tr500
		case 46:
			goto tr1178
		case 47:
			goto tr501
		case 116:
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
			goto tr501
		case 45:
			goto tr500
		case 46:
			goto tr1178
		case 47:
			goto tr501
		case 101:
			goto st824
		}
		goto st0
	st824:
		if p++; p == pe {
			goto _test_eof824
		}
	st_case_824:
		if data[p] == 109 {
			goto st825
		}
		goto st0
	st825:
		if p++; p == pe {
			goto _test_eof825
		}
	st_case_825:
		if data[p] == 98 {
			goto st826
		}
		goto st0
	st826:
		if p++; p == pe {
			goto _test_eof826
		}
	st_case_826:
		if data[p] == 101 {
			goto st827
		}
		goto st0
	st827:
		if p++; p == pe {
			goto _test_eof827
		}
	st_case_827:
		if data[p] == 114 {
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
			goto tr501
		case 45:
			goto tr500
		case 46:
			goto tr1178
		case 47:
			goto tr501
		}
		goto st0
tr1102:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st829
	st829:
		if p++; p == pe {
			goto _test_eof829
		}
	st_case_829:
//line ragel_parse_datetime.go:28088
		if data[p] == 100 {
			goto st830
		}
		goto st0
	st830:
		if p++; p == pe {
			goto _test_eof830
		}
	st_case_830:
		switch data[p] {
		case 32:
			goto st429
		case 45:
			goto st759
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto st429
		}
		goto st0
tr1103:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st831
	st831:
		if p++; p == pe {
			goto _test_eof831
		}
	st_case_831:
//line ragel_parse_datetime.go:28124
		if data[p] == 116 {
			goto st830
		}
		goto st0
tr1104:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st832
	st832:
		if p++; p == pe {
			goto _test_eof832
		}
	st_case_832:
//line ragel_parse_datetime.go:28145
		if data[p] == 104 {
			goto st830
		}
		goto st0
tr1089:
//line ragel/datetime.rl:5
 pb = p 
	goto st833
	st833:
		if p++; p == pe {
			goto _test_eof833
		}
	st_case_833:
//line ragel_parse_datetime.go:28159
		switch data[p] {
		case 32:
			goto tr595
		case 45:
			goto tr1101
		case 110:
			goto tr1102
		case 114:
			goto tr1102
		case 115:
			goto tr1103
		case 116:
			goto tr1104
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st758
			}
		case data[p] >= 46:
			goto tr595
		}
		goto st0
tr1090:
//line ragel/datetime.rl:5
 pb = p 
	goto st834
	st834:
		if p++; p == pe {
			goto _test_eof834
		}
	st_case_834:
//line ragel_parse_datetime.go:28192
		switch data[p] {
		case 32:
			goto tr595
		case 45:
			goto tr1101
		case 110:
			goto tr1102
		case 114:
			goto tr1102
		case 115:
			goto tr1103
		case 116:
			goto tr1104
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st758
			}
		case data[p] >= 46:
			goto tr595
		}
		goto st0
	st835:
		if p++; p == pe {
			goto _test_eof835
		}
	st_case_835:
		switch data[p] {
		case 112:
			goto st836
		case 117:
			goto st850
		}
		goto st0
	st836:
		if p++; p == pe {
			goto _test_eof836
		}
	st_case_836:
		if data[p] == 114 {
			goto st837
		}
		goto st0
	st837:
		if p++; p == pe {
			goto _test_eof837
		}
	st_case_837:
		switch data[p] {
		case 32:
			goto tr1190
		case 46:
			goto tr1191
		case 105:
			goto st848
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1190
		}
		goto st0
tr1190:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st838
tr1205:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st838
tr1212:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st838
tr1221:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st838
tr1231:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st838
tr1239:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st838
tr1242:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st838
tr1248:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st838
tr1252:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st838
tr1256:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st838
tr1265:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st838
tr1273:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st838
	st838:
		if p++; p == pe {
			goto _test_eof838
		}
	st_case_838:
//line ragel_parse_datetime.go:28307
		switch data[p] {
		case 48:
			goto tr1193
		case 51:
			goto tr1195
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr1196
			}
		case data[p] >= 49:
			goto tr1194
		}
		goto st0
tr1193:
//line ragel/datetime.rl:5
 pb = p 
	goto st839
	st839:
		if p++; p == pe {
			goto _test_eof839
		}
	st_case_839:
//line ragel_parse_datetime.go:28332
		if 49 <= data[p] && data[p] <= 57 {
			goto st840
		}
		goto st0
tr1196:
//line ragel/datetime.rl:5
 pb = p 
	goto st840
	st840:
		if p++; p == pe {
			goto _test_eof840
		}
	st_case_840:
//line ragel_parse_datetime.go:28346
		switch data[p] {
		case 32:
			goto tr610
		case 110:
			goto tr1198
		case 114:
			goto tr1198
		case 115:
			goto tr1199
		case 116:
			goto tr1200
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr610
		}
		goto st0
tr1198:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st841
	st841:
		if p++; p == pe {
			goto _test_eof841
		}
	st_case_841:
//line ragel_parse_datetime.go:28379
		if data[p] == 100 {
			goto st842
		}
		goto st0
	st842:
		if p++; p == pe {
			goto _test_eof842
		}
	st_case_842:
		if data[p] == 32 {
			goto st198
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto st198
		}
		goto st0
tr1199:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st843
	st843:
		if p++; p == pe {
			goto _test_eof843
		}
	st_case_843:
//line ragel_parse_datetime.go:28412
		if data[p] == 116 {
			goto st842
		}
		goto st0
tr1200:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st844
	st844:
		if p++; p == pe {
			goto _test_eof844
		}
	st_case_844:
//line ragel_parse_datetime.go:28433
		if data[p] == 104 {
			goto st842
		}
		goto st0
tr1194:
//line ragel/datetime.rl:5
 pb = p 
	goto st845
	st845:
		if p++; p == pe {
			goto _test_eof845
		}
	st_case_845:
//line ragel_parse_datetime.go:28447
		switch data[p] {
		case 32:
			goto tr610
		case 110:
			goto tr1198
		case 114:
			goto tr1198
		case 115:
			goto tr1199
		case 116:
			goto tr1200
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto st840
			}
		case data[p] >= 45:
			goto tr610
		}
		goto st0
tr1195:
//line ragel/datetime.rl:5
 pb = p 
	goto st846
	st846:
		if p++; p == pe {
			goto _test_eof846
		}
	st_case_846:
//line ragel_parse_datetime.go:28478
		switch data[p] {
		case 32:
			goto tr610
		case 110:
			goto tr1198
		case 114:
			goto tr1198
		case 115:
			goto tr1199
		case 116:
			goto tr1200
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 49 {
				goto st840
			}
		case data[p] >= 45:
			goto tr610
		}
		goto st0
tr1191:
//line ragel/datetime.rl:97
 st.Month = 4 
	goto st847
tr1206:
//line ragel/datetime.rl:101
 st.Month = 8 
	goto st847
tr1213:
//line ragel/datetime.rl:105
 st.Month = 12 
	goto st847
tr1222:
//line ragel/datetime.rl:95
 st.Month = 2 
	goto st847
tr1232:
//line ragel/datetime.rl:94
 st.Month = 1 
	goto st847
tr1240:
//line ragel/datetime.rl:100
 st.Month = 7 
	goto st847
tr1243:
//line ragel/datetime.rl:99
 st.Month = 6 
	goto st847
tr1249:
//line ragel/datetime.rl:96
 st.Month = 3 
	goto st847
tr1253:
//line ragel/datetime.rl:98
 st.Month = 5 
	goto st847
tr1257:
//line ragel/datetime.rl:104
 st.Month = 11 
	goto st847
tr1266:
//line ragel/datetime.rl:103
 st.Month = 10 
	goto st847
tr1274:
//line ragel/datetime.rl:102
 st.Month = 9 
	goto st847
	st847:
		if p++; p == pe {
			goto _test_eof847
		}
	st_case_847:
//line ragel_parse_datetime.go:28553
		switch data[p] {
		case 32:
			goto st838
		case 48:
			goto tr1193
		case 51:
			goto tr1195
		}
		switch {
		case data[p] < 49:
			if 45 <= data[p] && data[p] <= 47 {
				goto st838
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr1196
			}
		default:
			goto tr1194
		}
		goto st0
	st848:
		if p++; p == pe {
			goto _test_eof848
		}
	st_case_848:
		if data[p] == 108 {
			goto st849
		}
		goto st0
	st849:
		if p++; p == pe {
			goto _test_eof849
		}
	st_case_849:
		switch data[p] {
		case 32:
			goto tr1190
		case 46:
			goto tr1191
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1190
		}
		goto st0
	st850:
		if p++; p == pe {
			goto _test_eof850
		}
	st_case_850:
		if data[p] == 103 {
			goto st851
		}
		goto st0
	st851:
		if p++; p == pe {
			goto _test_eof851
		}
	st_case_851:
		switch data[p] {
		case 32:
			goto tr1205
		case 46:
			goto tr1206
		case 117:
			goto st852
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1205
		}
		goto st0
	st852:
		if p++; p == pe {
			goto _test_eof852
		}
	st_case_852:
		if data[p] == 115 {
			goto st853
		}
		goto st0
	st853:
		if p++; p == pe {
			goto _test_eof853
		}
	st_case_853:
		if data[p] == 116 {
			goto st854
		}
		goto st0
	st854:
		if p++; p == pe {
			goto _test_eof854
		}
	st_case_854:
		switch data[p] {
		case 32:
			goto tr1205
		case 46:
			goto tr1206
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1205
		}
		goto st0
	st855:
		if p++; p == pe {
			goto _test_eof855
		}
	st_case_855:
		if data[p] == 101 {
			goto st856
		}
		goto st0
	st856:
		if p++; p == pe {
			goto _test_eof856
		}
	st_case_856:
		if data[p] == 99 {
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
			goto tr1212
		case 46:
			goto tr1213
		case 101:
			goto st858
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1212
		}
		goto st0
	st858:
		if p++; p == pe {
			goto _test_eof858
		}
	st_case_858:
		if data[p] == 109 {
			goto st859
		}
		goto st0
	st859:
		if p++; p == pe {
			goto _test_eof859
		}
	st_case_859:
		if data[p] == 98 {
			goto st860
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
		if data[p] == 114 {
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
			goto tr1212
		case 46:
			goto tr1213
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1212
		}
		goto st0
	st863:
		if p++; p == pe {
			goto _test_eof863
		}
	st_case_863:
		if data[p] == 101 {
			goto st864
		}
		goto st0
	st864:
		if p++; p == pe {
			goto _test_eof864
		}
	st_case_864:
		if data[p] == 98 {
			goto st865
		}
		goto st0
	st865:
		if p++; p == pe {
			goto _test_eof865
		}
	st_case_865:
		switch data[p] {
		case 32:
			goto tr1221
		case 46:
			goto tr1222
		case 114:
			goto st866
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1221
		}
		goto st0
	st866:
		if p++; p == pe {
			goto _test_eof866
		}
	st_case_866:
		if data[p] == 117 {
			goto st867
		}
		goto st0
	st867:
		if p++; p == pe {
			goto _test_eof867
		}
	st_case_867:
		if data[p] == 97 {
			goto st868
		}
		goto st0
	st868:
		if p++; p == pe {
			goto _test_eof868
		}
	st_case_868:
		if data[p] == 114 {
			goto st869
		}
		goto st0
	st869:
		if p++; p == pe {
			goto _test_eof869
		}
	st_case_869:
		if data[p] == 121 {
			goto st870
		}
		goto st0
	st870:
		if p++; p == pe {
			goto _test_eof870
		}
	st_case_870:
		switch data[p] {
		case 32:
			goto tr1221
		case 46:
			goto tr1222
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1221
		}
		goto st0
	st871:
		if p++; p == pe {
			goto _test_eof871
		}
	st_case_871:
		switch data[p] {
		case 97:
			goto st872
		case 117:
			goto st878
		}
		goto st0
	st872:
		if p++; p == pe {
			goto _test_eof872
		}
	st_case_872:
		if data[p] == 110 {
			goto st873
		}
		goto st0
	st873:
		if p++; p == pe {
			goto _test_eof873
		}
	st_case_873:
		switch data[p] {
		case 32:
			goto tr1231
		case 46:
			goto tr1232
		case 117:
			goto st874
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1231
		}
		goto st0
	st874:
		if p++; p == pe {
			goto _test_eof874
		}
	st_case_874:
		if data[p] == 97 {
			goto st875
		}
		goto st0
	st875:
		if p++; p == pe {
			goto _test_eof875
		}
	st_case_875:
		if data[p] == 114 {
			goto st876
		}
		goto st0
	st876:
		if p++; p == pe {
			goto _test_eof876
		}
	st_case_876:
		if data[p] == 121 {
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
			goto tr1231
		case 46:
			goto tr1232
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1231
		}
		goto st0
	st878:
		if p++; p == pe {
			goto _test_eof878
		}
	st_case_878:
		switch data[p] {
		case 108:
			goto st879
		case 110:
			goto st881
		}
		goto st0
	st879:
		if p++; p == pe {
			goto _test_eof879
		}
	st_case_879:
		switch data[p] {
		case 32:
			goto tr1239
		case 46:
			goto tr1240
		case 121:
			goto st880
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1239
		}
		goto st0
	st880:
		if p++; p == pe {
			goto _test_eof880
		}
	st_case_880:
		switch data[p] {
		case 32:
			goto tr1239
		case 46:
			goto tr1240
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1239
		}
		goto st0
	st881:
		if p++; p == pe {
			goto _test_eof881
		}
	st_case_881:
		switch data[p] {
		case 32:
			goto tr1242
		case 46:
			goto tr1243
		case 101:
			goto st882
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1242
		}
		goto st0
	st882:
		if p++; p == pe {
			goto _test_eof882
		}
	st_case_882:
		switch data[p] {
		case 32:
			goto tr1242
		case 46:
			goto tr1243
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1242
		}
		goto st0
	st883:
		if p++; p == pe {
			goto _test_eof883
		}
	st_case_883:
		if data[p] == 97 {
			goto st884
		}
		goto st0
	st884:
		if p++; p == pe {
			goto _test_eof884
		}
	st_case_884:
		switch data[p] {
		case 114:
			goto st885
		case 121:
			goto st888
		}
		goto st0
	st885:
		if p++; p == pe {
			goto _test_eof885
		}
	st_case_885:
		switch data[p] {
		case 32:
			goto tr1248
		case 46:
			goto tr1249
		case 99:
			goto st886
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1248
		}
		goto st0
	st886:
		if p++; p == pe {
			goto _test_eof886
		}
	st_case_886:
		if data[p] == 104 {
			goto st887
		}
		goto st0
	st887:
		if p++; p == pe {
			goto _test_eof887
		}
	st_case_887:
		switch data[p] {
		case 32:
			goto tr1248
		case 46:
			goto tr1249
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1248
		}
		goto st0
	st888:
		if p++; p == pe {
			goto _test_eof888
		}
	st_case_888:
		switch data[p] {
		case 32:
			goto tr1252
		case 46:
			goto tr1253
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1252
		}
		goto st0
	st889:
		if p++; p == pe {
			goto _test_eof889
		}
	st_case_889:
		if data[p] == 111 {
			goto st890
		}
		goto st0
	st890:
		if p++; p == pe {
			goto _test_eof890
		}
	st_case_890:
		if data[p] == 118 {
			goto st891
		}
		goto st0
	st891:
		if p++; p == pe {
			goto _test_eof891
		}
	st_case_891:
		switch data[p] {
		case 32:
			goto tr1256
		case 46:
			goto tr1257
		case 101:
			goto st892
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1256
		}
		goto st0
	st892:
		if p++; p == pe {
			goto _test_eof892
		}
	st_case_892:
		if data[p] == 109 {
			goto st893
		}
		goto st0
	st893:
		if p++; p == pe {
			goto _test_eof893
		}
	st_case_893:
		if data[p] == 98 {
			goto st894
		}
		goto st0
	st894:
		if p++; p == pe {
			goto _test_eof894
		}
	st_case_894:
		if data[p] == 101 {
			goto st895
		}
		goto st0
	st895:
		if p++; p == pe {
			goto _test_eof895
		}
	st_case_895:
		if data[p] == 114 {
			goto st896
		}
		goto st0
	st896:
		if p++; p == pe {
			goto _test_eof896
		}
	st_case_896:
		switch data[p] {
		case 32:
			goto tr1256
		case 46:
			goto tr1257
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1256
		}
		goto st0
	st897:
		if p++; p == pe {
			goto _test_eof897
		}
	st_case_897:
		if data[p] == 99 {
			goto st898
		}
		goto st0
	st898:
		if p++; p == pe {
			goto _test_eof898
		}
	st_case_898:
		if data[p] == 116 {
			goto st899
		}
		goto st0
	st899:
		if p++; p == pe {
			goto _test_eof899
		}
	st_case_899:
		switch data[p] {
		case 32:
			goto tr1265
		case 46:
			goto tr1266
		case 111:
			goto st900
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1265
		}
		goto st0
	st900:
		if p++; p == pe {
			goto _test_eof900
		}
	st_case_900:
		if data[p] == 98 {
			goto st901
		}
		goto st0
	st901:
		if p++; p == pe {
			goto _test_eof901
		}
	st_case_901:
		if data[p] == 101 {
			goto st902
		}
		goto st0
	st902:
		if p++; p == pe {
			goto _test_eof902
		}
	st_case_902:
		if data[p] == 114 {
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
			goto tr1265
		case 46:
			goto tr1266
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1265
		}
		goto st0
	st904:
		if p++; p == pe {
			goto _test_eof904
		}
	st_case_904:
		if data[p] == 101 {
			goto st905
		}
		goto st0
	st905:
		if p++; p == pe {
			goto _test_eof905
		}
	st_case_905:
		if data[p] == 112 {
			goto st906
		}
		goto st0
	st906:
		if p++; p == pe {
			goto _test_eof906
		}
	st_case_906:
		switch data[p] {
		case 32:
			goto tr1273
		case 46:
			goto tr1274
		case 116:
			goto st907
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1273
		}
		goto st0
	st907:
		if p++; p == pe {
			goto _test_eof907
		}
	st_case_907:
		switch data[p] {
		case 32:
			goto tr1273
		case 46:
			goto tr1274
		case 101:
			goto st908
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1273
		}
		goto st0
	st908:
		if p++; p == pe {
			goto _test_eof908
		}
	st_case_908:
		if data[p] == 109 {
			goto st909
		}
		goto st0
	st909:
		if p++; p == pe {
			goto _test_eof909
		}
	st_case_909:
		if data[p] == 98 {
			goto st910
		}
		goto st0
	st910:
		if p++; p == pe {
			goto _test_eof910
		}
	st_case_910:
		if data[p] == 101 {
			goto st911
		}
		goto st0
	st911:
		if p++; p == pe {
			goto _test_eof911
		}
	st_case_911:
		if data[p] == 114 {
			goto st912
		}
		goto st0
	st912:
		if p++; p == pe {
			goto _test_eof912
		}
	st_case_912:
		switch data[p] {
		case 32:
			goto tr1273
		case 46:
			goto tr1274
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1273
		}
		goto st0
	st913:
		if p++; p == pe {
			goto _test_eof913
		}
	st_case_913:
		if data[p] == 97 {
			goto st914
		}
		goto st0
	st914:
		if p++; p == pe {
			goto _test_eof914
		}
	st_case_914:
		if data[p] == 121 {
			goto st915
		}
		goto st0
	st915:
		if p++; p == pe {
			goto _test_eof915
		}
	st_case_915:
		switch data[p] {
		case 32:
			goto st514
		case 44:
			goto st755
		}
		goto st0
	st916:
		if p++; p == pe {
			goto _test_eof916
		}
	st_case_916:
		switch data[p] {
		case 97:
			goto st917
		case 117:
			goto st923
		}
		goto st0
	st917:
		if p++; p == pe {
			goto _test_eof917
		}
	st_case_917:
		if data[p] == 110 {
			goto st918
		}
		goto st0
	st918:
		if p++; p == pe {
			goto _test_eof918
		}
	st_case_918:
		switch data[p] {
		case 32:
			goto tr1286
		case 46:
			goto tr1287
		case 117:
			goto st919
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1030
		}
		goto st0
	st919:
		if p++; p == pe {
			goto _test_eof919
		}
	st_case_919:
		if data[p] == 97 {
			goto st920
		}
		goto st0
	st920:
		if p++; p == pe {
			goto _test_eof920
		}
	st_case_920:
		if data[p] == 114 {
			goto st921
		}
		goto st0
	st921:
		if p++; p == pe {
			goto _test_eof921
		}
	st_case_921:
		if data[p] == 121 {
			goto st922
		}
		goto st0
	st922:
		if p++; p == pe {
			goto _test_eof922
		}
	st_case_922:
		switch data[p] {
		case 32:
			goto tr1286
		case 46:
			goto tr1287
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1030
		}
		goto st0
	st923:
		if p++; p == pe {
			goto _test_eof923
		}
	st_case_923:
		switch data[p] {
		case 108:
			goto st924
		case 110:
			goto st926
		}
		goto st0
	st924:
		if p++; p == pe {
			goto _test_eof924
		}
	st_case_924:
		switch data[p] {
		case 32:
			goto tr1294
		case 46:
			goto tr1295
		case 121:
			goto st925
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1039
		}
		goto st0
	st925:
		if p++; p == pe {
			goto _test_eof925
		}
	st_case_925:
		switch data[p] {
		case 32:
			goto tr1294
		case 46:
			goto tr1295
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1039
		}
		goto st0
	st926:
		if p++; p == pe {
			goto _test_eof926
		}
	st_case_926:
		switch data[p] {
		case 32:
			goto tr1297
		case 46:
			goto tr1298
		case 101:
			goto st927
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1043
		}
		goto st0
	st927:
		if p++; p == pe {
			goto _test_eof927
		}
	st_case_927:
		switch data[p] {
		case 32:
			goto tr1297
		case 46:
			goto tr1298
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1043
		}
		goto st0
	st928:
		if p++; p == pe {
			goto _test_eof928
		}
	st_case_928:
		switch data[p] {
		case 97:
			goto st929
		case 111:
			goto st934
		}
		goto st0
	st929:
		if p++; p == pe {
			goto _test_eof929
		}
	st_case_929:
		switch data[p] {
		case 114:
			goto st930
		case 121:
			goto st933
		}
		goto st0
	st930:
		if p++; p == pe {
			goto _test_eof930
		}
	st_case_930:
		switch data[p] {
		case 32:
			goto tr1304
		case 46:
			goto tr1305
		case 99:
			goto st931
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1050
		}
		goto st0
	st931:
		if p++; p == pe {
			goto _test_eof931
		}
	st_case_931:
		if data[p] == 104 {
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
			goto tr1304
		case 46:
			goto tr1305
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1050
		}
		goto st0
	st933:
		if p++; p == pe {
			goto _test_eof933
		}
	st_case_933:
		switch data[p] {
		case 32:
			goto tr1308
		case 46:
			goto tr1309
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1055
		}
		goto st0
	st934:
		if p++; p == pe {
			goto _test_eof934
		}
	st_case_934:
		if data[p] == 110 {
			goto st513
		}
		goto st0
	st935:
		if p++; p == pe {
			goto _test_eof935
		}
	st_case_935:
		if data[p] == 111 {
			goto st936
		}
		goto st0
	st936:
		if p++; p == pe {
			goto _test_eof936
		}
	st_case_936:
		if data[p] == 118 {
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
			goto tr1312
		case 46:
			goto tr1313
		case 101:
			goto st938
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1060
		}
		goto st0
	st938:
		if p++; p == pe {
			goto _test_eof938
		}
	st_case_938:
		if data[p] == 109 {
			goto st939
		}
		goto st0
	st939:
		if p++; p == pe {
			goto _test_eof939
		}
	st_case_939:
		if data[p] == 98 {
			goto st940
		}
		goto st0
	st940:
		if p++; p == pe {
			goto _test_eof940
		}
	st_case_940:
		if data[p] == 101 {
			goto st941
		}
		goto st0
	st941:
		if p++; p == pe {
			goto _test_eof941
		}
	st_case_941:
		if data[p] == 114 {
			goto st942
		}
		goto st0
	st942:
		if p++; p == pe {
			goto _test_eof942
		}
	st_case_942:
		switch data[p] {
		case 32:
			goto tr1312
		case 46:
			goto tr1313
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1060
		}
		goto st0
	st943:
		if p++; p == pe {
			goto _test_eof943
		}
	st_case_943:
		if data[p] == 99 {
			goto st944
		}
		goto st0
	st944:
		if p++; p == pe {
			goto _test_eof944
		}
	st_case_944:
		if data[p] == 116 {
			goto st945
		}
		goto st0
	st945:
		if p++; p == pe {
			goto _test_eof945
		}
	st_case_945:
		switch data[p] {
		case 32:
			goto tr1321
		case 46:
			goto tr1322
		case 111:
			goto st946
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1070
		}
		goto st0
	st946:
		if p++; p == pe {
			goto _test_eof946
		}
	st_case_946:
		if data[p] == 98 {
			goto st947
		}
		goto st0
	st947:
		if p++; p == pe {
			goto _test_eof947
		}
	st_case_947:
		if data[p] == 101 {
			goto st948
		}
		goto st0
	st948:
		if p++; p == pe {
			goto _test_eof948
		}
	st_case_948:
		if data[p] == 114 {
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
			goto tr1321
		case 46:
			goto tr1322
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1070
		}
		goto st0
	st950:
		if p++; p == pe {
			goto _test_eof950
		}
	st_case_950:
		switch data[p] {
		case 97:
			goto st951
		case 101:
			goto st955
		case 117:
			goto st934
		}
		goto st0
	st951:
		if p++; p == pe {
			goto _test_eof951
		}
	st_case_951:
		if data[p] == 116 {
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
			goto st514
		case 44:
			goto st755
		case 117:
			goto st953
		}
		goto st0
	st953:
		if p++; p == pe {
			goto _test_eof953
		}
	st_case_953:
		if data[p] == 114 {
			goto st954
		}
		goto st0
	st954:
		if p++; p == pe {
			goto _test_eof954
		}
	st_case_954:
		if data[p] == 100 {
			goto st913
		}
		goto st0
	st955:
		if p++; p == pe {
			goto _test_eof955
		}
	st_case_955:
		if data[p] == 112 {
			goto st956
		}
		goto st0
	st956:
		if p++; p == pe {
			goto _test_eof956
		}
	st_case_956:
		switch data[p] {
		case 32:
			goto tr1333
		case 46:
			goto tr1334
		case 116:
			goto st957
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1079
		}
		goto st0
	st957:
		if p++; p == pe {
			goto _test_eof957
		}
	st_case_957:
		switch data[p] {
		case 32:
			goto tr1333
		case 46:
			goto tr1334
		case 101:
			goto st958
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1079
		}
		goto st0
	st958:
		if p++; p == pe {
			goto _test_eof958
		}
	st_case_958:
		if data[p] == 109 {
			goto st959
		}
		goto st0
	st959:
		if p++; p == pe {
			goto _test_eof959
		}
	st_case_959:
		if data[p] == 98 {
			goto st960
		}
		goto st0
	st960:
		if p++; p == pe {
			goto _test_eof960
		}
	st_case_960:
		if data[p] == 101 {
			goto st961
		}
		goto st0
	st961:
		if p++; p == pe {
			goto _test_eof961
		}
	st_case_961:
		if data[p] == 114 {
			goto st962
		}
		goto st0
	st962:
		if p++; p == pe {
			goto _test_eof962
		}
	st_case_962:
		switch data[p] {
		case 32:
			goto tr1333
		case 46:
			goto tr1334
		}
		if 45 <= data[p] && data[p] <= 47 {
			goto tr1079
		}
		goto st0
	st963:
		if p++; p == pe {
			goto _test_eof963
		}
	st_case_963:
		switch data[p] {
		case 104:
			goto st964
		case 117:
			goto st967
		}
		goto st0
	st964:
		if p++; p == pe {
			goto _test_eof964
		}
	st_case_964:
		if data[p] == 117 {
			goto st965
		}
		goto st0
	st965:
		if p++; p == pe {
			goto _test_eof965
		}
	st_case_965:
		switch data[p] {
		case 32:
			goto st514
		case 44:
			goto st755
		case 114:
			goto st966
		}
		goto st0
	st966:
		if p++; p == pe {
			goto _test_eof966
		}
	st_case_966:
		if data[p] == 115 {
			goto st954
		}
		goto st0
	st967:
		if p++; p == pe {
			goto _test_eof967
		}
	st_case_967:
		if data[p] == 101 {
			goto st968
		}
		goto st0
	st968:
		if p++; p == pe {
			goto _test_eof968
		}
	st_case_968:
		switch data[p] {
		case 32:
			goto st514
		case 44:
			goto st755
		case 115:
			goto st954
		}
		goto st0
	st969:
		if p++; p == pe {
			goto _test_eof969
		}
	st_case_969:
		if data[p] == 101 {
			goto st970
		}
		goto st0
	st970:
		if p++; p == pe {
			goto _test_eof970
		}
	st_case_970:
		if data[p] == 100 {
			goto st971
		}
		goto st0
	st971:
		if p++; p == pe {
			goto _test_eof971
		}
	st_case_971:
		switch data[p] {
		case 32:
			goto st514
		case 44:
			goto st755
		case 110:
			goto st972
		}
		goto st0
	st972:
		if p++; p == pe {
			goto _test_eof972
		}
	st_case_972:
		if data[p] == 101 {
			goto st966
		}
		goto st0
	st973:
		if p++; p == pe {
			goto _test_eof973
		}
	st_case_973:
		if data[p] == 101 {
			goto st505
		}
		goto st0
	st974:
		if p++; p == pe {
			goto _test_eof974
		}
	st_case_974:
		if data[p] == 97 {
			goto st929
		}
		goto st0
	st975:
		if p++; p == pe {
			goto _test_eof975
		}
	st_case_975:
		if data[p] == 101 {
			goto st955
		}
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof976: cs = 976; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof977: cs = 977; goto _test_eof
	_test_eof978: cs = 978; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof979: cs = 979; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof980: cs = 980; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof981: cs = 981; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof982: cs = 982; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof983: cs = 983; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof984: cs = 984; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof985: cs = 985; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof986: cs = 986; goto _test_eof
	_test_eof987: cs = 987; goto _test_eof
	_test_eof988: cs = 988; goto _test_eof
	_test_eof989: cs = 989; goto _test_eof
	_test_eof990: cs = 990; goto _test_eof
	_test_eof991: cs = 991; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof992: cs = 992; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof993: cs = 993; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof994: cs = 994; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof995: cs = 995; goto _test_eof
	_test_eof996: cs = 996; goto _test_eof
	_test_eof997: cs = 997; goto _test_eof
	_test_eof998: cs = 998; goto _test_eof
	_test_eof999: cs = 999; goto _test_eof
	_test_eof1000: cs = 1000; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof1001: cs = 1001; goto _test_eof
	_test_eof1002: cs = 1002; goto _test_eof
	_test_eof1003: cs = 1003; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof1004: cs = 1004; goto _test_eof
	_test_eof1005: cs = 1005; goto _test_eof
	_test_eof1006: cs = 1006; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof1007: cs = 1007; goto _test_eof
	_test_eof1008: cs = 1008; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof1009: cs = 1009; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof1010: cs = 1010; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof1011: cs = 1011; goto _test_eof
	_test_eof1012: cs = 1012; goto _test_eof
	_test_eof1013: cs = 1013; goto _test_eof
	_test_eof1014: cs = 1014; goto _test_eof
	_test_eof1015: cs = 1015; goto _test_eof
	_test_eof1016: cs = 1016; goto _test_eof
	_test_eof1017: cs = 1017; goto _test_eof
	_test_eof1018: cs = 1018; goto _test_eof
	_test_eof1019: cs = 1019; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof1020: cs = 1020; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof1021: cs = 1021; goto _test_eof
	_test_eof1022: cs = 1022; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof1023: cs = 1023; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof1024: cs = 1024; goto _test_eof
	_test_eof1025: cs = 1025; goto _test_eof
	_test_eof1026: cs = 1026; goto _test_eof
	_test_eof1027: cs = 1027; goto _test_eof
	_test_eof1028: cs = 1028; goto _test_eof
	_test_eof1029: cs = 1029; goto _test_eof
	_test_eof1030: cs = 1030; goto _test_eof
	_test_eof1031: cs = 1031; goto _test_eof
	_test_eof1032: cs = 1032; goto _test_eof
	_test_eof1033: cs = 1033; goto _test_eof
	_test_eof1034: cs = 1034; goto _test_eof
	_test_eof1035: cs = 1035; goto _test_eof
	_test_eof1036: cs = 1036; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof1037: cs = 1037; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof1038: cs = 1038; goto _test_eof
	_test_eof1039: cs = 1039; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof1040: cs = 1040; goto _test_eof
	_test_eof1041: cs = 1041; goto _test_eof
	_test_eof1042: cs = 1042; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof1043: cs = 1043; goto _test_eof
	_test_eof1044: cs = 1044; goto _test_eof
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
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof1045: cs = 1045; goto _test_eof
	_test_eof1046: cs = 1046; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof1047: cs = 1047; goto _test_eof
	_test_eof1048: cs = 1048; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof1049: cs = 1049; goto _test_eof
	_test_eof1050: cs = 1050; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof1051: cs = 1051; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof1052: cs = 1052; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof1053: cs = 1053; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof1054: cs = 1054; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof1055: cs = 1055; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof1056: cs = 1056; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof1057: cs = 1057; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof1058: cs = 1058; goto _test_eof
	_test_eof150: cs = 150; goto _test_eof
	_test_eof151: cs = 151; goto _test_eof
	_test_eof152: cs = 152; goto _test_eof
	_test_eof1059: cs = 1059; goto _test_eof
	_test_eof153: cs = 153; goto _test_eof
	_test_eof1060: cs = 1060; goto _test_eof
	_test_eof1061: cs = 1061; goto _test_eof
	_test_eof1062: cs = 1062; goto _test_eof
	_test_eof1063: cs = 1063; goto _test_eof
	_test_eof154: cs = 154; goto _test_eof
	_test_eof155: cs = 155; goto _test_eof
	_test_eof1064: cs = 1064; goto _test_eof
	_test_eof156: cs = 156; goto _test_eof
	_test_eof1065: cs = 1065; goto _test_eof
	_test_eof1066: cs = 1066; goto _test_eof
	_test_eof157: cs = 157; goto _test_eof
	_test_eof158: cs = 158; goto _test_eof
	_test_eof1067: cs = 1067; goto _test_eof
	_test_eof159: cs = 159; goto _test_eof
	_test_eof160: cs = 160; goto _test_eof
	_test_eof161: cs = 161; goto _test_eof
	_test_eof162: cs = 162; goto _test_eof
	_test_eof1068: cs = 1068; goto _test_eof
	_test_eof163: cs = 163; goto _test_eof
	_test_eof164: cs = 164; goto _test_eof
	_test_eof1069: cs = 1069; goto _test_eof
	_test_eof165: cs = 165; goto _test_eof
	_test_eof166: cs = 166; goto _test_eof
	_test_eof167: cs = 167; goto _test_eof
	_test_eof1070: cs = 1070; goto _test_eof
	_test_eof168: cs = 168; goto _test_eof
	_test_eof169: cs = 169; goto _test_eof
	_test_eof1071: cs = 1071; goto _test_eof
	_test_eof1072: cs = 1072; goto _test_eof
	_test_eof170: cs = 170; goto _test_eof
	_test_eof171: cs = 171; goto _test_eof
	_test_eof172: cs = 172; goto _test_eof
	_test_eof173: cs = 173; goto _test_eof
	_test_eof1073: cs = 1073; goto _test_eof
	_test_eof174: cs = 174; goto _test_eof
	_test_eof175: cs = 175; goto _test_eof
	_test_eof1074: cs = 1074; goto _test_eof
	_test_eof1075: cs = 1075; goto _test_eof
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
	_test_eof1076: cs = 1076; goto _test_eof
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
	_test_eof1077: cs = 1077; goto _test_eof
	_test_eof202: cs = 202; goto _test_eof
	_test_eof203: cs = 203; goto _test_eof
	_test_eof204: cs = 204; goto _test_eof
	_test_eof205: cs = 205; goto _test_eof
	_test_eof206: cs = 206; goto _test_eof
	_test_eof207: cs = 207; goto _test_eof
	_test_eof208: cs = 208; goto _test_eof
	_test_eof1078: cs = 1078; goto _test_eof
	_test_eof209: cs = 209; goto _test_eof
	_test_eof210: cs = 210; goto _test_eof
	_test_eof1079: cs = 1079; goto _test_eof
	_test_eof211: cs = 211; goto _test_eof
	_test_eof212: cs = 212; goto _test_eof
	_test_eof213: cs = 213; goto _test_eof
	_test_eof1080: cs = 1080; goto _test_eof
	_test_eof214: cs = 214; goto _test_eof
	_test_eof1081: cs = 1081; goto _test_eof
	_test_eof215: cs = 215; goto _test_eof
	_test_eof1082: cs = 1082; goto _test_eof
	_test_eof1083: cs = 1083; goto _test_eof
	_test_eof1084: cs = 1084; goto _test_eof
	_test_eof1085: cs = 1085; goto _test_eof
	_test_eof1086: cs = 1086; goto _test_eof
	_test_eof1087: cs = 1087; goto _test_eof
	_test_eof216: cs = 216; goto _test_eof
	_test_eof217: cs = 217; goto _test_eof
	_test_eof218: cs = 218; goto _test_eof
	_test_eof1088: cs = 1088; goto _test_eof
	_test_eof219: cs = 219; goto _test_eof
	_test_eof1089: cs = 1089; goto _test_eof
	_test_eof1090: cs = 1090; goto _test_eof
	_test_eof1091: cs = 1091; goto _test_eof
	_test_eof1092: cs = 1092; goto _test_eof
	_test_eof1093: cs = 1093; goto _test_eof
	_test_eof1094: cs = 1094; goto _test_eof
	_test_eof220: cs = 220; goto _test_eof
	_test_eof221: cs = 221; goto _test_eof
	_test_eof222: cs = 222; goto _test_eof
	_test_eof1095: cs = 1095; goto _test_eof
	_test_eof1096: cs = 1096; goto _test_eof
	_test_eof223: cs = 223; goto _test_eof
	_test_eof224: cs = 224; goto _test_eof
	_test_eof1097: cs = 1097; goto _test_eof
	_test_eof225: cs = 225; goto _test_eof
	_test_eof1098: cs = 1098; goto _test_eof
	_test_eof226: cs = 226; goto _test_eof
	_test_eof227: cs = 227; goto _test_eof
	_test_eof228: cs = 228; goto _test_eof
	_test_eof229: cs = 229; goto _test_eof
	_test_eof230: cs = 230; goto _test_eof
	_test_eof231: cs = 231; goto _test_eof
	_test_eof1099: cs = 1099; goto _test_eof
	_test_eof1100: cs = 1100; goto _test_eof
	_test_eof1101: cs = 1101; goto _test_eof
	_test_eof1102: cs = 1102; goto _test_eof
	_test_eof232: cs = 232; goto _test_eof
	_test_eof1103: cs = 1103; goto _test_eof
	_test_eof1104: cs = 1104; goto _test_eof
	_test_eof1105: cs = 1105; goto _test_eof
	_test_eof1106: cs = 1106; goto _test_eof
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
	_test_eof1107: cs = 1107; goto _test_eof
	_test_eof288: cs = 288; goto _test_eof
	_test_eof1108: cs = 1108; goto _test_eof
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
	_test_eof1109: cs = 1109; goto _test_eof
	_test_eof446: cs = 446; goto _test_eof
	_test_eof1110: cs = 1110; goto _test_eof
	_test_eof1111: cs = 1111; goto _test_eof
	_test_eof447: cs = 447; goto _test_eof
	_test_eof1112: cs = 1112; goto _test_eof
	_test_eof1113: cs = 1113; goto _test_eof
	_test_eof448: cs = 448; goto _test_eof
	_test_eof1114: cs = 1114; goto _test_eof
	_test_eof449: cs = 449; goto _test_eof
	_test_eof1115: cs = 1115; goto _test_eof
	_test_eof450: cs = 450; goto _test_eof
	_test_eof1116: cs = 1116; goto _test_eof
	_test_eof1117: cs = 1117; goto _test_eof
	_test_eof1118: cs = 1118; goto _test_eof
	_test_eof1119: cs = 1119; goto _test_eof
	_test_eof1120: cs = 1120; goto _test_eof
	_test_eof1121: cs = 1121; goto _test_eof
	_test_eof1122: cs = 1122; goto _test_eof
	_test_eof1123: cs = 1123; goto _test_eof
	_test_eof1124: cs = 1124; goto _test_eof
	_test_eof451: cs = 451; goto _test_eof
	_test_eof1125: cs = 1125; goto _test_eof
	_test_eof452: cs = 452; goto _test_eof
	_test_eof1126: cs = 1126; goto _test_eof
	_test_eof1127: cs = 1127; goto _test_eof
	_test_eof453: cs = 453; goto _test_eof
	_test_eof1128: cs = 1128; goto _test_eof
	_test_eof454: cs = 454; goto _test_eof
	_test_eof1129: cs = 1129; goto _test_eof
	_test_eof1130: cs = 1130; goto _test_eof
	_test_eof1131: cs = 1131; goto _test_eof
	_test_eof1132: cs = 1132; goto _test_eof
	_test_eof1133: cs = 1133; goto _test_eof
	_test_eof1134: cs = 1134; goto _test_eof
	_test_eof1135: cs = 1135; goto _test_eof
	_test_eof1136: cs = 1136; goto _test_eof
	_test_eof1137: cs = 1137; goto _test_eof
	_test_eof1138: cs = 1138; goto _test_eof
	_test_eof1139: cs = 1139; goto _test_eof
	_test_eof1140: cs = 1140; goto _test_eof
	_test_eof1141: cs = 1141; goto _test_eof
	_test_eof455: cs = 455; goto _test_eof
	_test_eof1142: cs = 1142; goto _test_eof
	_test_eof456: cs = 456; goto _test_eof
	_test_eof457: cs = 457; goto _test_eof
	_test_eof458: cs = 458; goto _test_eof
	_test_eof459: cs = 459; goto _test_eof
	_test_eof1143: cs = 1143; goto _test_eof
	_test_eof460: cs = 460; goto _test_eof
	_test_eof1144: cs = 1144; goto _test_eof
	_test_eof461: cs = 461; goto _test_eof
	_test_eof1145: cs = 1145; goto _test_eof
	_test_eof1146: cs = 1146; goto _test_eof
	_test_eof462: cs = 462; goto _test_eof
	_test_eof1147: cs = 1147; goto _test_eof
	_test_eof1148: cs = 1148; goto _test_eof
	_test_eof1149: cs = 1149; goto _test_eof
	_test_eof1150: cs = 1150; goto _test_eof
	_test_eof463: cs = 463; goto _test_eof
	_test_eof464: cs = 464; goto _test_eof
	_test_eof465: cs = 465; goto _test_eof
	_test_eof1151: cs = 1151; goto _test_eof
	_test_eof1152: cs = 1152; goto _test_eof
	_test_eof466: cs = 466; goto _test_eof
	_test_eof467: cs = 467; goto _test_eof
	_test_eof1153: cs = 1153; goto _test_eof
	_test_eof468: cs = 468; goto _test_eof
	_test_eof469: cs = 469; goto _test_eof
	_test_eof470: cs = 470; goto _test_eof
	_test_eof471: cs = 471; goto _test_eof
	_test_eof1154: cs = 1154; goto _test_eof
	_test_eof472: cs = 472; goto _test_eof
	_test_eof1155: cs = 1155; goto _test_eof
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
	_test_eof1156: cs = 1156; goto _test_eof
	_test_eof538: cs = 538; goto _test_eof
	_test_eof1157: cs = 1157; goto _test_eof
	_test_eof539: cs = 539; goto _test_eof
	_test_eof1158: cs = 1158; goto _test_eof
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
	_test_eof1159: cs = 1159; goto _test_eof
	_test_eof1160: cs = 1160; goto _test_eof
	_test_eof627: cs = 627; goto _test_eof
	_test_eof628: cs = 628; goto _test_eof
	_test_eof1161: cs = 1161; goto _test_eof
	_test_eof629: cs = 629; goto _test_eof
	_test_eof630: cs = 630; goto _test_eof
	_test_eof631: cs = 631; goto _test_eof
	_test_eof1162: cs = 1162; goto _test_eof
	_test_eof632: cs = 632; goto _test_eof
	_test_eof1163: cs = 1163; goto _test_eof
	_test_eof633: cs = 633; goto _test_eof
	_test_eof1164: cs = 1164; goto _test_eof
	_test_eof1165: cs = 1165; goto _test_eof
	_test_eof1166: cs = 1166; goto _test_eof
	_test_eof1167: cs = 1167; goto _test_eof
	_test_eof1168: cs = 1168; goto _test_eof
	_test_eof1169: cs = 1169; goto _test_eof
	_test_eof634: cs = 634; goto _test_eof
	_test_eof635: cs = 635; goto _test_eof
	_test_eof636: cs = 636; goto _test_eof
	_test_eof1170: cs = 1170; goto _test_eof
	_test_eof637: cs = 637; goto _test_eof
	_test_eof1171: cs = 1171; goto _test_eof
	_test_eof638: cs = 638; goto _test_eof
	_test_eof1172: cs = 1172; goto _test_eof
	_test_eof639: cs = 639; goto _test_eof
	_test_eof1173: cs = 1173; goto _test_eof
	_test_eof1174: cs = 1174; goto _test_eof
	_test_eof1175: cs = 1175; goto _test_eof
	_test_eof1176: cs = 1176; goto _test_eof
	_test_eof1177: cs = 1177; goto _test_eof
	_test_eof1178: cs = 1178; goto _test_eof
	_test_eof640: cs = 640; goto _test_eof
	_test_eof641: cs = 641; goto _test_eof
	_test_eof642: cs = 642; goto _test_eof
	_test_eof1179: cs = 1179; goto _test_eof
	_test_eof643: cs = 643; goto _test_eof
	_test_eof1180: cs = 1180; goto _test_eof
	_test_eof644: cs = 644; goto _test_eof
	_test_eof645: cs = 645; goto _test_eof
	_test_eof1181: cs = 1181; goto _test_eof
	_test_eof1182: cs = 1182; goto _test_eof
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
	_test_eof1183: cs = 1183; goto _test_eof
	_test_eof1184: cs = 1184; goto _test_eof
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
	_test_eof867: cs = 867; goto _test_eof
	_test_eof868: cs = 868; goto _test_eof
	_test_eof869: cs = 869; goto _test_eof
	_test_eof870: cs = 870; goto _test_eof
	_test_eof871: cs = 871; goto _test_eof
	_test_eof872: cs = 872; goto _test_eof
	_test_eof873: cs = 873; goto _test_eof
	_test_eof874: cs = 874; goto _test_eof
	_test_eof875: cs = 875; goto _test_eof
	_test_eof876: cs = 876; goto _test_eof
	_test_eof877: cs = 877; goto _test_eof
	_test_eof878: cs = 878; goto _test_eof
	_test_eof879: cs = 879; goto _test_eof
	_test_eof880: cs = 880; goto _test_eof
	_test_eof881: cs = 881; goto _test_eof
	_test_eof882: cs = 882; goto _test_eof
	_test_eof883: cs = 883; goto _test_eof
	_test_eof884: cs = 884; goto _test_eof
	_test_eof885: cs = 885; goto _test_eof
	_test_eof886: cs = 886; goto _test_eof
	_test_eof887: cs = 887; goto _test_eof
	_test_eof888: cs = 888; goto _test_eof
	_test_eof889: cs = 889; goto _test_eof
	_test_eof890: cs = 890; goto _test_eof
	_test_eof891: cs = 891; goto _test_eof
	_test_eof892: cs = 892; goto _test_eof
	_test_eof893: cs = 893; goto _test_eof
	_test_eof894: cs = 894; goto _test_eof
	_test_eof895: cs = 895; goto _test_eof
	_test_eof896: cs = 896; goto _test_eof
	_test_eof897: cs = 897; goto _test_eof
	_test_eof898: cs = 898; goto _test_eof
	_test_eof899: cs = 899; goto _test_eof
	_test_eof900: cs = 900; goto _test_eof
	_test_eof901: cs = 901; goto _test_eof
	_test_eof902: cs = 902; goto _test_eof
	_test_eof903: cs = 903; goto _test_eof
	_test_eof904: cs = 904; goto _test_eof
	_test_eof905: cs = 905; goto _test_eof
	_test_eof906: cs = 906; goto _test_eof
	_test_eof907: cs = 907; goto _test_eof
	_test_eof908: cs = 908; goto _test_eof
	_test_eof909: cs = 909; goto _test_eof
	_test_eof910: cs = 910; goto _test_eof
	_test_eof911: cs = 911; goto _test_eof
	_test_eof912: cs = 912; goto _test_eof
	_test_eof913: cs = 913; goto _test_eof
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
	_test_eof927: cs = 927; goto _test_eof
	_test_eof928: cs = 928; goto _test_eof
	_test_eof929: cs = 929; goto _test_eof
	_test_eof930: cs = 930; goto _test_eof
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
	_test_eof944: cs = 944; goto _test_eof
	_test_eof945: cs = 945; goto _test_eof
	_test_eof946: cs = 946; goto _test_eof
	_test_eof947: cs = 947; goto _test_eof
	_test_eof948: cs = 948; goto _test_eof
	_test_eof949: cs = 949; goto _test_eof
	_test_eof950: cs = 950; goto _test_eof
	_test_eof951: cs = 951; goto _test_eof
	_test_eof952: cs = 952; goto _test_eof
	_test_eof953: cs = 953; goto _test_eof
	_test_eof954: cs = 954; goto _test_eof
	_test_eof955: cs = 955; goto _test_eof
	_test_eof956: cs = 956; goto _test_eof
	_test_eof957: cs = 957; goto _test_eof
	_test_eof958: cs = 958; goto _test_eof
	_test_eof959: cs = 959; goto _test_eof
	_test_eof960: cs = 960; goto _test_eof
	_test_eof961: cs = 961; goto _test_eof
	_test_eof962: cs = 962; goto _test_eof
	_test_eof963: cs = 963; goto _test_eof
	_test_eof964: cs = 964; goto _test_eof
	_test_eof965: cs = 965; goto _test_eof
	_test_eof966: cs = 966; goto _test_eof
	_test_eof967: cs = 967; goto _test_eof
	_test_eof968: cs = 968; goto _test_eof
	_test_eof969: cs = 969; goto _test_eof
	_test_eof970: cs = 970; goto _test_eof
	_test_eof971: cs = 971; goto _test_eof
	_test_eof972: cs = 972; goto _test_eof
	_test_eof973: cs = 973; goto _test_eof
	_test_eof974: cs = 974; goto _test_eof
	_test_eof975: cs = 975; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 984, 988, 997, 1008, 1039, 1080, 1084, 1091, 1096, 1162, 1166, 1175, 1182, 1184:
//line ragel/datetime.rl:7
 st.Zoned = true 
		case 1046, 1047, 1048:
//line ragel/datetime.rl:9

    st.Month, _ = strconv.Atoi(data[pb:p])

		case 1075:
//line ragel/datetime.rl:13

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

		case 976, 1077, 1098, 1099, 1151, 1155, 1158, 1159:
//line ragel/datetime.rl:17

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

		case 1078, 1097, 1105, 1107, 1108, 1109, 1144, 1149, 1153, 1154, 1156, 1157:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

		case 1045, 1074:
//line ragel/datetime.rl:25

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

		case 1010, 1020, 1115, 1125:
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

		case 982, 994, 1007, 1172, 1181:
//line ragel/datetime.rl:67

    st.Ad_bc = ADBC_BC;

		case 1005, 1112:
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

		case 1058, 1059:
//line ragel/datetime.rl:94
 st.Month = 1 
		case 1056, 1057:
//line ragel/datetime.rl:95
 st.Month = 2 
		case 1064, 1065:
//line ragel/datetime.rl:96
 st.Month = 3 
		case 1049, 1051:
//line ragel/datetime.rl:97
 st.Month = 4 
		case 1066:
//line ragel/datetime.rl:98
 st.Month = 5 
		case 1062, 1063:
//line ragel/datetime.rl:99
 st.Month = 6 
		case 1060, 1061:
//line ragel/datetime.rl:100
 st.Month = 7 
		case 1052, 1053:
//line ragel/datetime.rl:101
 st.Month = 8 
		case 1071, 1072, 1073:
//line ragel/datetime.rl:102
 st.Month = 9 
		case 1069, 1070:
//line ragel/datetime.rl:103
 st.Month = 10 
		case 1067, 1068:
//line ragel/datetime.rl:104
 st.Month = 11 
		case 1054, 1055:
//line ragel/datetime.rl:105
 st.Month = 12 
		case 977, 1043, 1044:
//line ragel/datetime.rl:107

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

		case 1009, 1114:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

		case 1002, 1036, 1037, 1101, 1104, 1106, 1110, 1141, 1142, 1145, 1148, 1150:
//line ragel/datetime.rl:119

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

		case 1022, 1127:
//line ragel/datetime.rl:122

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

		case 1021, 1035, 1126, 1140:
//line ragel/datetime.rl:125

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

		case 1033, 1138:
//line ragel/datetime.rl:128

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

		case 1023, 1034, 1128, 1139:
//line ragel/datetime.rl:131

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

		case 1011, 1012, 1013, 1014, 1015, 1016, 1017, 1018, 1019, 1024, 1025, 1026, 1027, 1028, 1029, 1030, 1031, 1032, 1116, 1117, 1118, 1119, 1120, 1121, 1122, 1123, 1124, 1129, 1130, 1131, 1132, 1133, 1134, 1135, 1136, 1137:
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

		case 980, 981:
//line ragel/datetime.rl:164

    if dot_index := strings.IndexRune(data[pb:p], '.'); dot_index == -1 { // no '.'
        monotonic_offset_sec := parse_digits(data[pb:p])
        st.MonotonicOffsetNanosecond = int64(monotonic_offset_sec) * 1000000000
    }else {
        monotonic_offset_sec := parse_digits(data[pb:pb+dot_index])
        monotonic_offset_nsec := parse_digits(data[pb+dot_index+1:p])
        st.MonotonicOffsetNanosecond = int64(monotonic_offset_sec) * 1000000000 + int64(monotonic_offset_nsec)
    }

		case 1147:
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

		case 1146:
//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

		case 1103:
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

		case 1102:
//line ragel/datetime.rl:116

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:21

    st.Year = parse_year_2_digits(data[pb:pb+2])

		case 989, 990, 998, 999, 1085, 1086, 1092, 1093, 1167, 1168, 1176, 1177:
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
		case 983, 985, 986, 987, 991, 995, 996, 1000, 1079, 1081, 1082, 1083, 1087, 1089, 1090, 1094, 1161, 1163, 1164, 1165, 1169, 1173, 1174, 1178:
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
		case 992, 1001, 1041, 1042, 1088, 1095, 1170, 1179:
//line ragel/datetime.rl:227

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:7
 st.Zoned = true 
//line ragel_parse_datetime.go:31543
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

