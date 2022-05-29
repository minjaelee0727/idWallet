//
//  WalletInfoView.swift
//  idWallet
//
//  Created by Min Jae Lee on 2022-05-29.
//

import SwiftUI

struct WalletInfoView: View {
    @AppStorage(Constants.publicKey) var publicKey: String = ""
    @AppStorage(Constants.secretKey) var secretKey: String = ""

    var body: some View {
        ScrollView() {
            VStack(alignment: .leading) {
                Text("PUBLIC KEY")
                    .bold()
                    .padding(.bottom, 3)
                
                Text(publicKey)
                    .truncationMode(.tail)
                    .padding(.bottom)
                
                Text("SECRET KEY")
                    .bold()
                    .padding(.bottom, 3)
                
                Text(secretKey)
                    .truncationMode(.tail)
            }
            .padding()
            
            Spacer()
        }
        .navigationBarTitleDisplayMode(.inline)
        .navigationTitle("idWallet Information")
    }
}

struct WalletInfoView_Previews: PreviewProvider {
    static var previews: some View {
        WalletInfoView()
    }
}
