//
//  RegisterView.swift
//  idWallet
//
//  Created by Min Jae Lee on 2022-05-29.
//

import SwiftUI

struct RegisterView: View {
    @State var birthDate: String = ""
    @State var sin: String = ""
    
    let ns = NetworkService()
    
    var body: some View {
        VStack(alignment: .leading) {
            Text("You can have only one idWallet")
                .padding(.leading)
            
            Form {
                Section(header: Text("Your information")) {
                    TextField("Birthdate", text: $birthDate)
                    TextField("Sin", text: $sin)
                }
                
                Section {
                    Button(action: {
                        ns.getWallet(rq: WalletRequest.init(Birthdate: birthDate, Sin: sin))
                    }) {
                        Text("Register")
                    }
                }
            }
            .navigationTitle("Register")
        }
    }
}

struct RegisterView_Previews: PreviewProvider {
    static var previews: some View {
        RegisterView()
    }
}
