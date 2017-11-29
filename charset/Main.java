import static java.lang.Character.MIN_CODE_POINT;
import static java.lang.Character.MAX_CODE_POINT;
import static java.lang.Character.MIN_SURROGATE;
import static java.lang.Character.MAX_SURROGATE;

import java.nio.ByteBuffer;
import java.nio.CharBuffer;
import java.nio.charset.CharacterCodingException;
import java.nio.charset.Charset;
import java.nio.charset.CharsetEncoder;
import java.nio.charset.CodingErrorAction;

// IANA Character Set Registry:
// http://www.iana.org/assignments/character-sets/character-sets.xhtml

// Java Supported Encodings:
// https://docs.oracle.com/javase/8/docs/technotes/guides/intl/encoding.doc.html
// https://docs.oracle.com/javase/9/intl/supported-encodings.htm

// IANA                                                     java.nio        java.io / java.lang
// ----                                                     --------        -------------------
//   #15: (JIS_X0201)                                       JIS_X0201       JIS_X0201
//   #16: (JIS_Encoding)                                    .               .
//   #17: Shift_JIS                                       * Shift_JIS       SJIS
//   #18: EUC-JP                                          * EUC-JP          EUC_JP
//   #19: (Extended_UNIX_Code_Fixed_Width_for_Japanese)     .               .
//   #36: (KS_C_5601-1987)                                  .               .
//   #37: ISO-2022-KR                                     * ISO-2022-KR     ISO2022KR
//   #38: EUC-KR                                          * EUC-KR          EUC_KR
//   #39: ISO-2022-JP                                     * ISO-2022-JP     ISO2022JP
//   #40: ISO-2022-JP-2                                     .               .
//   #41: (JIS_C6220-1969-jp)                               .               .
//   #42: (JIS_C6220-1969-ro)                               .               .
//   #57: GB_2312-80                                        .               .
//   #67: (JIS_C6229-1984-a)                                .               .
//   #68: (JIS_C6229-1984-b)                                .               .
//   #69: (JIS_C6229-1984-b-add)                            .               .
//   #70: (JIS_C6229-1984-hand)                             .               .
//   #71: (JIS_C6229-1984-hand-add)                         .               .
//   #72: (JIS_C6229-1984-kana)                             .               .
//   #98: (JIS_X0212-1990)                                  JIS_X0212-1990   JIS_X0212-1990
//  #102: (KSC5636)                                         .               .
//  #104: (ISO-2022-CN)                                     ISO-2022-CN     ISO2022CN
//  #105: (ISO-2022-CN-EXT)                                 .               .
//  #113: GBK                                             * GBK             GBK
//  #114: GB18030                                         * GB18030         GB18030
// #1004: (ISO-10646-J-1)                                   .               .
// #2024: (Windows-31J)                                     windows-31j     MS932
// #2025: GB2312                                          * GB2312          EUC_CN
// #2026: Big5                                            * Big5            Big5
// #2036: (IBM281)                                          .               .
// #2039: (IBM290)                                          .               .
// #2101: (Big5-HKSCS)                                    * Big5-HKSCS      Big5_HKSCS
//
//      x-Big5-Solaris   Big5_Solaris
//      x-euc-jp-linux   EUC_JP_LINUX
//      x-EUC-TW         EUC_TW
//      x-eucJP-Open     EUC_JP_Solaris
//      x-IBM1381        Cp1381
//      x-IBM1383        Cp1383
//      x-IBM33722       Cp33722
//      x-IBM834         Cp834
//      x-IBM930         Cp930
//      x-IBM933         Cp933
//      x-IBM935         Cp935
//      x-IBM937         Cp937
//      x-IBM939         Cp939
//      x-IBM942         Cp942
//      x-IBM942C        Cp942C
//      x-IBM943         Cp943
//      x-IBM943C        Cp943C
//      x-IBM948         Cp948
//      x-IBM949         Cp949
//      x-IBM949C        Cp949C
//      x-IBM950         Cp950
//      x-IBM964         Cp964
//      x-IBM970         Cp970
//      x-ISO2022-CN-CNS ISO2022_CN_CNS
//      x-ISO2022-CN-GB  ISO2022_CN_GB
//      x-JIS0208        x-JIS0208
//      x-JISAutoDetect  JISAutoDetect
//      x-Johab          x-Johab
//      x-MS950-HKSCS    MS950_HKSCS
//      x-mswin-936      MS936
//      x-PCK            PCK
//    * x-SJIS_0213      x-SJIS_0213
//      x-windows-949    MS949
//      x-windows-950    MS950
//      x-windows-iso2022jp  x-windows-iso2022jp
//

public class Main {

    public static final String[] CHARSETS = new String[] {
        "Big5",
        "Big5-HKSCS",
        "GB2312",
        "GBK",
        // "GB18030",       // The result seems to indicate that it can encode the same range as Unicode
        // "ISO-2022-CN"    // Cannot do "newEncoder()", java.lang.UnsupportedOperationException
        "JIS_X0201",
        "Shift_JIS",
        "EUC-JP",
        "ISO-2022-JP",
        "EUC-KR",
        "ISO-2022-KR",
    };

