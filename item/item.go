package item

type Item struct {
	Name          string
	Exp           int
	MadeWithItems []*ItemUsed
}

func (s *Item) IsCollectable() bool {
	return len(s.MadeWithItems) == 0
}

type ItemUsed struct {
	Item
	Quantity int
}

func (s *Item) CalcItemsReport(quantity int) map[string]*ItemUsed {
	currentItemMap := make(map[string]*ItemUsed)

	currentItemMap[s.Name] = &ItemUsed{
		Item:     *s,
		Quantity: quantity,
	}

	for _, inner := range s.MadeWithItems {
		itemReports := inner.CalcItemsReport(quantity * inner.Quantity)
		currentItemMap = MergeMaps(currentItemMap, itemReports)
	}

	return currentItemMap
}

func MergeMaps(a map[string]*ItemUsed, b map[string]*ItemUsed) map[string]*ItemUsed {
	merged := make(map[string]*ItemUsed)

	for key, value := range a {
		if merged[key] == nil {
			merged[key] = value
		} else {
			original := merged[key]
			original.Quantity += value.Quantity
			merged[key] = original
		}
	}

	for key, value := range b {
		if merged[key] == nil {
			merged[key] = value
		} else {
			original := merged[key]
			original.Quantity += value.Quantity
			merged[key] = original
		}
	}

	return merged
}

func NewItemUsed(item Item, quantity int) *ItemUsed {
	return &ItemUsed{
		Item:     item,
		Quantity: quantity,
	}
}

// Reagentes
var Fundente = Item{
	Name: "Fundente",
}

var Lixa = Item{
	Name: "Lixa",
}

var Tanino = Item{
	Name: "Tanino",
}

var Tecedura = Item{
	Name: "Tecedura",
}

// Minerios
var MinerioDeFerro = Item{
	Name: "Minério de Ferro",
}

var MinerioDePrata = Item{
	Name: "Minério de Prata",
}

var MinerioDeOuro = Item{
	Name: "Minério de Ouro",
}

var MinerioDePlatina = Item{
	Name: "Minério de Platina",
}

var MinerioDeEstelaco = Item{
	Name: "Minério de Estelaço",
}

var MinerioDeOricalco = Item{
	Name: "Minério de Oricalco",
}

var Tolvium = Item{
	Name: "Tolvium",
}

var Cinabrio = Item{
	Name: "Cinábrio",
}

// Madeiras
var MadeiraVerde = Item{
	Name: "Madeira Verde",
}

var MadeiraEnvelhecida = Item{
	Name: "Madeira Envelhecida",
}

var Urdeira = Item{
	Name: "Urdeira",
}

var Juca = Item{
	Name: "Jucá",
}

var Tabua = Item{
	Name: "Tábua",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(MadeiraVerde, 4),
	},
}

var Lenha = Item{
	Name: "Lenha",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(MadeiraEnvelhecida, 4),
		NewItemUsed(Tabua, 2),
		NewItemUsed(Lixa, 1),
	},
}

var TabuaDeUrdeira = Item{
	Name: "Tábua de Urdeira",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(Urdeira, 6),
		NewItemUsed(Lenha, 2),
		NewItemUsed(Lixa, 1),
	},
}

var TabuaDeJuca = Item{
	Name: "Tábua de Jucá",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(Juca, 8),
		NewItemUsed(TabuaDeUrdeira, 2),
		NewItemUsed(Lixa, 1),
	},
}

var Carvao = Item{
	Name: "Carvão",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(MadeiraVerde, 2),
	},
}

// Couros
var CouroCru = Item{
	Name: "Couro Cru",
}

var PeleGrossa = Item{
	Name: "Pele Grossa",
}

var PeleDeFerro = Item{
	Name: "Pele de Ferro",
}

var PeleArdente = Item{
	Name: "Pele Ardente",
}

var PeleCicatrizada = Item{
	Name: "Pele Cicatrizada",
}

var CouroGrosseiro = Item{
	Name: "Couro Grosseito",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(CouroCru, 4),
	},
}

var CouroRustico = Item{
	Name: "Couro Rústico",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(CouroGrosseiro, 4),
		NewItemUsed(Tanino, 1),
	},
}

var CouroReforcado = Item{
	Name: "Couro Reforçado",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(PeleGrossa, 6),
		NewItemUsed(CouroRustico, 2),
		NewItemUsed(Tanino, 1),
	},
}

