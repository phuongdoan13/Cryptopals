use std::{io::Bytes, result};
use base64::{engine::general_purpose, Engine as _};
use hex;
use crate::consts::{ASCII_CHARACTER_FREQUENCY};
use hypors::chi_square::goodness_of_fit;


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

pub fn single_byte_xor_cipher(hexStr: &str) -> (String, u8) {
    let bytes_result = hex::decode(hexStr);
    assert!(bytes_result.is_ok(), "input string is not hex.");
    let input_bytes = bytes_result.unwrap();

    let mut highest_chi_stat: f64 = f64::NEG_INFINITY;;
    let mut most_probable_byte: u8 = 0 as u8;
    let mut final_decrypt_string: String = String::new();

    for i in 0..=255 {
        let candidate_byte = i as u8;

        // xor each byte of the input string with the candidate_byte, and
        // calculate the frequency of each after_xor_byte;
        let mut freq_of_input_bytes = vec![0f32; 256];
        let mut decrypt_string_byte= Vec::new();
        for b in input_bytes.iter(){
            let after_xor_byte: u8 = b ^ candidate_byte;
            decrypt_string_byte.push(after_xor_byte);
            freq_of_input_bytes[after_xor_byte as usize] += 1f32;
        }

        // normalise the frequency of each after_xor_byte;
        for i in 0..=255 {
            freq_of_input_bytes[i] /= 256f32;
        }

        // calculate of chi-squared goodness of fit
        // against the pre-defined distribution of ascii letter
        let alpha = 0.25; // magical number alpha
        let chi_gof = goodness_of_fit(freq_of_input_bytes.iter().copied(),
                                      ASCII_CHARACTER_FREQUENCY.iter().copied(),
                                      alpha).unwrap();

        if chi_gof.test_statistic > highest_chi_stat {
            most_probable_byte = candidate_byte;
            highest_chi_stat =  chi_gof.test_statistic;
            final_decrypt_string = match str::from_utf8(decrypt_string_byte.as_slice()) {
                Ok(v) => v.to_string(),
                Err(e) => panic!("Invalid UTF-8 sequence: {}", e),
            };
        }
    }

    (final_decrypt_string, most_probable_byte)
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

    #[test]
    fn set1_challenge3_single_byte_xor_cipher() {
        let input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736";
        let (decrypt_message, candidate_byte) = single_byte_xor_cipher(input);
        let (expected_decrypt_message, expected_byte) = single_byte_xor_cipher
            (input);
        assert_eq!(decrypt_message, expected_decrypt_message);
    }
}