<?php
$strs = file_get_contents("./adjson2162620.log");
$arrs = array_filter(explode("\n", $strs));

$adArrs = [];
$diffs = [];
foreach ($arrs as $key => $value) {
	// code...
	$tmp = json_decode($value, true);
	$adArrs[] = $tmp;

	// 第0个元素，跳过.
	if ($key == 0) {
		continue;
	}

	// 进行diff.
	$diff = [
		"time" => "before: ".$adArrs[$key-1]["time"]."; after: ".$tmp["time"],
		"campaign_serving_level"=> [
			"old" => $adArrs[$key-1]["campaign_serving_level"],
			"new" => $tmp["campaign_serving_level"]
		],
		"user_serving_level"=> [
			"old" => $adArrs[$key-1]["user_serving_level"],
			"new" => $tmp["user_serving_level"]
		]
	];
	$diffs[] = $diff;
}

var_dump($diffs);


