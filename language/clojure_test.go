package language_test

import (
	"testing"

	detection "github.com/mxgnus-de/go-code-detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestClojureHelloWorldLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Hello World language detection")
	code := `(binding [*out* *err*]
      (println "Hello world!"))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureGuessTheNumberLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Guess the Number language detection")
	code := `(def target (inc (rand-int 10))
   
   (loop [n 0]
      (println "Guess a number between 1 and 10 until you get it right:")
      (let [guess (read)]
      (if (= guess target)
         (printf "Correct on the %d guess.\n" n)
         (do
            (println "Try again")
            (recur (inc n))))))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureBubbleSortLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Bubble Sort language detection")
	code := `(defn- bubble-step
      "was-changed: whether any elements prior to the current first element
      were swapped;
      returns a two-element vector [partially-sorted-sequence is-sorted]"
      [less? xs was-changed]
      (if (< (count xs) 2)
         [xs (not was-changed)]
         (let [[x1 x2 & xr] xs
         first-is-smaller   (less? x1 x2)
         is-changed         (or was-changed (not first-is-smaller))
         [smaller larger]   (if first-is-smaller [x1 x2] [x2 x1])
         [result is-sorted] (bubble-step
               less? (cons larger xr) is-changed)]
         [(cons smaller result) is-sorted])))
      
   (defn bubble-sort
      "Takes an optional less-than predicate and a sequence.
      Returns the sorted sequence.
      Very inefficient (O(n²))"
      ([xs] (bubble-sort <= xs))
      ([less? xs] 
         (let [[result is-sorted] (bubble-step less? xs false)]
            (if is-sorted
      result
      (recur less? result)))))
      
   (println (bubble-sort [10 9 8 7 6 5 4 3 2 1]))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureHeadSortLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Head Sort language detection")
	code := `(defn- swap [a i j]
      (assoc a i (nth a j) j (nth a i)))

   (defn- sift [a pred k l]
      (loop [a a x k y (inc (* 2 k))]
      (if (< (inc (* 2 x)) l)
         (let [ch (if (and (< y (dec l)) (pred (nth a y) (nth a (inc y))))
                     (inc y)
                     y)]
            (if (pred (nth a x) (nth a ch))
            (recur (swap a x ch) ch (inc (* 2 ch)))
            a))
         a)))

   (defn- heapify[pred a len]
      (reduce (fn [c term] (sift (swap c term 0) pred 0 term))
            (reduce (fn [c i] (sift c pred i len))
                     (vec a)
                     (range (dec (int (/ len 2))) -1 -1))
            (range (dec len) 0 -1)))

   (defn heap-sort
      ([a pred]
      (let [len (count a)]
         (heapify pred a len)))
      ([a]
         (heap-sort a <)))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureMergeSortLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Merge Sort language detection")
	code := `(defn merge [left right]
      (cond (nil? left) right
            (nil? right) left
            :else (let [[l & *left] left
                        [r & *right] right]
                  (if (<= l r) (cons l (merge *left right))
                                 (cons r (merge left *right))))))

   (defn merge-sort [list]
      (if (< (count list) 2)
      list
      (let [[left right] (split-at (/ (count list) 2) list)]
         (merge (merge-sort left) (merge-sort right)))))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureQuickSortLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Quick Sort language detection")
	code := `(defn qsort [[pivot & xs]]
      (when pivot
      (let [smaller #(< % pivot)]
         (lazy-cat (qsort (filter smaller xs))
      [pivot]
      (qsort (remove smaller xs))))))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureIsStringNumericLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Is String Numeric language detection")
	code := `(defn numeric? [s]
      (if-let [s (seq s)]
      (let [s (if (= (first s) \\-) (next s) s)
            s (drop-while #(Character/isDigit %) s)
            s (if (= (first s) \\.) (next s) s)
            s (drop-while #(Character/isDigit %) s)]
         (empty? s))))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureLudicNumbersLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Ludic Numbers language detection")
	code := `(defn ints-from [n]
      (cons n (lazy-seq (ints-from (inc n)))))
      
   (defn drop-nth [n seq] 
      (cond 
         (zero?    n) seq
         (empty? seq) []
         :else (concat (take (dec n) seq) (lazy-seq (drop-nth n (drop n seq))))))
      
   (def ludic ((fn ludic
      ([] (ludic 1))
      ([n] (ludic n (ints-from (inc n))))
      ([n [f & r]] (cons n (lazy-seq (ludic f (drop-nth f r))))))))
      
   (defn ludic? [n]  (= (first (filter (partial <= n) ludic)) n))
      
   (print "First 25: ")
   (println (take 25 ludic))
   (print "Count below 1000: ")
   (println (count (take-while (partial > 1000) ludic)))
   (print "2000th through 2005th: ")
   (println (map (partial nth ludic) (range 1999 2005)))
   (print "Triplets < 250: ")
   (println (filter (partial every? ludic?) 
            (for [i (range 250)] (list i (+ i 2) (+ i 6)))))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureDateManipulationLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Date Manipulation language detection")
	code := `(import java.util.Date
      java.text.SimpleDateFormat)
      
   (defn time+12 [s]
      (let [sdf (SimpleDateFormat. "MMMM d yyyy h:mma zzz")]
         (-> (.parse sdf s)
      (.getTime ,)
      (+ , 43200000)
      long
      (Date. ,)
      (->> , (.format sdf ,)))))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojurePerfectShuffleLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Perfect Shuffle language detection")
	code := `(defn perfect-shuffle [deck]
      (let [half (split-at (/ (count deck) 2) deck)]
         (interleave (first half) (last half))))
      
   (defn solve [deck-size]
      (let [original (range deck-size) 
            trials (drop 1 (iterate perfect-shuffle original))
            predicate #(= original %)]
         (println (format "%5s: %s" deck-size
         (inc (some identity (map-indexed (fn [i x] (when (predicate x) i)) trials)))))))
      
   (map solve [8 24 52 100 1020 1024 10000])`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureConditinalLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Conditional language detection")
	code := `(cond
      (= 1 2) :no)
      
   (cond
      (= 1 2) :no
      (= 1 1) :yes)`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureCurryingLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Currying language detection")
	code := `(def plus-a-hundred (partial + 100))
   (assert (= 
               (plus-a-hundred 1)
               101))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojure100DoorsLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure 100 Doors language detection")
	code := `(defn doors []
      (reduce (fn [doors toggle-idx] (update-in doors [toggle-idx] not))
               (into [] (repeat 100 false))
               (for [pass   (range 1 101)
                     i      (range (dec pass) 100 pass) ]
               i)))
      
   (defn open-doors [] (for [[d n] (map vector (doors) (iterate inc 1)) :when d] n))
      
   (defn print-open-doors []
      (println 
         "Open doors after 100 passes:"
         (apply str (interpose ", " (open-doors)))))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureLoopLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Loop language detection")
	code := `(doseq [s (map #(str %1 %2 %3) \"abc\" \"ABC\" \"123\")])`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestClojureNestedLoopLanguageDetection(t *testing.T) {
	t.Log("Testing Clojure Nested Loop language detection")
	code := `(ns nested)
 
   (defn create-matrix [width height]
      (for [_ (range width)]
         (for [_ (range height)]
         (inc (rand-int 20)))))
      
   (defn print-matrix [matrix]
      (loop [[row & rs] matrix]
         (when (= (loop [[x & xs] row]
                  (cond (= x 20) :stop
                        xs (recur xs)
                        :else :continue))
                  :continue)
         (when rs (recur rs)))))
      
   (print-matrix (create-matrix 10 10))`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_CLOJURE {
		t.Errorf("Expected language to be %s, got %s", language.LANGUAGE_CLOJURE, result.Language)
		t.Errorf("Result: %v", result)
	}
}
