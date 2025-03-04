
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
const datetime_parser_first_final int = 660
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
	case 660:
		goto st_case_660
	case 661:
		goto st_case_661
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 662:
		goto st_case_662
	case 11:
		goto st_case_11
	case 663:
		goto st_case_663
	case 12:
		goto st_case_12
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
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 671:
		goto st_case_671
	case 672:
		goto st_case_672
	case 673:
		goto st_case_673
	case 16:
		goto st_case_16
	case 674:
		goto st_case_674
	case 675:
		goto st_case_675
	case 676:
		goto st_case_676
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 677:
		goto st_case_677
	case 678:
		goto st_case_678
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 679:
		goto st_case_679
	case 21:
		goto st_case_21
	case 680:
		goto st_case_680
	case 22:
		goto st_case_22
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
	case 23:
		goto st_case_23
	case 690:
		goto st_case_690
	case 24:
		goto st_case_24
	case 691:
		goto st_case_691
	case 692:
		goto st_case_692
	case 25:
		goto st_case_25
	case 693:
		goto st_case_693
	case 26:
		goto st_case_26
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
	case 27:
		goto st_case_27
	case 707:
		goto st_case_707
	case 708:
		goto st_case_708
	case 709:
		goto st_case_709
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 710:
		goto st_case_710
	case 711:
		goto st_case_711
	case 712:
		goto st_case_712
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
	case 713:
		goto st_case_713
	case 714:
		goto st_case_714
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
	case 715:
		goto st_case_715
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
	case 716:
		goto st_case_716
	case 193:
		goto st_case_193
	case 194:
		goto st_case_194
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
	case 195:
		goto st_case_195
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
	case 724:
		goto st_case_724
	case 725:
		goto st_case_725
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
	case 726:
		goto st_case_726
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
	case 727:
		goto st_case_727
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
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 48:
			goto tr1
		case 51:
			goto tr3
		case 65:
			goto st181
		case 68:
			goto st256
		case 70:
			goto st264
		case 74:
			goto st603
		case 77:
			goto st615
		case 78:
			goto st622
		case 79:
			goto st630
		case 83:
			goto st637
		case 84:
			goto st650
		case 87:
			goto st656
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
//line ragel/parse_datetime.go:1556
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
//line ragel/parse_datetime.go:1570
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
			goto tr20
		}
		switch {
		case data[p] > 47:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr21
			}
		case data[p] >= 45:
			goto tr19
		}
		goto tr0
tr19:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line ragel/parse_datetime.go:1624
		switch data[p] {
		case 48:
			goto tr22
		case 49:
			goto tr23
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
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr24
		}
		goto tr0
tr22:
//line ragel/datetime.rl:7
 pb = p 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line ragel/parse_datetime.go:1660
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
			goto st660
		}
		goto tr0
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
		switch data[p] {
		case 32:
			goto tr873
		case 43:
			goto tr875
		case 45:
			goto tr876
		case 47:
			goto tr877
		case 84:
			goto tr878
		case 90:
			goto tr879
		case 95:
			goto tr880
		case 116:
			goto tr880
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr877
			}
		case data[p] >= 65:
			goto tr877
		}
		goto st0
tr1019:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st661
tr990:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st661
tr873:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st661
tr998:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st661
tr1005:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st661
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
	st_case_661:
//line ragel/parse_datetime.go:1750
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr884
		case 50:
			goto tr61
		case 65:
			goto tr885
		case 66:
			goto tr886
		case 90:
			goto tr887
		case 95:
			goto tr884
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr60
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr884
				}
			case data[p] >= 67:
				goto tr884
			}
		default:
			goto tr62
		}
		goto st0
tr888:
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
tr892:
//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st9
tr895:
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
tr897:
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
//line ragel/parse_datetime.go:1852
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
			goto st662
		}
		goto tr0
	st662:
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 67 {
			goto st663
		}
		goto tr0
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
		goto st0
tr1020:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st12
tr991:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st12
tr899:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st12
tr911:
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
tr915:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st12
tr923:
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
tr930:
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
tr943:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st12
tr952:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st12
tr960:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st12
tr980:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st12
tr875:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st12
tr999:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st12
tr1006:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line ragel/parse_datetime.go:2040
		if data[p] == 50 {
			goto tr41
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr42
			}
		case data[p] >= 48:
			goto tr40
		}
		goto tr0
tr40:
//line ragel/datetime.rl:7
 pb = p 
	goto st664
tr43:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st664
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
//line ragel/parse_datetime.go:2068
		switch data[p] {
		case 32:
			goto tr888
		case 58:
			goto tr890
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st665
		}
		goto st0
tr42:
//line ragel/datetime.rl:7
 pb = p 
	goto st665
tr45:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st665
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
//line ragel/parse_datetime.go:2094
		switch data[p] {
		case 32:
			goto tr888
		case 58:
			goto tr890
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st666
		}
		goto st0
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
		if data[p] == 32 {
			goto tr888
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st666
		}
		goto st0
tr890:
//line ragel/datetime.rl:150

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st667
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
	st_case_667:
//line ragel/parse_datetime.go:2133
		if data[p] == 32 {
			goto tr892
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr894
			}
		case data[p] >= 48:
			goto tr893
		}
		goto st0
tr893:
//line ragel/datetime.rl:7
 pb = p 
	goto st668
	st668:
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
//line ragel/parse_datetime.go:2155
		if data[p] == 32 {
			goto tr895
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st669
		}
		goto st0
tr894:
//line ragel/datetime.rl:7
 pb = p 
	goto st669
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
//line ragel/parse_datetime.go:2172
		if data[p] == 32 {
			goto tr895
		}
		goto st0
tr41:
//line ragel/datetime.rl:7
 pb = p 
	goto st670
tr44:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st670
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
//line ragel/parse_datetime.go:2192
		switch data[p] {
		case 32:
			goto tr888
		case 58:
			goto tr890
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st666
			}
		case data[p] >= 48:
			goto st665
		}
		goto st0
tr1021:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st13
tr992:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st13
tr900:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st13
tr912:
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
tr916:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st13
tr924:
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
tr931:
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
tr944:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st13
tr953:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st13
tr961:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st13
tr981:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st13
tr876:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

	goto st13
tr1000:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

	goto st13
tr1007:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line ragel/parse_datetime.go:2358
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
tr884:
//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1022:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr901:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr917:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr945:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr954:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr963:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr932:
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
tr983:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr926:
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
tr877:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr993:
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
tr1001:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st14
tr1008:
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
//line ragel/parse_datetime.go:2526
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
tr988:
//line ragel/datetime.rl:7
 pb = p 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line ragel/parse_datetime.go:2551
		switch data[p] {
		case 47:
			goto st671
		case 95:
			goto st671
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st671
			}
		case data[p] >= 65:
			goto st671
		}
		goto tr0
tr913:
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
	goto st671
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
//line ragel/parse_datetime.go:2599
		switch data[p] {
		case 32:
			goto tr897
		case 47:
			goto st671
		case 95:
			goto st671
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st671
			}
		case data[p] >= 65:
			goto st671
		}
		goto st0
tr60:
//line ragel/datetime.rl:7
 pb = p 
	goto st672
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
	st_case_672:
//line ragel/parse_datetime.go:2626
		switch data[p] {
		case 32:
			goto tr898
		case 43:
			goto tr899
		case 45:
			goto tr900
		case 47:
			goto tr901
		case 58:
			goto tr903
		case 65:
			goto tr904
		case 80:
			goto tr904
		case 90:
			goto tr905
		case 95:
			goto tr901
		case 97:
			goto tr906
		case 112:
			goto tr906
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st679
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr901
			}
		default:
			goto tr901
		}
		goto st0
tr898:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st673
tr914:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st673
tr968:
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

	goto st673
tr942:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st673
tr951:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st673
tr959:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st673
tr979:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st673
	st673:
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
//line ragel/parse_datetime.go:2736
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr884
		case 65:
			goto tr907
		case 66:
			goto tr886
		case 80:
			goto tr908
		case 90:
			goto tr887
		case 95:
			goto tr884
		case 97:
			goto tr909
		case 112:
			goto tr909
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr884
			}
		case data[p] >= 67:
			goto tr884
		}
		goto st0
tr907:
//line ragel/datetime.rl:7
 pb = p 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line ragel/parse_datetime.go:2779
		switch data[p] {
		case 47:
			goto st15
		case 68:
			goto st674
		case 77:
			goto st675
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
	st674:
		if p++; p == pe {
			goto _test_eof674
		}
	st_case_674:
		switch data[p] {
		case 47:
			goto st671
		case 95:
			goto st671
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st671
			}
		case data[p] >= 65:
			goto st671
		}
		goto st0
	st675:
		if p++; p == pe {
			goto _test_eof675
		}
	st_case_675:
		switch data[p] {
		case 32:
			goto tr910
		case 43:
			goto tr911
		case 45:
			goto tr912
		case 47:
			goto tr913
		case 95:
			goto tr913
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr913
			}
		case data[p] >= 65:
			goto tr913
		}
		goto st0
tr910:
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

	goto st676
tr922:
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

	goto st676
tr929:
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

	goto st676
	st676:
		if p++; p == pe {
			goto _test_eof676
		}
	st_case_676:
//line ragel/parse_datetime.go:2923
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr884
		case 65:
			goto tr885
		case 66:
			goto tr886
		case 90:
			goto tr887
		case 95:
			goto tr884
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr884
			}
		case data[p] >= 67:
			goto tr884
		}
		goto st0
