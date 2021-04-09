package lexer

// Other func

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) readComment() string {
	position := l.position
	for {
		if l.peekChar() == '\r' || l.peekChar() == '\n' || l.ch == 0 {
			break
		}
		l.readChar()
	}

	return l.input[position:l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	postion := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[postion:l.position]
}

func (l *Lexer) readNumber() string {
	postion := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[postion:l.position]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
