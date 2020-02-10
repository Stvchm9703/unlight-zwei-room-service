use serde::Deserialize;
use std::env;
use std::fs::File;
use std::io::prelude::*;

#[derive(Debug, Deserialize)]
struct Feature {
    location: Location,
    name: String,
}

#[derive(Debug, Deserialize)]
struct Location {
    latitude: i32,
    longitude: i32,
}

#[allow(dead_code)]
pub fn load() -> Vec<crate::routeguide::Feature> {
    let args: Vec<_> = env::args().collect();

    assert_eq!(args.len(), 2, "unexpected arguments");

    let mut file = File::open(&args[1]).ok().expect("failed to open data file");
    let mut data = String::new();
    file.read_to_string(&mut data)
        .ok()
        .expect("failed to read data file");

    let decoded: Vec<Feature> = serde_json::from_str(&data).unwrap();

    decoded
        .into_iter()
        .map(|feature| crate::routeguide::Feature {
            name: feature.name,
            location: Some(crate::routeguide::Point {
                longitude: feature.location.longitude,
                latitude: feature.location.latitude,
            }),
        })
        .collect()
}
