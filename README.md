[eh21147@it106357 tests]$ go test ./test -v -run BenchmarkGol
# ./test
stat /home/eh21147/Downloads/CSGO-Game-of-Life-PAL-main/PAL221/tests/test: directory not found
FAIL    ./test [setup failed]
FAIL
[eh21147@it106357 tests]$ go test -bench="." -v
=== RUN   TestAlive
    alive_test.go:94: Alive Cells 6878
    alive_test.go:94: Alive Cells 6008
    alive_test.go:94: Alive Cells 6072
    alive_test.go:94: Alive Cells 5567
    alive_test.go:94: Alive Cells 5565
--- PASS: TestAlive (10.03s)
=== RUN   TestGol
=== RUN   TestGol/16x16x0-1
=== RUN   TestGol/16x16x0-2
=== RUN   TestGol/16x16x0-3
=== RUN   TestGol/16x16x0-4
=== RUN   TestGol/16x16x0-5
=== RUN   TestGol/16x16x0-6
=== RUN   TestGol/16x16x0-7
=== RUN   TestGol/16x16x0-8
=== RUN   TestGol/16x16x0-9
=== RUN   TestGol/16x16x0-10
=== RUN   TestGol/16x16x0-11
=== RUN   TestGol/16x16x0-12
=== RUN   TestGol/16x16x0-13
=== RUN   TestGol/16x16x0-14
=== RUN   TestGol/16x16x0-15
=== RUN   TestGol/16x16x0-16
=== RUN   TestGol/16x16x1-1
=== RUN   TestGol/16x16x1-2
=== RUN   TestGol/16x16x1-3
=== RUN   TestGol/16x16x1-4
=== RUN   TestGol/16x16x1-5
=== RUN   TestGol/16x16x1-6
=== RUN   TestGol/16x16x1-7
=== RUN   TestGol/16x16x1-8
=== RUN   TestGol/16x16x1-9
=== RUN   TestGol/16x16x1-10
=== RUN   TestGol/16x16x1-11
=== RUN   TestGol/16x16x1-12
=== RUN   TestGol/16x16x1-13
=== RUN   TestGol/16x16x1-14
=== RUN   TestGol/16x16x1-15
=== RUN   TestGol/16x16x1-16
=== RUN   TestGol/16x16x100-1
=== RUN   TestGol/16x16x100-2
=== RUN   TestGol/16x16x100-3
=== RUN   TestGol/16x16x100-4
=== RUN   TestGol/16x16x100-5
=== RUN   TestGol/16x16x100-6
=== RUN   TestGol/16x16x100-7
=== RUN   TestGol/16x16x100-8
=== RUN   TestGol/16x16x100-9
=== RUN   TestGol/16x16x100-10
=== RUN   TestGol/16x16x100-11
=== RUN   TestGol/16x16x100-12
=== RUN   TestGol/16x16x100-13
=== RUN   TestGol/16x16x100-14
=== RUN   TestGol/16x16x100-15
=== RUN   TestGol/16x16x100-16
=== RUN   TestGol/64x64x0-1
=== RUN   TestGol/64x64x0-2
=== RUN   TestGol/64x64x0-3
=== RUN   TestGol/64x64x0-4
=== RUN   TestGol/64x64x0-5
=== RUN   TestGol/64x64x0-6
=== RUN   TestGol/64x64x0-7
=== RUN   TestGol/64x64x0-8
=== RUN   TestGol/64x64x0-9
=== RUN   TestGol/64x64x0-10
=== RUN   TestGol/64x64x0-11
=== RUN   TestGol/64x64x0-12
=== RUN   TestGol/64x64x0-13
=== RUN   TestGol/64x64x0-14
=== RUN   TestGol/64x64x0-15
=== RUN   TestGol/64x64x0-16
=== RUN   TestGol/64x64x1-1
=== RUN   TestGol/64x64x1-2
=== RUN   TestGol/64x64x1-3
=== RUN   TestGol/64x64x1-4
=== RUN   TestGol/64x64x1-5
=== RUN   TestGol/64x64x1-6
=== RUN   TestGol/64x64x1-7
=== RUN   TestGol/64x64x1-8
=== RUN   TestGol/64x64x1-9
=== RUN   TestGol/64x64x1-10
=== RUN   TestGol/64x64x1-11
=== RUN   TestGol/64x64x1-12
=== RUN   TestGol/64x64x1-13
=== RUN   TestGol/64x64x1-14
=== RUN   TestGol/64x64x1-15
=== RUN   TestGol/64x64x1-16
=== RUN   TestGol/64x64x100-1
=== RUN   TestGol/64x64x100-2
=== RUN   TestGol/64x64x100-3
=== RUN   TestGol/64x64x100-4
=== RUN   TestGol/64x64x100-5
=== RUN   TestGol/64x64x100-6
=== RUN   TestGol/64x64x100-7
=== RUN   TestGol/64x64x100-8
=== RUN   TestGol/64x64x100-9
=== RUN   TestGol/64x64x100-10
=== RUN   TestGol/64x64x100-11
=== RUN   TestGol/64x64x100-12
=== RUN   TestGol/64x64x100-13
=== RUN   TestGol/64x64x100-14
=== RUN   TestGol/64x64x100-15
=== RUN   TestGol/64x64x100-16
=== RUN   TestGol/512x512x0-1
=== RUN   TestGol/512x512x0-2
=== RUN   TestGol/512x512x0-3
=== RUN   TestGol/512x512x0-4
=== RUN   TestGol/512x512x0-5
=== RUN   TestGol/512x512x0-6
=== RUN   TestGol/512x512x0-7
=== RUN   TestGol/512x512x0-8
=== RUN   TestGol/512x512x0-9
=== RUN   TestGol/512x512x0-10
=== RUN   TestGol/512x512x0-11
=== RUN   TestGol/512x512x0-12
=== RUN   TestGol/512x512x0-13
=== RUN   TestGol/512x512x0-14
=== RUN   TestGol/512x512x0-15
=== RUN   TestGol/512x512x0-16
=== RUN   TestGol/512x512x1-1
=== RUN   TestGol/512x512x1-2
=== RUN   TestGol/512x512x1-3
=== RUN   TestGol/512x512x1-4
=== RUN   TestGol/512x512x1-5
=== RUN   TestGol/512x512x1-6
=== RUN   TestGol/512x512x1-7
=== RUN   TestGol/512x512x1-8
=== RUN   TestGol/512x512x1-9
=== RUN   TestGol/512x512x1-10
=== RUN   TestGol/512x512x1-11
=== RUN   TestGol/512x512x1-12
=== RUN   TestGol/512x512x1-13
=== RUN   TestGol/512x512x1-14
=== RUN   TestGol/512x512x1-15
=== RUN   TestGol/512x512x1-16
=== RUN   TestGol/512x512x100-1
=== RUN   TestGol/512x512x100-2
=== RUN   TestGol/512x512x100-3
=== RUN   TestGol/512x512x100-4
=== RUN   TestGol/512x512x100-5
=== RUN   TestGol/512x512x100-6
=== RUN   TestGol/512x512x100-7
=== RUN   TestGol/512x512x100-8
=== RUN   TestGol/512x512x100-9
=== RUN   TestGol/512x512x100-10
=== RUN   TestGol/512x512x100-11
=== RUN   TestGol/512x512x100-12
=== RUN   TestGol/512x512x100-13
=== RUN   TestGol/512x512x100-14
=== RUN   TestGol/512x512x100-15
=== RUN   TestGol/512x512x100-16
--- PASS: TestGol (37.21s)
    --- PASS: TestGol/16x16x0-1 (0.00s)
    --- PASS: TestGol/16x16x0-2 (0.00s)
    --- PASS: TestGol/16x16x0-3 (0.00s)
    --- PASS: TestGol/16x16x0-4 (0.00s)
    --- PASS: TestGol/16x16x0-5 (0.00s)
    --- PASS: TestGol/16x16x0-6 (0.00s)
    --- PASS: TestGol/16x16x0-7 (0.00s)
    --- PASS: TestGol/16x16x0-8 (0.00s)
    --- PASS: TestGol/16x16x0-9 (0.00s)
    --- PASS: TestGol/16x16x0-10 (0.00s)
    --- PASS: TestGol/16x16x0-11 (0.00s)
    --- PASS: TestGol/16x16x0-12 (0.00s)
    --- PASS: TestGol/16x16x0-13 (0.00s)
    --- PASS: TestGol/16x16x0-14 (0.00s)
    --- PASS: TestGol/16x16x0-15 (0.00s)
    --- PASS: TestGol/16x16x0-16 (0.00s)
    --- PASS: TestGol/16x16x1-1 (0.00s)
    --- PASS: TestGol/16x16x1-2 (0.00s)
    --- PASS: TestGol/16x16x1-3 (0.00s)
    --- PASS: TestGol/16x16x1-4 (0.00s)
    --- PASS: TestGol/16x16x1-5 (0.00s)
    --- PASS: TestGol/16x16x1-6 (0.00s)
    --- PASS: TestGol/16x16x1-7 (0.00s)
    --- PASS: TestGol/16x16x1-8 (0.00s)
    --- PASS: TestGol/16x16x1-9 (0.00s)
    --- PASS: TestGol/16x16x1-10 (0.00s)
    --- PASS: TestGol/16x16x1-11 (0.00s)
    --- PASS: TestGol/16x16x1-12 (0.00s)
    --- PASS: TestGol/16x16x1-13 (0.00s)
    --- PASS: TestGol/16x16x1-14 (0.00s)
    --- PASS: TestGol/16x16x1-15 (0.00s)
    --- PASS: TestGol/16x16x1-16 (0.00s)
    --- PASS: TestGol/16x16x100-1 (0.00s)
    --- PASS: TestGol/16x16x100-2 (0.00s)
    --- PASS: TestGol/16x16x100-3 (0.00s)
    --- PASS: TestGol/16x16x100-4 (0.00s)
    --- PASS: TestGol/16x16x100-5 (0.00s)
    --- PASS: TestGol/16x16x100-6 (0.00s)
    --- PASS: TestGol/16x16x100-7 (0.00s)
    --- PASS: TestGol/16x16x100-8 (0.00s)
    --- PASS: TestGol/16x16x100-9 (0.00s)
    --- PASS: TestGol/16x16x100-10 (0.00s)
    --- PASS: TestGol/16x16x100-11 (0.00s)
    --- PASS: TestGol/16x16x100-12 (0.00s)
    --- PASS: TestGol/16x16x100-13 (0.00s)
    --- PASS: TestGol/16x16x100-14 (0.00s)
    --- PASS: TestGol/16x16x100-15 (0.00s)
    --- PASS: TestGol/16x16x100-16 (0.00s)
    --- PASS: TestGol/64x64x0-1 (0.01s)
    --- PASS: TestGol/64x64x0-2 (0.01s)
    --- PASS: TestGol/64x64x0-3 (0.01s)
    --- PASS: TestGol/64x64x0-4 (0.01s)
    --- PASS: TestGol/64x64x0-5 (0.01s)
    --- PASS: TestGol/64x64x0-6 (0.01s)
    --- PASS: TestGol/64x64x0-7 (0.01s)
    --- PASS: TestGol/64x64x0-8 (0.01s)
    --- PASS: TestGol/64x64x0-9 (0.01s)
    --- PASS: TestGol/64x64x0-10 (0.01s)
    --- PASS: TestGol/64x64x0-11 (0.01s)
    --- PASS: TestGol/64x64x0-12 (0.01s)
    --- PASS: TestGol/64x64x0-13 (0.01s)
    --- PASS: TestGol/64x64x0-14 (0.01s)
    --- PASS: TestGol/64x64x0-15 (0.01s)
    --- PASS: TestGol/64x64x0-16 (0.01s)
    --- PASS: TestGol/64x64x1-1 (0.01s)
    --- PASS: TestGol/64x64x1-2 (0.01s)
    --- PASS: TestGol/64x64x1-3 (0.01s)
    --- PASS: TestGol/64x64x1-4 (0.01s)
    --- PASS: TestGol/64x64x1-5 (0.01s)
    --- PASS: TestGol/64x64x1-6 (0.01s)
    --- PASS: TestGol/64x64x1-7 (0.01s)
    --- PASS: TestGol/64x64x1-8 (0.01s)
    --- PASS: TestGol/64x64x1-9 (0.01s)
    --- PASS: TestGol/64x64x1-10 (0.01s)
    --- PASS: TestGol/64x64x1-11 (0.01s)
    --- PASS: TestGol/64x64x1-12 (0.01s)
    --- PASS: TestGol/64x64x1-13 (0.01s)
    --- PASS: TestGol/64x64x1-14 (0.01s)
    --- PASS: TestGol/64x64x1-15 (0.01s)
    --- PASS: TestGol/64x64x1-16 (0.01s)
    --- PASS: TestGol/64x64x100-1 (0.03s)
    --- PASS: TestGol/64x64x100-2 (0.02s)
    --- PASS: TestGol/64x64x100-3 (0.02s)
    --- PASS: TestGol/64x64x100-4 (0.02s)
    --- PASS: TestGol/64x64x100-5 (0.02s)
    --- PASS: TestGol/64x64x100-6 (0.02s)
    --- PASS: TestGol/64x64x100-7 (0.02s)
    --- PASS: TestGol/64x64x100-8 (0.02s)
    --- PASS: TestGol/64x64x100-9 (0.02s)
    --- PASS: TestGol/64x64x100-10 (0.02s)
    --- PASS: TestGol/64x64x100-11 (0.02s)
    --- PASS: TestGol/64x64x100-12 (0.02s)
    --- PASS: TestGol/64x64x100-13 (0.02s)
    --- PASS: TestGol/64x64x100-14 (0.02s)
    --- PASS: TestGol/64x64x100-15 (0.02s)
    --- PASS: TestGol/64x64x100-16 (0.02s)
    --- PASS: TestGol/512x512x0-1 (0.53s)
    --- PASS: TestGol/512x512x0-2 (0.52s)
    --- PASS: TestGol/512x512x0-3 (0.48s)
    --- PASS: TestGol/512x512x0-4 (0.46s)
    --- PASS: TestGol/512x512x0-5 (0.48s)
    --- PASS: TestGol/512x512x0-6 (0.51s)
    --- PASS: TestGol/512x512x0-7 (0.48s)
    --- PASS: TestGol/512x512x0-8 (0.51s)
    --- PASS: TestGol/512x512x0-9 (0.51s)
    --- PASS: TestGol/512x512x0-10 (0.43s)
    --- PASS: TestGol/512x512x0-11 (0.46s)
    --- PASS: TestGol/512x512x0-12 (0.44s)
    --- PASS: TestGol/512x512x0-13 (0.44s)
    --- PASS: TestGol/512x512x0-14 (0.42s)
    --- PASS: TestGol/512x512x0-15 (0.43s)
    --- PASS: TestGol/512x512x0-16 (0.45s)
    --- PASS: TestGol/512x512x1-1 (0.51s)
    --- PASS: TestGol/512x512x1-2 (0.47s)
    --- PASS: TestGol/512x512x1-3 (0.52s)
    --- PASS: TestGol/512x512x1-4 (0.51s)
    --- PASS: TestGol/512x512x1-5 (0.48s)
    --- PASS: TestGol/512x512x1-6 (0.47s)
    --- PASS: TestGol/512x512x1-7 (0.45s)
    --- PASS: TestGol/512x512x1-8 (0.49s)
    --- PASS: TestGol/512x512x1-9 (0.53s)
    --- PASS: TestGol/512x512x1-10 (0.46s)
    --- PASS: TestGol/512x512x1-11 (0.44s)
    --- PASS: TestGol/512x512x1-12 (10.93s)
    --- PASS: TestGol/512x512x1-13 (0.43s)
    --- PASS: TestGol/512x512x1-14 (0.42s)
    --- PASS: TestGol/512x512x1-15 (0.43s)
    --- PASS: TestGol/512x512x1-16 (0.42s)
    --- PASS: TestGol/512x512x100-1 (1.28s)
    --- PASS: TestGol/512x512x100-2 (0.88s)
    --- PASS: TestGol/512x512x100-3 (0.75s)
    --- PASS: TestGol/512x512x100-4 (0.70s)
    --- PASS: TestGol/512x512x100-5 (0.67s)
    --- PASS: TestGol/512x512x100-6 (0.64s)
    --- PASS: TestGol/512x512x100-7 (0.62s)
    --- PASS: TestGol/512x512x100-8 (0.60s)
    --- PASS: TestGol/512x512x100-9 (0.60s)
    --- PASS: TestGol/512x512x100-10 (0.59s)
    --- PASS: TestGol/512x512x100-11 (0.60s)
    --- PASS: TestGol/512x512x100-12 (0.59s)
    --- PASS: TestGol/512x512x100-13 (0.59s)
    --- PASS: TestGol/512x512x100-14 (0.59s)
    --- PASS: TestGol/512x512x100-15 (0.57s)
    --- PASS: TestGol/512x512x100-16 (0.58s)
