
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


// Part 2
// Need more voltage4
// Turn on exactly twelve batteries



// See https://aka.ms/new-console-template for more information

// var data = File.ReadAllLines("input.txt");
// long totalJoltage = 0;
// foreach (var line in data) {
//     if (string.IsNullOrEmpty(line)) {
//         continue;
//     }
//     Console.WriteLine($"Line: {line}");
//     long highestJoltage = HighestJoltage(line.ToCharArray(), 0, 12, "");
//     Console.WriteLine($"Highest Joltage: {highestJoltage}");
//     totalJoltage += highestJoltage;
// }
// Console.WriteLine($"Total Joltage: {totalJoltage}");

// static long HighestJoltage(
//     char[] data, 
//     int position, 
//     int remaining, 
//     string current
// ) {
//     if (remaining == 0) {
//         long joltage = long.Parse(current);
//         return joltage;
//     }

//     if (position >= data.Length || (data.Length - position) < remaining) {
//         return -1;
//     }

//     long result1 = HighestJoltage(data, position + 1, remaining - 1, current + data[position].ToString());
//     long result2 = HighestJoltage(data, position + 1, remaining, current);
//     return result1 > result2 ? result1 : result2;
// }

var data = File.ReadAllLines("input.txt");
long totalJoltage = 0;
foreach (var line in data) {
    if (string.IsNullOrEmpty(line)) {
        continue;
    }
    Console.WriteLine($"Line: {line}");
    long highestJoltage = FindHighestNumber(line.ToCharArray(), 12);
    Console.WriteLine($"Highest Joltage: {highestJoltage}");
    totalJoltage += highestJoltage;
}
Console.WriteLine($"Total Joltage: {totalJoltage}");

static long FindHighestNumber(char[] data, int count) {
    char[] result = new char[count];
    int startPosition = 0;

    for (int i = 0; i < count; i++) {
        int remaining = count - i;
        int endPost = data.Length - remaining; // Last valid position to pick from

        // Find the maximum digit in the valid range
        int maxPost = startPosition;
        for (int j = startPosition +1 ; j <= endPost; j++) {
            if (data[j] > data[maxPost]) {
                maxPost = j;
            }
        }

        // Place the maximum digit in the current position
        result[i] = data[maxPost];
        startPosition = maxPost + 1;
    }

    return long.Parse(new string(result));
    }