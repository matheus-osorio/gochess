package board

type BoardSide struct {
	King  bool
	Queen bool
}

type castleRights struct {
	White BoardSide
	Black BoardSide
}

// The configuration file that holds castling rights and who is
type Configuration struct {
	WhiteTurn    bool
	CastleRights castleRights
}

func (conf Configuration) GetCastlingRights(whitePiece bool) BoardSide {
	if whitePiece {
		return conf.CastleRights.White
	} else {
		return conf.CastleRights.Black
	}
}

func (conf Configuration) ChangeTurn() Configuration {
	conf.WhiteTurn = !conf.WhiteTurn
	return conf
}

func (conf Configuration) RemoveCastlingRights(whitePiece bool) Configuration {
	side := BoardSide{
		King:  false,
		Queen: false,
	}
	if whitePiece {
		conf.CastleRights.White = side
	} else {
		conf.CastleRights.Black = side
	}

	return conf
}
