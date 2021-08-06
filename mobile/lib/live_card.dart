import 'package:flutter/material.dart';

class LiveCard extends StatelessWidget {
  final String title, url;

  LiveCard(this.title, this.url);

  Widget build(BuildContext context) {
    return Center(
      child: Card(
        child: Column(
          children: [
            new Expanded(child: Image.network(url, fit: BoxFit.fitHeight)),
            Align(
              alignment: Alignment.centerLeft,
              child: Text(title, textAlign: TextAlign.start),)
          ],
        ),
      ),
    );
  }
}
