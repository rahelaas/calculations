## math-skills

The purpose of this program is to calculate the following:

1. Average
2. Median
3. Variance
4. Standard Deviation

The program has been built in Golang.

For the audit:
1. Download the [math-skills test program](https://assets.01-edu.org/stats-projects/math-skills) and  place it in the math-skills folder.

2. To allow executable permissions:
```bash
chmod +x math-skills
```
To generate random set of numbers in data.txt and save the answers in file.txt:
```bash
./math-skills
```

On a Mac use Docker and use the following commands:
```bash
RUN ./math-skills > file.txt
```
```bash
docker cp math:./app/file.txt ./
```

3. Run the program and compare the results printed from the main.go file and in the file.txt. The answers should match.
```bash
go run main.go data.txt
```

Here you can find the [audit questions](https://github.com/01-edu/public/tree/master/subjects/math-skills/audit).