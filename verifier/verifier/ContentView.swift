//
//  ContentView.swift
//  verifier
//
//  Created by Min Jae Lee on 2022-05-29.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        VStack {
            Text("This web includes Adult Content")
            Text("Verify your identity with your idWallet")
            
            Button(action: {
                let idWallet = URL(string: "idWallet://")!
                if UIApplication.shared.canOpenURL(idWallet)
                {
                    UIApplication.shared.open(idWallet, options: [:])
                } else {
                    print("XX")
                }
            }) {
                RoundedRectangle(cornerRadius: 10)
                    .frame(width: 200, height: 200)
                    .foregroundColor(.red)
                    .padding()
                    .overlay(Text("VERIFY").foregroundColor(.white))
            }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
