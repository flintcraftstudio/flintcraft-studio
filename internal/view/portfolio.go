package view

// ProjectQuote holds a client testimonial.
type ProjectQuote struct {
	Text        string
	Attribution string
	Portrait    string
}

// ProjectResult holds a single metric.
type ProjectResult struct {
	Value string
	Unit  string
	Label string
}

// ProjectImage is a single work photo shown in the project gallery.
type ProjectImage struct {
	Src string
	Alt string
}

// Project holds all data for a portfolio page.
type Project struct {
	Slug             string
	Client           string
	Title            string
	Subtitle         string
	Location         string
	Tags             []string
	LiveURL          string
	Screenshot       string
	ScreenshotMobile string
	HeroVideo        string // optional mp4; takes precedence over Screenshot in the desktop frame
	HeroVideoWebm    string // optional webm; preferred source if browser supports it
	Story            []string
	Quote            *ProjectQuote
	Results          []ProjectResult
	Gallery          []ProjectImage // optional work photos; rendered as a "The work" section
	Deliverables     []string
	NextSlug         string
	NextClient       string
}

// ProjectOrder defines the display order on the work listing page.
var ProjectOrder = []string{
	"schluters-metal-art",
	"rockabilly-roasting",
	"lo-mo-outfitting",
	"nautilus-group",
	"a-team-asphalt",
	"a-team-gutters",
	"traver-hardwood-floors",
}

