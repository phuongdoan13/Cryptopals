use base64::{engine::general_purpose, Engine as _};
use hex;

pub fn hex_to_base64(hex_str: &str) -> String {
    let bytes = hex::decode(hex_str).expect("Invalid hex");
    general_purpose::STANDARD.encode(bytes)
}
