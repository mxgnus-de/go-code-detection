package language_test

import (
	"testing"

	"github.com/mxgnus-de/go-code-detection/detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestJavaHelloWorldLanguageDetection(t *testing.T) {
	t.Log("Testing Java Hello World Language Detection")
	code := `System.out.println("Hello world!");`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestJavaFizzBuzzLanguageDetection(t *testing.T) {
	t.Log("Testing Java FizzBuzz Language Detection")
	code := `public class FizzBuzz {
      public static void main(String[] args) {
         for(int i = 1; i <= 100; i++) {
         if (((i % 5) == 0) && ((i % 7) == 0))
            System.out.print("fizzbuzz");    
         else if ((i % 5) == 0) System.out.print("fizz");
         else if ((i % 7) == 0) System.out.print("buzz");
         else System.out.print(i);
         System.out.print(" "); 
         }
         System.out.println();
      }
   }`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestJavaGetterAndSetterLanguageDetection(t *testing.T) {
	t.Log("Testing Java Getter and Setter Language Detection")
	code := `
      Person person = people.get(0);
      person.setName("John");
      person.setAge(30);
   `
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestJavaListAndArrayListLanguageDetection(t *testing.T) {
	t.Log("Testing Java List and ArrayList Language Detection")
	code := `List<String> things = new ArrayList<>();`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestJavaQuickSortLanguageDetection(t *testing.T) {
	t.Log("Testing Java Quick Sort Language Detection")
	code := `public static <E extends Comparable<? super E>> List<E> quickSort(List<E> arr) {
      if (arr.isEmpty())
         return arr;
      else {
         E pivot = arr.get(0);
   
         List<E> less = new LinkedList<E>();
         List<E> pivotList = new LinkedList<E>();
         List<E> more = new LinkedList<E>();
   
         // Partition
         for (E i: arr) {
               if (i.compareTo(pivot) < 0)
                  less.add(i);
               else if (i.compareTo(pivot) > 0)
                  more.add(i);
               else
                  pivotList.add(i);
         }
   
         // Recursively sort sublists
         less = quickSort(less);
         more = quickSort(more);
   
         // Concatenate results
         less.addAll(pivotList);
         less.addAll(more);
         return less;
      }
   }
   `
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestJavaBubbleSortLanguageDetection(t *testing.T) {
	t.Log("Testing Java Bubble Sort Language Detection")
	code := `public static <E extends Comparable<? super E>> void bubbleSort(E[] comparable) {
      boolean changed = false;
      do {
         changed = false;
         for (int a = 0; a < comparable.length - 1; a++) {
               if (comparable[a].compareTo(comparable[a + 1]) > 0) {
                  E tmp = comparable[a];
                  comparable[a] = comparable[a + 1];
                  comparable[a + 1] = tmp;
                  changed = true;
               }
         }
      } while (changed);
   }`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestJavaHttpServerLanguageDetection(t *testing.T) {
	t.Log("Testing Java Http Server Language Detection")
	code := `import java.net.URI;
   import java.net.http.HttpClient;
   import java.net.http.HttpRequest;
   import java.net.http.HttpResponse;
   import java.nio.charset.Charset;
      
   public class Main {
         public static void main(String[] args) {
            var request = HttpRequest.newBuilder(URI.create("https://www.rosettacode.org"))
                     .GET()
                     .build();
      
            HttpClient.newHttpClient()
                     .sendAsync(request, HttpResponse.BodyHandlers.ofString(Charset.defaultCharset()))
                     .thenApply(HttpResponse::body)
                     .thenAccept(System.out::println)
                     .join();
         }
   }`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestJavaFloydWarshallAlgorithmLanguageDetection(t *testing.T) {
	t.Log("Testing Java Floyd Warshall Algorithm Language Detection")
	code := `import static java.lang.String.format;
   import java.util.Arrays;
      
   public class FloydWarshall {
      
         public static void main(String[] args) {
            int[][] weights = {{1, 3, -2}, {2, 1, 4}, {2, 3, 3}, {3, 4, 2}, {4, 2, -1}};
            int numVertices = 4;
      
            floydWarshall(weights, numVertices);
         }
      
         static void floydWarshall(int[][] weights, int numVertices) {
      
            double[][] dist = new double[numVertices][numVertices];
            for (double[] row : dist)
               Arrays.fill(row, Double.POSITIVE_INFINITY);
      
            for (int[] w : weights)
               dist[w[0] - 1][w[1] - 1] = w[2];
      
            int[][] next = new int[numVertices][numVertices];
            for (int i = 0; i < next.length; i++) {
               for (int j = 0; j < next.length; j++)
                     if (i != j)
                        next[i][j] = j + 1;
            }
      
            for (int k = 0; k < numVertices; k++)
               for (int i = 0; i < numVertices; i++)
                     for (int j = 0; j < numVertices; j++)
                        if (dist[i][k] + dist[k][j] < dist[i][j]) {
                           dist[i][j] = dist[i][k] + dist[k][j];
                           next[i][j] = next[i][k];
                        }
      
            printResult(dist, next);
         }
      
         static void printResult(double[][] dist, int[][] next) {
            System.out.println("pair     dist    path");
            for (int i = 0; i < next.length; i++) {
               for (int j = 0; j < next.length; j++) {
                     if (i != j) {
                        int u = i + 1;
                        int v = j + 1;
                        String path = format("%d -> %d    %2d     %s", u, v,
                                 (int) dist[i][j], u);
                        do {
                           u = next[u - 1][v - 1];
                           path += " -> " + u;
                        } while (u != v);
                        System.out.println(path);
                     }
               }
            }
         }
   }`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestJavaLudicNumbersLanguageDetection(t *testing.T) {
	t.Log("Testing Java Ludic Numbers Language Detection")
	code := `import java.util.ArrayList;
   import java.util.List;
      
   public class Ludic{
      public static List<Integer> ludicUpTo(int n){
         List<Integer> ludics = new ArrayList<Integer>(n);
         for(int i = 1; i <= n; i++){   //fill the initial list
         ludics.add(i);
         }
      
         //start at index 1 because the first ludic number is 1 and we don't remove anything for it
         for(int cursor = 1; cursor < ludics.size(); cursor++){
         int thisLudic = ludics.get(cursor); //the first item in the list is a ludic number
         int removeCursor = cursor + thisLudic; //start removing that many items later
         while(removeCursor < ludics.size()){
            ludics.remove(removeCursor);		     //remove the next item
            removeCursor = removeCursor + thisLudic - 1; //move the removal cursor up as many spaces as we need to
                           //then back one to make up for the item we just removed
         }
         }
         return ludics;
      }
      
      public static List<List<Integer>> getTriplets(List<Integer> ludics){
         List<List<Integer>> triplets = new ArrayList<List<Integer>>();
         for(int i = 0; i < ludics.size() - 2; i++){ //only need to check up to the third to last item
         int thisLudic = ludics.get(i);
         if(ludics.contains(thisLudic + 2) && ludics.contains(thisLudic + 6)){
            List<Integer> triplet = new ArrayList<Integer>(3);
            triplet.add(thisLudic);
            triplet.add(thisLudic + 2);
            triplet.add(thisLudic + 6);
            triplets.add(triplet);
         }
         }
         return triplets;
      }
      
      public static void main(String[] srgs){
         System.out.println("First 25 Ludics: " + ludicUpTo(110));				//110 will get us 25 numbers
         System.out.println("Ludics up to 1000: " + ludicUpTo(1000).size());
         System.out.println("2000th - 2005th Ludics: " + ludicUpTo(22000).subList(1999, 2005));  //22000 will get us 2005 numbers
         System.out.println("Triplets up to 250: " + getTriplets(ludicUpTo(250)));
      }
   }`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestJavaFivenumLanguageDetection(t *testing.T) {
	t.Log("Testing Java Fivenum Language Detection")
	code := `import java.util.Arrays;
   
   public class Fivenum {
      
         static double median(double[] x, int start, int endInclusive) {
            int size = endInclusive - start + 1;
            if (size <= 0) throw new IllegalArgumentException("Array slice cannot be empty");
            int m = start + size / 2;
            return (size % 2 == 1) ? x[m] : (x[m - 1] + x[m]) / 2.0;
         }
      
         static double[] fivenum(double[] x) {
            for (Double d : x) {
               if (d.isNaN())
                     throw new IllegalArgumentException("Unable to deal with arrays containing NaN");
            }
            double[] result = new double[5];
            Arrays.sort(x);
            result[0] = x[0];
            result[2] = median(x, 0, x.length - 1);
            result[4] = x[x.length - 1];
            int m = x.length / 2;
            int lowerEnd = (x.length % 2 == 1) ? m : m - 1;
            result[1] = median(x, 0, lowerEnd);
            result[3] = median(x, m, x.length - 1);
            return result;
         }
      
         public static void main(String[] args) {
            double xl[][] = {
               {15.0, 6.0, 42.0, 41.0, 7.0, 36.0, 49.0, 40.0, 39.0, 47.0, 43.0},
               {36.0, 40.0, 7.0, 39.0, 41.0, 15.0},
               {
                     0.14082834,  0.09748790,  1.73131507,  0.87636009, -1.95059594,  0.73438555,
                     -0.03035726,  1.46675970, -0.74621349, -0.72588772,  0.63905160,  0.61501527,
                     -0.98983780, -1.00447874, -0.62759469,  0.66206163,  1.04312009, -0.10305385,
                     0.75775634,  0.32566578
               }
            };
            for (double[] x : xl) System.out.printf("%s\n\n", Arrays.toString(fivenum(x)));
         }
   }`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_JAVA {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_JAVA, result.Language)
		t.Errorf("Result: %v", result)
	}
}
