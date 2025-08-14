use base64::{engine::general_purpose, Engine as _};
use hex;

pub fn hex_to_base64(hex_str: &str) -> String {
    let bytes = hex::decode(hex_str).expect("Invalid hex");
    general_purpose::STANDARD.encode(bytes)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn set1_challenge1_hex_to_base64() {
        let input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d";
        let result = hex_to_base64(input);
        let expect = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t";
        assert_eq!(result, expect);
    }
}