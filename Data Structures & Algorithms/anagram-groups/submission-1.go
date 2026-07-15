func groupAnagrams(strs []string) [][]string {
    /* 1.
     * Declare a map
     * key: [1, 0, 1, 0, ..., 0]
     * value: ['act', 'cat']
     */    
    res := map[[26]int][]string{}

    // 2. For each word in strs...
    for _, s:= range strs {
        // 2.1 Declare the key for our map, initialized with 0 values
        var count [26]int

        /* 2.2
         * For each char in the word, index into its position
         * in the count array, and increment the value.
         * This determines the frequency of each character
         * appearing in the word.
         */
        for _, c := range s {
            count[c-'a']++
        }

        // 2.3 Append the string to the value array
        res[count] = append(res[count], s)
    }

    // 3. Declare an output variable
    var output [][]string

    // 4. For each key-value pair in our map...
    for _, group := range res {
        // 4.1 Simply append the value to the output array
        output = append(output, group)
    }

    // 5. Return the output
    return output
}