=== RUN   TestKeyboard
=== RUN   TestKeyboard/p
    util_test.go:409: Testing for first StateChange Executing event
    util_test.go:278: [Event] Completed Turns 0        Executing
    util_test.go:460: Testing for StateChange Paused event
    util_test.go:275: [Event] Completed Turns 20       Final Turn Complete
    keyboard_test.go:45: ERROR Your program has returned from the gol.Run function before it was unpaused
    util_test.go:278: [Event] Completed Turns 20       Quitting
    util_test.go:637: ERROR No StateChange Paused events received in 2 seconds
    util_test.go:438: Testing for StateChange Executing event
    util_test.go:637: ERROR No StateChange Executing events received in 2 seconds
=== RUN   TestKeyboard/s
    util_test.go:409: Testing for first StateChange Executing event
    util_test.go:278: [Event] Completed Turns 0        Executing
    util_test.go:568: Testing image output
    util_test.go:265: [Event] Completed Turns 1349     Alive Cells 6599     Avg  665 turns/sec
    util_test.go:265: [Event] Completed Turns 2710     Alive Cells 5932     Avg  673 turns/sec
    util_test.go:637: ERROR No ImageOutput events received in 4 seconds
                        If tests are running in WSL2, please make sure project is located within WSL2 file system rather than Windows!
                        i.e. Your path must not start with /mnt/...
    util_test.go:652: WARN Your program has not returned from the gol.Run function
                                                Continuing with other tests, leaving your program executing
                                                You may get unexpected behaviour
