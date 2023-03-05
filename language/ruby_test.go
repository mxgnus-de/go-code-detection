package language_test

import (
	"testing"

	"github.com/mxgnus-de/go-code-detection/detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestRubyHelloWorldLanguageDetection(t *testing.T) {
	t.Log("Testing Ruby Hello World Language Detection")
	code := `puts "Hello World!"`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_RUBY {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestRubyFizzBuzzLanguageDetection(t *testing.T) {
	t.Log("Testing Ruby FizzBuzz Language Detection")
	code := `1.step(100,1) do |i|
  if (i % 5) == 0 && (i % 3) ==0
      puts 'FizzBuzz'
  elsif (i % 5) == 0
      puts 'Buzz'
  elsif (i % 3) == 0
      puts 'Fizz'
  else
      puts i
  end
end`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_RUBY {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestRubyBubbleSortLanguageDetection(t *testing.T) {
	t.Log("Testing Ruby Bubble Sort Language Detection")
	code := `class Array
    def bubblesort1!
      length.times do |j|
        for i in 1...(length - j)
          if self[i] < self[i - 1]
            self[i], self[i - 1] = self[i - 1], self[i]
          end
        end
      end
      self
    end
    def bubblesort2!
      each_index do |index|
        (length - 1).downto( index ) do |i|
          self[i-1], self[i] = self[i], self[i-1] if self[i-1] < self[i]
        end
      end
      self
    end
  end
  ary = [3, 78, 4, 23, 6, 8, 6]
  ary.bubblesort1!
  p ary`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_RUBY {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestRubyHttpRequestLanguageDetection(t *testing.T) {
	t.Log("Testing Ruby HTTP Request Language Detection")
	code := `require 'fileutils'
  require 'open-uri'
   
  open("http://rosettacode.org/") {|f| FileUtils.copy_stream(f, $stdout)}`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_RUBY {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestRubyFloydWarshallAlgorithmLanguageDetection(t *testing.T) {
	t.Log("Testing Ruby Floyd-Warshall Algorithm Language Detection")
	code := `def floyd_warshall(n, edge)
    dist = Array.new(n){|i| Array.new(n){|j| i==j ? 0 : Float::INFINITY}}
    nxt = Array.new(n){Array.new(n)}
    edge.each do |u,v,w|
      dist[u-1][v-1] = w
      nxt[u-1][v-1] = v-1
    end
  
    n.times do |k|
      n.times do |i|
        n.times do |j|
          if dist[i][j] > dist[i][k] + dist[k][j]
            dist[i][j] = dist[i][k] + dist[k][j]
            nxt[i][j] = nxt[i][k]
          end
        end
      end
    end
  
    puts "pair     dist    path"
    n.times do |i|
      n.times do |j|
        next  if i==j
        u = i
        path = [u]
        path << (u = nxt[u][j])  while u != j
        path = path.map{|u| u+1}.join(" -> ")
        puts "%d -> %d  %4d     %s" % [i+1, j+1, dist[i][j], path]
      end
    end
  end
  
  n = 4
  edge = [[1, 3, -2], [2, 1, 4], [2, 3, 3], [3, 4, 2], [4, 2, -1]]
  floyd_warshall(n, edge)`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_RUBY {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestRubyLudicNumbersLanguageDetection(t *testing.T) {
	t.Log("Testing Ruby Ludic Numbers Language Detection")
	code := `def ludic(nmax=100000)
    Enumerator.new do |y|
      y << 1
      ary = *2..nmax
      until ary.empty?
        y << (n = ary.first)
        (0...ary.size).step(n){|i| ary[i] = nil}
        ary.compact!
      end
    end
  end
  
  puts "First 25 Ludic numbers:", ludic.first(25).to_s
  
  puts "Ludics below 1000:", ludic(1000).count
  
  puts "Ludic numbers 2000 to 2005:", ludic.first(2005).last(6).to_s
  
  ludics = ludic(250).to_a
  puts "Ludic triples below 250:",
      ludics.select{|x| ludics.include?(x+2) and ludics.include?(x+6)}.map{|x| [x, x+2, x+6]}.to_s`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_RUBY {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestRubyConflictsWithOtherLanguagesLanguageDetection(t *testing.T) {
	t.Log("Testing Ruby Conflicts With Other Languages Language Detection")
	code := `module XYZ
  class A
  end
  class B < A
  end
  end`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_RUBY {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestRubyBEGINAndENDLanguageDetection(t *testing.T) {
	t.Log("Testing Ruby BEGIN and END Language Detection")

	t.Run("1", func(t *testing.T) {
		t.Log("Testing Ruby BEGIN and END Language Detection")
		code := `#!/usr/bin/ruby

  puts "This is main Ruby Program"
  
  END {
     puts "Terminating Ruby Program"
  }
  BEGIN {
     puts "Initializing Ruby Program"
  }`
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_RUBY {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
			t.Errorf("Result: %v", result)
		}
	})

	t.Run("2", func(t *testing.T) {
		t.Log("Testing Ruby BEGIN and END Language Detection")
		code := `# Ruby Program of BEGIN and END Block

  # BEGIN block
  BEGIN {
  
  a = 4
  b = 3
  c = a + b
      
  # BEGIN block code
  puts "This is BEGIN block code"
  puts c
  
  }
    
  # END block
  END {
  
  a = 4
  b = 3
  c = a * b
      
  # END block code
  puts "This is END block code"
  puts c
  }
    
  # Code will execute before END block
  puts "Main Block"
  `
		result := detection.DetectLanguage(code, detection.NewDetectionOptions())

		if result.Language != language.LANGUAGE_RUBY {
			t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_RUBY, result.Language)
			t.Errorf("Result: %v", result)
		}
	})
}
