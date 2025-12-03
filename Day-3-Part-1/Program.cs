
// Batteries that supply emergency power to an escelator
// Each battery is labaled with joltage rating
// Joltage - a value from 1 to 9
// Puzzle input:
// 987654321111111
// 811111111111119
// 234234234234278
// 818181911112111

// Batteries are arranged into banks
// Eachc line of digits is 1 bank
// Turn on EXACTLY 2 batteries
// The joltage should equal the number formed by the digits on the batteries you've turned on.
// Example:
// 12345 - turn on 2 and 4
// the bank would produce 24 jolts (You cannot rearrange batteries
// Find the largest combination
// 987654321111111 = 98
// 811111111111119 = 89
// 234234234234278 = 78
// 818181911112111 = 92

// Answer - total output joltage (98 + 89 + 78 + 92 = 357)


// See https://aka.ms/new-console-template for more information

var data = File.ReadAllLines("input.txt");
int totalJoltage = 0;
foreach (var line in data) {
    if (string.IsNullOrEmpty(line)) {
        continue;
    }
    Console.WriteLine(line);
    int highestJoltage = HighestJoltage(line.ToCharArray());
    Console.WriteLine(highestJoltage);
    totalJoltage += highestJoltage;
}
Console.WriteLine($"Total Joltage: {totalJoltage}");

static int HighestJoltage(char[] data) {
    int highestJoltage = 0;
    for (int i = 0; i < data.Length -1; i++) {
        for (int j = i + 1; j < data.Length; j++) {
            var firstChar = data[i];
            var secondChar = data[j];
            int joltage = int.Parse(new string(new[] { firstChar, secondChar }));
            if (joltage > highestJoltage) {
                highestJoltage = joltage;
            }
        }
    }
    return highestJoltage;
}