=== RUN   TestKeyboard/q
    util_test.go:409: Testing for first StateChange Executing event
    util_test.go:278: [Event] Completed Turns 0        Executing
    util_test.go:568: Testing image output
    util_test.go:265: [Event] Completed Turns 1140     Alive Cells 6845     Avg  559 turns/sec
    util_test.go:265: [Event] Completed Turns 2253     Alive Cells 6708     Avg  558 turns/sec
    util_test.go:637: ERROR No ImageOutput events received in 4 seconds
                        If tests are running in WSL2, please make sure project is located within WSL2 file system rather than Windows!
                        i.e. Your path must not start with /mnt/...
    util_test.go:524: Testing for StateChange Quitting event
    util_test.go:265: [Event] Completed Turns 3340     Alive Cells 6184     Avg  553 turns/sec
    util_test.go:637: ERROR No StateChange Quitting events received in 2 seconds
    util_test.go:637: ERROR Your program has not returned from the gol.Run function
                                                Continuing with other tests, leaving your program executing
                                                You may get unexpected behaviour
=== RUN   TestKeyboard/p+s
    util_test.go:409: Testing for first StateChange Executing event
    util_test.go:278: [Event] Completed Turns 0        Executing
    util_test.go:460: Testing for StateChange Paused event
    util_test.go:265: [Event] Completed Turns 643      Alive Cells 6469     Avg  313 turns/sec
    util_test.go:637: ERROR No StateChange Paused events received in 2 seconds
    util_test.go:568: Testing image output
    util_test.go:265: [Event] Completed Turns 1299     Alive Cells 6821     Avg  320 turns/sec
    util_test.go:265: [Event] Completed Turns 1954     Alive Cells 6686     Avg  323 turns/sec
    util_test.go:637: ERROR No ImageOutput events received in 4 seconds
                        If tests are running in WSL2, please make sure project is located within WSL2 file system rather than Windows!
                        i.e. Your path must not start with /mnt/...
    util_test.go:652: WARN Your program has not returned from the gol.Run function
                                                Continuing with other tests, leaving your program executing
                                                You may get unexpected behaviour