tr885:
//line ragel/datetime.rl:7
 pb = p 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line ragel/parse_datetime.go:2960
		switch data[p] {
		case 47:
			goto st15
		case 68:
			goto st674
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
tr886:
//line ragel/datetime.rl:7
 pb = p 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line ragel/parse_datetime.go:2987
		switch data[p] {
		case 47:
			goto st15
		case 67:
			goto st677
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
	st677:
		if p++; p == pe {
			goto _test_eof677
		}
	st_case_677:
		switch data[p] {
		case 47:
			goto st671
		case 95:
			goto st671
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st671
			}
		case data[p] >= 65:
			goto st671
		}
		goto st0
tr887:
//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr1024:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr905:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr920:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr949:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr957:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr966:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr934:
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
	goto st678
tr985:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr928:
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
	goto st678
tr879:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr995:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr1003:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
tr1010:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st678
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
//line ragel/parse_datetime.go:3180
		switch data[p] {
		case 32:
			goto tr892
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
tr908:
//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr904:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr919:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr948:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr956:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr965:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st19
tr970:
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
tr984:
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
//line ragel/parse_datetime.go:3288
		switch data[p] {
		case 47:
			goto st15
		case 77:
			goto st675
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
tr909:
//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr906:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr921:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr950:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr958:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr967:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st20
tr971:
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
tr986:
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
//line ragel/parse_datetime.go:3396
		switch data[p] {
		case 47:
			goto st15
		case 95:
			goto st15
		case 109:
			goto st675
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
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
		switch data[p] {
		case 32:
			goto tr914
		case 43:
			goto tr915
		case 45:
			goto tr916
		case 47:
			goto tr917
		case 58:
			goto tr918
		case 65:
			goto tr919
		case 80:
			goto tr919
		case 90:
			goto tr920
		case 95:
			goto tr917
		case 97:
			goto tr921
		case 112:
			goto tr921
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st21
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr917
			}
		default:
			goto tr917
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= data[p] && data[p] <= 57 {
			goto st680
		}
		goto tr0
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
		switch data[p] {
		case 32:
			goto tr922
		case 43:
			goto tr923
		case 45:
			goto tr924
		case 46:
			goto tr925
		case 47:
			goto tr926
		case 90:
			goto tr928
		case 95:
			goto tr926
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr926
			}
		default:
			goto tr926
		}
		goto st0
tr925:
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
//line ragel/parse_datetime.go:3521
		if 48 <= data[p] && data[p] <= 57 {
			goto tr52
		}
		goto tr0
tr52:
//line ragel/datetime.rl:7
 pb = p 
	goto st681
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
	st_case_681:
//line ragel/parse_datetime.go:3535
		switch data[p] {
		case 32:
			goto tr929
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 90:
			goto tr934
		case 95:
			goto tr932
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st682
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st682:
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
		switch data[p] {
		case 32:
			goto tr929
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 90:
			goto tr934
		case 95:
			goto tr932
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st683
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st683:
		if p++; p == pe {
			goto _test_eof683
		}
	st_case_683:
		switch data[p] {
		case 32:
			goto tr929
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 90:
			goto tr934
		case 95:
			goto tr932
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st684
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st684:
		if p++; p == pe {
			goto _test_eof684
		}
	st_case_684:
		switch data[p] {
		case 32:
			goto tr929
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 90:
			goto tr934
		case 95:
			goto tr932
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st685
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
	st_case_685:
		switch data[p] {
		case 32:
			goto tr929
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 90:
			goto tr934
		case 95:
			goto tr932
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st686
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st686:
		if p++; p == pe {
			goto _test_eof686
		}
	st_case_686:
		switch data[p] {
		case 32:
			goto tr929
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 90:
			goto tr934
		case 95:
			goto tr932
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st687
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st687:
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
		switch data[p] {
		case 32:
			goto tr929
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 90:
			goto tr934
		case 95:
			goto tr932
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st688
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
		switch data[p] {
		case 32:
			goto tr929
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 90:
			goto tr934
		case 95:
			goto tr932
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st689
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
		switch data[p] {
		case 32:
			goto tr929
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 90:
			goto tr934
		case 95:
			goto tr932
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		case data[p] >= 65:
			goto tr932
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if 48 <= data[p] && data[p] <= 57 {
			goto st690
		}
		goto tr0
	st690:
		if p++; p == pe {
			goto _test_eof690
		}
	st_case_690:
		switch data[p] {
		case 32:
			goto tr922
		case 43:
			goto tr923
		case 45:
			goto tr924
		case 46:
			goto tr925
		case 47:
			goto tr926
		case 90:
			goto tr928
		case 95:
			goto tr926
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr926
			}
		case data[p] >= 65:
			goto tr926
		}
		goto st0
tr903:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st24
tr918:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line ragel/parse_datetime.go:3871
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr55
			}
		case data[p] >= 48:
			goto tr54
		}
		goto tr0
tr54:
//line ragel/datetime.rl:7
 pb = p 
	goto st691
	st691:
		if p++; p == pe {
			goto _test_eof691
		}
	st_case_691:
//line ragel/parse_datetime.go:3890
		switch data[p] {
		case 32:
			goto tr942
		case 43:
			goto tr943
		case 45:
			goto tr944
		case 47:
			goto tr945
		case 58:
			goto tr947
		case 65:
			goto tr948
		case 80:
			goto tr948
		case 90:
			goto tr949
		case 95:
			goto tr945
		case 97:
			goto tr950
		case 112:
			goto tr950
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st692
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr945
			}
		default:
			goto tr945
		}
		goto st0
	st692:
		if p++; p == pe {
			goto _test_eof692
		}
	st_case_692:
		switch data[p] {
		case 32:
			goto tr951
		case 43:
			goto tr952
		case 45:
			goto tr953
		case 47:
			goto tr954
		case 58:
			goto tr955
		case 65:
			goto tr956
		case 80:
			goto tr956
		case 90:
			goto tr957
		case 95:
			goto tr954
		case 97:
			goto tr958
		case 112:
			goto tr958
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr954
			}
		case data[p] >= 66:
			goto tr954
		}
		goto st0
tr947:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st25
tr955:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line ragel/parse_datetime.go:3983
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr57
			}
		case data[p] >= 48:
			goto tr56
		}
		goto tr0
tr56:
//line ragel/datetime.rl:7
 pb = p 
	goto st693
	st693:
		if p++; p == pe {
			goto _test_eof693
		}
	st_case_693:
//line ragel/parse_datetime.go:4002
		switch data[p] {
		case 32:
			goto tr959
		case 43:
			goto tr960
		case 45:
			goto tr961
		case 46:
			goto tr962
		case 47:
			goto tr963
		case 65:
			goto tr965
		case 80:
			goto tr965
		case 90:
			goto tr966
		case 95:
			goto tr963
		case 97:
			goto tr967
		case 112:
			goto tr967
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st703
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr963
			}
		default:
			goto tr963
		}
		goto st0
tr962:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st26
tr982:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line ragel/parse_datetime.go:4057
		if 48 <= data[p] && data[p] <= 57 {
			goto tr58
		}
		goto tr0
tr58:
//line ragel/datetime.rl:7
 pb = p 
	goto st694
	st694:
		if p++; p == pe {
			goto _test_eof694
		}
	st_case_694:
//line ragel/parse_datetime.go:4071
		switch data[p] {
		case 32:
			goto tr968
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 65:
			goto tr970
		case 80:
			goto tr970
		case 90:
			goto tr934
		case 95:
			goto tr932
		case 97:
			goto tr971
		case 112:
			goto tr971
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st695
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st695:
		if p++; p == pe {
			goto _test_eof695
		}
	st_case_695:
		switch data[p] {
		case 32:
			goto tr968
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 65:
			goto tr970
		case 80:
			goto tr970
		case 90:
			goto tr934
		case 95:
			goto tr932
		case 97:
			goto tr971
		case 112:
			goto tr971
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st696
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st696:
		if p++; p == pe {
			goto _test_eof696
		}
	st_case_696:
		switch data[p] {
		case 32:
			goto tr968
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 65:
			goto tr970
		case 80:
			goto tr970
		case 90:
			goto tr934
		case 95:
			goto tr932
		case 97:
			goto tr971
		case 112:
			goto tr971
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st697
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st697:
		if p++; p == pe {
			goto _test_eof697
		}
	st_case_697:
		switch data[p] {
		case 32:
			goto tr968
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 65:
			goto tr970
		case 80:
			goto tr970
		case 90:
			goto tr934
		case 95:
			goto tr932
		case 97:
			goto tr971
		case 112:
			goto tr971
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st698
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st698:
		if p++; p == pe {
			goto _test_eof698
		}
	st_case_698:
		switch data[p] {
		case 32:
			goto tr968
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 65:
			goto tr970
		case 80:
			goto tr970
		case 90:
			goto tr934
		case 95:
			goto tr932
		case 97:
			goto tr971
		case 112:
			goto tr971
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st699
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st699:
		if p++; p == pe {
			goto _test_eof699
		}
	st_case_699:
		switch data[p] {
		case 32:
			goto tr968
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 65:
			goto tr970
		case 80:
			goto tr970
		case 90:
			goto tr934
		case 95:
			goto tr932
		case 97:
			goto tr971
		case 112:
			goto tr971
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st700
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st700:
		if p++; p == pe {
			goto _test_eof700
		}
	st_case_700:
		switch data[p] {
		case 32:
			goto tr968
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 65:
			goto tr970
		case 80:
			goto tr970
		case 90:
			goto tr934
		case 95:
			goto tr932
		case 97:
			goto tr971
		case 112:
			goto tr971
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st701
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st701:
		if p++; p == pe {
			goto _test_eof701
		}
	st_case_701:
		switch data[p] {
		case 32:
			goto tr968
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 65:
			goto tr970
		case 80:
			goto tr970
		case 90:
			goto tr934
		case 95:
			goto tr932
		case 97:
			goto tr971
		case 112:
			goto tr971
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st702
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		default:
			goto tr932
		}
		goto st0
	st702:
		if p++; p == pe {
			goto _test_eof702
		}
	st_case_702:
		switch data[p] {
		case 32:
			goto tr968
		case 43:
			goto tr930
		case 45:
			goto tr931
		case 47:
			goto tr932
		case 65:
			goto tr970
		case 80:
			goto tr970
		case 90:
			goto tr934
		case 95:
			goto tr932
		case 97:
			goto tr971
		case 112:
			goto tr971
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr932
			}
		case data[p] >= 66:
			goto tr932
		}
		goto st0
	st703:
		if p++; p == pe {
			goto _test_eof703
		}
	st_case_703:
		switch data[p] {
		case 32:
			goto tr979
		case 43:
			goto tr980
		case 45:
			goto tr981
		case 46:
			goto tr982
		case 47:
			goto tr983
		case 65:
			goto tr984
		case 80:
			goto tr984
		case 90:
			goto tr985
		case 95:
			goto tr983
		case 97:
			goto tr986
		case 112:
			goto tr986
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr983
			}
		case data[p] >= 66:
			goto tr983
		}
		goto st0
tr57:
//line ragel/datetime.rl:7
 pb = p 
	goto st704
	st704:
		if p++; p == pe {
			goto _test_eof704
		}
	st_case_704:
//line ragel/parse_datetime.go:4470
		switch data[p] {
		case 32:
			goto tr959
		case 43:
			goto tr960
		case 45:
			goto tr961
		case 46:
			goto tr962
		case 47:
			goto tr963
		case 65:
			goto tr965
		case 80:
			goto tr965
		case 90:
			goto tr966
		case 95:
			goto tr963
		case 97:
			goto tr967
		case 112:
			goto tr967
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr963
			}
		case data[p] >= 66:
			goto tr963
		}
		goto st0
tr55:
//line ragel/datetime.rl:7
 pb = p 
	goto st705
	st705:
		if p++; p == pe {
			goto _test_eof705
		}
	st_case_705:
//line ragel/parse_datetime.go:4513
		switch data[p] {
		case 32:
			goto tr942
		case 43:
			goto tr943
		case 45:
			goto tr944
		case 47:
			goto tr945
		case 58:
			goto tr947
		case 65:
			goto tr948
		case 80:
			goto tr948
		case 90:
			goto tr949
		case 95:
			goto tr945
		case 97:
			goto tr950
		case 112:
			goto tr950
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr945
			}
		case data[p] >= 66:
			goto tr945
		}
		goto st0
tr61:
//line ragel/datetime.rl:7
 pb = p 
	goto st706
	st706:
		if p++; p == pe {
			goto _test_eof706
		}
	st_case_706:
//line ragel/parse_datetime.go:4556
		switch data[p] {
		case 32:
			goto tr898
		case 43:
			goto tr899
		case 45:
			goto tr900
		case 47:
			goto tr901
		case 58:
			goto tr903
		case 65:
			goto tr904
		case 80:
			goto tr904
		case 90:
			goto tr905
		case 95:
			goto tr901
		case 97:
			goto tr906
		case 112:
			goto tr906
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st679
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr901
				}
			case data[p] >= 66:
				goto tr901
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
tr62:
//line ragel/datetime.rl:7
 pb = p 
	goto st707
	st707:
		if p++; p == pe {
			goto _test_eof707
		}
	st_case_707:
//line ragel/parse_datetime.go:4617
		switch data[p] {
		case 32:
			goto tr898
		case 43:
			goto tr899
		case 45:
			goto tr900
		case 47:
			goto tr901
		case 58:
			goto tr903
		case 65:
			goto tr904
		case 80:
			goto tr904
		case 90:
			goto tr905
		case 95:
			goto tr901
		case 97:
			goto tr906
		case 112:
			goto tr906
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st27
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr901
			}
		default:
			goto tr901
		}
		goto st0
tr1023:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st708
tr878:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st708
tr994:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

//line ragel/datetime.rl:7
 pb = p 
	goto st708
tr1002:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st708
tr1009:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st708
	st708:
		if p++; p == pe {
			goto _test_eof708
		}
	st_case_708:
//line ragel/parse_datetime.go:4706
		switch data[p] {
		case 32:
			goto st9
		case 43:
			goto st12
		case 45:
			goto st13
		case 47:
			goto tr988
		case 50:
			goto tr61
		case 90:
			goto tr989
		case 95:
			goto tr988
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr60
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr988
				}
			case data[p] >= 65:
				goto tr988
			}
		default:
			goto tr62
		}
		goto st0
tr989:
//line ragel/datetime.rl:7
 pb = p 
	goto st709
	st709:
		if p++; p == pe {
			goto _test_eof709
		}
	st_case_709:
//line ragel/parse_datetime.go:4750
		switch data[p] {
		case 32:
			goto tr892
		case 47:
			goto st671
		case 95:
			goto st671
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st671
			}
		case data[p] >= 65:
			goto st671
		}
		goto st0
tr1025:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st28
tr880:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

//line ragel/datetime.rl:7
 pb = p 
	goto st28
tr996:
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
tr1004:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

//line ragel/datetime.rl:7
 pb = p 
	goto st28
tr1011:
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
//line ragel/parse_datetime.go:4819
		switch data[p] {
		case 47:
			goto st15
		case 50:
			goto tr61
		case 95:
			goto st15
		}
		switch {
		case data[p] < 51:
			if 48 <= data[p] && data[p] <= 49 {
				goto tr60
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
			goto tr62
		}
		goto tr0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 45:
			goto tr63
		case 47:
			goto tr63
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st660
		}
		goto tr0
tr63:
//line ragel/datetime.rl:11

    st.Month, _ = strconv.Atoi(data[pb:p])

	goto st30
tr72:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st30
tr76:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st30
tr82:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st30
tr90:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st30
tr99:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st30
tr106:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st30
tr108:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st30
tr113:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st30
tr116:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st30
tr119:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st30
tr127:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st30
tr134:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line ragel/parse_datetime.go:4920
		switch data[p] {
		case 48:
			goto tr64
		case 51:
			goto tr66
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr67
			}
		case data[p] >= 49:
			goto tr65
		}
		goto tr0
tr64:
//line ragel/datetime.rl:7
 pb = p 
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line ragel/parse_datetime.go:4945
		if 49 <= data[p] && data[p] <= 57 {
			goto st710
		}
		goto tr0
tr67:
//line ragel/datetime.rl:7
 pb = p 
	goto st710
	st710:
		if p++; p == pe {
			goto _test_eof710
		}
	st_case_710:
//line ragel/parse_datetime.go:4959
		switch data[p] {
		case 32:
			goto tr990
		case 43:
			goto tr991
		case 45:
			goto tr992
		case 47:
			goto tr993
		case 84:
			goto tr994
		case 90:
			goto tr995
		case 95:
			goto tr996
		case 116:
			goto tr996
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr993
			}
		case data[p] >= 65:
			goto tr993
		}
		goto st0
tr65:
//line ragel/datetime.rl:7
 pb = p 
	goto st711
	st711:
		if p++; p == pe {
			goto _test_eof711
		}
	st_case_711:
//line ragel/parse_datetime.go:4996
		switch data[p] {
		case 32:
			goto tr990
		case 43:
			goto tr991
		case 45:
			goto tr992
		case 47:
			goto tr993
		case 84:
			goto tr994
		case 90:
			goto tr995
		case 95:
			goto tr996
		case 116:
			goto tr996
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st710
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr993
			}
		default:
			goto tr993
		}
		goto st0
tr66:
//line ragel/datetime.rl:7
 pb = p 
	goto st712
	st712:
		if p++; p == pe {
			goto _test_eof712
		}
	st_case_712:
//line ragel/parse_datetime.go:5037
		switch data[p] {
		case 32:
			goto tr990
		case 43:
			goto tr991
		case 45:
			goto tr992
		case 47:
			goto tr993
		case 84:
			goto tr994
		case 90:
			goto tr995
		case 95:
			goto tr996
		case 116:
			goto tr996
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 49 {
				goto st710
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr993
			}
		default:
			goto tr993
		}
		goto st0
