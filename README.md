# go-snum
golang string number utility
implement ericlagergren/decimal ( github.com/ericlagergren/decimal )

# range limit
1. bt__sorted
1-1. 양수 : 65 byte (header 1 bt + body 64 bt)
1-2. 음수 : 66 byte (header 1 bt + body 64 bt + 0xFF 1 bt)

2. bt__unsorted
 2-1. 양수 : 55 byte (header 1 bt + body 54 bt)
 2-2. 음수 : 55 byte (header 1 bt + body 54 bt)


# 헤더 구성
 - 255 = 정수 자릿수가 96 자리인 양수 ( 1e95 <= x < 1e96 )
 - 254 = 정수 자릿수가 95 자리인 양수 ( 1e94 <= x < 1e95 )
 ..... ( 양의 정수 )
 - 160 = 정수 자릿수가 1 자리인 양수 ( 1 <= x < 10 )
 - 159 = 정수 자릿수가 -1 자리인 양수 ( 0.1 <= x < 1 )
 ..... ( 양의 소수 )
 - 128 = 정수 자릿수가 -32 자리인 양수 ( 0 <= x < 1e-31 ) - !! 0 포함 !!

 - 127 = 정수 자릿수가 -32 자리인 음수 ( -1e-31 < x <= -1e-32 )
 ..... ( 음의 소수 )
 - 96 = 정수 자릿수가 -1 자리인 음수 ( -1 < x <= -0.1 )
 - 95 = 정수 자릿수가 1 자리인 음수 ( -10 < x <= -1 )
 ..... ( 음의 정수 )
 - 1 = 정수 자릿수가 95 자리인 음수 ( -1e95 < x <= -1e94 )
 - 0 = 정수 자릿수가 96 자리인 음수 ( -1e96 < x <= -1e95 )

buf[1:]
 - 정수 + 소수 big.Int 에 담아 2자릿수 당 1바이트 로 압축한 byte array
 - 음수일 경우 보수로 저장