=== RUN   TestKeyboard/p+q
    util_test.go:409: Testing for first StateChange Executing event
    util_test.go:278: [Event] Completed Turns 0        Executing
    util_test.go:460: Testing for StateChange Paused event
    util_test.go:265: [Event] Completed Turns 556      Alive Cells 6691     Avg  270 turns/sec
    util_test.go:637: ERROR No StateChange Paused events received in 2 seconds
    util_test.go:568: Testing image output
    util_test.go:265: [Event] Completed Turns 1110     Alive Cells 6776     Avg  273 turns/sec
    util_test.go:265: [Event] Completed Turns 1678     Alive Cells 6428     Avg  277 turns/sec
    util_test.go:637: ERROR No ImageOutput events received in 4 seconds
                        If tests are running in WSL2, please make sure project is located within WSL2 file system rather than Windows!
                        i.e. Your path must not start with /mnt/...
    util_test.go:524: Testing for StateChange Quitting event
    util_test.go:265: [Event] Completed Turns 2230     Alive Cells 6650     Avg  279 turns/sec
    util_test.go:637: ERROR No StateChange Quitting events received in 2 seconds
    util_test.go:637: ERROR Your program has not returned from the gol.Run function
                                                Continuing with other tests, leaving your program executing
                                                You may get unexpected behaviour
--- FAIL: TestKeyboard (43.23s)
    --- FAIL: TestKeyboard/p (9.03s)
    --- FAIL: TestKeyboard/s (6.53s)
    --- FAIL: TestKeyboard/q (8.54s)
    --- FAIL: TestKeyboard/p+s (8.56s)
    --- FAIL: TestKeyboard/p+q (10.56s)
