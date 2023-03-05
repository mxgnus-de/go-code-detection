package language_test

import (
	"testing"

	"github.com/mxgnus-de/go-code-detection/detection"
	"github.com/mxgnus-de/go-code-detection/language"
)

func TestHTMLHelloWorldLanguageDetection(t *testing.T) {
	t.Log("Testing HTML Hello World Language Detection")
	code := `<!DOCTYPE html>
  <html>
    <head>
      <title>Page Title</title>
    </head>
    <body>
      <h1>Hello world</h1>
      <p>This is a tiny HTML page.</p>
    </body>
  </html>
  `
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_HTML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_HTML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestHTMLAnimationWithJS(t *testing.T) {
	t.Log("Testing HTML Animation with JS Language Detection")
	code := `<html> <head>
      <title>RC: Basic Animation</title>
      <script type="text/javascript">
         function animate(id) {
               var element = document.getElementById(id);
               var textNode = element.childNodes[0]; // assuming no other children

               var text = textNode.data;
               var reverse = false;

               element.onclick = function () { reverse = !reverse; };

               setInterval(function () {
                  if (reverse)
                     text = text.substring(1) + text[0];
                  else
                     text = text[text.length - 1] + text.substring(0, text.length - 1);
                  textNode.data = text;
               }, 100);
         }
      </script>
   </head> <body onload="animate('target')">
      <pre id="target">Hello World! </pre>
   </body> </html>`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_HTML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_HTML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestHTMLQuineWithCSSLanguageDetection(t *testing.T) {
	t.Log("Testing HTML Quine with CSS Language Detection")
	code := `<!DOCTYPE html>
   <html>
   <head>
      <title>HTML/CSS Quine</title>
      <style type="text/css">
      * { font: 10pt monospace; }
      
      head, style { display: block; }
      style { white-space: pre; }
      
      style:before {
         content:
         "<""!DOCTYPE html>"
         "A<html>A"
         "<head>A"
         "<title>""HTML/CSS Quine""</title>A"
         "<style type="text/css">";
      }
      style:after {
         content:
         "</style>A"
         "</head>A"
         "<""body></body>A"
         "</html>";
      }
      </style>
   </head>
   <body></body>
   </html>`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_HTML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_HTML, result.Language)
		t.Errorf("Result: %v", result)
	}
}

func TestHTMLCommentLanguageDetection(t *testing.T) {
	t.Log("Testing HTML Comment Language Detection")
	code := `<!-- This is a comment -->`
	result := detection.DetectLanguage(code, detection.NewDetectionOptions())

	if result.Language != language.LANGUAGE_HTML {
		t.Errorf("Expected language to be %s, but got %s", language.LANGUAGE_HTML, result.Language)
		t.Errorf("Result: %v", result)
	}
}
