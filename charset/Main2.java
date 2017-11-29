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

// Java Supported Encodings:
// https://docs.oracle.com/javase/8/docs/technotes/guides/intl/encoding.doc.html
// https://docs.oracle.com/javase/9/intl/supported-encodings.htm

public class Main2 {
    public static void main(String[] args) {
        final Charset charset = Charset.forName(args[0]);
        final CharsetEncoder charsetEncoder = charset.newEncoder()
            .onMalformedInput(CodingErrorAction.REPORT)
            .onUnmappableCharacter(CodingErrorAction.REPORT);

        for (int codePoint = MIN_CODE_POINT; codePoint <= MAX_CODE_POINT; ++codePoint) {
            if (codePoint >= MIN_SURROGATE && codePoint <= MAX_SURROGATE) {
                continue;
            }
            final char[] codePointChars = Character.toChars(codePoint);
            final CharBuffer charBuffer = CharBuffer.wrap(codePointChars);
            try {
                final ByteBuffer byteBuffer = charsetEncoder.reset().encode(charBuffer);
                byte[] bytes = byteBuffer.array();
                System.out.printf("U+%04x: ", codePoint);
                System.out.print("0x");
                for (final byte b : bytes) {
                    System.out.printf("%02X", b);
                }
                System.out.print(": ");
                System.out.write(bytes);
                System.out.write('\n');
            } catch (final CharacterCodingException e) {
            } catch (final IOException e) {
            }
        }
    }
}