=== RUN   TestPgm
=== RUN   TestPgm/16x16x0-1
=== RUN   TestPgm/16x16x0-2
=== RUN   TestPgm/16x16x0-3
=== RUN   TestPgm/16x16x0-4
=== RUN   TestPgm/16x16x0-5
=== RUN   TestPgm/16x16x0-6
=== RUN   TestPgm/16x16x0-7
=== RUN   TestPgm/16x16x0-8
=== RUN   TestPgm/16x16x0-9
=== RUN   TestPgm/16x16x0-10
=== RUN   TestPgm/16x16x0-11
=== RUN   TestPgm/16x16x0-12
=== RUN   TestPgm/16x16x0-13
=== RUN   TestPgm/16x16x0-14
=== RUN   TestPgm/16x16x0-15
=== RUN   TestPgm/16x16x0-16
=== RUN   TestPgm/16x16x1-1
=== RUN   TestPgm/16x16x1-2
=== RUN   TestPgm/16x16x1-3
=== RUN   TestPgm/16x16x1-4
=== RUN   TestPgm/16x16x1-5
=== RUN   TestPgm/16x16x1-6
=== RUN   TestPgm/16x16x1-7
=== RUN   TestPgm/16x16x1-8
=== RUN   TestPgm/16x16x1-9
=== RUN   TestPgm/16x16x1-10
=== RUN   TestPgm/16x16x1-11
=== RUN   TestPgm/16x16x1-12
=== RUN   TestPgm/16x16x1-13
=== RUN   TestPgm/16x16x1-14
=== RUN   TestPgm/16x16x1-15
=== RUN   TestPgm/16x16x1-16
=== RUN   TestPgm/16x16x100-1
=== RUN   TestPgm/16x16x100-2
=== RUN   TestPgm/16x16x100-3
=== RUN   TestPgm/16x16x100-4
=== RUN   TestPgm/16x16x100-5
=== RUN   TestPgm/16x16x100-6
=== RUN   TestPgm/16x16x100-7
=== RUN   TestPgm/16x16x100-8
=== RUN   TestPgm/16x16x100-9
=== RUN   TestPgm/16x16x100-10
=== RUN   TestPgm/16x16x100-11
=== RUN   TestPgm/16x16x100-12
=== RUN   TestPgm/16x16x100-13
=== RUN   TestPgm/16x16x100-14
=== RUN   TestPgm/16x16x100-15
=== RUN   TestPgm/16x16x100-16
=== RUN   TestPgm/64x64x0-1
=== RUN   TestPgm/64x64x0-2
=== RUN   TestPgm/64x64x0-3
=== RUN   TestPgm/64x64x0-4
=== RUN   TestPgm/64x64x0-5
=== RUN   TestPgm/64x64x0-6
=== RUN   TestPgm/64x64x0-7
=== RUN   TestPgm/64x64x0-8
=== RUN   TestPgm/64x64x0-9
=== RUN   TestPgm/64x64x0-10
=== RUN   TestPgm/64x64x0-11
=== RUN   TestPgm/64x64x0-12
=== RUN   TestPgm/64x64x0-13
=== RUN   TestPgm/64x64x0-14
=== RUN   TestPgm/64x64x0-15
=== RUN   TestPgm/64x64x0-16
=== RUN   TestPgm/64x64x1-1
=== RUN   TestPgm/64x64x1-2
=== RUN   TestPgm/64x64x1-3
=== RUN   TestPgm/64x64x1-4
=== RUN   TestPgm/64x64x1-5
=== RUN   TestPgm/64x64x1-6
=== RUN   TestPgm/64x64x1-7
=== RUN   TestPgm/64x64x1-8
=== RUN   TestPgm/64x64x1-9
=== RUN   TestPgm/64x64x1-10
=== RUN   TestPgm/64x64x1-11
=== RUN   TestPgm/64x64x1-12
=== RUN   TestPgm/64x64x1-13
=== RUN   TestPgm/64x64x1-14
=== RUN   TestPgm/64x64x1-15
=== RUN   TestPgm/64x64x1-16
=== RUN   TestPgm/64x64x100-1
=== RUN   TestPgm/64x64x100-2
=== RUN   TestPgm/64x64x100-3
=== RUN   TestPgm/64x64x100-4
=== RUN   TestPgm/64x64x100-5
=== RUN   TestPgm/64x64x100-6
=== RUN   TestPgm/64x64x100-7
=== RUN   TestPgm/64x64x100-8
=== RUN   TestPgm/64x64x100-9
=== RUN   TestPgm/64x64x100-10
=== RUN   TestPgm/64x64x100-11
=== RUN   TestPgm/64x64x100-12
=== RUN   TestPgm/64x64x100-13
=== RUN   TestPgm/64x64x100-14
=== RUN   TestPgm/64x64x100-15
=== RUN   TestPgm/64x64x100-16
=== RUN   TestPgm/512x512x0-1
=== RUN   TestPgm/512x512x0-2
=== RUN   TestPgm/512x512x0-3
=== RUN   TestPgm/512x512x0-4
=== RUN   TestPgm/512x512x0-5
=== RUN   TestPgm/512x512x0-6
=== RUN   TestPgm/512x512x0-7
=== RUN   TestPgm/512x512x0-8
=== RUN   TestPgm/512x512x0-9
=== RUN   TestPgm/512x512x0-10
=== RUN   TestPgm/512x512x0-11
=== RUN   TestPgm/512x512x0-12
=== RUN   TestPgm/512x512x0-13
=== RUN   TestPgm/512x512x0-14
=== RUN   TestPgm/512x512x0-15
=== RUN   TestPgm/512x512x0-16
=== RUN   TestPgm/512x512x1-1
=== RUN   TestPgm/512x512x1-2
=== RUN   TestPgm/512x512x1-3
=== RUN   TestPgm/512x512x1-4
=== RUN   TestPgm/512x512x1-5
=== RUN   TestPgm/512x512x1-6
=== RUN   TestPgm/512x512x1-7
=== RUN   TestPgm/512x512x1-8
=== RUN   TestPgm/512x512x1-9
=== RUN   TestPgm/512x512x1-10
=== RUN   TestPgm/512x512x1-11
=== RUN   TestPgm/512x512x1-12
=== RUN   TestPgm/512x512x1-13
=== RUN   TestPgm/512x512x1-14
=== RUN   TestPgm/512x512x1-15
=== RUN   TestPgm/512x512x1-16
=== RUN   TestPgm/512x512x100-1
=== RUN   TestPgm/512x512x100-2
=== RUN   TestPgm/512x512x100-3
=== RUN   TestPgm/512x512x100-4
=== RUN   TestPgm/512x512x100-5
=== RUN   TestPgm/512x512x100-6
=== RUN   TestPgm/512x512x100-7
=== RUN   TestPgm/512x512x100-8
=== RUN   TestPgm/512x512x100-9
=== RUN   TestPgm/512x512x100-10
=== RUN   TestPgm/512x512x100-11
=== RUN   TestPgm/512x512x100-12
=== RUN   TestPgm/512x512x100-13
=== RUN   TestPgm/512x512x100-14
=== RUN   TestPgm/512x512x100-15
=== RUN   TestPgm/512x512x100-16
--- PASS: TestPgm (74.36s)
    --- PASS: TestPgm/16x16x0-1 (0.01s)
    --- PASS: TestPgm/16x16x0-2 (0.01s)
    --- PASS: TestPgm/16x16x0-3 (0.01s)
    --- PASS: TestPgm/16x16x0-4 (0.00s)
    --- PASS: TestPgm/16x16x0-5 (0.00s)
    --- PASS: TestPgm/16x16x0-6 (0.01s)
    --- PASS: TestPgm/16x16x0-7 (0.00s)
    --- PASS: TestPgm/16x16x0-8 (0.01s)
    --- PASS: TestPgm/16x16x0-9 (0.00s)
    --- PASS: TestPgm/16x16x0-10 (0.01s)
    --- PASS: TestPgm/16x16x0-11 (0.01s)
    --- PASS: TestPgm/16x16x0-12 (0.00s)
    --- PASS: TestPgm/16x16x0-13 (0.01s)
    --- PASS: TestPgm/16x16x0-14 (0.00s)
    --- PASS: TestPgm/16x16x0-15 (0.00s)
    --- PASS: TestPgm/16x16x0-16 (0.01s)
    --- PASS: TestPgm/16x16x1-1 (0.00s)
    --- PASS: TestPgm/16x16x1-2 (0.00s)
    --- PASS: TestPgm/16x16x1-3 (0.00s)
    --- PASS: TestPgm/16x16x1-4 (0.00s)
    --- PASS: TestPgm/16x16x1-5 (0.00s)
    --- PASS: TestPgm/16x16x1-6 (0.00s)
    --- PASS: TestPgm/16x16x1-7 (0.00s)
    --- PASS: TestPgm/16x16x1-8 (0.00s)
    --- PASS: TestPgm/16x16x1-9 (0.01s)
    --- PASS: TestPgm/16x16x1-10 (0.00s)
    --- PASS: TestPgm/16x16x1-11 (0.01s)
    --- PASS: TestPgm/16x16x1-12 (0.01s)
    --- PASS: TestPgm/16x16x1-13 (0.01s)
    --- PASS: TestPgm/16x16x1-14 (0.00s)
    --- PASS: TestPgm/16x16x1-15 (0.01s)
    --- PASS: TestPgm/16x16x1-16 (0.00s)
    --- PASS: TestPgm/16x16x100-1 (0.01s)
    --- PASS: TestPgm/16x16x100-2 (0.01s)
    --- PASS: TestPgm/16x16x100-3 (0.01s)
    --- PASS: TestPgm/16x16x100-4 (0.01s)
    --- PASS: TestPgm/16x16x100-5 (0.01s)
    --- PASS: TestPgm/16x16x100-6 (0.01s)
    --- PASS: TestPgm/16x16x100-7 (0.01s)
    --- PASS: TestPgm/16x16x100-8 (0.01s)
    --- PASS: TestPgm/16x16x100-9 (0.01s)
    --- PASS: TestPgm/16x16x100-10 (0.01s)
    --- PASS: TestPgm/16x16x100-11 (0.01s)
    --- PASS: TestPgm/16x16x100-12 (0.01s)
    --- PASS: TestPgm/16x16x100-13 (0.01s)
    --- PASS: TestPgm/16x16x100-14 (0.01s)
    --- PASS: TestPgm/16x16x100-15 (0.01s)
    --- PASS: TestPgm/16x16x100-16 (0.01s)
    --- PASS: TestPgm/64x64x0-1 (0.02s)
    --- PASS: TestPgm/64x64x0-2 (0.02s)
    --- PASS: TestPgm/64x64x0-3 (0.03s)
    --- PASS: TestPgm/64x64x0-4 (0.03s)
    --- PASS: TestPgm/64x64x0-5 (0.02s)
    --- PASS: TestPgm/64x64x0-6 (0.02s)
    --- PASS: TestPgm/64x64x0-7 (0.02s)
    --- PASS: TestPgm/64x64x0-8 (0.02s)
    --- PASS: TestPgm/64x64x0-9 (0.02s)
    --- PASS: TestPgm/64x64x0-10 (0.02s)
    --- PASS: TestPgm/64x64x0-11 (0.03s)
    --- PASS: TestPgm/64x64x0-12 (0.02s)
    --- PASS: TestPgm/64x64x0-13 (0.02s)
    --- PASS: TestPgm/64x64x0-14 (0.03s)
    --- PASS: TestPgm/64x64x0-15 (0.03s)
    --- PASS: TestPgm/64x64x0-16 (0.02s)
    --- PASS: TestPgm/64x64x1-1 (0.02s)
    --- PASS: TestPgm/64x64x1-2 (0.02s)
    --- PASS: TestPgm/64x64x1-3 (0.02s)
    --- PASS: TestPgm/64x64x1-4 (0.02s)
    --- PASS: TestPgm/64x64x1-5 (0.02s)
    --- PASS: TestPgm/64x64x1-6 (0.02s)
    --- PASS: TestPgm/64x64x1-7 (0.02s)
    --- PASS: TestPgm/64x64x1-8 (0.02s)
    --- PASS: TestPgm/64x64x1-9 (0.02s)
    --- PASS: TestPgm/64x64x1-10 (0.02s)
    --- PASS: TestPgm/64x64x1-11 (0.02s)
    --- PASS: TestPgm/64x64x1-12 (0.02s)
    --- PASS: TestPgm/64x64x1-13 (0.02s)
    --- PASS: TestPgm/64x64x1-14 (0.02s)
    --- PASS: TestPgm/64x64x1-15 (0.02s)
    --- PASS: TestPgm/64x64x1-16 (0.02s)
    --- PASS: TestPgm/64x64x100-1 (0.04s)
    --- PASS: TestPgm/64x64x100-2 (0.04s)
    --- PASS: TestPgm/64x64x100-3 (0.03s)
    --- PASS: TestPgm/64x64x100-4 (0.04s)
    --- PASS: TestPgm/64x64x100-5 (0.04s)
    --- PASS: TestPgm/64x64x100-6 (0.04s)
    --- PASS: TestPgm/64x64x100-7 (0.04s)
    --- PASS: TestPgm/64x64x100-8 (0.04s)
    --- PASS: TestPgm/64x64x100-9 (0.04s)
    --- PASS: TestPgm/64x64x100-10 (0.03s)
    --- PASS: TestPgm/64x64x100-11 (0.03s)
    --- PASS: TestPgm/64x64x100-12 (0.04s)
    --- PASS: TestPgm/64x64x100-13 (0.03s)
    --- PASS: TestPgm/64x64x100-14 (0.04s)
    --- PASS: TestPgm/64x64x100-15 (0.03s)
    --- PASS: TestPgm/64x64x100-16 (15.52s)
    --- PASS: TestPgm/512x512x0-1 (0.87s)
    --- PASS: TestPgm/512x512x0-2 (0.89s)
    --- PASS: TestPgm/512x512x0-3 (1.00s)
    --- PASS: TestPgm/512x512x0-4 (0.95s)
    --- PASS: TestPgm/512x512x0-5 (0.97s)
    --- PASS: TestPgm/512x512x0-6 (0.99s)
    --- PASS: TestPgm/512x512x0-7 (1.00s)
    --- PASS: TestPgm/512x512x0-8 (1.01s)
    --- PASS: TestPgm/512x512x0-9 (0.97s)
    --- PASS: TestPgm/512x512x0-10 (0.91s)
    --- PASS: TestPgm/512x512x0-11 (0.99s)
    --- PASS: TestPgm/512x512x0-12 (0.92s)
    --- PASS: TestPgm/512x512x0-13 (0.91s)
    --- PASS: TestPgm/512x512x0-14 (0.95s)
    --- PASS: TestPgm/512x512x0-15 (0.98s)
    --- PASS: TestPgm/512x512x0-16 (1.01s)
    --- PASS: TestPgm/512x512x1-1 (0.94s)
    --- PASS: TestPgm/512x512x1-2 (1.00s)
    --- PASS: TestPgm/512x512x1-3 (0.97s)
    --- PASS: TestPgm/512x512x1-4 (0.98s)
    --- PASS: TestPgm/512x512x1-5 (1.03s)
    --- PASS: TestPgm/512x512x1-6 (1.00s)
    --- PASS: TestPgm/512x512x1-7 (0.99s)
    --- PASS: TestPgm/512x512x1-8 (1.00s)
    --- PASS: TestPgm/512x512x1-9 (0.97s)
    --- PASS: TestPgm/512x512x1-10 (1.01s)
    --- PASS: TestPgm/512x512x1-11 (1.02s)
    --- PASS: TestPgm/512x512x1-12 (1.00s)
    --- PASS: TestPgm/512x512x1-13 (0.99s)
    --- PASS: TestPgm/512x512x1-14 (0.99s)
    --- PASS: TestPgm/512x512x1-15 (1.04s)
    --- PASS: TestPgm/512x512x1-16 (0.96s)
    --- PASS: TestPgm/512x512x100-1 (2.82s)
    --- PASS: TestPgm/512x512x100-2 (2.07s)
    --- PASS: TestPgm/512x512x100-3 (1.87s)
    --- PASS: TestPgm/512x512x100-4 (1.70s)
    --- PASS: TestPgm/512x512x100-5 (1.60s)
    --- PASS: TestPgm/512x512x100-6 (1.55s)
    --- PASS: TestPgm/512x512x100-7 (1.49s)
    --- PASS: TestPgm/512x512x100-8 (1.52s)
    --- PASS: TestPgm/512x512x100-9 (1.43s)
    --- PASS: TestPgm/512x512x100-10 (1.47s)
    --- PASS: TestPgm/512x512x100-11 (1.50s)
    --- PASS: TestPgm/512x512x100-12 (1.42s)
    --- PASS: TestPgm/512x512x100-13 (1.38s)
    --- PASS: TestPgm/512x512x100-14 (1.42s)
    --- PASS: TestPgm/512x512x100-15 (1.40s)
    --- PASS: TestPgm/512x512x100-16 (1.40s)
