package com.hackerranksolution;

import static org.junit.Assert.assertArrayEquals;
import static org.junit.Assert.assertEquals;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.IntStream;
import java.util.stream.Stream;
import static java.util.stream.Collectors.toList;

import org.junit.Test;

public class FrequencyQueriesTest {

    @Test
    public void testFirstSample() {
        Integer[] operations = new Integer[] { 1, 1, 2, 2, 3, 2, 1, 1, 1, 1, 2, 1, 3, 2 };
        List<List<Integer>> input = new ArrayList<>();
        for (int i = 0; i < operations.length; i += 2) {
            input.add(Arrays.asList(operations[i], operations[i + 1]));
        }
        List<Integer> freqQuery = FrequencyQueries.freqQuery(input);
        assertEquals(freqQuery.get(0), Integer.valueOf(0));
        assertEquals(freqQuery.get(1), Integer.valueOf(1));
    }

    @Test
    public void testLargeInput() throws NumberFormatException, IOException {
        List<List<Integer>> queries = getInputListFromFile("src/test/java/assets/FrequentQueriesInput8.txt");

        List<Integer> freqQuery = FrequencyQueries.freqQuery(queries);
        Integer[] array = new Integer[freqQuery.size()];
        freqQuery.toArray(array); // fill the array

        List<Integer> outputList = getOutputListFromFile("src/test/java/assets/AnswerInput8.txt");
        Integer[] output = new Integer[outputList.size()];
        outputList.toArray(output); // fill the array

        assertArrayEquals(array, output);
    }

    private List<List<Integer>> getInputListFromFile(String filePath) throws NumberFormatException, IOException{
        FileReader fr = new FileReader(filePath);
        BufferedReader br = new BufferedReader(fr);

        int q = Integer.parseInt(br.readLine().trim());

        List<List<Integer>> queries = new ArrayList<>();

        IntStream.range(0, q).forEach(i -> {
            try {
                queries.add(
                        Stream.of(br.readLine().replaceAll("\\s+$", "").split(" "))
                                .map(Integer::parseInt)
                                .collect(toList()));
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        });

        br.close();
        return queries;
    }

    private List<Integer> getOutputListFromFile(String filePath) throws NumberFormatException, IOException{
        FileReader fr = new FileReader(filePath);
        BufferedReader br = new BufferedReader(fr);

        int q = Integer.parseInt(br.readLine().trim());

        List<Integer> results = new ArrayList<>();

        IntStream.range(0, q).forEach(i -> {
            try {
                results.add(Integer.parseInt(br.readLine()));
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        });

        br.close();
        return results;
    }
}
