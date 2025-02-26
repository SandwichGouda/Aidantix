- Meticulously follow [Module version docs](https://go.dev/doc/tutorial/create-module) to manage module versions
- Change versions whenever module is updated
- Add a package for sorting the `result.csv` file
  * Create a **new** file `sorted_result.csv`, to allow linear writing too
- Add other word databases
- Automatically remove unknown words from databases
- Use goroutines to cann the scrapers and retrieve their given values, such that if they encounter an error and exit (`log.Fatal(...)`), the error gets printed and another method is used (use default scraper, then goquery scraper, then ask user to hardcode the day number)
  * Also, GoQuery scraper should not use `Fatal()` (which is equivalent to `fmt.Print()` followed by an `exit()`) if only one day is retrieved but not the other
  * Return `Errors` instead
- Comment the code more, explaining function roles, block chunks and variable roles
- Make two separate outputs `result.csv` and `result_long.csv`, and read from `result_long.csv` to get where to start from
- Add Pedantix version. Add english version.