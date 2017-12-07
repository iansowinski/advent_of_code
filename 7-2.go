package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Program struct {
	name string
	weight int
	children []Program
	childrenStrings []string
}

func main() {
	// input := "pbga (66)\nxhth (57)\nebii (61)\nhavc (66)\nktlj (57)\nfwft (72) -> ktlj, cntj, xhth\nqoyq (66)\npadx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\njptl (61)\nugml (68) -> gyxo, ebii, jptl\ngyxo (61)\ncntj (57)"
	input := "czlmv (78)\nfiwbd (57)\ntwxnswy (98)\ndhrnu (62)\nyptrqp (64)\nvikplll (9)\nlgxjcy (359) -> vsbfmt, kbdcl\ntvisu (70)\nimxraee (33)\nqfbqiv (55) -> gwlyvbx, uanby\neolilw (78)\ngywnt (52)\ndnybu (88)\nyiyhuxk (15)\naxrep (33)\najtjym (11)\nzjmhab (68)\nxcstcy (20) -> iequdmw, fozfiqf, ytkmda, ywawq, nfgmyd, bzumy\ndgscbm (29)\nsghuxq (9)\njfmuzo (164) -> istoj, jyzrmnp\nzjklq (52)\nzwcucws (104) -> vqghe, alycm\nldfhqjw (809) -> wpprys, heotf, dybitt\nckbjuyy (63) -> ygvwdmg, pwxorqj, hshbiaq, oxqvzzv\nxnalv (11)\nowkily (191) -> cqnmo, fhqah\nwrbknnc (80)\nimakkq (99)\niymrfq (57) -> iwdjjf, zamje, eduxyw\numzsfe (66)\nruvmsg (55)\newtdu (66)\nuwcqi (91) -> umzsfe, tqjzv\nnqqxz (123) -> jrkbh, hyyeb, abwmmuw\nmetyk (77) -> euriiux, zztvgwh\nsmxwnsk (16)\nxhodzi (31)\nunscz (67) -> dpavyu, lrdmltp\nyrluhjr (50) -> qdpcslx, echdbt, aeeas, uodyv\nvvqtv (91)\nsigxauh (64) -> dnybu, kpnbj\nnqjmi (49)\nrsocykv (55)\nvqghe (93)\nlumomlj (5)\newods (83)\nqggycuv (26)\nzqbtgtw (94) -> ewtdu, xpmpulq, kpnjmwz\nqnbnflf (53)\naintn (73)\nibnho (16)\nvsbfmt (25)\nzwake (64)\nswpvgrn (36)\nvlsyqg (93) -> lrauqzk, gdhekde\nzawbrvg (10)\nwqgczyl (200) -> ilgue, bwdfntd, zlgst\nqhyan (391) -> obxyz, ahutsc\nvrppl (78)\nrqebmc (176) -> snnxhgn, anucwoj\nhhpib (98)\niqdjcwm (57)\nxpmpulq (66)\nafckhxn (16)\nuynmvcu (304) -> imxraee, rmfwij\ncybhcko (256) -> pgfomf, dacrm\nylrxgoz (35)\nedmzea (60)\nxuokqm (35)\nrmfwij (33)\ngtezger (64)\ndmjech (59)\navbnxz (47)\naapssr (72) -> twimhx, wfdiqkg, guuri, qwada, mwsivlf\nguuri (38) -> ydmxb, lndfaxo, sjfue, hufrsq, rxfma\nuyaarr (55)\njtudcya (98)\nfwgdynj (92)\nkpnbj (88)\nanucwoj (13)\npaxdij (10)\ncadrzx (29)\nodbbac (614) -> lnkowzc, zqcluqc, pvatk, thymhif, gchcumo, rffqotx, juscmde\nmaiiqc (133) -> apagba, mhhbej\nvoeps (50)\napagba (36)\ndrpgz (590) -> zabqp, vercm, maiiqc, egyfwlq\nhgtyjhv (23) -> ffqwo, yqzdjj, yrwcj, ilxgo, bvoxahz, iymrfq\nzjevsui (23) -> feiitan, upgfr\nbtzlwo (37)\nvzpqvfx (51)\nliwqqqi (117) -> ckdeo, movesh, abgfgav, mrlnt\nudslj (75)\ntwhcdv (58) -> uqzoo, ytwvl, gjxgd, ehcrq\nqtfzj (92)\nxlufcc (1391) -> qlsdtum, swkor, ndsvlr\nhghpgj (11)\noxckpuw (218) -> nssexu, uwogu\nrlwyzv (86)\nwontek (77) -> lqfsw, ejljdxo, jtudcya\nzpeuqlo (57)\ncoibvjw (65)\nxuqozk (9)\nwdyxwia (45)\nzetdtq (93)\nenymia (66) -> ckurzc, lcohspm\nheotf (5) -> gminpu, jodhx\nbttlx (37)\nsnnxhgn (13)\nuqcxvh (70)\nzvxhi (92)\nyfzypu (50)\nxhrttm (17)\nerhiag (13)\nwbnux (92)\nupgfr (87)\nhxkeypq (43)\ndexabuc (5)\ncjjwz (27)\nkbdcl (25)\nrrmtlri (80) -> jkknc, vvkrst, vblyk, ghiftw, tilvzsd, drzoi\nujaiypo (31)\nmgcyq (91)\nmqrgut (255) -> dhfrnu, bzijz\nzobnnps (411)\ncxjmggj (286) -> kmkqsvl, remwhfn\nxpinqek (50)\nfzvyv (14)\nvvkrst (264)\nlrdmltp (64)\nmjgilv (163) -> ifegka, olwcj, eheeelo, cvjhub\nrhqfa (80)\nwswjfu (52)\ntqbouv (443) -> dtqxiic, chsuzvr\nxxyjm (267)\njofbj (256) -> vikplll, vjaqo\nzbtptb (394) -> hiuuk, diriz, liudq\nixuyx (92)\nyfkgike (56) -> dhrnu, ehzvm\nrchhwx (46)\ngatdc (35)\nzheomzg (12) -> kkimdb, hhomxd\nlzstd (47)\nzqzrxe (10)\nkpnjmwz (66)\nqvpob (322) -> vtzfsci, uexfxv\nckurzc (85)\nrpxhqmz (240)\nmhhyoxf (86)\nryiytg (63)\nkmxek (226) -> rbgonm, qgfxzo, sghuxq, hauul\nidprbpo (146) -> wkurl, plwdhth, staoh\nliudq (19)\ncmouh (82) -> hmulqgo, qjsjrz, owkily, jkqbf, bipzxj\npneftv (97)\nulwzhy (54)\nvbdycf (931) -> uwrmetx, lruhvsn, hssfwi\ngdckdzx (37)\ndueof (87)\nhooxvp (40) -> rahsrp, uqcxvh\nnxgnumb (91) -> hdmdxi, oxxse, jhmvp\nvekdbjs (15) -> etznlt, kvwqrd, udslj\nptamwey (245) -> laovxxw, cpdtw\nictgyp (241) -> zbiypv, rrmtlri, odbbac, tkqke\ndybitt (37) -> ulwzhy, wrumij\ncfuned (17)\nnvgnls (54)\niwdjjf (79)\nfvifqfr (91) -> uxgbnjk, phckimp, apeba, udzzv\ncbtgtqb (24)\nyidnjq (10)\nbbdayq (86)\nbzzsss (92)\nzedqbr (26)\nyljjx (276) -> wumlwia, rnvlepv\nsarwcau (48)\nrquaisv (389) -> fbbkkxh, xhodzi\nzlgst (54)\ndxxzpr (16)\nvatlgyb (16)\nqojqcd (5)\nazpmzli (262) -> xodsc, tzlzl\nhusns (217) -> btzlwo, sntovuq\nwltlu (59)\nnzzwzmh (38)\napeba (90)\ncflrkd (31)\nfbmalu (275) -> qtqbdp, ldrfyi\nqnvby (57)\novgpa (29)\nweejuk (64)\nourium (26)\nyckuqt (43)\nhyyeb (53)\nhgrwjoy (221)\ncfqczei (39)\nbhwnch (184) -> qhyan, fvifqfr, rquaisv\niibalc (73) -> dgmevpb, mojye, zcfztz, qsvyfv, ulxxt, hzxiv, spovj\ntlpydjm (48)\npfvoay (31)\nvmpev (37)\neheeelo (22)\ntzlzl (6)\noncjf (68)\nkirowa (37)\nuqzqmo (124) -> pgkeahi, owgile\npwxorqj (91)\ndhfrnu (66)\nogjqrb (199)\nmwsivlf (18042) -> xcstcy, jluccf, pftdcqs\ncqbtoo (11)\nhlysw (14)\nekgsba (57)\nzcfztz (228) -> scyuq, vggwo\nowgile (18)\nkeczktv (10)\ngtixbu (85)\nhgywlzc (135) -> osysik, fmmgboi\nbwdfntd (54)\natoveeu (297) -> crtnt, qnvby\nmuecmt (51)\ntqjzv (66)\nzsraq (47)\nidoml (79)\nmwgimvy (49)\ngchcumo (138) -> msqtvwu, feqefvb\nzrmqrif (32)\nqsrfkq (118) -> sgxil, tiugk\nlrauqzk (88)\ndckozkt (89) -> kbefh, pnqri, zdnvz, tuzquw\nevsnlxn (99)\nvwwvjg (1005) -> nqqxz, dfewulo, qsrfkq\ndacrm (12)\nwnrqkg (362)\nnimeytn (147) -> difavgr, idcvpt, mjgrhot, glllgqs, qsguy\nbdrerk (146) -> oswqr, gjjyumz\nzfdhua (65) -> deewhp, nerzb\nkfjbjl (243) -> hatmf, tywedfu\nvsijk (213) -> hvtsr, qwhelvj\nlcohspm (85)\nbrqvmh (70)\nsinhju (87)\nuilshuf (20)\njodhx (70)\npuiks (20)\nwvuqq (65)\nzamje (79)\nrqfjzox (86)\nyjfazsr (44)\nkufsr (425) -> sdeana, byrxbgo, tvarspr\nnbmat (39)\ngwrred (131) -> aujwg, touzm\nsvkcvg (53) -> rbzqh, abnhpi, zrmqrif, kvtpc\ndasbhc (80) -> pmlim, gwoiid\nopnbls (202)\nknqqz (123) -> vebli, axwdrai\njyzrmnp (49)\ndifavgr (236) -> ukzfd, rxkpkvi\nwpjaa (76)\nsbfhxsx (99) -> pusdii, qqodh, tlpydjm, jgoaja\nqaedcl (61)\noxwnu (24)\ngtvpf (65) -> knqqz, qzfwf, vlsyqg, anvxtbs, idprbpo\nhuansqc (29)\noaabog (705) -> nmbql, yodnj, yrluhjr\nghiftw (182) -> gpikz, mvkxva\nqfhfg (50)\nzqcluqc (30) -> kdwzx, yflpz, aokvbbi\ntwimhx (25773) -> xsmdo, gpormp, akioek\nbrdbeoe (23)\ncfshobe (91)\njojnh (241)\nettsd (32) -> rzpmulz, fmjfhoo\nliymh (86)\nhmulqgo (275) -> xjrheul, cykhmjj\nzbptsb (80)\nvjaqo (9)\nwmfro (410) -> mcpyfmy, rzlxx\nwpprys (113) -> tkcmrl, dxxzpr\nfarpozz (13) -> wepvw, zzjmukx, ooenol\nwpwxn (51)\nizphu (97)\nepuiwg (54)\nefcqep (79)\ndnvqlog (53)\nknwlsi (54)\nzfrsmm (763) -> aapsp, gfsnb, jojnh\nzjxcrv (43)\nlnkowzc (88) -> aaqsdv, ujaiypo\nvgtxty (199) -> fcbxsf, acchih\ngrgjs (345) -> uvntywl, nzzwzmh\nlewfa (45) -> aouljs, swvae, ybwgz, spdsz\nnfgmyd (1390) -> yljjx, lcdleci, zcoyhde\nzzjmukx (69)\nvjwptwp (36)\nspfds (86)\nvqbsvk (72)\neyslsq (84) -> oncjf, nuibgd\ngvegx (90) -> qhujoiu, dpvmn, cyhnayc\nopfey (37)\nypiwj (27) -> pupaehu, hekhbd, kgfldlb, atoveeu, zobnnps, iukfd\nvzkneq (52) -> gvegx, mqrgut, bsvmale, fjueoko, cdoed\nkynrf (104) -> bnmgqen, iuftirx\nwibgm (83) -> zetdtq, dcytmqp, xfxqvk\nalycm (93)\nrcnhfcz (13)\nsmbkvwd (276) -> rvwvqij, avbnxz\nkitggwr (91)\nkbsmiq (70)\ntouzm (32)\nxyfeczu (317)\nslnezfv (83)\nxpizz (60)\nseicwz (964) -> izpxbq, gdomj, jrsbjr\nmbjsjy (51)\naoltxw (51)\nzztvgwh (87)\nrznqdxq (60)\nvggwo (13)\ntlxdyxj (1022) -> cnztmm, dkjwj, drigjl\nqjsjrz (125) -> aqrxpru, polnw\nkhpntc (92) -> hxkeypq, wfefbdt, mhbov\nqrdtzc (35)\nglllgqs (348) -> paxdij, yidnjq\njnghyrm (107) -> soebcx, opfey\nabnhpi (32)\ntziqot (226) -> peyyn, chwoyn\nhgdbip (10)\ngoqhcu (60)\nkddex (87) -> goqhcu, rxohrzb, ucsvcy\nzfzwko (20)\nuhjxtrp (16)\nzmzwttu (54)\nvdyftz (50) -> ypttuns, gfcfsf, ymonx, fjdyms\nctien (177) -> cqbtoo, hghpgj\njfqlxa (86)\nnhwgy (44)\nlruhvsn (86) -> ljeiizi, opctg\nuxvgnv (29)\ncwdoo (164) -> eozev, gohdd\nghkmb (165) -> pnuvuf, nskayc\nwumlwia (34)\nxryqzl (52)\nfbhnbdw (90)\nwuijd (50)\ndannmef (61)\naokgycq (50)\nwfefbdt (43)\nzmqxmtf (92)\nbiioe (70)\nwwcqz (55)\nechdbt (83)\ntdtxcm (35)\ntvarspr (6)\nzkhtua (187) -> zmbimk, kmxek, tziqot, aewbhax\nphckimp (90)\nrcomsru (73)\naaztou (24)\npwdoio (184) -> nxcus, bpbmbq, zbtptb\nimpqm (64) -> yumbba, zltzea\nbkjcc (41)\ncnztmm (255)\njdbel (87)\nvcjavrw (60)\nudkba (87)\njretzi (202)\ndonza (96)\nstaoh (41)\nsubnd (77)\njzrlfrl (86)\nsgxil (82)\nwahmvp (49)\nlvtxk (35)\ncxrukf (5)\natcole (135) -> ykwvu, yckuqt\nbjeto (71) -> msoip, wjkky\nqsvyfv (220) -> cfuned, vzayzaq\nztodsnq (344) -> lqrde, ukvzqny, svzqi\npkditpt (99)\nrgshogg (33)\ndrefe (29)\nrnqvaa (91) -> wvuqq, uifug\ncpxkqg (96)\nbstznzu (45)\niflhr (7)\npczpj (60)\njhmvp (30)\naouljs (68)\nndqomrl (368) -> fhyfhjg, qnbnflf\nwrumij (54)\nehcrq (299) -> cpllg, asnsvv\nauvkb (40)\ngdomj (191)\ndnnlgt (77)\nwasnz (63)\nrzpmulz (85)\ntiugk (82)\nsdjzc (27)\njdxth (14)\naeeas (83)\ncrtnt (57)\nytwvl (313) -> nqjmi, mwgimvy\navmhq (86)\njuscmde (150)\nkbpvvf (72)\nwegmegy (280) -> oaantta, fzvyv\nkfdtg (68)\nmatvb (55)\nehzvm (62)\nebxyi (54)\nqgfxzo (9)\nzoodui (220)\ndapjhx (93)\nywawq (1093) -> ptcby, kufsr, pxsuagy\nlaovxxw (27)\nfurny (8) -> lgxjcy, hcsiqp, vsijk\ngzocll (70)\nkifypgu (63)\nkywpr (91)\neikjv (10)\nwjkky (57)\nfkwsxxk (1446) -> eyixz, bjeto, ghkmb\nweeriq (19)\nrxkpkvi (66)\nbipzxj (81) -> bpgafw, ilidonx, uxswi\ntrjird (51)\nistoj (49)\nidcvpt (71) -> zzdgmmd, imakkq, jbodnl\npehio (202) -> bstznzu, xzdhzoa\nexxbgm (77)\ntuzquw (83)\nydmxb (1610) -> udoiqj, isfyuxu, hgtyjhv, tlxdyxj, inreau, qiinez\njmsuq (31)\ndkjwj (155) -> nbbiqh, nbfbkao\nxxbwj (122) -> xpinqek, xqmsvth\ndfewulo (48) -> phmmydn, eyjdqm, qpblhez\nukvzqny (47)\ngpormp (3465) -> ymogkwe, ywtwrq, hixtw, zmssxbl, twhcdv\ndiriz (19)\nzltzea (47)\njkqbf (87) -> wpwxn, vzpqvfx, muecmt, trjird\nuodyv (83)\nfhqah (50)\naapsp (5) -> nvhbnlc, dzqax, ndjjag, dmjech\nbxfwhrd (51)\nknoziw (29) -> sinhju, otbxf, udkba\nwbupk (63)\nixvxtb (87)\nrnvlepv (34)\nnerzb (81)\nbrtmsqe (59)\ngjjyumz (81)\nogpvd (53)\nypttuns (56)\nyzkpvn (50)\nhssfwi (102) -> qfhfg, uvkecqz\nooenol (69)\nmucpe (31)\nbpbmbq (83) -> irkhv, navnyn, bzzsss, wbnux\nkkimdb (84)\nsfgatoz (40)\nbqfor (66)\ntobsrs (89)\npolnw (83)\nxlqycb (20)\nuwrmetx (202)\njtjzsr (13)\ntehdy (16)\nfeiitan (87)\nhcsiqp (65) -> rlwyzv, liymh, huzjje, spfds\nudghm (35)\niequdmw (572) -> cwxgg, geqnwz, smbkvwd, swhnbvw, uynmvcu\nmwxiwty (90) -> lvtxk, ylrxgoz\nxqmsvth (50)\nulszh (64)\ncpllg (56)\noswpicl (51)\ngcglasa (143) -> nbmat, cfqczei\nkdwzx (40)\nigumiu (373) -> ggynw, lrodp\njptqm (200) -> avrfgx, mucpe\nyrwcj (112) -> cfshobe, mgcyq\nuvkecqz (50)\nmjgrhot (314) -> sdjzc, cjjwz\nohuvi (89)\nuplxy (60)\nymfurx (26)\nvrkrr (55)\ntkxuj (86)\nmydixq (60)\nobgin (293) -> lewfa, ukwydlj, xyfeczu\nswvae (68)\nadlvipk (168) -> uilshuf, qxkilb\nvxjhric (220)\nmhbeqm (99)\nvcyfwm (88) -> qrdtzc, gatdc\nxujof (5)\nrjkhl (257) -> mogla, nlewmv, mxxecyh\noxqvzzv (91)\nrwxmoi (265) -> dnvqlog, ogpvd\ntcybjv (178) -> iuneyv, cgszcig\nxtpgpbx (15)\numtpolq (98)\nzkdwxd (64)\ngqiiycg (97)\npftdcqs (2087) -> xeefzwg, vyudgjy, ylvvgmj, fefnhby, ypiwj\nvyudgjy (1079) -> ettsd, knastam, jretzi, rqebmc, kngspv, onudby, opnbls\nmicfuva (85)\ndgmevpb (120) -> xcmre, upslw\nncghncy (36)\noxtckvl (144) -> erhiag, kvcak\ntilvzsd (98) -> slnezfv, sxick\nsnaebf (11)\nsdnzxl (2345) -> xhkrx, tcmkx\nqdpcslx (83)\nymogkwe (597) -> srqbs, hgrwjoy, khpntc, atcole, spdyx\ntywedfu (92)\nchhmaxr (10)\nonudby (18) -> zvxhi, zmqxmtf\nxdwzfh (45)\navrfgx (31)\nctafsp (26)\ncnbvz (360) -> zpeuqlo, iqdjcwm\nuuijd (98)\ngwlyvbx (86)\naewbhax (70) -> glhbvi, ulszh, rbawcc\nfcbxsf (15)\nyqzxa (306)\nbbmbn (65)\nnskayc (10)\nywuxmil (54) -> wasnz, ryiytg, raaxwq, jtkrr\ndtrxzld (131) -> wontek, rjkhl, rwxmoi\nshetkt (9) -> jqmcro, rqfjzox\nmovesh (26)\nuylnr (46)\nuexfxv (43)\ntjxms (92) -> cfwox, vpcfu, boiuyt, nvgnls\ncyhnayc (99)\nrwtdhr (49)\nglxtiz (80)\nhhmrwj (178) -> oswpicl, bxfwhrd\nnwkxd (19)\ndxgntli (46)\nxodsc (6)\nffgplgk (117) -> bkjcc, bclkir, cwhjt\nfdcnhr (92)\nwaojp (14)\njtkrr (63)\nmhbov (43)\nsxick (83)\nqwhelvj (98)\nthrmrw (134) -> tliov, xhpvb, vjbjy\nnavnyn (92)\nrpjiwuz (85) -> jmsuq, ctqgj\ncfwox (54)\nghwwfca (131) -> fsmidq, aaztou, bpcrtwl, vfojca\npnqri (83)\nlqdqct (98)\netznlt (75)\ngeqnwz (370)\norydf (91)\nywtwrq (1021) -> qfbqiv, ghwwfca, zfdhua\nayrucb (14)\njkmernz (94)\nrxohrzb (60)\nrffqotx (84) -> rgshogg, axrep\nxxtdti (83)\njluccf (8332) -> xlpzp, obgin, dtrxzld, nzbknfm, ldfhqjw\npnuvuf (10)\niqqevm (80)\niuftirx (33)\nsprno (14)\nzcyfth (24)\npoenou (76) -> evsnlxn, mhbeqm\npxifep (91)\ntlgew (49)\nxriiu (15)\nhxzthym (303) -> zsbcn, nwejb, dckozkt, grgjs\njagoxt (10)\nnuibgd (68)\nemqqx (17)\nerzayz (611) -> ywuxmil, clhyd, yqzxa, cxjmggj, kcqvnd, cvipqzv\numjwkfc (94) -> wpcry, pkbxhn, xmnreh\ntjnxnd (123) -> gnitpo, yfzypu\nkxftjs (43) -> tjxkxom, hrbsamg, qvpob\ndbrlat (411) -> sprno, waojp\ndqggz (76)\naqrxpru (83)\nqedsq (894) -> sscfkq, waxrcby, fkwsxxk\nkngspv (60) -> hhuhid, mhoaev\nyonxye (177) -> qallt, vmpev\ngoqdmgy (85) -> fiwbd, ekgsba\nnboczvq (55)\nzlrvs (78)\nzabqp (147) -> uxvgnv, huansqc\nndsvlr (88) -> mbjsjy, mjxtbvx\nhnnjspe (79) -> wuijd, aokgycq, yzkpvn\nhjjri (1300) -> vrkrr, jaxae\nbvoxahz (184) -> kmvkkze, nboczvq\nbpareqw (391) -> pehio, pxjve, zqbtgtw\nafpxssv (45)\nibzpdf (119) -> gywnt, jogksw\nsdeana (6)\nsgzijsv (76)\npeyyn (18)\nxdszea (185) -> kbpvvf, vqbsvk\nkvwqrd (75)\ntyuyffp (228) -> tlcseu, xryqzl, hraby\nfistzz (1760) -> vgtxty, bffvo, hnnjspe\nzqqyhq (60)\nqwada (40305) -> wfkcsb, qlboef, pkowhq\nremwhfn (10)\nvatqulm (66)\npsirxj (42) -> rcomsru, aintn\nmhblqw (10)\nuwogu (38)\nfsmidq (24)\nbsvmale (79) -> exxbgm, trvjrbi, subnd, dnnlgt\nzsbcn (421)\nilgue (54)\nldsbhlt (153) -> ayrucb, hlysw\nscyuq (13)\npiqgu (833) -> dkwde, dkhjdb, dqoatul, psirxj, mbmaqh, gxdqpcd\nqsemjfo (241) -> yjfazsr, aeiktt\nbfxmh (94)\nrrnjpdu (280) -> lmuwznh, iflhr\njaxae (55)\nnbbiqh (50)\nvwisj (44)\ngwoiid (78)\nctidma (26) -> ncghncy, swpvgrn, xwtgr, vjwptwp\noovlaqz (69)\nfezhnw (71) -> ixuyx, pkesm, fdcnhr, qtfzj\nhauul (9)\ndkhjdb (66) -> dannmef, qaedcl\nrtlde (120) -> glxtiz, rhqfa\nqtfuwru (24)\nmqyipy (23)\nhvrmmxf (94)\nswhnbvw (54) -> zodbq, qmiurj, idoml, efcqep\nzcoyhde (92) -> esjuqh, vxucz, kifypgu, bveaw\nqhujoiu (99)\nsophzzn (407) -> uhjxtrp, vatlgyb\nwfhixl (15)\neyixz (175) -> hmeolu, lumomlj\nldrzzzz (11)\nhdmdxi (30)\nazrcn (98)\njbodnl (99)\ndkrdqpo (15)\nwpcry (62)\nivnufgs (93)\nmgreg (68)\nvxobef (73)\nnxcus (361) -> wdyxwia, nhhyvm\nbdgwdt (280)\nmculw (54)\njmavfy (35)\nasnsvv (56)\nwpezq (224) -> qppamj, pokjgkl\ngsaci (313) -> tkxuj, jzrlfrl\neeuohnm (70)\nzzdgmmd (99)\nmsoip (57)\nspovj (58) -> lqdqct, omxny\ncovpz (123) -> wbzwxs, hvlfx\nukzfd (66)\nygldlhw (17) -> bbmbn, coibvjw\nxmnreh (62)\nnssexu (38)\nzibnjxk (11)\nfhyfhjg (53)\njaebowl (50)\nvfojca (24)\ngminpu (70)\nylpahav (29)\nzdnvz (83)\ntlulph (94)\nhvlfx (49)\nsadhkpb (204) -> edmzea, pczpj, jxevrgz\nthymhif (150)\newfql (37)\ncykhmjj (8)\nxlpzp (158) -> jnghyrm, nxgnumb, ldsbhlt, svkcvg, avnosfl, shetkt\nylvvgmj (1785) -> csvxlo, enymia, dasbhc\ntdfgfqn (55)\nrnyarzh (44)\nexgsu (94)\nheqvbi (48)\nybwgz (68)\nrbzqh (32)\niwwngux (98) -> hvrmmxf, nebwj, bfxmh, fkrdkav\nwolyko (13)\njrsbjr (45) -> vxobef, zsjoz\nwowgkro (1847) -> vcvjtf, biioe\nxjrheul (8)\ngohdd (63)\ntlfbyq (168) -> zqzrxe, jagoxt, zawbrvg, chhmaxr\njgoaja (48)\nghmxos (43)\nbvqcnc (195)\nffqwo (116) -> ohuvi, ygtbl\nqallt (37)\naujwg (32)\nupslw (67)\naxwdrai (73)\nhyeetiq (70)\ninreau (23) -> rosop, baimp, oxckpuw, nehwool, rrnjpdu, rqecouz\nwfdiqkg (27213) -> ictgyp, yffxaiy, prvckio, qedsq, fjmjc\nobxyz (30)\nhatmf (92)\nxffmho (423) -> xdszea, qsemjfo, fvfehj\nomxny (98)\nkfjqheo (1683) -> gwrred, bvqcnc, unscz\nbffvo (169) -> puiks, zfzwko, xlqycb\ntpbeg (6) -> sgzijsv, ysgnao\nonwqp (22) -> mowkbi, gsaci, wuogj, ztodsnq, tqbouv\nfcfbxmb (73)\niuneyv (48)\nqybyrm (98)\ncszchsg (73) -> ieqgh, fezhnw, sophzzn, dbrlat, igumiu\nlrodp (33)\nqmiurj (79)\nrbawcc (64)\nhhomxd (84)\neuriiux (87)\nesjuqh (63)\npxsuagy (369) -> ewfql, gdckdzx\nxhpvb (52)\njroea (42) -> jmavfy, udghm, tdtxcm\ntmfyif (15)\nstfqd (59) -> zqqyhq, uplxy, rznqdxq, xpizz\nvcwlaw (68) -> zeaxycf, izphu\neyrgmc (102) -> bbdayq, avmhq\npusdii (48)\ncgszcig (48)\nbyrxbgo (6)\nukwydlj (135) -> vvqtv, ffjie\nboiuyt (54)\nmnbkaa (93) -> pneftv, gqiiycg, hrqaph\nvxucz (63)\nejljdxo (98)\ndkwde (50) -> shwsv, oovlaqz\naktswyh (51)\ndvvkzv (40)\nqppamj (8)\nbftoy (221)\ndzqax (59)\nrlrnjnz (1115) -> dqggz, wpjaa\nqzfwf (247) -> ldrzzzz, cxfqupv\nffjie (91)\nmowkbi (93) -> uuijd, umtpolq, twxnswy, ivowf\nljeiizi (58)\nemsjxz (47) -> jaebowl, vyxgc, voeps\nhrqgfhg (415) -> vcwlaw, dsrostg, tfzsp, jptqm, vmuqpj, jfmuzo\nxwtgr (36)\nctqgj (31)\nbwqxk (220) -> xtpgpbx, dkrdqpo, wfhixl, yiyhuxk\nkpwfqfc (93)\nulxxt (236) -> xuqozk, tyxfxu\nyumbba (47)\nrvwvqij (47)\nprvckio (93) -> zqjqic, kfjqheo, cszchsg\nuifug (65)\nzeaxycf (97)\ncxfqupv (11)\nkcqvnd (306)\npxjve (246) -> rizbeea, brdbeoe\nmhoaev (71)\nrzlxx (32)\nrnoaimc (99)\nolwcj (22)\nfmmgboi (31)\nbclkir (41)\nhraby (52)\njqmcro (86)\nkvcak (13)\nmsqtvwu (6)\ntrvjrbi (77)\njogksw (52)\ncdoed (240) -> tlgew, rwtdhr, wahmvp\nqpblhez (78)\nrqecouz (260) -> emqqx, xhrttm\nzmbimk (22) -> zbptsb, wrbknnc, iqqevm\nzmssxbl (1033) -> tjnxnd, ibzpdf, uwcqi\nqlsdtum (158) -> ibnho, smxwnsk\nhyzzwx (93)\nbhvvs (173) -> zcyfth, oxwnu\nsoebcx (37)\nlgltgc (98)\nlcbljqb (226) -> bcpbir, zlubkqa\njrkbh (53)\nicziaym (1507) -> uqzqmo, mwxiwty, ihskp\nlqrde (47)\nhrqaph (97)\nqyrhb (486) -> wegmegy, tjxms, bdrerk\nzwvjv (89)\nrizbeea (23)\nmjxtbvx (51)\ndrfeinj (30) -> cpxkqg, donza\nspdyx (189) -> tehdy, afckhxn\ndqxwazb (54)\nvblyk (160) -> zjklq, wswjfu\nmcpyfmy (32)\nxzormv (5)\ncwhjt (41)\ndrigjl (6) -> xxtdti, ewods, dkysxdp\nfmjfhoo (85)\ntliov (52)\ngjxgd (371) -> hgdbip, eikjv, keczktv, mhblqw\ncpdtw (27)\nxaxuyg (372) -> bvfdo, aoltxw\ndpavyu (64)\negyfwlq (133) -> cbtgtqb, qtfuwru, mnjirpy\nllrntn (106) -> rcnhfcz, nqzvz, wolyko, jtjzsr\nfjdyms (56)\nsjfue (11700) -> vcyfwm, impqm, llrntn, tpbeg\nnmbql (250) -> bqfor, vatqulm\nudzzv (90)\ndsrostg (232) -> xriiu, tmfyif\nihskp (42) -> wltlu, brtmsqe\nzsxjxs (23)\nlqfsw (98)\nqsguy (252) -> drefe, ylpahav, dgscbm, emznzp\nilidonx (70)\nhekhbd (411)\nzodbq (79)\nfjueoko (105) -> jkmernz, exgsu, tlulph\nidqaieo (88) -> sfgatoz, dvvkzv, auvkb\nudoiqj (914) -> sbfhxsx, fbmalu, husns\nfbbkkxh (31)\nnzbknfm (84) -> thrmrw, cwdoo, zwcucws, knoziw\nhixtw (582) -> rtlde, hhmrwj, lcbljqb, umjwkfc\nifegka (22)\nbnmgqen (33)\nyodnj (320) -> pfvoay, cflrkd\nhiuuk (19)\nmlafk (38) -> wqgczyl, wnrqkg, wibgm, tpnvx\ndrzoi (68) -> lgltgc, cpess\ndqoatul (42) -> fcfbxmb, qququw\nbibly (37)\nopctg (58)\nabwmmuw (53)\nucsvcy (60)\nanvxtbs (97) -> jfqlxa, mhhyoxf\nhwevm (11)\nqiinez (107) -> wpezq, sigxauh, rpxhqmz, xxlwrd, ffgplgk, hvsywx, vekdbjs\nrhnota (37)\nmrlnt (26)\nytkmda (1825) -> ctien, ogjqrb, goqdmgy\nuanby (86)\noxxse (30)\njkknc (44) -> uyaarr, wwcqz, ruvmsg, tdfgfqn\nxxxxbnh (299)\nfefnhby (1167) -> liwqqqi, bftoy, covpz, bhvvs, gcglasa, rnqvaa\nxsqvwa (274)\nxmgank (374) -> dexabuc, xzormv\nxxlwrd (194) -> zsxjxs, mqyipy\nyyyel (68)\nfvfehj (49) -> tvisu, hyeetiq, kbsmiq, brqvmh\nivowf (98)\nisfyuxu (1346) -> ygldlhw, jroea, rpjiwuz\ncqnmo (50)\nmxxecyh (38)\naaqsdv (31)\ndkysxdp (83)\nchsuzvr (21)\nvjbjy (52)\npyzwl (90)\nizpxbq (7) -> wennxvo, fwgdynj\nfkrdkav (94)\nkmkqsvl (10)\noswqr (81)\npvatk (18) -> nhwgy, rnyarzh, vwisj\ngxdqpcd (188)\nggynw (33)\nacchih (15)\nldrfyi (8)\ntjxkxom (230) -> zwvjv, tobsrs\nvyxgc (50)\nwfkcsb (6251) -> vxjhric, farpozz, eyslsq, zoodui\notbxf (87)\nshwsv (69)\nuxswi (70)\ncvjhub (22)\nzkyiauj (41) -> pqnkr, mnbkaa, xmgank, sadhkpb, tyuyffp\nakioek (53) -> wowgkro, icziaym, nimeytn, hrqgfhg, hxzthym, vzkneq\nemznzp (29)\nhvtsr (98)\nyqzdjj (33) -> dueof, jdbel, ixvxtb\nfqkbscn (418) -> kddex, nylhp, xxyjm, agkgi\nplwdhth (41)\nabgfgav (26)\nwbzwxs (49)\ntlcseu (52)\nvercm (79) -> aejbixg, wbupk\nuvntywl (38)\nhmeolu (5)\nswkor (80) -> matvb, rsocykv\ndeewhp (81)\nbveaw (63)\nqlboef (81) -> gtvpf, qyrhb, xffmho, drpgz, hjjri\ncvipqzv (296) -> xujof, vsffvp\nmvkxva (41)\nosysik (31)\ntkcmrl (16)\nxfxqvk (93)\nsscfkq (720) -> ckbjuyy, zdlxv, kfjbjl\naejbixg (63)\nlndfaxo (97) -> rqsjqrg, onwqp, sdnzxl, erzayz, fistzz\nbcpbir (27)\nfeqefvb (6)\nxsmdo (4571) -> vwwvjg, iibalc, ysgotr, oaabog\nknastam (180) -> zibnjxk, snaebf\nwaxrcby (83) -> tcybjv, xsqvwa, poenou, eyrgmc, azpmzli, vdyftz, jofbj\nbvfdo (51)\nnvhbnlc (59)\nvcvjtf (70)\nmojye (244) -> qojqcd, cxrukf\nqqodh (48)\nahutsc (30)\nvtzfsci (43)\nwennxvo (92)\nykwvu (43)\ntfzsp (126) -> zjmhab, kfdtg\nmogla (38)\nrxfma (8531) -> kxftjs, bpareqw, rlrnjnz\nckdeo (26)\nwkurl (41)\nhuzjje (86)\nlcdleci (344)\nnqxju (90)\ncwxgg (232) -> uylnr, rchhwx, dxgntli\ngfcfsf (56)\nysgnao (76)\nkvtpc (32)\nraaxwq (63)\nfewkfib (965) -> pyzwl, fbhnbdw, nqxju\nyffxaiy (1014) -> piqgu, zkyiauj, xlufcc\nxeefzwg (1740) -> mjgilv, metyk, yonxye\ngnitpo (50)\nnlewmv (38)\ngdhekde (88)\ntyxfxu (9)\nlmuwznh (7)\nchwoyn (18)\nmbmaqh (188)\nxsulc (19)\nhrbsamg (16) -> hhpib, qybyrm, agcpul, azrcn\ntpnvx (284) -> qggycuv, ctafsp, ourium\nuxgbnjk (90)\nnqzvz (13)\nkgfldlb (378) -> ajtjym, xnalv, hwevm\nrosop (294)\nwuogj (411) -> rhnota, gdawyc\nhzxiv (176) -> zedqbr, ymfurx, xfapfj\nzbiypv (1040) -> tlfbyq, idqaieo, adlvipk\ndcytmqp (93)\nsaykbn (91)\ngpikz (41)\nrqsjqrg (1937) -> ctidma, kynrf, oxtckvl\nbzijz (66)\nozxpbo (137) -> bwqxk, bdgwdt, imzegu, dvjssd, cybhcko\npkowhq (1187) -> zfrsmm, tlskukk, fqkbscn, mlafk\nnylhp (171) -> heqvbi, sarwcau\nnwejb (313) -> epuiwg, mculw\nxfapfj (26)\nsntovuq (37)\nnehwool (78) -> knwlsi, zmzwttu, dqxwazb, ebxyi\ndvjssd (98) -> nbccggn, orydf\noaantta (14)\nvzayzaq (17)\nzlubkqa (27)\nxhkrx (51)\ntkqke (1073) -> zjevsui, emsjxz, hgywlzc\nnhhyvm (45)\npokjgkl (8)\ndpvmn (99)\naokvbbi (40)\nbzumy (1522) -> yfkgike, nqata, zheomzg, hooxvp, glrrnl\nbpcrtwl (24)\nirkhv (92)\ntlskukk (1464) -> ixoiuh, jdxth\nnbccggn (91)\nsvzqi (47)\ntcmkx (51)\nclhyd (220) -> ghmxos, zjxcrv\nzdlxv (287) -> eeuohnm, gzocll\neyjdqm (78)\npmlim (78)\nnbfbkao (50)\naeiktt (44)\nrbgonm (9)\nphmmydn (78)\nagkgi (156) -> bibly, bttlx, kirowa\nzqjqic (1602) -> humqvl, drfeinj, xxbwj\naghou (51)\nysgotr (954) -> xxxxbnh, stfqd, ptamwey\nkmvkkze (55)\nglhbvi (64)\nbaimp (158) -> yyyel, mgreg\nxcmre (67)\nmhhbej (36)\ngdawyc (37)\nspdsz (68)\npupaehu (99) -> eolilw, czlmv, zlrvs, vrppl\nvmuqpj (160) -> aghou, aktswyh\nagcpul (98)\nhhuhid (71)\npgkeahi (18)\nnebwj (94)\nglrrnl (52) -> zwake, gtezger\ndtqxiic (21)\npkbxhn (62)\nqququw (73)\nilxgo (237) -> nwkxd, xsulc, weeriq\nymonx (56)\nnqata (122) -> ovgpa, cadrzx\nrahsrp (70)\nwepvw (69)\nkbefh (83)\ncpess (98)\npqnkr (186) -> rnoaimc, pkditpt\nuqzoo (47) -> kitggwr, kywpr, saykbn, pxifep\npgfomf (12)\nvpcfu (54)\nsrqbs (221)\nixoiuh (14)\neozev (63)\nqtqbdp (8)\nzsjoz (73)\nyflpz (40)\nbpgafw (70)\nimzegu (190) -> afpxssv, xdwzfh\nygvwdmg (91)\nvebli (73)\nhshbiaq (91)\nhumqvl (222)\nfozfiqf (52) -> ndqomrl, iwwngux, xaxuyg, cnbvz, wmfro\niukfd (291) -> vcjavrw, mydixq\neduxyw (79)\nygtbl (89)\nfjmjc (3192) -> zkhtua, furny, fewkfib\ncsvxlo (236)\nxzdhzoa (45)\navnosfl (87) -> zsraq, lzstd\nptcby (71) -> dapjhx, hyzzwx, kpwfqfc, ivnufgs\nhvsywx (170) -> tvhftq, xuokqm\ngfsnb (71) -> gtixbu, micfuva\nqxkilb (20)\npkesm (92)\nndjjag (59)\nvsffvp (5)\nieqgh (247) -> zkdwxd, weejuk, yptrqp\njxevrgz (60)\nmnjirpy (24)\nhufrsq (3110) -> cmouh, vbdycf, seicwz, pwdoio, bhwnch, ozxpbo\ntvhftq (35)"
	tempInput := strings.Split(input, "\n")

	allPrograms := make([]Program, 0)
	for _, item := range tempInput {
		var newProgram Program
		firstSplit := strings.Split(item, ") -> ")
		secondSplit := strings.Split(firstSplit[0], " (")
		newProgram.name = secondSplit[0]
		if len(firstSplit) > 1 {
			newProgram.weight, _ = strconv.Atoi(secondSplit[1])
			newProgram.childrenStrings = strings.Split(firstSplit[1], ", ")
		} else {
			number := strings.Split(secondSplit[1], ")")
			newProgram.weight, _ = strconv.Atoi(number[0])
		}
		allPrograms = append(allPrograms, newProgram)
	}
	firstParent := findFirstParent(allPrograms)
	allPrograms = deleteByItem(firstParent, allPrograms)
	firstParent, _ = growTree(firstParent, allPrograms)
	fmt.Println(findDefect(firstParent).weight)
	// fmt.Println(item.name)
			// fmt.Println(rightWeight)
			// fmt.Println(item.weight)
			// fmt.Println(anwser)
}