tr23:
//line ragel/datetime.rl:7
 pb = p 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line ragel/parse_datetime.go:5078
		switch data[p] {
		case 45:
			goto tr63
		case 47:
			goto tr63
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
tr24:
//line ragel/datetime.rl:7
 pb = p 
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line ragel/parse_datetime.go:5103
		switch data[p] {
		case 45:
			goto tr63
		case 47:
			goto tr63
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
			goto tr72
		case 47:
			goto tr72
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
			goto tr72
		case 47:
			goto tr72
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
			goto tr76
		case 47:
			goto tr76
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
			goto tr76
		case 47:
			goto tr76
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
			goto tr82
		case 47:
			goto tr82
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
			goto tr82
		case 47:
			goto tr82
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
			goto tr90
		case 47:
			goto tr90
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
			goto tr90
		case 47:
			goto tr90
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
			goto tr99
		case 47:
			goto tr99
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
			goto tr99
		case 47:
			goto tr99
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
			goto tr106
		case 47:
			goto tr106
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
			goto tr106
		case 47:
			goto tr106
		}
		goto tr0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 45:
			goto tr108
		case 47:
			goto tr108
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
			goto tr108
		case 47:
			goto tr108
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
			goto tr113
		case 47:
			goto tr113
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
			goto tr113
		case 47:
			goto tr113
		}
		goto tr0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 45:
			goto tr116
		case 47:
			goto tr116
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
			goto tr119
		case 47:
			goto tr119
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
			goto tr119
		case 47:
			goto tr119
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
			goto tr127
		case 47:
			goto tr127
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
			goto tr127
		case 47:
			goto tr127
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
			goto tr134
		case 47:
			goto tr134
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
			goto tr134
		case 47:
			goto tr134
		}
		goto tr0
tr20:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line ragel/parse_datetime.go:5840
		if 48 <= data[p] && data[p] <= 57 {
			goto tr141
		}
		goto tr0
tr141:
//line ragel/datetime.rl:7
 pb = p 
	goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
//line ragel/parse_datetime.go:5854
		if 48 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto tr0
tr21:
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
//line ragel/parse_datetime.go:5872
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
			goto st713
		}
		goto tr0
	st713:
		if p++; p == pe {
			goto _test_eof713
		}
	st_case_713:
		switch data[p] {
		case 32:
			goto tr873
		case 43:
			goto tr875
		case 45:
			goto tr876
		case 47:
			goto tr877
		case 84:
			goto tr878
		case 90:
			goto tr879
		case 95:
			goto tr880
		case 116:
			goto tr880
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st714
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr877
			}
		default:
			goto tr877
		}
		goto st0
	st714:
		if p++; p == pe {
			goto _test_eof714
		}
	st_case_714:
		switch data[p] {
		case 32:
			goto tr998
		case 43:
			goto tr999
		case 45:
			goto tr1000
		case 47:
			goto tr1001
		case 84:
			goto tr1002
		case 90:
			goto tr1003
		case 95:
			goto tr1004
		case 116:
			goto tr1004
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1001
			}
		case data[p] >= 65:
			goto tr1001
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		if data[p] == 32 {
			goto tr144
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto tr0
tr144:
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
//line ragel/parse_datetime.go:5982
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
			goto tr156
		case 105:
			goto st113
		}
		goto tr0
tr156:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st111
tr162:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st111
tr168:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st111
tr176:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st111
tr185:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st111
tr192:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st111
tr194:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st111
tr199:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st111
tr202:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st111
tr205:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st111
tr213:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st111
tr220:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line ragel/parse_datetime.go:6088
		if 48 <= data[p] && data[p] <= 57 {
			goto tr158
		}
		goto tr0
tr158:
//line ragel/datetime.rl:7
 pb = p 
	goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line ragel/parse_datetime.go:6102
		if 48 <= data[p] && data[p] <= 57 {
			goto st715
		}
		goto tr0
	st715:
		if p++; p == pe {
			goto _test_eof715
		}
	st_case_715:
		switch data[p] {
		case 32:
			goto tr1005
		case 43:
			goto tr1006
		case 45:
			goto tr1007
		case 47:
			goto tr1008
		case 84:
			goto tr1009
		case 90:
			goto tr1010
		case 95:
			goto tr1011
		case 116:
			goto tr1011
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1008
			}
		case data[p] >= 65:
			goto tr1008
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
			goto tr156
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
			goto tr162
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
			goto tr162
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
			goto tr168
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
			goto tr168
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
			goto tr176
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
			goto tr176
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
			goto tr185
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
			goto tr185
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
			goto tr192
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
			goto tr192
		}
		goto tr0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 32:
			goto tr194
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
			goto tr194
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
			goto tr199
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
			goto tr199
		}
		goto tr0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		if data[p] == 32 {
			goto tr202
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
			goto tr205
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
			goto tr205
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
			goto tr213
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
			goto tr213
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
			goto tr220
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
			goto tr220
		}
		goto tr0
tr2:
//line ragel/datetime.rl:7
 pb = p 
	goto st178
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
//line ragel/parse_datetime.go:6772
		if data[p] == 32 {
			goto tr144
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st106
		}
		goto tr0
tr3:
//line ragel/datetime.rl:7
 pb = p 
	goto st179
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
//line ragel/parse_datetime.go:6789
		if data[p] == 32 {
			goto tr144
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
	goto st180
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
//line ragel/parse_datetime.go:6811
		if data[p] == 32 {
			goto tr144
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto tr0
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 112:
			goto st182
		case 117:
			goto st251
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
		switch data[p] {
		case 32:
			goto tr230
		case 45:
			goto tr231
		case 47:
			goto tr231
		case 105:
			goto st249
		}
		goto tr0
tr230:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st184
tr333:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st184
tr340:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st184
tr350:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st184
tr810:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st184
tr818:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st184
tr821:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st184
tr828:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st184
tr832:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st184
tr836:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st184
tr845:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st184
tr857:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st184
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
//line ragel/parse_datetime.go:6909
		switch data[p] {
		case 48:
			goto tr233
		case 51:
			goto tr235
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto tr234
		}
		goto tr0
tr233:
//line ragel/datetime.rl:7
 pb = p 
	goto st185
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
//line ragel/parse_datetime.go:6929
		if 49 <= data[p] && data[p] <= 57 {
			goto st186
		}
		goto tr0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		if data[p] == 32 {
			goto tr237
		}
		goto tr0
tr237:
//line ragel/datetime.rl:92

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
//line ragel/parse_datetime.go:6959
		if data[p] == 50 {
			goto tr239
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr240
			}
		case data[p] >= 48:
			goto tr238
		}
		goto tr0
tr238:
//line ragel/datetime.rl:7
 pb = p 
	goto st188
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
//line ragel/parse_datetime.go:6981
		switch data[p] {
		case 32:
			goto tr241
		case 58:
			goto tr243
		case 65:
			goto tr244
		case 80:
			goto tr244
		case 97:
			goto tr245
		case 112:
			goto tr245
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st202
		}
		goto tr0
tr241:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st189
tr266:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st189
tr305:
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

	goto st189
tr288:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st189
tr293:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st189
tr299:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st189
tr316:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st189
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
//line ragel/parse_datetime.go:7072
		switch data[p] {
		case 65:
			goto tr247
		case 80:
			goto tr247
		case 97:
			goto tr248
		case 112:
			goto tr248
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr246
		}
		goto tr0
tr246:
//line ragel/datetime.rl:7
 pb = p 
	goto st190
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
//line ragel/parse_datetime.go:7096
		if 48 <= data[p] && data[p] <= 57 {
			goto st191
		}
		goto tr0
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		if 48 <= data[p] && data[p] <= 57 {
			goto st192
		}
		goto tr0
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		if 48 <= data[p] && data[p] <= 57 {
			goto st716
		}
		goto tr0
	st716:
		if p++; p == pe {
			goto _test_eof716
		}
	st_case_716:
		if data[p] == 32 {
			goto tr1012
		}
		goto st0
tr1012:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

	goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
//line ragel/parse_datetime.go:7139
		switch data[p] {
		case 43:
			goto st194
		case 45:
			goto st195
		case 47:
			goto tr254
		case 90:
			goto tr255
		case 95:
			goto tr254
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr254
			}
		case data[p] >= 65:
			goto tr254
		}
		goto tr0
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		if data[p] == 50 {
			goto tr257
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr258
			}
		case data[p] >= 48:
			goto tr256
		}
		goto tr0
tr256:
//line ragel/datetime.rl:7
 pb = p 
	goto st717
tr259:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st717
	st717:
		if p++; p == pe {
			goto _test_eof717
		}
	st_case_717:
//line ragel/parse_datetime.go:7193
		if data[p] == 58 {
			goto tr1014
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st718
		}
		goto st0
tr258:
//line ragel/datetime.rl:7
 pb = p 
	goto st718
tr261:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st718
	st718:
		if p++; p == pe {
			goto _test_eof718
		}
	st_case_718:
//line ragel/parse_datetime.go:7216
		if data[p] == 58 {
			goto tr1014
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st719
		}
		goto st0
	st719:
		if p++; p == pe {
			goto _test_eof719
		}
	st_case_719:
		if 48 <= data[p] && data[p] <= 57 {
			goto st719
		}
		goto st0
tr1014:
//line ragel/datetime.rl:150

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st720
	st720:
		if p++; p == pe {
			goto _test_eof720
		}
	st_case_720:
//line ragel/parse_datetime.go:7249
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr1017
			}
		case data[p] >= 48:
			goto tr1016
		}
		goto st0
tr1016:
//line ragel/datetime.rl:7
 pb = p 
	goto st721
	st721:
		if p++; p == pe {
			goto _test_eof721
		}
	st_case_721:
//line ragel/parse_datetime.go:7268
		if 48 <= data[p] && data[p] <= 57 {
			goto st722
		}
		goto st0
tr1017:
//line ragel/datetime.rl:7
 pb = p 
	goto st722
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
//line ragel/parse_datetime.go:7282
		goto st0
tr257:
//line ragel/datetime.rl:7
 pb = p 
	goto st723
tr260:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st723
	st723:
		if p++; p == pe {
			goto _test_eof723
		}
	st_case_723:
//line ragel/parse_datetime.go:7299
		if data[p] == 58 {
			goto tr1014
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st719
			}
		case data[p] >= 48:
			goto st718
		}
		goto st0
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		if data[p] == 50 {
			goto tr260
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr261
			}
		case data[p] >= 48:
			goto tr259
		}
		goto tr0
tr254:
//line ragel/datetime.rl:7
 pb = p 
	goto st196
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
//line ragel/parse_datetime.go:7338
		switch data[p] {
		case 47:
			goto st197
		case 95:
			goto st197
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st197
			}
		case data[p] >= 65:
			goto st197
		}
		goto tr0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 47:
			goto st724
		case 95:
			goto st724
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st724
			}
		case data[p] >= 65:
			goto st724
		}
		goto tr0
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
	st_case_724:
		switch data[p] {
		case 47:
			goto st724
		case 95:
			goto st724
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st724
			}
		case data[p] >= 65:
			goto st724
		}
		goto st0
tr255:
//line ragel/datetime.rl:7
 pb = p 
	goto st725
	st725:
		if p++; p == pe {
			goto _test_eof725
		}
	st_case_725:
//line ragel/parse_datetime.go:7403
		switch data[p] {
		case 47:
			goto st197
		case 95:
			goto st197
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st197
			}
		case data[p] >= 65:
			goto st197
		}
		goto st0
tr247:
//line ragel/datetime.rl:7
 pb = p 
	goto st198
tr244:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st198
tr269:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st198
tr291:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st198
tr295:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st198
tr302:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st198
tr307:
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
	goto st198
tr318:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st198
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
//line ragel/parse_datetime.go:7509
		if data[p] == 77 {
			goto st199
		}
		goto tr0
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		if data[p] == 32 {
			goto tr265
		}
		goto tr0
tr265:
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

	goto st200
tr272:
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

	goto st200
tr276:
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

	goto st200
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
//line ragel/parse_datetime.go:7601
		if 48 <= data[p] && data[p] <= 57 {
			goto tr246
		}
		goto tr0
tr248:
//line ragel/datetime.rl:7
 pb = p 
	goto st201
tr245:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st201
tr270:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st201
tr292:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st201
tr296:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st201
tr303:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st201
tr308:
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
	goto st201
tr319:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st201
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
//line ragel/parse_datetime.go:7696
		if data[p] == 109 {
			goto st199
		}
		goto tr0
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		switch data[p] {
		case 32:
			goto tr266
		case 58:
			goto tr268
		case 65:
			goto tr269
		case 80:
			goto tr269
		case 97:
			goto tr270
		case 112:
			goto tr270
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st203
		}
		goto tr0
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		if 48 <= data[p] && data[p] <= 57 {
			goto st204
		}
		goto tr0
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		switch data[p] {
		case 32:
			goto tr272
		case 46:
			goto tr273
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st215
		}
		goto tr0
tr273:
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

	goto st205
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
//line ragel/parse_datetime.go:7770
		if 48 <= data[p] && data[p] <= 57 {
			goto tr275
		}
		goto tr0
tr275:
//line ragel/datetime.rl:7
 pb = p 
	goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
//line ragel/parse_datetime.go:7784
		if data[p] == 32 {
			goto tr276
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st207
		}
		goto tr0
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		if data[p] == 32 {
			goto tr276
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st208
		}
		goto tr0
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		if data[p] == 32 {
			goto tr276
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st209
		}
		goto tr0
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		if data[p] == 32 {
			goto tr276
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st210
		}
		goto tr0
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		if data[p] == 32 {
			goto tr276
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st211
		}
		goto tr0
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		if data[p] == 32 {
			goto tr276
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st212
		}
		goto tr0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		if data[p] == 32 {
			goto tr276
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st213
		}
		goto tr0
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		if data[p] == 32 {
			goto tr276
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st214
		}
		goto tr0
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		if data[p] == 32 {
			goto tr276
		}
		goto tr0
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		if 48 <= data[p] && data[p] <= 57 {
			goto st216
		}
		goto tr0
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 32:
			goto tr272
		case 46:
			goto tr273
		}
		goto tr0
tr243:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st217
tr268:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st217
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
//line ragel/parse_datetime.go:7923
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr287
			}
		case data[p] >= 48:
			goto tr286
		}
		goto tr0
