package gopay

/*

// Java SDK 获取Root证书SN
public static String getRootCertSN(String rootCertContent){
   String rootCertSN = NULL;
   try {
	   X509Certificate[] x509Certificates = readPemCertChain(rootCertContent);
	   MessageDigest md = MessageDigest.getInstance("MD5");
	   for (X509Certificate c : x509Certificates) {
		   if(c.getSigAlgOID().startsWith("1.2.840.113549.1.1")){
			   md.update((c.getIssuerX500Principal().getName() + c.getSerialNumber()).getBytes());
			   String certSN = new BigInteger(1,md.digest()).toString(16);
			   certSN = fillMD5(certSN);
			   if(StringUtils.isEmpty(rootCertSN)){
				   rootCertSN = certSN;
			   }else {
				   rootCertSN = rootCertSN + "_" + certSN;
			   }
		   }

	   }
   }catch (Exception e){
	   AlipayLogger.logBizError(("err"));
   }
   return rootCertSN;
}

// Java SDK 获取证书
private X509Certificate getCert(String  certPath) throws AlipayApiException{
	InputStream inputStream = NULL;
	try {
		inputStream = new FileInputStream(certPath);
		CertificateFactory cf = CertificateFactory.getInstance("X.509");
		X509Certificate cert = (X509Certificate)cf.generateCertificate(inputStream);
		return cert;

	}catch (IOException e) {
		throw new AlipayApiException(e);
	}catch (CertificateException e){
		throw new AlipayApiException(e);
	}finally {
		try {
			if (inputStream != NULL) {
				inputStream.close();
			}
		}catch (IOException e) {
			throw new AlipayApiException(e);
		}
	}
}

// Java SDK 获取证书SN
private String getCertSN(X509Certificate cf) throws AlipayApiException{
	try {
		MessageDigest md = MessageDigest.getInstance("MD5");
		md.update((cf.getIssuerX500Principal().getName() + cf.getSerialNumber()).getBytes());
		String certSN = new BigInteger(1,md.digest()).toString(16);
		certSN = fillMD5(certSN);
		return certSN;
	}catch (NoSuchAlgorithmException e){
		throw new AlipayApiException(e);
	}
}

*/
