use std::{io::Bytes, result};
use base64::{engine::general_purpose, Engine as _};
use hex;

pub fn hex_to_base64(hex_str: &str) -> String {
    let bytes = hex::decode(hex_str).expect("Invalid hex");
    general_purpose::STANDARD.encode(bytes)
}

pub fn fixed_xor(hexStr1: & str, hexStr2: & str) -> String {
    let byte1_result = hex::decode(hexStr1);
    assert!(byte1_result.is_ok(), "input string 1 is not hex.");

    let byte2_result = hex::decode(hexStr2);
    assert!(byte2_result.is_ok(), "input string 2 is not hex");

    let byte1 = byte1_result.unwrap();
    let byte2 = byte2_result.unwrap();

    let resulted_byte = xor_bytes_vector_with_different_length(byte1, byte2);
    return hex::encode(resulted_byte);
}

fn xor_bytes_vector_with_different_length(longerBytes: Vec<u8>, shorterBytes: Vec<u8>) -> Vec<u8> {
    if longerBytes.len() < shorterBytes.len(){
        return xor_bytes_vector_with_different_length(shorterBytes, longerBytes);
    }

    let length_difference = longerBytes.len() - shorterBytes.len();
    let mut result : Vec<u8> = Vec::new();
    for i in 0..longerBytes.len(){
        if i < length_difference {
            result.push(longerBytes[i] ^ 0);
        } else {
            result.push(longerBytes[i] ^ shorterBytes[i - length_difference]);
        }
    }
    return result;
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

    #[test]
    fn set1_challenge2_fixed_xor() {
        let input1 = "1c0111001f010100061a024b53535009181c";
        let input2 = "686974207468652062756c6c277320657965";
        let result = fixed_xor(input1, input2);
        let expect = "746865206b696420646f6e277420706c6179";
        assert_eq!(result, expect);
    }
}