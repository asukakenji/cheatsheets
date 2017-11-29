import static java.lang.Character.MIN_CODE_POINT;
import static java.lang.Character.MAX_CODE_POINT;
import static java.lang.Character.MIN_SURROGATE;
import static java.lang.Character.MAX_SURROGATE;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.nio.CharBuffer;
import java.nio.charset.CharacterCodingException;
import java.nio.charset.Charset;
import java.nio.charset.CharsetEncoder;
import java.nio.charset.CodingErrorAction;
import java.util.Formatter;

// Java Supported Encodings:
// https://docs.oracle.com/javase/8/docs/technotes/guides/intl/encoding.doc.html
// https://docs.oracle.com/javase/9/intl/supported-encodings.htm

// Common Encodings:
// - UTF-8
// - UTF-16, UTF-16BE, UTF-16LE, x-UTF-16BE-BOM (not available), x-UTF-16LE-BOM
// - UTF-32, UTF-32BE, UTF-32LE, x-UTF-32BE-BOM, x-UTF-32LE-BOM
// - Big5, Big5-HKSCS
public class Main2 {

    private static String byteArrayToHexString(byte[] bytes) {
        final Formatter formatter = new Formatter();
        for (final byte b : bytes) {
            formatter.format("%02X", b);
        }
        return formatter.toString();
    }

    public static void main(String[] args) {
        final String charsetToBePrinted = args[0];
        final String printedInCharset = args[1];

        final Charset inCharset = Charset.forName(charsetToBePrinted);
        final Charset outCharset = Charset.forName(printedInCharset);

        if (inCharset.equals(outCharset)) {
            final CharsetEncoder charsetEncoder = inCharset.newEncoder()
                .onMalformedInput(CodingErrorAction.REPORT)
                .onUnmappableCharacter(CodingErrorAction.REPORT);
            for (int codePoint = MIN_CODE_POINT; codePoint <= MAX_CODE_POINT; ++codePoint) {
                if (codePoint >= MIN_SURROGATE && codePoint <= MAX_SURROGATE) {
                    continue;
                }
                final char[] codePointChars = Character.toChars(codePoint);
                final CharBuffer charBuffer = CharBuffer.wrap(codePointChars);
                try {
                    final ByteBuffer byteBuffer = charsetEncoder.reset().encode(charBuffer.rewind());
                    byte[] bytes = byteBuffer.array();
                    System.out.printf("U+%04x: 0x%s: ", codePoint, byteArrayToHexString(bytes));
                    System.out.write(bytes);
                    System.out.write('\n');
                } catch (final CharacterCodingException e) {
                } catch (final IOException e) {
                }
            }
        } else {
            final CharsetEncoder inCharsetEncoder = inCharset.newEncoder()
                .onMalformedInput(CodingErrorAction.REPORT)
                .onUnmappableCharacter(CodingErrorAction.REPORT);
            final CharsetEncoder outCharsetEncoder = outCharset.newEncoder()
                .onMalformedInput(CodingErrorAction.REPORT)
                .onUnmappableCharacter(CodingErrorAction.REPORT);
            for (int codePoint = MIN_CODE_POINT; codePoint <= MAX_CODE_POINT; ++codePoint) {
                if (codePoint >= MIN_SURROGATE && codePoint <= MAX_SURROGATE) {
                    continue;
                }
                final char[] codePointChars = Character.toChars(codePoint);
                final CharBuffer charBuffer = CharBuffer.wrap(codePointChars);
                try {
                    final ByteBuffer inByteBuffer = inCharsetEncoder.reset().encode(charBuffer.rewind());
                    final ByteBuffer outByteBuffer = outCharsetEncoder.reset().encode(charBuffer.rewind());
                    byte[] inBytes = inByteBuffer.array();
                    byte[] outBytes = outByteBuffer.array();
                    System.out.printf("U+%04x: 0x%s: 0x%s:", codePoint, byteArrayToHexString(inBytes), byteArrayToHexString(outBytes));
                    System.out.write(outBytes);
                    System.out.write('\n');
                } catch (final CharacterCodingException e) {
                } catch (final IOException e) {
                }
            }
        }
    }
}
