package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Team struct {
	Name       string
	Title      string
	PhotoUrl   string
	LinkdinUrl string
}

func main() {

	companies := []string{"https://hax.co/company/algoma/", "https://hax.co/company/altiro-energy/", "https://hax.co/company/aurasense/", "https://hax.co/company/cocoon/", "https://hax.co/company/cool-amps/", "https://hax.co/company/lighthearted-ai/", "https://hax.co/company/material/", "https://hax.co/company/mitico/", "https://hax.co/company/oli/", "https://hax.co/company/pureli/", "https://hax.co/company/silana/", "https://hax.co/company/trellisense/", "https://hax.co/company/verdex-technologies/", "https://hax.co/company/3dk-tech/", "https://hax.co/company/aciist-smart-networks-ltd/", "https://hax.co/company/aeromutable/", "https://hax.co/company/aisight/", "https://hax.co/company/allozymes/", "https://hax.co/company/amatec/", "https://hax.co/company/amber-agriculture/", "https://hax.co/company/amper/", "https://hax.co/company/arculus-solutions/", "https://hax.co/company/argentum/", "https://hax.co/company/artyc/", "https://hax.co/company/atome/", "https://hax.co/company/avidbots/", "https://hax.co/company/axem-neurotechnology-inc/", "https://hax.co/company/ayrton-energy/", "https://hax.co/company/babybe-gmbh/", "https://hax.co/company/bartesian/", "https://hax.co/company/biometallica/", "https://hax.co/company/bitome-inc/", "https://hax.co/company/breezi/", "https://hax.co/company/butlr/", "https://hax.co/company/caldo-restaurant-technologies/", "https://hax.co/company/canyon-magnet/", "https://hax.co/company/carbon-bridge/", "https://hax.co/company/cargo-kite/", "https://hax.co/company/chirp/", "https://hax.co/company/chronos-dx/", "https://hax.co/company/circadia/", "https://hax.co/company/clarity/", "https://hax.co/company/cleanrobotics/", "https://hax.co/company/cradlewise/", "https://hax.co/company/criamtech/", "https://hax.co/company/darma/", "https://hax.co/company/deepspin/", "https://hax.co/company/dia/", "https://hax.co/company/digi-bio/", "https://hax.co/company/dispatch/", "https://hax.co/company/divigas/", "https://hax.co/company/drip-ai/", "https://hax.co/company/fairmart/", "https://hax.co/company/feel-therapeutics/", "https://hax.co/company/feetme/", "https://hax.co/company/flair/", "https://hax.co/company/flow-neuroscience/", "https://hax.co/company/flowbio/", "https://hax.co/company/fluidai/", "https://hax.co/company/forward-robotics-inc/", "https://hax.co/company/gaia-ai/", "https://hax.co/company/green-li-ion/", "https://hax.co/company/harae-dx/", "https://hax.co/company/hausbots/", "https://hax.co/company/hearth-labs/", "https://hax.co/company/hydrostasis/", "https://hax.co/company/hyperlume/", "https://hax.co/company/inbolt/", "https://hax.co/company/japet-medical/", "https://hax.co/company/kegg/", "https://hax.co/company/kinexcs/", "https://hax.co/company/kniterate/", "https://hax.co/company/kokoon-technology/", "https://hax.co/company/kolibri/", "https://hax.co/company/labrador-systems/", "https://hax.co/company/leadoptik/", "https://hax.co/company/lief-therapeutics/", "https://hax.co/company/lilu/", "https://hax.co/company/livin-farms/", "https://hax.co/company/lura-health/", "https://hax.co/company/makeblock/", "https://hax.co/company/mazlite/", "https://hax.co/company/mdc/", "https://hax.co/company/mechasys/", "https://hax.co/company/mesa-quantum/", "https://hax.co/company/mesh/", "https://hax.co/company/metal-light/", "https://hax.co/company/mimic-systems/", "https://hax.co/company/minut/", "https://hax.co/company/motion-metrics-limited/", "https://hax.co/company/mowito/", "https://hax.co/company/neptune-robotics/", "https://hax.co/company/neupeak-robotics/", "https://hax.co/company/neurocess/", "https://hax.co/company/neurode/", "https://hax.co/company/nordetect/", "https://hax.co/company/nuada/", "https://hax.co/company/nura/", "https://hax.co/company/openshelf/", "https://hax.co/company/opentrons/", "https://hax.co/company/particle/", "https://hax.co/company/perimeter/", "https://hax.co/company/petcube/", "https://hax.co/company/petronics/", "https://hax.co/company/pix-moving/", "https://hax.co/company/pons/", "https://hax.co/company/portable-diagnostic-systems/", "https://hax.co/company/preemadonna/", "https://hax.co/company/presso/", "https://hax.co/company/pulse-industrial/", "https://hax.co/company/pulsenics/", "https://hax.co/company/pushme/", "https://hax.co/company/q5d/", "https://hax.co/company/qnetic/", "https://hax.co/company/r-zero/", "https://hax.co/company/ray-iot/", "https://hax.co/company/reach-industries/", "https://hax.co/company/renewell-energy/", "https://hax.co/company/renovate-robotics/", "https://hax.co/company/revol-technologies-inc/", "https://hax.co/company/rightbot/", "https://hax.co/company/robo-deck/", "https://hax.co/company/robomart/", "https://hax.co/company/rockmass-technologies-inc/", "https://hax.co/company/rxall-inc/", "https://hax.co/company/samphire-neuroscience/", "https://hax.co/company/sana/", "https://hax.co/company/seppure/", "https://hax.co/company/sidework/", "https://hax.co/company/simbe-robotics/", "https://hax.co/company/skygauge-robotics/", "https://hax.co/company/smartex/", "https://hax.co/company/solano-energy/", "https://hax.co/company/somatic/", "https://hax.co/company/still-bright/", "https://hax.co/company/strados-labs/", "https://hax.co/company/striemo/", "https://hax.co/company/sungreenh2/", "https://hax.co/company/swap-robotics/", "https://hax.co/company/tandem/", "https://hax.co/company/tensorfield-agriculture/", "https://hax.co/company/terran-robotics/", "https://hax.co/company/the-last-gameboard/", "https://hax.co/company/the-vit-initiative-llc/", "https://hax.co/company/tumi-robotics/", "https://hax.co/company/unboxrobotics/", "https://hax.co/company/unicorn-biotechnologies/", "https://hax.co/company/unspun/", "https://hax.co/company/untap/", "https://hax.co/company/vertiq/", "https://hax.co/company/viabot/", "https://hax.co/company/vici-robotics/", "https://hax.co/company/voltstorage/", "https://hax.co/company/vue-smart-glasses/", "https://hax.co/company/wakecap/", "https://hax.co/company/wazer/", "https://hax.co/company/wild/", "https://hax.co/company/xfuel/", "https://hax.co/company/xn-health/", "https://hax.co/company/xrobotics/", "https://hax.co/company/yeelight/", "https://hax.co/company/zio-health/"}

	fileName := "result-scrape.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("ERROR: Could not create file %q: %s\n", fileName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write column headers of the text file
	writer.Write([]string{"company_url", "logo_url", "company_header_tag_line", "company_website", "company_social",
		"company_name", "company_details", "tags", "team_details"})

	for _, companyurl := range companies {
		fetchURL := companyurl

		// Instantiate the default Collector
		c := colly.NewCollector()

		// Before making a request, print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting: ", r.URL)
		})

		logoUrl := ""
		c.OnHTML(`.single-company-header__logo-inner`, func(e *colly.HTMLElement) {
			logoUrl = e.ChildAttr("img", "src")
		})

		companyHeaderTagline := ""
		c.OnHTML(`.single-company-header__tagline`, func(h *colly.HTMLElement) {
			companyHeaderTagline = h.Text
		})

		companyWebsite := ""
		c.OnHTML(`.single-company-header__website`, func(h *colly.HTMLElement) {
			companyWebsite = h.ChildAttr("a", "href")
		})

		companySocial := ""
		c.OnHTML(`.single-company-header__social`, func(h *colly.HTMLElement) {
			companySocial = h.ChildAttr("a", "href")
		})

		companyName := ""
		c.OnHTML(`.single-company-header__title`, func(h *colly.HTMLElement) {
			companyName = h.Text
		})

		companyDetails := ""
		c.OnHTML(`.single-company-details__content`, func(h *colly.HTMLElement) {
			companyDetails = h.ChildText("p")
		})

		tags := ""
		c.OnHTML(`.single-company-details__terms`, func(h *colly.HTMLElement) {
			h.ForEach("a", func(i int, el *colly.HTMLElement) {
				if i == 0 {
					tags = el.Text
				} else {
					tags = tags + "," + el.Text
				}

			})
		})

		var founderName []string
		var position []string
		c.OnHTML(`.single-company-team__item`, func(h *colly.HTMLElement) {
			founderName = append(founderName, h.Attr("h3"))
			position = append(position, h.ChildText("div"))
		})
		var linkdinUrl []string
		c.OnHTML(`.single-company-team__item-social`, func(h *colly.HTMLElement) {
			linkdinUrl = append(linkdinUrl, h.ChildAttr("a", "href"))
		})

		// start scraping the page under the given URL
		c.Visit(fetchURL)

		var teams []Team
		for index, identifier := range linkdinUrl {
			teams = append(teams, Team{
				Name:       founderName[index],
				Title:      position[index],
				LinkdinUrl: identifier,
			})
		}
		mTeams, _ := json.Marshal(teams)

		// 	//Write all scraped pieces of information to output text file
		writer.Write([]string{
			fetchURL,
			logoUrl,
			companyHeaderTagline,
			companyWebsite,
			companySocial,
			companyName,
			companyDetails,
			tags,
			string(mTeams),
		})
		fmt.Println("End of scraping: ", fetchURL)
	}

}
