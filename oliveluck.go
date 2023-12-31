package oliveluck

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/castillobgr/sententia"
)

var (
	randSource rand.Source
	rng        *rand.Rand
	funcSlice  = []func() string{}
)

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
	rng = rand.New(randSource)

	funcSlice = []func() string{
		func() string {
			str1, err := sententia.Make("{{ noun }}")
			if err != nil {
				panic(err)
			}

			str2, err := sententia.Make("{{ adjective }}")
			if err != nil {
				panic(err)
			}
			return clean(str1, str2)
		},
		func() string { return cleanAndCombine(gofakeit.Animal, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.BeerName, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.BeerStyle, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.BeerYeast, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Bird, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.BuzzWord, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.CarMaker, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.CarModel, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.CarType, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.FarmAnimal, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Fruit, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Gender, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.HackerAdjective, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.HackerNoun, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.HackerVerb, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Hobby, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Hobby, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.JobLevel, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.JobTitle, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Language, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.MonthString, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.MovieGenre, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.NounAbstract, filteredColor) },
		func() string { return cleanAndCombine(gofakeit.NounAbstract, gofakeit.HackerAdjective) },
		func() string { return cleanAndCombine(gofakeit.PetName, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Pronoun, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.State, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.StreetPrefix, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.StreetSuffix, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Vegetable, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Verb, gofakeit.Adverb) },
		func() string { return cleanAndCombine(gofakeit.Verb, gofakeit.AdverbDegree) },
		func() string { return cleanAndCombine(gofakeit.Verb, gofakeit.AdverbFrequencyDefinite) },
		func() string { return cleanAndCombine(gofakeit.Verb, gofakeit.AdverbManner) },
		func() string { return cleanAndCombine(gofakeit.Verb, gofakeit.AdverbPlace) },
		func() string { return cleanAndCombine(randomdata.Noun, randomdata.Adjective) },
	}
}

func Main() int {
	test1()

	return 0
}

func cleanAndCombine(f1, f2 func() string) string {
	return clean(f1(), f2())
}

func clean(str1, str2 string) string {
	r := strings.ToLower(str2 + str1)
	str := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(r, "")

	return str
}

func filteredColor() string {
	color := gofakeit.SafeColor()
	for {
		if color != "black" {
			break
		}

		color = gofakeit.SafeColor()
	}

	return color
}

func test1() {
	i := 0
	for i < 10 {
		namer := GetRandNamer()
		names := GenRandomNames(namer, 1)
		for _, name := range names {
			fmt.Fprintf(os.Stdout, "%s\n", name)
		}
		i++
	}
}

func GenRandomNames(namer func() string, maxNames int) []string {
	seen := make(map[string]string)
	names := make([]string, 0, maxNames)

	for count := 0; count < maxNames; {
		name := namer()

		_, found := seen[name]

		if found {
			continue
		}

		names = append(names, name)

		count++
		seen[name] = name
	}

	return names
}

func GetRandNamer() func() string {
	rand := rng.Intn(len(funcSlice))

	return funcSlice[rand]
}