tr286:
//line ragel/datetime.rl:7
 pb = p 
	goto st218
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
//line ragel/parse_datetime.go:7942
		switch data[p] {
		case 32:
			goto tr288
		case 58:
			goto tr290
		case 65:
			goto tr291
		case 80:
			goto tr291
		case 97:
			goto tr292
		case 112:
			goto tr292
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st219
		}
		goto tr0
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		switch data[p] {
		case 32:
			goto tr293
		case 58:
			goto tr294
		case 65:
			goto tr295
		case 80:
			goto tr295
		case 97:
			goto tr296
		case 112:
			goto tr296
		}
		goto tr0
tr290:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st220
tr294:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st220
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
//line ragel/parse_datetime.go:7998
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr298
			}
		case data[p] >= 48:
			goto tr297
		}
		goto tr0
tr297:
//line ragel/datetime.rl:7
 pb = p 
	goto st221
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
//line ragel/parse_datetime.go:8017
		switch data[p] {
		case 32:
			goto tr299
		case 46:
			goto tr300
		case 65:
			goto tr302
		case 80:
			goto tr302
		case 97:
			goto tr303
		case 112:
			goto tr303
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st232
		}
		goto tr0
tr300:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st222
tr317:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st222
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
//line ragel/parse_datetime.go:8053
		if 48 <= data[p] && data[p] <= 57 {
			goto tr304
		}
		goto tr0
tr304:
//line ragel/datetime.rl:7
 pb = p 
	goto st223
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
//line ragel/parse_datetime.go:8067
		switch data[p] {
		case 32:
			goto tr305
		case 65:
			goto tr307
		case 80:
			goto tr307
		case 97:
			goto tr308
		case 112:
			goto tr308
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st224
		}
		goto tr0
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 32:
			goto tr305
		case 65:
			goto tr307
		case 80:
			goto tr307
		case 97:
			goto tr308
		case 112:
			goto tr308
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st225
		}
		goto tr0
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 32:
			goto tr305
		case 65:
			goto tr307
		case 80:
			goto tr307
		case 97:
			goto tr308
		case 112:
			goto tr308
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st226
		}
		goto tr0
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 32:
			goto tr305
		case 65:
			goto tr307
		case 80:
			goto tr307
		case 97:
			goto tr308
		case 112:
			goto tr308
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st227
		}
		goto tr0
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		switch data[p] {
		case 32:
			goto tr305
		case 65:
			goto tr307
		case 80:
			goto tr307
		case 97:
			goto tr308
		case 112:
			goto tr308
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st228
		}
		goto tr0
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		switch data[p] {
		case 32:
			goto tr305
		case 65:
			goto tr307
		case 80:
			goto tr307
		case 97:
			goto tr308
		case 112:
			goto tr308
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st229
		}
		goto tr0
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
		switch data[p] {
		case 32:
			goto tr305
		case 65:
			goto tr307
		case 80:
			goto tr307
		case 97:
			goto tr308
		case 112:
			goto tr308
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st230
		}
		goto tr0
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 32:
			goto tr305
		case 65:
			goto tr307
		case 80:
			goto tr307
		case 97:
			goto tr308
		case 112:
			goto tr308
		}
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
			goto tr305
		case 65:
			goto tr307
		case 80:
			goto tr307
		case 97:
			goto tr308
		case 112:
			goto tr308
		}
		goto tr0
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		switch data[p] {
		case 32:
			goto tr316
		case 46:
			goto tr317
		case 65:
			goto tr318
		case 80:
			goto tr318
		case 97:
			goto tr319
		case 112:
			goto tr319
		}
		goto tr0
tr298:
//line ragel/datetime.rl:7
 pb = p 
	goto st233
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
//line ragel/parse_datetime.go:8278
		switch data[p] {
		case 32:
			goto tr299
		case 46:
			goto tr300
		case 65:
			goto tr302
		case 80:
			goto tr302
		case 97:
			goto tr303
		case 112:
			goto tr303
		}
		goto tr0
tr287:
//line ragel/datetime.rl:7
 pb = p 
	goto st234
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
//line ragel/parse_datetime.go:8303
		switch data[p] {
		case 32:
			goto tr288
		case 58:
			goto tr290
		case 65:
			goto tr291
		case 80:
			goto tr291
		case 97:
			goto tr292
		case 112:
			goto tr292
		}
		goto tr0
tr239:
//line ragel/datetime.rl:7
 pb = p 
	goto st235
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
//line ragel/parse_datetime.go:8328
		switch data[p] {
		case 32:
			goto tr241
		case 58:
			goto tr243
		case 65:
			goto tr244
		case 80:
			goto tr244
		case 97:
			goto tr245
		case 112:
			goto tr245
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st236
			}
		case data[p] >= 48:
			goto st202
		}
		goto tr0
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		if 48 <= data[p] && data[p] <= 57 {
			goto st203
		}
		goto tr0
tr240:
//line ragel/datetime.rl:7
 pb = p 
	goto st237
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
//line ragel/parse_datetime.go:8370
		switch data[p] {
		case 32:
			goto tr241
		case 58:
			goto tr243
		case 65:
			goto tr244
		case 80:
			goto tr244
		case 97:
			goto tr245
		case 112:
			goto tr245
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st236
		}
		goto tr0
tr234:
//line ragel/datetime.rl:7
 pb = p 
	goto st238
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
//line ragel/parse_datetime.go:8398
		if 48 <= data[p] && data[p] <= 57 {
			goto st186
		}
		goto tr0
tr235:
//line ragel/datetime.rl:7
 pb = p 
	goto st239
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
//line ragel/parse_datetime.go:8412
		if 48 <= data[p] && data[p] <= 49 {
			goto st186
		}
		goto tr0
tr231:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st240
tr334:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st240
tr341:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st240
tr351:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st240
tr811:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st240
tr819:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st240
tr822:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st240
tr829:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st240
tr833:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st240
tr837:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st240
tr846:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st240
tr858:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st240
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
//line ragel/parse_datetime.go:8470
		switch data[p] {
		case 48:
			goto tr321
		case 51:
			goto tr323
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr324
			}
		case data[p] >= 49:
			goto tr322
		}
		goto tr0
tr321:
//line ragel/datetime.rl:7
 pb = p 
	goto st241
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
//line ragel/parse_datetime.go:8495
		if 49 <= data[p] && data[p] <= 57 {
			goto st242
		}
		goto tr0
tr324:
//line ragel/datetime.rl:7
 pb = p 
	goto st242
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
//line ragel/parse_datetime.go:8509
		switch data[p] {
		case 45:
			goto tr326
		case 47:
			goto tr326
		}
		goto tr0
tr668:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st243
tr672:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st243
tr678:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st243
tr686:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st243
tr695:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st243
tr702:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st243
tr704:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st243
tr709:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st243
tr712:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st243
tr715:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st243
tr723:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st243
tr730:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st243
tr326:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st243
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
//line ragel/parse_datetime.go:8581
		if 48 <= data[p] && data[p] <= 57 {
			goto tr327
		}
		goto tr0
tr327:
//line ragel/datetime.rl:7
 pb = p 
	goto st244
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
//line ragel/parse_datetime.go:8595
		if 48 <= data[p] && data[p] <= 57 {
			goto st245
		}
		goto tr0
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		if 48 <= data[p] && data[p] <= 57 {
			goto st246
		}
		goto tr0
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		if 48 <= data[p] && data[p] <= 57 {
			goto st726
		}
		goto tr0
	st726:
		if p++; p == pe {
			goto _test_eof726
		}
	st_case_726:
		switch data[p] {
		case 32:
			goto tr1019
		case 43:
			goto tr1020
		case 45:
			goto tr1021
		case 47:
			goto tr1022
		case 84:
			goto tr1023
		case 90:
			goto tr1024
		case 95:
			goto tr1025
		case 116:
			goto tr1025
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr1022
			}
		case data[p] >= 65:
			goto tr1022
		}
		goto st0
tr322:
//line ragel/datetime.rl:7
 pb = p 
	goto st247
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
//line ragel/parse_datetime.go:8659
		switch data[p] {
		case 45:
			goto tr326
		case 47:
			goto tr326
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st242
		}
		goto tr0
tr323:
//line ragel/datetime.rl:7
 pb = p 
	goto st248
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
//line ragel/parse_datetime.go:8679
		switch data[p] {
		case 45:
			goto tr326
		case 47:
			goto tr326
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st242
		}
		goto tr0
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		if data[p] == 108 {
			goto st250
		}
		goto tr0
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 32:
			goto tr230
		case 45:
			goto tr231
		case 47:
			goto tr231
		}
		goto tr0
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		if data[p] == 103 {
			goto st252
		}
		goto tr0
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		switch data[p] {
		case 32:
			goto tr333
		case 45:
			goto tr334
		case 47:
			goto tr334
		case 117:
			goto st253
		}
		goto tr0
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		if data[p] == 115 {
			goto st254
		}
		goto tr0
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		if data[p] == 116 {
			goto st255
		}
		goto tr0
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 32:
			goto tr333
		case 45:
			goto tr334
		case 47:
			goto tr334
		}
		goto tr0
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		if data[p] == 101 {
			goto st257
		}
		goto tr0
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		if data[p] == 99 {
			goto st258
		}
		goto tr0
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		switch data[p] {
		case 32:
			goto tr340
		case 45:
			goto tr341
		case 47:
			goto tr341
		case 101:
			goto st259
		}
		goto tr0
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		if data[p] == 109 {
			goto st260
		}
		goto tr0
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		if data[p] == 98 {
			goto st261
		}
		goto tr0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		if data[p] == 101 {
			goto st262
		}
		goto tr0
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		if data[p] == 114 {
			goto st263
		}
		goto tr0
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		switch data[p] {
		case 32:
			goto tr340
		case 45:
			goto tr341
		case 47:
			goto tr341
		}
		goto tr0
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 101:
			goto st265
		case 114:
			goto st272
		}
		goto tr0
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		if data[p] == 98 {
			goto st266
		}
		goto tr0
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
		switch data[p] {
		case 32:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr351
		case 114:
			goto st267
		}
		goto tr0
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		if data[p] == 117 {
			goto st268
		}
		goto tr0
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		if data[p] == 97 {
			goto st269
		}
		goto tr0
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		if data[p] == 114 {
			goto st270
		}
		goto tr0
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		if data[p] == 121 {
			goto st271
		}
		goto tr0
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		switch data[p] {
		case 32:
			goto tr350
		case 45:
			goto tr351
		case 47:
			goto tr351
		}
		goto tr0
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		if data[p] == 105 {
			goto st273
		}
		goto tr0
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
		switch data[p] {
		case 32:
			goto st274
		case 44:
			goto st456
		case 100:
			goto st600
		}
		goto tr0
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		switch data[p] {
		case 65:
			goto st275
		case 68:
			goto st398
		case 70:
			goto st406
		case 74:
			goto st414
		case 77:
			goto st426
		case 78:
			goto st432
		case 79:
			goto st440
		case 83:
			goto st447
		}
		goto tr0
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		switch data[p] {
		case 112:
			goto st276
		case 117:
			goto st393
		}
		goto tr0
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		if data[p] == 114 {
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
			goto tr372
		case 105:
			goto st391
		}
		goto tr0
tr372:
//line ragel/datetime.rl:82
 st.Month = 4 
	goto st278
tr584:
//line ragel/datetime.rl:86
 st.Month = 8 
	goto st278
tr590:
//line ragel/datetime.rl:90
 st.Month = 12 
	goto st278
tr598:
//line ragel/datetime.rl:80
 st.Month = 2 
	goto st278
tr607:
//line ragel/datetime.rl:79
 st.Month = 1 
	goto st278
tr614:
//line ragel/datetime.rl:85
 st.Month = 7 
	goto st278
tr616:
//line ragel/datetime.rl:84
 st.Month = 6 
	goto st278
tr621:
//line ragel/datetime.rl:81
 st.Month = 3 
	goto st278
tr624:
//line ragel/datetime.rl:83
 st.Month = 5 
	goto st278
tr627:
//line ragel/datetime.rl:89
 st.Month = 11 
	goto st278
tr635:
//line ragel/datetime.rl:88
 st.Month = 10 
	goto st278
tr642:
//line ragel/datetime.rl:87
 st.Month = 9 
	goto st278
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
//line ragel/parse_datetime.go:9074
		switch data[p] {
		case 32:
			goto st279
		case 48:
			goto tr375
		case 51:
			goto tr377
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr378
			}
		case data[p] >= 49:
			goto tr376
		}
		goto tr0
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		switch data[p] {
		case 48:
			goto tr379
		case 51:
			goto tr381
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr378
			}
		case data[p] >= 49:
			goto tr380
		}
		goto tr0
tr379:
//line ragel/datetime.rl:7
 pb = p 
	goto st280
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
//line ragel/parse_datetime.go:9121
		if 49 <= data[p] && data[p] <= 57 {
			goto st281
		}
		goto tr0
tr378:
//line ragel/datetime.rl:7
 pb = p 
	goto st281
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
//line ragel/parse_datetime.go:9135
		if data[p] == 32 {
			goto tr383
		}
		goto tr0
tr383:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st282
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
//line ragel/parse_datetime.go:9156
		if data[p] == 50 {
			goto tr385
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr386
			}
		case data[p] >= 48:
			goto tr384
		}
		goto tr0
tr384:
//line ragel/datetime.rl:7
 pb = p 
	goto st283
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
//line ragel/parse_datetime.go:9178
		switch data[p] {
		case 32:
			goto tr387
		case 43:
			goto tr388
		case 45:
			goto tr389
		case 47:
			goto tr390
		case 58:
			goto tr392
		case 65:
			goto tr393
		case 80:
			goto tr393
		case 90:
			goto tr394
		case 95:
			goto tr390
		case 97:
			goto tr395
		case 112:
			goto tr395
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st306
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr390
			}
		default:
			goto tr390
		}
		goto tr0
tr387:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st284
tr430:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st284
tr493:
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

	goto st284
