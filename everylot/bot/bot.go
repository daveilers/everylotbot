package main

import "flag"

func main() {
	// screenName := flag.String( `screen_name`, `Twitter screen name (without @)`)
	// dB :=  flag.String(`database`, ``, `path to SQLite lots database`)
	iD := flag.String(`id`, "", `tweet the entry in the lots table with this id`)
	config := flag.String(`config`, "bots.yaml", ``)
	dryRun := flag.Bool(`dry-run`, false, ``)
	// verbose := flag.Bool(`verbose`,  false, ``)
	// quiet := flag.Bool(`quiet`,  true, ``)

	//  flag.String(`search-format`, ``,` format string use for searching Google`)
	//  flag.String(`print-format`, ``,` format string use for poster to Twitter`)

	flag.Parse()

	// logger.debug('everylot starting with %s, %s', args.screen_name, args.database)

	// el = EveryLot(args.database,
	// 			  logger=logger,
	// 			  print_format=args.print_format,
	// 			  search_format=args.search_format,
	// 			  id_=args.id)

	// if not el.lot:
	// 	logger.error('No lot found')
	// 	return

	// logger.debug('%s addresss: %s zip: %s', el.lot['id'], el.lot.get('address'), el.lot.get('zip'))
	// logger.debug('db location %s,%s', el.lot['lat'], el.lot['lon'])

	// // Get the streetview image and upload it
	// // ("sv.jpg" is a dummy value, since filename is a required parameter).
	// image = el.get_streetview_image(api.config['streetview'])
	// media = api.media_upload('sv.jpg', file=image)

	// // compose an update with all the good parameters
	// // including the media string.
	// update = el.compose(media.media_id_string)
	// logger.info(update['status'])

	// if not args.dry_run:
	// 	logger.debug("posting")
	// 	status = api.update_status(**update)
	// 	try:
	// 		el.mark_as_tweeted(status.id)
	// 	except AttributeError:
	// 		el.mark_as_tweeted('1')

}
