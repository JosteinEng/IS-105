package quote

import "rsc.io/quote"

func TestQuote() string {
	return quote.Glass() + "\n" + quote.Go()
}