var CouroInfuso = Item{
	Name: "Couro Reforçado",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(PeleDeFerro, 8),
		NewItemUsed(CouroReforcado, 2),
		NewItemUsed(Tanino, 1),
	},
}

var CouroRunico = Item{
	Name: "Couro Rúnico",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(CouroInfuso, 5),
		NewItemUsed(PeleArdente, 1),
		NewItemUsed(PeleCicatrizada, 1),
		NewItemUsed(Tanino, 1),
	},
}

// Tecidos
var Fibra = Item{
	Name: "Fibra",
}

var FioDeSeda = Item{
	Name: "Fio De Seda",
}

var FibraDeFio = Item{
	Name: "Fibra de Fio",
}

var TecidoDeEscama = Item{
	Name: "Tecido de Escama",
}

var TramaCalica = Item{
	Name: "Trama Cálica",
}

var Linho = Item{
	Name: "Linho",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(Fibra, 4),
	},
}

var Cetim = Item{
	Name: "Cetim",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(Linho, 4),
		NewItemUsed(Tecedura, 1),
	},
}

var Seda = Item{
	Name: "Seda",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(FioDeSeda, 6),
		NewItemUsed(Cetim, 2),
		NewItemUsed(Tecedura, 1),
	},
}

var SedaInfusa = Item{
	Name: "Seda Infusa",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(FibraDeFio, 8),
		NewItemUsed(Seda, 2),
		NewItemUsed(Tecedura, 1),
	},
}

var FioDeFenix = Item{
	Name: "Fio De Fenix",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(SedaInfusa, 5),
		NewItemUsed(TecidoDeEscama, 1),
		NewItemUsed(TramaCalica, 1),
		NewItemUsed(Tecedura, 1),
	},
}

// Lingotes
var LingoteDeFerro = Item{
	Name: "Lingote de Ferro",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(MinerioDeFerro, 4),
	},
}

var LingoteDePrata = Item{
	Name: "Lingote de Prata",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(MinerioDePrata, 4),
	},
}

var LingoteDeAco = Item{
	Name: "Lingote de Aço",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(LingoteDeFerro, 3),
		NewItemUsed(Fundente, 1),
		NewItemUsed(Carvao, 2),
	},
}

var LingoteDeOuro = Item{
	Name: "Lingote de Ouro",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(MinerioDeOuro, 5),
		NewItemUsed(LingoteDePrata, 2),
		NewItemUsed(Fundente, 1),
	},
}

var LingoteDeEstelaco = Item{
	Name: "Lingote de Estelaço",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(MinerioDeEstelaco, 6),
		NewItemUsed(LingoteDeAco, 2),
		NewItemUsed(Fundente, 1),
		NewItemUsed(Carvao, 2),
	},
}

var LingoteDePlatina = Item{
	Name: "Lingote de Platina",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(MinerioDePlatina, 6),
		NewItemUsed(LingoteDeOuro, 2),
		NewItemUsed(Fundente, 1),
	},
}

var LingoteDeOricalco = Item{
	Name: "Lingote de Oricalco",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(MinerioDeOricalco, 8),
		NewItemUsed(LingoteDeEstelaco, 2),
		NewItemUsed(Fundente, 1),
		NewItemUsed(Carvao, 2),
	},
}

var LingoteDeAsmodeo = Item{
	Name: "Lingote de Asmódeo",
	MadeWithItems: []*ItemUsed{
		NewItemUsed(LingoteDeOricalco, 5),
		NewItemUsed(Tolvium, 1),
		NewItemUsed(Cinabrio, 1),
		NewItemUsed(Fundente, 1),
		NewItemUsed(Carvao, 2),
	},
}

// Machados
var MachadaoDeFerro = Item{
	Name: "Machadão de Ferro",
	Exp:  204,
	MadeWithItems: []*ItemUsed{
		NewItemUsed(LingoteDeFerro, 12),
		NewItemUsed(Tabua, 3),
		NewItemUsed(CouroGrosseiro, 2),
	},
}

var MachadaoDeAco = Item{
	Name: "Machadão de Aço",
	Exp:  540,
	MadeWithItems: []*ItemUsed{
		NewItemUsed(LingoteDeAco, 12),
		NewItemUsed(Tabua, 3),
		NewItemUsed(CouroGrosseiro, 2),
	},
}
