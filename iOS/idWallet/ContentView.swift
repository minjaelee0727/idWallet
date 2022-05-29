//
//  ContentView.swift
//  idWalletIOS
//
//  Created by Min Jae Lee on 2022-05-28.
//

import SwiftUI

public var screenWidth: CGFloat {
    return UIScreen.main.bounds.width
}

public var screenHeight: CGFloat {
    return UIScreen.main.bounds.height
}

struct Constants {
    static let wallet = "Wallet"
}

struct ContentView: View {
    @AppStorage("walletExists") private var walletExists: Bool = true
    @State private var showPublicKey: Bool = false
    var body: some View {
        NavigationView {
            VStack {
                if walletExists {
                    VStack(alignment: .leading) {
                        Text("PUBLIC KEY")
                            .bold()
                            .padding(.bottom, 3)
                        HStack {
                            if showPublicKey {
                                Text("aslkdfjl;askdfjl;kasdjfi39u09ualskdjf2u340a;lks ")
                                    .minimumScaleFactor(0.8)
                                    .lineLimit(1)
                            } else {
                                Text("----------------------------------------------")
                                    .foregroundColor(.secondary)
                                    .minimumScaleFactor(0.8)
                                    .lineLimit(1)
                            }
                            
                            Button(action: {
                                showPublicKey.toggle()
                            }) {
                                Image(systemName: "eye")
                            }
                        }
                        
                        Spacer()
                    }
                } else {
                    VStack {
                        Text("You don't have credential yet\nPlease register it")
                        
                        NavigationLink(destination: RegisterView()) {
                            Circle()
                                .frame(width: screenWidth * 0.5, height: screenHeight * 0.5)
                            .foregroundColor(.blue)
                            .overlay(Text("REGISTER").foregroundColor(.white))
                        }
                    }
                }
            }
            .padding()
            .navigationTitle("idWallet")
            .navigationBarColor(backgroundColor: .blue, tintColor: .white)
        }
    }
}

//https://medium.com/swlh/custom-navigationview-bar-in-swiftui-4b782eb68e94
extension View {
  func navigationBarColor(backgroundColor: UIColor, tintColor: UIColor) -> some View {
    self.modifier(NavigationBarColor(backgroundColor: backgroundColor, tintColor: tintColor))
  }
}

struct NavigationBarColor: ViewModifier {

  init(backgroundColor: UIColor, tintColor: UIColor) {
    let coloredAppearance = UINavigationBarAppearance()
    coloredAppearance.configureWithOpaqueBackground()
    coloredAppearance.backgroundColor = backgroundColor
    coloredAppearance.titleTextAttributes = [.foregroundColor: tintColor]
    coloredAppearance.largeTitleTextAttributes = [.foregroundColor: tintColor]
                   
    UINavigationBar.appearance().standardAppearance = coloredAppearance
    UINavigationBar.appearance().scrollEdgeAppearance = coloredAppearance
    UINavigationBar.appearance().compactAppearance = coloredAppearance
    UINavigationBar.appearance().tintColor = tintColor
  }

  func body(content: Content) -> some View {
    content
  }
}

struct RegisterView: View {
    @State var fullname: String = ""
    @State var age: String = ""

    
    var body: some View {
        VStack(alignment: .leading) {
            Text("You can have only one idWallet")
                .padding(.leading)
            
            Form {
                Section(header: Text("Your information")) {
                    TextField("Full name", text: $fullname)
                    TextField("Age", text: $age)
                }
                
                Section {
                    Button(action: {
                        
                    }) {
                        Text("Register")
                    }
                }
            }
            .navigationTitle("Register")
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