// 1458
func findDefect(program Program) Program {
	if len(program.children) == 0 {
		return program
	} 
	rightWeight := recurrentSum(program.children[0])
	defect := program.children[0]
	for index, item := range program.children {
		if index != 0 {
			if rightWeight == recurrentSum(item) {
				break
			} else {
				rightWeight = recurrentSum(item)
			}
		}
	}
	for _, item := range program.children {
		if rightWeight != recurrentSum(item) {
			defect = item
			break
			// anwser := item.weight - (recurrentSum(item) - rightWeight)
		}
	}
	if recurrentSum(defect) != rightWeight {
		return findDefect(defect)
	}
	return program
}

// func findDifference(program Program) Program {
// 	rightWeight := recurrentSum(program.children[0])
// 	defect := program.children[0]
// 	for index, item := range program.children {
// 		if index != 0 {
// 			if rightWeight == recurrentSum(item) {
// 				break
// 			} else {
// 				rightWeight = recurrentSum(item)
// 			}
// 		}
// 	}
// 	for _, item := range program.children {
// 		if rightWeight != recurrentSum(item) {
// 			defect = item
// 			break
// 		}
// 	}
// 	return recurrentSum(defect) - rightWeight
// }

func recurrentSum(program Program) int {
	sum := program.weight
	if len(program.children) != 0 {
		for _, child := range program.children {
			sum += recurrentSum(child)
		}
	}
	return sum
}


