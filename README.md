# Aidantix

Un coup de pouce pour le Cémantix.

## Installation and usage

1. Build (`go build main.go`) and run (`./main`)

2. Enter the requested time buff between requests (to avoid DoS ban), for instance `1000` milliseconds.

3. Let him cook.

4. The results are dumped in in `result.csv`.

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