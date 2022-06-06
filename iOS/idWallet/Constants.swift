//
//  Constants.swift
//  idWallet
//
//  Created by Min Jae Lee on 2022-05-29.
//

import UIKit

public var screenWidth: CGFloat {
    return UIScreen.main.bounds.width
}

public var screenHeight: CGFloat {
    return UIScreen.main.bounds.height
}

enum Constants {
    static let registerURL = URL(string: "http://0.0.0.0:8080/register")
    static let publicKey: String = "k_publicKey"
    static let secretKey: String = "k_secretKey"
}

struct IdWallet: Codable {
    let publicKey: String
    let secretKey: String
}

struct WalletRequest: Codable {
    let Birthdate: String
    let Sin: String
}
