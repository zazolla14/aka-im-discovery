package db

import (
	"fmt"

	"gorm.io/gorm"

	entity "github.com/1nterdigital/aka-im-discover/internal/model"
)

const (
	PositionDefault1 = 1
	PositionDefault2 = 2
	PositionDefault3 = 3

	PositionFacebook      = 1 + iota // 1
	PositionTwitter                  // 2
	PositionInstagram                // 3
	PositionLinkedIn                 // 4
	PositionYouTube                  // 5
	PositionReddit                   // 6
	PositionPinterest                // 7
	PositionTikTok                   // 8
	PositionSnapchat                 // 9
	PositionWhatsApp                 // 10
	PositionTelegram                 // 11
	PositionGitHub                   // 12
	PositionStackOverflow            // 13
	PositionMedium                   // 14
	PositionNetflix                  // 15
	PositionAmazon                   // 16
	PositionEbay                     // 17
	PositionWikipedia                // 18
	PositionGoogle                   // 19
	PositionYahoo                    // 20
)

func intPtr(i int) *int {
	return &i
}

func InitData(gormDB *gorm.DB) error {
	if err := seedCarousels(gormDB); err != nil {
		return err
	}

	if err := seedArticles(gormDB); err != nil {
		return err
	}

	return nil
}

func seedCarousels(gormDB *gorm.DB) error {
	var count int64
	gormDB.Model(&entity.DiscoverCarousels{}).Count(&count)
	if count > 0 {
		return nil
	}

	//nolint:lll // long URL
	seeds := []entity.DiscoverCarousels{
		{
			Title:     "Default Carousel 1",
			ImageURL:  "https://media.istockphoto.com/id/473082752/id/foto/bunga-segar-dalam-es-krim-kerucut-masih-hidup.jpg?s=1024x1024&w=is&k=20&c=bRgYey2MywT7iJFE7W-FtMR9XIpxRd-PM0oqpzDFr2I=",
			LinkURL:   "https://www.facebook.com",
			Position:  intPtr(PositionDefault1),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Default Carousel 2",
			ImageURL:  "https://media.istockphoto.com/id/1136116781/id/foto/telur-paskah-bunga-berwarna-warni-di-latar-belakang-biru-pastel-paskah-konsep-musim-semi-rata.jpg?s=2048x2048&w=is&k=20&c=3pPm6EMFmyeY3CyifRxyOH2DU2dViRe1R-50xNpH1-I=",
			LinkURL:   "https://www.twitter.com",
			Position:  intPtr(PositionDefault2),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Default Carousel 3",
			ImageURL:  "https://media.istockphoto.com/id/1148668425/id/foto/celengan-merah-muda-kecil-terikat-di-atas-mobil-tua.jpg?s=2048x2048&w=is&k=20&c=5kPK-ReLQom11iZC-zPfxEd1a1dQ_2JY07J7zVnApKA=",
			LinkURL:   "https://www.instagram.com",
			Position:  intPtr(PositionDefault3),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
	}

	if err := gormDB.Create(&seeds).Error; err != nil {
		return fmt.Errorf("failed to seed carousels: %w", err)
	}
	return nil
}

func seedArticles(gormDB *gorm.DB) error {
	var count int64
	gormDB.Model(&entity.DiscoverArticles{}).Count(&count)
	if count > 0 {
		return nil // already seeded
	}

	seeds := []entity.DiscoverArticles{
		{
			Title:     "Facebook",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626269.png",
			LinkURL:   "https://www.facebook.com",
			Position:  intPtr(PositionFacebook),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Twitter",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626271.png",
			LinkURL:   "https://www.twitter.com",
			Position:  intPtr(PositionTwitter),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Instagram",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626270.png",
			LinkURL:   "https://www.instagram.com",
			Position:  intPtr(PositionInstagram),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "LinkedIn",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626273.png",
			LinkURL:   "https://www.linkedin.com",
			Position:  intPtr(PositionLinkedIn),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "YouTube",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626292.png",
			LinkURL:   "https://www.youtube.com",
			Position:  intPtr(PositionYouTube),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Reddit",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626300.png",
			LinkURL:   "https://www.reddit.com",
			Position:  intPtr(PositionReddit),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Pinterest",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626275.png",
			LinkURL:   "https://www.pinterest.com",
			Position:  intPtr(PositionPinterest),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "TikTok",
			ImageURL:  "https://cdn-icons-png.freepik.com/512/3046/3046121.png",
			LinkURL:   "https://www.tiktok.com",
			Position:  intPtr(PositionTikTok),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Snapchat",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626276.png",
			LinkURL:   "https://www.snapchat.com",
			Position:  intPtr(PositionSnapchat),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "WhatsApp",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626279.png",
			LinkURL:   "https://www.whatsapp.com",
			Position:  intPtr(PositionWhatsApp),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Telegram",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626281.png",
			LinkURL:   "https://telegram.org",
			Position:  intPtr(PositionTelegram),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "GitHub",
			ImageURL:  "https://cdn-icons-png.freepik.com/512/2175/2175377.png",
			LinkURL:   "https://github.com",
			Position:  intPtr(PositionGitHub),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Stack Overflow",
			ImageURL:  "https://cdn-icons-png.freepik.com/256/2626/2626299.png",
			LinkURL:   "https://stackoverflow.com",
			Position:  intPtr(PositionStackOverflow),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Medium",
			ImageURL:  "https://cdn-icons-png.freepik.com/512/2504/2504925.png",
			LinkURL:   "https://medium.com",
			Position:  intPtr(PositionMedium),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Netflix",
			ImageURL:  "https://cdn-icons-png.freepik.com/512/2504/2504929.png",
			LinkURL:   "https://www.netflix.com",
			Position:  intPtr(PositionNetflix),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Amazon",
			ImageURL:  "https://cdn-icons-png.freepik.com/512/14063/14063250.png",
			LinkURL:   "https://www.amazon.com",
			Position:  intPtr(PositionAmazon),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "eBay",
			ImageURL:  "https://cdn-icons-png.freepik.com/512/14083/14083029.png",
			LinkURL:   "https://www.ebay.com",
			Position:  intPtr(PositionEbay),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Wikipedia",
			ImageURL:  "https://cdn-icons-png.freepik.com/512/14064/14064552.png",
			LinkURL:   "https://www.wikipedia.org",
			Position:  intPtr(PositionWikipedia),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Google",
			ImageURL:  "https://cdn-icons-png.freepik.com/512/14063/14063276.png",
			LinkURL:   "https://www.google.com",
			Position:  intPtr(PositionGoogle),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
		{
			Title:     "Yahoo",
			ImageURL:  "https://cdn-icons-png.freepik.com/512/2175/2175361.png",
			LinkURL:   "https://www.yahoo.com",
			Position:  intPtr(PositionYahoo),
			CreatedBy: "system",
			UpdatedBy: "system",
		},
	}

	if err := gormDB.Create(&seeds).Error; err != nil {
		return fmt.Errorf("failed to seed articles: %w", err)
	}
	return nil
}
