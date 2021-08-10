import 'package:flutter/foundation.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

const Config = {
  "HOLOVN_SERVER": "https://holovn.herokuapp.com"
};

String readEnvironment(String key){
  if (kReleaseMode){
    return Config[key]!;
  }
  return dotenv.env[key]!;
}