    public static final int FORMAT_MD = 0;
    public static final int FORMAT_HTML = 1;
    public static final int SHOW_CHECK = 0;
    public static final int SHOW_ENCODED = 1;

    public static void main(String[] args) {
        final int format = ("html".equals(args[0]) ? FORMAT_HTML : FORMAT_MD);
        final int show = ("encoded".equals(args[1]) ? SHOW_ENCODED : SHOW_CHECK);

        final Charset[] css = new Charset[CHARSETS.length];
        final CharsetEncoder[] cses = new CharsetEncoder[CHARSETS.length];
        for (int index = 0; index < CHARSETS.length; ++index) {
            css[index] = Charset.forName(CHARSETS[index]);
            cses[index] = css[index].newEncoder()
                .onMalformedInput(CodingErrorAction.REPORT)
                .onUnmappableCharacter(CodingErrorAction.REPORT);
        }

        // Document Header
        if (format == FORMAT_HTML) {
            System.out.println("<!DOCTYPE html>");
            System.out.print("<html><head><title>Charset Table</title></head><body><table border='1'>");
        }

        // Table Header
        final String tableHeaderStart = (format == FORMAT_MD ? "| %s | %s | %s |" : "<tr><th>%s</th><th>%s</th><th>%s</th>");
        final String tableHeader = (format == FORMAT_MD ? " %s |" : "<th>%s</th>");
        final String tableHeaderEnd = (format == FORMAT_MD ? "\n" : "</tr>");
        System.out.printf(tableHeaderStart, "DEC", "HEX", "CHAR");
        for (int index = 0; index < CHARSETS.length; ++index) {
            System.out.printf(tableHeader, CHARSETS[index]);
        }
        System.out.print(tableHeaderEnd);

        // Table Header Separator
        if (format == FORMAT_MD) {
            System.out.print("| --- | --- | --- |");
            for (int index = 0; index < CHARSETS.length; ++index) {
                System.out.print(" --- |");
            }
            System.out.println();
        }

        // Table Table Data
        final String tableDataStart = (format == FORMAT_MD ? "|%1$d|%1$X|&#%1$d;|" : "<tr><td>%1$d</td><td>%1$X</td><td>&#%1$d;</td>");
        final String tableData = (format == FORMAT_MD ? "%s|" : "<td>%s</td>");
        final String tableDataEnd = (format == FORMAT_MD ? "\n" : "</tr>");
        for (int codePoint = MIN_CODE_POINT; codePoint <= MAX_CODE_POINT; ++codePoint) {
            // if (codePoint >= MIN_SURROGATE && codePoint <= MAX_SURROGATE) {
            //     final char[] chars = Character.toChars(codePoint);
            //     continue;
            // }
            final char[] chars = Character.toChars(codePoint);
            final CharBuffer charBuffer = CharBuffer.wrap(chars);
            final boolean[] isEncodeable = new boolean[CHARSETS.length];
            String[] encoded = null;
            boolean hasLanguage = false;
            if (show == SHOW_CHECK) {
                for (int index = 0; index < CHARSETS.length; ++index) {
                    try {
                        final ByteBuffer byteBuffer = cses[index].reset().encode((CharBuffer) charBuffer.rewind());
                        isEncodeable[index] = true;
                        hasLanguage = true;
                    } catch (final CharacterCodingException e) {
                        //System.out.println(codePoint);
                    }
                }
            } else {
                encoded = new String[CHARSETS.length];
                for (int index = 0; index < CHARSETS.length; ++index) {
                    try {
                        final ByteBuffer byteBuffer = cses[index].reset().encode((CharBuffer) charBuffer.rewind());
                        final int limit = byteBuffer.limit();
                        final StringBuilder sb = new StringBuilder();
                        for (int index2 = 0; index2 < limit; ++index2) {
                            final byte b = byteBuffer.get();
                            final String s = String.format("%02X", b);
                            sb.append(s);
                        }
                        encoded[index] = sb.toString();
                        isEncodeable[index] = true;
                        hasLanguage = true;
                    } catch (final CharacterCodingException e) {
                        //System.out.println(codePoint);
                    }
                }
            }
            if (hasLanguage) {
                System.out.printf(tableDataStart, codePoint);
                for (int index = 0; index < CHARSETS.length; ++index) {
                    if (isEncodeable[index]) {
                        System.out.printf(tableData, (show == SHOW_CHECK) ? "✓" : encoded[index]);
                    } else {
                        System.out.printf(tableData, "");
                    }
                }
                System.out.print(tableDataEnd);
            }
        }

        // Document Header
        if (format == FORMAT_HTML) {
            System.out.print("</table></body>");
        }
    }
}
