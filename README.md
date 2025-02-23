# Aidantix

Un coup de pouce pour le Cémantix.

## Installation and usage

1. Get [Go](https://go.dev). `[packagemanager] install go` should work. Othervise, visit https://go.dev

2. Build (`go build main.go`) and run (`./main`).

3. Enter the requested time buff between requests (to avoid DoS ban), for instance `1000` milliseconds.

4. Answer the prompt asking if you want verbose or non-verbose mode. With verbose mode turned on, all requests and requests answers will be printed to standard output. Without, the command prompt will shut up and work silently.

5. Let him cook.

6. The results are dumped in in `result.csv`.

```csv
Word,Rank,Score
chaleur,328,1.4500
feu,928,0.7280
flamme,454,0.6550
chauffer,228,0.3010
température,112,0.132
chaud,550,0.4410
incendie,808,0.9410
Brûlure,799,0.9330
```

## Dependencies

- github.com/andybalholm/cascadia
- golang.org/x/net