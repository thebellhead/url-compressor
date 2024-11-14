CREATE DATABASE IF NOT EXISTS url_compressor;

CREATE TABLE IF NOT EXISTS url_compressor.url_map (
    short_url String,
    long_url  String
) ENGINE = ReplacingMergeTree()
ORDER BY short_url;
