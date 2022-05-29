//
//  ContentView.swift
//  idWalletIOS
//
//  Created by Min Jae Lee on 2022-05-28.
//

import SwiftUI

struct ContentView: View {
    @AppStorage(Constants.publicKey) private var publicKey: String = ""
    @State private var showPublicKey: Bool = false
    
    var body: some View {
        NavigationView {
            VStack(alignment: .center) {
                VStack(spacing: 0){
                    
                    HStack {
                        Text("idWallet")
                            .font(.largeTitle)
                            .bold()
                        
                        Spacer()
                        
                        NavigationLink(destination: WalletInfoView()) {
                            Image(systemName: "info.circle")
                                .foregroundColor(.white)
                                .font(.body)
                        }
                    }
                    .padding(.horizontal)
                    .padding(.bottom, 5)
                    
                    Divider()
                }
                .padding(.top, screenWidth * 0.15)
                .background(Color.blue)

                Spacer()
                
                if !publicKey.isEmpty {
                    Button(action: {
                        
                    }) {
                        RoundedRectangle(cornerRadius: 15)
                            .foregroundColor(.green)
                            .frame(width: screenWidth * 0.7, height: screenWidth * 0.7)
                            .overlay(Text("Verify").foregroundColor(.white))
                    }
                    
                    
                    Spacer()
                    
                    Button(action: {
                        UserDefaults.standard.set("", forKey: Constants.publicKey)
                    }) {
                        Text("REMOVE idWallet")
                            .foregroundColor(.red)
                            .padding()
                    }
                } else {
                    Text("You don't have credential yet\nPlease register it")
                    
                    NavigationLink(destination: RegisterView()) {
                        Circle()
                            .frame(width: screenWidth * 0.5, height: screenHeight * 0.5)
                            .foregroundColor(.blue)
                            .overlay(Text("REGISTER").foregroundColor(.white))
                    }
                }
                
                Spacer()
            }
            .navigationBarHidden(true)
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