tr464:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st284
tr473:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st284
tr483:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st284
tr504:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st284
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
//line ragel/parse_datetime.go:9288
		switch data[p] {
		case 32:
			goto st285
		case 43:
			goto st289
		case 45:
			goto st297
		case 47:
			goto tr399
		case 65:
			goto tr401
		case 80:
			goto tr401
		case 90:
			goto tr402
		case 95:
			goto tr399
		case 97:
			goto tr403
		case 112:
			goto tr403
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr400
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr399
			}
		default:
			goto tr399
		}
		goto tr0
tr410:
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
	goto st285
tr414:
//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st285
tr417:
//line ragel/datetime.rl:159

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st285
tr424:
//line ragel/datetime.rl:196

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:9
 st.Zoned = true 
	goto st285
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
//line ragel/parse_datetime.go:9387
		if 48 <= data[p] && data[p] <= 57 {
			goto tr400
		}
		goto tr0
tr400:
//line ragel/datetime.rl:7
 pb = p 
	goto st286
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
//line ragel/parse_datetime.go:9401
		if 48 <= data[p] && data[p] <= 57 {
			goto st287
		}
		goto tr0
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		if 48 <= data[p] && data[p] <= 57 {
			goto st288
		}
		goto tr0
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		if 48 <= data[p] && data[p] <= 57 {
			goto st727
		}
		goto tr0
	st727:
		if p++; p == pe {
			goto _test_eof727
		}
	st_case_727:
		goto st0
tr388:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st289
tr427:
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

	goto st289
tr431:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st289
tr441:
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

	goto st289
tr449:
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

	goto st289
tr465:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st289
tr474:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st289
tr484:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st289
tr505:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st289
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
//line ragel/parse_datetime.go:9544
		if data[p] == 50 {
			goto tr408
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr409
			}
		case data[p] >= 48:
			goto tr407
		}
		goto tr0
tr407:
//line ragel/datetime.rl:7
 pb = p 
	goto st290
tr419:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st290
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
//line ragel/parse_datetime.go:9572
		switch data[p] {
		case 32:
			goto tr410
		case 58:
			goto tr412
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st291
		}
		goto tr0
tr409:
//line ragel/datetime.rl:7
 pb = p 
	goto st291
tr421:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st291
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
//line ragel/parse_datetime.go:9598
		switch data[p] {
		case 32:
			goto tr410
		case 58:
			goto tr412
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
		if data[p] == 32 {
			goto tr410
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st292
		}
		goto tr0
tr412:
//line ragel/datetime.rl:150

    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }

	goto st293
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
//line ragel/parse_datetime.go:9637
		if data[p] == 32 {
			goto tr414
		}
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr416
			}
		case data[p] >= 48:
			goto tr415
		}
		goto tr0
tr415:
//line ragel/datetime.rl:7
 pb = p 
	goto st294
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
//line ragel/parse_datetime.go:9659
		if data[p] == 32 {
			goto tr417
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st295
		}
		goto tr0
tr416:
//line ragel/datetime.rl:7
 pb = p 
	goto st295
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
//line ragel/parse_datetime.go:9676
		if data[p] == 32 {
			goto tr417
		}
		goto tr0
tr408:
//line ragel/datetime.rl:7
 pb = p 
	goto st296
tr420:
//line ragel/datetime.rl:148
 st.NegtiveZoneOffset = true 
//line ragel/datetime.rl:7
 pb = p 
	goto st296
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
//line ragel/parse_datetime.go:9696
		switch data[p] {
		case 32:
			goto tr410
		case 58:
			goto tr412
		}
		switch {
		case data[p] > 51:
			if 52 <= data[p] && data[p] <= 57 {
				goto st292
			}
		case data[p] >= 48:
			goto st291
		}
		goto tr0
tr389:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st297
tr428:
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

	goto st297
tr432:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st297
tr442:
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

	goto st297
tr450:
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

	goto st297
tr466:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st297
tr475:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st297
tr485:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st297
tr506:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st297
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
//line ragel/parse_datetime.go:9826
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
		goto tr0
tr399:
//line ragel/datetime.rl:7
 pb = p 
	goto st298
tr390:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st298
tr433:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st298
tr467:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st298
tr476:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st298
tr487:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st298
tr451:
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
	goto st298
tr508:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st298
tr444:
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
	goto st298
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
//line ragel/parse_datetime.go:9948
		switch data[p] {
		case 47:
			goto st299
		case 95:
			goto st299
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st299
			}
		case data[p] >= 65:
			goto st299
		}
		goto tr0
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 47:
			goto st300
		case 95:
			goto st300
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st300
			}
		case data[p] >= 65:
			goto st300
		}
		goto tr0
tr429:
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
	goto st300
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
//line ragel/parse_datetime.go:10016
		switch data[p] {
		case 32:
			goto tr424
		case 47:
			goto st300
		case 95:
			goto st300
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st300
			}
		case data[p] >= 65:
			goto st300
		}
		goto tr0
tr401:
//line ragel/datetime.rl:7
 pb = p 
	goto st301
tr393:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st301
tr436:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st301
tr470:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st301
tr478:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st301
tr489:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st301
tr495:
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
	goto st301
tr509:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st301
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
//line ragel/parse_datetime.go:10124
		switch data[p] {
		case 47:
			goto st299
		case 77:
			goto st302
		case 95:
			goto st299
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st299
			}
		case data[p] >= 65:
			goto st299
		}
		goto tr0
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		switch data[p] {
		case 32:
			goto tr426
		case 43:
			goto tr427
		case 45:
			goto tr428
		case 47:
			goto tr429
		case 95:
			goto tr429
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr429
			}
		case data[p] >= 65:
			goto tr429
		}
		goto tr0
tr426:
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

	goto st303
tr440:
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

	goto st303
tr448:
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

	goto st303
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
//line ragel/parse_datetime.go:10246
		switch data[p] {
		case 32:
			goto st285
		case 43:
			goto st289
		case 45:
			goto st297
		case 47:
			goto tr399
		case 90:
			goto tr402
		case 95:
			goto tr399
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr400
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr399
			}
		default:
			goto tr399
		}
		goto tr0
tr402:
//line ragel/datetime.rl:7
 pb = p 
	goto st304
tr394:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st304
tr437:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st304
tr471:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st304
tr479:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st304
tr490:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st304
tr453:
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
	goto st304
tr510:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st304
tr446:
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
	goto st304
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
//line ragel/parse_datetime.go:10383
		switch data[p] {
		case 32:
			goto tr414
		case 47:
			goto st299
		case 95:
			goto st299
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st299
			}
		case data[p] >= 65:
			goto st299
		}
		goto tr0
tr403:
//line ragel/datetime.rl:7
 pb = p 
	goto st305
tr395:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st305
tr438:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st305
tr472:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st305
tr480:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st305
tr491:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st305
tr496:
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
	goto st305
tr511:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st305
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
//line ragel/parse_datetime.go:10491
		switch data[p] {
		case 47:
			goto st299
		case 95:
			goto st299
		case 109:
			goto st302
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st299
			}
		case data[p] >= 65:
			goto st299
		}
		goto tr0
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		switch data[p] {
		case 32:
			goto tr430
		case 43:
			goto tr431
		case 45:
			goto tr432
		case 47:
			goto tr433
		case 58:
			goto tr435
		case 65:
			goto tr436
		case 80:
			goto tr436
		case 90:
			goto tr437
		case 95:
			goto tr433
		case 97:
			goto tr438
		case 112:
			goto tr438
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st307
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr433
			}
		default:
			goto tr433
		}
		goto tr0
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		if 48 <= data[p] && data[p] <= 57 {
			goto st308
		}
		goto tr0
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		switch data[p] {
		case 32:
			goto tr440
		case 43:
			goto tr441
		case 45:
			goto tr442
		case 46:
			goto tr443
		case 47:
			goto tr444
		case 90:
			goto tr446
		case 95:
			goto tr444
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st319
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr444
			}
		default:
			goto tr444
		}
		goto tr0
tr443:
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

	goto st309
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
//line ragel/parse_datetime.go:10616
		if 48 <= data[p] && data[p] <= 57 {
			goto tr447
		}
		goto tr0
tr447:
//line ragel/datetime.rl:7
 pb = p 
	goto st310
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
//line ragel/parse_datetime.go:10630
		switch data[p] {
		case 32:
			goto tr448
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st311
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
		switch data[p] {
		case 32:
			goto tr448
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st312
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		switch data[p] {
		case 32:
			goto tr448
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st313
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 32:
			goto tr448
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st314
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		switch data[p] {
		case 32:
			goto tr448
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st315
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		switch data[p] {
		case 32:
			goto tr448
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st316
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		switch data[p] {
		case 32:
			goto tr448
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st317
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		switch data[p] {
		case 32:
			goto tr448
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st318
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		switch data[p] {
		case 32:
			goto tr448
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		case data[p] >= 65:
			goto tr451
		}
		goto tr0
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		if 48 <= data[p] && data[p] <= 57 {
			goto st320
		}
		goto tr0
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		switch data[p] {
		case 32:
			goto tr440
		case 43:
			goto tr441
		case 45:
			goto tr442
		case 46:
			goto tr443
		case 47:
			goto tr444
		case 90:
			goto tr446
		case 95:
			goto tr444
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr444
			}
		case data[p] >= 65:
			goto tr444
		}
		goto tr0
tr392:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st321
tr435:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st321
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
//line ragel/parse_datetime.go:10966
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr463
			}
		case data[p] >= 48:
			goto tr462
		}
		goto tr0
tr462:
//line ragel/datetime.rl:7
 pb = p 
	goto st322
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
//line ragel/parse_datetime.go:10985
		switch data[p] {
		case 32:
			goto tr464
		case 43:
			goto tr465
		case 45:
			goto tr466
		case 47:
			goto tr467
		case 58:
			goto tr469
		case 65:
			goto tr470
		case 80:
			goto tr470
		case 90:
			goto tr471
		case 95:
			goto tr467
		case 97:
			goto tr472
		case 112:
			goto tr472
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st323
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr467
			}
		default:
			goto tr467
		}
		goto tr0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 32:
			goto tr473
		case 43:
			goto tr474
		case 45:
			goto tr475
		case 47:
			goto tr476
		case 58:
			goto tr477
		case 65:
			goto tr478
		case 80:
			goto tr478
		case 90:
			goto tr479
		case 95:
			goto tr476
		case 97:
			goto tr480
		case 112:
			goto tr480
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr476
			}
		case data[p] >= 66:
			goto tr476
		}
		goto tr0
tr469:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st324
tr477:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st324
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
//line ragel/parse_datetime.go:11078
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr482
			}
		case data[p] >= 48:
			goto tr481
		}
		goto tr0
tr481:
//line ragel/datetime.rl:7
 pb = p 
	goto st325
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
//line ragel/parse_datetime.go:11097
		switch data[p] {
		case 32:
			goto tr483
		case 43:
			goto tr484
		case 45:
			goto tr485
		case 46:
			goto tr486
		case 47:
			goto tr487
		case 65:
			goto tr489
		case 80:
			goto tr489
		case 90:
			goto tr490
		case 95:
			goto tr487
		case 97:
			goto tr491
		case 112:
			goto tr491
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st336
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr487
			}
		default:
			goto tr487
		}
		goto tr0
tr486:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st326
tr507:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st326
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
//line ragel/parse_datetime.go:11152
		if 48 <= data[p] && data[p] <= 57 {
			goto tr492
		}
		goto tr0