func growTree(parent Program, allPrograms []Program) (Program, []Program) {
	if len(parent.childrenStrings) == 0 {
		return parent, allPrograms
	}
	for _, itemParent := range parent.childrenStrings {
		for _, item := range allPrograms {
			if item.name == itemParent {
				clearedAllPrograms := deleteByItem(item, allPrograms)
				var childToAppend Program
				childToAppend, allPrograms = growTree(item, clearedAllPrograms)
				parent.children = append(parent.children, childToAppend)
			}
		}
	}
	return parent, allPrograms
}

func deleteByItem(itemToDelete Program, array []Program) []Program {
	indexToDelete := 0
	for index, item := range array {
		if itemToDelete.name == item.name {
			indexToDelete = index
			break
		}
	}
	array = array[:indexToDelete+copy(array[indexToDelete:], array[indexToDelete+1:])]
	return array
}

func findFirstParent(allPrograms []Program) Program {
	var lastParent Program
	currentParent := allPrograms[0]
	for !(lastParent.name == currentParent.name) {
		lastParent = currentParent
		currentParent = findParent(allPrograms, currentParent)
	}
	return currentParent
}

func findParent(allPrograms []Program, currentParent Program) Program {
	for _, item := range allPrograms {
		if currentParent.name != item.name && len(item.childrenStrings) > 0 {
			for _, currentChild := range item.childrenStrings {
				if currentChild == currentParent.name {
					return item
				}
			}
		}
	}
	return currentParent
}