// Projects is the canonical list of portfolio entries.
var Projects = map[string]Project{
	"schluters-metal-art": {
		Slug:             "schluters-metal-art",
		Client:           "Schluter’s Metal Art",
		Title:            "Custom-only metalwork. A site that’s custom too.",
		Subtitle:         "Custom site, brand, and a self-serve photo gallery for a one-man fabrication shop in Helena.",
		Location:         "Custom metal fabrication · Helena, MT",
		Tags:             []string{"Marketing site", "Brand", "SEO"},
		LiveURL:          "https://schlutersmetalart.com",
		Screenshot:       "/static/images/portfolio/schluters-metal-art_screenshot.webp",
		ScreenshotMobile: "/static/images/portfolio/schluters-metal-art_screenshot_mobile.webp",
		NextSlug:         "rockabilly-roasting",
		NextClient:       "Rockabilly Roasting",
		Story: []string{
			"Colton Schluter has been fabricating custom metalwork in Helena since 2012 — railings, gates, structural steel, decorative pieces, all of it one-off. Online, he had a Facebook page and not much else. He’d had a website once, but he lost it, and access to every photo of his old work went with it — years of finished projects, gone from his control.",
			"We built him a new one from the ground up — brand, copy, and a custom site — and made sure the photos couldn’t disappear again. It runs on cloud storage Colton owns permanently, with a backend built so he can upload, organize, and manage the gallery himself. No agency in the loop, no platform holding his work hostage.",
			"It fits the shop. Nothing Colton makes comes off a shelf, and neither does the site. “Shaped, welded & coated in Helena” — the work is custom-only, and now the thing that shows it off is too.",
		},
		Gallery: []ProjectImage{
			{Src: "/static/images/portfolio/schluters-fabrication-1.webp", Alt: "Custom steel fabrication by Schluter's Metal Art"},
			{Src: "/static/images/portfolio/schluters-fabrication-2.webp", Alt: "Custom steel fabrication by Schluter's Metal Art"},
			{Src: "/static/images/portfolio/schluters-staircase-1.webp", Alt: "Custom metal staircase by Schluter's Metal Art"},
			{Src: "/static/images/portfolio/schluters-staircase-2.webp", Alt: "Custom metal staircase by Schluter's Metal Art"},
			{Src: "/static/images/portfolio/schluters-furniture-1.webp", Alt: "Custom metal furniture by Schluter's Metal Art"},
			{Src: "/static/images/portfolio/schluters-furniture-2.webp", Alt: "Custom metal furniture by Schluter's Metal Art"},
			{Src: "/static/images/portfolio/schluters-offroad-1.webp", Alt: "Off-road metal fabrication by Schluter's Metal Art"},
			{Src: "/static/images/portfolio/schluters-powdercoat-1.webp", Alt: "Powder-coated metalwork by Schluter's Metal Art"},
		},
		Deliverables: []string{
			"Custom-designed marketing site — no templates, no page builders",
			"Brand identity — logo, color palette, typography",
			"Custom photo-management backend — upload, organize, and update the gallery directly",
			"Permanent cloud storage — every project photo owned and controlled by the shop",
			"Filterable project gallery — railings, gates, structural, decorative",
			"Full copywriting — every page written from scratch",
			"SEO foundation — meta tags, schema markup, sitemap",
			"Hosting and SSL on FlintCraft infrastructure",
		},
	},
	"nautilus-group": {
		Slug:             "nautilus-group",
		Client:           "Nautilus Group Cleaning Services",
		Title:            "No web presence. Now they have a pipeline.",
		Subtitle:         "Custom site, brand, and SEO for a cleaning company starting from zero.",
		Location:         "Commercial cleaning \u00b7 Helena, MT",
		Tags:             []string{"Marketing site", "Brand", "SEO"},
		LiveURL:          "https://nautiluscleaning.com",
		Screenshot:       "/static/images/portfolio/nautilus_screenshot.webp",
		ScreenshotMobile: "/static/images/portfolio/nautilus_screenshot_mobile.webp",
		NextSlug:         "a-team-asphalt",
		NextClient:       "A-Team Asphalt",
		Story: []string{
			"Nautilus had nothing \u2014 no site, no listings, no digital footprint. A husband-and-wife team with real credentials and real work ethic, but completely invisible online. They needed everything: a brand identity, a website, a Google Business Profile, and an SEO foundation that would start generating leads.",
			"We built it all from scratch. The site was designed to communicate trustworthiness and professionalism \u2014 insured, bonded, licensed \u2014 without looking corporate. The copy speaks directly to property managers and business owners who need reliable commercial cleaning.",
			"Within four weeks of launch they had received a commercial contract and five new inquiries. They had nothing before. Now they have a pipeline.",
		},
		Quote: &ProjectQuote{
			Text:        "We went from invisible to getting calls. That\u2019s what we needed.",
			Attribution: "Nautilus Group Cleaning Services",
		},
		Results: []ProjectResult{
			{Value: "0", Unit: "\u21921", Label: "Web presence built from nothing"},
			{Value: "4", Unit: "wks", Label: "Launch to first commercial contract"},
			{Value: "5", Unit: "leads", Label: "Inquiries in month one"},
		},
		Deliverables: []string{
			"Custom-designed marketing site \u2014 no templates, no page builders",
			"Brand identity \u2014 logo, color palette, typography",
			"Full copywriting \u2014 every page written from scratch",
			"Google Business Profile setup and optimization",
			"SEO foundation \u2014 meta tags, schema markup, sitemap",
			"Hosting and SSL on FlintCraft infrastructure",
			"Ongoing maintenance and content updates",
		},
	},
	"a-team-asphalt": {
		Slug:             "a-team-asphalt",
		Client:           "A-Team Asphalt",
		Title:            "Twenty years of reputation. A website that finally matched it.",
		Subtitle:         "Full custom site with brand identity, copy, and SEO for a Pacific Northwest paving company.",
		Location:         "Asphalt & paving \u00b7 Sumner, WA",
		Tags:             []string{"Marketing site", "SEO"},
		LiveURL:          "https://ateamasphalt.com",
		Screenshot:       "/static/images/portfolio/ateam_screenshot.webp",
		ScreenshotMobile: "/static/images/portfolio/ateam_screenshot_mobile.webp",
		NextSlug:         "a-team-gutters",
		NextClient:       "A-Team Gutters",
		Story: []string{
			"Don and Gary had built a real business the hard way \u2014 word of mouth, quality work, repeat clients across the Puyallup Valley. Their website hadn\u2019t kept up. It didn\u2019t reflect the scale of the operation or the quality of the work.",
			"We rebuilt it from the ground up \u2014 brand, copy, and code \u2014 and fixed an NAP inconsistency that was quietly suppressing their local search ranking. The new site was designed to convert: clear service pages, strong calls to action, and copy that speaks to both commercial property managers and residential homeowners.",
			"The site launched in six weeks. Every line of code was written for them. Every word of copy was written about their specific business, their specific market, their specific strengths.",
		},
		Quote: &ProjectQuote{
			Text:        "They understood what we do and built something we\u2019re actually proud to hand out.",
			Attribution: "Don \u2014 A-Team Asphalt",
		},
		Results: []ProjectResult{
			{Value: "6", Unit: "wks", Label: "Discovery to launch"},
			{Value: "0", Unit: "templates", Label: "Every line custom"},
			{Value: "1", Unit: "call", Label: "To reach the builder"},
		},
		Deliverables: []string{
			"Custom-designed marketing site \u2014 dark, industrial aesthetic matching the trade",
			"Brand identity \u2014 logo refinement, color system, typography",
			"Full copywriting \u2014 service pages, about, contact",
			"SEO strategy \u2014 NAP correction, local search optimization",
			"Google Business Profile setup",
			"Email infrastructure",
			"Hosting, SSL, and ongoing maintenance",
		},
	},
	"a-team-gutters": {
		Slug:             "a-team-gutters",
		Client:           "A-Team Gutters",
		Title:            "New trade, new brand, built to generate leads from day one.",
		Subtitle:         "Brand identity and custom site for a new gutter installation company.",
		Location:         "Gutter installation \u00b7 Bonney Lake, WA",
		Tags:             []string{"Marketing site", "Brand"},
		LiveURL:          "https://ateamgutter.com",
		Screenshot:       "/static/images/portfolio/ateam_gutters_screenshot.webp",
		ScreenshotMobile: "/static/images/portfolio/ateam_gutters_screenshot_mobile.webp",
		NextSlug:         "traver-hardwood-floors",
		NextClient:       "Traver Hardwood Floors",
		Story: []string{
			"A-Team Gutters launched in 2024 with strong trade credentials and no web presence. They needed a brand and a site that would generate residential leads in a competitive Pacific Northwest market from day one.",
			"We built the brand from the ground up \u2014 identity, copy, and a custom site designed around the specific services they offer and the specific areas they serve. No templates, no generic contractor design. A site built for this business, in this market.",
		},
		Results: []ProjectResult{
			{Value: "0", Unit: "\u21921", Label: "Brand created from scratch"},
			{Value: "1", Unit: "site", Label: "Custom built"},
			{Value: "2024", Unit: "", Label: "Launch year"},
		},
		Deliverables: []string{
			"Brand identity \u2014 logo, colors, typography",
			"Custom-designed marketing site",
			"Copywriting \u2014 all pages",
			"SEO foundation",
			"Hosting and SSL",
		},
	},
	"traver-hardwood-floors": {
		Slug:             "traver-hardwood-floors",
		Client:           "Traver Hardwood Floors",
		Title:            "Every board placed with intent. Now the site is too.",
		Subtitle:         "Chris Traver has been laying hardwood floors in western Montana for 18 years. The redesign started with the floors themselves.",
		Location:         "Hardwood flooring \u00b7 Helena, MT",
		Tags:             []string{"Marketing site", "Brand", "Design"},
		LiveURL:          "https://traverhardwoodfloors.com",
		Screenshot:       "/static/images/portfolio/thf_screenshot.webp",
		ScreenshotMobile: "/static/images/portfolio/thf_screenshot_mobile.webp",
		NextSlug:         "schluters-metal-art",
		NextClient:       "Schluter’s Metal Art",
		Story: []string{
			"Chris Traver has been laying hardwood floors in western Montana for 18 years. His site didn\u2019t show it. The previous design was functional. But it read like a flooring store, not a craftsman. On a site where the portfolio is the only argument that matters, there were two projects in it.",
			"The redesign started with the floors themselves \u2014 warm espresso darks, amber light, a palette pulled from aged fir and walnut endgrain. The typography got weight. The layout got wide enough to let photography breathe. Dark sections alternate with light ones so the page has pace, not just content.",
			"A design critique after the first build caught five things worth fixing before it shipped. The logo had gone missing. The About section was doing double duty. The service cards had no photography on a site built around photography. Small things. The kind that matter.",
			"The animations are themed around the job. Each element on the hero slides in with a two-pixel overshoot before settling \u2014 like a board nudging into its neighbor and locking flush. The proof bar counts up. The amber border on a service card sweeps left to right on hover. You notice it without noticing you noticed it.",
		},
		Results: []ProjectResult{
			{Value: "18", Unit: "yrs", Label: "Experience, now visible"},
			{Value: "0", Unit: "deps", Label: "Vanilla JS, no libraries"},
			{Value: "5", Unit: "fixes", Label: "Caught in design critique"},
		},
		Deliverables: []string{
			"Full site redesign \u2014 dark/light alternating sections, photography-first layout",
			"Color palette derived from hardwood \u2014 espresso, amber, aged fir tones",
			"Custom animations \u2014 board-settle overshoot, counting proof bar, sweep hovers",
			"Typography overhaul \u2014 Playfair Display + Source Sans 3",
			"Portfolio expansion and photography integration",
			"Hugo static site with Go API, Caddy, Docker, GitHub Actions",
			"Zero-dependency frontend \u2014 IntersectionObserver, requestAnimationFrame",
		},
	},
	"lo-mo-outfitting": {
		Slug:             "lo-mo-outfitting",
		Client:           "Lo Mo Outfitting",
		Title:            "Born on the river. Now the site feels like it.",
		Subtitle:         "Custom site, brand voice, and booking flow for a Missouri River fly fishing outfitter run by guides who actually live there.",
		Location:         "Fly fishing outfitter · Craig, Montana",
		Tags:             []string{"Marketing site", "Brand", "Design"},
		LiveURL:          "https://lomooutfitting.com",
		Screenshot:       "/static/images/portfolio/lomo_screenshot.webp",
		ScreenshotMobile: "/static/images/portfolio/lomo_screenshot_mobile.webp",
		NextSlug:         "nautilus-group",
		NextClient:       "Nautilus Group",
		Story: []string{
			"Matt Mohar has been fishing the Missouri for over twenty years. He started guiding after a decade in concrete — not because someone told him fly fishing was a good brand play, but because taking people fishing is what he wanted to do. Lo Mo runs seven licensed guides, all local, all on the water in their off hours. The operation didn't need inventing. It needed a site that matched.",
			"The previous site wasn't telling that story. We rebuilt it around what makes Lo Mo different: no booking platforms, no middlemen, no pretense. Matt answers the phone. He picks the guide. He picks the stretch of river. The site needed to communicate that directness without losing the weight of twenty years and thirty-five miles of blue-ribbon water.",
			"The design pulls from the river itself — deep slate greens, muted golds, photography that breathes. The navigation puts trips and guides up front. The copy is direct and honest, the way Matt talks. Pricing is visible, not buried. The heroes program and seasonal rates are surfaced, not hidden behind a contact form.",
		},
		Quote: &ProjectQuote{
			Text:        "Booking has never been more simple.",
			Attribution: "Matt Mohar — Lo Mo Outfitting",
			Portrait:    "/static/images/portfolio/matt-mohar.webp",
		},
		Results: []ProjectResult{
			{Value: "20", Unit: "yrs", Label: "On the Missouri River"},
			{Value: "7", Unit: "guides", Label: "All local, all licensed"},
			{Value: "5K+", Unit: "trout/mi", Label: "Blue-ribbon water"},
		},
		Deliverables: []string{
			"Custom-designed marketing site — dark, river-inspired palette",
			"Brand voice and full copywriting — every page written from scratch",
			"Trip booking flow — pricing, seasonal rates, heroes program",
			"Guide profiles — individual bios and photography",
			"SEO foundation — meta tags, schema markup, sitemap",
			"Hosting and SSL on FlintCraft infrastructure",
			"Ongoing maintenance and content updates",
		},
	},
	"rockabilly-roasting": {
		Slug:             "rockabilly-roasting",
		Client:           "Rockabilly Roasting",
		Title:            "Inherited a plugin store. Replaced it with one worthy of the roast.",
		Subtitle:         "Custom e-commerce, subscription management, and brand identity for a small-batch coffee roaster in eastern Washington.",
		Location:         "Coffee roaster · Kennewick, WA",
		Tags:             []string{"E-commerce", "Marketing site", "Brand"},
		LiveURL:          "https://rockabillyroasting.com",
		Screenshot:       "/static/images/portfolio/rockabilly_screenshot.webp",
		ScreenshotMobile: "/static/images/portfolio/rockabilly_screenshot_mobile.webp",
		NextSlug:         "lo-mo-outfitting",
		NextClient:       "Lo Mo Outfitting",
		Story: []string{
			"The previous storefront was inherited — a stock WooCommerce build that made every part of the customer journey harder than it needed to be. Sign-up friction. Subscription management buried behind plugin defaults. A checkout flow that didn’t feel like the brand it was selling. For a roaster pulling genuinely great coffee, the storefront was the weakest link in the operation.",
			"We rebuilt it from the ground up — storefront, checkout, and subscription handling, all custom. No WooCommerce. No plugin sprawl. Designed around how a small-batch roaster actually sells coffee: one-time bags, recurring shipments, gift subscriptions, account management that takes one click instead of three.",
			"And the brand itself got the treatment it deserved. Identity, typography, photography — all rebuilt to match the quality in the bag. A storefront finally worthy of the coffee.",
		},
		Quote: &ProjectQuote{
			Text:        "FlintCraft Studio is an absolutely must have company working for you in your corner. They are super approachable, helpful and friendly. If you are looking to take your company to the next level I would highly recommend you give FlintCraft a call.",
			Attribution: "Kagen Cox · Owner, Rockabilly Roasting Co.",
		},
		Deliverables: []string{
			"Custom-designed e-commerce storefront — no WooCommerce, no plugins",
			"Custom subscription management — sign-up, pause, modify, cancel",
			"Custom checkout flow — one-time, recurring, and gift purchases",
			"Brand identity — logo, color palette, typography",
			"Product photography integration",
			"SEO foundation — meta tags, schema markup, sitemap",
			"Hosting and SSL on FlintCraft infrastructure",
			"Ongoing maintenance and content updates",
		},
	},
}