tr492:
//line ragel/datetime.rl:7
 pb = p 
	goto st327
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
//line ragel/parse_datetime.go:11166
		switch data[p] {
		case 32:
			goto tr493
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr495
		case 80:
			goto tr495
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr496
		case 112:
			goto tr496
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st328
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		switch data[p] {
		case 32:
			goto tr493
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr495
		case 80:
			goto tr495
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr496
		case 112:
			goto tr496
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st329
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		switch data[p] {
		case 32:
			goto tr493
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr495
		case 80:
			goto tr495
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr496
		case 112:
			goto tr496
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st330
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		switch data[p] {
		case 32:
			goto tr493
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr495
		case 80:
			goto tr495
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr496
		case 112:
			goto tr496
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st331
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		switch data[p] {
		case 32:
			goto tr493
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr495
		case 80:
			goto tr495
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr496
		case 112:
			goto tr496
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st332
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		switch data[p] {
		case 32:
			goto tr493
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr495
		case 80:
			goto tr495
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr496
		case 112:
			goto tr496
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st333
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 32:
			goto tr493
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr495
		case 80:
			goto tr495
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr496
		case 112:
			goto tr496
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st334
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		switch data[p] {
		case 32:
			goto tr493
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr495
		case 80:
			goto tr495
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr496
		case 112:
			goto tr496
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st335
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		switch data[p] {
		case 32:
			goto tr493
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr495
		case 80:
			goto tr495
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr496
		case 112:
			goto tr496
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		case data[p] >= 66:
			goto tr451
		}
		goto tr0
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		switch data[p] {
		case 32:
			goto tr504
		case 43:
			goto tr505
		case 45:
			goto tr506
		case 46:
			goto tr507
		case 47:
			goto tr508
		case 65:
			goto tr509
		case 80:
			goto tr509
		case 90:
			goto tr510
		case 95:
			goto tr508
		case 97:
			goto tr511
		case 112:
			goto tr511
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr508
			}
		case data[p] >= 66:
			goto tr508
		}
		goto tr0
tr482:
//line ragel/datetime.rl:7
 pb = p 
	goto st337
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
//line ragel/parse_datetime.go:11565
		switch data[p] {
		case 32:
			goto tr483
		case 43:
			goto tr484
		case 45:
			goto tr485
		case 46:
			goto tr486
		case 47:
			goto tr487
		case 65:
			goto tr489
		case 80:
			goto tr489
		case 90:
			goto tr490
		case 95:
			goto tr487
		case 97:
			goto tr491
		case 112:
			goto tr491
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr487
			}
		case data[p] >= 66:
			goto tr487
		}
		goto tr0
tr463:
//line ragel/datetime.rl:7
 pb = p 
	goto st338
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
//line ragel/parse_datetime.go:11608
		switch data[p] {
		case 32:
			goto tr464
		case 43:
			goto tr465
		case 45:
			goto tr466
		case 47:
			goto tr467
		case 58:
			goto tr469
		case 65:
			goto tr470
		case 80:
			goto tr470
		case 90:
			goto tr471
		case 95:
			goto tr467
		case 97:
			goto tr472
		case 112:
			goto tr472
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr467
			}
		case data[p] >= 66:
			goto tr467
		}
		goto tr0
tr385:
//line ragel/datetime.rl:7
 pb = p 
	goto st339
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
//line ragel/parse_datetime.go:11651
		switch data[p] {
		case 32:
			goto tr387
		case 43:
			goto tr388
		case 45:
			goto tr389
		case 47:
			goto tr390
		case 58:
			goto tr392
		case 65:
			goto tr393
		case 80:
			goto tr393
		case 90:
			goto tr394
		case 95:
			goto tr390
		case 97:
			goto tr395
		case 112:
			goto tr395
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st306
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr390
				}
			case data[p] >= 66:
				goto tr390
			}
		default:
			goto st340
		}
		goto tr0
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		if 48 <= data[p] && data[p] <= 57 {
			goto st307
		}
		goto tr0
tr386:
//line ragel/datetime.rl:7
 pb = p 
	goto st341
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
//line ragel/parse_datetime.go:11712
		switch data[p] {
		case 32:
			goto tr387
		case 43:
			goto tr388
		case 45:
			goto tr389
		case 47:
			goto tr390
		case 58:
			goto tr392
		case 65:
			goto tr393
		case 80:
			goto tr393
		case 90:
			goto tr394
		case 95:
			goto tr390
		case 97:
			goto tr395
		case 112:
			goto tr395
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st340
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr390
			}
		default:
			goto tr390
		}
		goto tr0
tr380:
//line ragel/datetime.rl:7
 pb = p 
	goto st342
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
//line ragel/parse_datetime.go:11759
		if data[p] == 32 {
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st281
		}
		goto tr0
tr381:
//line ragel/datetime.rl:7
 pb = p 
	goto st343
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
//line ragel/parse_datetime.go:11776
		if data[p] == 32 {
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st281
		}
		goto tr0
tr375:
//line ragel/datetime.rl:7
 pb = p 
	goto st344
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
//line ragel/parse_datetime.go:11793
		if 49 <= data[p] && data[p] <= 57 {
			goto st345
		}
		goto tr0
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		if data[p] == 32 {
			goto tr514
		}
		goto tr0
tr514:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st346
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
//line ragel/parse_datetime.go:11823
		if data[p] == 50 {
			goto tr516
		}
		switch {
		case data[p] > 49:
			if 51 <= data[p] && data[p] <= 57 {
				goto tr517
			}
		case data[p] >= 48:
			goto tr515
		}
		goto tr0
tr515:
//line ragel/datetime.rl:7
 pb = p 
	goto st347
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
//line ragel/parse_datetime.go:11845
		switch data[p] {
		case 32:
			goto tr518
		case 43:
			goto tr388
		case 45:
			goto tr389
		case 47:
			goto tr390
		case 58:
			goto tr520
		case 65:
			goto tr521
		case 80:
			goto tr521
		case 90:
			goto tr394
		case 95:
			goto tr390
		case 97:
			goto tr522
		case 112:
			goto tr522
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st353
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr390
			}
		default:
			goto tr390
		}
		goto tr0
tr518:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st348
tr527:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st348
tr566:
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

	goto st348
tr549:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st348
tr554:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st348
tr560:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st348
tr577:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st348
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
//line ragel/parse_datetime.go:11955
		switch data[p] {
		case 32:
			goto st285
		case 43:
			goto st289
		case 45:
			goto st297
		case 47:
			goto tr399
		case 65:
			goto tr523
		case 80:
			goto tr523
		case 90:
			goto tr402
		case 95:
			goto tr399
		case 97:
			goto tr524
		case 112:
			goto tr524
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr246
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr399
			}
		default:
			goto tr399
		}
		goto tr0
tr523:
//line ragel/datetime.rl:7
 pb = p 
	goto st349
tr521:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st349
tr530:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st349
tr552:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st349
tr556:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st349
tr563:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st349
tr568:
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
	goto st349
tr579:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st349
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
//line ragel/parse_datetime.go:12081
		switch data[p] {
		case 47:
			goto st299
		case 77:
			goto st350
		case 95:
			goto st299
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st299
			}
		case data[p] >= 65:
			goto st299
		}
		goto tr0
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		switch data[p] {
		case 32:
			goto tr526
		case 43:
			goto tr427
		case 45:
			goto tr428
		case 47:
			goto tr429
		case 95:
			goto tr429
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr429
			}
		case data[p] >= 65:
			goto tr429
		}
		goto tr0
tr526:
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

	goto st351
tr533:
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

	goto st351
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

	goto st351
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
//line ragel/parse_datetime.go:12203
		switch data[p] {
		case 32:
			goto st285
		case 43:
			goto st289
		case 45:
			goto st297
		case 47:
			goto tr399
		case 90:
			goto tr402
		case 95:
			goto tr399
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr246
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr399
			}
		default:
			goto tr399
		}
		goto tr0
tr524:
//line ragel/datetime.rl:7
 pb = p 
	goto st352
tr522:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st352
tr531:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st352
tr553:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st352
tr557:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st352
tr564:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

//line ragel/datetime.rl:7
 pb = p 
	goto st352
tr569:
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
	goto st352
tr580:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

//line ragel/datetime.rl:7
 pb = p 
	goto st352
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
//line ragel/parse_datetime.go:12321
		switch data[p] {
		case 47:
			goto st299
		case 95:
			goto st299
		case 109:
			goto st350
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st299
			}
		case data[p] >= 65:
			goto st299
		}
		goto tr0
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		switch data[p] {
		case 32:
			goto tr527
		case 43:
			goto tr431
		case 45:
			goto tr432
		case 47:
			goto tr433
		case 58:
			goto tr529
		case 65:
			goto tr530
		case 80:
			goto tr530
		case 90:
			goto tr437
		case 95:
			goto tr433
		case 97:
			goto tr531
		case 112:
			goto tr531
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st354
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr433
			}
		default:
			goto tr433
		}
		goto tr0
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		if 48 <= data[p] && data[p] <= 57 {
			goto st355
		}
		goto tr0
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		switch data[p] {
		case 32:
			goto tr533
		case 43:
			goto tr441
		case 45:
			goto tr442
		case 46:
			goto tr534
		case 47:
			goto tr444
		case 90:
			goto tr446
		case 95:
			goto tr444
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st366
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr444
			}
		default:
			goto tr444
		}
		goto tr0
tr534:
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

	goto st356
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
//line ragel/parse_datetime.go:12446
		if 48 <= data[p] && data[p] <= 57 {
			goto tr536
		}
		goto tr0
tr536:
//line ragel/datetime.rl:7
 pb = p 
	goto st357
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
//line ragel/parse_datetime.go:12460
		switch data[p] {
		case 32:
			goto tr537
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st358
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		switch data[p] {
		case 32:
			goto tr537
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st359
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		switch data[p] {
		case 32:
			goto tr537
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st360
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		switch data[p] {
		case 32:
			goto tr537
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st361
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		switch data[p] {
		case 32:
			goto tr537
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st362
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		switch data[p] {
		case 32:
			goto tr537
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st363
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		switch data[p] {
		case 32:
			goto tr537
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st364
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		switch data[p] {
		case 32:
			goto tr537
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st365
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		switch data[p] {
		case 32:
			goto tr537
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 90:
			goto tr453
		case 95:
			goto tr451
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		case data[p] >= 65:
			goto tr451
		}
		goto tr0
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		if 48 <= data[p] && data[p] <= 57 {
			goto st367
		}
		goto tr0
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		switch data[p] {
		case 32:
			goto tr533
		case 43:
			goto tr441
		case 45:
			goto tr442
		case 46:
			goto tr534
		case 47:
			goto tr444
		case 90:
			goto tr446
		case 95:
			goto tr444
		}
		switch {
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr444
			}
		case data[p] >= 65:
			goto tr444
		}
		goto tr0
tr520:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

	goto st368
tr529:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

	goto st368
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
//line ragel/parse_datetime.go:12796
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr548
			}
		case data[p] >= 48:
			goto tr547
		}
		goto tr0
tr547:
//line ragel/datetime.rl:7
 pb = p 
	goto st369
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
//line ragel/parse_datetime.go:12815
		switch data[p] {
		case 32:
			goto tr549
		case 43:
			goto tr465
		case 45:
			goto tr466
		case 47:
			goto tr467
		case 58:
			goto tr551
		case 65:
			goto tr552
		case 80:
			goto tr552
		case 90:
			goto tr471
		case 95:
			goto tr467
		case 97:
			goto tr553
		case 112:
			goto tr553
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st370
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr467
			}
		default:
			goto tr467
		}
		goto tr0
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		switch data[p] {
		case 32:
			goto tr554
		case 43:
			goto tr474
		case 45:
			goto tr475
		case 47:
			goto tr476
		case 58:
			goto tr555
		case 65:
			goto tr556
		case 80:
			goto tr556
		case 90:
			goto tr479
		case 95:
			goto tr476
		case 97:
			goto tr557
		case 112:
			goto tr557
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr476
			}
		case data[p] >= 66:
			goto tr476
		}
		goto tr0
tr551:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

	goto st371
tr555:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

	goto st371
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
//line ragel/parse_datetime.go:12908
		switch {
		case data[p] > 53:
			if 54 <= data[p] && data[p] <= 57 {
				goto tr559
			}
		case data[p] >= 48:
			goto tr558
		}
		goto tr0
tr558:
//line ragel/datetime.rl:7
 pb = p 
	goto st372
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
//line ragel/parse_datetime.go:12927
		switch data[p] {
		case 32:
			goto tr560
		case 43:
			goto tr484
		case 45:
			goto tr485
		case 46:
			goto tr561
		case 47:
			goto tr487
		case 65:
			goto tr563
		case 80:
			goto tr563
		case 90:
			goto tr490
		case 95:
			goto tr487
		case 97:
			goto tr564
		case 112:
			goto tr564
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st383
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr487
			}
		default:
			goto tr487
		}
		goto tr0
tr561:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

	goto st373
tr578:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

	goto st373
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
//line ragel/parse_datetime.go:12982
		if 48 <= data[p] && data[p] <= 57 {
			goto tr565
		}
		goto tr0
tr565:
//line ragel/datetime.rl:7
 pb = p 
	goto st374
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
//line ragel/parse_datetime.go:12996
		switch data[p] {
		case 32:
			goto tr566
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr568
		case 80:
			goto tr568
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr569
		case 112:
			goto tr569
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st375
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		switch data[p] {
		case 32:
			goto tr566
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr568
		case 80:
			goto tr568
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr569
		case 112:
			goto tr569
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st376
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		switch data[p] {
		case 32:
			goto tr566
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr568
		case 80:
			goto tr568
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr569
		case 112:
			goto tr569
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st377
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		switch data[p] {
		case 32:
			goto tr566
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr568
		case 80:
			goto tr568
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr569
		case 112:
			goto tr569
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st378
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		switch data[p] {
		case 32:
			goto tr566
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr568
		case 80:
			goto tr568
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr569
		case 112:
			goto tr569
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st379
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		switch data[p] {
		case 32:
			goto tr566
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr568
		case 80:
			goto tr568
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr569
		case 112:
			goto tr569
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st380
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		switch data[p] {
		case 32:
			goto tr566
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr568
		case 80:
			goto tr568
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr569
		case 112:
			goto tr569
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st381
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		switch data[p] {
		case 32:
			goto tr566
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr568
		case 80:
			goto tr568
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr569
		case 112:
			goto tr569
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st382
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto tr0
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		switch data[p] {
		case 32:
			goto tr566
		case 43:
			goto tr449
		case 45:
			goto tr450
		case 47:
			goto tr451
		case 65:
			goto tr568
		case 80:
			goto tr568
		case 90:
			goto tr453
		case 95:
			goto tr451
		case 97:
			goto tr569
		case 112:
			goto tr569
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr451
			}
		case data[p] >= 66:
			goto tr451
		}
		goto tr0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 32:
			goto tr577
		case 43:
			goto tr505
		case 45:
			goto tr506
		case 46:
			goto tr578
		case 47:
			goto tr508
		case 65:
			goto tr579
		case 80:
			goto tr579
		case 90:
			goto tr510
		case 95:
			goto tr508
		case 97:
			goto tr580
		case 112:
			goto tr580
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr508
			}
		case data[p] >= 66:
			goto tr508
		}
		goto tr0
tr559:
//line ragel/datetime.rl:7
 pb = p 
	goto st384
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
//line ragel/parse_datetime.go:13395
		switch data[p] {
		case 32:
			goto tr560
		case 43:
			goto tr484
		case 45:
			goto tr485
		case 46:
			goto tr561
		case 47:
			goto tr487
		case 65:
			goto tr563
		case 80:
			goto tr563
		case 90:
			goto tr490
		case 95:
			goto tr487
		case 97:
			goto tr564
		case 112:
			goto tr564
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr487
			}
		case data[p] >= 66:
			goto tr487
		}
		goto tr0
tr548:
//line ragel/datetime.rl:7
 pb = p 
	goto st385
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
//line ragel/parse_datetime.go:13438
		switch data[p] {
		case 32:
			goto tr549
		case 43:
			goto tr465
		case 45:
			goto tr466
		case 47:
			goto tr467
		case 58:
			goto tr551
		case 65:
			goto tr552
		case 80:
			goto tr552
		case 90:
			goto tr471
		case 95:
			goto tr467
		case 97:
			goto tr553
		case 112:
			goto tr553
		}
		switch {
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr467
			}
		case data[p] >= 66:
			goto tr467
		}
		goto tr0
tr516:
//line ragel/datetime.rl:7
 pb = p 
	goto st386
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
//line ragel/parse_datetime.go:13481
		switch data[p] {
		case 32:
			goto tr518
		case 43:
			goto tr388
		case 45:
			goto tr389
		case 47:
			goto tr390
		case 58:
			goto tr520
		case 65:
			goto tr521
		case 80:
			goto tr521
		case 90:
			goto tr394
		case 95:
			goto tr390
		case 97:
			goto tr522
		case 112:
			goto tr522
		}
		switch {
		case data[p] < 52:
			if 48 <= data[p] && data[p] <= 51 {
				goto st353
			}
		case data[p] > 57:
			switch {
			case data[p] > 89:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr390
				}
			case data[p] >= 66:
				goto tr390
			}
		default:
			goto st387
		}
		goto tr0
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		if 48 <= data[p] && data[p] <= 57 {
			goto st354
		}
		goto tr0
tr517:
//line ragel/datetime.rl:7
 pb = p 
	goto st388
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
//line ragel/parse_datetime.go:13542
		switch data[p] {
		case 32:
			goto tr518
		case 43:
			goto tr388
		case 45:
			goto tr389
		case 47:
			goto tr390
		case 58:
			goto tr520
		case 65:
			goto tr521
		case 80:
			goto tr521
		case 90:
			goto tr394
		case 95:
			goto tr390
		case 97:
			goto tr522
		case 112:
			goto tr522
		}
		switch {
		case data[p] < 66:
			if 48 <= data[p] && data[p] <= 57 {
				goto st387
			}
		case data[p] > 89:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr390
			}
		default:
			goto tr390
		}
		goto tr0
tr376:
//line ragel/datetime.rl:7
 pb = p 
	goto st389
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
//line ragel/parse_datetime.go:13589
		if data[p] == 32 {
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st345
		}
		goto tr0
tr377:
//line ragel/datetime.rl:7
 pb = p 
	goto st390
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
//line ragel/parse_datetime.go:13606
		if data[p] == 32 {
			goto tr383
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st345
		}
		goto tr0
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
		if data[p] == 108 {
			goto st392
		}
		goto tr0
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		if data[p] == 32 {
			goto tr372
		}
		goto tr0
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		if data[p] == 103 {
			goto st394
		}
		goto tr0
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		switch data[p] {
		case 32:
			goto tr584
		case 117:
			goto st395
		}
		goto tr0
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		if data[p] == 115 {
			goto st396
		}
		goto tr0
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		if data[p] == 116 {
			goto st397
		}
		goto tr0
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		if data[p] == 32 {
			goto tr584
		}
		goto tr0
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		if data[p] == 101 {
			goto st399
		}
		goto tr0
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		if data[p] == 99 {
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
			goto tr590
		case 101:
			goto st401
		}
		goto tr0
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		if data[p] == 109 {
			goto st402
		}
		goto tr0
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
		if data[p] == 98 {
			goto st403
		}
		goto tr0
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		if data[p] == 101 {
			goto st404
		}
		goto tr0
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		if data[p] == 114 {
			goto st405
		}
		goto tr0
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		if data[p] == 32 {
			goto tr590
		}
		goto tr0
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		if data[p] == 101 {
			goto st407
		}
		goto tr0
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		if data[p] == 98 {
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
			goto tr598
		case 114:
			goto st409
		}
		goto tr0
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		if data[p] == 117 {
			goto st410
		}
		goto tr0
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		if data[p] == 97 {
			goto st411
		}
		goto tr0
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		if data[p] == 114 {
			goto st412
		}
		goto tr0
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		if data[p] == 121 {
			goto st413
		}
		goto tr0
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
		if data[p] == 32 {
			goto tr598
		}
		goto tr0
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		switch data[p] {
		case 97:
			goto st415
		case 117:
			goto st421
		}
		goto tr0
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		if data[p] == 110 {
			goto st416
		}
		goto tr0
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		switch data[p] {
		case 32:
			goto tr607
		case 117:
			goto st417
		}
		goto tr0
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		if data[p] == 97 {
			goto st418
		}
		goto tr0
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		if data[p] == 114 {
			goto st419
		}
		goto tr0
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		if data[p] == 121 {
			goto st420
		}
		goto tr0
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
		if data[p] == 32 {
			goto tr607
		}
		goto tr0
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		switch data[p] {
		case 108:
			goto st422
		case 110:
			goto st424
		}
		goto tr0
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		switch data[p] {
		case 32:
			goto tr614
		case 121:
			goto st423
		}
		goto tr0
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		if data[p] == 32 {
			goto tr614
		}
		goto tr0
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 32:
			goto tr616
		case 101:
			goto st425
		}
		goto tr0
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
		if data[p] == 32 {
			goto tr616
		}
		goto tr0
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
		if data[p] == 97 {
			goto st427
		}
		goto tr0
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 114:
			goto st428
		case 121:
			goto st431
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
		case 99:
			goto st429
		}
		goto tr0
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
		if data[p] == 104 {
			goto st430
		}
		goto tr0
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
		if data[p] == 32 {
			goto tr621
		}
		goto tr0
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		if data[p] == 32 {
			goto tr624
		}
		goto tr0
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
		if data[p] == 111 {
			goto st433
		}
		goto tr0
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		if data[p] == 118 {
			goto st434
		}
		goto tr0
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		switch data[p] {
		case 32:
			goto tr627
		case 101:
			goto st435
		}
		goto tr0
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		if data[p] == 109 {
			goto st436
		}
		goto tr0
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		if data[p] == 98 {
			goto st437
		}
		goto tr0
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		if data[p] == 101 {
			goto st438
		}
		goto tr0
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		if data[p] == 114 {
			goto st439
		}
		goto tr0
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		if data[p] == 32 {
			goto tr627
		}
		goto tr0
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
		if data[p] == 99 {
			goto st441
		}
		goto tr0
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		if data[p] == 116 {
			goto st442
		}
		goto tr0
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		switch data[p] {
		case 32:
			goto tr635
		case 111:
			goto st443
		}
		goto tr0
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
		if data[p] == 98 {
			goto st444
		}
		goto tr0
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		if data[p] == 101 {
			goto st445
		}
		goto tr0
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
		if data[p] == 114 {
			goto st446
		}
		goto tr0
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		if data[p] == 32 {
			goto tr635
		}
		goto tr0
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
		if data[p] == 101 {
			goto st448
		}
		goto tr0
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		if data[p] == 112 {
			goto st449
		}
		goto tr0
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		switch data[p] {
		case 32:
			goto tr642
		case 116:
			goto st450
		}
		goto tr0
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
		if data[p] == 101 {
			goto st451
		}
		goto tr0
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		if data[p] == 109 {
			goto st452
		}
		goto tr0
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		if data[p] == 98 {
			goto st453
		}
		goto tr0
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		if data[p] == 101 {
			goto st454
		}
		goto tr0
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		if data[p] == 114 {
			goto st455
		}
		goto tr0
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
		if data[p] == 32 {
			goto tr642
		}
		goto tr0
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		if data[p] == 32 {
			goto st457
		}
		goto tr0
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 48:
			goto tr650
		case 51:
			goto tr652
		}
		switch {
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto tr653
			}
		case data[p] >= 49:
			goto tr651
		}
		goto tr0
