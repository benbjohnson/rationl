package rationl

import (
	"fmt"
	"io"
)

//line head.ego:1
func head(w io.Writer) error {
//line head.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line head.ego:3
	if _, err := fmt.Fprintf(w, "<head>\n  "); err != nil {
		return err
	}
//line head.ego:4
	if _, err := fmt.Fprintf(w, "<meta charset=\"utf-8\">\n  "); err != nil {
		return err
	}
//line head.ego:5
	if _, err := fmt.Fprintf(w, "<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n  "); err != nil {
		return err
	}
//line head.ego:6
	if _, err := fmt.Fprintf(w, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n  "); err != nil {
		return err
	}
//line head.ego:7
	if _, err := fmt.Fprintf(w, "<title>Rationl"); err != nil {
		return err
	}
//line head.ego:7
	if _, err := fmt.Fprintf(w, "</title>\n  "); err != nil {
		return err
	}
//line head.ego:8
	if _, err := fmt.Fprintf(w, "<link href=\"/assets/bootstrap.min.css\" rel=\"stylesheet\">\n  "); err != nil {
		return err
	}
//line head.ego:9
	if _, err := fmt.Fprintf(w, "<link href=\"/assets/application.css\" rel=\"stylesheet\">\n  "); err != nil {
		return err
	}
//line head.ego:10
	if _, err := fmt.Fprintf(w, "<script src=\"/assets/jquery-2.1.0.min.js\">"); err != nil {
		return err
	}
//line head.ego:10
	if _, err := fmt.Fprintf(w, "</script>\n  "); err != nil {
		return err
	}
//line head.ego:11
	if _, err := fmt.Fprintf(w, "<script src=\"/assets/bootstrap.min.js\">"); err != nil {
		return err
	}
//line head.ego:11
	if _, err := fmt.Fprintf(w, "</script>\n"); err != nil {
		return err
	}
//line head.ego:12
	if _, err := fmt.Fprintf(w, "</head>\n"); err != nil {
		return err
	}
	return nil
}

//line index.ego:1
func index(w io.Writer, session *Session) error {
//line index.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:3
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line index.ego:4
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line index.ego:5
	head(w)
//line index.ego:6
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line index.ego:7
	if _, err := fmt.Fprintf(w, "<body class=\"index\">\n    "); err != nil {
		return err
	}
//line index.ego:8
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line index.ego:9
	if _, err := fmt.Fprintf(w, "<div class=\"header\">\n        "); err != nil {
		return err
	}
//line index.ego:10
	if _, err := fmt.Fprintf(w, "<h3 class=\"text-muted\">Rationl"); err != nil {
		return err
	}
//line index.ego:10
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line index.ego:11
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line index.ego:13
	if _, err := fmt.Fprintf(w, "<div class=\"jumbotron\">\n        "); err != nil {
		return err
	}
//line index.ego:14
	if _, err := fmt.Fprintf(w, "<h1>Experiment Tracking"); err != nil {
		return err
	}
//line index.ego:14
	if _, err := fmt.Fprintf(w, "</h1>\n        "); err != nil {
		return err
	}
//line index.ego:15
	if _, err := fmt.Fprintf(w, "<p class=\"lead\">\n            Rationl is a tool for experimentally investigating and fixing software issues.\n        "); err != nil {
		return err
	}
//line index.ego:17
	if _, err := fmt.Fprintf(w, "</p>\n        "); err != nil {
		return err
	}
//line index.ego:18
	if _, err := fmt.Fprintf(w, "<p>\n            "); err != nil {
		return err
	}
//line index.ego:19
	if _, err := fmt.Fprintf(w, "<a class=\"btn btn-lg btn-success\" href=\"/signup\" role=\"button\">Sign in with GitHub"); err != nil {
		return err
	}
//line index.ego:19
	if _, err := fmt.Fprintf(w, "</a>\n        "); err != nil {
		return err
	}
//line index.ego:20
	if _, err := fmt.Fprintf(w, "</p>\n      "); err != nil {
		return err
	}
//line index.ego:21
	if _, err := fmt.Fprintf(w, "</div>\n\n    "); err != nil {
		return err
	}
//line index.ego:23
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line index.ego:23
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line index.ego:24
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line index.ego:25
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}