=== RUN   TestSdl
=== RUN   TestSdl/turn
    util_test.go:483: Testing for FinalTurnComplete event
    util_test.go:278: [Event] Completed Turns 0        Executing
    util_test.go:275: [Event] Completed Turns 100      Final Turn Complete
    util_test.go:504: Testing number of TurnComplete events sent
=== RUN   TestSdl/images
    util_test.go:409: Testing for first StateChange Executing event
    util_test.go:278: [Event] Completed Turns 0        Executing
    util_test.go:367: Checking SDL image at turn 0
    util_test.go:367: Checking SDL image at turn 1
    util_test.go:367: Checking SDL image at turn 100
=== RUN   TestSdl/alive
    util_test.go:409: Testing for first StateChange Executing event
    util_test.go:278: [Event] Completed Turns 0        Executing
    util_test.go:336: Checking number of alive cells in the SDL window at turn 1
    util_test.go:336: Checking number of alive cells in the SDL window at turn 2
    util_test.go:336: Checking number of alive cells in the SDL window at turn 3
    util_test.go:336: Checking number of alive cells in the SDL window at turn 4
    util_test.go:336: Checking number of alive cells in the SDL window at turn 5
    util_test.go:336: Checking number of alive cells in the SDL window at turn 6
    util_test.go:336: Checking number of alive cells in the SDL window at turn 7
    util_test.go:336: Checking number of alive cells in the SDL window at turn 8
    util_test.go:336: Checking number of alive cells in the SDL window at turn 9
    util_test.go:336: Checking number of alive cells in the SDL window at turn 10
    util_test.go:336: Checking number of alive cells in the SDL window at turn 11
    util_test.go:336: Checking number of alive cells in the SDL window at turn 12
    util_test.go:336: Checking number of alive cells in the SDL window at turn 13
    util_test.go:336: Checking number of alive cells in the SDL window at turn 14
    util_test.go:336: Checking number of alive cells in the SDL window at turn 15
    util_test.go:336: Checking number of alive cells in the SDL window at turn 16
    util_test.go:336: Checking number of alive cells in the SDL window at turn 17
    util_test.go:336: Checking number of alive cells in the SDL window at turn 18
    util_test.go:336: Checking number of alive cells in the SDL window at turn 19
    util_test.go:336: Checking number of alive cells in the SDL window at turn 20
    util_test.go:336: Checking number of alive cells in the SDL window at turn 21
    util_test.go:336: Checking number of alive cells in the SDL window at turn 22
    util_test.go:336: Checking number of alive cells in the SDL window at turn 23
    util_test.go:336: Checking number of alive cells in the SDL window at turn 24
    util_test.go:336: Checking number of alive cells in the SDL window at turn 25
    util_test.go:336: Checking number of alive cells in the SDL window at turn 26
    util_test.go:336: Checking number of alive cells in the SDL window at turn 27
    util_test.go:336: Checking number of alive cells in the SDL window at turn 28
    util_test.go:336: Checking number of alive cells in the SDL window at turn 29
    util_test.go:336: Checking number of alive cells in the SDL window at turn 30
    util_test.go:336: Checking number of alive cells in the SDL window at turn 31
    util_test.go:336: Checking number of alive cells in the SDL window at turn 32
    util_test.go:336: Checking number of alive cells in the SDL window at turn 33
    util_test.go:336: Checking number of alive cells in the SDL window at turn 34
    util_test.go:336: Checking number of alive cells in the SDL window at turn 35
    util_test.go:336: Checking number of alive cells in the SDL window at turn 36
    util_test.go:336: Checking number of alive cells in the SDL window at turn 37
    util_test.go:336: Checking number of alive cells in the SDL window at turn 38
    util_test.go:336: Checking number of alive cells in the SDL window at turn 39
    util_test.go:336: Checking number of alive cells in the SDL window at turn 40
    util_test.go:336: Checking number of alive cells in the SDL window at turn 41
    util_test.go:336: Checking number of alive cells in the SDL window at turn 42
    util_test.go:336: Checking number of alive cells in the SDL window at turn 43
    util_test.go:336: Checking number of alive cells in the SDL window at turn 44
    util_test.go:336: Checking number of alive cells in the SDL window at turn 45
    util_test.go:336: Checking number of alive cells in the SDL window at turn 46
    util_test.go:336: Checking number of alive cells in the SDL window at turn 47
    util_test.go:336: Checking number of alive cells in the SDL window at turn 48
    util_test.go:336: Checking number of alive cells in the SDL window at turn 49
    util_test.go:336: Checking number of alive cells in the SDL window at turn 50
    util_test.go:336: Checking number of alive cells in the SDL window at turn 51
    util_test.go:336: Checking number of alive cells in the SDL window at turn 52
    util_test.go:336: Checking number of alive cells in the SDL window at turn 53
    util_test.go:336: Checking number of alive cells in the SDL window at turn 54
    util_test.go:336: Checking number of alive cells in the SDL window at turn 55
    util_test.go:336: Checking number of alive cells in the SDL window at turn 56
    util_test.go:336: Checking number of alive cells in the SDL window at turn 57
    util_test.go:336: Checking number of alive cells in the SDL window at turn 58
    util_test.go:336: Checking number of alive cells in the SDL window at turn 59
    util_test.go:336: Checking number of alive cells in the SDL window at turn 60
    util_test.go:336: Checking number of alive cells in the SDL window at turn 61
    util_test.go:336: Checking number of alive cells in the SDL window at turn 62
    util_test.go:336: Checking number of alive cells in the SDL window at turn 63
    util_test.go:336: Checking number of alive cells in the SDL window at turn 64
    util_test.go:336: Checking number of alive cells in the SDL window at turn 65
    util_test.go:336: Checking number of alive cells in the SDL window at turn 66
    util_test.go:336: Checking number of alive cells in the SDL window at turn 67
    util_test.go:336: Checking number of alive cells in the SDL window at turn 68
    util_test.go:336: Checking number of alive cells in the SDL window at turn 69
    util_test.go:336: Checking number of alive cells in the SDL window at turn 70
    util_test.go:336: Checking number of alive cells in the SDL window at turn 71
    util_test.go:336: Checking number of alive cells in the SDL window at turn 72
    util_test.go:336: Checking number of alive cells in the SDL window at turn 73
    util_test.go:336: Checking number of alive cells in the SDL window at turn 74
    util_test.go:336: Checking number of alive cells in the SDL window at turn 75
    util_test.go:336: Checking number of alive cells in the SDL window at turn 76
    util_test.go:336: Checking number of alive cells in the SDL window at turn 77
    util_test.go:336: Checking number of alive cells in the SDL window at turn 78
    util_test.go:336: Checking number of alive cells in the SDL window at turn 79
    util_test.go:336: Checking number of alive cells in the SDL window at turn 80
    util_test.go:336: Checking number of alive cells in the SDL window at turn 81
    util_test.go:336: Checking number of alive cells in the SDL window at turn 82
    util_test.go:336: Checking number of alive cells in the SDL window at turn 83
    util_test.go:336: Checking number of alive cells in the SDL window at turn 84
    util_test.go:336: Checking number of alive cells in the SDL window at turn 85
    util_test.go:336: Checking number of alive cells in the SDL window at turn 86
    util_test.go:336: Checking number of alive cells in the SDL window at turn 87
    util_test.go:336: Checking number of alive cells in the SDL window at turn 88
    util_test.go:336: Checking number of alive cells in the SDL window at turn 89
    util_test.go:336: Checking number of alive cells in the SDL window at turn 90
    util_test.go:336: Checking number of alive cells in the SDL window at turn 91
    util_test.go:336: Checking number of alive cells in the SDL window at turn 92
    util_test.go:336: Checking number of alive cells in the SDL window at turn 93
    util_test.go:336: Checking number of alive cells in the SDL window at turn 94
    util_test.go:336: Checking number of alive cells in the SDL window at turn 95
    util_test.go:336: Checking number of alive cells in the SDL window at turn 96
    util_test.go:336: Checking number of alive cells in the SDL window at turn 97
    util_test.go:336: Checking number of alive cells in the SDL window at turn 98
    util_test.go:336: Checking number of alive cells in the SDL window at turn 99
    util_test.go:336: Checking number of alive cells in the SDL window at turn 100
--- PASS: TestSdl (4.34s)
    --- PASS: TestSdl/turn (1.44s)
    --- PASS: TestSdl/images (1.44s)
    --- PASS: TestSdl/alive (1.46s)
=== RUN   TestTrace
--- PASS: TestTrace (0.04s)
FAIL
exit status 1
FAIL    uk.ac.bris.cs/gameoflife/tests  169.230s