tr650:
//line ragel/datetime.rl:7
 pb = p 
	goto st458
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
//line ragel/parse_datetime.go:14276
		if 49 <= data[p] && data[p] <= 57 {
			goto st459
		}
		goto tr0
tr653:
//line ragel/datetime.rl:7
 pb = p 
	goto st459
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
//line ragel/parse_datetime.go:14290
		switch data[p] {
		case 32:
			goto tr655
		case 45:
			goto tr656
		}
		goto tr0
tr655:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

	goto st460
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
//line ragel/parse_datetime.go:14314
		switch data[p] {
		case 65:
			goto st461
		case 68:
			goto st471
		case 70:
			goto st479
		case 74:
			goto st487
		case 77:
			goto st499
		case 78:
			goto st505
		case 79:
			goto st513
		case 83:
			goto st520
		}
		goto tr0
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		switch data[p] {
		case 112:
			goto st462
		case 117:
			goto st466
		}
		goto tr0
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
		if data[p] == 114 {
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
		case 105:
			goto st464
		}
		goto tr0
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
		if data[p] == 108 {
			goto st465
		}
		goto tr0
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
		if data[p] == 32 {
			goto tr668
		}
		goto tr0
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		if data[p] == 103 {
			goto st467
		}
		goto tr0
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
		switch data[p] {
		case 32:
			goto tr672
		case 117:
			goto st468
		}
		goto tr0
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
		if data[p] == 115 {
			goto st469
		}
		goto tr0
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		if data[p] == 116 {
			goto st470
		}
		goto tr0
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		if data[p] == 32 {
			goto tr672
		}
		goto tr0
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		if data[p] == 101 {
			goto st472
		}
		goto tr0
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
		if data[p] == 99 {
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
			goto tr678
		case 101:
			goto st474
		}
		goto tr0
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		if data[p] == 109 {
			goto st475
		}
		goto tr0
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		if data[p] == 98 {
			goto st476
		}
		goto tr0
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
		if data[p] == 101 {
			goto st477
		}
		goto tr0
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		if data[p] == 114 {
			goto st478
		}
		goto tr0
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		if data[p] == 32 {
			goto tr678
		}
		goto tr0
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		if data[p] == 101 {
			goto st480
		}
		goto tr0
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		if data[p] == 98 {
			goto st481
		}
		goto tr0
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		switch data[p] {
		case 32:
			goto tr686
		case 114:
			goto st482
		}
		goto tr0
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
		if data[p] == 117 {
			goto st483
		}
		goto tr0
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
		if data[p] == 97 {
			goto st484
		}
		goto tr0
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
		if data[p] == 114 {
			goto st485
		}
		goto tr0
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
		if data[p] == 121 {
			goto st486
		}
		goto tr0
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		if data[p] == 32 {
			goto tr686
		}
		goto tr0
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
		switch data[p] {
		case 97:
			goto st488
		case 117:
			goto st494
		}
		goto tr0
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
		if data[p] == 110 {
			goto st489
		}
		goto tr0
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
		switch data[p] {
		case 32:
			goto tr695
		case 117:
			goto st490
		}
		goto tr0
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		if data[p] == 97 {
			goto st491
		}
		goto tr0
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		if data[p] == 114 {
			goto st492
		}
		goto tr0
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
		if data[p] == 121 {
			goto st493
		}
		goto tr0
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
		if data[p] == 32 {
			goto tr695
		}
		goto tr0
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		switch data[p] {
		case 108:
			goto st495
		case 110:
			goto st497
		}
		goto tr0
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		switch data[p] {
		case 32:
			goto tr702
		case 121:
			goto st496
		}
		goto tr0
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
		if data[p] == 32 {
			goto tr702
		}
		goto tr0
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		switch data[p] {
		case 32:
			goto tr704
		case 101:
			goto st498
		}
		goto tr0
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		if data[p] == 32 {
			goto tr704
		}
		goto tr0
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		if data[p] == 97 {
			goto st500
		}
		goto tr0
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		switch data[p] {
		case 114:
			goto st501
		case 121:
			goto st504
		}
		goto tr0
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		switch data[p] {
		case 32:
			goto tr709
		case 99:
			goto st502
		}
		goto tr0
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		if data[p] == 104 {
			goto st503
		}
		goto tr0
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		if data[p] == 32 {
			goto tr709
		}
		goto tr0
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		if data[p] == 32 {
			goto tr712
		}
		goto tr0
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		if data[p] == 111 {
			goto st506
		}
		goto tr0
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		if data[p] == 118 {
			goto st507
		}
		goto tr0
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		switch data[p] {
		case 32:
			goto tr715
		case 101:
			goto st508
		}
		goto tr0
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		if data[p] == 109 {
			goto st509
		}
		goto tr0
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		if data[p] == 98 {
			goto st510
		}
		goto tr0
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
		if data[p] == 101 {
			goto st511
		}
		goto tr0
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		if data[p] == 114 {
			goto st512
		}
		goto tr0
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		if data[p] == 32 {
			goto tr715
		}
		goto tr0
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
		if data[p] == 99 {
			goto st514
		}
		goto tr0
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		if data[p] == 116 {
			goto st515
		}
		goto tr0
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
		switch data[p] {
		case 32:
			goto tr723
		case 111:
			goto st516
		}
		goto tr0
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		if data[p] == 98 {
			goto st517
		}
		goto tr0
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		if data[p] == 101 {
			goto st518
		}
		goto tr0
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		if data[p] == 114 {
			goto st519
		}
		goto tr0
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		if data[p] == 32 {
			goto tr723
		}
		goto tr0
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
		if data[p] == 101 {
			goto st521
		}
		goto tr0
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		if data[p] == 112 {
			goto st522
		}
		goto tr0
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		switch data[p] {
		case 32:
			goto tr730
		case 116:
			goto st523
		}
		goto tr0
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
		if data[p] == 101 {
			goto st524
		}
		goto tr0
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
		if data[p] == 109 {
			goto st525
		}
		goto tr0
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		if data[p] == 98 {
			goto st526
		}
		goto tr0
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
		if data[p] == 101 {
			goto st527
		}
		goto tr0
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
		if data[p] == 114 {
			goto st528
		}
		goto tr0
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
		if data[p] == 32 {
			goto tr730
		}
		goto tr0
