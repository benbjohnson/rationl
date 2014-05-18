package templates

import (
	"fmt"
	"github.com/benbjohnson/rationl"
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
func Index(w io.Writer, session *rationl.Session) error {
//line index.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:4
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:5
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line index.ego:6
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line index.ego:7
	head(w)
//line index.ego:8
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line index.ego:9
	if _, err := fmt.Fprintf(w, "<body class=\"index\">\n    "); err != nil {
		return err
	}
//line index.ego:10
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line index.ego:11
	if _, err := fmt.Fprintf(w, "<div class=\"header\">\n        "); err != nil {
		return err
	}
//line index.ego:12
	if _, err := fmt.Fprintf(w, "<h3 class=\"text-muted\">Rationl"); err != nil {
		return err
	}
//line index.ego:12
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line index.ego:13
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line index.ego:15
	if _, err := fmt.Fprintf(w, "<div class=\"jumbotron\">\n        "); err != nil {
		return err
	}
//line index.ego:16
	if _, err := fmt.Fprintf(w, "<h1>Experiment Tracking"); err != nil {
		return err
	}
//line index.ego:16
	if _, err := fmt.Fprintf(w, "</h1>\n        "); err != nil {
		return err
	}
//line index.ego:17
	if _, err := fmt.Fprintf(w, "<p class=\"lead\">\n            Rationl is a tool for experimentally investigating and fixing software issues.\n        "); err != nil {
		return err
	}
//line index.ego:19
	if _, err := fmt.Fprintf(w, "</p>\n        "); err != nil {
		return err
	}
//line index.ego:20
	if _, err := fmt.Fprintf(w, "<p>\n            "); err != nil {
		return err
	}
//line index.ego:21
	if _, err := fmt.Fprintf(w, "<a class=\"btn btn-lg btn-success\" href=\"/authorize\" role=\"button\">Sign in with GitHub"); err != nil {
		return err
	}
//line index.ego:21
	if _, err := fmt.Fprintf(w, "</a>\n        "); err != nil {
		return err
	}
//line index.ego:22
	if _, err := fmt.Fprintf(w, "</p>\n      "); err != nil {
		return err
	}
//line index.ego:23
	if _, err := fmt.Fprintf(w, "</div>\n\n    "); err != nil {
		return err
	}
//line index.ego:25
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line index.ego:25
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line index.ego:26
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line index.ego:27
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}

//line investigations.ego:1
func Investigations(w io.Writer, session *rationl.Session, investigations []*rationl.Investigation) error {
//line investigations.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line investigations.ego:4
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line investigations.ego:5
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line investigations.ego:6
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line investigations.ego:7
	head(w)
//line investigations.ego:8
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line investigations.ego:9
	if _, err := fmt.Fprintf(w, "<body class=\"investigations\">\n    "); err != nil {
		return err
	}
//line investigations.ego:10
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line investigations.ego:11
	if _, err := fmt.Fprintf(w, "<div class=\"page-header\">\n        "); err != nil {
		return err
	}
//line investigations.ego:12
	if _, err := fmt.Fprintf(w, "<h3>Investigations"); err != nil {
		return err
	}
//line investigations.ego:12
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line investigations.ego:13
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line investigations.ego:15
	if len(investigations) == 0 {
//line investigations.ego:16
		if _, err := fmt.Fprintf(w, "\n        "); err != nil {
			return err
		}
//line investigations.ego:16
		if _, err := fmt.Fprintf(w, "<div class=\"row\">\n          "); err != nil {
			return err
		}
//line investigations.ego:17
		if _, err := fmt.Fprintf(w, "<div class=\"col-lg-12\">\n            "); err != nil {
			return err
		}
//line investigations.ego:18
		if _, err := fmt.Fprintf(w, "<p>\n              Rationl groups your experiments together based on what you're trying\n              to investigate. For example, if you're trying to figure out why your\n              application is running slow, you might have many experiments you \n              want to try but they're all related to your application performance.\n            "); err != nil {
			return err
		}
//line investigations.ego:23
		if _, err := fmt.Fprintf(w, "</p>\n          "); err != nil {
			return err
		}
//line investigations.ego:24
		if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
			return err
		}
//line investigations.ego:25
		if _, err := fmt.Fprintf(w, "</div>\n      "); err != nil {
			return err
		}
//line investigations.ego:26
	} else {
//line investigations.ego:27
		if _, err := fmt.Fprintf(w, "\n        "); err != nil {
			return err
		}
//line investigations.ego:27
		if _, err := fmt.Fprintf(w, "<table class=\"table\">\n          "); err != nil {
			return err
		}
//line investigations.ego:28
		if _, err := fmt.Fprintf(w, "<thead>\n            "); err != nil {
			return err
		}
//line investigations.ego:29
		if _, err := fmt.Fprintf(w, "<tr>\n              "); err != nil {
			return err
		}
//line investigations.ego:30
		if _, err := fmt.Fprintf(w, "<th>Name"); err != nil {
			return err
		}
//line investigations.ego:30
		if _, err := fmt.Fprintf(w, "</th>\n              "); err != nil {
			return err
		}
//line investigations.ego:31
		if _, err := fmt.Fprintf(w, "<th>Date"); err != nil {
			return err
		}
//line investigations.ego:31
		if _, err := fmt.Fprintf(w, "</th>\n            "); err != nil {
			return err
		}
//line investigations.ego:32
		if _, err := fmt.Fprintf(w, "</tr>\n          "); err != nil {
			return err
		}
//line investigations.ego:33
		if _, err := fmt.Fprintf(w, "</thead>\n          "); err != nil {
			return err
		}
//line investigations.ego:34
		if _, err := fmt.Fprintf(w, "<tbody>\n            "); err != nil {
			return err
		}
//line investigations.ego:35
		for _, investigation := range investigations {
//line investigations.ego:36
			if _, err := fmt.Fprintf(w, "\n              "); err != nil {
				return err
			}
//line investigations.ego:36
			if _, err := fmt.Fprintf(w, "<tr>\n                "); err != nil {
				return err
			}
//line investigations.ego:37
			if _, err := fmt.Fprintf(w, "<td>\n                  "); err != nil {
				return err
			}
//line investigations.ego:38
			if _, err := fmt.Fprintf(w, "%v", investigation.GetName()); err != nil {
				return err
			}
//line investigations.ego:39
			if _, err := fmt.Fprintf(w, "\n                "); err != nil {
				return err
			}
//line investigations.ego:39
			if _, err := fmt.Fprintf(w, "</td>\n                "); err != nil {
				return err
			}
//line investigations.ego:40
			if _, err := fmt.Fprintf(w, "<td>Jan 01, 2014"); err != nil {
				return err
			}
//line investigations.ego:40
			if _, err := fmt.Fprintf(w, "</td>\n              "); err != nil {
				return err
			}
//line investigations.ego:41
			if _, err := fmt.Fprintf(w, "</tr>\n            "); err != nil {
				return err
			}
//line investigations.ego:42
		}
//line investigations.ego:43
		if _, err := fmt.Fprintf(w, "\n          "); err != nil {
			return err
		}
//line investigations.ego:43
		if _, err := fmt.Fprintf(w, "</tbody>\n        "); err != nil {
			return err
		}
//line investigations.ego:44
		if _, err := fmt.Fprintf(w, "</table>\n      "); err != nil {
			return err
		}
//line investigations.ego:45
	}
//line investigations.ego:46
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line investigations.ego:47
	if _, err := fmt.Fprintf(w, "<div class=\"row\">\n        "); err != nil {
		return err
	}
//line investigations.ego:48
	if _, err := fmt.Fprintf(w, "<div class=\"col-md-12\">\n          "); err != nil {
		return err
	}
//line investigations.ego:49
	if _, err := fmt.Fprintf(w, "<a type=\"button\" class=\"btn btn-primary\" href=\"/investigations/new\">New Investigation"); err != nil {
		return err
	}
//line investigations.ego:49
	if _, err := fmt.Fprintf(w, "</a>\n        "); err != nil {
		return err
	}
//line investigations.ego:50
	if _, err := fmt.Fprintf(w, "</div>\n      "); err != nil {
		return err
	}
//line investigations.ego:51
	if _, err := fmt.Fprintf(w, "</div>\n    "); err != nil {
		return err
	}
//line investigations.ego:52
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line investigations.ego:52
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line investigations.ego:53
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line investigations.ego:54
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}
