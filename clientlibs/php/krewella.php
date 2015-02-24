<?php
// krewella_message takes in a network, channel and message string and returns
// a bool, string pair describing if it worked and why not if it didn't.
// This figures the host and key out based on a define() or an envvar for you.
function krewella_message($network, $channel, $message) {
	if(!defined(KREWELLA_HOST)) {
		if(!getenv('KREWELLA_HOST')) {
			return array(FALSE, 'No host defined');
		} else {
			$host = getenv('KREWELLA_HOST');
		}
	} else {
		$host = constant('KREWELLA_HOST');
	}

	if(!defined(KREWELLA_KEY)) {
		if(!getenv('KREWELLA_KEY')) {
			return array(FALSE, 'No key defined');
		} else {
			$key = getenv('KREWELLA_KEY');
		}
	} else {
		$key = constant('KREWELLA_KEY');
	}

	$ch = curl_init($host . '/' . $network . '/' . $channel);

	curl_setopt($ch, CURLOPT_HTTPHEADER, array(
		"X-Krewella-Auth: $key",
		'Content-Type: application/json',
	));
	curl_setopt($ch, CURLOPT_POSTFIELDS, $message);
	curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
	curl_setopt($ch, CURLOPT_POST, TRUE);

	$reply = curl_exec($ch);
	$code = curl_getinfo($ch, CURLINFO_HTTP_CODE);

	curl_close($ch);

	return array($code == 200, $reply);
}