tr656:
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
//line ragel/parse_datetime.go:15007
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
		case 45:
			goto tr156
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
		if data[p] == 45 {
			goto tr156
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
		case 45:
			goto tr162
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
		if data[p] == 45 {
			goto tr162
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
		case 45:
			goto tr168
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
		if data[p] == 45 {
			goto tr168
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
		case 45:
			goto tr176
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
		if data[p] == 45 {
			goto tr176
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
		case 45:
			goto tr185
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
		if data[p] == 45 {
			goto tr185
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
		case 45:
			goto tr192
		case 121:
			goto st565
		}
		goto tr0
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
		if data[p] == 45 {
			goto tr192
		}
		goto tr0
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
		switch data[p] {
		case 45:
			goto tr194
		case 101:
			goto st567
		}
		goto tr0
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		if data[p] == 45 {
			goto tr194
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
		case 45:
			goto tr199
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
		if data[p] == 45 {
			goto tr199
		}
		goto tr0
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		if data[p] == 45 {
			goto tr202
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
		case 45:
			goto tr205
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
		if data[p] == 45 {
			goto tr205
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
		case 45:
			goto tr213
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
		if data[p] == 45 {
			goto tr213
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
		case 45:
			goto tr220
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
		if data[p] == 45 {
			goto tr220
		}
		goto tr0
tr651:
//line ragel/datetime.rl:7
 pb = p 
	goto st598
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
//line ragel/parse_datetime.go:15693
		switch data[p] {
		case 32:
			goto tr655
		case 45:
			goto tr656
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st459
		}
		goto tr0
tr652:
//line ragel/datetime.rl:7
 pb = p 
	goto st599
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
//line ragel/parse_datetime.go:15713
		switch data[p] {
		case 32:
			goto tr655
		case 45:
			goto tr656
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st459
		}
		goto tr0
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		if data[p] == 97 {
			goto st601
		}
		goto tr0
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
		if data[p] == 121 {
			goto st602
		}
		goto tr0
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
		switch data[p] {
		case 32:
			goto st274
		case 44:
			goto st456
		}
		goto tr0
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
		switch data[p] {
		case 97:
			goto st604
		case 117:
			goto st610
		}
		goto tr0
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		if data[p] == 110 {
			goto st605
		}
		goto tr0
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
		switch data[p] {
		case 32:
			goto tr810
		case 45:
			goto tr811
		case 47:
			goto tr811
		case 117:
			goto st606
		}
		goto tr0
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
		if data[p] == 97 {
			goto st607
		}
		goto tr0
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
		if data[p] == 114 {
			goto st608
		}
		goto tr0
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		if data[p] == 121 {
			goto st609
		}
		goto tr0
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
		switch data[p] {
		case 32:
			goto tr810
		case 45:
			goto tr811
		case 47:
			goto tr811
		}
		goto tr0
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		switch data[p] {
		case 108:
			goto st611
		case 110:
			goto st613
		}
		goto tr0
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		switch data[p] {
		case 32:
			goto tr818
		case 45:
			goto tr819
		case 47:
			goto tr819
		case 121:
			goto st612
		}
		goto tr0
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
		switch data[p] {
		case 32:
			goto tr818
		case 45:
			goto tr819
		case 47:
			goto tr819
		}
		goto tr0
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		switch data[p] {
		case 32:
			goto tr821
		case 45:
			goto tr822
		case 47:
			goto tr822
		case 101:
			goto st614
		}
		goto tr0
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
		switch data[p] {
		case 32:
			goto tr821
		case 45:
			goto tr822
		case 47:
			goto tr822
		}
		goto tr0
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
		switch data[p] {
		case 97:
			goto st616
		case 111:
			goto st621
		}
		goto tr0
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
		switch data[p] {
		case 114:
			goto st617
		case 121:
			goto st620
		}
		goto tr0
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
		switch data[p] {
		case 32:
			goto tr828
		case 45:
			goto tr829
		case 47:
			goto tr829
		case 99:
			goto st618
		}
		goto tr0
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
		if data[p] == 104 {
			goto st619
		}
		goto tr0
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		switch data[p] {
		case 32:
			goto tr828
		case 45:
			goto tr829
		case 47:
			goto tr829
		}
		goto tr0
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		switch data[p] {
		case 32:
			goto tr832
		case 45:
			goto tr833
		case 47:
			goto tr833
		}
		goto tr0
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		if data[p] == 110 {
			goto st273
		}
		goto tr0
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		if data[p] == 111 {
			goto st623
		}
		goto tr0
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
		if data[p] == 118 {
			goto st624
		}
		goto tr0
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
		switch data[p] {
		case 32:
			goto tr836
		case 45:
			goto tr837
		case 47:
			goto tr837
		case 101:
			goto st625
		}
		goto tr0
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
		if data[p] == 109 {
			goto st626
		}
		goto tr0
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
		if data[p] == 98 {
			goto st627
		}
		goto tr0
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
		if data[p] == 101 {
			goto st628
		}
		goto tr0
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
		if data[p] == 114 {
			goto st629
		}
		goto tr0
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
		switch data[p] {
		case 32:
			goto tr836
		case 45:
			goto tr837
		case 47:
			goto tr837
		}
		goto tr0
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
		if data[p] == 99 {
			goto st631
		}
		goto tr0
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		if data[p] == 116 {
			goto st632
		}
		goto tr0
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
		switch data[p] {
		case 32:
			goto tr845
		case 45:
			goto tr846
		case 47:
			goto tr846
		case 111:
			goto st633
		}
		goto tr0
	st633:
		if p++; p == pe {
			goto _test_eof633
		}
	st_case_633:
		if data[p] == 98 {
			goto st634
		}
		goto tr0
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
		if data[p] == 101 {
			goto st635
		}
		goto tr0
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
		if data[p] == 114 {
			goto st636
		}
		goto tr0
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
		switch data[p] {
		case 32:
			goto tr845
		case 45:
			goto tr846
		case 47:
			goto tr846
		}
		goto tr0
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
		switch data[p] {
		case 97:
			goto st638
		case 101:
			goto st642
		case 117:
			goto st621
		}
		goto tr0
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
		if data[p] == 116 {
			goto st639
		}
		goto tr0
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
		switch data[p] {
		case 32:
			goto st274
		case 44:
			goto st456
		case 117:
			goto st640
		}
		goto tr0
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
		if data[p] == 114 {
			goto st641
		}
		goto tr0
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
		if data[p] == 100 {
			goto st600
		}
		goto tr0
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
		if data[p] == 112 {
			goto st643
		}
		goto tr0
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
		switch data[p] {
		case 32:
			goto tr857
		case 45:
			goto tr858
		case 47:
			goto tr858
		case 116:
			goto st644
		}
		goto tr0
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
		if data[p] == 101 {
			goto st645
		}
		goto tr0
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
		if data[p] == 109 {
			goto st646
		}
		goto tr0
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
		if data[p] == 98 {
			goto st647
		}
		goto tr0
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
		if data[p] == 101 {
			goto st648
		}
		goto tr0
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
		if data[p] == 114 {
			goto st649
		}
		goto tr0
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
		switch data[p] {
		case 32:
			goto tr857
		case 45:
			goto tr858
		case 47:
			goto tr858
		}
		goto tr0
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
		switch data[p] {
		case 104:
			goto st651
		case 117:
			goto st654
		}
		goto tr0
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		if data[p] == 117 {
			goto st652
		}
		goto tr0
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
		switch data[p] {
		case 32:
			goto st274
		case 44:
			goto st456
		case 114:
			goto st653
		}
		goto tr0
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		if data[p] == 115 {
			goto st641
		}
		goto tr0
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
		if data[p] == 101 {
			goto st655
		}
		goto tr0
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
		switch data[p] {
		case 32:
			goto st274
		case 44:
			goto st456
		case 115:
			goto st641
		}
		goto tr0
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
		if data[p] == 101 {
			goto st657
		}
		goto tr0
	st657:
		if p++; p == pe {
			goto _test_eof657
		}
	st_case_657:
		if data[p] == 100 {
			goto st658
		}
		goto tr0
	st658:
		if p++; p == pe {
			goto _test_eof658
		}
	st_case_658:
		switch data[p] {
		case 32:
			goto st274
		case 44:
			goto st456
		case 110:
			goto st659
		}
		goto tr0
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
		if data[p] == 101 {
			goto st653
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
	_test_eof660: cs = 660; goto _test_eof
	_test_eof661: cs = 661; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof662: cs = 662; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof663: cs = 663; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof664: cs = 664; goto _test_eof
	_test_eof665: cs = 665; goto _test_eof
	_test_eof666: cs = 666; goto _test_eof
	_test_eof667: cs = 667; goto _test_eof
	_test_eof668: cs = 668; goto _test_eof
	_test_eof669: cs = 669; goto _test_eof
	_test_eof670: cs = 670; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof671: cs = 671; goto _test_eof
	_test_eof672: cs = 672; goto _test_eof
	_test_eof673: cs = 673; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof674: cs = 674; goto _test_eof
	_test_eof675: cs = 675; goto _test_eof
	_test_eof676: cs = 676; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof677: cs = 677; goto _test_eof
	_test_eof678: cs = 678; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof679: cs = 679; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof680: cs = 680; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof681: cs = 681; goto _test_eof
	_test_eof682: cs = 682; goto _test_eof
	_test_eof683: cs = 683; goto _test_eof
	_test_eof684: cs = 684; goto _test_eof
	_test_eof685: cs = 685; goto _test_eof
	_test_eof686: cs = 686; goto _test_eof
	_test_eof687: cs = 687; goto _test_eof
	_test_eof688: cs = 688; goto _test_eof
	_test_eof689: cs = 689; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof690: cs = 690; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof691: cs = 691; goto _test_eof
	_test_eof692: cs = 692; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof693: cs = 693; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
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
	_test_eof27: cs = 27; goto _test_eof
	_test_eof707: cs = 707; goto _test_eof
	_test_eof708: cs = 708; goto _test_eof
	_test_eof709: cs = 709; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof710: cs = 710; goto _test_eof
	_test_eof711: cs = 711; goto _test_eof
	_test_eof712: cs = 712; goto _test_eof
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
	_test_eof713: cs = 713; goto _test_eof
	_test_eof714: cs = 714; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof715: cs = 715; goto _test_eof
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
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
	_test_eof189: cs = 189; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof716: cs = 716; goto _test_eof
	_test_eof193: cs = 193; goto _test_eof
	_test_eof194: cs = 194; goto _test_eof
	_test_eof717: cs = 717; goto _test_eof
	_test_eof718: cs = 718; goto _test_eof
	_test_eof719: cs = 719; goto _test_eof
	_test_eof720: cs = 720; goto _test_eof
	_test_eof721: cs = 721; goto _test_eof
	_test_eof722: cs = 722; goto _test_eof
	_test_eof723: cs = 723; goto _test_eof
	_test_eof195: cs = 195; goto _test_eof
	_test_eof196: cs = 196; goto _test_eof
	_test_eof197: cs = 197; goto _test_eof
	_test_eof724: cs = 724; goto _test_eof
	_test_eof725: cs = 725; goto _test_eof
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
	_test_eof726: cs = 726; goto _test_eof
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
	_test_eof727: cs = 727; goto _test_eof
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

	_test_eof: {}
	if p == eof {
		switch cs {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279, 280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299, 300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339, 340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 383, 384, 385, 386, 387, 388, 389, 390, 391, 392, 393, 394, 395, 396, 397, 398, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422, 423, 424, 425, 426, 427, 428, 429, 430, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 446, 447, 448, 449, 450, 451, 452, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 495, 496, 497, 498, 499, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 510, 511, 512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 524, 525, 526, 527, 528, 529, 530, 531, 532, 533, 534, 535, 536, 537, 538, 539, 540, 541, 542, 543, 544, 545, 546, 547, 548, 549, 550, 551, 552, 553, 554, 555, 556, 557, 558, 559, 560, 561, 562, 563, 564, 565, 566, 567, 568, 569, 570, 571, 572, 573, 574, 575, 576, 577, 578, 579, 580, 581, 582, 583, 584, 585, 586, 587, 588, 589, 590, 591, 592, 593, 594, 595, 596, 597, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 608, 609, 610, 611, 612, 613, 614, 615, 616, 617, 618, 619, 620, 621, 622, 623, 624, 625, 626, 627, 628, 629, 630, 631, 632, 633, 634, 635, 636, 637, 638, 639, 640, 641, 642, 643, 644, 645, 646, 647, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 658, 659:
//line ragel/datetime.rl:5
 return st, err 
		case 667, 678, 709, 720, 725:
//line ragel/datetime.rl:9
 st.Zoned = true 
		case 714:
//line ragel/datetime.rl:15

    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])

		case 716, 726, 727:
//line ragel/datetime.rl:19

    st.Year, _ = strconv.Atoi(data[pb:pb+4])

		case 715:
//line ragel/datetime.rl:23

    st.Year = parse_year_2_digits(data[pb:pb+2])

		case 660, 713:
//line ragel/datetime.rl:27

    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])

		case 680, 690:
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

		case 663, 677:
//line ragel/datetime.rl:52

    st.Ad_bc = ADBC_BC;

		case 675:
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

		case 710, 711, 712:
//line ragel/datetime.rl:92

    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }

		case 679:
//line ragel/datetime.rl:101

    st.Hour, _ = strconv.Atoi(data[pb:pb+2])

		case 672, 706, 707:
//line ragel/datetime.rl:104

    st.Hour, _ = strconv.Atoi(data[pb:pb+1])

		case 692:
//line ragel/datetime.rl:107

    st.Minute, _ = strconv.Atoi(data[pb:pb+2])

		case 691, 705:
//line ragel/datetime.rl:110

    st.Minute, _ = strconv.Atoi(data[pb:pb+1])

		case 703:
//line ragel/datetime.rl:113

    st.Second, _ = strconv.Atoi(data[pb:pb+2])

		case 693, 704:
//line ragel/datetime.rl:116

    st.Second, _ = strconv.Atoi(data[pb:pb+1])

		case 681, 682, 683, 684, 685, 686, 687, 688, 689, 694, 695, 696, 697, 698, 699, 700, 701, 702:
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

		case 668, 669, 721, 722:
//line ragel/datetime.rl:159

    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }

//line ragel/datetime.rl:9
 st.Zoned = true 
		case 664, 665, 666, 670, 717, 718, 719, 723:
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
		case 671, 724:
//line ragel/datetime.rl:196

    st.ZoneName = data[pb:p]
    st.Zoned = true

//line ragel/datetime.rl:9
 st.Zoned = true 
//line ragel/parse_datetime.go:17320
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

